# \ContainerActionsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployContainer**](ContainerActionsAPI.md#DeployContainer) | **Post** /container/{containerId}/deploy | Deploy container
[**RebootContainer**](ContainerActionsAPI.md#RebootContainer) | **Post** /container/{containerId}/restart-service | Reboot container
[**RedeployContainer**](ContainerActionsAPI.md#RedeployContainer) | **Post** /container/{containerId}/redeploy | Redeploy container
[**StopContainer**](ContainerActionsAPI.md#StopContainer) | **Post** /container/{containerId}/stop | Stop container
[**UninstallContainer**](ContainerActionsAPI.md#UninstallContainer) | **Post** /container/{containerId}/uninstall | Uninstall container



## DeployContainer

> Status DeployContainer(ctx, containerId).ContainerDeployRequest(containerDeployRequest).Execute()

Deploy container



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
	containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
	containerDeployRequest := *openapiclient.NewContainerDeployRequest("ImageTag_example") // ContainerDeployRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerActionsAPI.DeployContainer(context.Background(), containerId).ContainerDeployRequest(containerDeployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerActionsAPI.DeployContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployContainer`: Status
	fmt.Fprintf(os.Stdout, "Response from `ContainerActionsAPI.DeployContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerDeployRequest** | [**ContainerDeployRequest**](ContainerDeployRequest.md) |  | 

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


## RebootContainer

> Status RebootContainer(ctx, containerId).Execute()

Reboot container

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
	containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerActionsAPI.RebootContainer(context.Background(), containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerActionsAPI.RebootContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootContainer`: Status
	fmt.Fprintf(os.Stdout, "Response from `ContainerActionsAPI.RebootContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootContainerRequest struct via the builder pattern


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


## RedeployContainer

> Status RedeployContainer(ctx, containerId).Execute()

Redeploy container

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
	containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerActionsAPI.RedeployContainer(context.Background(), containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerActionsAPI.RedeployContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeployContainer`: Status
	fmt.Fprintf(os.Stdout, "Response from `ContainerActionsAPI.RedeployContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployContainerRequest struct via the builder pattern


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


## StopContainer

> Status StopContainer(ctx, containerId).Execute()

Stop container

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
	containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerActionsAPI.StopContainer(context.Background(), containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerActionsAPI.StopContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopContainer`: Status
	fmt.Fprintf(os.Stdout, "Response from `ContainerActionsAPI.StopContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopContainerRequest struct via the builder pattern


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


## UninstallContainer

> Status UninstallContainer(ctx, containerId).Body(body).Execute()

Uninstall container



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
	containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerActionsAPI.UninstallContainer(context.Background(), containerId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerActionsAPI.UninstallContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UninstallContainer`: Status
	fmt.Fprintf(os.Stdout, "Response from `ContainerActionsAPI.UninstallContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallContainerRequest struct via the builder pattern


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

