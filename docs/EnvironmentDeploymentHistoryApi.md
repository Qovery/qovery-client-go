# \EnvironmentDeploymentHistoryApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEnvironmentDeploymentHistory**](EnvironmentDeploymentHistoryApi.md#ListEnvironmentDeploymentHistory) | **Get** /environment/{environmentId}/deploymentHistory | List environment deploys



## ListEnvironmentDeploymentHistory

> DeploymentHistoryEnvironmentPaginatedResponseList ListEnvironmentDeploymentHistory(ctx, environmentId).StartId(startId).Execute()

List environment deploys



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentId := TODO // string | Environment ID
    startId := TODO // string | Starting point after which to return results (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentDeploymentHistoryApi.ListEnvironmentDeploymentHistory(context.Background(), environmentId).StartId(startId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentDeploymentHistoryApi.ListEnvironmentDeploymentHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnvironmentDeploymentHistory`: DeploymentHistoryEnvironmentPaginatedResponseList
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentDeploymentHistoryApi.ListEnvironmentDeploymentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentDeploymentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startId** | [**string**](string.md) | Starting point after which to return results | 

### Return type

[**DeploymentHistoryEnvironmentPaginatedResponseList**](DeploymentHistoryEnvironmentPaginatedResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

