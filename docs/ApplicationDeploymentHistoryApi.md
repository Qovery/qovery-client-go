# \ApplicationDeploymentHistoryApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListApplicationDeploymentHistory**](ApplicationDeploymentHistoryApi.md#ListApplicationDeploymentHistory) | **Get** /application/{applicationId}/deploymentHistory | List application deploys



## ListApplicationDeploymentHistory

> DeploymentHistoryPaginatedResponseList ListApplicationDeploymentHistory(ctx, applicationId).StartId(startId).Execute()

List application deploys



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    startId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Starting point after which to return results (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationDeploymentHistoryApi.ListApplicationDeploymentHistory(context.Background(), applicationId).StartId(startId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDeploymentHistoryApi.ListApplicationDeploymentHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationDeploymentHistory`: DeploymentHistoryPaginatedResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationDeploymentHistoryApi.ListApplicationDeploymentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationDeploymentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startId** | **string** | Starting point after which to return results | 

### Return type

[**DeploymentHistoryPaginatedResponseList**](DeploymentHistoryPaginatedResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

