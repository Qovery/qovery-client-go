# \JobDeploymentHistoryApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListJobDeploymentHistory**](JobDeploymentHistoryApi.md#ListJobDeploymentHistory) | **Get** /job/{jobId}/deploymentHistory | List job deployments



## ListJobDeploymentHistory

> ListJobDeploymentHistory200Response ListJobDeploymentHistory(ctx).Execute()

List job deployments



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobDeploymentHistoryApi.ListJobDeploymentHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobDeploymentHistoryApi.ListJobDeploymentHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJobDeploymentHistory`: ListJobDeploymentHistory200Response
    fmt.Fprintf(os.Stdout, "Response from `JobDeploymentHistoryApi.ListJobDeploymentHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListJobDeploymentHistoryRequest struct via the builder pattern


### Return type

[**ListJobDeploymentHistory200Response**](ListJobDeploymentHistory200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

