// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Fibre Channel session is created when Fibre Channel initiator connects to this group.
const (
    fibreChannelSessionPath = "fibre_channel_sessions"
)

// FibreChannelSessionObjectSet
type FibreChannelSessionObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new FibreChannelSession object
func (objectSet *FibreChannelSessionObjectSet) CreateObject(payload *model.FibreChannelSession) (*model.FibreChannelSession, error) {
	fibreChannelSessionObjectSetResp, err := objectSet.Client.Post(fibreChannelSessionPath, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if fibreChannelSessionObjectSetResp == nil {
		return nil,nil
	}
	return fibreChannelSessionObjectSetResp.(*model.FibreChannelSession), err
}

// UpdateObject Modify existing FibreChannelSession object
func (objectSet *FibreChannelSessionObjectSet) UpdateObject(id string, payload *model.FibreChannelSession) (*model.FibreChannelSession, error) {
	fibreChannelSessionObjectSetResp, err := objectSet.Client.Put(fibreChannelSessionPath, id, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if fibreChannelSessionObjectSetResp == nil {
		return nil,nil
	}
	return fibreChannelSessionObjectSetResp.(*model.FibreChannelSession), err
}

// DeleteObject deletes the FibreChannelSession object with the specified ID
func (objectSet *FibreChannelSessionObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(fibreChannelSessionPath, id)
	if err !=nil {
		return err
	}
	return nil
}

// GetObject returns a FibreChannelSession object with the given ID
func (objectSet *FibreChannelSessionObjectSet) GetObject(id string) (*model.FibreChannelSession, error) {
	fibreChannelSessionObjectSetResp, err := objectSet.Client.Get(fibreChannelSessionPath, id, model.FibreChannelSession{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if fibreChannelSessionObjectSetResp == nil {
		return nil,nil
	}
	return fibreChannelSessionObjectSetResp.(*model.FibreChannelSession), err
}

// GetObjectList returns the list of FibreChannelSession objects
func (objectSet *FibreChannelSessionObjectSet) GetObjectList() ([]*model.FibreChannelSession, error) {
	fibreChannelSessionObjectSetResp, err := objectSet.Client.List(fibreChannelSessionPath)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelSessionObjectSet(fibreChannelSessionObjectSetResp), err
}

// GetObjectListFromParams returns the list of FibreChannelSession objects using the given params query info
func (objectSet *FibreChannelSessionObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.FibreChannelSession, error) {
	fibreChannelSessionObjectSetResp, err := objectSet.Client.ListFromParams(fibreChannelSessionPath, params)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelSessionObjectSet(fibreChannelSessionObjectSetResp), err
}

// generated function to build the appropriate response types
func buildFibreChannelSessionObjectSet(response interface{}) ([]*model.FibreChannelSession) {
	values := reflect.ValueOf(response)
	results := make([]*model.FibreChannelSession, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.FibreChannelSession{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}