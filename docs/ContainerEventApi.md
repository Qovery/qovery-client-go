# \ContainerEventApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListContainerEvent**](ContainerEventApi.md#ListContainerEvent) | **Get** /container/{containerId}/event | NOT YET IMPLEMENTED - List container events



## ListContainerEvent

> EventPaginatedResponseList ListContainerEvent(ctx, containerId).StartId(startId).Execute()

NOT YET IMPLEMENTED - List container events



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    startId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Starting point after which to return results (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerEventApi.ListContainerEvent(context.Background(), containerId).StartId(startId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerEventApi.ListContainerEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerEvent`: EventPaginatedResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerEventApi.ListContainerEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startId** | **string** | Starting point after which to return results | 

### Return type

[**EventPaginatedResponseList**](EventPaginatedResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

