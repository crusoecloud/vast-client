# ReplicationstreamsBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to null]
**SourceDir** | **string** | path to replicate | [optional] [default to null]
**PolicyId** | **string** | replication policy id | [optional] [default to null]
**Enabled** | **bool** | enable/pause replication stream | [optional] [default to null]
**ProtectionPolicyId** | **string** | protection policy id | [optional] [default to null]
**ProtectedPathId** | **string** | protected path id | [optional] [default to null]
**PriorityScore** | **float64** | Indicates how close replication is to missing its RPO in relation to their current interval. A lower score means a higher priority. | [optional] [default to null]
**IsManualPriorityScore** | **bool** | Indicates whether the priority score is set manually by a user. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

