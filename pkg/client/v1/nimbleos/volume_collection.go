// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos


// VolumeCollection - Manage volume collections. Volume collections are logical groups of volumes that share protection characteristics such as snapshot and replication schedules. Volume collections can be created from scratch or based on predefined protection templates.
// Export VolumeCollectionFields for advance operations like search filter etc.
var VolumeCollectionFields *VolumeCollection

func init(){
 IDfield:= "id"
 ProttmplIdfield:= "prottmpl_id"
 Namefield:= "name"
 FullNamefield:= "full_name"
 SearchNamefield:= "search_name"
 Descriptionfield:= "description"
 PolOwnerNamefield:= "pol_owner_name"
 AppServerfield:= "app_server"
 AppClusterNamefield:= "app_cluster_name"
 AppServiceNamefield:= "app_service_name"
 VcenterHostnamefield:= "vcenter_hostname"
 VcenterUsernamefield:= "vcenter_username"
 VcenterPasswordfield:= "vcenter_password"
 AgentHostnamefield:= "agent_hostname"
 AgentUsernamefield:= "agent_username"
 AgentPasswordfield:= "agent_password"
 ReplicationPartnerfield:= "replication_partner"
 HandoverReplicationPartnerfield:= "handover_replication_partner"

 VolumeCollectionFields= &VolumeCollection{
  ID:                           &IDfield,
  ProttmplId:                   &ProttmplIdfield,
  Name:                         &Namefield,
  FullName:                     &FullNamefield,
  SearchName:                   &SearchNamefield,
  Description:                  &Descriptionfield,
  PolOwnerName:                 &PolOwnerNamefield,
  AppServer:                    &AppServerfield,
  AppClusterName:               &AppClusterNamefield,
  AppServiceName:               &AppServiceNamefield,
  VcenterHostname:              &VcenterHostnamefield,
  VcenterUsername:              &VcenterUsernamefield,
  VcenterPassword:              &VcenterPasswordfield,
  AgentHostname:                &AgentHostnamefield,
  AgentUsername:                &AgentUsernamefield,
  AgentPassword:                &AgentPasswordfield,
  ReplicationPartner:           &ReplicationPartnerfield,
  HandoverReplicationPartner:   &HandoverReplicationPartnerfield,
 }
}


type VolumeCollection struct {
 // ID - Identifier for volume collection.
  ID *string `json:"id,omitempty"`
 // ProttmplId - Identifier of the protection template whose attributes will be used to create this volume collection. This attribute is only used for input when creating a volume collection and is not outputed.
  ProttmplId *string `json:"prottmpl_id,omitempty"`
 // Name - Name of volume collection.
  Name *string `json:"name,omitempty"`
 // FullName - Fully qualified name of volume collection.
  FullName *string `json:"full_name,omitempty"`
 // SearchName - Name of volume collection used for object search.
  SearchName *string `json:"search_name,omitempty"`
 // Description - Text description of volume collection.
  Description *string `json:"description,omitempty"`
 // ReplPriority - Replication priority for the volume collection with the following choices: {normal | high}.
    ReplPriority *NsReplPriorityType `json:"repl_priority,omitempty"`
 // PolOwnerName - Owner group.
  PolOwnerName *string `json:"pol_owner_name,omitempty"`
 // ReplicationType - Type of replication configured for the volume collection.
    ReplicationType *NsReplicationType `json:"replication_type,omitempty"`
 // SynchronousReplicationType - Type of synchronous replication configured for the volume collection.
    SynchronousReplicationType *NsSynchronousReplicationType `json:"synchronous_replication_type,omitempty"`
 // SynchronousReplicationState - State of synchronous replication on the volume collection.
    SynchronousReplicationState *NsSynchronousReplicationState `json:"synchronous_replication_state,omitempty"`
 // AppSync - Application Synchronization.
    AppSync *NsAppSyncType `json:"app_sync,omitempty"`
 // AppServer - Application server hostname.
  AppServer *string `json:"app_server,omitempty"`
 // AppId - Application ID running on the server. Application ID can only be specified if application synchronization is \\"vss\\".
    AppId *NsAppIdType `json:"app_id,omitempty"`
 // AppClusterName - If the application is running within a Windows cluster environment, this is the cluster name.
  AppClusterName *string `json:"app_cluster_name,omitempty"`
 // AppServiceName - If the application is running within a Windows cluster environment then this is the instance name of the service running within the cluster environment.
  AppServiceName *string `json:"app_service_name,omitempty"`
 // VcenterHostname - VMware vCenter hostname. Custom port number can be specified with vCenter hostname using \\":\\".
  VcenterHostname *string `json:"vcenter_hostname,omitempty"`
 // VcenterUsername - Application VMware vCenter username.
  VcenterUsername *string `json:"vcenter_username,omitempty"`
 // VcenterPassword - Application VMware vCenter password.
  VcenterPassword *string `json:"vcenter_password,omitempty"`
 // AgentHostname - Generic backup agent hostname. Custom port number can be specified with agent hostname using \\":\\".
  AgentHostname *string `json:"agent_hostname,omitempty"`
 // AgentUsername - Generic backup agent username.
  AgentUsername *string `json:"agent_username,omitempty"`
 // AgentPassword - Generic backup agent password.
  AgentPassword *string `json:"agent_password,omitempty"`
 // CreationTime - Time when this volume collection was created.
    CreationTime *int64 `json:"creation_time,omitempty"`
 // LastModifiedTime - Time when this volume collection was last modified.
    LastModifiedTime *int64 `json:"last_modified_time,omitempty"`
 // VolumeList - List of volumes associated with the volume collection.
    VolumeList []*NsVolumeSummary `json:"volume_list,omitempty"`
 // DownstreamVolumeList - List of downstream volumes associated with the volume collection.
    DownstreamVolumeList []*NsVolumePoolInfo `json:"downstream_volume_list,omitempty"`
 // UpstreamVolumeList - List of upstream volumes associated with the volume collection.
    UpstreamVolumeList []*NsVolumePoolInfo `json:"upstream_volume_list,omitempty"`
 // VolumeCount - Count of volumes associated with the volume collection.
    VolumeCount *int64 `json:"volume_count,omitempty"`
 // CachePinnedVolumeList - List of cache pinned volumes associated with volume collection.
    CachePinnedVolumeList []*NsVolumeSummary `json:"cache_pinned_volume_list,omitempty"`
 // LastSnapcoll - Last snapshot collection on this volume collection.
 LastSnapcoll *NsSnapcollSummary `json:"last_snapcoll,omitempty"`
 // SnapcollCount - Count of snapshot collections associated with volume collection.
    SnapcollCount *int64 `json:"snapcoll_count,omitempty"`
 // ScheduleList - List of snapshot schedules associated with volume collection.
    ScheduleList []*NsSchedule `json:"schedule_list,omitempty"`
 // ReplicationPartner - Replication partner for this volume collection.
  ReplicationPartner *string `json:"replication_partner,omitempty"`
 // LastReplicatedSnapcoll - Last replicated snapshot collection on this volume collection.
 LastReplicatedSnapcoll *NsSnapcollSummary `json:"last_replicated_snapcoll,omitempty"`
 // LastReplicatedSnapcollList - List of snapshot collection information for the last replicated snapshot collection per schedule.
    LastReplicatedSnapcollList []*NsReplicatedSnapcollSummary `json:"last_replicated_snapcoll_list,omitempty"`
 // ProtectionType - Specifies if volume collection is protected with schedules. If protected, indicated whether replication is setup.
    ProtectionType *NsProtectionType `json:"protection_type,omitempty"`
 // LagTime - Replication lag time for volume collection.
    LagTime *int64 `json:"lag_time,omitempty"`
 // IsStandaloneVolcoll - Indicates whether this is a standalone volume collection.
    IsStandaloneVolcoll *bool `json:"is_standalone_volcoll,omitempty"`
 // TotalReplBytes - Total size of volumes to be replicated for this volume collection.
    TotalReplBytes *int64 `json:"total_repl_bytes,omitempty"`
 // ReplBytesTransferred - Total size of volumes replicated for this volume collection.
    ReplBytesTransferred *int64 `json:"repl_bytes_transferred,omitempty"`
 // IsHandingOver - Indicates whether a handover operation is in progress on this volume collection.
    IsHandingOver *bool `json:"is_handing_over,omitempty"`
 // HandoverReplicationPartner - Replication partner to which ownership is being transferred as part of handover operation.
  HandoverReplicationPartner *string `json:"handover_replication_partner,omitempty"`
 // Metadata - Key-value pairs that augment a volume collection's attributes.
    Metadata []*NsKeyValue `json:"metadata,omitempty"`
 // SrepLastSync - Time when a synchronously replicated volume collection was last synchronized.
    SrepLastSync *int64 `json:"srep_last_sync,omitempty"`
 // SrepResyncPercent - Percentage of the resync progress for a synchronously replicated volume collection.
    SrepResyncPercent *int64 `json:"srep_resync_percent,omitempty"`
}
