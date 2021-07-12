# \ApplicationEventApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListApplicationEvent**](ApplicationEventApi.md#ListApplicationEvent) | **Get** /application/{applicationId}/event | List application events



## ListApplicationEvent

> EventPaginatedResponseList ListApplicationEvent(ctx, applicationId).StartId(startId).Execute()

List application events



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
    applicationId := TODO // string | Application ID
    startId := TODO // string | Starting point after which to return results (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationEventApi.ListApplicationEvent(context.Background(), applicationId).StartId(startId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationEventApi.ListApplicationEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationEvent`: EventPaginatedResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationEventApi.ListApplicationEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | [**string**](.md) | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startId** | [**string**](string.md) | Starting point after which to return results | 

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

