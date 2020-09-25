# {{classname}}

All URIs are relative to *https://slingshot.egoiapp.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeAlertTemplate**](AlertsApi.md#ChangeAlertTemplate) | **Patch** /v2/alert/template/{id} | Change Single Alert Template
[**CreateAlertTemplate**](AlertsApi.md#CreateAlertTemplate) | **Post** /v2/alert/template | Create Alert Template
[**ExecuteEmailAlert**](AlertsApi.md#ExecuteEmailAlert) | **Post** /v2/alert/action/execute | Execute Alert
[**ListMultiChannel**](AlertsApi.md#ListMultiChannel) | **Get** /v2/alert/template/{id} | Get Single Alert Template
[**ListMultiChannel1**](AlertsApi.md#ListMultiChannel1) | **Get** /v2/alert/template | Get All Alert Template
[**RemoveAlertTemplate**](AlertsApi.md#RemoveAlertTemplate) | **Delete** /v2/alert/template/{id} | Remove Alert Template
[**StopAlert**](AlertsApi.md#StopAlert) | **Post** /v2/alert/action/stop | Stop Alert

# **ChangeAlertTemplate**
> AlertTemplateResponse ChangeAlertTemplate(ctx, id, optional)
Change Single Alert Template

This method will change information about the template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the Alert Template to edit | 
 **optional** | ***AlertsApiChangeAlertTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertsApiChangeAlertTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AlertTemplateRequest**](AlertTemplateRequest.md)|  | 

### Return type

[**AlertTemplateResponse**](AlertTemplateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAlertTemplate**
> AlertTemplateIdResponse CreateAlertTemplate(ctx, optional)
Create Alert Template

This method will create and add an alert template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertsApiCreateAlertTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertsApiCreateAlertTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AlertTemplateRequest**](AlertTemplateRequest.md)|  | 

### Return type

[**AlertTemplateIdResponse**](AlertTemplateIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteEmailAlert**
> AlertIdResponse ExecuteEmailAlert(ctx, optional)
Execute Alert

This method executes an alert flow , using a pre-created alert template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertsApiExecuteEmailAlertOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertsApiExecuteEmailAlertOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body**](Body.md)| default response | 

### Return type

[**AlertIdResponse**](AlertIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMultiChannel**
> AlertTemplateResponse ListMultiChannel(ctx, id)
Get Single Alert Template

This method returns the Alert Template wih the id in the path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the Alert Template | 

### Return type

[**AlertTemplateResponse**](AlertTemplateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMultiChannel1**
> []AlertTemplateResponse ListMultiChannel1(ctx, )
Get All Alert Template

This method returns all the Alert Templates

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AlertTemplateResponse**](AlertTemplateResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAlertTemplate**
> RemoveAlertTemplate(ctx, id)
Remove Alert Template

This method removes an Alert Template from the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the Alert Template to remove | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopAlert**
> AlertIdResponse StopAlert(ctx, optional)
Stop Alert

This method will stop any alerts that are being processed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertsApiStopAlertOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertsApiStopAlertOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CancelAlertRequest**](CancelAlertRequest.md)|  | 

### Return type

[**AlertIdResponse**](AlertIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

