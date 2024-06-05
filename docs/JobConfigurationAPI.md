# \JobConfigurationAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditJobAdvancedSettings**](JobConfigurationAPI.md#EditJobAdvancedSettings) | **Put** /job/{jobId}/advancedSettings | Edit advanced settings
[**GetJobAdvancedSettings**](JobConfigurationAPI.md#GetJobAdvancedSettings) | **Get** /job/{jobId}/advancedSettings | Get advanced settings



## EditJobAdvancedSettings

> JobAdvancedSettings EditJobAdvancedSettings(ctx, jobId).JobAdvancedSettings(jobAdvancedSettings).Execute()

Edit advanced settings



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
	jobAdvancedSettings := *openapiclient.NewJobAdvancedSettings() // JobAdvancedSettings |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobConfigurationAPI.EditJobAdvancedSettings(context.Background(), jobId).JobAdvancedSettings(jobAdvancedSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobConfigurationAPI.EditJobAdvancedSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditJobAdvancedSettings`: JobAdvancedSettings
	fmt.Fprintf(os.Stdout, "Response from `JobConfigurationAPI.EditJobAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditJobAdvancedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobAdvancedSettings** | [**JobAdvancedSettings**](JobAdvancedSettings.md) |  | 

### Return type

[**JobAdvancedSettings**](JobAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobAdvancedSettings

> JobAdvancedSettings GetJobAdvancedSettings(ctx, jobId).Execute()

Get advanced settings



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
	resp, r, err := apiClient.JobConfigurationAPI.GetJobAdvancedSettings(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobConfigurationAPI.GetJobAdvancedSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobAdvancedSettings`: JobAdvancedSettings
	fmt.Fprintf(os.Stdout, "Response from `JobConfigurationAPI.GetJobAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobAdvancedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobAdvancedSettings**](JobAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

