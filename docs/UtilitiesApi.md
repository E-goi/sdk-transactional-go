# {{classname}}

All URIs are relative to *https://slingshot.egoiapp.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailValidator**](UtilitiesApi.md#EmailValidator) | **Post** /v2/emailValidator | Email Validator
[**HtmlToPdf**](UtilitiesApi.md#HtmlToPdf) | **Post** /v2/htmlToPdfConverter | HTML to PDF converter
[**PingApi**](UtilitiesApi.md#PingApi) | **Get** /v2/ping | Ping API
[**SmsLengthCalculator**](UtilitiesApi.md#SmsLengthCalculator) | **Post** /v2/smsLengthCalculator | SMS Length Calculator
[**ValidatePhone**](UtilitiesApi.md#ValidatePhone) | **Get** /v2/validate-phone/{number} | Validate Phone

# **EmailValidator**
> EmailValidatorResponse EmailValidator(ctx, optional)
Email Validator

Validates status and heath of the email address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UtilitiesApiEmailValidatorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesApiEmailValidatorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmailValidatorRequest**](EmailValidatorRequest.md)|  | 

### Return type

[**EmailValidatorResponse**](EmailValidatorResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HtmlToPdf**
> HtmlToPdfResponse HtmlToPdf(ctx, optional)
HTML to PDF converter

Converts an HTML input (either via html string or url) to a PDF returned at base 64. Before using, please validate if your HTML content isdone accordingly to the best standards [HERE](https://validator.w3.org/)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UtilitiesApiHtmlToPdfOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesApiHtmlToPdfOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of HtmlToPdfRequest**](HtmlToPdfRequest.md)|  | 

### Return type

[**HtmlToPdfResponse**](HtmlToPdfResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PingApi**
> PingResponse PingApi(ctx, )
Ping API

Send an innocuous request to the API, in order to confirm its working as intended 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PingResponse**](PingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmsLengthCalculator**
> SmsLengthCalculatorResponse SmsLengthCalculator(ctx, optional)
SMS Length Calculator

Calculates the sizes of an SMS text

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UtilitiesApiSmsLengthCalculatorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesApiSmsLengthCalculatorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SmsLengthCalculatorRequest**](SmsLengthCalculatorRequest.md)|  | 

### Return type

[**SmsLengthCalculatorResponse**](SmsLengthCalculatorResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidatePhone**
> ValidatePhoneResponse ValidatePhone(ctx, number)
Validate Phone

Checks if a phone number is valid. Can be either land line or mobile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **number** | **string**| The phone number to validate. must be in format &#x60;&#x60;countryCode-mobileNumber&#x60;&#x60; | 

### Return type

[**ValidatePhoneResponse**](ValidatePhoneResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

