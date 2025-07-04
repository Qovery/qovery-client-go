# \ContainerDeploymentHistoryAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListContainerDeploymentHistory**](ContainerDeploymentHistoryAPI.md#ListContainerDeploymentHistory) | **Get** /container/{containerId}/deploymentHistory | List container deployments
[**ListContainerDeploymentHistoryV2**](ContainerDeploymentHistoryAPI.md#ListContainerDeploymentHistoryV2) | **Get** /container/{containerId}/deploymentHistoryV2 | List container deployments



## ListContainerDeploymentHistory

> ListContainerDeploymentHistory200Response ListContainerDeploymentHistory(ctx, containerId).Execute()

List container deployments



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
	containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerDeploymentHistoryAPI.ListContainerDeploymentHistory(context.Background(), containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerDeploymentHistoryAPI.ListContainerDeploymentHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContainerDeploymentHistory`: ListContainerDeploymentHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `ContainerDeploymentHistoryAPI.ListContainerDeploymentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerDeploymentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListContainerDeploymentHistory200Response**](ListContainerDeploymentHistory200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerDeploymentHistoryV2

> DeploymentHistoryServicePaginatedResponseListV2 ListContainerDeploymentHistoryV2(ctx, containerId).PageSize(pageSize).Execute()

List container deployments



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
	containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
	pageSize := float32(8.14) // float32 | The number of deployments to return in the current page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerDeploymentHistoryAPI.ListContainerDeploymentHistoryV2(context.Background(), containerId).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerDeploymentHistoryAPI.ListContainerDeploymentHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContainerDeploymentHistoryV2`: DeploymentHistoryServicePaginatedResponseListV2
	fmt.Fprintf(os.Stdout, "Response from `ContainerDeploymentHistoryAPI.ListContainerDeploymentHistoryV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerDeploymentHistoryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **float32** | The number of deployments to return in the current page | [default to 20]

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

