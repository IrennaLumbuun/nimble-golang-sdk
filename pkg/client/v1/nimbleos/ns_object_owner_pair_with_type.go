// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsObjectOwnerPairWithType - List of objects of a given type along with their owners.
// Export NsObjectOwnerPairWithTypeFields for advance operations like search filter etc.
var NsObjectOwnerPairWithTypeFields *NsObjectOwnerPairWithType

func init() {

	NsObjectOwnerPairWithTypeFields = &NsObjectOwnerPairWithType{}
}

type NsObjectOwnerPairWithType struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ObjOwnerPairList - List of object names and owners.
	ObjOwnerPairList []*NsObjectOwnerPair `json:"obj_owner_pair_list,omitempty"`
}

// sdk internal struct
type nsObjectOwnerPairWithType struct {
	ObjType          *NsObjectType        `json:"obj_type,omitempty"`
	ObjOwnerPairList []*NsObjectOwnerPair `json:"obj_owner_pair_list,omitempty"`
}

// EncodeNsObjectOwnerPairWithType - Transform NsObjectOwnerPairWithType to nsObjectOwnerPairWithType type
func EncodeNsObjectOwnerPairWithType(request interface{}) (*nsObjectOwnerPairWithType, error) {
	reqNsObjectOwnerPairWithType := request.(*NsObjectOwnerPairWithType)
	byte, err := json.Marshal(reqNsObjectOwnerPairWithType)
	if err != nil {
		return nil, err
	}
	respNsObjectOwnerPairWithTypePtr := &nsObjectOwnerPairWithType{}
	err = json.Unmarshal(byte, respNsObjectOwnerPairWithTypePtr)
	return respNsObjectOwnerPairWithTypePtr, err
}

// DecodeNsObjectOwnerPairWithType Transform nsObjectOwnerPairWithType to NsObjectOwnerPairWithType type
func DecodeNsObjectOwnerPairWithType(request interface{}) (*NsObjectOwnerPairWithType, error) {
	reqNsObjectOwnerPairWithType := request.(*nsObjectOwnerPairWithType)
	byte, err := json.Marshal(reqNsObjectOwnerPairWithType)
	if err != nil {
		return nil, err
	}
	respNsObjectOwnerPairWithType := &NsObjectOwnerPairWithType{}
	err = json.Unmarshal(byte, respNsObjectOwnerPairWithType)
	return respNsObjectOwnerPairWithType, err
}
