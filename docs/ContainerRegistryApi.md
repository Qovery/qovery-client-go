# \ContainerRegistryApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditContainerRegistry**](ContainerRegistryApi.md#EditContainerRegistry) | **Put** /organization/{organizationId}/containerRegistry/{containerRegistryId} | Edit a container registry



## EditContainerRegistry

> ContainerRegistryResponse EditContainerRegistry(ctx, organizationId).ContainerRegistryRequest(containerRegistryRequest).Execute()

Edit a container registry

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
    resp, r, err := apiClient.ContainerRegistryApi.EditContainerRegistry(context.Background(), organizationId).ContainerRegistryRequest(containerRegistryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistryApi.EditContainerRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditContainerRegistry`: ContainerRegistryResponse
    fmt.Fprintf(os.Stdout, "Response from `ContainerRegistryApi.EditContainerRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditContainerRegistryRequest struct via the builder pattern


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

