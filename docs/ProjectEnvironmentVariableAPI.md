# \ProjectEnvironmentVariableAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectEnvironmentVariable**](ProjectEnvironmentVariableAPI.md#CreateProjectEnvironmentVariable) | **Post** /project/{projectId}/environmentVariable | Add an environment variable to the project
[**CreateProjectEnvironmentVariableAlias**](ProjectEnvironmentVariableAPI.md#CreateProjectEnvironmentVariableAlias) | **Post** /project/{projectId}/environmentVariable/{environmentVariableId}/alias | Create an environment variable alias at the project level
[**CreateProjectEnvironmentVariableOverride**](ProjectEnvironmentVariableAPI.md#CreateProjectEnvironmentVariableOverride) | **Post** /project/{projectId}/environmentVariable/{environmentVariableId}/override | Create an environment variable override at the project level
[**DeleteProjectEnvironmentVariable**](ProjectEnvironmentVariableAPI.md#DeleteProjectEnvironmentVariable) | **Delete** /project/{projectId}/environmentVariable/{environmentVariableId} | Delete an environment variable from a project
[**EditProjectEnvironmentVariable**](ProjectEnvironmentVariableAPI.md#EditProjectEnvironmentVariable) | **Put** /project/{projectId}/environmentVariable/{environmentVariableId} | Edit an environment variable belonging to the project
[**ListProjectEnvironmentVariable**](ProjectEnvironmentVariableAPI.md#ListProjectEnvironmentVariable) | **Get** /project/{projectId}/environmentVariable | List project environment variables



## CreateProjectEnvironmentVariable

> EnvironmentVariable CreateProjectEnvironmentVariable(ctx, projectId).EnvironmentVariableRequest(environmentVariableRequest).Execute()

Add an environment variable to the project



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    environmentVariableRequest := *openapiclient.NewEnvironmentVariableRequest("Key_example") // EnvironmentVariableRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectEnvironmentVariableAPI.CreateProjectEnvironmentVariable(context.Background(), projectId).EnvironmentVariableRequest(environmentVariableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEnvironmentVariableAPI.CreateProjectEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectEnvironmentVariable`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `ProjectEnvironmentVariableAPI.CreateProjectEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentVariableRequest** | [**EnvironmentVariableRequest**](EnvironmentVariableRequest.md) |  | 

### Return type

[**EnvironmentVariable**](EnvironmentVariable.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectEnvironmentVariableAlias

> EnvironmentVariable CreateProjectEnvironmentVariableAlias(ctx, projectId, environmentVariableId).Key(key).Execute()

Create an environment variable alias at the project level



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
    key := *openapiclient.NewKey("Key_example") // Key |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectEnvironmentVariableAPI.CreateProjectEnvironmentVariableAlias(context.Background(), projectId, environmentVariableId).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEnvironmentVariableAPI.CreateProjectEnvironmentVariableAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectEnvironmentVariableAlias`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `ProjectEnvironmentVariableAPI.CreateProjectEnvironmentVariableAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectEnvironmentVariableAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | [**Key**](Key.md) |  | 

### Return type

[**EnvironmentVariable**](EnvironmentVariable.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectEnvironmentVariableOverride

> EnvironmentVariable CreateProjectEnvironmentVariableOverride(ctx, projectId, environmentVariableId).Value(value).Execute()

Create an environment variable override at the project level



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
    value := *openapiclient.NewValue() // Value |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectEnvironmentVariableAPI.CreateProjectEnvironmentVariableOverride(context.Background(), projectId, environmentVariableId).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEnvironmentVariableAPI.CreateProjectEnvironmentVariableOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectEnvironmentVariableOverride`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `ProjectEnvironmentVariableAPI.CreateProjectEnvironmentVariableOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectEnvironmentVariableOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | [**Value**](Value.md) |  | 

### Return type

[**EnvironmentVariable**](EnvironmentVariable.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectEnvironmentVariable

> DeleteProjectEnvironmentVariable(ctx, projectId, environmentVariableId).Execute()

Delete an environment variable from a project



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectEnvironmentVariableAPI.DeleteProjectEnvironmentVariable(context.Background(), projectId, environmentVariableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEnvironmentVariableAPI.DeleteProjectEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectEnvironmentVariableRequest struct via the builder pattern


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


## EditProjectEnvironmentVariable

> EnvironmentVariable EditProjectEnvironmentVariable(ctx, projectId, environmentVariableId).EnvironmentVariableEditRequest(environmentVariableEditRequest).Execute()

Edit an environment variable belonging to the project



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
    environmentVariableEditRequest := *openapiclient.NewEnvironmentVariableEditRequest("Key_example") // EnvironmentVariableEditRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectEnvironmentVariableAPI.EditProjectEnvironmentVariable(context.Background(), projectId, environmentVariableId).EnvironmentVariableEditRequest(environmentVariableEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEnvironmentVariableAPI.EditProjectEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProjectEnvironmentVariable`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `ProjectEnvironmentVariableAPI.EditProjectEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditProjectEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environmentVariableEditRequest** | [**EnvironmentVariableEditRequest**](EnvironmentVariableEditRequest.md) |  | 

### Return type

[**EnvironmentVariable**](EnvironmentVariable.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectEnvironmentVariable

> EnvironmentVariableResponseList ListProjectEnvironmentVariable(ctx, projectId).Execute()

List project environment variables

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectEnvironmentVariableAPI.ListProjectEnvironmentVariable(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectEnvironmentVariableAPI.ListProjectEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProjectEnvironmentVariable`: EnvironmentVariableResponseList
    fmt.Fprintf(os.Stdout, "Response from `ProjectEnvironmentVariableAPI.ListProjectEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentVariableResponseList**](EnvironmentVariableResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

