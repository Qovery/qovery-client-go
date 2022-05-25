# \ContainerDatabaseApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachDatabaseToContainer**](ContainerDatabaseApi.md#AttachDatabaseToContainer) | **Post** /container/{containerId}/database/{targetDatabaseId} | Link a database to the container
[**AttachLogicalDatabaseToContainer**](ContainerDatabaseApi.md#AttachLogicalDatabaseToContainer) | **Post** /container/{containerId}/logicalDatabase/{targetLogicalDatabaseId} | Link a logical database to the container
[**ListContainerDatabase**](ContainerDatabaseApi.md#ListContainerDatabase) | **Get** /container/{containerId}/database | List linked databases
[**ListContainerLogicalDatabase**](ContainerDatabaseApi.md#ListContainerLogicalDatabase) | **Get** /container/{containerId}/logicalDatabase | List linked logical databases
[**RemoveDatabaseFromContainer**](ContainerDatabaseApi.md#RemoveDatabaseFromContainer) | **Delete** /container/{containerId}/database/{targetDatabaseId} | Remove database link to this container.
[**RemoveLogicalDatabaseFromContainer**](ContainerDatabaseApi.md#RemoveLogicalDatabaseFromContainer) | **Delete** /container/{containerId}/logicalDatabase/{targetLogicalDatabaseId} | Remove logical database link to this container.



## AttachDatabaseToContainer

> Database AttachDatabaseToContainer(ctx, containerId, targetDatabaseId).Execute()

Link a database to the container

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
    targetDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Target database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerDatabaseApi.AttachDatabaseToContainer(context.Background(), containerId, targetDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerDatabaseApi.AttachDatabaseToContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachDatabaseToContainer`: Database
    fmt.Fprintf(os.Stdout, "Response from `ContainerDatabaseApi.AttachDatabaseToContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**targetDatabaseId** | **string** | Target database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachDatabaseToContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Database**](Database.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachLogicalDatabaseToContainer

> LogicalDatabase AttachLogicalDatabaseToContainer(ctx, containerId, targetLogicalDatabaseId).Execute()

Link a logical database to the container

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
    targetLogicalDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Target database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerDatabaseApi.AttachLogicalDatabaseToContainer(context.Background(), containerId, targetLogicalDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerDatabaseApi.AttachLogicalDatabaseToContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachLogicalDatabaseToContainer`: LogicalDatabase
    fmt.Fprintf(os.Stdout, "Response from `ContainerDatabaseApi.AttachLogicalDatabaseToContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**targetLogicalDatabaseId** | **string** | Target database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachLogicalDatabaseToContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LogicalDatabase**](LogicalDatabase.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerDatabase

> DatabaseResponseList ListContainerDatabase(ctx, containerId).Execute()

List linked databases

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
    resp, r, err := apiClient.ContainerDatabaseApi.ListContainerDatabase(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerDatabaseApi.ListContainerDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerDatabase`: DatabaseResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerDatabaseApi.ListContainerDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatabaseResponseList**](DatabaseResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerLogicalDatabase

> LogicalDatabaseResponseList ListContainerLogicalDatabase(ctx, containerId).Execute()

List linked logical databases

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
    resp, r, err := apiClient.ContainerDatabaseApi.ListContainerLogicalDatabase(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerDatabaseApi.ListContainerLogicalDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerLogicalDatabase`: LogicalDatabaseResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerDatabaseApi.ListContainerLogicalDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerLogicalDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogicalDatabaseResponseList**](LogicalDatabaseResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDatabaseFromContainer

> RemoveDatabaseFromContainer(ctx, containerId, targetDatabaseId).Execute()

Remove database link to this container.

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
    targetDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Target database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerDatabaseApi.RemoveDatabaseFromContainer(context.Background(), containerId, targetDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerDatabaseApi.RemoveDatabaseFromContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**targetDatabaseId** | **string** | Target database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDatabaseFromContainerRequest struct via the builder pattern


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


## RemoveLogicalDatabaseFromContainer

> RemoveLogicalDatabaseFromContainer(ctx, containerId, targetLogicalDatabaseId).Execute()

Remove logical database link to this container.

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
    targetLogicalDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Target database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerDatabaseApi.RemoveLogicalDatabaseFromContainer(context.Background(), containerId, targetLogicalDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerDatabaseApi.RemoveLogicalDatabaseFromContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**targetLogicalDatabaseId** | **string** | Target database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLogicalDatabaseFromContainerRequest struct via the builder pattern


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

