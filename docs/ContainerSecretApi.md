# \ContainerSecretApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContainerSecret**](ContainerSecretApi.md#CreateContainerSecret) | **Post** /container/{containerId}/secret | NOT YET IMPLEMENTED - Add a secret to the container
[**CreateContainerSecretAlias**](ContainerSecretApi.md#CreateContainerSecretAlias) | **Post** /container/{containerId}/secret/{secretId}/alias | NOT YET IMPLEMENTED - Create a secret alias at the container level
[**DeleteContainerSecret**](ContainerSecretApi.md#DeleteContainerSecret) | **Delete** /container/{containerId}/secret/{secretId} | NOT YET IMPLEMENTED - Delete a secret from an container
[**EditContainerSecret**](ContainerSecretApi.md#EditContainerSecret) | **Put** /container/{containerId}/secret/{secretId} | NOT YET IMPLEMENTED - Edit a secret belonging to the container
[**ListContainerSecrets**](ContainerSecretApi.md#ListContainerSecrets) | **Get** /container/{containerId}/secret | NOT YET IMPLEMENTED - List container secrets



## CreateContainerSecret

> Secret CreateContainerSecret(ctx, containerId).SecretRequest(secretRequest).Execute()

NOT YET IMPLEMENTED - Add a secret to the container



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
    secretRequest := *openapiclient.NewSecretRequest("Key_example", "Value_example") // SecretRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerSecretApi.CreateContainerSecret(context.Background(), containerId).SecretRequest(secretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerSecretApi.CreateContainerSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainerSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `ContainerSecretApi.CreateContainerSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secretRequest** | [**SecretRequest**](SecretRequest.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContainerSecretAlias

> Secret CreateContainerSecretAlias(ctx, containerId, secretId).Key(key).Execute()

NOT YET IMPLEMENTED - Create a secret alias at the container level



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
    secretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Secret ID
    key := *openapiclient.NewKey("Key_example") // Key |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerSecretApi.CreateContainerSecretAlias(context.Background(), containerId, secretId).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerSecretApi.CreateContainerSecretAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainerSecretAlias`: Secret
    fmt.Fprintf(os.Stdout, "Response from `ContainerSecretApi.CreateContainerSecretAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**secretId** | **string** | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerSecretAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | [**Key**](Key.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerSecret

> DeleteContainerSecret(ctx, containerId, secretId).Execute()

NOT YET IMPLEMENTED - Delete a secret from an container



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
    secretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Secret ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerSecretApi.DeleteContainerSecret(context.Background(), containerId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerSecretApi.DeleteContainerSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**secretId** | **string** | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerSecretRequest struct via the builder pattern


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


## EditContainerSecret

> Secret EditContainerSecret(ctx, containerId, secretId).SecretEditRequest(secretEditRequest).Execute()

NOT YET IMPLEMENTED - Edit a secret belonging to the container



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
    secretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Secret ID
    secretEditRequest := *openapiclient.NewSecretEditRequest() // SecretEditRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerSecretApi.EditContainerSecret(context.Background(), containerId, secretId).SecretEditRequest(secretEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerSecretApi.EditContainerSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditContainerSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `ContainerSecretApi.EditContainerSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**secretId** | **string** | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditContainerSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **secretEditRequest** | [**SecretEditRequest**](SecretEditRequest.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerSecrets

> SecretResponseList ListContainerSecrets(ctx, containerId).Execute()

NOT YET IMPLEMENTED - List container secrets



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
    resp, r, err := apiClient.ContainerSecretApi.ListContainerSecrets(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerSecretApi.ListContainerSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerSecrets`: SecretResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerSecretApi.ListContainerSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretResponseList**](SecretResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

