# \JobDeploymentRestrictionAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJobDeploymentRestriction**](JobDeploymentRestrictionAPI.md#CreateJobDeploymentRestriction) | **Post** /job/{jobId}/deploymentRestriction | Create a job deployment restriction
[**DeleteJobDeploymentRestriction**](JobDeploymentRestrictionAPI.md#DeleteJobDeploymentRestriction) | **Delete** /job/{jobId}/deploymentRestriction/{deploymentRestrictionId} | Delete a job deployment restriction
[**EditJobDeploymentRestriction**](JobDeploymentRestrictionAPI.md#EditJobDeploymentRestriction) | **Put** /job/{jobId}/deploymentRestriction/{deploymentRestrictionId} | Edit a job deployment restriction
[**GetJobDeploymentRestrictions**](JobDeploymentRestrictionAPI.md#GetJobDeploymentRestrictions) | **Get** /job/{jobId}/deploymentRestriction | Get job deployment restrictions



## CreateJobDeploymentRestriction

> JobDeploymentRestrictionResponse CreateJobDeploymentRestriction(ctx, jobId).JobDeploymentRestrictionRequest(jobDeploymentRestrictionRequest).Execute()

Create a job deployment restriction



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
	jobDeploymentRestrictionRequest := *openapiclient.NewJobDeploymentRestrictionRequest(openapiclient.DeploymentRestrictionModeEnum("EXCLUDE"), openapiclient.DeploymentRestrictionTypeEnum("PATH"), "job1/src/") // JobDeploymentRestrictionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobDeploymentRestrictionAPI.CreateJobDeploymentRestriction(context.Background(), jobId).JobDeploymentRestrictionRequest(jobDeploymentRestrictionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobDeploymentRestrictionAPI.CreateJobDeploymentRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJobDeploymentRestriction`: JobDeploymentRestrictionResponse
	fmt.Fprintf(os.Stdout, "Response from `JobDeploymentRestrictionAPI.CreateJobDeploymentRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobDeploymentRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobDeploymentRestrictionRequest** | [**JobDeploymentRestrictionRequest**](JobDeploymentRestrictionRequest.md) |  | 

### Return type

[**JobDeploymentRestrictionResponse**](JobDeploymentRestrictionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJobDeploymentRestriction

> DeleteJobDeploymentRestriction(ctx, jobId, deploymentRestrictionId).Execute()

Delete a job deployment restriction



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
	deploymentRestrictionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Restriction ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobDeploymentRestrictionAPI.DeleteJobDeploymentRestriction(context.Background(), jobId, deploymentRestrictionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobDeploymentRestrictionAPI.DeleteJobDeploymentRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 
**deploymentRestrictionId** | **string** | Deployment Restriction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobDeploymentRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditJobDeploymentRestriction

> JobDeploymentRestrictionResponse EditJobDeploymentRestriction(ctx, jobId, deploymentRestrictionId).JobDeploymentRestrictionRequest(jobDeploymentRestrictionRequest).Execute()

Edit a job deployment restriction



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
	deploymentRestrictionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Restriction ID
	jobDeploymentRestrictionRequest := *openapiclient.NewJobDeploymentRestrictionRequest(openapiclient.DeploymentRestrictionModeEnum("EXCLUDE"), openapiclient.DeploymentRestrictionTypeEnum("PATH"), "job1/src/") // JobDeploymentRestrictionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobDeploymentRestrictionAPI.EditJobDeploymentRestriction(context.Background(), jobId, deploymentRestrictionId).JobDeploymentRestrictionRequest(jobDeploymentRestrictionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobDeploymentRestrictionAPI.EditJobDeploymentRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditJobDeploymentRestriction`: JobDeploymentRestrictionResponse
	fmt.Fprintf(os.Stdout, "Response from `JobDeploymentRestrictionAPI.EditJobDeploymentRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 
**deploymentRestrictionId** | **string** | Deployment Restriction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditJobDeploymentRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jobDeploymentRestrictionRequest** | [**JobDeploymentRestrictionRequest**](JobDeploymentRestrictionRequest.md) |  | 

### Return type

[**JobDeploymentRestrictionResponse**](JobDeploymentRestrictionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobDeploymentRestrictions

> JobDeploymentRestrictionResponseList GetJobDeploymentRestrictions(ctx, jobId).Execute()

Get job deployment restrictions



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
	resp, r, err := apiClient.JobDeploymentRestrictionAPI.GetJobDeploymentRestrictions(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobDeploymentRestrictionAPI.GetJobDeploymentRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobDeploymentRestrictions`: JobDeploymentRestrictionResponseList
	fmt.Fprintf(os.Stdout, "Response from `JobDeploymentRestrictionAPI.GetJobDeploymentRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobDeploymentRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobDeploymentRestrictionResponseList**](JobDeploymentRestrictionResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

