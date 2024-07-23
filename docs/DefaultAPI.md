# \DefaultAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContainerRegistryAssociatedServices**](DefaultAPI.md#GetContainerRegistryAssociatedServices) | **Get** /organization/{organizationId}/containerRegistry/{containerRegistryId}/associatedServices | Get organization container registry associated services



## GetContainerRegistryAssociatedServices

> ContainerRegistryAssociatedServicesResponseList GetContainerRegistryAssociatedServices(ctx, organizationId, containerRegistryId).Execute()

Get organization container registry associated services



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
	containerRegistryId := "containerRegistryId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetContainerRegistryAssociatedServices(context.Background(), organizationId, containerRegistryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetContainerRegistryAssociatedServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerRegistryAssociatedServices`: ContainerRegistryAssociatedServicesResponseList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetContainerRegistryAssociatedServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**containerRegistryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerRegistryAssociatedServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContainerRegistryAssociatedServicesResponseList**](ContainerRegistryAssociatedServicesResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

