# \DatabaseApplicationAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDatabaseApplication**](DatabaseApplicationAPI.md#ListDatabaseApplication) | **Get** /database/{databaseId}/application | List applications using the database
[**RemoveApplicationFromDatabase**](DatabaseApplicationAPI.md#RemoveApplicationFromDatabase) | **Delete** /database/{databaseId}/application/{targetApplicationId} | Remove an application from this database 



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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabaseApplicationAPI.ListDatabaseApplication(context.Background(), databaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApplicationAPI.ListDatabaseApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatabaseApplication`: ApplicationResponseList
	fmt.Fprintf(os.Stdout, "Response from `DatabaseApplicationAPI.ListDatabaseApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationResponseList**](ApplicationResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID
	targetApplicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Target application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DatabaseApplicationAPI.RemoveApplicationFromDatabase(context.Background(), databaseId, targetApplicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseApplicationAPI.RemoveApplicationFromDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 
**targetApplicationId** | **string** | Target application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveApplicationFromDatabaseRequest struct via the builder pattern


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

