# \HelmDeploymentRestrictionAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHelmDeploymentRestriction**](HelmDeploymentRestrictionAPI.md#CreateHelmDeploymentRestriction) | **Post** /helm/{helmId}/deploymentRestriction | Create a helm deployment restriction
[**DeleteHelmDeploymentRestriction**](HelmDeploymentRestrictionAPI.md#DeleteHelmDeploymentRestriction) | **Delete** /helm/{helmId}/deploymentRestriction/{deploymentRestrictionId} | Delete a helm deployment restriction
[**EditHelmDeploymentRestriction**](HelmDeploymentRestrictionAPI.md#EditHelmDeploymentRestriction) | **Put** /helm/{helmId}/deploymentRestriction/{deploymentRestrictionId} | Edit a helm deployment restriction
[**GetHelmDeploymentRestrictions**](HelmDeploymentRestrictionAPI.md#GetHelmDeploymentRestrictions) | **Get** /helm/{helmId}/deploymentRestriction | Get helm deployment restrictions



## CreateHelmDeploymentRestriction

> HelmDeploymentRestrictionResponse CreateHelmDeploymentRestriction(ctx, helmId).HelmDeploymentRestrictionRequest(helmDeploymentRestrictionRequest).Execute()

Create a helm deployment restriction



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
    helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID
    helmDeploymentRestrictionRequest := *openapiclient.NewHelmDeploymentRestrictionRequest(openapiclient.DeploymentRestrictionModeEnum("EXCLUDE"), openapiclient.DeploymentRestrictionTypeEnum("PATH"), "helm1/src/") // HelmDeploymentRestrictionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmDeploymentRestrictionAPI.CreateHelmDeploymentRestriction(context.Background(), helmId).HelmDeploymentRestrictionRequest(helmDeploymentRestrictionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmDeploymentRestrictionAPI.CreateHelmDeploymentRestriction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHelmDeploymentRestriction`: HelmDeploymentRestrictionResponse
    fmt.Fprintf(os.Stdout, "Response from `HelmDeploymentRestrictionAPI.CreateHelmDeploymentRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHelmDeploymentRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **helmDeploymentRestrictionRequest** | [**HelmDeploymentRestrictionRequest**](HelmDeploymentRestrictionRequest.md) |  | 

### Return type

[**HelmDeploymentRestrictionResponse**](HelmDeploymentRestrictionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHelmDeploymentRestriction

> DeleteHelmDeploymentRestriction(ctx, helmId, deploymentRestrictionId).Execute()

Delete a helm deployment restriction



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
    helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID
    deploymentRestrictionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Restriction ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HelmDeploymentRestrictionAPI.DeleteHelmDeploymentRestriction(context.Background(), helmId, deploymentRestrictionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmDeploymentRestrictionAPI.DeleteHelmDeploymentRestriction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 
**deploymentRestrictionId** | **string** | Deployment Restriction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHelmDeploymentRestrictionRequest struct via the builder pattern


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


## EditHelmDeploymentRestriction

> HelmDeploymentRestrictionResponse EditHelmDeploymentRestriction(ctx, helmId, deploymentRestrictionId).HelmDeploymentRestrictionRequest(helmDeploymentRestrictionRequest).Execute()

Edit a helm deployment restriction



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
    helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID
    deploymentRestrictionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Restriction ID
    helmDeploymentRestrictionRequest := *openapiclient.NewHelmDeploymentRestrictionRequest(openapiclient.DeploymentRestrictionModeEnum("EXCLUDE"), openapiclient.DeploymentRestrictionTypeEnum("PATH"), "helm1/src/") // HelmDeploymentRestrictionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmDeploymentRestrictionAPI.EditHelmDeploymentRestriction(context.Background(), helmId, deploymentRestrictionId).HelmDeploymentRestrictionRequest(helmDeploymentRestrictionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmDeploymentRestrictionAPI.EditHelmDeploymentRestriction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditHelmDeploymentRestriction`: HelmDeploymentRestrictionResponse
    fmt.Fprintf(os.Stdout, "Response from `HelmDeploymentRestrictionAPI.EditHelmDeploymentRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 
**deploymentRestrictionId** | **string** | Deployment Restriction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditHelmDeploymentRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **helmDeploymentRestrictionRequest** | [**HelmDeploymentRestrictionRequest**](HelmDeploymentRestrictionRequest.md) |  | 

### Return type

[**HelmDeploymentRestrictionResponse**](HelmDeploymentRestrictionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHelmDeploymentRestrictions

> HelmDeploymentRestrictionResponseList GetHelmDeploymentRestrictions(ctx, helmId).Execute()

Get helm deployment restrictions



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
    helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmDeploymentRestrictionAPI.GetHelmDeploymentRestrictions(context.Background(), helmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmDeploymentRestrictionAPI.GetHelmDeploymentRestrictions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHelmDeploymentRestrictions`: HelmDeploymentRestrictionResponseList
    fmt.Fprintf(os.Stdout, "Response from `HelmDeploymentRestrictionAPI.GetHelmDeploymentRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHelmDeploymentRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HelmDeploymentRestrictionResponseList**](HelmDeploymentRestrictionResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

