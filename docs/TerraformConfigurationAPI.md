# \TerraformConfigurationAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditTerraformAdvancedSettings**](TerraformConfigurationAPI.md#EditTerraformAdvancedSettings) | **Put** /terraform/{terraformId}/advancedSettings | Edit Advanced settings
[**GetTerraformAdvancedSettings**](TerraformConfigurationAPI.md#GetTerraformAdvancedSettings) | **Get** /terraform/{terraformId}/advancedSettings | Get Advanced settings



## EditTerraformAdvancedSettings

> TerraformAdvancedSettings EditTerraformAdvancedSettings(ctx, terraformId).TerraformAdvancedSettings(terraformAdvancedSettings).Execute()

Edit Advanced settings

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
	terraformId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Terraform ID
	terraformAdvancedSettings := *openapiclient.NewTerraformAdvancedSettings() // TerraformAdvancedSettings |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerraformConfigurationAPI.EditTerraformAdvancedSettings(context.Background(), terraformId).TerraformAdvancedSettings(terraformAdvancedSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformConfigurationAPI.EditTerraformAdvancedSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditTerraformAdvancedSettings`: TerraformAdvancedSettings
	fmt.Fprintf(os.Stdout, "Response from `TerraformConfigurationAPI.EditTerraformAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTerraformAdvancedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terraformAdvancedSettings** | [**TerraformAdvancedSettings**](TerraformAdvancedSettings.md) |  | 

### Return type

[**TerraformAdvancedSettings**](TerraformAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTerraformAdvancedSettings

> TerraformAdvancedSettings GetTerraformAdvancedSettings(ctx, terraformId).Execute()

Get Advanced settings

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
	terraformId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Terraform ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerraformConfigurationAPI.GetTerraformAdvancedSettings(context.Background(), terraformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformConfigurationAPI.GetTerraformAdvancedSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTerraformAdvancedSettings`: TerraformAdvancedSettings
	fmt.Fprintf(os.Stdout, "Response from `TerraformConfigurationAPI.GetTerraformAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTerraformAdvancedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TerraformAdvancedSettings**](TerraformAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

