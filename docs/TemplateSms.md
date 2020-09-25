# TemplateSms

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **[]string** |  | [default to null]
**From** | **string** | The id of the sender. | [default to null]
**TemplateId** | **string** | The id of the template. | [default to null]
**MergeTags** | **[]string** | A list of the merge tags to replace information gaps in the email. The format of each tag should be &#x60;&#x60;tagName:tagValue&#x60;&#x60; | [optional] [default to null]
**Priority** | **string** | The priority of the message. | [optional] [default to PRIORITY.NON_URGENT]
**CustomData** | **string** | Extra Data related to the campaign to be re-sent via WebHook. | [optional] [default to null]
**Registered** | **bool** | option to register the message. Its an exclusive feature for senders that require it. | [optional] [default to false]
**Group** | **string** | The group of the message. | [optional] [default to default]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

