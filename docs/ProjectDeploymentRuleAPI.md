# \ProjectDeploymentRuleAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploymentRule**](ProjectDeploymentRuleAPI.md#CreateDeploymentRule) | **Post** /project/{projectId}/deploymentRule | Create a deployment rule
[**DeleteProjectDeploymentRule**](ProjectDeploymentRuleAPI.md#DeleteProjectDeploymentRule) | **Delete** /project/{projectId}/deploymentRule/{deploymentRuleId} | Delete a project deployment rule
[**EditProjectDeployemtnRule**](ProjectDeploymentRuleAPI.md#EditProjectDeployemtnRule) | **Put** /project/{projectId}/deploymentRule/{deploymentRuleId} | Edit a project deployment rule
[**GetProjectDeploymentRule**](ProjectDeploymentRuleAPI.md#GetProjectDeploymentRule) | **Get** /project/{projectId}/deploymentRule/{deploymentRuleId} | Get a project deployment rule
[**ListProjectDeploymentRules**](ProjectDeploymentRuleAPI.md#ListProjectDeploymentRules) | **Get** /project/{projectId}/deploymentRule | List project deployment rules
[**UpdateDeploymentRulesPriorityOrder**](ProjectDeploymentRuleAPI.md#UpdateDeploymentRulesPriorityOrder) | **Put** /project/{projectId}/deploymentRule/order | Update deployment rules priority order



## CreateDeploymentRule

> ProjectDeploymentRule CreateDeploymentRule(ctx, projectId).ProjectDeploymentRuleRequest(projectDeploymentRuleRequest).Execute()

Create a deployment rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    projectDeploymentRuleRequest := *openapiclient.NewProjectDeploymentRuleRequest("project-rule", openapiclient.EnvironmentModeEnum("DEVELOPMENT"), "ClusterId_example", "UTC", time.Now(), time.Now(), []openapiclient.WeekdayEnum{openapiclient.WeekdayEnum("MONDAY")}, "Wildcard_example") // ProjectDeploymentRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectDeploymentRuleAPI.CreateDeploymentRule(context.Background(), projectId).ProjectDeploymentRuleRequest(projectDeploymentRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentRuleAPI.CreateDeploymentRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeploymentRule`: ProjectDeploymentRule
    fmt.Fprintf(os.Stdout, "Response from `ProjectDeploymentRuleAPI.CreateDeploymentRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectDeploymentRuleRequest** | [**ProjectDeploymentRuleRequest**](ProjectDeploymentRuleRequest.md) |  | 

### Return type

[**ProjectDeploymentRule**](ProjectDeploymentRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectDeploymentRule

> DeleteProjectDeploymentRule(ctx, projectId, deploymentRuleId).Execute()

Delete a project deployment rule



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    deploymentRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Rule ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectDeploymentRuleAPI.DeleteProjectDeploymentRule(context.Background(), projectId, deploymentRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentRuleAPI.DeleteProjectDeploymentRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**deploymentRuleId** | **string** | Deployment Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectDeploymentRuleRequest struct via the builder pattern


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


## EditProjectDeployemtnRule

> ProjectDeploymentRule EditProjectDeployemtnRule(ctx, projectId, deploymentRuleId).ProjectDeploymentRuleRequest(projectDeploymentRuleRequest).Execute()

Edit a project deployment rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    deploymentRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Rule ID
    projectDeploymentRuleRequest := *openapiclient.NewProjectDeploymentRuleRequest("project-rule", openapiclient.EnvironmentModeEnum("DEVELOPMENT"), "ClusterId_example", "UTC", time.Now(), time.Now(), []openapiclient.WeekdayEnum{openapiclient.WeekdayEnum("MONDAY")}, "Wildcard_example") // ProjectDeploymentRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectDeploymentRuleAPI.EditProjectDeployemtnRule(context.Background(), projectId, deploymentRuleId).ProjectDeploymentRuleRequest(projectDeploymentRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentRuleAPI.EditProjectDeployemtnRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProjectDeployemtnRule`: ProjectDeploymentRule
    fmt.Fprintf(os.Stdout, "Response from `ProjectDeploymentRuleAPI.EditProjectDeployemtnRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**deploymentRuleId** | **string** | Deployment Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditProjectDeployemtnRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectDeploymentRuleRequest** | [**ProjectDeploymentRuleRequest**](ProjectDeploymentRuleRequest.md) |  | 

### Return type

[**ProjectDeploymentRule**](ProjectDeploymentRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectDeploymentRule

> ProjectDeploymentRule GetProjectDeploymentRule(ctx, projectId, deploymentRuleId).Execute()

Get a project deployment rule



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    deploymentRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Rule ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectDeploymentRuleAPI.GetProjectDeploymentRule(context.Background(), projectId, deploymentRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentRuleAPI.GetProjectDeploymentRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectDeploymentRule`: ProjectDeploymentRule
    fmt.Fprintf(os.Stdout, "Response from `ProjectDeploymentRuleAPI.GetProjectDeploymentRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 
**deploymentRuleId** | **string** | Deployment Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectDeploymentRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectDeploymentRule**](ProjectDeploymentRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectDeploymentRules

> ProjectDeploymentRuleResponseList ListProjectDeploymentRules(ctx, projectId).Execute()

List project deployment rules



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectDeploymentRuleAPI.ListProjectDeploymentRules(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentRuleAPI.ListProjectDeploymentRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProjectDeploymentRules`: ProjectDeploymentRuleResponseList
    fmt.Fprintf(os.Stdout, "Response from `ProjectDeploymentRuleAPI.ListProjectDeploymentRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectDeploymentRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectDeploymentRuleResponseList**](ProjectDeploymentRuleResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploymentRulesPriorityOrder

> UpdateDeploymentRulesPriorityOrder(ctx, projectId).ProjectDeploymentRulesPriorityOrderRequest(projectDeploymentRulesPriorityOrderRequest).Execute()

Update deployment rules priority order



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID
    projectDeploymentRulesPriorityOrderRequest := *openapiclient.NewProjectDeploymentRulesPriorityOrderRequest() // ProjectDeploymentRulesPriorityOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectDeploymentRuleAPI.UpdateDeploymentRulesPriorityOrder(context.Background(), projectId).ProjectDeploymentRulesPriorityOrderRequest(projectDeploymentRulesPriorityOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectDeploymentRuleAPI.UpdateDeploymentRulesPriorityOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentRulesPriorityOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectDeploymentRulesPriorityOrderRequest** | [**ProjectDeploymentRulesPriorityOrderRequest**](ProjectDeploymentRulesPriorityOrderRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

