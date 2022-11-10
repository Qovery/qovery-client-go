# \JobActionsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployJob**](JobActionsApi.md#DeployJob) | **Post** /job/{jobId}/deploy | Deploy job
[**RestartJob**](JobActionsApi.md#RestartJob) | **Post** /job/{jobId}/restart | Restart job
[**StopJob**](JobActionsApi.md#StopJob) | **Post** /job/{jobId}/stop | Stop job



## DeployJob

> Status DeployJob(ctx).Force(force).JobDeployRequest(jobDeployRequest).Execute()

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
    force := true // bool | Enable or Disable the force trigger of the job (optional) (default to false)
    jobDeployRequest := *openapiclient.NewJobDeployRequest("ImageTag_example") // JobDeployRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobActionsApi.DeployJob(context.Background()).Force(force).JobDeployRequest(jobDeployRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobActionsApi.DeployJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployJob`: Status
    fmt.Fprintf(os.Stdout, "Response from `JobActionsApi.DeployJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeployJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **bool** | Enable or Disable the force trigger of the job | [default to false]
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

> Status RestartJob(ctx).Force(force).Execute()

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
    force := true // bool | Enable or Disable the force trigger of the job (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobActionsApi.RestartJob(context.Background()).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobActionsApi.RestartJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartJob`: Status
    fmt.Fprintf(os.Stdout, "Response from `JobActionsApi.RestartJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestartJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **bool** | Enable or Disable the force trigger of the job | [default to false]

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

> Status StopJob(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobActionsApi.StopJob(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobActionsApi.StopJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopJob`: Status
    fmt.Fprintf(os.Stdout, "Response from `JobActionsApi.StopJob`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStopJobRequest struct via the builder pattern


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

