# {{classname}}

All URIs are relative to *https://slingshot.egoiapp.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProduceEmailReport**](ReportsApi.md#ProduceEmailReport) | **Get** /v2/email/reports | Get Email Report
[**ProduceSmsReport**](ReportsApi.md#ProduceSmsReport) | **Get** /v2/sms/reports | Get Sms Report

# **ProduceEmailReport**
> []EmailReportResponse ProduceEmailReport(ctx, optional)
Get Email Report

Retrieve Mail Reports and Statistics. Entries can be customizable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsApiProduceEmailReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiProduceEmailReportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **optional.String**| The name of the group to search | 
 **groupBy** | **optional.String**| The option to group the results. Can be either &#x60;&#x60;day&#x60;&#x60; or &#x60;&#x60;month&#x60;&#x60;. Defaults to &#x60;&#x60;day&#x60;&#x60; if none are present | 
 **dateStart** | **optional.String**| The date to start the report. Date must be in format &#x60;&#x60;yyyy-MM-dd&#x60;&#x60; | 
 **dateEnd** | **optional.String**| The date to end the report. Date should be in format &#x60;&#x60;yyyy-MM-dd&#x60;&#x60; | 

### Return type

[**[]EmailReportResponse**](EmailReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProduceSmsReport**
> []SmsReportResponse ProduceSmsReport(ctx, optional)
Get Sms Report

Retrieve Sms Reports and Statistics. Entries can be customizable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsApiProduceSmsReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiProduceSmsReportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **optional.String**| The id of the group to search | 
 **groupBy** | **optional.String**| The option to group the results. Can be either &#x60;&#x60;day&#x60;&#x60; or &#x60;&#x60;month&#x60;&#x60;. Defaults to &#x60;&#x60;day&#x60;&#x60; if none are present | 
 **dateStart** | **optional.String**| The date to start the report. Date must be in format &#x60;&#x60;yyyy-MM-dd&#x60;&#x60; | 
 **dateEnd** | **optional.String**| The date to end the report. Date should be in format &#x60;&#x60;yyyy-MM-dd&#x60;&#x60; | 

### Return type

[**[]SmsReportResponse**](SmsReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

