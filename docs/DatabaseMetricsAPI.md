# \DatabaseMetricsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDatabaseCurrentMetric**](DatabaseMetricsAPI.md#GetDatabaseCurrentMetric) | **Get** /database/{databaseId}/currentMetric | Get current metric consumption of the database 
[**GetDatabaseMetricCpu**](DatabaseMetricsAPI.md#GetDatabaseMetricCpu) | **Get** /database/{databaseId}/metric/cpu | Get CPU consumption metric over time for the database
[**GetDatabaseMetricHealthCheck**](DatabaseMetricsAPI.md#GetDatabaseMetricHealthCheck) | **Get** /database/{databaseId}/metric/healthCheck | Get Health Check latency  metric over time for the database
[**GetDatabaseMetricMemory**](DatabaseMetricsAPI.md#GetDatabaseMetricMemory) | **Get** /database/{databaseId}/metric/memory | Get Memory consumption metric over time for the database
[**GetDatabaseMetricStorage**](DatabaseMetricsAPI.md#GetDatabaseMetricStorage) | **Get** /database/{databaseId}/metric/storage | Get Storage consumption metric over time for the database



## GetDatabaseCurrentMetric

> DatabaseCurrentMetric GetDatabaseCurrentMetric(ctx, databaseId).Execute()

Get current metric consumption of the database 

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
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseMetricsAPI.GetDatabaseCurrentMetric(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMetricsAPI.GetDatabaseCurrentMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseCurrentMetric`: DatabaseCurrentMetric
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMetricsAPI.GetDatabaseCurrentMetric`: %v\n", resp)
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

[**DatabaseCurrentMetric**](DatabaseCurrentMetric.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

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
    openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseMetricsAPI.GetDatabaseMetricCpu(context.Background(), databaseId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMetricsAPI.GetDatabaseMetricCpu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseMetricCpu`: MetricCPUDatapointResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMetricsAPI.GetDatabaseMetricCpu`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

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
    openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseMetricsAPI.GetDatabaseMetricHealthCheck(context.Background(), databaseId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMetricsAPI.GetDatabaseMetricHealthCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseMetricHealthCheck`: MetricGenericResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMetricsAPI.GetDatabaseMetricHealthCheck`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

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
    openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseMetricsAPI.GetDatabaseMetricMemory(context.Background(), databaseId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMetricsAPI.GetDatabaseMetricMemory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseMetricMemory`: MetricMemoryDatapointResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMetricsAPI.GetDatabaseMetricMemory`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

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
    openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID
    lastSeconds := float32(8.14) // float32 | Up to how many seconds in the past to ask analytics results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseMetricsAPI.GetDatabaseMetricStorage(context.Background(), databaseId).LastSeconds(lastSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMetricsAPI.GetDatabaseMetricStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseMetricStorage`: MetricStorageDatapointResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMetricsAPI.GetDatabaseMetricStorage`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

