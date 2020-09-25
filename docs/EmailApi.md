# {{classname}}

All URIs are relative to *https://slingshot.egoiapp.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendEmailMessages**](EmailApi.md#SendEmailMessages) | **Post** /v2/email/messages/action/send | Send an Email Message

# **SendEmailMessages**
> []SendMessageIdResponse SendEmailMessages(ctx, optional)
Send an Email Message

Send an email message to a recipient

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmailApiSendEmailMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmailApiSendEmailMessagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []interface{}**](interface{}.md)| default response | 

### Return type

[**[]SendMessageIdResponse**](SendMessageIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

