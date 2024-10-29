# \EnvironmentDeploymentHistoryAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEnvironmentDeploymentHistory**](EnvironmentDeploymentHistoryAPI.md#ListEnvironmentDeploymentHistory) | **Get** /environment/{environmentId}/deploymentHistory | List environment deployments
[**ListEnvironmentDeploymentHistoryV2**](EnvironmentDeploymentHistoryAPI.md#ListEnvironmentDeploymentHistoryV2) | **Get** /environment/{environmentId}/deploymentHistoryV2 | List environment deployments



## ListEnvironmentDeploymentHistory

> DeploymentHistoryEnvironmentPaginatedResponseList ListEnvironmentDeploymentHistory(ctx, environmentId).StartId(startId).Execute()

List environment deployments



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
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
	startId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Starting point after which to return results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentDeploymentHistoryAPI.ListEnvironmentDeploymentHistory(context.Background(), environmentId).StartId(startId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentDeploymentHistoryAPI.ListEnvironmentDeploymentHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnvironmentDeploymentHistory`: DeploymentHistoryEnvironmentPaginatedResponseList
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentDeploymentHistoryAPI.ListEnvironmentDeploymentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentDeploymentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startId** | **string** | Starting point after which to return results | 

### Return type

[**DeploymentHistoryEnvironmentPaginatedResponseList**](DeploymentHistoryEnvironmentPaginatedResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironmentDeploymentHistoryV2

> DeploymentHistoryEnvironmentPaginatedResponseListV2 ListEnvironmentDeploymentHistoryV2(ctx, environmentId).StartId(startId).Execute()

List environment deployments



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
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
	startId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Starting point after which to return results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentDeploymentHistoryAPI.ListEnvironmentDeploymentHistoryV2(context.Background(), environmentId).StartId(startId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentDeploymentHistoryAPI.ListEnvironmentDeploymentHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnvironmentDeploymentHistoryV2`: DeploymentHistoryEnvironmentPaginatedResponseListV2
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentDeploymentHistoryAPI.ListEnvironmentDeploymentHistoryV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentDeploymentHistoryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startId** | **string** | Starting point after which to return results | 

### Return type

[**DeploymentHistoryEnvironmentPaginatedResponseListV2**](DeploymentHistoryEnvironmentPaginatedResponseListV2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

