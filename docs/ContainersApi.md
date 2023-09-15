# \ContainersApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoDeployContainerEnvironments**](ContainersApi.md#AutoDeployContainerEnvironments) | **Post** /organization/{organizationId}/container/deploy | Auto deploy containers
[**CloneContainer**](ContainersApi.md#CloneContainer) | **Post** /container/{containerId}/clone | Clone container
[**CreateContainer**](ContainersApi.md#CreateContainer) | **Post** /environment/{environmentId}/container | Create a container
[**GetContainerRegistryContainerStatus**](ContainersApi.md#GetContainerRegistryContainerStatus) | **Get** /organization/{organizationId}/containerRegistry/{containerRegistryId}/container/status | List all container registry container statuses
[**GetDefaultContainerAdvancedSettings**](ContainersApi.md#GetDefaultContainerAdvancedSettings) | **Get** /defaultContainerAdvancedSettings | List default container advanced settings
[**GetEnvironmentContainerCurrentScale**](ContainersApi.md#GetEnvironmentContainerCurrentScale) | **Get** /environment/{environmentId}/container/currentScale | List current scaling information for each container
[**GetEnvironmentContainerCurrentStorage**](ContainersApi.md#GetEnvironmentContainerCurrentStorage) | **Get** /environment/{environmentId}/container/currentStorage | List current storage disk usage for each containers
[**GetEnvironmentContainerStatus**](ContainersApi.md#GetEnvironmentContainerStatus) | **Get** /environment/{environmentId}/container/status | List all environment container statuses
[**ListContainer**](ContainersApi.md#ListContainer) | **Get** /environment/{environmentId}/container | List containers
[**PreviewContainerEnvironments**](ContainersApi.md#PreviewContainerEnvironments) | **Post** /organization/{organizationId}/container/preview | Preview container environments



## AutoDeployContainerEnvironments

> Status AutoDeployContainerEnvironments(ctx, organizationId).OrganizationContainerAutoDeployRequest(organizationContainerAutoDeployRequest).Execute()

Auto deploy containers



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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    organizationContainerAutoDeployRequest := *openapiclient.NewOrganizationContainerAutoDeployRequest() // OrganizationContainerAutoDeployRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.AutoDeployContainerEnvironments(context.Background(), organizationId).OrganizationContainerAutoDeployRequest(organizationContainerAutoDeployRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.AutoDeployContainerEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutoDeployContainerEnvironments`: Status
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.AutoDeployContainerEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoDeployContainerEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationContainerAutoDeployRequest** | [**OrganizationContainerAutoDeployRequest**](OrganizationContainerAutoDeployRequest.md) |  | 

### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneContainer

> ContainerResponse CloneContainer(ctx, containerId).CloneContainerRequest(cloneContainerRequest).Execute()

Clone container



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
    cloneContainerRequest := *openapiclient.NewCloneContainerRequest("Name_example", "EnvironmentId_example") // CloneContainerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.CloneContainer(context.Background(), containerId).CloneContainerRequest(cloneContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.CloneContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneContainer`: ContainerResponse
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.CloneContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneContainerRequest** | [**CloneContainerRequest**](CloneContainerRequest.md) |  | 

### Return type

[**ContainerResponse**](ContainerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContainer

> ContainerResponse CreateContainer(ctx, environmentId).ContainerRequest(containerRequest).Execute()

Create a container

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
    containerRequest := *openapiclient.NewContainerRequest("Name_example", "RegistryId_example", "ImageName_example", "Tag_example", *openapiclient.NewHealthcheck()) // ContainerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.CreateContainer(context.Background(), environmentId).ContainerRequest(containerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.CreateContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainer`: ContainerResponse
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.CreateContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerRequest** | [**ContainerRequest**](ContainerRequest.md) |  | 

### Return type

[**ContainerResponse**](ContainerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerRegistryContainerStatus

> ReferenceObjectStatusResponseList GetContainerRegistryContainerStatus(ctx, organizationId, containerRegistryId).Execute()

List all container registry container statuses



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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    containerRegistryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container Registry ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.GetContainerRegistryContainerStatus(context.Background(), organizationId, containerRegistryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.GetContainerRegistryContainerStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerRegistryContainerStatus`: ReferenceObjectStatusResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.GetContainerRegistryContainerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**containerRegistryId** | **string** | Container Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerRegistryContainerStatusRequest struct via the builder pattern


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


## GetDefaultContainerAdvancedSettings

> ContainerAdvancedSettings GetDefaultContainerAdvancedSettings(ctx).Execute()

List default container advanced settings

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.GetDefaultContainerAdvancedSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.GetDefaultContainerAdvancedSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultContainerAdvancedSettings`: ContainerAdvancedSettings
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.GetDefaultContainerAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultContainerAdvancedSettingsRequest struct via the builder pattern


### Return type

[**ContainerAdvancedSettings**](ContainerAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentContainerCurrentScale

> EnvironmentContainersCurrentScaleResponseList GetEnvironmentContainerCurrentScale(ctx, environmentId).Execute()

List current scaling information for each container



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.GetEnvironmentContainerCurrentScale(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.GetEnvironmentContainerCurrentScale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentContainerCurrentScale`: EnvironmentContainersCurrentScaleResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.GetEnvironmentContainerCurrentScale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentContainerCurrentScaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentContainersCurrentScaleResponseList**](EnvironmentContainersCurrentScaleResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentContainerCurrentStorage

> EnvironmentContainersStorageResponseList GetEnvironmentContainerCurrentStorage(ctx, environmentId).Execute()

List current storage disk usage for each containers

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.GetEnvironmentContainerCurrentStorage(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.GetEnvironmentContainerCurrentStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentContainerCurrentStorage`: EnvironmentContainersStorageResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.GetEnvironmentContainerCurrentStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentContainerCurrentStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentContainersStorageResponseList**](EnvironmentContainersStorageResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentContainerStatus

> ReferenceObjectStatusResponseList GetEnvironmentContainerStatus(ctx, environmentId).Execute()

List all environment container statuses



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.GetEnvironmentContainerStatus(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.GetEnvironmentContainerStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentContainerStatus`: ReferenceObjectStatusResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.GetEnvironmentContainerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentContainerStatusRequest struct via the builder pattern


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


## ListContainer

> ContainerResponseList ListContainer(ctx, environmentId).Execute()

List containers

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.ListContainer(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.ListContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainer`: ContainerResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.ListContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerResponseList**](ContainerResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewContainerEnvironments

> Status PreviewContainerEnvironments(ctx, organizationId).OrganizationContainerPreviewRequest(organizationContainerPreviewRequest).Execute()

Preview container environments



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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    organizationContainerPreviewRequest := *openapiclient.NewOrganizationContainerPreviewRequest() // OrganizationContainerPreviewRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.PreviewContainerEnvironments(context.Background(), organizationId).OrganizationContainerPreviewRequest(organizationContainerPreviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.PreviewContainerEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewContainerEnvironments`: Status
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.PreviewContainerEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewContainerEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationContainerPreviewRequest** | [**OrganizationContainerPreviewRequest**](OrganizationContainerPreviewRequest.md) |  | 

### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

