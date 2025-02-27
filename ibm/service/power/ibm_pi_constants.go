package power

import "time"

const (
	// Arguments
	Arg_Action                              = "pi_action"
	Arg_AffinityInstance                    = "pi_affinity_instance"
	Arg_AffinityPolicy                      = "pi_affinity_policy"
	Arg_AffinityVolume                      = "pi_affinity_volume"
	Arg_AntiAffinityInstances               = "pi_anti_affinity_instances"
	Arg_AntiAffinityVolumes                 = "pi_anti_affinity_volumes"
	Arg_BootVolumeReplicationEnabled        = "pi_boot_volume_replication_enabled"
	Arg_Cidr                                = "pi_cidr"
	Arg_CloudConnectionID                   = "pi_cloud_connection_id"
	Arg_CloudConnectionName                 = "pi_cloud_connection_name"
	Arg_CloudInstanceID                     = "pi_cloud_instance_id"
	Arg_ConsistencyGroupName                = "pi_consistency_group_name"
	Arg_Datacenter                          = "pi_datacenter"
	Arg_DatacenterZone                      = "pi_datacenter_zone"
	Arg_DeploymentTarget                    = "pi_deployment_target"
	Arg_DeploymentType                      = "pi_deployment_type"
	Arg_Description                         = "pi_description"
	Arg_DestinationPorts                    = "pi_destination_ports"
	Arg_DhcpID                              = "pi_dhcp_id"
	Arg_DhcpName                            = "pi_dhcp_name"
	Arg_DhcpSnatEnabled                     = "pi_dhcp_snat_enabled"
	Arg_DnsServer                           = "pi_dns_server"
	Arg_HealthStatus                        = "pi_health_status"
	Arg_Host                                = "pi_host"
	Arg_HostGroupID                         = "pi_host_group_id"
	Arg_HostID                              = "pi_host_id"
	Arg_Hosts                               = "pi_hosts"
	Arg_IBMiCSS                             = "pi_ibmi_css"
	Arg_IBMiPHA                             = "pi_ibmi_pha"
	Arg_IBMiRDSUsers                        = "pi_ibmi_rds_users"
	Arg_ImageID                             = "pi_image_id"
	Arg_ImageImportDetails                  = "pi_image_import_details"
	Arg_ImageName                           = "pi_image_name"
	Arg_InstanceID                          = "pi_instance_id"
	Arg_InstanceName                        = "pi_instance_name"
	Arg_IPAddress                           = "pi_ip_address"
	Arg_Key                                 = "pi_ssh_key"
	Arg_KeyName                             = "pi_key_name"
	Arg_KeyPairName                         = "pi_key_pair_name"
	Arg_LanguageCode                        = "pi_language_code"
	Arg_LicenseRepositoryCapacity           = "pi_license_repository_capacity"
	Arg_Memory                              = "pi_memory"
	Arg_Name                                = "pi_name"
	Arg_Network                             = "pi_network"
	Arg_NetworkAddressGroupID               = "pi_network_address_group_id"
	Arg_NetworkAddressGroupMemberID         = "pi_network_address_group_member_id"
	Arg_NetworkID                           = "pi_network_id"
	Arg_NetworkInterfaceID                  = "pi_network_interface_id"
	Arg_NetworkName                         = "pi_network_name"
	Arg_NetworkSecurityGroupID              = "pi_network_security_group_id"
	Arg_NetworkSecurityGroupMemberID        = "pi_network_security_group_member_id"
	Arg_NetworkSecurityGroupRuleID          = "pi_network_security_group_rule_id"
	Arg_PinPolicy                           = "pi_pin_policy"
	Arg_PlacementGroupID                    = "pi_placement_group_id"
	Arg_PlacementGroupName                  = "pi_placement_group_name"
	Arg_PlacementGroupPolicy                = "pi_placement_group_policy"
	Arg_Plan                                = "pi_plan"
	Arg_Processors                          = "pi_processors"
	Arg_ProcType                            = "pi_proc_type"
	Arg_Protocol                            = "pi_protocol"
	Arg_Remote                              = "pi_remote"
	Arg_Remove                              = "pi_remove"
	Arg_Replicants                          = "pi_replicants"
	Arg_ReplicationEnabled                  = "pi_replication_enabled"
	Arg_ReplicationPolicy                   = "pi_replication_policy"
	Arg_ReplicationScheme                   = "pi_replication_scheme"
	Arg_ReplicationSites                    = "pi_replication_sites"
	Arg_ResourceGroupID                     = "pi_resource_group_id"
	Arg_SAP                                 = "sap"
	Arg_SAPDeploymentType                   = "pi_sap_deployment_type"
	Arg_SAPProfileID                        = "pi_sap_profile_id"
	Arg_Secondaries                         = "pi_secondaries"
	Arg_SharedProcessorPool                 = "pi_shared_processor_pool"
	Arg_SharedProcessorPoolHostGroup        = "pi_shared_processor_pool_host_group"
	Arg_SharedProcessorPoolID               = "pi_shared_processor_pool_id"
	Arg_SharedProcessorPoolName             = "pi_shared_processor_pool_name"
	Arg_SharedProcessorPoolPlacementGroupID = "pi_shared_processor_pool_placement_group_id"
	Arg_SharedProcessorPoolPlacementGroups  = "pi_shared_processor_pool_placement_groups"
	Arg_SharedProcessorPoolReservedCores    = "pi_shared_processor_pool_reserved_cores"
	Arg_SnapshotID                          = "pi_snapshot_id"
	Arg_SnapShotName                        = "pi_snap_shot_name"
	Arg_SourcePorts                         = "pi_source_ports"
	Arg_SPPPlacementGroupID                 = "pi_spp_placement_group_id"
	Arg_SPPPlacementGroupName               = "pi_spp_placement_group_name"
	Arg_SPPPlacementGroupPolicy             = "pi_spp_placement_group_policy"
	Arg_SSHKey                              = "pi_ssh_key"
	Arg_StorageConnection                   = "pi_storage_connection"
	Arg_StoragePool                         = "pi_storage_pool"
	Arg_StoragePoolAffinity                 = "pi_storage_pool_affinity"
	Arg_StorageType                         = "pi_storage_type"
	Arg_SysType                             = "pi_sys_type"
	Arg_Target                              = "pi_target"
	Arg_TargetStorageTier                   = "pi_target_storage_tier"
	Arg_Type                                = "pi_type"
	Arg_UserData                            = "pi_user_data"
	Arg_UserTags                            = "pi_user_tags"
	Arg_VirtualCoresAssigned                = "pi_virtual_cores_assigned"
	Arg_VirtualOpticalDevice                = "pi_virtual_optical_device"
	Arg_VolumeCloneName                     = "pi_volume_clone_name"
	Arg_VolumeCloneTaskID                   = "pi_volume_clone_task_id"
	Arg_VolumeGroupID                       = "pi_volume_group_id"
	Arg_VolumeGroupName                     = "pi_volume_group_name"
	Arg_VolumeID                            = "pi_volume_id"
	Arg_VolumeIDs                           = "pi_volume_ids"
	Arg_VolumeName                          = "pi_volume_name"
	Arg_VolumeOnboardingID                  = "pi_volume_onboarding_id"
	Arg_VolumePool                          = "pi_volume_pool"
	Arg_VolumeShareable                     = "pi_volume_shareable"
	Arg_VolumeSize                          = "pi_volume_size"
	Arg_VolumeSnapshotID                    = "pi_volume_snapshot_id"
	Arg_VolumeType                          = "pi_volume_type"
	Arg_VTL                                 = "vtl"

	// Attributes
	Attr_Access                             = "access"
	Attr_AccessConfig                       = "access_config"
	Attr_Action                             = "action"
	Attr_AllocatedCores                     = "allocated_cores"
	Attr_Architecture                       = "architecture"
	Attr_AsynchronousReplication            = "asynchronous_replication"
	Attr_Auxiliary                          = "auxiliary"
	Attr_AuxiliaryChangedVolumeName         = "auxiliary_changed_volume_name"
	Attr_AuxiliaryVolumeName                = "auxiliary_volume_name"
	Attr_AvailabilityZone                   = "availability_zone"
	Attr_AvailableCores                     = "available_cores"
	Attr_AvailableHosts                     = "available_hosts"
	Attr_AvailableIPCount                   = "available_ip_count"
	Attr_AvailableMemory                    = "available_memory"
	Attr_Bootable                           = "bootable"
	Attr_BootVolumeID                       = "boot_volume_id"
	Attr_Capabilities                       = "capabilities"
	Attr_CapabilityDetails                  = "capability_details"
	Attr_Capacity                           = "capacity"
	Attr_Certified                          = "certified"
	Attr_CIDR                               = "cidr"
	Attr_ClassicEnabled                     = "classic_enabled"
	Attr_ClonedVolumes                      = "clone_volumes"
	Attr_CloneVolumeID                      = "clone_volume_id"
	Attr_CloudConnectionID                  = "cloud_connection_id"
	Attr_CloudInstanceID                    = "cloud_instance_id"
	Attr_CloudInstances                     = "cloud_instances"
	Attr_Code                               = "code"
	Attr_ConnectionMode                     = "connection_mode"
	Attr_Connections                        = "connections"
	Attr_ConsistencyGroupName               = "consistency_group_name"
	Attr_ConsoleLanguages                   = "console_languages"
	Attr_ContainerFormat                    = "container_format"
	Attr_CopyRate                           = "copy_rate"
	Attr_CopyType                           = "copy_type"
	Attr_CoreMemoryRatio                    = "core_memory_ratio"
	Attr_Cores                              = "cores"
	Attr_Count                              = "count"
	Attr_CPUs                               = "cpus"
	Attr_Created                            = "created"
	Attr_CreateTime                         = "create_time"
	Attr_CreationDate                       = "creation_date"
	Attr_CRN                                = "crn"
	Attr_CyclePeriodSeconds                 = "cycle_period_seconds"
	Attr_CyclingMode                        = "cycling_mode"
	Attr_DatacenterCapabilities             = "pi_datacenter_capabilities"
	Attr_DatacenterHref                     = "pi_datacenter_href"
	Attr_DatacenterLocation                 = "pi_datacenter_location"
	Attr_Datacenters                        = "datacenters"
	Attr_DatacenterStatus                   = "pi_datacenter_status"
	Attr_DatacenterType                     = "pi_datacenter_type"
	Attr_Dedicated                          = "dedicated"
	Attr_Default                            = "default"
	Attr_DeleteOnTermination                = "delete_on_termination"
	Attr_DeploymentType                     = "deployment_type"
	Attr_Description                        = "description"
	Attr_DestinationPort                    = "destination_port"
	Attr_Details                            = "details"
	Attr_DhcpID                             = "dhcp_id"
	Attr_DhcpManaged                        = "dhcp_managed"
	Attr_DisasterRecovery                   = "disaster_recovery"
	Attr_DisasterRecoveryLocations          = "disaster_recovery_locations"
	Attr_DiskFormat                         = "disk_format"
	Attr_DiskType                           = "disk_type"
	Attr_DisplayName                        = "display_name"
	Attr_DNS                                = "dns"
	Attr_Enabled                            = "enabled"
	Attr_Endianness                         = "endianness"
	Attr_ExternalIP                         = "external_ip"
	Attr_FailureMessage                     = "failure_message"
	Attr_FailureReason                      = "failure_reason"
	Attr_Fault                              = "fault"
	Attr_Flag                               = "flag"
	Attr_FlashCopyMappings                  = "flash_copy_mappings"
	Attr_FlashCopyName                      = "flash_copy_name"
	Attr_FreezeTime                         = "freeze_time"
	Attr_FullSystemProfile                  = "full_system_profile"
	Attr_Gateway                            = "gateway"
	Attr_General                            = "general"
	Attr_GlobalRouting                      = "global_routing"
	Attr_GreDestinationAddress              = "gre_destination_address"
	Attr_GreSourceAddress                   = "gre_source_address"
	Attr_GroupID                            = "group_id"
	Attr_HealthStatus                       = "health_status"
	Attr_HostGroup                          = "host_group"
	Attr_HostGroupID                        = "host_group_id"
	Attr_HostGroups                         = "host_groups"
	Attr_HostID                             = "host_id"
	Attr_Hosts                              = "hosts"
	Attr_Href                               = "href"
	Attr_Hypervisor                         = "hypervisor"
	Attr_HypervisorType                     = "hypervisor_type"
	Attr_IBMiCSS                            = "ibmi_css"
	Attr_IBMIPAddress                       = "ibm_ip_address"
	Attr_IBMiPHA                            = "ibmi_pha"
	Attr_IBMiRDS                            = "ibmi_rds"
	Attr_IBMiRDSUsers                       = "ibmi_rds_users"
	Attr_ICMPType                           = "icmp_type"
	Attr_ID                                 = "id"
	Attr_ImageID                            = "image_id"
	Attr_ImageInfo                          = "image_info"
	Attr_Images                             = "images"
	Attr_ImageType                          = "image_type"
	Attr_InputVolumes                       = "input_volumes"
	Attr_Instance                           = "instance"
	Attr_InstanceID                         = "instance_id"
	Attr_InstanceIP                         = "instance_ip"
	Attr_InstanceMac                        = "instance_mac"
	Attr_Instances                          = "instances"
	Attr_InstanceSnapshots                  = "instance_snapshots"
	Attr_InstanceVolumes                    = "instance_volumes"
	Attr_Interfaces                         = "interfaces"
	Attr_IOThrottleRate                     = "io_throttle_rate"
	Attr_IP                                 = "ip"
	Attr_IPAddress                          = "ip_address"
	Attr_IPaddress                          = "ipaddress"
	Attr_IPOctet                            = "ipoctet"
	Attr_IsActive                           = "is_active"
	Attr_Jumbo                              = "jumbo"
	Attr_Key                                = "key"
	Attr_Keys                               = "keys"
	Attr_Language                           = "language"
	Attr_LastUpdateDate                     = "last_update_date"
	Attr_LastUpdatedDate                    = "last_updated_date"
	Attr_Leases                             = "leases"
	Attr_LicenseRepositoryCapacity          = "license_repository_capacity"
	Attr_LicenseType                        = "license_type"
	Attr_Location                           = "location"
	Attr_MacAddress                         = "mac_address"
	Attr_Macaddress                         = "macaddress"
	Attr_MasterChangedVolumeName            = "master_changed_volume_name"
	Attr_MasterVolumeName                   = "master_volume_name"
	Attr_Max                                = "max"
	Attr_MaxAllocationSize                  = "max_allocation_size"
	Attr_MaxAvailable                       = "max_available"
	Attr_MaxCoresAvailable                  = "max_cores_available"
	Attr_Maximum                            = "maximum"
	Attr_MaximumStorageAllocation           = "max_storage_allocation"
	Attr_MaxMem                             = "maxmem"
	Attr_MaxMemory                          = "max_memory"
	Attr_MaxMemoryAvailable                 = "max_memory_available"
	Attr_MaxProc                            = "maxproc"
	Attr_MaxProcessors                      = "max_processors"
	Attr_MaxVirtualCores                    = "max_virtual_cores"
	Attr_Members                            = "members"
	Attr_Memory                             = "memory"
	Attr_Message                            = "message"
	Attr_Metered                            = "metered"
	Attr_MigrationStatus                    = "migration_status"
	Attr_Min                                = "min"
	Attr_Minimum                            = "minimum"
	Attr_MinMem                             = "minmem"
	Attr_MinMemory                          = "min_memory"
	Attr_MinProc                            = "minproc"
	Attr_MinProcessors                      = "min_processors"
	Attr_MinVirtualCores                    = "min_virtual_cores"
	Attr_MirroringState                     = "mirroring_state"
	Attr_MTU                                = "mtu"
	Attr_Name                               = "name"
	Attr_NetworkAddressGroupID              = "network_address_group_id"
	Attr_NetworkAddressGroups               = "network_address_groups"
	Attr_NetworkID                          = "network_id"
	Attr_NetworkInterfaceID                 = "network_interface_id"
	Attr_NetworkName                        = "network_name"
	Attr_NetworkPorts                       = "network_ports"
	Attr_Networks                           = "networks"
	Attr_NetworkSecurityGroupID             = "network_security_group_id"
	Attr_NetworkSecurityGroupMemberID       = "network_security_group_member_id"
	Attr_NetworkSecurityGroups              = "network_security_groups"
	Attr_NumberOfVolumes                    = "number_of_volumes"
	Attr_Onboardings                        = "onboardings"
	Attr_OperatingSystem                    = "operating_system"
	Attr_OSType                             = "os_type"
	Attr_PercentComplete                    = "percent_complete"
	Attr_PinPolicy                          = "pin_policy"
	Attr_PlacementGroupID                   = "placement_group_id"
	Attr_PlacementGroups                    = "placement_groups"
	Attr_Policy                             = "policy"
	Attr_Pool                               = "pool"
	Attr_PoolName                           = "pool_name"
	Attr_Port                               = "port"
	Attr_PortID                             = "portid"
	Attr_PowerEdgeRouter                    = "power_edge_router"
	Attr_Primary                            = "primary"
	Attr_PrimaryRole                        = "primary_role"
	Attr_Processors                         = "processors"
	Attr_ProcType                           = "proctype"
	Attr_Product                            = "product"
	Attr_ProfileID                          = "profile_id"
	Attr_Profiles                           = "profiles"
	Attr_Progress                           = "progress"
	Attr_Protocol                           = "protocol"
	Attr_PublicIP                           = "public_ip"
	Attr_PVMInstanceID                      = "pvm_instance_id"
	Attr_PVMInstances                       = "pvm_instances"
	Attr_PVMSnapshots                       = "pvm_snapshots"
	Attr_Region                             = "region"
	Attr_RegionStorageTiers                 = "region_storage_tiers"
	Attr_Remote                             = "remote"
	Attr_RemoteCopyID                       = "remote_copy_id"
	Attr_RemoteCopyRelationshipNames        = "remote_copy_relationship_names"
	Attr_RemoteCopyRelationships            = "remote_copy_relationships"
	Attr_RemotePool                         = "remote_pool"
	Attr_ReplicationEnabled                 = "replication_enabled"
	Attr_ReplicationPoolMap                 = "replication_pool_map"
	Attr_ReplicationSites                   = "replication_sites"
	Attr_ReplicationStatus                  = "replication_status"
	Attr_ReplicationType                    = "replication_type"
	Attr_ReservedCore                       = "reserved_core"
	Attr_ReservedCores                      = "reserved_cores"
	Attr_ReservedMemory                     = "reserved_memory"
	Attr_ResultsOnboardedVolumes            = "results_onboarded_volumes"
	Attr_ResultsVolumeOnboardingFailures    = "results_volume_onboarding_failures"
	Attr_Rules                              = "rules"
	Attr_SAPS                               = "saps"
	Attr_Secondaries                        = "secondaries"
	Attr_ServerName                         = "server_name"
	Attr_Servers                            = "servers"
	Attr_Shareable                          = "shreable"
	Attr_SharedCoreRatio                    = "shared_core_ratio"
	Attr_SharedProcessorPool                = "shared_processor_pool"
	Attr_SharedProcessorPoolID              = "shared_processor_pool_id"
	Attr_SharedProcessorPoolPlacementGroups = "spp_placement_groups"
	Attr_SharedProcessorPools               = "shared_processor_pools"
	Attr_Size                               = "size"
	Attr_SnapshotID                         = "snapshot_id"
	Attr_SourceChecksum                     = "source_checksum"
	Attr_SourcePort                         = "source_port"
	Attr_SourceVolumeID                     = "source_volume_id"
	Attr_SourceVolumeName                   = "source_volume_name"
	Attr_Speed                              = "speed"
	Attr_SPPPlacementGroupID                = "spp_placement_group_id"
	Attr_SPPPlacementGroupMembers           = "members"
	Attr_SPPPlacementGroupName              = "name"
	Attr_SPPPlacementGroupPolicy            = "policy"
	Attr_SPPPlacementGroups                 = "spp_placement_groups"
	Attr_SSHKey                             = "ssh_key"
	Attr_StartTime                          = "start_time"
	Attr_State                              = "state"
	Attr_Status                             = "status"
	Attr_StatusDescriptionErrors            = "status_description_errors"
	Attr_StatusDetail                       = "status_detail"
	Attr_StorageConnection                  = "storage_connection"
	Attr_StoragePool                        = "storage_pool"
	Attr_StoragePoolAffinity                = "storage_pool_affinity"
	Attr_StoragePoolsCapacity               = "storage_pools_capacity"
	Attr_StorageType                        = "storage_type"
	Attr_StorageTypesCapacity               = "storage_types_capacity"
	Attr_SupportedSystems                   = "supported_systems"
	Attr_Synchronized                       = "synchronized"
	Attr_SynchronousReplication             = "synchronous_replication"
	Attr_SystemPoolName                     = "system_pool_name"
	Attr_SystemPools                        = "system_pools"
	Attr_Systems                            = "systems"
	Attr_SysType                            = "sys_type"
	Attr_Systype                            = "systype"
	Attr_Target                             = "target"
	Attr_TargetLocations                    = "target_locations"
	Attr_TargetVolumeName                   = "target_volume_name"
	Attr_TaskID                             = "task_id"
	Attr_TCPFlags                           = "tcp_flags"
	Attr_TenantID                           = "tenant_id"
	Attr_TenantName                         = "tenant_name"
	Attr_TotalCapacity                      = "total_capacity"
	Attr_TotalCore                          = "total_core"
	Attr_TotalInstances                     = "total_instances"
	Attr_TotalMemory                        = "total_memory"
	Attr_TotalMemoryConsumed                = "total_memory_consumed"
	Attr_TotalProcessorsConsumed            = "total_processors_consumed"
	Attr_TotalSSDStorageConsumed            = "total_ssd_storage_consumed"
	Attr_TotalStandardStorageConsumed       = "total_standard_storage_consumed"
	Attr_Type                               = "type"
	Attr_Uncapped                           = "uncapped"
	Attr_UpdatedDate                        = "updated_date"
	Attr_URL                                = "url"
	Attr_UsedCore                           = "used_core"
	Attr_UsedIPCount                        = "used_ip_count"
	Attr_UsedIPPercent                      = "used_ip_percent"
	Attr_UsedMemory                         = "used_memory"
	Attr_UserIPAddress                      = "user_ip_address"
	Attr_UserTags                           = "user_tags"
	Attr_VCPUs                              = "vcpus"
	Attr_Vendor                             = "vendor"
	Attr_VirtualCoresAssigned               = "virtual_cores_assigned"
	Attr_VLanID                             = "vlan_id"
	Attr_VolumeGroupID                      = "volume_group_id"
	Attr_VolumeGroupName                    = "volume_group_name"
	Attr_VolumeGroups                       = "volume_groups"
	Attr_VolumeGroupStatus                  = "volume_group_status"
	Attr_VolumeID                           = "volume_id"
	Attr_VolumeIDs                          = "volume_ids"
	Attr_VolumePool                         = "volume_pool"
	Attr_Volumes                            = "volumes"
	Attr_VolumeSnapshots                    = "volume_snapshots"
	Attr_VolumesSnapshots                   = "volume_snapshots"
	Attr_VolumeStatus                       = "volume_status"
	Attr_VPCCRNs                            = "vpc_crns"
	Attr_VPCEnabled                         = "vpc_enabled"
	Attr_WorkloadType                       = "workload_type"
	Attr_Workspace                          = "workspace"
	Attr_WorkspaceCapabilities              = "pi_workspace_capabilities"
	Attr_WorkspaceDetails                   = "pi_workspace_details"
	Attr_WorkspaceID                        = "pi_workspace_id"
	Attr_WorkspaceLocation                  = "pi_workspace_location"
	Attr_WorkspaceName                      = "pi_workspace_name"
	Attr_Workspaces                         = "workspaces"
	Attr_WorkspaceStatus                    = "pi_workspace_status"
	Attr_WorkspaceType                      = "pi_workspace_type"
	Attr_WWN                                = "wwn"

	// Duplicate Attributes, will be removed as refactoring take course.
	PICloudConnectionClassicGreSource      = "gre_source_address"
	PICloudConnectionConnectionMode        = "connection_mode"
	PICloudConnectionIBMIPAddress          = "ibm_ip_address"
	PICloudConnectionId                    = "cloud_connection_id"
	PICloudConnectionPort                  = "port"
	PICloudConnectionUserIPAddress         = "user_ip_address"
	PIVPNConnectionDeadPeerDetectionAction = "action"

	// OS Type
	OS_IBMI  = "ibmi"
	StockVTL = "stock-vtl"

	// Allowed Values
	Affinity                  = "affinity"
	All                       = "all"
	Allow                     = "allow"
	AntiAffinity              = "anti-affinity"
	Attach                    = "attach"
	BYOL                      = "byol"
	Capped                    = "capped"
	Critical                  = "CRITICAL"
	CUSTOM_VIRTUAL_CORES      = "custom-virtualcores"
	Dedicated                 = "dedicated"
	DefaultNAG                = "default-network-address-group"
	Deny                      = "deny"
	DeploymentTypeEpic        = "EPIC"
	DeploymentTypeVMNoStorage = "VMNoStorage"
	DestinationUnreach        = "destination-unreach"
	DHCPVlan                  = "dhcp-vlan"
	Disable                   = "disable"
	Echo                      = "echo"
	EchoReply                 = "echo-reply"
	Enable                    = "enable"
	Hana                      = "Hana"
	Hard                      = "hard"
	Host                      = "host"
	HostGroup                 = "hostGroup"
	ICMP                      = "icmp"
	IPV4_Address              = "ipv4-address"
	NAG                       = "network-address-group"
	MaxVolumeSupport          = "maxVolumeSupport"
	Netweaver                 = "Netweaver"
	Network_Interface         = "network-interface"
	None                      = "none"
	NSG                       = "network-security-group"
	OK                        = "OK"
	PER                       = "power-edge-router"
	Prefix                    = "prefix"
	Private                   = "private"
	Public                    = "public"
	PubVlan                   = "pub-vlan"
	SAP                       = "SAP"
	Shared                    = "shared"
	Soft                      = "soft"
	SourceQuench              = "source-quench"
	Suffix                    = "suffix"
	TCP                       = "tcp"
	TimeExceeded              = "time-exceeded"
	UDP                       = "udp"
	UserTagType               = "user"
	Vlan                      = "vlan"
	vSCSI                     = "vSCSI"
	Warning                   = "WARNING"

	// Actions
	Action_HardReboot        = "hard-reboot"
	Action_ImmediateShutdown = "immediate-shutdown"
	Action_ResetState        = "reset-state"
	Action_SoftReboot        = "soft-reboot"
	Action_Start             = "start"
	Action_Stop              = "stop"

	// States
	NotFound                 = "not found"
	State_Active             = "active"
	State_ACTIVE             = "ACTIVE"
	State_Added              = "added"
	State_Adding             = "adding"
	State_Available          = "available"
	State_Build              = "build"
	State_Building           = "building"
	State_Completed          = "completed"
	State_Configuring        = "configuring"
	State_Creating           = "creating"
	State_Deleted            = "deleted"
	State_Deleting           = "deleting"
	State_Down               = "down"
	State_Error              = "error"
	State_ERROR              = "ERROR"
	State_Failed             = "failed"
	State_Found              = "Found"
	State_Inactive           = "inactive"
	State_InProgress         = "in progress"
	State_InUse              = "in-use"
	State_NotFound           = "not found"
	State_Pending            = "pending"
	State_PENDING            = "PENDING"
	State_PendingReclamation = "pending_reclamation"
	State_Provisioning       = "provisioning"
	State_Removed            = "removed"
	State_Removing           = "removing"
	State_Resize             = "resize"
	State_RESIZE             = "RESIZE"
	State_Retry              = "retry"
	State_Shutoff            = "shutoff"
	State_SHUTOFF            = "SHUTOFF"
	State_Stopping           = "stopping"
	State_Up                 = "up"
	State_Updating           = "updating"
	State_VerifyResize       = "verify_resize"

	// Timeout values
	Timeout_Active  = 2 * time.Minute
	Timeout_Delay   = 60 * time.Second
	Timeout_Warning = 60 * time.Second

	// TODO: Second Half Cleanup, remove extra variables

	// IBM PI Volume Group
	PIVolumeGroupAction = "pi_volume_group_action"
	PIVolumeGroupID     = "pi_volume_group_id"

	// VPN
	PIVPNConnectionId                         = "connection_id"
	PIVPNConnectionStatus                     = "connection_status"
	PIVPNConnectionDeadPeerDetection          = "dead_peer_detections"
	PIVPNConnectionDeadPeerDetectionInterval  = "interval"
	PIVPNConnectionDeadPeerDetectionThreshold = "threshold"
	PIVPNConnectionLocalGatewayAddress        = "local_gateway_address"
	PIVPNConnectionVpnGatewayAddress          = "gateway_address"

	// Cloud Connections
	PICloudConnectionTransitEnabled = "pi_cloud_connection_transit_enabled"
)
