# \ContainerDependencyApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContainerDependency**](ContainerDependencyApi.md#CreateContainerDependency) | **Post** /container/{containerId}/dependency/{targetContainerId} | Add container dependency to this application.
[**ListContainerDependency**](ContainerDependencyApi.md#ListContainerDependency) | **Get** /container/{containerId}/dependency | List container dependencies
[**RemoveContainerDependency**](ContainerDependencyApi.md#RemoveContainerDependency) | **Delete** /container/{containerId}/dependency/{targetContainerId} | Remove container dependency to this container.



## CreateContainerDependency

> ContainerResponse CreateContainerDependency(ctx, containerId, targetContainerId).Execute()

Add container dependency to this application.



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    targetContainerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Target container ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerDependencyApi.CreateContainerDependency(context.Background(), containerId, targetContainerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerDependencyApi.CreateContainerDependency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainerDependency`: ContainerResponse
    fmt.Fprintf(os.Stdout, "Response from `ContainerDependencyApi.CreateContainerDependency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**targetContainerId** | **string** | Target container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContainerResponse**](ContainerResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerDependency

> ContainerResponseList ListContainerDependency(ctx, containerId).Execute()

List container dependencies

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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerDependencyApi.ListContainerDependency(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerDependencyApi.ListContainerDependency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerDependency`: ContainerResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerDependencyApi.ListContainerDependency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerResponseList**](ContainerResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveContainerDependency

> RemoveContainerDependency(ctx, containerId).Execute()

Remove container dependency to this container.

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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerDependencyApi.RemoveContainerDependency(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerDependencyApi.RemoveContainerDependency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveContainerDependencyRequest struct via the builder pattern


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

