# \DefaultAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterLogs**](DefaultAPI.md#GetClusterLogs) | **Get** /cluster/{clusterId}/logs | Fetch cluster logs
[**GetClusterTokenByClusterId**](DefaultAPI.md#GetClusterTokenByClusterId) | **Get** /cluster/{clusterId}/token | Get cluster token by clusterId
[**GetDeploymentStatusByDeploymentRequestId**](DefaultAPI.md#GetDeploymentStatusByDeploymentRequestId) | **Get** /environment/deploymentStatus | Get Deployment Status By DeploymentRequestId



## GetClusterLogs

> ClusterLogsResponse GetClusterLogs(ctx, clusterId).Endpoint(endpoint).Query(query).Start(start).End(end).Limit(limit).Since(since).Step(step).Interval(interval).Direction(direction).Execute()

Fetch cluster logs



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
	clusterId := "clusterId_example" // string | 
	endpoint := "endpoint_example" // string | 
	query := "query_example" // string | 
	start := "start_example" // string |  (optional)
	end := "end_example" // string |  (optional)
	limit := "limit_example" // string |  (optional)
	since := "since_example" // string |  (optional)
	step := "step_example" // string |  (optional)
	interval := "interval_example" // string |  (optional)
	direction := "direction_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetClusterLogs(context.Background(), clusterId).Endpoint(endpoint).Query(query).Start(start).End(end).Limit(limit).Since(since).Step(step).Interval(interval).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetClusterLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterLogs`: ClusterLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetClusterLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpoint** | **string** |  | 
 **query** | **string** |  | 
 **start** | **string** |  | 
 **end** | **string** |  | 
 **limit** | **string** |  | 
 **since** | **string** |  | 
 **step** | **string** |  | 
 **interval** | **string** |  | 
 **direction** | **string** |  | 

### Return type

[**ClusterLogsResponse**](ClusterLogsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterTokenByClusterId

> GetClusterTokenByClusterId200Response GetClusterTokenByClusterId(ctx, clusterId).Execute()

Get cluster token by clusterId

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
	clusterId := "clusterId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetClusterTokenByClusterId(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetClusterTokenByClusterId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterTokenByClusterId`: GetClusterTokenByClusterId200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetClusterTokenByClusterId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterTokenByClusterIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetClusterTokenByClusterId200Response**](GetClusterTokenByClusterId200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentStatusByDeploymentRequestId

> EnvDeploymentStatus GetDeploymentStatusByDeploymentRequestId(ctx).DeploymentRequestId(deploymentRequestId).Execute()

Get Deployment Status By DeploymentRequestId

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
	deploymentRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetDeploymentStatusByDeploymentRequestId(context.Background()).DeploymentRequestId(deploymentRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetDeploymentStatusByDeploymentRequestId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentStatusByDeploymentRequestId`: EnvDeploymentStatus
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetDeploymentStatusByDeploymentRequestId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentStatusByDeploymentRequestIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentRequestId** | **string** |  | 

### Return type

[**EnvDeploymentStatus**](EnvDeploymentStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

