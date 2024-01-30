# QospoliciesBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to null]
**Mode** | **string** | QoS provisioning mode | [optional] [default to null]
**StaticLimits** | [***RequestQosStaticLimits**](RequestQOSStaticLimits.md) |  | [optional] [default to null]
**CapacityLimits** | [***RequestQosDynamicLimits**](RequestQOSDynamicLimits.md) |  | [optional] [default to null]
**LimitBy** | **string** | Parameter to limit by | [optional] [default to null]
**PolicyType** | **string** | QOS Policy type - VIEW or USER | [optional] [default to null]
**AttachedUsers** | [***interface{}**](interface{}.md) | Attached users for USER QOS Policy | [optional] [default to null]
**TenantId** | **int32** | Tenant ID | [optional] [default to null]
**IsDefault** | **bool** | Is default User QOS Policy | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

