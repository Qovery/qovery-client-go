# \DatabaseApplicationApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDatabaseApplication**](DatabaseApplicationApi.md#ListDatabaseApplication) | **Get** /database/{databaseId}/application | List applications using the database
[**RemoveApplicationFromDatabase**](DatabaseApplicationApi.md#RemoveApplicationFromDatabase) | **Delete** /database/{databaseId}/application/{targetApplicationId} | Remove an application from this database 



## ListDatabaseApplication

> ApplicationResponseList ListDatabaseApplication(ctx, databaseId).Execute()

List applications using the database

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
    databaseId := TODO // string | Database ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseApplicationApi.ListDatabaseApplication(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApplicationApi.ListDatabaseApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabaseApplication`: ApplicationResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseApplicationApi.ListDatabaseApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | [**string**](.md) | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationResponseList**](ApplicationResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveApplicationFromDatabase

> RemoveApplicationFromDatabase(ctx, databaseId, targetApplicationId).Execute()

Remove an application from this database 

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
    databaseId := TODO // string | Database ID
    targetApplicationId := TODO // string | Target application ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DatabaseApplicationApi.RemoveApplicationFromDatabase(context.Background(), databaseId, targetApplicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApplicationApi.RemoveApplicationFromDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | [**string**](.md) | Database ID | 
**targetApplicationId** | [**string**](.md) | Target application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveApplicationFromDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

