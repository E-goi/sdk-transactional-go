# {{classname}}

All URIs are relative to *https://slingshot.egoiapp.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWebHook**](WebhooksApi.md#AddWebHook) | **Post** /v2/webhooks | Create new Webhook
[**ListWebHooks**](WebhooksApi.md#ListWebHooks) | **Get** /v2/webhooks | Get All Webhooks
[**RemoveWebHook**](WebhooksApi.md#RemoveWebHook) | **Delete** /v2/webhooks | Remove Webhook

# **AddWebHook**
> WebhookIdResponse AddWebHook(ctx, optional)
Create new Webhook

Create a new Webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhooksApiAddWebHookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiAddWebHookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body1**](Body1.md)| default response | 

### Return type

[**WebhookIdResponse**](WebhookIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebHooks**
> []GetWebhooksResponse ListWebHooks(ctx, optional)
Get All Webhooks

Return all Webhooks defined by the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhooksApiListWebHooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiListWebHooksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **optional.String**| The channel of the groups. Can be &#x60;email&#x60; or &#x60;sms&#x60;, otherwise returns all groups. | 

### Return type

[**[]GetWebhooksResponse**](GetWebhooksResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveWebHook**
> RemoveWebHook(ctx, id, optional)
Remove Webhook

Remove a Webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the contact to be delete. Cannot be empty. | 
 **optional** | ***WebhooksApiRemoveWebHookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiRemoveWebHookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | **optional.String**| The channel of the Webhook. Can be &#x60;email&#x60; or &#x60;sms&#x60;, otherwise removes all the requested webhooks. | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

