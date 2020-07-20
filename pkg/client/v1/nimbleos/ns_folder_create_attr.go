// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsFolderCreateAttr - Attributes for folder creation.
// Export NsFolderCreateAttrFields for advance operations like search filter etc.
var NsFolderCreateAttrFields *NsFolderCreateAttr

func init() {

	NsFolderCreateAttrFields = &NsFolderCreateAttr{
		Name:   "name",
		PoolId: "pool_id",
	}
}

type NsFolderCreateAttr struct {
	// Name - Name of folder.
	Name string `json:"name,omitempty"`
	// PoolId - ID of pool to create the folder in.
	PoolId string `json:"pool_id,omitempty"`
}

// sdk internal struct
type nsFolderCreateAttr struct {
	Name   *string `json:"name,omitempty"`
	PoolId *string `json:"pool_id,omitempty"`
}

// EncodeNsFolderCreateAttr - Transform NsFolderCreateAttr to nsFolderCreateAttr type
func EncodeNsFolderCreateAttr(request interface{}) (*nsFolderCreateAttr, error) {
	reqNsFolderCreateAttr := request.(*NsFolderCreateAttr)
	byte, err := json.Marshal(reqNsFolderCreateAttr)
	if err != nil {
		return nil, err
	}
	respNsFolderCreateAttrPtr := &nsFolderCreateAttr{}
	err = json.Unmarshal(byte, respNsFolderCreateAttrPtr)
	return respNsFolderCreateAttrPtr, err
}

// DecodeNsFolderCreateAttr Transform nsFolderCreateAttr to NsFolderCreateAttr type
func DecodeNsFolderCreateAttr(request interface{}) (*NsFolderCreateAttr, error) {
	reqNsFolderCreateAttr := request.(*nsFolderCreateAttr)
	byte, err := json.Marshal(reqNsFolderCreateAttr)
	if err != nil {
		return nil, err
	}
	respNsFolderCreateAttr := &NsFolderCreateAttr{}
	err = json.Unmarshal(byte, respNsFolderCreateAttr)
	return respNsFolderCreateAttr, err
}
