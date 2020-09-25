# {{classname}}

All URIs are relative to *https://slingshot.egoiapp.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMultiChannelFlow**](MultiChannelApi.md#CreateMultiChannelFlow) | **Post** /v2/multi-channel/flow | Create Multi-Channel Flow
[**ExecuteMultiChannel**](MultiChannelApi.md#ExecuteMultiChannel) | **Post** /v2/multi-channel/execute | Execute Multi-Channel
[**ListMultiChannel2**](MultiChannelApi.md#ListMultiChannel2) | **Get** /v2/multi-channel/flow/{id} | Get Single Multi-Channel Flow
[**ListMultiChannels**](MultiChannelApi.md#ListMultiChannels) | **Get** /v2/multi-channel/flow | Get All Multi-Channel Flows
[**RemoveMultiChannel**](MultiChannelApi.md#RemoveMultiChannel) | **Delete** /v2/multi-channel/flow/{id} | Remove Single Multi-Channel

# **CreateMultiChannelFlow**
> MultiChannelFlowIdResponse CreateMultiChannelFlow(ctx, optional)
Create Multi-Channel Flow

This method will create and add a Multi-Channel Flow. First message doesn't need interval or condition, since it doesn't depend on other messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MultiChannelApiCreateMultiChannelFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MultiChannelApiCreateMultiChannelFlowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MultiChannelFlowRequest**](MultiChannelFlowRequest.md)|  | 

### Return type

[**MultiChannelFlowIdResponse**](MultiChannelFlowIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteMultiChannel**
> MultiChannelIdResponse ExecuteMultiChannel(ctx, optional)
Execute Multi-Channel

This method will create and start a flow, sending the Multi-Channel of messages to the appropriate recipient

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MultiChannelApiExecuteMultiChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MultiChannelApiExecuteMultiChannelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ExecuteMultiChannelRequest**](ExecuteMultiChannelRequest.md)|  | 

### Return type

[**MultiChannelIdResponse**](MultiChannelIdResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMultiChannel2**
> []MultiChannelFlowResponse ListMultiChannel2(ctx, id)
Get Single Multi-Channel Flow

This method returns the Multi-Channel Flow wih the id in the path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the Multi-Channel Flow | 

### Return type

[**[]MultiChannelFlowResponse**](MultiChannelFlowResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMultiChannels**
> []MultiChannelFlowResponse ListMultiChannels(ctx, )
Get All Multi-Channel Flows

This method returns all Multi-Channels Flows set in the system

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]MultiChannelFlowResponse**](MultiChannelFlowResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveMultiChannel**
> RemoveMultiChannel(ctx, id)
Remove Single Multi-Channel

This method removes a Multi-Channel Flow from the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the Multi-Channel Flow to remove | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

