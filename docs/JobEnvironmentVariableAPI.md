# \JobEnvironmentVariableAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJobEnvironmentVariable**](JobEnvironmentVariableAPI.md#CreateJobEnvironmentVariable) | **Post** /job/{jobId}/environmentVariable | Add an environment variable to the job
[**CreateJobEnvironmentVariableAlias**](JobEnvironmentVariableAPI.md#CreateJobEnvironmentVariableAlias) | **Post** /job/{jobId}/environmentVariable/{environmentVariableId}/alias | Create an environment variable alias at the job level
[**CreateJobEnvironmentVariableOverride**](JobEnvironmentVariableAPI.md#CreateJobEnvironmentVariableOverride) | **Post** /job/{jobId}/environmentVariable/{environmentVariableId}/override | Create an environment variable override at the job level
[**DeleteJobEnvironmentVariable**](JobEnvironmentVariableAPI.md#DeleteJobEnvironmentVariable) | **Delete** /job/{jobId}/environmentVariable/{environmentVariableId} | Delete an environment variable from a job
[**EditJobEnvironmentVariable**](JobEnvironmentVariableAPI.md#EditJobEnvironmentVariable) | **Put** /job/{jobId}/environmentVariable/{environmentVariableId} | Edit an environment variable belonging to the job
[**ImportJobEnvironmentVariable**](JobEnvironmentVariableAPI.md#ImportJobEnvironmentVariable) | **Post** /job/{jobId}/environmentVariable/import | Import variables
[**ListJobEnvironmentVariable**](JobEnvironmentVariableAPI.md#ListJobEnvironmentVariable) | **Get** /job/{jobId}/environmentVariable | List environment variables



## CreateJobEnvironmentVariable

> EnvironmentVariable CreateJobEnvironmentVariable(ctx, jobId).EnvironmentVariableRequest(environmentVariableRequest).Execute()

Add an environment variable to the job



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
    environmentVariableRequest := *openapiclient.NewEnvironmentVariableRequest("Key_example") // EnvironmentVariableRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobEnvironmentVariableAPI.CreateJobEnvironmentVariable(context.Background(), jobId).EnvironmentVariableRequest(environmentVariableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobEnvironmentVariableAPI.CreateJobEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJobEnvironmentVariable`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `JobEnvironmentVariableAPI.CreateJobEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobEnvironmentVariableRequest struct via the builder pattern


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


## CreateJobEnvironmentVariableAlias

> EnvironmentVariable CreateJobEnvironmentVariableAlias(ctx, jobId, environmentVariableId).Key(key).Execute()

Create an environment variable alias at the job level



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
    key := *openapiclient.NewKey("Key_example") // Key |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobEnvironmentVariableAPI.CreateJobEnvironmentVariableAlias(context.Background(), jobId, environmentVariableId).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobEnvironmentVariableAPI.CreateJobEnvironmentVariableAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJobEnvironmentVariableAlias`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `JobEnvironmentVariableAPI.CreateJobEnvironmentVariableAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobEnvironmentVariableAliasRequest struct via the builder pattern


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


## CreateJobEnvironmentVariableOverride

> EnvironmentVariable CreateJobEnvironmentVariableOverride(ctx, jobId, environmentVariableId).Value(value).Execute()

Create an environment variable override at the job level



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
    value := *openapiclient.NewValue() // Value |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobEnvironmentVariableAPI.CreateJobEnvironmentVariableOverride(context.Background(), jobId, environmentVariableId).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobEnvironmentVariableAPI.CreateJobEnvironmentVariableOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJobEnvironmentVariableOverride`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `JobEnvironmentVariableAPI.CreateJobEnvironmentVariableOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobEnvironmentVariableOverrideRequest struct via the builder pattern


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


## DeleteJobEnvironmentVariable

> DeleteJobEnvironmentVariable(ctx, jobId, environmentVariableId).Execute()

Delete an environment variable from a job



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.JobEnvironmentVariableAPI.DeleteJobEnvironmentVariable(context.Background(), jobId, environmentVariableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobEnvironmentVariableAPI.DeleteJobEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobEnvironmentVariableRequest struct via the builder pattern


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


## EditJobEnvironmentVariable

> EnvironmentVariable EditJobEnvironmentVariable(ctx, jobId, environmentVariableId).EnvironmentVariableEditRequest(environmentVariableEditRequest).Execute()

Edit an environment variable belonging to the job



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
    environmentVariableEditRequest := *openapiclient.NewEnvironmentVariableEditRequest("Key_example") // EnvironmentVariableEditRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobEnvironmentVariableAPI.EditJobEnvironmentVariable(context.Background(), jobId, environmentVariableId).EnvironmentVariableEditRequest(environmentVariableEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobEnvironmentVariableAPI.EditJobEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditJobEnvironmentVariable`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `JobEnvironmentVariableAPI.EditJobEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditJobEnvironmentVariableRequest struct via the builder pattern


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


## ImportJobEnvironmentVariable

> VariableImport ImportJobEnvironmentVariable(ctx, jobId).VariableImportRequest(variableImportRequest).Execute()

Import variables



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
    variableImportRequest := *openapiclient.NewVariableImportRequest(false, []openapiclient.VariableImportRequestVarsInner{*openapiclient.NewVariableImportRequestVarsInner("Name_example", "Value_example", openapiclient.APIVariableScopeEnum("APPLICATION"), false)}) // VariableImportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobEnvironmentVariableAPI.ImportJobEnvironmentVariable(context.Background(), jobId).VariableImportRequest(variableImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobEnvironmentVariableAPI.ImportJobEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportJobEnvironmentVariable`: VariableImport
    fmt.Fprintf(os.Stdout, "Response from `JobEnvironmentVariableAPI.ImportJobEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportJobEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variableImportRequest** | [**VariableImportRequest**](VariableImportRequest.md) |  | 

### Return type

[**VariableImport**](VariableImport.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJobEnvironmentVariable

> EnvironmentVariableResponseList ListJobEnvironmentVariable(ctx, jobId).Execute()

List environment variables

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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobEnvironmentVariableAPI.ListJobEnvironmentVariable(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobEnvironmentVariableAPI.ListJobEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJobEnvironmentVariable`: EnvironmentVariableResponseList
    fmt.Fprintf(os.Stdout, "Response from `JobEnvironmentVariableAPI.ListJobEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListJobEnvironmentVariableRequest struct via the builder pattern


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

