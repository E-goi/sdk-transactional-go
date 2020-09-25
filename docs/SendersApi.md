# {{classname}}

All URIs are relative to *https://slingshot.egoiapp.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmailSenders**](SendersApi.md#GetEmailSenders) | **Get** /v2/email/senders | Get All Email Senders
[**GetSmsSenders**](SendersApi.md#GetSmsSenders) | **Get** /v2/sms/senders | Get All Sms Senders

# **GetEmailSenders**
> []EmailSenderResponse GetEmailSenders(ctx, )
Get All Email Senders

Retrieve the list of available Email senders

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EmailSenderResponse**](EmailSenderResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmsSenders**
> []SmsSenderResponse GetSmsSenders(ctx, )
Get All Sms Senders

Return all Sms Senders defined by the user

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SmsSenderResponse**](SmsSenderResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

