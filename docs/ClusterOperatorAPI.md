# \ClusterOperatorAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachClusterOperator**](ClusterOperatorAPI.md#AttachClusterOperator) | **Post** /organization/{organizationId}/cluster/{clusterId}/operator/attach | Attach a cluster to the Qovery Operator execution path
[**GetClusterOperatorBootstrap**](ClusterOperatorAPI.md#GetClusterOperatorBootstrap) | **Get** /organization/{organizationId}/cluster/{clusterId}/operator/bootstrap | Get the Qovery Operator bootstrap
[**GetClusterOperatorStatus**](ClusterOperatorAPI.md#GetClusterOperatorStatus) | **Get** /organization/{organizationId}/cluster/{clusterId}/operator/status | Get the Qovery Operator status for a cluster
[**ListClusterOperatorFleet**](ClusterOperatorAPI.md#ListClusterOperatorFleet) | **Get** /admin/operator/clusters | List the Qovery Operator fleet



## AttachClusterOperator

> AttachClusterOperator(ctx, organizationId, clusterId).Execute()

Attach a cluster to the Qovery Operator execution path



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClusterOperatorAPI.AttachClusterOperator(context.Background(), organizationId, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterOperatorAPI.AttachClusterOperator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachClusterOperatorRequest struct via the builder pattern


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


## GetClusterOperatorBootstrap

> ClusterOperatorBootstrapResponse GetClusterOperatorBootstrap(ctx, organizationId, clusterId).Execute()

Get the Qovery Operator bootstrap



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterOperatorAPI.GetClusterOperatorBootstrap(context.Background(), organizationId, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterOperatorAPI.GetClusterOperatorBootstrap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterOperatorBootstrap`: ClusterOperatorBootstrapResponse
	fmt.Fprintf(os.Stdout, "Response from `ClusterOperatorAPI.GetClusterOperatorBootstrap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterOperatorBootstrapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterOperatorBootstrapResponse**](ClusterOperatorBootstrapResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterOperatorStatus

> ClusterOperatorStatusResponse GetClusterOperatorStatus(ctx, organizationId, clusterId).Execute()

Get the Qovery Operator status for a cluster



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterOperatorAPI.GetClusterOperatorStatus(context.Background(), organizationId, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterOperatorAPI.GetClusterOperatorStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterOperatorStatus`: ClusterOperatorStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ClusterOperatorAPI.GetClusterOperatorStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterOperatorStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterOperatorStatusResponse**](ClusterOperatorStatusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterOperatorFleet

> ClusterOperatorFleetInventoryResponseList ListClusterOperatorFleet(ctx).Execute()

List the Qovery Operator fleet



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
	resp, r, err := apiClient.ClusterOperatorAPI.ListClusterOperatorFleet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterOperatorAPI.ListClusterOperatorFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterOperatorFleet`: ClusterOperatorFleetInventoryResponseList
	fmt.Fprintf(os.Stdout, "Response from `ClusterOperatorAPI.ListClusterOperatorFleet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterOperatorFleetRequest struct via the builder pattern


### Return type

[**ClusterOperatorFleetInventoryResponseList**](ClusterOperatorFleetInventoryResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

