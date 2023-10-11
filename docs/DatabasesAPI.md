# \DatabasesAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneDatabase**](DatabasesAPI.md#CloneDatabase) | **Post** /database/{databaseId}/clone | Clone database
[**CreateDatabase**](DatabasesAPI.md#CreateDatabase) | **Post** /environment/{environmentId}/database | Create a database
[**GetEnvironmentDatabaseStatus**](DatabasesAPI.md#GetEnvironmentDatabaseStatus) | **Get** /environment/{environmentId}/database/status | List all environment databases statuses
[**ListDatabase**](DatabasesAPI.md#ListDatabase) | **Get** /environment/{environmentId}/database | List environment databases
[**ListEnvironmentDatabaseConfig**](DatabasesAPI.md#ListEnvironmentDatabaseConfig) | **Get** /environment/{environmentId}/databaseConfiguration | List eligible database types, versions and modes for the environment
[**ListEnvironmentDatabaseCurrentMetric**](DatabasesAPI.md#ListEnvironmentDatabaseCurrentMetric) | **Get** /environment/{environmentId}/database/currentMetric | List current metric consumption for each database



## CloneDatabase

> Database CloneDatabase(ctx, databaseId).CloneDatabaseRequest(cloneDatabaseRequest).Execute()

Clone database



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
    cloneDatabaseRequest := *openapiclient.NewCloneDatabaseRequest("Name_example", "EnvironmentId_example") // CloneDatabaseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.CloneDatabase(context.Background(), databaseId).CloneDatabaseRequest(cloneDatabaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.CloneDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneDatabase`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.CloneDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneDatabaseRequest** | [**CloneDatabaseRequest**](CloneDatabaseRequest.md) |  | 

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


## CreateDatabase

> Database CreateDatabase(ctx, environmentId).DatabaseRequest(databaseRequest).Execute()

Create a database

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
    databaseRequest := *openapiclient.NewDatabaseRequest("Name_example", openapiclient.DatabaseTypeEnum("MONGODB"), "10.1", openapiclient.DatabaseModeEnum("CONTAINER")) // DatabaseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.CreateDatabase(context.Background(), environmentId).DatabaseRequest(databaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.CreateDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabase`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.CreateDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **databaseRequest** | [**DatabaseRequest**](DatabaseRequest.md) |  | 

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


## GetEnvironmentDatabaseStatus

> ReferenceObjectStatusResponseList GetEnvironmentDatabaseStatus(ctx, environmentId).Execute()

List all environment databases statuses



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.GetEnvironmentDatabaseStatus(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetEnvironmentDatabaseStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentDatabaseStatus`: ReferenceObjectStatusResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetEnvironmentDatabaseStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentDatabaseStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReferenceObjectStatusResponseList**](ReferenceObjectStatusResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabase

> DatabaseResponseList ListDatabase(ctx, environmentId).Execute()

List environment databases

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.ListDatabase(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.ListDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabase`: DatabaseResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.ListDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatabaseResponseList**](DatabaseResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironmentDatabaseConfig

> DatabaseConfigurationResponseList ListEnvironmentDatabaseConfig(ctx, environmentId).Execute()

List eligible database types, versions and modes for the environment

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.ListEnvironmentDatabaseConfig(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.ListEnvironmentDatabaseConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnvironmentDatabaseConfig`: DatabaseConfigurationResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.ListEnvironmentDatabaseConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentDatabaseConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatabaseConfigurationResponseList**](DatabaseConfigurationResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironmentDatabaseCurrentMetric

> EnvironmentDatabasesCurrentMetricResponseList ListEnvironmentDatabaseCurrentMetric(ctx, environmentId).Execute()

List current metric consumption for each database

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.ListEnvironmentDatabaseCurrentMetric(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.ListEnvironmentDatabaseCurrentMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnvironmentDatabaseCurrentMetric`: EnvironmentDatabasesCurrentMetricResponseList
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.ListEnvironmentDatabaseCurrentMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentDatabaseCurrentMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentDatabasesCurrentMetricResponseList**](EnvironmentDatabasesCurrentMetricResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

