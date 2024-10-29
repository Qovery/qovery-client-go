# \HelmDeploymentHistoryAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListHelmDeploymentHistory**](HelmDeploymentHistoryAPI.md#ListHelmDeploymentHistory) | **Get** /helm/{helmId}/deploymentHistory | List helm deployments
[**ListHelmDeploymentHistoryV2**](HelmDeploymentHistoryAPI.md#ListHelmDeploymentHistoryV2) | **Get** /helm/{helmId}/deploymentHistoryV2 | List helm deployments



## ListHelmDeploymentHistory

> ListHelmDeploymentHistory200Response ListHelmDeploymentHistory(ctx, helmId).Execute()

List helm deployments



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
	helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmDeploymentHistoryAPI.ListHelmDeploymentHistory(context.Background(), helmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmDeploymentHistoryAPI.ListHelmDeploymentHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHelmDeploymentHistory`: ListHelmDeploymentHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `HelmDeploymentHistoryAPI.ListHelmDeploymentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHelmDeploymentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListHelmDeploymentHistory200Response**](ListHelmDeploymentHistory200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHelmDeploymentHistoryV2

> DeploymentHistoryServicePaginatedResponseListV2 ListHelmDeploymentHistoryV2(ctx, helmId).Execute()

List helm deployments



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
	helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmDeploymentHistoryAPI.ListHelmDeploymentHistoryV2(context.Background(), helmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmDeploymentHistoryAPI.ListHelmDeploymentHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHelmDeploymentHistoryV2`: DeploymentHistoryServicePaginatedResponseListV2
	fmt.Fprintf(os.Stdout, "Response from `HelmDeploymentHistoryAPI.ListHelmDeploymentHistoryV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHelmDeploymentHistoryV2Request struct via the builder pattern


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

