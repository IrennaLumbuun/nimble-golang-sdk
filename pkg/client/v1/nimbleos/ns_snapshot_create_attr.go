// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsSnapshotCreateAttr - Select fields containing volume info.
// Export NsSnapshotCreateAttrFields for advance operations like search filter etc.
var NsSnapshotCreateAttrFields *NsSnapshotCreateAttr

func init() {

	NsSnapshotCreateAttrFields = &NsSnapshotCreateAttr{
		VolId:       "vol_id",
		Name:        "name",
		Description: "description",
	}
}

type NsSnapshotCreateAttr struct {
	// VolId - ID of volume.
	VolId string `json:"vol_id,omitempty"`
	// Name - Snapshot name.
	Name string `json:"name,omitempty"`
	// Description - Snapshot description.
	Description string `json:"description,omitempty"`
	// Online - Snapshot is online.
	Online *bool `json:"online,omitempty"`
	// Writable - Snapshot is writable.
	Writable *bool `json:"writable,omitempty"`
	// AgentType - External management agent type.
	AgentType *NsAgentType `json:"agent_type,omitempty"`
	// Metadata - Key-value pairs that augment a snapshot's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
}

// sdk internal struct
type nsSnapshotCreateAttr struct {
	VolId       *string       `json:"vol_id,omitempty"`
	Name        *string       `json:"name,omitempty"`
	Description *string       `json:"description,omitempty"`
	Online      *bool         `json:"online,omitempty"`
	Writable    *bool         `json:"writable,omitempty"`
	AgentType   *NsAgentType  `json:"agent_type,omitempty"`
	Metadata    []*NsKeyValue `json:"metadata,omitempty"`
}

// EncodeNsSnapshotCreateAttr - Transform NsSnapshotCreateAttr to nsSnapshotCreateAttr type
func EncodeNsSnapshotCreateAttr(request interface{}) (*nsSnapshotCreateAttr, error) {
	reqNsSnapshotCreateAttr := request.(*NsSnapshotCreateAttr)
	byte, err := json.Marshal(reqNsSnapshotCreateAttr)
	if err != nil {
		return nil, err
	}
	respNsSnapshotCreateAttrPtr := &nsSnapshotCreateAttr{}
	err = json.Unmarshal(byte, respNsSnapshotCreateAttrPtr)
	return respNsSnapshotCreateAttrPtr, err
}

// DecodeNsSnapshotCreateAttr Transform nsSnapshotCreateAttr to NsSnapshotCreateAttr type
func DecodeNsSnapshotCreateAttr(request interface{}) (*NsSnapshotCreateAttr, error) {
	reqNsSnapshotCreateAttr := request.(*nsSnapshotCreateAttr)
	byte, err := json.Marshal(reqNsSnapshotCreateAttr)
	if err != nil {
		return nil, err
	}
	respNsSnapshotCreateAttr := &NsSnapshotCreateAttr{}
	err = json.Unmarshal(byte, respNsSnapshotCreateAttr)
	return respNsSnapshotCreateAttr, err
}
