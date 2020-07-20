// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsGroupMergeReturn - Response of group merge validation.
// Export NsGroupMergeReturnFields for advance operations like search filter etc.
var NsGroupMergeReturnFields *NsGroupMergeReturn

func init() {

	NsGroupMergeReturnFields = &NsGroupMergeReturn{
		SrcSid:             "src_sid",
		SrcGroupName:       "src_group_name",
		DstGroupName:       "dst_group_name",
		DstGroupSwversion:  "dst_group_swversion",
		ValidationErrorMsg: "validation_error_msg",
	}
}

type NsGroupMergeReturn struct {
	// OnlineVols - List of volumes which are online.
	OnlineVols []*string `json:"online_vols,omitempty"`
	// OnlineSnaps - List of snapshots which are online.
	OnlineSnaps []*NsVolAndSnapName `json:"online_snaps,omitempty"`
	// ActivePartners - List of partners on groupB which are not paused/stopped.
	ActivePartners []*string `json:"active_partners,omitempty"`
	// DstNatPartners - List of partners on groupA that think we are behind a NAT.
	DstNatPartners []*string `json:"dst_nat_partners,omitempty"`
	// SrcThrottles - List of partner throttles on groupB.
	SrcThrottles []*string `json:"src_throttles,omitempty"`
	// DstThrottles - List of partner throttles on groupA.
	DstThrottles []*string `json:"dst_throttles,omitempty"`
	// ReplObjs - List of volumes which have the same UID (replicated).
	ReplObjs []*NsReplPairListWithObjectType `json:"repl_objs,omitempty"`
	// NameConflicts - List of objects with conflicting names.
	NameConflicts []*NsObjectNameListWithType `json:"name_conflicts,omitempty"`
	// NameConflictsManualResolve - List of objects with conflicting names. These object need to be resolved manually.
	NameConflictsManualResolve []*NsObjectNameListWithType `json:"name_conflicts_manual_resolve,omitempty"`
	// SerialConflicts - List of objects with conflicting serial numbers.
	SerialConflicts []*NsObjectNameListWithType `json:"serial_conflicts,omitempty"`
	// AppUuidConflicts - List of objects with conflicting app_uuid.
	AppUuidConflicts []*NsObjectNameListWithType `json:"app_uuid_conflicts,omitempty"`
	// NameConflictsAndOwners - List of objects with conflicting names and their owners.
	NameConflictsAndOwners []*NsObjectOwnerPairWithType `json:"name_conflicts_and_owners,omitempty"`
	// LimitViolations - List of object types whose total number limit would be violated.
	LimitViolations []*NsObjectLimit `json:"limit_violations,omitempty"`
	// SnapRetainLimitViolations - List of snapshot retainment param limit violations.
	SnapRetainLimitViolations []*NsSnapRetainLimit `json:"snap_retain_limit_violations,omitempty"`
	// NetworkErrorList - List of validation errors with details.
	NetworkErrorList []*NsErrorWithArguments `json:"network_error_list,omitempty"`
	// AutoSwitchoverConflicts - List of validation errors for automatic switchover of Group Management.
	AutoSwitchoverConflicts []*NsErrorWithArguments `json:"auto_switchover_conflicts,omitempty"`
	// ErrorList - List of validation errors.
	ErrorList []*NsErrorWithArguments `json:"error_list,omitempty"`
	// AliasConflicts - List of WWPN/alias conflicts (same initiator WWPN, but different aliases).
	AliasConflicts []*NsAliasConflictPair `json:"alias_conflicts,omitempty"`
	// OnlineFcIntfs - List of Fibre Channel interfaces which are online.
	OnlineFcIntfs []*NsFibreChannelInterfaceFullName `json:"online_fc_intfs,omitempty"`
	// LunConflicts - List of LUN conflicts.
	LunConflicts []*NsLunConflictPair `json:"lun_conflicts,omitempty"`
	// SrcSid - Session ID for source group.
	SrcSid string `json:"src_sid,omitempty"`
	// SrcGroupName - Name of the group we are merging with.
	SrcGroupName string `json:"src_group_name,omitempty"`
	// DstGroupName - Name of the destination group.
	DstGroupName string `json:"dst_group_name,omitempty"`
	// DstGroupSwversion - Software version of the destination group.
	DstGroupSwversion string `json:"dst_group_swversion,omitempty"`
	// HtypeNameConflicts - List host type with different bit mask.
	HtypeNameConflicts []*string `json:"htype_name_conflicts,omitempty"`
	// HostTypeConflicts - List of initiators and igroups with different host types.
	HostTypeConflicts []*NsHostType `json:"host_type_conflicts,omitempty"`
	// ValidationError - Umbrella error code for group merge validation.
	ValidationError []*NsErrorWithArguments `json:"validation_error,omitempty"`
	// ValidationErrorMsg - Detailed error message.
	ValidationErrorMsg string `json:"validation_error_msg,omitempty"`
}

// sdk internal struct
type nsGroupMergeReturn struct {
	OnlineVols                 []*string                          `json:"online_vols,omitempty"`
	OnlineSnaps                []*NsVolAndSnapName                `json:"online_snaps,omitempty"`
	ActivePartners             []*string                          `json:"active_partners,omitempty"`
	DstNatPartners             []*string                          `json:"dst_nat_partners,omitempty"`
	SrcThrottles               []*string                          `json:"src_throttles,omitempty"`
	DstThrottles               []*string                          `json:"dst_throttles,omitempty"`
	ReplObjs                   []*NsReplPairListWithObjectType    `json:"repl_objs,omitempty"`
	NameConflicts              []*NsObjectNameListWithType        `json:"name_conflicts,omitempty"`
	NameConflictsManualResolve []*NsObjectNameListWithType        `json:"name_conflicts_manual_resolve,omitempty"`
	SerialConflicts            []*NsObjectNameListWithType        `json:"serial_conflicts,omitempty"`
	AppUuidConflicts           []*NsObjectNameListWithType        `json:"app_uuid_conflicts,omitempty"`
	NameConflictsAndOwners     []*NsObjectOwnerPairWithType       `json:"name_conflicts_and_owners,omitempty"`
	LimitViolations            []*NsObjectLimit                   `json:"limit_violations,omitempty"`
	SnapRetainLimitViolations  []*NsSnapRetainLimit               `json:"snap_retain_limit_violations,omitempty"`
	NetworkErrorList           []*NsErrorWithArguments            `json:"network_error_list,omitempty"`
	AutoSwitchoverConflicts    []*NsErrorWithArguments            `json:"auto_switchover_conflicts,omitempty"`
	ErrorList                  []*NsErrorWithArguments            `json:"error_list,omitempty"`
	AliasConflicts             []*NsAliasConflictPair             `json:"alias_conflicts,omitempty"`
	OnlineFcIntfs              []*NsFibreChannelInterfaceFullName `json:"online_fc_intfs,omitempty"`
	LunConflicts               []*NsLunConflictPair               `json:"lun_conflicts,omitempty"`
	SrcSid                     *string                            `json:"src_sid,omitempty"`
	SrcGroupName               *string                            `json:"src_group_name,omitempty"`
	DstGroupName               *string                            `json:"dst_group_name,omitempty"`
	DstGroupSwversion          *string                            `json:"dst_group_swversion,omitempty"`
	HtypeNameConflicts         []*string                          `json:"htype_name_conflicts,omitempty"`
	HostTypeConflicts          []*NsHostType                      `json:"host_type_conflicts,omitempty"`
	ValidationError            []*NsErrorWithArguments            `json:"validation_error,omitempty"`
	ValidationErrorMsg         *string                            `json:"validation_error_msg,omitempty"`
}

// EncodeNsGroupMergeReturn - Transform NsGroupMergeReturn to nsGroupMergeReturn type
func EncodeNsGroupMergeReturn(request interface{}) (*nsGroupMergeReturn, error) {
	reqNsGroupMergeReturn := request.(*NsGroupMergeReturn)
	byte, err := json.Marshal(reqNsGroupMergeReturn)
	if err != nil {
		return nil, err
	}
	respNsGroupMergeReturnPtr := &nsGroupMergeReturn{}
	err = json.Unmarshal(byte, respNsGroupMergeReturnPtr)
	return respNsGroupMergeReturnPtr, err
}

// DecodeNsGroupMergeReturn Transform nsGroupMergeReturn to NsGroupMergeReturn type
func DecodeNsGroupMergeReturn(request interface{}) (*NsGroupMergeReturn, error) {
	reqNsGroupMergeReturn := request.(*nsGroupMergeReturn)
	byte, err := json.Marshal(reqNsGroupMergeReturn)
	if err != nil {
		return nil, err
	}
	respNsGroupMergeReturn := &NsGroupMergeReturn{}
	err = json.Unmarshal(byte, respNsGroupMergeReturn)
	return respNsGroupMergeReturn, err
}
