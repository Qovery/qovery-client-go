# \LogicalDatabaseApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogicalDatabaseOnDatabase**](LogicalDatabaseApi.md#CreateLogicalDatabaseOnDatabase) | **Post** /database/{databaseId}/logicalDatabase | Create a logical database on the database
[**DeleteLogicalDatabase**](LogicalDatabaseApi.md#DeleteLogicalDatabase) | **Delete** /logicalDatabase/{logicalDatabaseId} | Delete a Logical database
[**EditLogicalDatabase**](LogicalDatabaseApi.md#EditLogicalDatabase) | **Put** /logicalDatabase/{logicalDatabaseId} | Edit a logical database
[**EditLogicalDatabaseCredentials**](LogicalDatabaseApi.md#EditLogicalDatabaseCredentials) | **Put** /logicalDatabase/{logicalDatabaseId}/credentials | Edit logical database credentials
[**GetLogicalDatabase**](LogicalDatabaseApi.md#GetLogicalDatabase) | **Get** /logicalDatabase/{logicalDatabaseId} | Get logical database by ID
[**GetLogicalDatabaseCredentials**](LogicalDatabaseApi.md#GetLogicalDatabaseCredentials) | **Get** /logicalDatabase/{logicalDatabaseId}/credentials | Get  credentials of the logical database
[**ListLogicalDatabaseApplication**](LogicalDatabaseApi.md#ListLogicalDatabaseApplication) | **Get** /logicalDatabase/{logicalDatabaseId}/application | List linked applications
[**ListLogicalDatabaseContainer**](LogicalDatabaseApi.md#ListLogicalDatabaseContainer) | **Get** /logicalDatabase/{logicalDatabaseId}/container | List linked containers
[**ListLogicalDatabaseDatabase**](LogicalDatabaseApi.md#ListLogicalDatabaseDatabase) | **Get** /database/{databaseId}/logicalDatabase | List logical databases of a database



## CreateLogicalDatabaseOnDatabase

> LogicalDatabase CreateLogicalDatabaseOnDatabase(ctx, databaseId).LogicalDatabaseRequest(logicalDatabaseRequest).Execute()

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
    resp, r, err := apiClient.LogicalDatabaseApi.CreateLogicalDatabaseOnDatabase(context.Background(), databaseId).LogicalDatabaseRequest(logicalDatabaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogicalDatabaseApi.CreateLogicalDatabaseOnDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogicalDatabaseOnDatabase`: LogicalDatabase
    fmt.Fprintf(os.Stdout, "Response from `LogicalDatabaseApi.CreateLogicalDatabaseOnDatabase`: %v\n", resp)
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

[**LogicalDatabase**](LogicalDatabase.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    logicalDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Logical Database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogicalDatabaseApi.DeleteLogicalDatabase(context.Background(), logicalDatabaseId).Execute()
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
**logicalDatabaseId** | **string** | Logical Database ID | 

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

> LogicalDatabase EditLogicalDatabase(ctx, logicalDatabaseId).LogicalDatabaseRequest(logicalDatabaseRequest).Execute()

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
    logicalDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Logical Database ID
    logicalDatabaseRequest := *openapiclient.NewLogicalDatabaseRequest("Name_example") // LogicalDatabaseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogicalDatabaseApi.EditLogicalDatabase(context.Background(), logicalDatabaseId).LogicalDatabaseRequest(logicalDatabaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogicalDatabaseApi.EditLogicalDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditLogicalDatabase`: LogicalDatabase
    fmt.Fprintf(os.Stdout, "Response from `LogicalDatabaseApi.EditLogicalDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalDatabaseId** | **string** | Logical Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLogicalDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logicalDatabaseRequest** | [**LogicalDatabaseRequest**](LogicalDatabaseRequest.md) |  | 

### Return type

[**LogicalDatabase**](LogicalDatabase.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditLogicalDatabaseCredentials

> Credentials EditLogicalDatabaseCredentials(ctx, logicalDatabaseId).CredentialsRequest(credentialsRequest).Execute()

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
    logicalDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Logical Database ID
    credentialsRequest := *openapiclient.NewCredentialsRequest("Login_example", "Password_example") // CredentialsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogicalDatabaseApi.EditLogicalDatabaseCredentials(context.Background(), logicalDatabaseId).CredentialsRequest(credentialsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogicalDatabaseApi.EditLogicalDatabaseCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditLogicalDatabaseCredentials`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `LogicalDatabaseApi.EditLogicalDatabaseCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalDatabaseId** | **string** | Logical Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLogicalDatabaseCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialsRequest** | [**CredentialsRequest**](CredentialsRequest.md) |  | 

### Return type

[**Credentials**](Credentials.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalDatabase

> LogicalDatabase GetLogicalDatabase(ctx, logicalDatabaseId).Execute()

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
    logicalDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Logical Database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogicalDatabaseApi.GetLogicalDatabase(context.Background(), logicalDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogicalDatabaseApi.GetLogicalDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogicalDatabase`: LogicalDatabase
    fmt.Fprintf(os.Stdout, "Response from `LogicalDatabaseApi.GetLogicalDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalDatabaseId** | **string** | Logical Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogicalDatabase**](LogicalDatabase.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogicalDatabaseCredentials

> Credentials GetLogicalDatabaseCredentials(ctx, logicalDatabaseId).Execute()

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
    logicalDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Logical Database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogicalDatabaseApi.GetLogicalDatabaseCredentials(context.Background(), logicalDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogicalDatabaseApi.GetLogicalDatabaseCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogicalDatabaseCredentials`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `LogicalDatabaseApi.GetLogicalDatabaseCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalDatabaseId** | **string** | Logical Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalDatabaseCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Credentials**](Credentials.md)

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
    logicalDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Logical Database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogicalDatabaseApi.ListLogicalDatabaseApplication(context.Background(), logicalDatabaseId).Execute()
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
**logicalDatabaseId** | **string** | Logical Database ID | 

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


## ListLogicalDatabaseContainer

> ContainerResponseList ListLogicalDatabaseContainer(ctx, logicalDatabaseId).Execute()

List linked containers

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
    logicalDatabaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Logical Database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogicalDatabaseApi.ListLogicalDatabaseContainer(context.Background(), logicalDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogicalDatabaseApi.ListLogicalDatabaseContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogicalDatabaseContainer`: ContainerResponseList
    fmt.Fprintf(os.Stdout, "Response from `LogicalDatabaseApi.ListLogicalDatabaseContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logicalDatabaseId** | **string** | Logical Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogicalDatabaseContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerResponseList**](ContainerResponseList.md)

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
    databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogicalDatabaseApi.ListLogicalDatabaseDatabase(context.Background(), databaseId).Execute()
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
**databaseId** | **string** | Database ID | 

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

