# \JobConfigurationApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditJobAdvancedSettings**](JobConfigurationApi.md#EditJobAdvancedSettings) | **Put** /job/{jobId}/advancedSettings | Edit advanced settings
[**GetJobAdvancedSettings**](JobConfigurationApi.md#GetJobAdvancedSettings) | **Get** /job/{jobId}/advancedSettings | Get advanced settings



## EditJobAdvancedSettings

> JobAdvancedSettings EditJobAdvancedSettings(ctx).JobAdvancedSettings(jobAdvancedSettings).Execute()

Edit advanced settings



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
    jobAdvancedSettings := *openapiclient.NewJobAdvancedSettings() // JobAdvancedSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobConfigurationApi.EditJobAdvancedSettings(context.Background()).JobAdvancedSettings(jobAdvancedSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobConfigurationApi.EditJobAdvancedSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditJobAdvancedSettings`: JobAdvancedSettings
    fmt.Fprintf(os.Stdout, "Response from `JobConfigurationApi.EditJobAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditJobAdvancedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobAdvancedSettings** | [**JobAdvancedSettings**](JobAdvancedSettings.md) |  | 

### Return type

[**JobAdvancedSettings**](JobAdvancedSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobAdvancedSettings

> JobAdvancedSettings GetJobAdvancedSettings(ctx).Execute()

Get advanced settings



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
    resp, r, err := apiClient.JobConfigurationApi.GetJobAdvancedSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobConfigurationApi.GetJobAdvancedSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobAdvancedSettings`: JobAdvancedSettings
    fmt.Fprintf(os.Stdout, "Response from `JobConfigurationApi.GetJobAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobAdvancedSettingsRequest struct via the builder pattern


### Return type

[**JobAdvancedSettings**](JobAdvancedSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

