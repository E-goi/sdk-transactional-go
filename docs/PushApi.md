# {{classname}}

All URIs are relative to *https://slingshot.egoiapp.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApp**](PushApi.md#AddApp) | **Post** /v2/apps | Create App
[**AddPushTemplate**](PushApi.md#AddPushTemplate) | **Post** /v2/push/templates | Create Push Template
[**DeleteApp**](PushApi.md#DeleteApp) | **Delete** /v2/apps/{id} | Remove App
[**DeletePushTemplate**](PushApi.md#DeletePushTemplate) | **Delete** /v2/push/templates/{id} | Remove Push Template
[**ListApp**](PushApi.md#ListApp) | **Get** /v2/apps/{id} | Get an App
[**ListApps**](PushApi.md#ListApps) | **Get** /v2/apps | Get All Apps
[**ListPushTemplate**](PushApi.md#ListPushTemplate) | **Get** /v2/push/templates/{id} | Get an Push Template
[**ListPushTemplates**](PushApi.md#ListPushTemplates) | **Get** /v2/push/templates | Get All Push Templates
[**SendPushMessages**](PushApi.md#SendPushMessages) | **Post** /v2/push/messages/action/send | Send a Push Message
[**UpdateApp**](PushApi.md#UpdateApp) | **Put** /v2/apps/{id} | Update App
[**UpdatePushTemplate**](PushApi.md#UpdatePushTemplate) | **Patch** /v2/push/templates/{id} | Update Push Template

# **AddApp**
> TemplateIdResponse AddApp(ctx, optional)
Create App

Creates a new App

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PushApiAddAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PushApiAddAppOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MobileAppRequest**](MobileAppRequest.md)|  | 

### Return type

[**TemplateIdResponse**](TemplateIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPushTemplate**
> TemplateIdResponse AddPushTemplate(ctx, optional)
Create Push Template

Creates a new template record for Push

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PushApiAddPushTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PushApiAddPushTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PushTemplateRequest**](PushTemplateRequest.md)|  | 

### Return type

[**TemplateIdResponse**](TemplateIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApp**
> DeleteApp(ctx, id)
Remove App

Remove App

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the app. Cannot be empty. | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePushTemplate**
> DeletePushTemplate(ctx, id)
Remove Push Template

Remove Push Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the Push Template. Cannot be empty. | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApp**
> MobileAppIdResponse ListApp(ctx, id)
Get an App

Gets an App

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the App. Cannot be empty. | 

### Return type

[**MobileAppIdResponse**](MobileAppIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApps**
> []MobileAppResponse ListApps(ctx, )
Get All Apps

Obtains all the Apps

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]MobileAppResponse**](MobileAppResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPushTemplate**
> SinglePushTemplateResponse ListPushTemplate(ctx, id)
Get an Push Template

Gets an Push Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the Push Template. Cannot be empty. | 

### Return type

[**SinglePushTemplateResponse**](SinglePushTemplateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPushTemplates**
> []SinglePushTemplateResponse ListPushTemplates(ctx, )
Get All Push Templates

Obtains all the Push templates

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SinglePushTemplateResponse**](SinglePushTemplateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendPushMessages**
> []SendMessageIdResponse SendPushMessages(ctx, optional)
Send a Push Message

Send an push message to a recipient

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PushApiSendPushMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PushApiSendPushMessagesOpts struct
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

# **UpdateApp**
> MobileAppIdResponse UpdateApp(ctx, id, optional)
Update App

Updates an App, given a template id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the Push Template. Cannot be empty. | 
 **optional** | ***PushApiUpdateAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PushApiUpdateAppOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MobileAppRequest**](MobileAppRequest.md)|  | 

### Return type

[**MobileAppIdResponse**](MobileAppIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePushTemplate**
> SinglePushTemplateResponse UpdatePushTemplate(ctx, id, optional)
Update Push Template

Updates an Push template, given a template id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the Push Template. Cannot be empty. | 
 **optional** | ***PushApiUpdatePushTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PushApiUpdatePushTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PushTemplateRequest**](PushTemplateRequest.md)|  | 

### Return type

[**SinglePushTemplateResponse**](SinglePushTemplateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

