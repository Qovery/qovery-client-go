# \ProjectSecretApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectSecret**](ProjectSecretApi.md#CreateProjectSecret) | **Post** /project/{projectId}/secret | Add a secret to the project
[**CreateProjectSecretAlias**](ProjectSecretApi.md#CreateProjectSecretAlias) | **Post** /project/{projectId}/secret/{secretId}/alias | Create a secret alias at the project level
[**CreateProjectSecretOverride**](ProjectSecretApi.md#CreateProjectSecretOverride) | **Post** /project/{projectId}/secret/{secretId}/override | Create a secret override at the project level
[**EditProjectSecret**](ProjectSecretApi.md#EditProjectSecret) | **Put** /project/{projectId}/secret/{secretId} | Edit a secret belonging to the project
[**ListProjectSecrets**](ProjectSecretApi.md#ListProjectSecrets) | **Get** /project/{projectId}/secret | List project secrets
[**ProjectProjectIdSecretSecretIdDelete**](ProjectSecretApi.md#ProjectProjectIdSecretSecretIdDelete) | **Delete** /project/{projectId}/secret/{secretId} | Delete a secret from a project



## CreateProjectSecret

> SecretResponse CreateProjectSecret(ctx, projectId).SecretRequest(secretRequest).Execute()

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
    projectId := TODO // string | Project ID
    secretRequest := *openapiclient.NewSecretRequest("Key_example", "Value_example") // SecretRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectSecretApi.CreateProjectSecret(context.Background(), projectId).SecretRequest(secretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecretApi.CreateProjectSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectSecret`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectSecretApi.CreateProjectSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secretRequest** | [**SecretRequest**](SecretRequest.md) |  | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectSecretAlias

> SecretResponse CreateProjectSecretAlias(ctx, projectId, secretId).Key(key).Execute()

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
    projectId := TODO // string | Project ID
    secretId := TODO // string | Secret ID
    key := *openapiclient.NewKey("Key_example") // Key |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectSecretApi.CreateProjectSecretAlias(context.Background(), projectId, secretId).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecretApi.CreateProjectSecretAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectSecretAlias`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectSecretApi.CreateProjectSecretAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project ID | 
**secretId** | [**string**](.md) | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectSecretAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | [**Key**](Key.md) |  | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectSecretOverride

> SecretResponse CreateProjectSecretOverride(ctx, projectId, secretId).Value(value).Execute()

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
    projectId := TODO // string | Project ID
    secretId := TODO // string | Secret ID
    value := *openapiclient.NewValue("Value_example") // Value |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectSecretApi.CreateProjectSecretOverride(context.Background(), projectId, secretId).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecretApi.CreateProjectSecretOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectSecretOverride`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectSecretApi.CreateProjectSecretOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project ID | 
**secretId** | [**string**](.md) | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectSecretOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | [**Value**](Value.md) |  | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditProjectSecret

> SecretResponse EditProjectSecret(ctx, projectId, secretId).SecretEditRequest(secretEditRequest).Execute()

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
    projectId := TODO // string | Project ID
    secretId := TODO // string | Secret ID
    secretEditRequest := *openapiclient.NewSecretEditRequest() // SecretEditRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectSecretApi.EditProjectSecret(context.Background(), projectId, secretId).SecretEditRequest(secretEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecretApi.EditProjectSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProjectSecret`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectSecretApi.EditProjectSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project ID | 
**secretId** | [**string**](.md) | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditProjectSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **secretEditRequest** | [**SecretEditRequest**](SecretEditRequest.md) |  | 

### Return type

[**SecretResponse**](SecretResponse.md)

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
    projectId := TODO // string | Project ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectSecretApi.ListProjectSecrets(context.Background(), projectId).Execute()
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
**projectId** | [**string**](.md) | Project ID | 

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


## ProjectProjectIdSecretSecretIdDelete

> ProjectProjectIdSecretSecretIdDelete(ctx, projectId, secretId).Execute()

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
    projectId := TODO // string | Project ID
    secretId := TODO // string | Secret ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectSecretApi.ProjectProjectIdSecretSecretIdDelete(context.Background(), projectId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecretApi.ProjectProjectIdSecretSecretIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | [**string**](.md) | Project ID | 
**secretId** | [**string**](.md) | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectProjectIdSecretSecretIdDeleteRequest struct via the builder pattern


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

