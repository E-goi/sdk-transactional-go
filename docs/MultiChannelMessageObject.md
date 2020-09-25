# MultiChannelMessageObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Available communications channels in the API. Currently there are Emails and Sms. | [default to null]
**Interval** | **int32** | the interval in seconds between the message and the previous. Defaults to 3600s (1h). | [optional] [default to null]
**Condition** | **string** | The condition of the previous message to trigger the next one. In list of messages in the plan the condition of the first message is discarded (doesnt have previous condition). Can be:  Name | Description | Channel --- | --- | --- none | no condition of previous message | email, sms not_open | previous message hasnt been opened | email open | previous message has been opened | email click | previous message has a link that has been clicked | email not_click | previous message has a link that has not been clicked | email | [optional] [default to null]
**TemplateId** | **int32** | the template id already in the system. overrides subject, textBody and htmlBody. | [optional] [default to null]
**Subject** | **string** | The subject of the message. Only applies to email messages. | [optional] [default to null]
**TextBody** | **string** | The text body of the message. | [optional] [default to null]
**HtmlBody** | **string** | The html body of the message. Only applies to email messages. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
