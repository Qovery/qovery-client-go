# \ApplicationMetricsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationCurrentInstance**](ApplicationMetricsApi.md#GetApplicationCurrentInstance) | **Get** /application/{applicationId}/instance | List currently running instances of the application with their CPU and RAM metrics
[**GetApplicationCurrentScale**](ApplicationMetricsApi.md#GetApplicationCurrentScale) | **Get** /application/{applicationId}/currentScale | Get current scaling of the application
[**GetApplicationCurrentStorageDisk**](ApplicationMetricsApi.md#GetApplicationCurrentStorageDisk) | **Get** /application/{applicationId}/currentStorage | List current storage disk usage
[**GetApplicationMetricCpu**](ApplicationMetricsApi.md#GetApplicationMetricCpu) | **Get** /application/{applicationId}/metric/cpu | Get CPU consumption metric over time for the application
[**GetApplicationMetricHealthCheck**](ApplicationMetricsApi.md#GetApplicationMetricHealthCheck) | **Get** /application/{applicationId}/metric/healthCheck | Get Health Check latency  metric over time for the application
[**GetApplicationMetricMemory**](ApplicationMetricsApi.md#GetApplicationMetricMemory) | **Get** /application/{applicationId}/metric/memory | Get Memory consumption metric over time for the application
[**GetApplicationMetricStorage**](ApplicationMetricsApi.md#GetApplicationMetricStorage) | **Get** /application/{applicationId}/metric/storage | Get Storage consumption metric over time for the application



## GetApplicationCurrentInstance

> InstanceResponseList GetApplicationCurrentInstance(ctx, applicationId).Execute()

List currently running instances of the application with their CPU and RAM metrics

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationMetricsApi.GetApplicationCurrentInstance(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationMetricsApi.GetApplicationCurrentInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationCurrentInstance`: InstanceResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationMetricsApi.GetApplicationCurrentInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCurrentInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstanceResponseList**](InstanceResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCurrentScale

> ApplicationCurrentScale GetApplicationCurrentScale(ctx, applicationId).Execute()

Get current scaling of the application



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationMetricsApi.GetApplicationCurrentScale(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationMetricsApi.GetApplicationCurrentScale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationCurrentScale`: ApplicationCurrentScale
    fmt.Fprintf(os.Stdout, "Response from `ApplicationMetricsApi.GetApplicationCurrentScale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCurrentScaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationCurrentScale**](ApplicationCurrentScale.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCurrentStorageDisk

> StorageDiskResponseList GetApplicationCurrentStorageDisk(ctx, applicationId).Execute()

List current storage disk usage

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationMetricsApi.GetApplicationCurrentStorageDisk(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationMetricsApi.GetApplicationCurrentStorageDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationCurrentStorageDisk`: StorageDiskResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationMetricsApi.GetApplicationCurrentStorageDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCurrentStorageDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StorageDiskResponseList**](StorageDiskResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationMetricCpu

> MetricCPUResponseList GetApplicationMetricCpu(ctx, applicationId).LastSeconds(lastSeconds).Execute()

Get CPU consumption metric over time for the application

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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationMetricsApi.GetApplicationMetricCpu(context.Background(), applicationId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationMetricsApi.GetApplicationMetricCpu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationMetricCpu`: MetricCPUResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationMetricsApi.GetApplicationMetricCpu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationMetricCpuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeconds** | **float32** | Up to how many seconds in the past to ask analytics results | 

### Return type

[**MetricCPUResponseList**](MetricCPUResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationMetricHealthCheck

> MetricGenericResponseList GetApplicationMetricHealthCheck(ctx, applicationId).LastSeconds(lastSeconds).Execute()

Get Health Check latency  metric over time for the application



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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationMetricsApi.GetApplicationMetricHealthCheck(context.Background(), applicationId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationMetricsApi.GetApplicationMetricHealthCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationMetricHealthCheck`: MetricGenericResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationMetricsApi.GetApplicationMetricHealthCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationMetricHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeconds** | **float32** | Up to how many seconds in the past to ask analytics results | 

### Return type

[**MetricGenericResponseList**](MetricGenericResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationMetricMemory

> MetricMemoryResponseList GetApplicationMetricMemory(ctx, applicationId).LastSeconds(lastSeconds).Execute()

Get Memory consumption metric over time for the application

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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationMetricsApi.GetApplicationMetricMemory(context.Background(), applicationId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationMetricsApi.GetApplicationMetricMemory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationMetricMemory`: MetricMemoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationMetricsApi.GetApplicationMetricMemory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationMetricMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeconds** | **float32** | Up to how many seconds in the past to ask analytics results | 

### Return type

[**MetricMemoryResponseList**](MetricMemoryResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationMetricStorage

> MetricStorageResponseList GetApplicationMetricStorage(ctx, applicationId).LastSeconds(lastSeconds).Execute()

Get Storage consumption metric over time for the application

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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationMetricsApi.GetApplicationMetricStorage(context.Background(), applicationId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationMetricsApi.GetApplicationMetricStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationMetricStorage`: MetricStorageResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationMetricsApi.GetApplicationMetricStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationMetricStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeconds** | **float32** | Up to how many seconds in the past to ask analytics results | 

### Return type

[**MetricStorageResponseList**](MetricStorageResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

