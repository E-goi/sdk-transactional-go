# {{classname}}

All URIs are relative to *https://slingshot.egoiapp.com/api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDomain**](DomainsApi.md#AddDomain) | **Post** /v2/domain | Create New Domain
[**ChangeDefaultDomain**](DomainsApi.md#ChangeDefaultDomain) | **Post** /v2/domain/default | Change Default Domain
[**ListDefaultDomain**](DomainsApi.md#ListDefaultDomain) | **Get** /v2/domain/default | Get Default Domain
[**ListDomain**](DomainsApi.md#ListDomain) | **Get** /v2/domain/{id} | Get Single Domain
[**ListDomains**](DomainsApi.md#ListDomains) | **Get** /v2/domain | Get All Domains
[**ValidateDomain**](DomainsApi.md#ValidateDomain) | **Get** /v2/domain/{id}/validate | Validate Single Domain

# **AddDomain**
> DomainResponse AddDomain(ctx, optional)
Create New Domain

Add a new Domain to the list of available Domains for the sender to the recipient.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DomainsApiAddDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainsApiAddDomainOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DomainRequest**](DomainRequest.md)|  | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeDefaultDomain**
> DomainResponse ChangeDefaultDomain(ctx, optional)
Change Default Domain

Change the currently configured default Domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DomainsApiChangeDefaultDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainsApiChangeDefaultDomainOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DefaultDomainRequest**](DefaultDomainRequest.md)|  | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDefaultDomain**
> DomainResponse ListDefaultDomain(ctx, )
Get Default Domain

Retrieve the currently configured default Domain. If no default domain is set, the list will be empty.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDomain**
> DomainResponse ListDomain(ctx, id)
Get Single Domain

Retrieve the Domain with the id in the path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the Verify Request | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDomains**
> []DomainResponse ListDomains(ctx, )
Get All Domains

Retrieve the list of all Domains available for the sender to the recipients.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]DomainResponse**](DomainResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateDomain**
> ValidateDomainResponse ValidateDomain(ctx, id)
Validate Single Domain

Clears the validation cache and validates the supplied configured Domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the Domain | 

### Return type

[**ValidateDomainResponse**](ValidateDomainResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

