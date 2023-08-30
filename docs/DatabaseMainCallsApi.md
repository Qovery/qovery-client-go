# \DatabaseMainCallsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDatabase**](DatabaseMainCallsApi.md#DeleteDatabase) | **Delete** /database/{databaseId} | Delete a database 
[**EditDatabase**](DatabaseMainCallsApi.md#EditDatabase) | **Put** /database/{databaseId} | Edit a database 
[**EditDatabaseCredentials**](DatabaseMainCallsApi.md#EditDatabaseCredentials) | **Put** /database/{databaseId}/masterCredentials | Edit database  master credentials
[**GetDatabase**](DatabaseMainCallsApi.md#GetDatabase) | **Get** /database/{databaseId} | Get database by ID
[**GetDatabaseMasterCredentials**](DatabaseMainCallsApi.md#GetDatabaseMasterCredentials) | **Get** /database/{databaseId}/masterCredentials | Get master credentials of the database
[**GetDatabaseStatus**](DatabaseMainCallsApi.md#GetDatabaseStatus) | **Get** /database/{databaseId}/status | Get database status
[**ListDatabaseVersion**](DatabaseMainCallsApi.md#ListDatabaseVersion) | **Get** /database/{databaseId}/version | List eligible versions for the database



## DeleteDatabase

> DeleteDatabase(ctx, databaseId).Execute()

Delete a database 



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
    resp, r, err := apiClient.DatabaseMainCallsApi.DeleteDatabase(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMainCallsApi.DeleteDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseRequest struct via the builder pattern


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


## EditDatabase

> Database EditDatabase(ctx, databaseId).DatabaseEditRequest(databaseEditRequest).Execute()

Edit a database 



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
    databaseEditRequest := *openapiclient.NewDatabaseEditRequest() // DatabaseEditRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseMainCallsApi.EditDatabase(context.Background(), databaseId).DatabaseEditRequest(databaseEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMainCallsApi.EditDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditDatabase`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMainCallsApi.EditDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **databaseEditRequest** | [**DatabaseEditRequest**](DatabaseEditRequest.md) |  | 

### Return type

[**Database**](Database.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditDatabaseCredentials

> Credentials EditDatabaseCredentials(ctx, databaseId).CredentialsRequest(credentialsRequest).Execute()

Edit database  master credentials

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
    credentialsRequest := *openapiclient.NewCredentialsRequest("Login_example", "Password_example") // CredentialsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabaseMainCallsApi.EditDatabaseCredentials(context.Background(), databaseId).CredentialsRequest(credentialsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMainCallsApi.EditDatabaseCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditDatabaseCredentials`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMainCallsApi.EditDatabaseCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditDatabaseCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialsRequest** | [**CredentialsRequest**](CredentialsRequest.md) |  | 

### Return type

[**Credentials**](Credentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabase

> Database GetDatabase(ctx, databaseId).Execute()

Get database by ID

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
    resp, r, err := apiClient.DatabaseMainCallsApi.GetDatabase(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMainCallsApi.GetDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabase`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMainCallsApi.GetDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Database**](Database.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseMasterCredentials

> Credentials GetDatabaseMasterCredentials(ctx, databaseId).Execute()

Get master credentials of the database

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
    resp, r, err := apiClient.DatabaseMainCallsApi.GetDatabaseMasterCredentials(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMainCallsApi.GetDatabaseMasterCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseMasterCredentials`: Credentials
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMainCallsApi.GetDatabaseMasterCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseMasterCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Credentials**](Credentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseStatus

> Status GetDatabaseStatus(ctx, databaseId).Execute()

Get database status

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
    resp, r, err := apiClient.DatabaseMainCallsApi.GetDatabaseStatus(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMainCallsApi.GetDatabaseStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseStatus`: Status
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMainCallsApi.GetDatabaseStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabaseVersion

> VersionResponseList ListDatabaseVersion(ctx, databaseId).Execute()

List eligible versions for the database

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
    resp, r, err := apiClient.DatabaseMainCallsApi.ListDatabaseVersion(context.Background(), databaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabaseMainCallsApi.ListDatabaseVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabaseVersion`: VersionResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabaseMainCallsApi.ListDatabaseVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VersionResponseList**](VersionResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

