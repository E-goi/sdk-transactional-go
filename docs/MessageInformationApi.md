# {{classname}}

All URIs are relative to *https://slingshot.egoiapp.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllEmailMessages**](MessageInformationApi.md#ListAllEmailMessages) | **Get** /v2/email/messages | Get All Email Messages
[**ListAllSmsMessages**](MessageInformationApi.md#ListAllSmsMessages) | **Get** /v2/sms/messages | Get All Sms Messages
[**ListSingleEmailMessage**](MessageInformationApi.md#ListSingleEmailMessage) | **Get** /v2/email/messages/{id} | Get Single Email Messages
[**ListSingleSmsMessage**](MessageInformationApi.md#ListSingleSmsMessage) | **Get** /v2/sms/messages/{id} | Get Single Sms Messages

# **ListAllEmailMessages**
> []PartialEmailMessageResponse ListAllEmailMessages(ctx, optional)
Get All Email Messages

Retrieve the list of all sms messages. Results can be filtered by messages ids, group or channel. The result will be a combination of all parameters present

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessageInformationApiListAllEmailMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessageInformationApiListAllEmailMessagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **optional.String**| The name of the group to search | 
 **dateStart** | **optional.String**| The date to start the report. Date must be in format &#x60;&#x60;yyyy-MM-dd&#x60;&#x60; | 
 **dateEnd** | **optional.String**| The date to end the report. Date should be in format &#x60;&#x60;yyyy-MM-dd&#x60;&#x60; | 
 **destination** | **optional.String**| The mobile number of a recipient. | 
 **status** | **optional.String**| The status to check the message. | 

### Return type

[**[]PartialEmailMessageResponse**](PartialEmailMessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllSmsMessages**
> []PartialSmsMessageResponse ListAllSmsMessages(ctx, optional)
Get All Sms Messages

Retrieve the list of all sms messages. Results can be filtered by messages ids, group or channel. The result will be a combination of all parameters present

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessageInformationApiListAllSmsMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessageInformationApiListAllSmsMessagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **optional.String**| The name of the group to search | 
 **dateStart** | **optional.String**| The date to start the report. Date must be in format &#x60;&#x60;yyyy-MM-dd&#x60;&#x60; | 
 **dateEnd** | **optional.String**| The date to end the report. Date should be in format &#x60;&#x60;yyyy-MM-dd&#x60;&#x60; | 
 **destination** | **optional.String**| The mobile number of a recipient. | 
 **status** | **optional.String**| The status to check the message. | 

### Return type

[**[]PartialSmsMessageResponse**](PartialSmsMessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSingleEmailMessage**
> CompleteEmailMessageResponse ListSingleEmailMessage(ctx, id)
Get Single Email Messages

Retrieve the message with the id from path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The id of the message. | 

### Return type

[**CompleteEmailMessageResponse**](CompleteEmailMessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSingleSmsMessage**
> CompleteSmsMessageResponse ListSingleSmsMessage(ctx, id)
Get Single Sms Messages

Retrieve the message with the id from path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The id of the message. | 

### Return type

[**CompleteSmsMessageResponse**](CompleteSmsMessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

