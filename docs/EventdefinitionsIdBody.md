# EventdefinitionsIdBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severity** | **string** | The severity of an alarm triggered by this event. INFO means no alarm is triggered. | [optional] [default to null]
**TriggerOn** | **string** | For &#x27;Object Modified&#x27; alarms: a list of values | For &#x27;Threshold/Rate&#x27; alarms: a list of 2 members. The first is an operator like gt/ge/lte and the second is a number | [optional] [default to null]
**TriggerOff** | **string** | For &#x27;Object Modified&#x27; alarms: a list of values | [optional] [default to null]
**TimeFrame** | **string** | For rate alarms, the The time frame over which to monitor the property. | [optional] [default to null]
**EmailRecipients** | **[]string** | List of emails you want to notify in case this alarm occurs | [optional] [default to null]
**WebhookUrl** | **string** | The URL of the API endpoint of an external application to trigger on alarms. | [optional] [default to null]
**WebhookMethod** | **string** | The HTTP method to invoke with the webhook trigger. | [optional] [default to null]
**WebhookData** | **string** | The payload, if required, to send with a POST command. You can use the $event variable to include the event message. | [optional] [default to null]
**WebhookParams** | **string** | The URL parameters to send with a GET command. | [optional] [default to null]
**DisableActions** | **bool** | Set to true to disable alert actions for the event definition. | [optional] [default to null]
**Enabled** | **bool** | Set to true to enable events, alarms and actions. | [optional] [default to null]
**Internal** | **bool** |  | [optional] [default to null]
**Cooldown** | **int32** | Minimal time to wait between two consecutive events | [optional] [default to null]
**RaiseAtCount** | **int32** | Raise an alarm after a specific number of recurrences | [optional] [default to null]
**AlarmOnly** | **bool** | Set to true for only alarms to trigger configured actions such as email and webhook. Set to false for all events of the definition to trigger the configured actions. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

