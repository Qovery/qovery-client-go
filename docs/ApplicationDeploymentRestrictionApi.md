# \ApplicationDeploymentRestrictionApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationDeploymentRestriction**](ApplicationDeploymentRestrictionApi.md#CreateApplicationDeploymentRestriction) | **Post** /application/{applicationId}/deploymentRestriction | Create an application deployment restriction
[**DeleteApplicationDeploymentRestriction**](ApplicationDeploymentRestrictionApi.md#DeleteApplicationDeploymentRestriction) | **Delete** /application/{applicationId}/deploymentRestriction/{deploymentRestrictionId} | Delete an application deployment restriction
[**EditApplicationDeploymentRestriction**](ApplicationDeploymentRestrictionApi.md#EditApplicationDeploymentRestriction) | **Put** /application/{applicationId}/deploymentRestriction/{deploymentRestrictionId} | Edit an application deployment restriction
[**GetApplicationDeploymentRestrictions**](ApplicationDeploymentRestrictionApi.md#GetApplicationDeploymentRestrictions) | **Get** /application/{applicationId}/deploymentRestriction | Get application deployment restrictions



## CreateApplicationDeploymentRestriction

> ApplicationDeploymentRestriction CreateApplicationDeploymentRestriction(ctx, applicationId).ApplicationDeploymentRestrictionRequest(applicationDeploymentRestrictionRequest).Execute()

Create an application deployment restriction



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
    applicationDeploymentRestrictionRequest := *openapiclient.NewApplicationDeploymentRestrictionRequest(openapiclient.DeploymentRestrictionModeEnum("EXCLUDE"), openapiclient.DeploymentRestrictionTypeEnum("PATH"), "application1/src/") // ApplicationDeploymentRestrictionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationDeploymentRestrictionApi.CreateApplicationDeploymentRestriction(context.Background(), applicationId).ApplicationDeploymentRestrictionRequest(applicationDeploymentRestrictionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDeploymentRestrictionApi.CreateApplicationDeploymentRestriction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationDeploymentRestriction`: ApplicationDeploymentRestriction
    fmt.Fprintf(os.Stdout, "Response from `ApplicationDeploymentRestrictionApi.CreateApplicationDeploymentRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationDeploymentRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDeploymentRestrictionRequest** | [**ApplicationDeploymentRestrictionRequest**](ApplicationDeploymentRestrictionRequest.md) |  | 

### Return type

[**ApplicationDeploymentRestriction**](ApplicationDeploymentRestriction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationDeploymentRestriction

> DeleteApplicationDeploymentRestriction(ctx, applicationId, deploymentRestrictionId).Execute()

Delete an application deployment restriction



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
    deploymentRestrictionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Restriction ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationDeploymentRestrictionApi.DeleteApplicationDeploymentRestriction(context.Background(), applicationId, deploymentRestrictionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDeploymentRestrictionApi.DeleteApplicationDeploymentRestriction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**deploymentRestrictionId** | **string** | Deployment Restriction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationDeploymentRestrictionRequest struct via the builder pattern


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


## EditApplicationDeploymentRestriction

> ApplicationDeploymentRestriction EditApplicationDeploymentRestriction(ctx, applicationId, deploymentRestrictionId).ApplicationDeploymentRestrictionRequest(applicationDeploymentRestrictionRequest).Execute()

Edit an application deployment restriction



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
    deploymentRestrictionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Restriction ID
    applicationDeploymentRestrictionRequest := *openapiclient.NewApplicationDeploymentRestrictionRequest(openapiclient.DeploymentRestrictionModeEnum("EXCLUDE"), openapiclient.DeploymentRestrictionTypeEnum("PATH"), "application1/src/") // ApplicationDeploymentRestrictionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationDeploymentRestrictionApi.EditApplicationDeploymentRestriction(context.Background(), applicationId, deploymentRestrictionId).ApplicationDeploymentRestrictionRequest(applicationDeploymentRestrictionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDeploymentRestrictionApi.EditApplicationDeploymentRestriction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditApplicationDeploymentRestriction`: ApplicationDeploymentRestriction
    fmt.Fprintf(os.Stdout, "Response from `ApplicationDeploymentRestrictionApi.EditApplicationDeploymentRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**deploymentRestrictionId** | **string** | Deployment Restriction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditApplicationDeploymentRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationDeploymentRestrictionRequest** | [**ApplicationDeploymentRestrictionRequest**](ApplicationDeploymentRestrictionRequest.md) |  | 

### Return type

[**ApplicationDeploymentRestriction**](ApplicationDeploymentRestriction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationDeploymentRestrictions

> ApplicationDeploymentRestrictionResponseList GetApplicationDeploymentRestrictions(ctx, applicationId).Execute()

Get application deployment restrictions



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
    resp, r, err := apiClient.ApplicationDeploymentRestrictionApi.GetApplicationDeploymentRestrictions(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDeploymentRestrictionApi.GetApplicationDeploymentRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationDeploymentRestrictions`: ApplicationDeploymentRestrictionResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationDeploymentRestrictionApi.GetApplicationDeploymentRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationDeploymentRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationDeploymentRestrictionResponseList**](ApplicationDeploymentRestrictionResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

