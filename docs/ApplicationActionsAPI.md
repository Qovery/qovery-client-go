# \ApplicationActionsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployApplication**](ApplicationActionsAPI.md#DeployApplication) | **Post** /application/{applicationId}/deploy | Deploy application
[**RebootApplication**](ApplicationActionsAPI.md#RebootApplication) | **Post** /application/{applicationId}/restart-service | Reboot application
[**RedeployApplication**](ApplicationActionsAPI.md#RedeployApplication) | **Post** /application/{applicationId}/redeploy | Redeploy application
[**StopApplication**](ApplicationActionsAPI.md#StopApplication) | **Post** /application/{applicationId}/stop | Stop application
[**UninstallApplication**](ApplicationActionsAPI.md#UninstallApplication) | **Post** /application/{applicationId}/uninstall | Uninstall application



## DeployApplication

> Status DeployApplication(ctx, applicationId).DeployRequest(deployRequest).Execute()

Deploy application



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
	applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
	deployRequest := *openapiclient.NewDeployRequest("GitCommitId_example") // DeployRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationActionsAPI.DeployApplication(context.Background(), applicationId).DeployRequest(deployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationActionsAPI.DeployApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployApplication`: Status
	fmt.Fprintf(os.Stdout, "Response from `ApplicationActionsAPI.DeployApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deployRequest** | [**DeployRequest**](DeployRequest.md) |  | 

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


## RebootApplication

> Status RebootApplication(ctx, applicationId).Execute()

Reboot application

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
	applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationActionsAPI.RebootApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationActionsAPI.RebootApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootApplication`: Status
	fmt.Fprintf(os.Stdout, "Response from `ApplicationActionsAPI.RebootApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootApplicationRequest struct via the builder pattern


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


## RedeployApplication

> Status RedeployApplication(ctx, applicationId).Execute()

Redeploy application

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
	applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationActionsAPI.RedeployApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationActionsAPI.RedeployApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeployApplication`: Status
	fmt.Fprintf(os.Stdout, "Response from `ApplicationActionsAPI.RedeployApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployApplicationRequest struct via the builder pattern


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


## StopApplication

> Status StopApplication(ctx, applicationId).Execute()

Stop application

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
	applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationActionsAPI.StopApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationActionsAPI.StopApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopApplication`: Status
	fmt.Fprintf(os.Stdout, "Response from `ApplicationActionsAPI.StopApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopApplicationRequest struct via the builder pattern


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


## UninstallApplication

> Status UninstallApplication(ctx, applicationId).Body(body).Execute()

Uninstall application



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
	applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationActionsAPI.UninstallApplication(context.Background(), applicationId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationActionsAPI.UninstallApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UninstallApplication`: Status
	fmt.Fprintf(os.Stdout, "Response from `ApplicationActionsAPI.UninstallApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

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

