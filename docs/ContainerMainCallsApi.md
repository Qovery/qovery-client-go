# \ContainerMainCallsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteContainer**](ContainerMainCallsApi.md#DeleteContainer) | **Delete** /container/{containerId} | Delete container
[**EditContainer**](ContainerMainCallsApi.md#EditContainer) | **Put** /container/{containerId} | Edit container
[**GetContainer**](ContainerMainCallsApi.md#GetContainer) | **Get** /container/{containerId} | Get container by ID
[**GetContainerStatus**](ContainerMainCallsApi.md#GetContainerStatus) | **Get** /container/{containerId}/status | Get container status
[**ListContainerLinks**](ContainerMainCallsApi.md#ListContainerLinks) | **Get** /container/{containerId}/link | List all URLs of the container



## DeleteContainer

> DeleteContainer(ctx, containerId).Execute()

Delete container



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
    resp, r, err := apiClient.ContainerMainCallsApi.DeleteContainer(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerMainCallsApi.DeleteContainer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteContainerRequest struct via the builder pattern


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


## EditContainer

> ContainerResponse EditContainer(ctx, containerId).ContainerEditRequest(containerEditRequest).Execute()

Edit container



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
    containerEditRequest := *openapiclient.NewContainerEditRequest("Name_example", "RegistryId_example", "ImageName_example", "Tag_example") // ContainerEditRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerMainCallsApi.EditContainer(context.Background(), containerId).ContainerEditRequest(containerEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerMainCallsApi.EditContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditContainer`: ContainerResponse
    fmt.Fprintf(os.Stdout, "Response from `ContainerMainCallsApi.EditContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerEditRequest** | [**ContainerEditRequest**](ContainerEditRequest.md) |  | 

### Return type

[**ContainerResponse**](ContainerResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainer

> ContainerResponse GetContainer(ctx, containerId).Execute()

Get container by ID

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
    resp, r, err := apiClient.ContainerMainCallsApi.GetContainer(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerMainCallsApi.GetContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainer`: ContainerResponse
    fmt.Fprintf(os.Stdout, "Response from `ContainerMainCallsApi.GetContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerRequest struct via the builder pattern


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


## GetContainerStatus

> Status GetContainerStatus(ctx, containerId).Execute()

Get container status

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
    resp, r, err := apiClient.ContainerMainCallsApi.GetContainerStatus(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerMainCallsApi.GetContainerStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerStatus`: Status
    fmt.Fprintf(os.Stdout, "Response from `ContainerMainCallsApi.GetContainerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerLinks

> LinkResponseList ListContainerLinks(ctx, containerId).Execute()

List all URLs of the container



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
    resp, r, err := apiClient.ContainerMainCallsApi.ListContainerLinks(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerMainCallsApi.ListContainerLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerLinks`: LinkResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerMainCallsApi.ListContainerLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LinkResponseList**](LinkResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

