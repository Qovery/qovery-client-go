# \ApplicationDatabaseApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachDatabasetoApplication**](ApplicationDatabaseApi.md#AttachDatabasetoApplication) | **Post** /application/{applicationId}/database/{targetDatabaseId} | Link a database to the application
[**AttachLogicalDatabasetoApplication**](ApplicationDatabaseApi.md#AttachLogicalDatabasetoApplication) | **Post** /application/{applicationId}/logicalDatabase/{targetLogicalDatabaseId} | Link a logical database to the application
[**ListApplicationDatabase**](ApplicationDatabaseApi.md#ListApplicationDatabase) | **Get** /application/{applicationId}/database | List linked databases
[**ListApplicationLogicalDatabase**](ApplicationDatabaseApi.md#ListApplicationLogicalDatabase) | **Get** /application/{applicationId}/logicalDatabase | List linked logical databases
[**RemoveDatabaseFromApplication**](ApplicationDatabaseApi.md#RemoveDatabaseFromApplication) | **Delete** /application/{applicationId}/database/{targetDatabaseId} | Remove database link to this application.
[**RemoveLogicalDatabaseFromApplication**](ApplicationDatabaseApi.md#RemoveLogicalDatabaseFromApplication) | **Delete** /application/{applicationId}/logicalDatabase/{targetLogicalDatabaseId} | Remove logical database link to this application.



## AttachDatabasetoApplication

> DatabaseResponse AttachDatabasetoApplication(ctx, applicationId, targetDatabaseId).Execute()

Link a database to the application

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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    targetDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Target database ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationDatabaseApi.AttachDatabasetoApplication(context.Background(), applicationId, targetDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDatabaseApi.AttachDatabasetoApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachDatabasetoApplication`: DatabaseResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationDatabaseApi.AttachDatabasetoApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**targetDatabaseId** | **string** | Target database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachDatabasetoApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DatabaseResponse**](DatabaseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachLogicalDatabasetoApplication

> LogicalDatabaseResponse AttachLogicalDatabasetoApplication(ctx, applicationId, targetLogicalDatabaseId).Execute()

Link a logical database to the application

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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    targetLogicalDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Target database ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationDatabaseApi.AttachLogicalDatabasetoApplication(context.Background(), applicationId, targetLogicalDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDatabaseApi.AttachLogicalDatabasetoApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachLogicalDatabasetoApplication`: LogicalDatabaseResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationDatabaseApi.AttachLogicalDatabasetoApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**targetLogicalDatabaseId** | **string** | Target database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachLogicalDatabasetoApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LogicalDatabaseResponse**](LogicalDatabaseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationDatabase

> DatabaseResponseList ListApplicationDatabase(ctx, applicationId).Execute()

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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationDatabaseApi.ListApplicationDatabase(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDatabaseApi.ListApplicationDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationDatabase`: DatabaseResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationDatabaseApi.ListApplicationDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationDatabaseRequest struct via the builder pattern


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


## ListApplicationLogicalDatabase

> LogicalDatabaseResponseList ListApplicationLogicalDatabase(ctx, applicationId).Execute()

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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationDatabaseApi.ListApplicationLogicalDatabase(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDatabaseApi.ListApplicationLogicalDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationLogicalDatabase`: LogicalDatabaseResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationDatabaseApi.ListApplicationLogicalDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationLogicalDatabaseRequest struct via the builder pattern


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


## RemoveDatabaseFromApplication

> RemoveDatabaseFromApplication(ctx, applicationId, targetDatabaseId).Execute()

Remove database link to this application.

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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    targetDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Target database ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationDatabaseApi.RemoveDatabaseFromApplication(context.Background(), applicationId, targetDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDatabaseApi.RemoveDatabaseFromApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**targetDatabaseId** | **string** | Target database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDatabaseFromApplicationRequest struct via the builder pattern


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


## RemoveLogicalDatabaseFromApplication

> RemoveLogicalDatabaseFromApplication(ctx, applicationId, targetLogicalDatabaseId).Execute()

Remove logical database link to this application.

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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    targetLogicalDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Target database ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationDatabaseApi.RemoveLogicalDatabaseFromApplication(context.Background(), applicationId, targetLogicalDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDatabaseApi.RemoveLogicalDatabaseFromApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**targetLogicalDatabaseId** | **string** | Target database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLogicalDatabaseFromApplicationRequest struct via the builder pattern


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

