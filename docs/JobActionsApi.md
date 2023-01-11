# \JobActionsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployJob**](JobActionsApi.md#DeployJob) | **Post** /job/{jobId}/deploy | Deploy job
[**RestartJob**](JobActionsApi.md#RestartJob) | **Post** /job/{jobId}/restart | Restart job
[**StopJob**](JobActionsApi.md#StopJob) | **Post** /job/{jobId}/stop | Stop job



## DeployJob

> Status DeployJob(ctx, jobId).ForceEvent(forceEvent).JobDeployRequest(jobDeployRequest).Execute()

Deploy job



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
    forceEvent := openapiclient.JobForceEvent("START") // JobForceEvent | When filled, it indicates the target event to be deployed.   If the concerned job hasn't the target event provided, the job won't be deployed.  (optional)
    jobDeployRequest := *openapiclient.NewJobDeployRequest() // JobDeployRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobActionsApi.DeployJob(context.Background(), jobId).ForceEvent(forceEvent).JobDeployRequest(jobDeployRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobActionsApi.DeployJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployJob`: Status
    fmt.Fprintf(os.Stdout, "Response from `JobActionsApi.DeployJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceEvent** | [**JobForceEvent**](JobForceEvent.md) | When filled, it indicates the target event to be deployed.   If the concerned job hasn&#39;t the target event provided, the job won&#39;t be deployed.  | 
 **jobDeployRequest** | [**JobDeployRequest**](JobDeployRequest.md) |  | 

### Return type

[**Status**](Status.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartJob

> Status RestartJob(ctx, jobId).ForceEvent(forceEvent).Execute()

Restart job

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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
    forceEvent := openapiclient.JobForceEvent("START") // JobForceEvent | When filled, it indicates the target event to be deployed.   If the concerned job hasn't the target event provided, the job won't be deployed.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobActionsApi.RestartJob(context.Background(), jobId).ForceEvent(forceEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobActionsApi.RestartJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartJob`: Status
    fmt.Fprintf(os.Stdout, "Response from `JobActionsApi.RestartJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceEvent** | [**JobForceEvent**](JobForceEvent.md) | When filled, it indicates the target event to be deployed.   If the concerned job hasn&#39;t the target event provided, the job won&#39;t be deployed.  | 

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


## StopJob

> Status StopJob(ctx, jobId).Execute()

Stop job

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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobActionsApi.StopJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobActionsApi.StopJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopJob`: Status
    fmt.Fprintf(os.Stdout, "Response from `JobActionsApi.StopJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopJobRequest struct via the builder pattern


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
