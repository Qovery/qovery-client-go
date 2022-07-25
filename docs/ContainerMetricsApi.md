# \ContainerMetricsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContainerCurrentInstance**](ContainerMetricsApi.md#GetContainerCurrentInstance) | **Get** /container/{containerId}/instance | NOT YET IMPLEMENTED - List currently running instances of the container with their CPU and RAM metrics
[**GetContainerCurrentScale**](ContainerMetricsApi.md#GetContainerCurrentScale) | **Get** /container/{containerId}/currentScale | NOT YET IMPLEMENTED - Get current scaling of the container
[**GetContainerCurrentStorageDisk**](ContainerMetricsApi.md#GetContainerCurrentStorageDisk) | **Get** /container/{containerId}/currentStorage | NOT YET IMPLEMENTED - List current storage disk usage
[**GetContainerMetricCpu**](ContainerMetricsApi.md#GetContainerMetricCpu) | **Get** /container/{containerId}/metric/cpu | NOT YET IMPLEMENTED - Get CPU consumption metric over time for the container
[**GetContainerMetricHealthCheck**](ContainerMetricsApi.md#GetContainerMetricHealthCheck) | **Get** /container/{containerId}/metric/healthCheck | NOT YET IMPLEMENTED - Get Health Check latency  metric over time for the container
[**GetContainerMetricMemory**](ContainerMetricsApi.md#GetContainerMetricMemory) | **Get** /container/{containerId}/metric/memory | NOT YET IMPLEMENTED - Get Memory consumption metric over time for the container
[**GetContainerMetricRestart**](ContainerMetricsApi.md#GetContainerMetricRestart) | **Get** /container/{containerId}/metric/restart | NOT YET IMPLEMENTED - List container restarts
[**GetContainerMetricStorage**](ContainerMetricsApi.md#GetContainerMetricStorage) | **Get** /container/{containerId}/metric/storage | NOT YET IMPLEMENTED - Get Storage consumption metric over time for the container 



## GetContainerCurrentInstance

> InstanceResponseList GetContainerCurrentInstance(ctx, containerId).Execute()

NOT YET IMPLEMENTED - List currently running instances of the container with their CPU and RAM metrics

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
    resp, r, err := apiClient.ContainerMetricsApi.GetContainerCurrentInstance(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerMetricsApi.GetContainerCurrentInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerCurrentInstance`: InstanceResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerMetricsApi.GetContainerCurrentInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerCurrentInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstanceResponseList**](InstanceResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerCurrentScale

> ContainerCurrentScale GetContainerCurrentScale(ctx, containerId).Execute()

NOT YET IMPLEMENTED - Get current scaling of the container



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
    resp, r, err := apiClient.ContainerMetricsApi.GetContainerCurrentScale(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerMetricsApi.GetContainerCurrentScale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerCurrentScale`: ContainerCurrentScale
    fmt.Fprintf(os.Stdout, "Response from `ContainerMetricsApi.GetContainerCurrentScale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerCurrentScaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerCurrentScale**](ContainerCurrentScale.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerCurrentStorageDisk

> StorageDiskResponseList GetContainerCurrentStorageDisk(ctx, containerId).Execute()

NOT YET IMPLEMENTED - List current storage disk usage

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
    resp, r, err := apiClient.ContainerMetricsApi.GetContainerCurrentStorageDisk(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerMetricsApi.GetContainerCurrentStorageDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerCurrentStorageDisk`: StorageDiskResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerMetricsApi.GetContainerCurrentStorageDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerCurrentStorageDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StorageDiskResponseList**](StorageDiskResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerMetricCpu

> MetricCPUResponseList GetContainerMetricCpu(ctx, containerId).LastSeconds(lastSeconds).Execute()

NOT YET IMPLEMENTED - Get CPU consumption metric over time for the container

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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerMetricsApi.GetContainerMetricCpu(context.Background(), containerId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerMetricsApi.GetContainerMetricCpu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerMetricCpu`: MetricCPUResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerMetricsApi.GetContainerMetricCpu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerMetricCpuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeconds** | **float32** | Up to how many seconds in the past to ask analytics results | 

### Return type

[**MetricCPUResponseList**](MetricCPUResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerMetricHealthCheck

> MetricGenericResponseList GetContainerMetricHealthCheck(ctx, containerId).LastSeconds(lastSeconds).Execute()

NOT YET IMPLEMENTED - Get Health Check latency  metric over time for the container



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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerMetricsApi.GetContainerMetricHealthCheck(context.Background(), containerId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerMetricsApi.GetContainerMetricHealthCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerMetricHealthCheck`: MetricGenericResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerMetricsApi.GetContainerMetricHealthCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerMetricHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeconds** | **float32** | Up to how many seconds in the past to ask analytics results | 

### Return type

[**MetricGenericResponseList**](MetricGenericResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerMetricMemory

> MetricMemoryResponseList GetContainerMetricMemory(ctx, containerId).LastSeconds(lastSeconds).Execute()

NOT YET IMPLEMENTED - Get Memory consumption metric over time for the container

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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerMetricsApi.GetContainerMetricMemory(context.Background(), containerId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerMetricsApi.GetContainerMetricMemory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerMetricMemory`: MetricMemoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerMetricsApi.GetContainerMetricMemory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerMetricMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeconds** | **float32** | Up to how many seconds in the past to ask analytics results | 

### Return type

[**MetricMemoryResponseList**](MetricMemoryResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerMetricRestart

> MetricRestart GetContainerMetricRestart(ctx, containerId).LastSeconds(lastSeconds).Execute()

NOT YET IMPLEMENTED - List container restarts



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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerMetricsApi.GetContainerMetricRestart(context.Background(), containerId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerMetricsApi.GetContainerMetricRestart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerMetricRestart`: MetricRestart
    fmt.Fprintf(os.Stdout, "Response from `ContainerMetricsApi.GetContainerMetricRestart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerMetricRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeconds** | **float32** | Up to how many seconds in the past to ask analytics results | 

### Return type

[**MetricRestart**](MetricRestart.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerMetricStorage

> MetricStorageResponseList GetContainerMetricStorage(ctx, containerId).LastSeconds(lastSeconds).Execute()

NOT YET IMPLEMENTED - Get Storage consumption metric over time for the container 

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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerMetricsApi.GetContainerMetricStorage(context.Background(), containerId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerMetricsApi.GetContainerMetricStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerMetricStorage`: MetricStorageResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerMetricsApi.GetContainerMetricStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerMetricStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeconds** | **float32** | Up to how many seconds in the past to ask analytics results | 

### Return type

[**MetricStorageResponseList**](MetricStorageResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

