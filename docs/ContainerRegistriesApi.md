# \ContainerRegistriesApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContainerRegistry**](ContainerRegistriesApi.md#CreateContainerRegistry) | **Post** /organization/{organizationId}/containerRegistry | Create a container registry
[**DeleteContainerRegistry**](ContainerRegistriesApi.md#DeleteContainerRegistry) | **Delete** /organization/{organizationId}/containerRegistry/{containerRegistryId} | Delete a container registry
[**ListAvailableContainerRegistry**](ContainerRegistriesApi.md#ListAvailableContainerRegistry) | **Get** /organization/{organizationId}/availableContainerRegistry | List supported container registries
[**ListContainerRegistry**](ContainerRegistriesApi.md#ListContainerRegistry) | **Get** /organization/{organizationId}/containerRegistry | List organization container registries



## CreateContainerRegistry

> ContainerRegistryResponse CreateContainerRegistry(ctx, organizationId).ContainerRegistryRequest(containerRegistryRequest).Execute()

Create a container registry

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
    containerRegistryRequest := *openapiclient.NewContainerRegistryRequest("Name_example", openapiclient.ContainerRegistryKindEnum("ECR"), "Url_example", map[string]interface{}{"key": interface{}(123)}) // ContainerRegistryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerRegistriesApi.CreateContainerRegistry(context.Background(), organizationId).ContainerRegistryRequest(containerRegistryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistriesApi.CreateContainerRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainerRegistry`: ContainerRegistryResponse
    fmt.Fprintf(os.Stdout, "Response from `ContainerRegistriesApi.CreateContainerRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerRegistryRequest** | [**ContainerRegistryRequest**](ContainerRegistryRequest.md) |  | 

### Return type

[**ContainerRegistryResponse**](ContainerRegistryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerRegistry

> DeleteContainerRegistry(ctx, organizationId).Execute()

Delete a container registry

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
    resp, r, err := apiClient.ContainerRegistriesApi.DeleteContainerRegistry(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistriesApi.DeleteContainerRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableContainerRegistry

> AvailableContainerRegistryResponse ListAvailableContainerRegistry(ctx, organizationId).Execute()

List supported container registries



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
    resp, r, err := apiClient.ContainerRegistriesApi.ListAvailableContainerRegistry(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistriesApi.ListAvailableContainerRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailableContainerRegistry`: AvailableContainerRegistryResponse
    fmt.Fprintf(os.Stdout, "Response from `ContainerRegistriesApi.ListAvailableContainerRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableContainerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AvailableContainerRegistryResponse**](AvailableContainerRegistryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerRegistry

> ListContainerRegistry200Response ListContainerRegistry(ctx, organizationId).Execute()

List organization container registries

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
    resp, r, err := apiClient.ContainerRegistriesApi.ListContainerRegistry(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistriesApi.ListContainerRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerRegistry`: ListContainerRegistry200Response
    fmt.Fprintf(os.Stdout, "Response from `ContainerRegistriesApi.ListContainerRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListContainerRegistry200Response**](ListContainerRegistry200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

