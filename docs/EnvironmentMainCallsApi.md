# \EnvironmentMainCallsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEnvironment**](EnvironmentMainCallsApi.md#DeleteEnvironment) | **Delete** /environment/{environmentId} | Delete an environment
[**EditEnvironment**](EnvironmentMainCallsApi.md#EditEnvironment) | **Put** /environment/{environmentId} | Edit an environment
[**GetEnvironment**](EnvironmentMainCallsApi.md#GetEnvironment) | **Get** /environment/{environmentId} | Get environment by ID
[**GetEnvironmentStatus**](EnvironmentMainCallsApi.md#GetEnvironmentStatus) | **Get** /environment/{environmentId}/status | Get environment status
[**GetEnvironmentStatuses**](EnvironmentMainCallsApi.md#GetEnvironmentStatuses) | **Get** /environment/{environmentId}/statuses | Get environment statuses with services status
[**GetEnvironmentStatusesWithStages**](EnvironmentMainCallsApi.md#GetEnvironmentStatusesWithStages) | **Get** /environment/{environmentId}/statusesWithStages | Get environment statuses with stages



## DeleteEnvironment

> DeleteEnvironment(ctx, environmentId).Execute()

Delete an environment



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentMainCallsApi.DeleteEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentMainCallsApi.DeleteEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentRequest struct via the builder pattern


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


## EditEnvironment

> Environment EditEnvironment(ctx, environmentId).EnvironmentEditRequest(environmentEditRequest).Execute()

Edit an environment



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
    environmentEditRequest := *openapiclient.NewEnvironmentEditRequest() // EnvironmentEditRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentMainCallsApi.EditEnvironment(context.Background(), environmentId).EnvironmentEditRequest(environmentEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentMainCallsApi.EditEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentMainCallsApi.EditEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentEditRequest** | [**EnvironmentEditRequest**](EnvironmentEditRequest.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironment

> Environment GetEnvironment(ctx, environmentId).Execute()

Get environment by ID

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentMainCallsApi.GetEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentMainCallsApi.GetEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentMainCallsApi.GetEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Environment**](Environment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentStatus

> EnvironmentStatus GetEnvironmentStatus(ctx, environmentId).Execute()

Get environment status

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentMainCallsApi.GetEnvironmentStatus(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentMainCallsApi.GetEnvironmentStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentStatus`: EnvironmentStatus
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentMainCallsApi.GetEnvironmentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentStatus**](EnvironmentStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentStatuses

> EnvironmentStatuses GetEnvironmentStatuses(ctx, environmentId).Execute()

Get environment statuses with services status

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentMainCallsApi.GetEnvironmentStatuses(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentMainCallsApi.GetEnvironmentStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentStatuses`: EnvironmentStatuses
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentMainCallsApi.GetEnvironmentStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentStatuses**](EnvironmentStatuses.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentStatusesWithStages

> GetEnvironmentStatusesWithStages200Response GetEnvironmentStatusesWithStages(ctx, environmentId).Execute()

Get environment statuses with stages

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentMainCallsApi.GetEnvironmentStatusesWithStages(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentMainCallsApi.GetEnvironmentStatusesWithStages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentStatusesWithStages`: GetEnvironmentStatusesWithStages200Response
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentMainCallsApi.GetEnvironmentStatusesWithStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentStatusesWithStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEnvironmentStatusesWithStages200Response**](GetEnvironmentStatusesWithStages200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

