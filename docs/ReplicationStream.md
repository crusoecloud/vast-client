# ReplicationStream

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Guid** | **string** | guid | [optional] [default to null]
**State** | **string** | state | [optional] [default to null]
**Role** | **string** | current role in the replication | [optional] [default to null]
**SourceDir** | **string** | path to replicate | [optional] [default to null]
**ReplicationPolicy** | **string** | replication policy id | [optional] [default to null]
**ProtectionPolicy** | **string** | protection policy id | [optional] [default to null]
**Enabled** | **bool** | start/pause replication | [optional] [default to null]
**Bw** | **int64** | Replication Bandwidth | [optional] [default to null]
**Internal** | **bool** |  | [optional] [default to null]
**RemoteTenantName** | **string** | remote tenant name | [optional] [default to null]
**RemoteTenantGuid** | **string** | remote tenant guid | [optional] [default to null]
**ReplicationGroupId** | **int32** | replication group id | [optional] [default to null]
**ProtectedPathId** | **int32** | protected path id | [optional] [default to null]
**ProtectedPathGuid** | **string** | protected path guid | [optional] [default to null]
**PriorityScore** | **float64** | Indicates how close replication is to missing its RPO in relation to their current interval. A lower score means a higher priority. | [optional] [default to null]
**IsManualPriorityScore** | **bool** | Indicates whether the priority score is set manually by a user. | [optional] [default to null]
**PriorityNumber** | **int32** | Priority of replication job. A lower number means higher priority. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

