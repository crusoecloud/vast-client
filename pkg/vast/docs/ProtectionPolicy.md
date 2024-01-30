# ProtectionPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Title** | **string** |  | [optional] [default to null]
**Guid** | **string** | unique identifier | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**IsOnSchedule** | **bool** |  | [optional] [default to null]
**PrettySchedules** | **[]string** |  | [optional] [default to null]
**ReplicationTarget** | **string** |  | [optional] [default to null]
**Handle** | **string** |  | [optional] [default to null]
**Created** | **string** |  | [optional] [default to null]
**ScheduleMiss** | **int32** |  | [optional] [default to null]
**NativeReplicationRemoteTarget** | **string** |  | [optional] [default to null]
**IsLocal** | **bool** |  | [optional] [default to null]
**TargetGuid** | **string** |  | [optional] [default to null]
**TargetObjectId** | **int32** | target object id | [optional] [default to null]
**TargetName** | **string** | Target Name | [optional] [default to null]
**Frames** | [**[]Frame**](Frame.md) | Defines the schedule for snapshot creation and the local and remote retention policies. | [optional] [default to null]
**Prefix** | **string** | The prefix of the snapshot that will be created | [optional] [default to null]
**CloneType** | **string** | Specify the type of data protection. CLOUD_REPLICATION is S3 backup. LOCAL means local snapshots without replication. (LOCAL | NATIVE_REPLICATION | CLOUD_REPLICATION) | [optional] [default to null]
**State** | **string** | State of Protection Policy | [optional] [default to null]
**SyncInterval** | **int32** | A sync point is a common restore point for all group members. This value guarantees such a sync point exists in this duration. In other words, this is the maximal sync duration gap between other members | [optional] [default to null]
**Internal** | **bool** |  | [optional] [default to null]
**Indestructible** | **bool** | Indestructible the Protection policy from being deleted | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

