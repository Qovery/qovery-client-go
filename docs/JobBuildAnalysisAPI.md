# \JobBuildAnalysisAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJobBuildAnalysis**](JobBuildAnalysisAPI.md#GetJobBuildAnalysis) | **Get** /job/{jobId}/buildAnalysis | Get the build analysis of a job



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

