# \DatabaseMetricsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDatabaseCurrentMetric**](DatabaseMetricsApi.md#GetDatabaseCurrentMetric) | **Get** /database/{databaseId}/currentMetric | Get current metric consumption of the database 
[**GetDatabaseMetricCpu**](DatabaseMetricsApi.md#GetDatabaseMetricCpu) | **Get** /database/{databaseId}/metric/cpu | Get CPU consumption metric over time for the database
[**GetDatabaseMetricHealthCheck**](DatabaseMetricsApi.md#GetDatabaseMetricHealthCheck) | **Get** /database/{databaseId}/metric/healthCheck | Get Health Check latency  metric over time for the database
[**GetDatabaseMetricMemory**](DatabaseMetricsApi.md#GetDatabaseMetricMemory) | **Get** /database/{databaseId}/metric/memory | Get Memory consumption metric over time for the database
[**GetDatabaseMetricRestart**](DatabaseMetricsApi.md#GetDatabaseMetricRestart) | **Get** /database/{databaseId}/metric/restart | List database restarts
[**GetDatabaseMetricStorage**](DatabaseMetricsApi.md#GetDatabaseMetricStorage) | **Get** /database/{databaseId}/metric/storage | Get Storage consumption metric over time for the database



## GetDatabaseCurrentMetric

> DatabaseCurrentMetricResponse GetDatabaseCurrentMetric(ctx, databaseId).Execute()

Get current metric consumption of the database 

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
    resp, r, err := apiClient.DatabaseMetricsApi.GetDatabaseCurrentMetric(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMetricsApi.GetDatabaseCurrentMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseCurrentMetric`: DatabaseCurrentMetricResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMetricsApi.GetDatabaseCurrentMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseCurrentMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatabaseCurrentMetricResponse**](DatabaseCurrentMetricResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseMetricCpu

> MetricCPUDatapointResponseList GetDatabaseMetricCpu(ctx, databaseId).LastSeconds(lastSeconds).Execute()

Get CPU consumption metric over time for the database

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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseMetricsApi.GetDatabaseMetricCpu(context.Background(), databaseId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMetricsApi.GetDatabaseMetricCpu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseMetricCpu`: MetricCPUDatapointResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMetricsApi.GetDatabaseMetricCpu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseMetricCpuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeconds** | **float32** | Up to how many seconds in the past to ask analytics results | 

### Return type

[**MetricCPUDatapointResponseList**](MetricCPUDatapointResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseMetricHealthCheck

> MetricGenericResponseList GetDatabaseMetricHealthCheck(ctx, databaseId).LastSeconds(lastSeconds).Execute()

Get Health Check latency  metric over time for the database



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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseMetricsApi.GetDatabaseMetricHealthCheck(context.Background(), databaseId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMetricsApi.GetDatabaseMetricHealthCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseMetricHealthCheck`: MetricGenericResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMetricsApi.GetDatabaseMetricHealthCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseMetricHealthCheckRequest struct via the builder pattern


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


## GetDatabaseMetricMemory

> MetricMemoryDatapointResponseList GetDatabaseMetricMemory(ctx, databaseId).LastSeconds(lastSeconds).Execute()

Get Memory consumption metric over time for the database

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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseMetricsApi.GetDatabaseMetricMemory(context.Background(), databaseId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMetricsApi.GetDatabaseMetricMemory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseMetricMemory`: MetricMemoryDatapointResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMetricsApi.GetDatabaseMetricMemory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseMetricMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeconds** | **float32** | Up to how many seconds in the past to ask analytics results | 

### Return type

[**MetricMemoryDatapointResponseList**](MetricMemoryDatapointResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseMetricRestart

> MetricRestartResponse GetDatabaseMetricRestart(ctx, databaseId).LastSeconds(lastSeconds).Execute()

List database restarts



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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseMetricsApi.GetDatabaseMetricRestart(context.Background(), databaseId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMetricsApi.GetDatabaseMetricRestart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseMetricRestart`: MetricRestartResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMetricsApi.GetDatabaseMetricRestart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseMetricRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeconds** | **float32** | Up to how many seconds in the past to ask analytics results | 

### Return type

[**MetricRestartResponse**](MetricRestartResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseMetricStorage

> MetricStorageDatapointResponseList GetDatabaseMetricStorage(ctx, databaseId).LastSeconds(lastSeconds).Execute()

Get Storage consumption metric over time for the database

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
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseMetricsApi.GetDatabaseMetricStorage(context.Background(), databaseId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMetricsApi.GetDatabaseMetricStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseMetricStorage`: MetricStorageDatapointResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMetricsApi.GetDatabaseMetricStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseMetricStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastSeconds** | **float32** | Up to how many seconds in the past to ask analytics results | 

### Return type

[**MetricStorageDatapointResponseList**](MetricStorageDatapointResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

