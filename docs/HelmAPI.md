# \HelmAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDefaultHelmAdvancedSettings**](HelmAPI.md#GetDefaultHelmAdvancedSettings) | **Get** /defaultHelmAdvancedSettings | List default helm advanced settings



## GetDefaultHelmAdvancedSettings

> map[string]interface{} GetDefaultHelmAdvancedSettings(ctx).Execute()

List default helm advanced settings

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmAPI.GetDefaultHelmAdvancedSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmAPI.GetDefaultHelmAdvancedSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultHelmAdvancedSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HelmAPI.GetDefaultHelmAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultHelmAdvancedSettingsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

