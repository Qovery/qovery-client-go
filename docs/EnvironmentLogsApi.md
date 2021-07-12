# \EnvironmentLogsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEnvironmentLog**](EnvironmentLogsApi.md#ListEnvironmentLog) | **Get** /environment/{environmentId}/log | List environment deployment logs



## ListEnvironmentLog

> EnvironmentLogResponseList ListEnvironmentLog(ctx, environmentId).Execute()

List environment deployment logs



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentLogsApi.ListEnvironmentLog(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentLogsApi.ListEnvironmentLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnvironmentLog`: EnvironmentLogResponseList
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentLogsApi.ListEnvironmentLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentLogResponseList**](EnvironmentLogResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

