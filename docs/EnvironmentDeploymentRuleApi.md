# \EnvironmentDeploymentRuleApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditEnvironmentDeploymentRule**](EnvironmentDeploymentRuleApi.md#EditEnvironmentDeploymentRule) | **Put** /environment/{environmentId}/deploymentRule/{deploymentRuleId} | Edit an environment deployment rule
[**GetEnvironmentDeploymentRule**](EnvironmentDeploymentRuleApi.md#GetEnvironmentDeploymentRule) | **Get** /environment/{environmentId}/deploymentRule | Get environment deployment rule



## EditEnvironmentDeploymentRule

> EnvironmentDeploymentRuleResponse EditEnvironmentDeploymentRule(ctx, environmentId, deploymentRuleId).EnvironmentDeploymentRuleEditRequest(environmentDeploymentRuleEditRequest).Execute()

Edit an environment deployment rule

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
    deploymentRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Rule ID
    environmentDeploymentRuleEditRequest := *openapiclient.NewEnvironmentDeploymentRuleEditRequest() // EnvironmentDeploymentRuleEditRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentDeploymentRuleApi.EditEnvironmentDeploymentRule(context.Background(), environmentId, deploymentRuleId).EnvironmentDeploymentRuleEditRequest(environmentDeploymentRuleEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentDeploymentRuleApi.EditEnvironmentDeploymentRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditEnvironmentDeploymentRule`: EnvironmentDeploymentRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentDeploymentRuleApi.EditEnvironmentDeploymentRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 
**deploymentRuleId** | **string** | Deployment Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditEnvironmentDeploymentRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environmentDeploymentRuleEditRequest** | [**EnvironmentDeploymentRuleEditRequest**](EnvironmentDeploymentRuleEditRequest.md) |  | 

### Return type

[**EnvironmentDeploymentRuleResponse**](EnvironmentDeploymentRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentDeploymentRule

> EnvironmentDeploymentRuleResponse GetEnvironmentDeploymentRule(ctx, environmentId).Execute()

Get environment deployment rule

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
    resp, r, err := apiClient.EnvironmentDeploymentRuleApi.GetEnvironmentDeploymentRule(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentDeploymentRuleApi.GetEnvironmentDeploymentRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentDeploymentRule`: EnvironmentDeploymentRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentDeploymentRuleApi.GetEnvironmentDeploymentRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentDeploymentRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentDeploymentRuleResponse**](EnvironmentDeploymentRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

