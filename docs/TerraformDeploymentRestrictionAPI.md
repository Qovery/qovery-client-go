# \TerraformDeploymentRestrictionAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTerraformDeploymentRestriction**](TerraformDeploymentRestrictionAPI.md#CreateTerraformDeploymentRestriction) | **Post** /terraform/{terraformId}/deploymentRestriction | Create a terraform deployment restriction
[**GetTerraformDeploymentRestrictions**](TerraformDeploymentRestrictionAPI.md#GetTerraformDeploymentRestrictions) | **Get** /terraform/{terraformId}/deploymentRestriction | Get terraform deployment restrictions



## CreateTerraformDeploymentRestriction

> TerraformDeploymentRestrictionResponse CreateTerraformDeploymentRestriction(ctx, terraformId).TerraformDeploymentRestrictionRequest(terraformDeploymentRestrictionRequest).Execute()

Create a terraform deployment restriction



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
	terraformDeploymentRestrictionRequest := *openapiclient.NewTerraformDeploymentRestrictionRequest(openapiclient.DeploymentRestrictionModeEnum("EXCLUDE"), openapiclient.DeploymentRestrictionTypeEnum("PATH")) // TerraformDeploymentRestrictionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerraformDeploymentRestrictionAPI.CreateTerraformDeploymentRestriction(context.Background(), terraformId).TerraformDeploymentRestrictionRequest(terraformDeploymentRestrictionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformDeploymentRestrictionAPI.CreateTerraformDeploymentRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTerraformDeploymentRestriction`: TerraformDeploymentRestrictionResponse
	fmt.Fprintf(os.Stdout, "Response from `TerraformDeploymentRestrictionAPI.CreateTerraformDeploymentRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTerraformDeploymentRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terraformDeploymentRestrictionRequest** | [**TerraformDeploymentRestrictionRequest**](TerraformDeploymentRestrictionRequest.md) |  | 

### Return type

[**TerraformDeploymentRestrictionResponse**](TerraformDeploymentRestrictionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTerraformDeploymentRestrictions

> TerraformDeploymentRestrictionResponseList GetTerraformDeploymentRestrictions(ctx, terraformId).Execute()

Get terraform deployment restrictions



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
	resp, r, err := apiClient.TerraformDeploymentRestrictionAPI.GetTerraformDeploymentRestrictions(context.Background(), terraformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformDeploymentRestrictionAPI.GetTerraformDeploymentRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTerraformDeploymentRestrictions`: TerraformDeploymentRestrictionResponseList
	fmt.Fprintf(os.Stdout, "Response from `TerraformDeploymentRestrictionAPI.GetTerraformDeploymentRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTerraformDeploymentRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TerraformDeploymentRestrictionResponseList**](TerraformDeploymentRestrictionResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

