# \JobBuildAnalysisAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJobBuildAnalysis**](JobBuildAnalysisAPI.md#GetJobBuildAnalysis) | **Get** /job/{jobId}/buildAnalysis | Get the build analysis of a job
[**StartJobBuildAnalysis**](JobBuildAnalysisAPI.md#StartJobBuildAnalysis) | **Post** /job/{jobId}/buildAnalysis | Analyse the latest build of a job



## GetJobBuildAnalysis

> BuildAnalysisState GetJobBuildAnalysis(ctx, jobId).Execute()

Get the build analysis of a job



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
	resp, r, err := apiClient.JobBuildAnalysisAPI.GetJobBuildAnalysis(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobBuildAnalysisAPI.GetJobBuildAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobBuildAnalysis`: BuildAnalysisState
	fmt.Fprintf(os.Stdout, "Response from `JobBuildAnalysisAPI.GetJobBuildAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobBuildAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuildAnalysisState**](BuildAnalysisState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartJobBuildAnalysis

> BuildAnalysisResponse StartJobBuildAnalysis(ctx, jobId).BuildAnalysisStartRequest(buildAnalysisStartRequest).Execute()

Analyse the latest build of a job



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
	buildAnalysisStartRequest := *openapiclient.NewBuildAnalysisStartRequest("6e5c1d3a-6b6f-4c35-9a2f-8f1d2c3b4a5e-42") // BuildAnalysisStartRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobBuildAnalysisAPI.StartJobBuildAnalysis(context.Background(), jobId).BuildAnalysisStartRequest(buildAnalysisStartRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobBuildAnalysisAPI.StartJobBuildAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartJobBuildAnalysis`: BuildAnalysisResponse
	fmt.Fprintf(os.Stdout, "Response from `JobBuildAnalysisAPI.StartJobBuildAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartJobBuildAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildAnalysisStartRequest** | [**BuildAnalysisStartRequest**](BuildAnalysisStartRequest.md) |  | 

### Return type

[**BuildAnalysisResponse**](BuildAnalysisResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

