# \EnvironmentActionsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelEnvironmentDeployment**](EnvironmentActionsApi.md#CancelEnvironmentDeployment) | **Post** /environment/{environmentId}/cancelDeployment | Cancel environment deployment
[**CloneEnvironment**](EnvironmentActionsApi.md#CloneEnvironment) | **Post** /environment/{environmentId}/clone | Clone environment
[**DeleteSelectedServices**](EnvironmentActionsApi.md#DeleteSelectedServices) | **Post** /environment/{environmentId}/service/delete | Delete services
[**DeployAllServices**](EnvironmentActionsApi.md#DeployAllServices) | **Post** /environment/{environmentId}/service/deploy | Deploy services
[**DeployEnvironment**](EnvironmentActionsApi.md#DeployEnvironment) | **Post** /environment/{environmentId}/deploy | Deploy environment
[**RebootServices**](EnvironmentActionsApi.md#RebootServices) | **Post** /environment/{environmentId}/service/restart-service | Reboot services
[**RedeployEnvironment**](EnvironmentActionsApi.md#RedeployEnvironment) | **Post** /environment/{environmentId}/redeploy | Redeploy environment
[**RestartEnvironment**](EnvironmentActionsApi.md#RestartEnvironment) | **Post** /environment/{environmentId}/restart | Deprecated - Restart environment
[**StopEnvironment**](EnvironmentActionsApi.md#StopEnvironment) | **Post** /environment/{environmentId}/stop | Stop environment
[**StopSelectedServices**](EnvironmentActionsApi.md#StopSelectedServices) | **Post** /environment/{environmentId}/service/stop | Stop services



## CancelEnvironmentDeployment

> EnvironmentStatus CancelEnvironmentDeployment(ctx, environmentId).Execute()

Cancel environment deployment



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
    resp, r, err := apiClient.EnvironmentActionsApi.CancelEnvironmentDeployment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsApi.CancelEnvironmentDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelEnvironmentDeployment`: EnvironmentStatus
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsApi.CancelEnvironmentDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelEnvironmentDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentStatus**](EnvironmentStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneEnvironment

> Environment CloneEnvironment(ctx, environmentId).CloneRequest(cloneRequest).Execute()

Clone environment



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
    cloneRequest := *openapiclient.NewCloneRequest("Name_example") // CloneRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentActionsApi.CloneEnvironment(context.Background(), environmentId).CloneRequest(cloneRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsApi.CloneEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsApi.CloneEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneRequest** | [**CloneRequest**](CloneRequest.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSelectedServices

> DeleteSelectedServices(ctx, environmentId).EnvironmentServiceIdsAllRequest(environmentServiceIdsAllRequest).Execute()

Delete services



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
    environmentServiceIdsAllRequest := *openapiclient.NewEnvironmentServiceIdsAllRequest() // EnvironmentServiceIdsAllRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentActionsApi.DeleteSelectedServices(context.Background(), environmentId).EnvironmentServiceIdsAllRequest(environmentServiceIdsAllRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsApi.DeleteSelectedServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSelectedServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentServiceIdsAllRequest** | [**EnvironmentServiceIdsAllRequest**](EnvironmentServiceIdsAllRequest.md) |  | 

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


## DeployAllServices

> Status DeployAllServices(ctx, environmentId).DeployAllRequest(deployAllRequest).Execute()

Deploy services



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
    deployAllRequest := *openapiclient.NewDeployAllRequest() // DeployAllRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentActionsApi.DeployAllServices(context.Background(), environmentId).DeployAllRequest(deployAllRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsApi.DeployAllServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployAllServices`: Status
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsApi.DeployAllServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployAllServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deployAllRequest** | [**DeployAllRequest**](DeployAllRequest.md) |  | 

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


## DeployEnvironment

> Status DeployEnvironment(ctx, environmentId).Execute()

Deploy environment



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
    resp, r, err := apiClient.EnvironmentActionsApi.DeployEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsApi.DeployEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployEnvironment`: Status
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsApi.DeployEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootServices

> Status RebootServices(ctx, environmentId).RebootServicesRequest(rebootServicesRequest).Execute()

Reboot services



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
    rebootServicesRequest := *openapiclient.NewRebootServicesRequest() // RebootServicesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentActionsApi.RebootServices(context.Background(), environmentId).RebootServicesRequest(rebootServicesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsApi.RebootServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootServices`: Status
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsApi.RebootServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rebootServicesRequest** | [**RebootServicesRequest**](RebootServicesRequest.md) |  | 

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


## RedeployEnvironment

> EnvironmentStatus RedeployEnvironment(ctx, environmentId).Execute()

Redeploy environment

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
    resp, r, err := apiClient.EnvironmentActionsApi.RedeployEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsApi.RedeployEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RedeployEnvironment`: EnvironmentStatus
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsApi.RedeployEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentStatus**](EnvironmentStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartEnvironment

> EnvironmentStatus RestartEnvironment(ctx, environmentId).Execute()

Deprecated - Restart environment



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
    resp, r, err := apiClient.EnvironmentActionsApi.RestartEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsApi.RestartEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartEnvironment`: EnvironmentStatus
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsApi.RestartEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentStatus**](EnvironmentStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopEnvironment

> EnvironmentStatus StopEnvironment(ctx, environmentId).Execute()

Stop environment

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
    resp, r, err := apiClient.EnvironmentActionsApi.StopEnvironment(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsApi.StopEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopEnvironment`: EnvironmentStatus
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsApi.StopEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentStatus**](EnvironmentStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopSelectedServices

> StopSelectedServices(ctx, environmentId).EnvironmentServiceIdsAllRequest(environmentServiceIdsAllRequest).Execute()

Stop services



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
    environmentServiceIdsAllRequest := *openapiclient.NewEnvironmentServiceIdsAllRequest() // EnvironmentServiceIdsAllRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentActionsApi.StopSelectedServices(context.Background(), environmentId).EnvironmentServiceIdsAllRequest(environmentServiceIdsAllRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsApi.StopSelectedServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopSelectedServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentServiceIdsAllRequest** | [**EnvironmentServiceIdsAllRequest**](EnvironmentServiceIdsAllRequest.md) |  | 

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

