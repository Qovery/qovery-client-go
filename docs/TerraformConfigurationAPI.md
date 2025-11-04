# \TerraformConfigurationAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTerraformVariable**](TerraformConfigurationAPI.md#DeleteTerraformVariable) | **Delete** /terraform/{terraformId}/variables/{key} | Delete a terraform variable
[**EditTerraformAdvancedSettings**](TerraformConfigurationAPI.md#EditTerraformAdvancedSettings) | **Put** /terraform/{terraformId}/advancedSettings | Edit Advanced settings
[**GetTerraformAdvancedSettings**](TerraformConfigurationAPI.md#GetTerraformAdvancedSettings) | **Get** /terraform/{terraformId}/advancedSettings | Get Advanced settings
[**GetTerraformVariables**](TerraformConfigurationAPI.md#GetTerraformVariables) | **Get** /terraform/{terraformId}/variables | Get terraform variables
[**UpdateTerraformVariable**](TerraformConfigurationAPI.md#UpdateTerraformVariable) | **Post** /terraform/{terraformId}/variables | Create or update a terraform variable



## DeleteTerraformVariable

> DeleteTerraformVariable(ctx, terraformId, key).Execute()

Delete a terraform variable

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
	key := "key_example" // string | Variable key to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TerraformConfigurationAPI.DeleteTerraformVariable(context.Background(), terraformId, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformConfigurationAPI.DeleteTerraformVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 
**key** | **string** | Variable key to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTerraformVariableRequest struct via the builder pattern


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
	terraformId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
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
**terraformId** | **string** |  | 

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
	terraformId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
**terraformId** | **string** |  | 

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


## GetTerraformVariables

> TerraformVariablesResponse GetTerraformVariables(ctx, terraformId).Execute()

Get terraform variables

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
	resp, r, err := apiClient.TerraformConfigurationAPI.GetTerraformVariables(context.Background(), terraformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformConfigurationAPI.GetTerraformVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTerraformVariables`: TerraformVariablesResponse
	fmt.Fprintf(os.Stdout, "Response from `TerraformConfigurationAPI.GetTerraformVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTerraformVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TerraformVariablesResponse**](TerraformVariablesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTerraformVariable

> UpdateTerraformVariable(ctx, terraformId).TerraformVarKeyValue(terraformVarKeyValue).Execute()

Create or update a terraform variable

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
	terraformVarKeyValue := *openapiclient.NewTerraformVarKeyValue() // TerraformVarKeyValue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TerraformConfigurationAPI.UpdateTerraformVariable(context.Background(), terraformId).TerraformVarKeyValue(terraformVarKeyValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformConfigurationAPI.UpdateTerraformVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTerraformVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terraformVarKeyValue** | [**TerraformVarKeyValue**](TerraformVarKeyValue.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

