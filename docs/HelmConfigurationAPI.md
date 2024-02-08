# \HelmConfigurationAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditHelmAdvancedSettings**](HelmConfigurationAPI.md#EditHelmAdvancedSettings) | **Put** /helm/{helmId}/advancedSettings | Edit advanced settings
[**GetHelmAdvancedSettings**](HelmConfigurationAPI.md#GetHelmAdvancedSettings) | **Get** /helm/{helmId}/advancedSettings | Get advanced settings



## EditHelmAdvancedSettings

> HelmAdvancedSettings EditHelmAdvancedSettings(ctx, helmId).HelmAdvancedSettings(helmAdvancedSettings).Execute()

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
    helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID
    helmAdvancedSettings := *openapiclient.NewHelmAdvancedSettings() // HelmAdvancedSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmConfigurationAPI.EditHelmAdvancedSettings(context.Background(), helmId).HelmAdvancedSettings(helmAdvancedSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmConfigurationAPI.EditHelmAdvancedSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditHelmAdvancedSettings`: HelmAdvancedSettings
    fmt.Fprintf(os.Stdout, "Response from `HelmConfigurationAPI.EditHelmAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditHelmAdvancedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **helmAdvancedSettings** | [**HelmAdvancedSettings**](HelmAdvancedSettings.md) |  | 

### Return type

[**HelmAdvancedSettings**](HelmAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHelmAdvancedSettings

> HelmAdvancedSettings GetHelmAdvancedSettings(ctx, helmId).Execute()

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
    helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmConfigurationAPI.GetHelmAdvancedSettings(context.Background(), helmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmConfigurationAPI.GetHelmAdvancedSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHelmAdvancedSettings`: HelmAdvancedSettings
    fmt.Fprintf(os.Stdout, "Response from `HelmConfigurationAPI.GetHelmAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHelmAdvancedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HelmAdvancedSettings**](HelmAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

