# \TerraformActionsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployTerraform**](TerraformActionsAPI.md#DeployTerraform) | **Post** /terraform/{terraformId}/deploy | Deploy terraform
[**RedeployTerraform**](TerraformActionsAPI.md#RedeployTerraform) | **Post** /terraform/{terraformId}/redeploy | Redeploy terraform
[**UninstallTerraform**](TerraformActionsAPI.md#UninstallTerraform) | **Post** /terraform/{terraformId}/uninstall | Uninstall terraform



## DeployTerraform

> Status DeployTerraform(ctx, terraformId).TerraformDeployRequest(terraformDeployRequest).Execute()

Deploy terraform



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
	terraformDeployRequest := *openapiclient.NewTerraformDeployRequest() // TerraformDeployRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerraformActionsAPI.DeployTerraform(context.Background(), terraformId).TerraformDeployRequest(terraformDeployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformActionsAPI.DeployTerraform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployTerraform`: Status
	fmt.Fprintf(os.Stdout, "Response from `TerraformActionsAPI.DeployTerraform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployTerraformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terraformDeployRequest** | [**TerraformDeployRequest**](TerraformDeployRequest.md) |  | 

### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeployTerraform

> Status RedeployTerraform(ctx, terraformId).Execute()

Redeploy terraform

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
	terraformId := "terraformId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerraformActionsAPI.RedeployTerraform(context.Background(), terraformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformActionsAPI.RedeployTerraform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeployTerraform`: Status
	fmt.Fprintf(os.Stdout, "Response from `TerraformActionsAPI.RedeployTerraform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployTerraformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallTerraform

> Status UninstallTerraform(ctx, terraformId).ForceTerraformAction(forceTerraformAction).Body(body).Execute()

Uninstall terraform



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
	forceTerraformAction := openapiclient.DeleteTerraformAction("SKIP_DESTROY") // DeleteTerraformAction | Force a specific action to be executed by Terraform during uninstall. (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerraformActionsAPI.UninstallTerraform(context.Background(), terraformId).ForceTerraformAction(forceTerraformAction).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformActionsAPI.UninstallTerraform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UninstallTerraform`: Status
	fmt.Fprintf(os.Stdout, "Response from `TerraformActionsAPI.UninstallTerraform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallTerraformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceTerraformAction** | [**DeleteTerraformAction**](DeleteTerraformAction.md) | Force a specific action to be executed by Terraform during uninstall. | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

