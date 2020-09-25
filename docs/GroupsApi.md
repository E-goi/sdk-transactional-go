# {{classname}}

All URIs are relative to *https://slingshot.egoiapp.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGroups**](GroupsApi.md#ListGroups) | **Get** /v2/groups | Get All Groups

# **ListGroups**
> []GroupResponse ListGroups(ctx, optional)
Get All Groups

Retrieve all groups to sort messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiListGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiListGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **optional.String**| The channel of the groups. Can be &#x60;email&#x60; or &#x60;sms&#x60;, otherwise returns all groups. | 

### Return type

[**[]GroupResponse**](GroupResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

