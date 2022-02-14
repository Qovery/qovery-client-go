# \DatabaseApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogicalDatabaseOnDatabase**](DatabaseApi.md#CreateLogicalDatabaseOnDatabase) | **Post** /database/{databaseId}/logicalDatabase | Create a logical database on the database



## CreateLogicalDatabaseOnDatabase

> LogicalDatabaseResponse CreateLogicalDatabaseOnDatabase(ctx, databaseId).LogicalDatabaseRequest(logicalDatabaseRequest).Execute()

Create a logical database on the database



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
    logicalDatabaseRequest := *openapiclient.NewLogicalDatabaseRequest("Name_example") // LogicalDatabaseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseApi.CreateLogicalDatabaseOnDatabase(context.Background(), databaseId).LogicalDatabaseRequest(logicalDatabaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApi.CreateLogicalDatabaseOnDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogicalDatabaseOnDatabase`: LogicalDatabaseResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApi.CreateLogicalDatabaseOnDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalDatabaseOnDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logicalDatabaseRequest** | [**LogicalDatabaseRequest**](LogicalDatabaseRequest.md) |  | 

### Return type

[**LogicalDatabaseResponse**](LogicalDatabaseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

