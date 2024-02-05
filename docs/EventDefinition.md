# EventDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**ObjectType** | **string** |  | [optional] [default to null]
**EventType** | **string** |  | [optional] [default to null]
**Severity** | **string** | The severity of the alarm | [optional] [default to null]
**TriggerOn** | **string** | For &#x27;Object Modified&#x27; alarms: a list of values | For &#x27;Threshold/Rate&#x27; alarms: a list of 2 members. The first is an operator like gt/ge/lte and the second is a number  | [optional] [default to null]
**TriggerOff** | **string** | For &#x27;Object Modified&#x27; alarms: a list of values | [optional] [default to null]
**UserModified** | **bool** | Did a user modify this event definition | [optional] [default to null]
**EmailRecipients** | **[]string** | List of emails you want to notify in case this event occurs | [optional] [default to null]
**WebhookUrl** | **string** | The URL that the webhook will go to | [optional] [default to null]
**WebhookData** | **string** | Use $event as event message parameter | [optional] [default to null]
**WebhookMethod** | **string** | The method that the webhook will use | [optional] [default to null]
**Property** | **string** |  | [optional] [default to null]
**WebhookParams** | **string** |  | [optional] [default to null]
**AlarmDefinitions** | **string** |  | [optional] [default to null]
**ActionDefinitions** | **string** |  | [optional] [default to null]
**EventMessage** | **string** |  | [optional] [default to null]
**DisableActions** | **bool** |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**Internal** | **bool** |  | [optional] [default to null]
**AlarmOnly** | **bool** | When this is enabled, only alarms will lead to email and webhook actions | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

