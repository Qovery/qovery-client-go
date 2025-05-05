# \ServiceStatusAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIngressDeploymentStatus**](ServiceStatusAPI.md#GetIngressDeploymentStatus) | **Get** /{serviceType}/{serviceId}/ingressDeploymentStatus | Get Ingress Deployment Status By Service



## GetIngressDeploymentStatus

> IngressDeploymentStatusResponse GetIngressDeploymentStatus(ctx, serviceType, serviceId).Execute()

Get Ingress Deployment Status By Service

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
	serviceType := "serviceType_example" // string | 
	serviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceStatusAPI.GetIngressDeploymentStatus(context.Background(), serviceType, serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceStatusAPI.GetIngressDeploymentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIngressDeploymentStatus`: IngressDeploymentStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceStatusAPI.GetIngressDeploymentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceType** | **string** |  | 
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngressDeploymentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IngressDeploymentStatusResponse**](IngressDeploymentStatusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

