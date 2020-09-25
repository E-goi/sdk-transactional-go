# AlertTemplateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | the name to coloquially address the alert. | [optional] [default to null]
**Interval** | **int32** | The interval in seconds between messages to the recipient(s). Maximum interval is 86400 seconds (24 hours). | [default to null]
**MaxAttempts** | **int32** | The maximum amount of attempts to be made. | [default to null]
**TemplateId** | **string** | the id of the message template existent in the system. | [optional] [default to null]
**Subject** | **string** | the subject of the message. only used in email alerts. | [optional] [default to null]
**HtmlBody** | **string** | the html version of the body of the message. only used in email alerts. | [optional] [default to null]
**TextBody** | **string** | the text version of the message. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

