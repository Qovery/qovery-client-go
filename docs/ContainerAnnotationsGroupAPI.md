# \ContainerAnnotationsGroupAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAnnotationsGroupToContainer**](ContainerAnnotationsGroupAPI.md#AddAnnotationsGroupToContainer) | **Post** /container/{containerId}/annotationsGroup/{annotationsGroupId} | Add annotations group to container
[**DeleteAnnotationsGroupToContainer**](ContainerAnnotationsGroupAPI.md#DeleteAnnotationsGroupToContainer) | **Delete** /container/{containerId}/annotationsGroup/{annotationsGroupId} | Delete annotations group to container
[**ListContainerAnnotationsGroup**](ContainerAnnotationsGroupAPI.md#ListContainerAnnotationsGroup) | **Get** /container/{containerId}/annotationsGroup | List container annotations group



## AddAnnotationsGroupToContainer

> AddAnnotationsGroupToContainer(ctx, containerId, annotationsGroupId).Execute()

Add annotations group to container



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    annotationsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization annotations group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContainerAnnotationsGroupAPI.AddAnnotationsGroupToContainer(context.Background(), containerId, annotationsGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAnnotationsGroupAPI.AddAnnotationsGroupToContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**annotationsGroupId** | **string** | Organization annotations group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAnnotationsGroupToContainerRequest struct via the builder pattern


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


## DeleteAnnotationsGroupToContainer

> DeleteAnnotationsGroupToContainer(ctx, containerId, annotationsGroupId).Execute()

Delete annotations group to container



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    annotationsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization annotations group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContainerAnnotationsGroupAPI.DeleteAnnotationsGroupToContainer(context.Background(), containerId, annotationsGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAnnotationsGroupAPI.DeleteAnnotationsGroupToContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**annotationsGroupId** | **string** | Organization annotations group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnnotationsGroupToContainerRequest struct via the builder pattern


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


## ListContainerAnnotationsGroup

> []OrganizationAnnotationsGroupResponse ListContainerAnnotationsGroup(ctx, containerId).Execute()

List container annotations group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerAnnotationsGroupAPI.ListContainerAnnotationsGroup(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerAnnotationsGroupAPI.ListContainerAnnotationsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerAnnotationsGroup`: []OrganizationAnnotationsGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ContainerAnnotationsGroupAPI.ListContainerAnnotationsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerAnnotationsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OrganizationAnnotationsGroupResponse**](OrganizationAnnotationsGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

