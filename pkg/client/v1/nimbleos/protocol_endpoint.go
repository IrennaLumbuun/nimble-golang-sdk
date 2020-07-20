// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// ProtocolEndpoint - Protocol endpoints are administrative logical units (LUs) in an LU conglomerate to be used with VMware Virtual Volumes.
// Export ProtocolEndpointFields for advance operations like search filter etc.
var ProtocolEndpointFields *ProtocolEndpoint

func init() {

	ProtocolEndpointFields = &ProtocolEndpoint{
		ID:           "id",
		Name:         "name",
		Description:  "description",
		PoolName:     "pool_name",
		PoolId:       "pool_id",
		SerialNumber: "serial_number",
		TargetName:   "target_name",
	}
}

type ProtocolEndpoint struct {
	// ID - Identifier for the protocol endpoint.
	ID string `json:"id,omitempty"`
	// Name - Name of the protocol endpoint.
	Name string `json:"name,omitempty"`
	// Description - Text description of the protocol endpoint.
	Description string `json:"description,omitempty"`
	// PoolName - Name of the pool where the protocol endpoint resides. If pool option is not specified, protocol endpoint is assigned to the default pool.
	PoolName string `json:"pool_name,omitempty"`
	// PoolId - Identifier associated with the pool in the storage pool table.
	PoolId string `json:"pool_id,omitempty"`
	// State - Operational state of protocol endpoint.
	State *NsPEOpStateType `json:"state,omitempty"`
	// SerialNumber - Identifier associated with the protocol endpoint for the SCSI protocol.
	SerialNumber string `json:"serial_number,omitempty"`
	// TargetName - The iSCSI Qualified Name (IQN) or the Fibre Channel World Wide Node Name (WWNN) of the target protocol endpoint.
	TargetName string `json:"target_name,omitempty"`
	// GroupSpecificIds - External UID is used to compute the serial number and IQN which never change even if the running group changes (e.g. after group merge). Group-specific IDs determine whether external UID is used for computing serial number and IQN.
	GroupSpecificIds *bool `json:"group_specific_ids,omitempty"`
	// CreationTime - Time when this protocol endpoint was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this protocol endpoint was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
	// NumConnections - Number of connections via this protocol endpoint.
	NumConnections int64 `json:"num_connections,omitempty"`
	// NumIscsiConnections - Number of iSCSI connections via this protocol endpoint.
	NumIscsiConnections int64 `json:"num_iscsi_connections,omitempty"`
	// NumFcConnections - Number of FC connections via this protocol endpoint.
	NumFcConnections int64 `json:"num_fc_connections,omitempty"`
	// AccessControlRecords - List of access control records that apply to this protocol endpoint.
	AccessControlRecords []*NsAccessControlRecord `json:"access_control_records,omitempty"`
	// IscsiSessions - List of iSCSI sessions connected to this protocol endpoint.
	IscsiSessions []*NsISCSISession `json:"iscsi_sessions,omitempty"`
	// FcSessions - List of FC sessions connected to this protocol endpoint.
	FcSessions []*NsFCSession `json:"fc_sessions,omitempty"`
	// AccessProtocol - Access protocol of the protocol endpoint. Only initiator groups with the same access protocol can access the protocol endpoint. If not specified in the creation request, it will be the access protocol supported by the group. If the group supports multiple protocols, the default will be Fibre Channel.
	AccessProtocol *NsAccessProtocol `json:"access_protocol,omitempty"`
}

// sdk internal struct
type protocolEndpoint struct {
	ID                   *string                  `json:"id,omitempty"`
	Name                 *string                  `json:"name,omitempty"`
	Description          *string                  `json:"description,omitempty"`
	PoolName             *string                  `json:"pool_name,omitempty"`
	PoolId               *string                  `json:"pool_id,omitempty"`
	State                *NsPEOpStateType         `json:"state,omitempty"`
	SerialNumber         *string                  `json:"serial_number,omitempty"`
	TargetName           *string                  `json:"target_name,omitempty"`
	GroupSpecificIds     *bool                    `json:"group_specific_ids,omitempty"`
	CreationTime         *int64                   `json:"creation_time,omitempty"`
	LastModified         *int64                   `json:"last_modified,omitempty"`
	NumConnections       *int64                   `json:"num_connections,omitempty"`
	NumIscsiConnections  *int64                   `json:"num_iscsi_connections,omitempty"`
	NumFcConnections     *int64                   `json:"num_fc_connections,omitempty"`
	AccessControlRecords []*NsAccessControlRecord `json:"access_control_records,omitempty"`
	IscsiSessions        []*NsISCSISession        `json:"iscsi_sessions,omitempty"`
	FcSessions           []*NsFCSession           `json:"fc_sessions,omitempty"`
	AccessProtocol       *NsAccessProtocol        `json:"access_protocol,omitempty"`
}

// EncodeProtocolEndpoint - Transform ProtocolEndpoint to protocolEndpoint type
func EncodeProtocolEndpoint(request interface{}) (*protocolEndpoint, error) {
	reqProtocolEndpoint := request.(*ProtocolEndpoint)
	byte, err := json.Marshal(reqProtocolEndpoint)
	if err != nil {
		return nil, err
	}
	respProtocolEndpointPtr := &protocolEndpoint{}
	err = json.Unmarshal(byte, respProtocolEndpointPtr)
	return respProtocolEndpointPtr, err
}

// DecodeProtocolEndpoint Transform protocolEndpoint to ProtocolEndpoint type
func DecodeProtocolEndpoint(request interface{}) (*ProtocolEndpoint, error) {
	reqProtocolEndpoint := request.(*protocolEndpoint)
	byte, err := json.Marshal(reqProtocolEndpoint)
	if err != nil {
		return nil, err
	}
	respProtocolEndpoint := &ProtocolEndpoint{}
	err = json.Unmarshal(byte, respProtocolEndpoint)
	return respProtocolEndpoint, err
}
