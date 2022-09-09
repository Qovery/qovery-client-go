# \ProjectSecretApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectSecret**](ProjectSecretApi.md#CreateProjectSecret) | **Post** /project/{projectId}/secret | Add a secret to the project
[**CreateProjectSecretAlias**](ProjectSecretApi.md#CreateProjectSecretAlias) | **Post** /project/{projectId}/secret/{secretId}/alias | Create a secret alias at the project level
[**CreateProjectSecretOverride**](ProjectSecretApi.md#CreateProjectSecretOverride) | **Post** /project/{projectId}/secret/{secretId}/override | Create a secret override at the project level
[**DeleteProjectSecret**](ProjectSecretApi.md#DeleteProjectSecret) | **Delete** /project/{projectId}/secret/{secretId} | Delete a secret from a project
[**EditProjectSecret**](ProjectSecretApi.md#EditProjectSecret) | **Put** /project/{projectId}/secret/{secretId} | Edit a secret belonging to the project
[**ListProjectSecrets**](ProjectSecretApi.md#ListProjectSecrets) | **Get** /project/{projectId}/secret | List project secrets



## CreateProjectSecret

> Secret CreateProjectSecret(ctx, projectId).SecretRequest(secretRequest).Execute()

Add a secret to the project



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    secretRequest := *openapiclient.NewSecretRequest("Key_example", "Value_example") // SecretRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectSecretApi.CreateProjectSecret(context.Background(), projectId).SecretRequest(secretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecretApi.CreateProjectSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `ProjectSecretApi.CreateProjectSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectSecretRequest struct via the builder pattern


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


## CreateProjectSecretAlias

> Secret CreateProjectSecretAlias(ctx, projectId, secretId).Key(key).Execute()

Create a secret alias at the project level



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    secretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Secret ID
    key := *openapiclient.NewKey("Key_example") // Key |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectSecretApi.CreateProjectSecretAlias(context.Background(), projectId, secretId).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecretApi.CreateProjectSecretAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectSecretAlias`: Secret
    fmt.Fprintf(os.Stdout, "Response from `ProjectSecretApi.CreateProjectSecretAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**secretId** | **string** | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectSecretAliasRequest struct via the builder pattern


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


## CreateProjectSecretOverride

> Secret CreateProjectSecretOverride(ctx, projectId, secretId).Value(value).Execute()

Create a secret override at the project level



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    secretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Secret ID
    value := *openapiclient.NewValue("Value_example") // Value |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectSecretApi.CreateProjectSecretOverride(context.Background(), projectId, secretId).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecretApi.CreateProjectSecretOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectSecretOverride`: Secret
    fmt.Fprintf(os.Stdout, "Response from `ProjectSecretApi.CreateProjectSecretOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**secretId** | **string** | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectSecretOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | [**Value**](Value.md) |  | 

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


## DeleteProjectSecret

> DeleteProjectSecret(ctx, projectId, secretId).Execute()

Delete a secret from a project



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    secretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Secret ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectSecretApi.DeleteProjectSecret(context.Background(), projectId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecretApi.DeleteProjectSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**secretId** | **string** | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectSecretRequest struct via the builder pattern


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


## EditProjectSecret

> Secret EditProjectSecret(ctx, projectId, secretId).SecretEditRequest(secretEditRequest).Execute()

Edit a secret belonging to the project



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    secretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Secret ID
    secretEditRequest := *openapiclient.NewSecretEditRequest("Value_example", "Key_example") // SecretEditRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectSecretApi.EditProjectSecret(context.Background(), projectId, secretId).SecretEditRequest(secretEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecretApi.EditProjectSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProjectSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `ProjectSecretApi.EditProjectSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**secretId** | **string** | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditProjectSecretRequest struct via the builder pattern


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


## ListProjectSecrets

> SecretResponseList ListProjectSecrets(ctx, projectId).Execute()

List project secrets

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectSecretApi.ListProjectSecrets(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecretApi.ListProjectSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProjectSecrets`: SecretResponseList
    fmt.Fprintf(os.Stdout, "Response from `ProjectSecretApi.ListProjectSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectSecretsRequest struct via the builder pattern


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

