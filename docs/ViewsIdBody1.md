# ViewsIdBody1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | View name | [optional] [default to null]
**Path** | **string** | View path | [optional] [default to null]
**CreateDir** | **bool** | Create a directory at the specified path | [optional] [default to null]
**Alias** | **string** | For NFS-enabled views, a view alias for NFSv3 clients. | [optional] [default to null]
**Bucket** | **string** | Not yet implemented | [optional] [default to null]
**S3ObjectOwnershipRule** | **string** |  | [optional] [default to null]
**Protocols** | **[]string** | Enabled client access protocols | [optional] [default to null]
**Share** | **string** | SMB share name | [optional] [default to null]
**PolicyId** | **int32** | View policy ID. | [optional] [default to null]
**S3Versioning** | **bool** | S3 Versioning | [optional] [default to null]
**S3UnverifiedLookup** | **bool** | S3 Unverified Lookup | [optional] [default to null]
**AllowAnonymousAccess** | **bool** | Allow S3 anonymous access | [optional] [default to null]
**AllowS3AnonymousAccess** | **bool** | Allow S3 anonymous access | [optional] [default to null]
**BucketOwner** | **string** | S3 Bucket owner | [optional] [default to null]
**Locking** | **bool** | Object Locking | [optional] [default to null]
**S3LocksRetentionMode** | **string** | The retention mode for new object versions stored in this bucket. You can override this if you upload a new object version with an explicit retention mode and period. | [optional] [default to null]
**DefaultRetentionPeriod** | **string** | The amount of time in which the file/object is protected. This is mandatory if an S3 retention mode is set. | [optional] [default to null]
**NfsInteropFlags** | **string** | Indicates whether the view should support simultaneous access to NFS3/NFS4/SMB protocols. | [optional] [default to null]
**BucketCreators** | **[]string** |  | [optional] [default to null]
**BucketCreatorsGroups** | **[]string** |  | [optional] [default to null]
**ShareAcl** | [***interface{}**](interface{}.md) | Share-level ACL details | [optional] [default to null]
**SelectForLiveMonitoring** | **bool** |  | [optional] [default to null]
**QosPolicyId** | **int32** | QoS Policy ID | [optional] [default to null]
**QosPolicy** | **string** | QoS Policy | [optional] [default to null]
**AbeProtocols** | **[]string** | Protocols for which ABE is enabled | [optional] [default to null]
**AbeMaxDepth** | **int32** | The maximum depth of folders under the view on which the ABE feature is enabled | [optional] [default to null]
**IsSeamless** | **bool** | supports seamless connection between peers | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

