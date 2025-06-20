# \JobActionsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CleanFailedJob**](JobActionsAPI.md#CleanFailedJob) | **Post** /job/{jobId}/cleanFailedJob | Clean a failed job
[**DeployJob**](JobActionsAPI.md#DeployJob) | **Post** /job/{jobId}/deploy | Deploy job
[**RedeployJob**](JobActionsAPI.md#RedeployJob) | **Post** /job/{jobId}/redeploy | Redeploy job
[**StopJob**](JobActionsAPI.md#StopJob) | **Post** /job/{jobId}/stop | Stop job
[**UninstallJob**](JobActionsAPI.md#UninstallJob) | **Post** /job/{jobId}/uninstall | Uninstall job



## CleanFailedJob

> CleanFailedJob200Response CleanFailedJob(ctx, jobId).Execute()

Clean a failed job

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
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobActionsAPI.CleanFailedJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobActionsAPI.CleanFailedJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CleanFailedJob`: CleanFailedJob200Response
	fmt.Fprintf(os.Stdout, "Response from `JobActionsAPI.CleanFailedJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCleanFailedJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CleanFailedJob200Response**](CleanFailedJob200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
	forceEvent := openapiclient.JobForceEvent("START") // JobForceEvent | When filled, it indicates the target event to be deployed.   If the concerned job hasn't the target event provided, the job won't be deployed.  (optional)
	jobDeployRequest := *openapiclient.NewJobDeployRequest() // JobDeployRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobActionsAPI.DeployJob(context.Background(), jobId).ForceEvent(forceEvent).JobDeployRequest(jobDeployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobActionsAPI.DeployJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployJob`: Status
	fmt.Fprintf(os.Stdout, "Response from `JobActionsAPI.DeployJob`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeployJob

> Status RedeployJob(ctx, jobId).ForceEvent(forceEvent).Execute()

Redeploy job

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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
	forceEvent := openapiclient.JobForceEvent("START") // JobForceEvent | When filled, it indicates the target event to be deployed.   If the concerned job hasn't the target event provided, the job won't be deployed.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobActionsAPI.RedeployJob(context.Background(), jobId).ForceEvent(forceEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobActionsAPI.RedeployJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeployJob`: Status
	fmt.Fprintf(os.Stdout, "Response from `JobActionsAPI.RedeployJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceEvent** | [**JobForceEvent**](JobForceEvent.md) | When filled, it indicates the target event to be deployed.   If the concerned job hasn&#39;t the target event provided, the job won&#39;t be deployed.  | 

### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobActionsAPI.StopJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobActionsAPI.StopJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopJob`: Status
	fmt.Fprintf(os.Stdout, "Response from `JobActionsAPI.StopJob`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallJob

> Status UninstallJob(ctx, jobId).Body(body).Execute()

Uninstall job



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobActionsAPI.UninstallJob(context.Background(), jobId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobActionsAPI.UninstallJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UninstallJob`: Status
	fmt.Fprintf(os.Stdout, "Response from `JobActionsAPI.UninstallJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

