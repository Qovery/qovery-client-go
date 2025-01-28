# \DefaultAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterTokenByClusterId**](DefaultAPI.md#GetClusterTokenByClusterId) | **Get** /cluster/{clusterId}/token | Get cluster token by clusterId
[**GetDeploymentStatusByDeploymentRequestId**](DefaultAPI.md#GetDeploymentStatusByDeploymentRequestId) | **Get** /environment/deploymentStatus | Get Deployment Status By DeploymentRequestId
[**ListServicesByOrganizationId**](DefaultAPI.md#ListServicesByOrganizationId) | **Get** /organization/{organizationId}/services | List Services By OrganizationId



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


## ListServicesByOrganizationId

> ListServicesByOrganizationId200Response ListServicesByOrganizationId(ctx, organizationId).ProjectId(projectId).EnvironmentId(environmentId).Execute()

List Services By OrganizationId

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
	organizationId := "organizationId_example" // string | 
	projectId := "projectId_example" // string |  (optional)
	environmentId := "environmentId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListServicesByOrganizationId(context.Background(), organizationId).ProjectId(projectId).EnvironmentId(environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListServicesByOrganizationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServicesByOrganizationId`: ListServicesByOrganizationId200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListServicesByOrganizationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServicesByOrganizationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **string** |  | 
 **environmentId** | **string** |  | 

### Return type

[**ListServicesByOrganizationId200Response**](ListServicesByOrganizationId200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

