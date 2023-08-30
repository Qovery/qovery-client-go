# \OrganizationApiTokenApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationApiToken**](OrganizationApiTokenApi.md#CreateOrganizationApiToken) | **Post** /organization/{organizationId}/apiToken | Create an organization api token
[**DeleteOrganizationApiToken**](OrganizationApiTokenApi.md#DeleteOrganizationApiToken) | **Delete** /organization/{organizationId}/apiToken/{apiTokenId} | Delete organization api token
[**ListOrganizationApiTokens**](OrganizationApiTokenApi.md#ListOrganizationApiTokens) | **Get** /organization/{organizationId}/apiToken | List organization api tokens



## CreateOrganizationApiToken

> OrganizationApiTokenCreate CreateOrganizationApiToken(ctx, organizationId).OrganizationApiTokenCreateRequest(organizationApiTokenCreateRequest).Execute()

Create an organization api token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    organizationApiTokenCreateRequest := *openapiclient.NewOrganizationApiTokenCreateRequest("Name_example", openapiclient.OrganizationApiTokenScope("ADMIN")) // OrganizationApiTokenCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApiTokenApi.CreateOrganizationApiToken(context.Background(), organizationId).OrganizationApiTokenCreateRequest(organizationApiTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApiTokenApi.CreateOrganizationApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApiToken`: OrganizationApiTokenCreate
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApiTokenApi.CreateOrganizationApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationApiTokenCreateRequest** | [**OrganizationApiTokenCreateRequest**](OrganizationApiTokenCreateRequest.md) |  | 

### Return type

[**OrganizationApiTokenCreate**](OrganizationApiTokenCreate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationApiToken

> DeleteOrganizationApiToken(ctx, organizationId, apiTokenId).Execute()

Delete organization api token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    apiTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization Api Token ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApiTokenApi.DeleteOrganizationApiToken(context.Background(), organizationId, apiTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApiTokenApi.DeleteOrganizationApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**apiTokenId** | **string** | Organization Api Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationApiTokens

> OrganizationApiTokenResponseList ListOrganizationApiTokens(ctx, organizationId).Execute()

List organization api tokens



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationApiTokenApi.ListOrganizationApiTokens(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApiTokenApi.ListOrganizationApiTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizationApiTokens`: OrganizationApiTokenResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApiTokenApi.ListOrganizationApiTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationApiTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationApiTokenResponseList**](OrganizationApiTokenResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

