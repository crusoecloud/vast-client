# PartialUserQuota

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotaSystemId** | **int32** |  | [optional] [default to null]
**SoftLimit** | **int32** | Soft quota limit | [optional] [default to null]
**HardLimit** | **int32** | Hard quota limit | [optional] [default to null]
**HardLimitInodes** | **int32** | Hard inodes quota limit | [optional] [default to null]
**SoftLimitInodes** | **int32** | Soft inodes quota limit | [optional] [default to null]
**GracePeriod** | **string** | Quota enforcement grace period in seconds, minutes, hours or days. Example: 90m | [optional] [default to null]
**TimeToBlock** | **string** | Grace period expiration time | [optional] [default to null]
**Entity** | [***PartialQuotaEntityInfo**](PartialQuotaEntityInfo.md) |  | [optional] [default to null]
**IsAccountable** | **bool** |  | [optional] [default to null]
**State** | **string** |  | [optional] [default to null]
**UsedCapacity** | **int64** | Used capacity in bytes | [optional] [default to null]
**UsedInodes** | **int32** | Used inodes | [optional] [default to null]
**PercentInodes** | **int32** |  | [optional] [default to null]
**PercentCapacity** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

