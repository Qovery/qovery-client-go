# \TerraformDeploymentHistoryAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListTerraformDeploymentHistoryV2**](TerraformDeploymentHistoryAPI.md#ListTerraformDeploymentHistoryV2) | **Get** /terraform/{terraformId}/deploymentHistoryV2 | List terraform deployments



## ListTerraformDeploymentHistoryV2

> DeploymentHistoryServicePaginatedResponseListV2 ListTerraformDeploymentHistoryV2(ctx, terraformId).Execute()

List terraform deployments



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
	resp, r, err := apiClient.TerraformDeploymentHistoryAPI.ListTerraformDeploymentHistoryV2(context.Background(), terraformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformDeploymentHistoryAPI.ListTerraformDeploymentHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTerraformDeploymentHistoryV2`: DeploymentHistoryServicePaginatedResponseListV2
	fmt.Fprintf(os.Stdout, "Response from `TerraformDeploymentHistoryAPI.ListTerraformDeploymentHistoryV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTerraformDeploymentHistoryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeploymentHistoryServicePaginatedResponseListV2**](DeploymentHistoryServicePaginatedResponseListV2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

