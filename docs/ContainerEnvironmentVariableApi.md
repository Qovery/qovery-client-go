# \ContainerEnvironmentVariableApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContainerEnvironmentVariable**](ContainerEnvironmentVariableApi.md#CreateContainerEnvironmentVariable) | **Post** /container/{containerId}/environmentVariable | Add an environment variable to the container
[**CreateContainerEnvironmentVariableAlias**](ContainerEnvironmentVariableApi.md#CreateContainerEnvironmentVariableAlias) | **Post** /container/{containerId}/environmentVariable/{environmentVariableId}/alias | Create an environment variable alias at the container level
[**CreateContainerEnvironmentVariableOverride**](ContainerEnvironmentVariableApi.md#CreateContainerEnvironmentVariableOverride) | **Post** /container/{containerId}/environmentVariable/{environmentVariableId}/override | Create an environment variable override at the container level
[**DeleteContainerEnvironmentVariable**](ContainerEnvironmentVariableApi.md#DeleteContainerEnvironmentVariable) | **Delete** /container/{containerId}/environmentVariable/{environmentVariableId} | Delete an environment variable from a container
[**EditContainerEnvironmentVariable**](ContainerEnvironmentVariableApi.md#EditContainerEnvironmentVariable) | **Put** /container/{containerId}/environmentVariable/{environmentVariableId} | Edit an environment variable belonging to the container
[**ImportContainerEnvironmentVariable**](ContainerEnvironmentVariableApi.md#ImportContainerEnvironmentVariable) | **Post** /container/{containerId}/environmentVariable/import | Import variables
[**ListContainerEnvironmentVariable**](ContainerEnvironmentVariableApi.md#ListContainerEnvironmentVariable) | **Get** /container/{containerId}/environmentVariable | List environment variables



## CreateContainerEnvironmentVariable

> EnvironmentVariable CreateContainerEnvironmentVariable(ctx, containerId).EnvironmentVariableRequest(environmentVariableRequest).Execute()

Add an environment variable to the container



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    environmentVariableRequest := *openapiclient.NewEnvironmentVariableRequest("Key_example", "Value_example") // EnvironmentVariableRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerEnvironmentVariableApi.CreateContainerEnvironmentVariable(context.Background(), containerId).EnvironmentVariableRequest(environmentVariableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerEnvironmentVariableApi.CreateContainerEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainerEnvironmentVariable`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `ContainerEnvironmentVariableApi.CreateContainerEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerEnvironmentVariableRequest struct via the builder pattern


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


## CreateContainerEnvironmentVariableAlias

> EnvironmentVariable CreateContainerEnvironmentVariableAlias(ctx, containerId, environmentVariableId).Key(key).Execute()

Create an environment variable alias at the container level



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
    key := *openapiclient.NewKey("Key_example") // Key |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerEnvironmentVariableApi.CreateContainerEnvironmentVariableAlias(context.Background(), containerId, environmentVariableId).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerEnvironmentVariableApi.CreateContainerEnvironmentVariableAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainerEnvironmentVariableAlias`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `ContainerEnvironmentVariableApi.CreateContainerEnvironmentVariableAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerEnvironmentVariableAliasRequest struct via the builder pattern


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


## CreateContainerEnvironmentVariableOverride

> EnvironmentVariable CreateContainerEnvironmentVariableOverride(ctx, containerId, environmentVariableId).Value(value).Execute()

Create an environment variable override at the container level



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
    value := *openapiclient.NewValue("Value_example") // Value |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerEnvironmentVariableApi.CreateContainerEnvironmentVariableOverride(context.Background(), containerId, environmentVariableId).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerEnvironmentVariableApi.CreateContainerEnvironmentVariableOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainerEnvironmentVariableOverride`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `ContainerEnvironmentVariableApi.CreateContainerEnvironmentVariableOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerEnvironmentVariableOverrideRequest struct via the builder pattern


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


## DeleteContainerEnvironmentVariable

> DeleteContainerEnvironmentVariable(ctx, containerId, environmentVariableId).Execute()

Delete an environment variable from a container



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerEnvironmentVariableApi.DeleteContainerEnvironmentVariable(context.Background(), containerId, environmentVariableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerEnvironmentVariableApi.DeleteContainerEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerEnvironmentVariableRequest struct via the builder pattern


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


## EditContainerEnvironmentVariable

> EnvironmentVariable EditContainerEnvironmentVariable(ctx, containerId, environmentVariableId).EnvironmentVariableEditRequest(environmentVariableEditRequest).Execute()

Edit an environment variable belonging to the container



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
    environmentVariableEditRequest := *openapiclient.NewEnvironmentVariableEditRequest("Key_example", "Value_example") // EnvironmentVariableEditRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerEnvironmentVariableApi.EditContainerEnvironmentVariable(context.Background(), containerId, environmentVariableId).EnvironmentVariableEditRequest(environmentVariableEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerEnvironmentVariableApi.EditContainerEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditContainerEnvironmentVariable`: EnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `ContainerEnvironmentVariableApi.EditContainerEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditContainerEnvironmentVariableRequest struct via the builder pattern


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


## ImportContainerEnvironmentVariable

> VariableImport ImportContainerEnvironmentVariable(ctx, containerId).VariableImportRequest(variableImportRequest).Execute()

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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    variableImportRequest := *openapiclient.NewVariableImportRequest(false, []openapiclient.VariableImportRequestVars{*openapiclient.NewVariableImportRequestVars("Name_example", "Value_example", openapiclient.EnvironmentVariableScopeEnum("APPLICATION"), false)}) // VariableImportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerEnvironmentVariableApi.ImportContainerEnvironmentVariable(context.Background(), containerId).VariableImportRequest(variableImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerEnvironmentVariableApi.ImportContainerEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportContainerEnvironmentVariable`: VariableImport
    fmt.Fprintf(os.Stdout, "Response from `ContainerEnvironmentVariableApi.ImportContainerEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportContainerEnvironmentVariableRequest struct via the builder pattern


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


## ListContainerEnvironmentVariable

> EnvironmentVariableResponseList ListContainerEnvironmentVariable(ctx, containerId).Execute()

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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerEnvironmentVariableApi.ListContainerEnvironmentVariable(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerEnvironmentVariableApi.ListContainerEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerEnvironmentVariable`: EnvironmentVariableResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerEnvironmentVariableApi.ListContainerEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerEnvironmentVariableRequest struct via the builder pattern


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

