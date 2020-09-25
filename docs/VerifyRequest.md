# VerifyRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderId** | **string** | The id of the sender. | [default to null]
**To** | **string** | The phone number of the recipient of the code. Format should be &#x60;countryCode&#x60;-&#x60;phoneNumber&#x60; | [default to null]
**Message** | **string** | the message that the user is going to receive. location of the code should be given by using the tag {{ code }}. All instances of {{ code }} will be replaced with a code that is in accordance with the settings. If the tag is not present, the code will be added at the end of the message. | [optional] [default to null]
**CodeLength** | **int32** | the length of the code in characters. Maximum is 20 characters. Default is 6 characters. | [optional] [default to 6]
**CodeValidity** | **int32** | the period of time the code is valid after the message has been sent, in seconds. Defaults to 600 seconds (10 min). Maximum validity is 3600 seconds (1 hour). | [optional] [default to 600]
**CodeType** | **string** | the format of the code. can be:  Format | Description | RegEx --- | --- | --- numeric | numbers | [0-9] alpha | letters | [A-Z] alphanumeric | numbers and letters | [0-9A-Z] hexa | numbers and letters up to F | [0-9A-F] | [optional] [default to CODE_TYPE.NUMERIC]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

