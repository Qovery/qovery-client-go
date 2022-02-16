# \DatabaseEventApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDatabaseEvent**](DatabaseEventApi.md#ListDatabaseEvent) | **Get** /database/{databaseId}/event | List database  events



## ListDatabaseEvent

> EventPaginatedResponseList ListDatabaseEvent(ctx, databaseId).StartId(startId).Execute()

List database  events



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
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID
    startId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Starting point after which to return results (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseEventApi.ListDatabaseEvent(context.Background(), databaseId).StartId(startId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseEventApi.ListDatabaseEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabaseEvent`: EventPaginatedResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseEventApi.ListDatabaseEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseEventRequest struct via the builder pattern


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

