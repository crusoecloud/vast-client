# IdUpgradeWithoutFileBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableDr** | **bool** | Enables data reduction (DR) for a cluster without DR enabled prior to upgrade | [default to false]
**Force** | **bool** | Forces upgrade regardless of version or upgrade state | [default to false]
**FwUpgrade** | **bool** | Upgrade FWs: BMC, MCU, PCI, NIC | [optional] [default to false]
**Isolcpus** | **bool** | Resets the configuration of isolated CPUs according to a formula | [optional] [default to false]
**VmsOnlyUpgrade** | **bool** | Forces upgrade of VMS only | [optional] [default to false]
**SkipHwCheck** | **bool** | Skips validation of hardware component health. Use with caution since component redundancy is important in NDU. Do not use with OS upgrade. | [optional] [default to false]
**SkipSwValidations** | **bool** | Skips SW validations (e.g. deny list) | [optional] [default to false]
**OsUpgrade** | **bool** | Performs OS upgrade on CNodes and DNodes in addition to upgrading core platform build | [optional] [default to false]
**CnodesBatchSizePercentage** | **int32** | Overrides default percentage of CNodes to upgrade in parallel. Max 50 | [optional] [default to null]
**DnodesBatchSizePercentage** | **float32** | Overrides default percentage of DNodes to upgrade in parallel. Max 37.5. Not relevant during os upgrade | [optional] [default to 20.0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

