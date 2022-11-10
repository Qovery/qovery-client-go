# \JobMetricsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJobCurrentInstance**](JobMetricsApi.md#GetJobCurrentInstance) | **Get** /job/{jobId}/instance | List currently running instances of the job with their CPU and RAM metrics



## GetJobCurrentInstance

> InstanceResponseList GetJobCurrentInstance(ctx).Execute()

List currently running instances of the job with their CPU and RAM metrics

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
    resp, r, err := apiClient.JobMetricsApi.GetJobCurrentInstance(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobMetricsApi.GetJobCurrentInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobCurrentInstance`: InstanceResponseList
    fmt.Fprintf(os.Stdout, "Response from `JobMetricsApi.GetJobCurrentInstance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobCurrentInstanceRequest struct via the builder pattern


### Return type

[**InstanceResponseList**](InstanceResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

