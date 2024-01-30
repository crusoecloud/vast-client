# ClusterAuditing

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditDirName** | **string** | Name of the audit directory, which is located under the root directory. | [optional] [default to null]
**ReadAccessUsers** | **[]string** | Enter users here to grant them read access to all files in the audit directory. To make the audit directory accessible to clients, create a view on the directory. | [optional] [default to null]
**ReadAccessUsersGroups** | **[]string** | Enter groups here to grant them read access to all files in the audit directory. To make the audit directory accessible to clients, create a view on the directory. | [optional] [default to null]
**MaxFileSize** | **int32** | Maximum audit file size for each CNode core | [optional] [default to null]
**MaxRetentionPeriod** | **int32** | Max retention period for audit files | [optional] [default to null]
**MaxRetentionTimeunit** | **string** | Max retention period timeunit for audit files | [optional] [default to null]
**ProtocolsAudit** | [***interface{}**](interface{}.md) | Map of protocols audit configurations | [optional] [default to null]
**Protocols** | **[]string** | Protocols to audit | [optional] [default to null]
**MaxAuditDirSize** | **string** | Maximum audit directory size | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

