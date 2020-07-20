// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsArrayNet - Array network config.
// Export NsArrayNetFields for advance operations like search filter etc.
var NsArrayNetFields *NsArrayNet

func init() {

	NsArrayNetFields = &NsArrayNet{
		Name:            "name",
		CtrlrASupportIp: "ctrlr_a_support_ip",
		CtrlrBSupportIp: "ctrlr_b_support_ip",
	}
}

type NsArrayNet struct {
	// Name - Name of the array.
	Name string `json:"name,omitempty"`
	// MemberGid - GID of member. This field cannot be updated.
	MemberGid int64 `json:"member_gid,omitempty"`
	// CtrlrASupportIp - IP address of controller A.
	CtrlrASupportIp string `json:"ctrlr_a_support_ip,omitempty"`
	// CtrlrBSupportIp - IP address of controller B.
	CtrlrBSupportIp string `json:"ctrlr_b_support_ip,omitempty"`
	// NicList - List of NICs.
	NicList []*NsNIC `json:"nic_list,omitempty"`
}

// sdk internal struct
type nsArrayNet struct {
	Name            *string  `json:"name,omitempty"`
	MemberGid       *int64   `json:"member_gid,omitempty"`
	CtrlrASupportIp *string  `json:"ctrlr_a_support_ip,omitempty"`
	CtrlrBSupportIp *string  `json:"ctrlr_b_support_ip,omitempty"`
	NicList         []*NsNIC `json:"nic_list,omitempty"`
}

// EncodeNsArrayNet - Transform NsArrayNet to nsArrayNet type
func EncodeNsArrayNet(request interface{}) (*nsArrayNet, error) {
	reqNsArrayNet := request.(*NsArrayNet)
	byte, err := json.Marshal(reqNsArrayNet)
	if err != nil {
		return nil, err
	}
	respNsArrayNetPtr := &nsArrayNet{}
	err = json.Unmarshal(byte, respNsArrayNetPtr)
	return respNsArrayNetPtr, err
}

// DecodeNsArrayNet Transform nsArrayNet to NsArrayNet type
func DecodeNsArrayNet(request interface{}) (*NsArrayNet, error) {
	reqNsArrayNet := request.(*nsArrayNet)
	byte, err := json.Marshal(reqNsArrayNet)
	if err != nil {
		return nil, err
	}
	respNsArrayNet := &NsArrayNet{}
	err = json.Unmarshal(byte, respNsArrayNet)
	return respNsArrayNet, err
}
