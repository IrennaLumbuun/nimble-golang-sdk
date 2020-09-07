// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Shelf Service - Disk shelf and head unit houses disks and controller.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// ShelfService type
type ShelfService struct {
	objectSet *client.ShelfObjectSet
}

// NewShelfService - method to initialize "ShelfService"
func NewShelfService(gs *NsGroupService) *ShelfService {
	objectSet := gs.client.GetShelfObjectSet()
	return &ShelfService{objectSet: objectSet}
}

// GetShelfs - method returns a array of pointers of type "Shelfs"
func (svc *ShelfService) GetShelfs(params *param.GetParams) ([]*nimbleos.Shelf, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	shelfResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return shelfResp, nil
}

// CreateShelf - method creates a "Shelf"
func (svc *ShelfService) CreateShelf(obj *nimbleos.Shelf) (*nimbleos.Shelf, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	shelfResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return shelfResp, nil
}

// UpdateShelf - method modifies  the "Shelf"
func (svc *ShelfService) UpdateShelf(id string, obj *nimbleos.Shelf) (*nimbleos.Shelf, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	shelfResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return shelfResp, nil
}

// GetShelfById - method returns a pointer to "Shelf"
func (svc *ShelfService) GetShelfById(id string) (*nimbleos.Shelf, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	shelfResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return shelfResp, nil
}

// DeleteShelf - deletes the "Shelf"
func (svc *ShelfService) DeleteShelf(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}

// IdentifyShelf - turn on chassis identifier for a controller.
//   Required parameters:
//       id - ID of shelf.
//       cid - Possible values: 'A', 'B'.
//       status - Status value of identifier to set.

func (svc *ShelfService) IdentifyShelf(id string, cid *nimbleos.NsControllerId, status bool) (*nimbleos.NsShelfIdentifyStatusReturn, error) {

	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified id:%v, cid:%v, status:%v ", id, cid, status)
	}

	resp, err := svc.objectSet.Identify(&id, cid, &status)
	return resp, err

}

// EvacuateShelf - perform shelf evacuation.
//   Required parameters:
//       id - ID of shelf.
//       driveset - Driveset to evacuate.

//   Optional parameters:
//       dryRun - Argument to perform a dry run, not the actual shelf evacuation.
//       start - Argument to perform the shelf evacuation.
//       cancel - Argument to cancel the shelf evacuation.
//       pause - Argument to pause the shelf evacuation.
//       resume - Argument to resume the shelf evacuation.

func (svc *ShelfService) EvacuateShelf(id string, driveset uint64, dryRun *bool, start *bool, cancel *bool, pause *bool, resume *bool) error {

	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified id:%v, driveset:%v, dryRun:%v, start:%v, cancel:%v, pause:%v, resume:%v ", id, driveset, dryRun, start, cancel, pause, resume)
	}

	err := svc.objectSet.Evacuate(&id, &driveset, dryRun, start, cancel, pause, resume)
	return err
}
