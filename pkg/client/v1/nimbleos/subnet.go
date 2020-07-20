// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// Subnet - Search subnets information. Many networking tasks require that objects such as replication partners are either on the same network or have a route to a secondary network. Subnets let you create logical addressing for selective routing.
// Export SubnetFields for advance operations like search filter etc.
var SubnetFields *Subnet

func init() {

	SubnetFields = &Subnet{
		ID:          "id",
		Name:        "name",
		Network:     "network",
		Netmask:     "netmask",
		DiscoveryIp: "discovery_ip",
	}
}

type Subnet struct {
	// ID - Identifier for the initiator group.
	ID string `json:"id,omitempty"`
	// Name - Name of subnet configuration.
	Name string `json:"name,omitempty"`
	// Network - Subnet network address.
	Network string `json:"network,omitempty"`
	// Netmask - Subnet netmask address.
	Netmask string `json:"netmask,omitempty"`
	// Type - Subnet type. Options include 'mgmt', 'data', and 'mgmt,data'.
	Type *NsSubnetType `json:"type,omitempty"`
	// AllowIscsi - Subnet type.
	AllowIscsi *bool `json:"allow_iscsi,omitempty"`
	// AllowGroup - Subnet type.
	AllowGroup *bool `json:"allow_group,omitempty"`
	// DiscoveryIp - Subnet network address.
	DiscoveryIp string `json:"discovery_ip,omitempty"`
	// Mtu - MTU for specified subnet. Valid MTU's are in the 512-16000 range.
	Mtu int64 `json:"mtu,omitempty"`
	// NetzoneType - Specify Network Affinity Zone type for iSCSI enabled subnets. Valid types are Single, Bisect, and EvenOdd for iSCSI subnets.
	NetzoneType *NsNetZoneType `json:"netzone_type,omitempty"`
	// VlanId - VLAN ID for specified subnet. Valid ID's are in the 1-4094 range.
	VlanId int64 `json:"vlan_id,omitempty"`
	// CreationTime - Time when this subnet configuration was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this subnet configuration was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
	// Failover - Failover setting of the subnet.
	Failover *bool `json:"failover,omitempty"`
	// FailoverEnableTime - Failover for this subnet will be enabled again at the time specified by failover_enable_time.
	FailoverEnableTime int64 `json:"failover_enable_time,omitempty"`
}

// sdk internal struct
type subnet struct {
	ID                 *string        `json:"id,omitempty"`
	Name               *string        `json:"name,omitempty"`
	Network            *string        `json:"network,omitempty"`
	Netmask            *string        `json:"netmask,omitempty"`
	Type               *NsSubnetType  `json:"type,omitempty"`
	AllowIscsi         *bool          `json:"allow_iscsi,omitempty"`
	AllowGroup         *bool          `json:"allow_group,omitempty"`
	DiscoveryIp        *string        `json:"discovery_ip,omitempty"`
	Mtu                *int64         `json:"mtu,omitempty"`
	NetzoneType        *NsNetZoneType `json:"netzone_type,omitempty"`
	VlanId             *int64         `json:"vlan_id,omitempty"`
	CreationTime       *int64         `json:"creation_time,omitempty"`
	LastModified       *int64         `json:"last_modified,omitempty"`
	Failover           *bool          `json:"failover,omitempty"`
	FailoverEnableTime *int64         `json:"failover_enable_time,omitempty"`
}

// EncodeSubnet - Transform Subnet to subnet type
func EncodeSubnet(request interface{}) (*subnet, error) {
	reqSubnet := request.(*Subnet)
	byte, err := json.Marshal(reqSubnet)
	if err != nil {
		return nil, err
	}
	respSubnetPtr := &subnet{}
	err = json.Unmarshal(byte, respSubnetPtr)
	return respSubnetPtr, err
}

// DecodeSubnet Transform subnet to Subnet type
func DecodeSubnet(request interface{}) (*Subnet, error) {
	reqSubnet := request.(*subnet)
	byte, err := json.Marshal(reqSubnet)
	if err != nil {
		return nil, err
	}
	respSubnet := &Subnet{}
	err = json.Unmarshal(byte, respSubnet)
	return respSubnet, err
}
