# \DatabaseDeploymentHistoryApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDatabaseDeploymentHistory**](DatabaseDeploymentHistoryApi.md#ListDatabaseDeploymentHistory) | **Get** /database/{databaseId}/deploymentHistory | List database deploys



## ListDatabaseDeploymentHistory

> ListDatabaseDeploymentHistory200Response ListDatabaseDeploymentHistory(ctx, databaseId).Execute()

List database deploys



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseDeploymentHistoryApi.ListDatabaseDeploymentHistory(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseDeploymentHistoryApi.ListDatabaseDeploymentHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabaseDeploymentHistory`: ListDatabaseDeploymentHistory200Response
    fmt.Fprintf(os.Stdout, "Response from `DatabaseDeploymentHistoryApi.ListDatabaseDeploymentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseDeploymentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListDatabaseDeploymentHistory200Response**](ListDatabaseDeploymentHistory200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

