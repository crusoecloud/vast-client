# QosPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [default to null]
**Guid** | **string** | QoS Policy guid | [default to null]
**Name** | **string** |  | [default to null]
**Mode** | **string** | QoS provisioning mode | [optional] [default to null]
**IoSizeBytes** | **int32** | Sets the size of IO for static and capacity limit definitions. The number of IOs per request is obtained by dividing request size by IO size. Default: 64K, Recommended range: 4K - 1M | [optional] [default to null]
**StaticLimits** | [***QosStaticLimits**](QOSStaticLimits.md) |  | [default to null]
**CapacityLimits** | [***QosDynamicLimits**](QOSDynamicLimits.md) |  | [optional] [default to null]
**LimitBy** | **string** | Parameter to limit by | [optional] [default to null]
**PolicyType** | **string** | QOS Policy type - VIEW or USER | [default to null]
**AttachedUsers** | [**[]AttachedUser**](AttachedUser.md) | Attached users for USER QOS Policy | [optional] [default to null]
**TenantId** | **int32** | Tenant ID | [optional] [default to null]
**TenantName** | **string** | Tenant Name | [optional] [default to null]
**IsDefault** | **bool** | Is default User QOS Policy | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

