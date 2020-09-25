# TemplateEmail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | the Domain to de used in the sender email. | [optional] [default to null]
**SenderId** | **string** | The code identification of the sender profile to be used. | [default to null]
**SenderName** | **string** | The name to be associated with the Domain to form the email. | [optional] [default to null]
**To** | **[]string** | The To field in the email. | [default to null]
**Cc** | **[]string** | The Cc field in the email. | [optional] [default to null]
**Bcc** | **[]string** | The Bcc field in the email. | [optional] [default to null]
**ReplyToId** | **string** | The email id for if/when the recipient responds to the email. | [optional] [default to null]
**TemplateId** | **string** | The id of the template to use. | [default to null]
**Attachment** | [**[]AttachmentRequest**](AttachmentRequest.md) | A list of files to be attached to the emails | [optional] [default to null]
**MergeTags** | **[]string** | A list of the merge tags to replace information gaps in the email. The format of each tag should be &#x60;&#x60;tagName:tagValue&#x60;&#x60; | [optional] [default to null]
**CustomData** | **string** | Extra Data related to the campaign to be re-sent via WebHook. | [optional] [default to null]
**Priority** | **string** | The priority of the message. | [optional] [default to PRIORITY.NON_URGENT]
**Registered** | **bool** | option to register the message. Its an exclusive feature for senders that require it. | [optional] [default to false]
**Header** | [***HeaderOptionsRequest**](HeaderOptionsRequest.md) |  | [optional] [default to null]
**Group** | **string** | The group of the message. | [optional] [default to default]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

