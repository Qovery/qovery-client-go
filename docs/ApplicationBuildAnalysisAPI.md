# \ApplicationBuildAnalysisAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationBuildAnalysis**](ApplicationBuildAnalysisAPI.md#GetApplicationBuildAnalysis) | **Get** /application/{applicationId}/buildAnalysis | Get the build analysis of an application



## GetApplicationBuildAnalysis

> BuildAnalysisState GetApplicationBuildAnalysis(ctx, applicationId).Execute()

Get the build analysis of an application



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
	applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAnalysisAPI.GetApplicationBuildAnalysis(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAnalysisAPI.GetApplicationBuildAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationBuildAnalysis`: BuildAnalysisState
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAnalysisAPI.GetApplicationBuildAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationBuildAnalysisRequest struct via the builder pattern


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

