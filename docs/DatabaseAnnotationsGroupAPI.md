# \DatabaseAnnotationsGroupAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAnnotationsGroupToDatabase**](DatabaseAnnotationsGroupAPI.md#AddAnnotationsGroupToDatabase) | **Post** /database/{databaseId}/annotationsGroup/{annotationsGroupId} | Add annotations group to database
[**DeleteAnnotationsGroupToDatabase**](DatabaseAnnotationsGroupAPI.md#DeleteAnnotationsGroupToDatabase) | **Delete** /database/{databaseId}/annotationsGroup/{annotationsGroupId} | Delete annotations group to database



## AddAnnotationsGroupToDatabase

> AddAnnotationsGroupToDatabase(ctx, databaseId, annotationsGroupId).Execute()

Add annotations group to database



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID
    annotationsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization annotations group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DatabaseAnnotationsGroupAPI.AddAnnotationsGroupToDatabase(context.Background(), databaseId, annotationsGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseAnnotationsGroupAPI.AddAnnotationsGroupToDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 
**annotationsGroupId** | **string** | Organization annotations group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAnnotationsGroupToDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnnotationsGroupToDatabase

> DeleteAnnotationsGroupToDatabase(ctx, databaseId, annotationsGroupId).Execute()

Delete annotations group to database



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID
    annotationsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization annotations group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DatabaseAnnotationsGroupAPI.DeleteAnnotationsGroupToDatabase(context.Background(), databaseId, annotationsGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseAnnotationsGroupAPI.DeleteAnnotationsGroupToDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 
**annotationsGroupId** | **string** | Organization annotations group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnnotationsGroupToDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

