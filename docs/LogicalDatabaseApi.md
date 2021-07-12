# \LogicalDatabaseApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLogicalDatabase**](LogicalDatabaseApi.md#DeleteLogicalDatabase) | **Delete** /logicalDatabase/{logicalDatabaseId} | Delete a Logical database
[**EditLogicalDatabase**](LogicalDatabaseApi.md#EditLogicalDatabase) | **Put** /logicalDatabase/{logicalDatabaseId} | Edit a logical database
[**EditLogicalDatabaseCredentials**](LogicalDatabaseApi.md#EditLogicalDatabaseCredentials) | **Put** /logicalDatabase/{logicalDatabaseId}/credentials | Edit logical database credentials
[**GetLogicalDatabase**](LogicalDatabaseApi.md#GetLogicalDatabase) | **Get** /logicalDatabase/{logicalDatabaseId} | Get logical database by ID
[**GetLogicalDatabaseCredentials**](LogicalDatabaseApi.md#GetLogicalDatabaseCredentials) | **Get** /logicalDatabase/{logicalDatabaseId}/credentials | Get  credentials of the logical database
[**ListLogicalDatabaseApplication**](LogicalDatabaseApi.md#ListLogicalDatabaseApplication) | **Get** /logicalDatabase/{logicalDatabaseId}/application | List linked applications
[**ListLogicalDatabaseDatabase**](LogicalDatabaseApi.md#ListLogicalDatabaseDatabase) | **Get** /database/{databaseId}/logicalDatabase | List logical databases of a database



## DeleteLogicalDatabase

> DeleteLogicalDatabase(ctx, logicalDatabaseId).Execute()

Delete a Logical database



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
    logicalDatabaseId := TODO // string | Logical Database ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogicalDatabaseApi.DeleteLogicalDatabase(context.Background(), logicalDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogicalDatabaseApi.DeleteLogicalDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalDatabaseId** | [**string**](.md) | Logical Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogicalDatabaseRequest struct via the builder pattern


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


## EditLogicalDatabase

> LogicalDatabaseResponse EditLogicalDatabase(ctx, logicalDatabaseId).LogicalDatabaseRequest(logicalDatabaseRequest).Execute()

Edit a logical database

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
    logicalDatabaseId := TODO // string | Logical Database ID
    logicalDatabaseRequest := *openapiclient.NewLogicalDatabaseRequest("Name_example") // LogicalDatabaseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogicalDatabaseApi.EditLogicalDatabase(context.Background(), logicalDatabaseId).LogicalDatabaseRequest(logicalDatabaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogicalDatabaseApi.EditLogicalDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditLogicalDatabase`: LogicalDatabaseResponse
    fmt.Fprintf(os.Stdout, "Response from `LogicalDatabaseApi.EditLogicalDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalDatabaseId** | [**string**](.md) | Logical Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLogicalDatabaseRequest struct via the builder pattern


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


## EditLogicalDatabaseCredentials

> CredentialsResponse EditLogicalDatabaseCredentials(ctx, logicalDatabaseId).CredentialsRequest(credentialsRequest).Execute()

Edit logical database credentials

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
    logicalDatabaseId := TODO // string | Logical Database ID
    credentialsRequest := *openapiclient.NewCredentialsRequest("Login_example", "Password_example") // CredentialsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogicalDatabaseApi.EditLogicalDatabaseCredentials(context.Background(), logicalDatabaseId).CredentialsRequest(credentialsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogicalDatabaseApi.EditLogicalDatabaseCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditLogicalDatabaseCredentials`: CredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `LogicalDatabaseApi.EditLogicalDatabaseCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalDatabaseId** | [**string**](.md) | Logical Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLogicalDatabaseCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialsRequest** | [**CredentialsRequest**](CredentialsRequest.md) |  | 

### Return type

[**CredentialsResponse**](CredentialsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalDatabase

> LogicalDatabaseResponse GetLogicalDatabase(ctx, logicalDatabaseId).Execute()

Get logical database by ID



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
    logicalDatabaseId := TODO // string | Logical Database ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogicalDatabaseApi.GetLogicalDatabase(context.Background(), logicalDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogicalDatabaseApi.GetLogicalDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogicalDatabase`: LogicalDatabaseResponse
    fmt.Fprintf(os.Stdout, "Response from `LogicalDatabaseApi.GetLogicalDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalDatabaseId** | [**string**](.md) | Logical Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogicalDatabaseResponse**](LogicalDatabaseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalDatabaseCredentials

> CredentialsResponse GetLogicalDatabaseCredentials(ctx, logicalDatabaseId).Execute()

Get  credentials of the logical database

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
    logicalDatabaseId := TODO // string | Logical Database ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogicalDatabaseApi.GetLogicalDatabaseCredentials(context.Background(), logicalDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogicalDatabaseApi.GetLogicalDatabaseCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogicalDatabaseCredentials`: CredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `LogicalDatabaseApi.GetLogicalDatabaseCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalDatabaseId** | [**string**](.md) | Logical Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalDatabaseCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CredentialsResponse**](CredentialsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogicalDatabaseApplication

> ApplicationResponseList ListLogicalDatabaseApplication(ctx, logicalDatabaseId).Execute()

List linked applications

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
    logicalDatabaseId := TODO // string | Logical Database ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LogicalDatabaseApi.ListLogicalDatabaseApplication(context.Background(), logicalDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogicalDatabaseApi.ListLogicalDatabaseApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogicalDatabaseApplication`: ApplicationResponseList
    fmt.Fprintf(os.Stdout, "Response from `LogicalDatabaseApi.ListLogicalDatabaseApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalDatabaseId** | [**string**](.md) | Logical Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogicalDatabaseApplicationRequest struct via the builder pattern


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


## ListLogicalDatabaseDatabase

> LogicalDatabaseResponseList ListLogicalDatabaseDatabase(ctx, databaseId).Execute()

List logical databases of a database



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
    resp, r, err := api_client.LogicalDatabaseApi.ListLogicalDatabaseDatabase(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogicalDatabaseApi.ListLogicalDatabaseDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogicalDatabaseDatabase`: LogicalDatabaseResponseList
    fmt.Fprintf(os.Stdout, "Response from `LogicalDatabaseApi.ListLogicalDatabaseDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | [**string**](.md) | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogicalDatabaseDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogicalDatabaseResponseList**](LogicalDatabaseResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

