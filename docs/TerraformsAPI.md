# \TerraformsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneTerraform**](TerraformsAPI.md#CloneTerraform) | **Post** /terraform/{terraformId}/clone | Clone terraform
[**CreateTerraform**](TerraformsAPI.md#CreateTerraform) | **Post** /environment/{environmentId}/terraform | Create a terraform
[**GetDefaultTerraformAdvancedSettings**](TerraformsAPI.md#GetDefaultTerraformAdvancedSettings) | **Get** /defaultTerraformAdvancedSettings | List default terraform advanced settings
[**ListTerraforms**](TerraformsAPI.md#ListTerraforms) | **Get** /environment/{environmentId}/terraform | List terraforms



## CloneTerraform

> TerraformResponse CloneTerraform(ctx, terraformId).CloneServiceRequest(cloneServiceRequest).Execute()

Clone terraform



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
	cloneServiceRequest := *openapiclient.NewCloneServiceRequest("Name_example", "EnvironmentId_example") // CloneServiceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerraformsAPI.CloneTerraform(context.Background(), terraformId).CloneServiceRequest(cloneServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformsAPI.CloneTerraform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneTerraform`: TerraformResponse
	fmt.Fprintf(os.Stdout, "Response from `TerraformsAPI.CloneTerraform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneTerraformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneServiceRequest** | [**CloneServiceRequest**](CloneServiceRequest.md) |  | 

### Return type

[**TerraformResponse**](TerraformResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTerraform

> TerraformResponse CreateTerraform(ctx, environmentId).TerraformRequest(terraformRequest).Execute()

Create a terraform

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
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	terraformRequest := *openapiclient.NewTerraformRequest("Name_example", "Description_example", false, false, openapiclient.TerraformRequest_terraform_files_source{TerraformRequestTerraformFilesSourceOneOf: openapiclient.NewTerraformRequestTerraformFilesSourceOneOf()}, *openapiclient.NewTerraformVariablesSourceRequest([]string{"TfVarFilePaths_example"}, []openapiclient.TerraformVarKeyValue{*openapiclient.NewTerraformVarKeyValue()}), openapiclient.TerraformBackend{TerraformBackendOneOf: openapiclient.NewTerraformBackendOneOf(map[string]interface{}(123))}, "Provider_example", *openapiclient.NewTerraformProviderVersion("ExplicitVersion_example"), *openapiclient.NewTerraformRequestJobResources(int32(123), int32(123), int32(123))) // TerraformRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerraformsAPI.CreateTerraform(context.Background(), environmentId).TerraformRequest(terraformRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformsAPI.CreateTerraform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTerraform`: TerraformResponse
	fmt.Fprintf(os.Stdout, "Response from `TerraformsAPI.CreateTerraform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTerraformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terraformRequest** | [**TerraformRequest**](TerraformRequest.md) |  | 

### Return type

[**TerraformResponse**](TerraformResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultTerraformAdvancedSettings

> TerraformAdvancedSettings GetDefaultTerraformAdvancedSettings(ctx).Execute()

List default terraform advanced settings

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
	resp, r, err := apiClient.TerraformsAPI.GetDefaultTerraformAdvancedSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformsAPI.GetDefaultTerraformAdvancedSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultTerraformAdvancedSettings`: TerraformAdvancedSettings
	fmt.Fprintf(os.Stdout, "Response from `TerraformsAPI.GetDefaultTerraformAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultTerraformAdvancedSettingsRequest struct via the builder pattern


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


## ListTerraforms

> TerraformResponseList ListTerraforms(ctx, environmentId).Execute()

List terraforms

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
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerraformsAPI.ListTerraforms(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformsAPI.ListTerraforms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTerraforms`: TerraformResponseList
	fmt.Fprintf(os.Stdout, "Response from `TerraformsAPI.ListTerraforms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTerraformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TerraformResponseList**](TerraformResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

