# \EnvironmentDeploymentRuleApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditEnvironmentDeploymentRule**](EnvironmentDeploymentRuleApi.md#EditEnvironmentDeploymentRule) | **Put** /environment/{environmentId}/deploymentRule/{deploymentRuleId} | Edit an environment deployment rule



## EditEnvironmentDeploymentRule

> EnvironmentDeploymentRule EditEnvironmentDeploymentRule(ctx, environmentId, deploymentRuleId).EnvironmentDeploymentRuleEditRequest(environmentDeploymentRuleEditRequest).Execute()

Edit an environment deployment rule

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
    deploymentRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Rule ID
    environmentDeploymentRuleEditRequest := *openapiclient.NewEnvironmentDeploymentRuleEditRequest("UTC", time.Now(), time.Now(), []openapiclient.WeekdayEnum{openapiclient.WeekdayEnum("MONDAY")}) // EnvironmentDeploymentRuleEditRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentDeploymentRuleApi.EditEnvironmentDeploymentRule(context.Background(), environmentId, deploymentRuleId).EnvironmentDeploymentRuleEditRequest(environmentDeploymentRuleEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentDeploymentRuleApi.EditEnvironmentDeploymentRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditEnvironmentDeploymentRule`: EnvironmentDeploymentRule
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

[**EnvironmentDeploymentRule**](EnvironmentDeploymentRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

