# {{classname}}

All URIs are relative to *https://slingshot.egoiapp.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelVerify**](VerifyApi.md#CancelVerify) | **Post** /v2/verify/{id}/cancel | Cancel Verify Request
[**GetVerify**](VerifyApi.md#GetVerify) | **Get** /v2/verify/{id} | Get Verify Request
[**RequestVerify**](VerifyApi.md#RequestVerify) | **Post** /v2/verify/request | Request Verify
[**ResendVerify**](VerifyApi.md#ResendVerify) | **Post** /v2/verify/{id}/resend | Resend Verify Code
[**ValidateVerifyCode**](VerifyApi.md#ValidateVerifyCode) | **Post** /v2/verify/{id}/validate | Validate Verify Code

# **CancelVerify**
> []VerifyIdResponse CancelVerify(ctx, id)
Cancel Verify Request

This method will cancel the requested verify. No codes will be further generated or confirmed within the context of the id of the message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the Verify Request | 

### Return type

[**[]VerifyIdResponse**](VerifyIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVerify**
> CheckVerifyResponse GetVerify(ctx, id)
Get Verify Request

This method gets a verify requests, including status of it

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the Verify Request | 

### Return type

[**CheckVerifyResponse**](CheckVerifyResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestVerify**
> VerifyIdResponse RequestVerify(ctx, optional)
Request Verify

This method will generate a code and send a message with it to a recipient

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VerifyApiRequestVerifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerifyApiRequestVerifyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VerifyRequest**](VerifyRequest.md)|  | 

### Return type

[**VerifyIdResponse**](VerifyIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResendVerify**
> []VerifyIdResponse ResendVerify(ctx, id)
Resend Verify Code

This method will resend the requested verify. A new code will be generated and sent to the recipient

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the Verify Request | 

### Return type

[**[]VerifyIdResponse**](VerifyIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateVerifyCode**
> []VerifyIdResponse ValidateVerifyCode(ctx, id, optional)
Validate Verify Code

This method validates the code submitted by the recipient. This request should be made by the sender of the Verify request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the Verify Request | 
 **optional** | ***VerifyApiValidateVerifyCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerifyApiValidateVerifyCodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ValidateCodeVerifyRequest**](ValidateCodeVerifyRequest.md)|  | 

### Return type

[**[]VerifyIdResponse**](VerifyIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

