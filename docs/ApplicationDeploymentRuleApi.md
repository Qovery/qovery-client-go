# \ApplicationDeploymentRuleApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditApplicationDeploymentRule**](ApplicationDeploymentRuleApi.md#EditApplicationDeploymentRule) | **Put** /application/{applicationId}/deploymentRule/{deploymentRuleId} | Edit an application deployment rule
[**GetApplicationDeploymentRule**](ApplicationDeploymentRuleApi.md#GetApplicationDeploymentRule) | **Get** /application/{applicationId}/deploymentRule | Get application deployment rule



## EditApplicationDeploymentRule

> ApplicationDeploymentRuleResponse EditApplicationDeploymentRule(ctx, applicationId, deploymentRuleId).ApplicationDeploymentRuleEditRequest(applicationDeploymentRuleEditRequest).Execute()

Edit an application deployment rule



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
    deploymentRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Rule ID
    applicationDeploymentRuleEditRequest := *openapiclient.NewApplicationDeploymentRuleEditRequest() // ApplicationDeploymentRuleEditRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationDeploymentRuleApi.EditApplicationDeploymentRule(context.Background(), applicationId, deploymentRuleId).ApplicationDeploymentRuleEditRequest(applicationDeploymentRuleEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDeploymentRuleApi.EditApplicationDeploymentRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditApplicationDeploymentRule`: ApplicationDeploymentRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationDeploymentRuleApi.EditApplicationDeploymentRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**deploymentRuleId** | **string** | Deployment Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditApplicationDeploymentRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationDeploymentRuleEditRequest** | [**ApplicationDeploymentRuleEditRequest**](ApplicationDeploymentRuleEditRequest.md) |  | 

### Return type

[**ApplicationDeploymentRuleResponse**](ApplicationDeploymentRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationDeploymentRule

> ApplicationDeploymentRuleResponse GetApplicationDeploymentRule(ctx, applicationId).Execute()

Get application deployment rule



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
    resp, r, err := apiClient.ApplicationDeploymentRuleApi.GetApplicationDeploymentRule(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDeploymentRuleApi.GetApplicationDeploymentRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationDeploymentRule`: ApplicationDeploymentRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationDeploymentRuleApi.GetApplicationDeploymentRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationDeploymentRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationDeploymentRuleResponse**](ApplicationDeploymentRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

