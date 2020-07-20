// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsPrivilege - Privilege info.
// Export NsPrivilegeFields for advance operations like search filter etc.
var NsPrivilegeFields *NsPrivilege

func init() {

	NsPrivilegeFields = &NsPrivilege{
		ObjectType: "object_type",
	}
}

type NsPrivilege struct {
	// ObjectType - Object type name associated with this privilege.
	ObjectType string `json:"object_type,omitempty"`
	// Operations - List of operations associated with the above object for this privilege.
	Operations []*string `json:"operations,omitempty"`
}

// sdk internal struct
type nsPrivilege struct {
	ObjectType *string   `json:"object_type,omitempty"`
	Operations []*string `json:"operations,omitempty"`
}

// EncodeNsPrivilege - Transform NsPrivilege to nsPrivilege type
func EncodeNsPrivilege(request interface{}) (*nsPrivilege, error) {
	reqNsPrivilege := request.(*NsPrivilege)
	byte, err := json.Marshal(reqNsPrivilege)
	if err != nil {
		return nil, err
	}
	respNsPrivilegePtr := &nsPrivilege{}
	err = json.Unmarshal(byte, respNsPrivilegePtr)
	return respNsPrivilegePtr, err
}

// DecodeNsPrivilege Transform nsPrivilege to NsPrivilege type
func DecodeNsPrivilege(request interface{}) (*NsPrivilege, error) {
	reqNsPrivilege := request.(*nsPrivilege)
	byte, err := json.Marshal(reqNsPrivilege)
	if err != nil {
		return nil, err
	}
	respNsPrivilege := &NsPrivilege{}
	err = json.Unmarshal(byte, respNsPrivilege)
	return respNsPrivilege, err
}
