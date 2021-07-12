# \EnvironmentVariableApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironmentEnvironmentVariable**](EnvironmentVariableApi.md#CreateEnvironmentEnvironmentVariable) | **Post** /environment/{environmentId}/environmentVariable | Add an environment variable to the environment
[**CreateEnvironmentEnvironmentVariableAlias**](EnvironmentVariableApi.md#CreateEnvironmentEnvironmentVariableAlias) | **Post** /environment/{environmentId}/environmentVariable/{environmentVariableId}/alias | Create an environment variable alias at the environment level
[**CreateEnvironmentEnvironmentVariableOverride**](EnvironmentVariableApi.md#CreateEnvironmentEnvironmentVariableOverride) | **Post** /environment/{environmentId}/environmentVariable/{environmentVariableId}/override | Create an environment variable override at the environment level
[**DeleteEnvironmentEnvironmentVariable**](EnvironmentVariableApi.md#DeleteEnvironmentEnvironmentVariable) | **Delete** /environment/{environmentId}/environmentVariable/{environmentVariableId} | Delete an environment variable from an environment
[**EditEnvironmentEnvironmentVariable**](EnvironmentVariableApi.md#EditEnvironmentEnvironmentVariable) | **Put** /environment/{environmentId}/environmentVariable/{environmentVariableId} | Edit an environment variable belonging to the environment
[**ListEnvironmentEnvironmentVariable**](EnvironmentVariableApi.md#ListEnvironmentEnvironmentVariable) | **Get** /environment/{environmentId}/environmentVariable | List environment variables



## CreateEnvironmentEnvironmentVariable

> EnvironmentVariableResponse CreateEnvironmentEnvironmentVariable(ctx, environmentId).EnvironmentVariableRequest(environmentVariableRequest).Execute()

Add an environment variable to the environment



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
    environmentId := TODO // string | Environment ID
    environmentVariableRequest := *openapiclient.NewEnvironmentVariableRequest("Key_example", "Value_example") // EnvironmentVariableRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentVariableApi.CreateEnvironmentEnvironmentVariable(context.Background(), environmentId).EnvironmentVariableRequest(environmentVariableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableApi.CreateEnvironmentEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentEnvironmentVariable`: EnvironmentVariableResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableApi.CreateEnvironmentEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentVariableRequest** | [**EnvironmentVariableRequest**](EnvironmentVariableRequest.md) |  | 

### Return type

[**EnvironmentVariableResponse**](EnvironmentVariableResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvironmentEnvironmentVariableAlias

> EnvironmentVariableResponse CreateEnvironmentEnvironmentVariableAlias(ctx, environmentId, environmentVariableId).Key(key).Execute()

Create an environment variable alias at the environment level



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
    environmentId := TODO // string | Environment ID
    environmentVariableId := TODO // string | Environment Variable ID
    key := *openapiclient.NewKey("Key_example") // Key |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentVariableApi.CreateEnvironmentEnvironmentVariableAlias(context.Background(), environmentId, environmentVariableId).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableApi.CreateEnvironmentEnvironmentVariableAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentEnvironmentVariableAlias`: EnvironmentVariableResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableApi.CreateEnvironmentEnvironmentVariableAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 
**environmentVariableId** | [**string**](.md) | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentEnvironmentVariableAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | [**Key**](Key.md) |  | 

### Return type

[**EnvironmentVariableResponse**](EnvironmentVariableResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvironmentEnvironmentVariableOverride

> EnvironmentVariableResponse CreateEnvironmentEnvironmentVariableOverride(ctx, environmentId, environmentVariableId).Value(value).Execute()

Create an environment variable override at the environment level



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
    environmentId := TODO // string | Environment ID
    environmentVariableId := TODO // string | Environment Variable ID
    value := *openapiclient.NewValue("Value_example") // Value |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentVariableApi.CreateEnvironmentEnvironmentVariableOverride(context.Background(), environmentId, environmentVariableId).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableApi.CreateEnvironmentEnvironmentVariableOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentEnvironmentVariableOverride`: EnvironmentVariableResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableApi.CreateEnvironmentEnvironmentVariableOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 
**environmentVariableId** | [**string**](.md) | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentEnvironmentVariableOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | [**Value**](Value.md) |  | 

### Return type

[**EnvironmentVariableResponse**](EnvironmentVariableResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentEnvironmentVariable

> DeleteEnvironmentEnvironmentVariable(ctx, environmentId, environmentVariableId).Execute()

Delete an environment variable from an environment



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
    environmentId := TODO // string | Environment ID
    environmentVariableId := TODO // string | Environment Variable ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentVariableApi.DeleteEnvironmentEnvironmentVariable(context.Background(), environmentId, environmentVariableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableApi.DeleteEnvironmentEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 
**environmentVariableId** | [**string**](.md) | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentEnvironmentVariableRequest struct via the builder pattern


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


## EditEnvironmentEnvironmentVariable

> EnvironmentVariableResponse EditEnvironmentEnvironmentVariable(ctx, environmentId, environmentVariableId).EnvironmentVariableEditRequest(environmentVariableEditRequest).Execute()

Edit an environment variable belonging to the environment



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
    environmentId := TODO // string | Environment ID
    environmentVariableId := TODO // string | Environment Variable ID
    environmentVariableEditRequest := *openapiclient.NewEnvironmentVariableEditRequest("Key_example", "Value_example") // EnvironmentVariableEditRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentVariableApi.EditEnvironmentEnvironmentVariable(context.Background(), environmentId, environmentVariableId).EnvironmentVariableEditRequest(environmentVariableEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableApi.EditEnvironmentEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditEnvironmentEnvironmentVariable`: EnvironmentVariableResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableApi.EditEnvironmentEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 
**environmentVariableId** | [**string**](.md) | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditEnvironmentEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environmentVariableEditRequest** | [**EnvironmentVariableEditRequest**](EnvironmentVariableEditRequest.md) |  | 

### Return type

[**EnvironmentVariableResponse**](EnvironmentVariableResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironmentEnvironmentVariable

> EnvironmentVariableResponseList ListEnvironmentEnvironmentVariable(ctx, environmentId).Execute()

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
    environmentId := TODO // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentVariableApi.ListEnvironmentEnvironmentVariable(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableApi.ListEnvironmentEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnvironmentEnvironmentVariable`: EnvironmentVariableResponseList
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableApi.ListEnvironmentEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | [**string**](.md) | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentEnvironmentVariableRequest struct via the builder pattern


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

