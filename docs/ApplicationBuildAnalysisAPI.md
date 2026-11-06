# \ApplicationBuildAnalysisAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationBuildAnalysis**](ApplicationBuildAnalysisAPI.md#GetApplicationBuildAnalysis) | **Get** /application/{applicationId}/buildAnalysis | Get the build analysis of an application
[**StartApplicationBuildAnalysis**](ApplicationBuildAnalysisAPI.md#StartApplicationBuildAnalysis) | **Post** /application/{applicationId}/buildAnalysis | Analyse the latest build of an application



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


## StartApplicationBuildAnalysis

> BuildAnalysisResponse StartApplicationBuildAnalysis(ctx, applicationId).BuildAnalysisStartRequest(buildAnalysisStartRequest).Execute()

Analyse the latest build of an application



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
	buildAnalysisStartRequest := *openapiclient.NewBuildAnalysisStartRequest("6e5c1d3a-6b6f-4c35-9a2f-8f1d2c3b4a5e-42") // BuildAnalysisStartRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationBuildAnalysisAPI.StartApplicationBuildAnalysis(context.Background(), applicationId).BuildAnalysisStartRequest(buildAnalysisStartRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationBuildAnalysisAPI.StartApplicationBuildAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartApplicationBuildAnalysis`: BuildAnalysisResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationBuildAnalysisAPI.StartApplicationBuildAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartApplicationBuildAnalysisRequest struct via the builder pattern


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

