# \ContainerDeploymentRestrictionApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContainerDeploymentRestriction**](ContainerDeploymentRestrictionApi.md#CreateContainerDeploymentRestriction) | **Post** /container/{containerId}/deploymentRestriction | NOT YET IMPLEMENTED - Create an container deployment restriction
[**DeleteContainerDeploymentRestriction**](ContainerDeploymentRestrictionApi.md#DeleteContainerDeploymentRestriction) | **Delete** /container/{containerId}/deploymentRestriction/{deploymentRestrictionId} | NOT YET IMPLEMENTED - Delete a container deployment restriction
[**EditContainerDeploymentRestriction**](ContainerDeploymentRestrictionApi.md#EditContainerDeploymentRestriction) | **Put** /container/{containerId}/deploymentRestriction/{deploymentRestrictionId} | NOT YET IMPLEMENTED - Edit a container deployment restriction
[**GetContainerDeploymentRestrictions**](ContainerDeploymentRestrictionApi.md#GetContainerDeploymentRestrictions) | **Get** /container/{containerId}/deploymentRestriction | NOT YET IMPLEMENTED - Get container deployment restrictions



## CreateContainerDeploymentRestriction

> ContainerDeploymentRestriction CreateContainerDeploymentRestriction(ctx, containerId).ContainerDeploymentRestrictionRequest(containerDeploymentRestrictionRequest).Execute()

NOT YET IMPLEMENTED - Create an container deployment restriction



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
    containerDeploymentRestrictionRequest := *openapiclient.NewContainerDeploymentRestrictionRequest(openapiclient.DeploymentRestrictionModeEnum("EXCLUDE"), openapiclient.DeploymentRestrictionTypeEnum("PATH"), "application1/src/") // ContainerDeploymentRestrictionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerDeploymentRestrictionApi.CreateContainerDeploymentRestriction(context.Background(), containerId).ContainerDeploymentRestrictionRequest(containerDeploymentRestrictionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerDeploymentRestrictionApi.CreateContainerDeploymentRestriction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainerDeploymentRestriction`: ContainerDeploymentRestriction
    fmt.Fprintf(os.Stdout, "Response from `ContainerDeploymentRestrictionApi.CreateContainerDeploymentRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerDeploymentRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerDeploymentRestrictionRequest** | [**ContainerDeploymentRestrictionRequest**](ContainerDeploymentRestrictionRequest.md) |  | 

### Return type

[**ContainerDeploymentRestriction**](ContainerDeploymentRestriction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerDeploymentRestriction

> DeleteContainerDeploymentRestriction(ctx, containerId).Execute()

NOT YET IMPLEMENTED - Delete a container deployment restriction



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
    resp, r, err := apiClient.ContainerDeploymentRestrictionApi.DeleteContainerDeploymentRestriction(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerDeploymentRestrictionApi.DeleteContainerDeploymentRestriction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerDeploymentRestrictionRequest struct via the builder pattern


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


## EditContainerDeploymentRestriction

> ContainerDeploymentRestriction EditContainerDeploymentRestriction(ctx, containerId, deploymentRestrictionId).ContainerDeploymentRestrictionRequest(containerDeploymentRestrictionRequest).Execute()

NOT YET IMPLEMENTED - Edit a container deployment restriction



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
    deploymentRestrictionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Restriction ID
    containerDeploymentRestrictionRequest := *openapiclient.NewContainerDeploymentRestrictionRequest(openapiclient.DeploymentRestrictionModeEnum("EXCLUDE"), openapiclient.DeploymentRestrictionTypeEnum("PATH"), "application1/src/") // ContainerDeploymentRestrictionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerDeploymentRestrictionApi.EditContainerDeploymentRestriction(context.Background(), containerId, deploymentRestrictionId).ContainerDeploymentRestrictionRequest(containerDeploymentRestrictionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerDeploymentRestrictionApi.EditContainerDeploymentRestriction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditContainerDeploymentRestriction`: ContainerDeploymentRestriction
    fmt.Fprintf(os.Stdout, "Response from `ContainerDeploymentRestrictionApi.EditContainerDeploymentRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**deploymentRestrictionId** | **string** | Deployment Restriction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditContainerDeploymentRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **containerDeploymentRestrictionRequest** | [**ContainerDeploymentRestrictionRequest**](ContainerDeploymentRestrictionRequest.md) |  | 

### Return type

[**ContainerDeploymentRestriction**](ContainerDeploymentRestriction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerDeploymentRestrictions

> ContainerDeploymentRestrictionResponseList GetContainerDeploymentRestrictions(ctx, containerId).Execute()

NOT YET IMPLEMENTED - Get container deployment restrictions



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
    resp, r, err := apiClient.ContainerDeploymentRestrictionApi.GetContainerDeploymentRestrictions(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerDeploymentRestrictionApi.GetContainerDeploymentRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerDeploymentRestrictions`: ContainerDeploymentRestrictionResponseList
    fmt.Fprintf(os.Stdout, "Response from `ContainerDeploymentRestrictionApi.GetContainerDeploymentRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerDeploymentRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerDeploymentRestrictionResponseList**](ContainerDeploymentRestrictionResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

