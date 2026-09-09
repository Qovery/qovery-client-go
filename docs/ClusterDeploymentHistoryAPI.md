# \ClusterDeploymentHistoryAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListClusterDeploymentHistoryV2**](ClusterDeploymentHistoryAPI.md#ListClusterDeploymentHistoryV2) | **Get** /organization/{organizationId}/cluster/{clusterId}/deploymentHistoryV2 | List cluster deployments
[**ListClusterDeploymentLogs**](ClusterDeploymentHistoryAPI.md#ListClusterDeploymentLogs) | **Get** /organization/{organizationId}/cluster/{clusterId}/deployment/{deploymentId}/logs | List logs for a specific cluster deployment



## ListClusterDeploymentHistoryV2

> ClusterDeploymentHistoryPaginatedResponseListV2 ListClusterDeploymentHistoryV2(ctx, organizationId, clusterId).PageSize(pageSize).Execute()

List cluster deployments



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	pageSize := float32(8.14) // float32 | The number of deployments to return in the current page. Must be greater than or equal to 1 (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterDeploymentHistoryAPI.ListClusterDeploymentHistoryV2(context.Background(), organizationId, clusterId).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterDeploymentHistoryAPI.ListClusterDeploymentHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterDeploymentHistoryV2`: ClusterDeploymentHistoryPaginatedResponseListV2
	fmt.Fprintf(os.Stdout, "Response from `ClusterDeploymentHistoryAPI.ListClusterDeploymentHistoryV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterDeploymentHistoryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **float32** | The number of deployments to return in the current page. Must be greater than or equal to 1 | [default to 20]

### Return type

[**ClusterDeploymentHistoryPaginatedResponseListV2**](ClusterDeploymentHistoryPaginatedResponseListV2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterDeploymentLogs

> ClusterLogsResponseList ListClusterDeploymentLogs(ctx, organizationId, clusterId, deploymentId).Execute()

List logs for a specific cluster deployment



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterDeploymentHistoryAPI.ListClusterDeploymentLogs(context.Background(), organizationId, clusterId, deploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterDeploymentHistoryAPI.ListClusterDeploymentLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterDeploymentLogs`: ClusterLogsResponseList
	fmt.Fprintf(os.Stdout, "Response from `ClusterDeploymentHistoryAPI.ListClusterDeploymentLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 
**deploymentId** | **string** | Deployment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterDeploymentLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ClusterLogsResponseList**](ClusterLogsResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

