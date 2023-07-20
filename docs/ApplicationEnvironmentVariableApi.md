# \ApplicationEnvironmentVariableApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationEnvironmentVariable**](ApplicationEnvironmentVariableApi.md#CreateApplicationEnvironmentVariable) | **Post** /application/{applicationId}/environmentVariable | Add an environment variable to the application
[**CreateApplicationEnvironmentVariableAlias**](ApplicationEnvironmentVariableApi.md#CreateApplicationEnvironmentVariableAlias) | **Post** /application/{applicationId}/environmentVariable/{environmentVariableId}/alias | Create an environment variable alias at the application level
[**CreateApplicationEnvironmentVariableOverride**](ApplicationEnvironmentVariableApi.md#CreateApplicationEnvironmentVariableOverride) | **Post** /application/{applicationId}/environmentVariable/{environmentVariableId}/override | Create an environment variable override at the application level
[**DeleteApplicationEnvironmentVariable**](ApplicationEnvironmentVariableApi.md#DeleteApplicationEnvironmentVariable) | **Delete** /application/{applicationId}/environmentVariable/{environmentVariableId} | Delete an environment variable from an application
[**EditApplicationEnvironmentVariable**](ApplicationEnvironmentVariableApi.md#EditApplicationEnvironmentVariable) | **Put** /application/{applicationId}/environmentVariable/{environmentVariableId} | Edit an environment variable belonging to the application
[**ImportEnvironmentVariable**](ApplicationEnvironmentVariableApi.md#ImportEnvironmentVariable) | **Post** /application/{applicationId}/environmentVariable/import | Import variables
[**ListApplicationEnvironmentVariable**](ApplicationEnvironmentVariableApi.md#ListApplicationEnvironmentVariable) | **Get** /application/{applicationId}/environmentVariable | List environment variables



## CreateApplicationEnvironmentVariable

> EnvironmentVariable CreateApplicationEnvironmentVariable(ctx, applicationId).EnvironmentVariableRequest(environmentVariableRequest).Execute()

Add an environment variable to the application



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    environmentVariableRequest := *openapiclient.NewEnvironmentVariableRequest("Key_example") // EnvironmentVariableRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationEnvironmentVariableApi.CreateApplicationEnvironmentVariable(context.Background(), applicationId).EnvironmentVariableRequest(environmentVariableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationEnvironmentVariableApi.CreateApplicationEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationEnvironmentVariable`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `ApplicationEnvironmentVariableApi.CreateApplicationEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentVariableRequest** | [**EnvironmentVariableRequest**](EnvironmentVariableRequest.md) |  | 

### Return type

[**EnvironmentVariable**](EnvironmentVariable.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationEnvironmentVariableAlias

> EnvironmentVariable CreateApplicationEnvironmentVariableAlias(ctx, applicationId, environmentVariableId).Key(key).Execute()

Create an environment variable alias at the application level



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
    key := *openapiclient.NewKey("Key_example") // Key |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationEnvironmentVariableApi.CreateApplicationEnvironmentVariableAlias(context.Background(), applicationId, environmentVariableId).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationEnvironmentVariableApi.CreateApplicationEnvironmentVariableAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationEnvironmentVariableAlias`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `ApplicationEnvironmentVariableApi.CreateApplicationEnvironmentVariableAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationEnvironmentVariableAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | [**Key**](Key.md) |  | 

### Return type

[**EnvironmentVariable**](EnvironmentVariable.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationEnvironmentVariableOverride

> EnvironmentVariable CreateApplicationEnvironmentVariableOverride(ctx, applicationId, environmentVariableId).Value(value).Execute()

Create an environment variable override at the application level



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
    value := *openapiclient.NewValue() // Value |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationEnvironmentVariableApi.CreateApplicationEnvironmentVariableOverride(context.Background(), applicationId, environmentVariableId).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationEnvironmentVariableApi.CreateApplicationEnvironmentVariableOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationEnvironmentVariableOverride`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `ApplicationEnvironmentVariableApi.CreateApplicationEnvironmentVariableOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationEnvironmentVariableOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | [**Value**](Value.md) |  | 

### Return type

[**EnvironmentVariable**](EnvironmentVariable.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationEnvironmentVariable

> DeleteApplicationEnvironmentVariable(ctx, applicationId, environmentVariableId).Execute()

Delete an environment variable from an application



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationEnvironmentVariableApi.DeleteApplicationEnvironmentVariable(context.Background(), applicationId, environmentVariableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationEnvironmentVariableApi.DeleteApplicationEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationEnvironmentVariableRequest struct via the builder pattern


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


## EditApplicationEnvironmentVariable

> EnvironmentVariable EditApplicationEnvironmentVariable(ctx, applicationId, environmentVariableId).EnvironmentVariableEditRequest(environmentVariableEditRequest).Execute()

Edit an environment variable belonging to the application



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
    environmentVariableEditRequest := *openapiclient.NewEnvironmentVariableEditRequest("Key_example") // EnvironmentVariableEditRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationEnvironmentVariableApi.EditApplicationEnvironmentVariable(context.Background(), applicationId, environmentVariableId).EnvironmentVariableEditRequest(environmentVariableEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationEnvironmentVariableApi.EditApplicationEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditApplicationEnvironmentVariable`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `ApplicationEnvironmentVariableApi.EditApplicationEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditApplicationEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environmentVariableEditRequest** | [**EnvironmentVariableEditRequest**](EnvironmentVariableEditRequest.md) |  | 

### Return type

[**EnvironmentVariable**](EnvironmentVariable.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportEnvironmentVariable

> VariableImport ImportEnvironmentVariable(ctx, applicationId).VariableImportRequest(variableImportRequest).Execute()

Import variables



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    variableImportRequest := *openapiclient.NewVariableImportRequest(false, []openapiclient.VariableImportRequestVarsInner{*openapiclient.NewVariableImportRequestVarsInner("Name_example", "Value_example", openapiclient.APIVariableScopeEnum("APPLICATION"), false)}) // VariableImportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationEnvironmentVariableApi.ImportEnvironmentVariable(context.Background(), applicationId).VariableImportRequest(variableImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationEnvironmentVariableApi.ImportEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportEnvironmentVariable`: VariableImport
    fmt.Fprintf(os.Stdout, "Response from `ApplicationEnvironmentVariableApi.ImportEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variableImportRequest** | [**VariableImportRequest**](VariableImportRequest.md) |  | 

### Return type

[**VariableImport**](VariableImport.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationEnvironmentVariable

> EnvironmentVariableResponseList ListApplicationEnvironmentVariable(ctx, applicationId).Execute()

List environment variables

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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationEnvironmentVariableApi.ListApplicationEnvironmentVariable(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationEnvironmentVariableApi.ListApplicationEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationEnvironmentVariable`: EnvironmentVariableResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationEnvironmentVariableApi.ListApplicationEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentVariableResponseList**](EnvironmentVariableResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

