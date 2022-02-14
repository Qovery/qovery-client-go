# \DatabaseActionsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployDatabase**](DatabaseActionsApi.md#DeployDatabase) | **Post** /database/{databaseId}/deploy | Deploy database 
[**RestartDatabase**](DatabaseActionsApi.md#RestartDatabase) | **Post** /database/{databaseId}/restart | Retart database
[**StopDatabase**](DatabaseActionsApi.md#StopDatabase) | **Post** /database/{databaseId}/stop | Stop database



## DeployDatabase

> Status DeployDatabase(ctx, databaseId).Execute()

Deploy database 

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
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseActionsApi.DeployDatabase(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseActionsApi.DeployDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployDatabase`: Status
    fmt.Fprintf(os.Stdout, "Response from `DatabaseActionsApi.DeployDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployDatabaseRequest struct via the builder pattern


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


## RestartDatabase

> Status RestartDatabase(ctx, databaseId).Execute()

Retart database

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
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseActionsApi.RestartDatabase(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseActionsApi.RestartDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartDatabase`: Status
    fmt.Fprintf(os.Stdout, "Response from `DatabaseActionsApi.RestartDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartDatabaseRequest struct via the builder pattern


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


## StopDatabase

> Status StopDatabase(ctx, databaseId).Execute()

Stop database

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
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseActionsApi.StopDatabase(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseActionsApi.StopDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopDatabase`: Status
    fmt.Fprintf(os.Stdout, "Response from `DatabaseActionsApi.StopDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopDatabaseRequest struct via the builder pattern


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

