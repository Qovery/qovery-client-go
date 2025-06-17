# \HelmActionsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployHelm**](HelmActionsAPI.md#DeployHelm) | **Post** /helm/{helmId}/deploy | Deploy helm
[**RedeployHelm**](HelmActionsAPI.md#RedeployHelm) | **Post** /helm/{helmId}/redeploy | Redeploy helm
[**StopHelm**](HelmActionsAPI.md#StopHelm) | **Post** /helm/{helmId}/stop | Stop helm
[**UninstallHelm**](HelmActionsAPI.md#UninstallHelm) | **Post** /helm/{helmId}/uninstall | Uninstall helm



## DeployHelm

> Status DeployHelm(ctx, helmId).ForceEvent(forceEvent).HelmDeployRequest(helmDeployRequest).Execute()

Deploy helm



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
	forceEvent := openapiclient.HelmForceEvent("DIFF") // HelmForceEvent | When filled, it indicates the target event to be deployed.   If the concerned helm hasn't the target event provided, the helm won't be deployed.  (optional)
	helmDeployRequest := *openapiclient.NewHelmDeployRequest() // HelmDeployRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmActionsAPI.DeployHelm(context.Background(), helmId).ForceEvent(forceEvent).HelmDeployRequest(helmDeployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmActionsAPI.DeployHelm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployHelm`: Status
	fmt.Fprintf(os.Stdout, "Response from `HelmActionsAPI.DeployHelm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployHelmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceEvent** | [**HelmForceEvent**](HelmForceEvent.md) | When filled, it indicates the target event to be deployed.   If the concerned helm hasn&#39;t the target event provided, the helm won&#39;t be deployed.  | 
 **helmDeployRequest** | [**HelmDeployRequest**](HelmDeployRequest.md) |  | 

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


## RedeployHelm

> Status RedeployHelm(ctx, helmId).ForceEvent(forceEvent).Execute()

Redeploy helm

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
	forceEvent := openapiclient.HelmForceEvent("DIFF") // HelmForceEvent | When filled, it indicates the target event to be deployed.   If the concerned helm hasn't the target event provided, the helm won't be deployed.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmActionsAPI.RedeployHelm(context.Background(), helmId).ForceEvent(forceEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmActionsAPI.RedeployHelm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeployHelm`: Status
	fmt.Fprintf(os.Stdout, "Response from `HelmActionsAPI.RedeployHelm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployHelmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceEvent** | [**HelmForceEvent**](HelmForceEvent.md) | When filled, it indicates the target event to be deployed.   If the concerned helm hasn&#39;t the target event provided, the helm won&#39;t be deployed.  | 

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


## StopHelm

> Status StopHelm(ctx, helmId).Execute()

Stop helm

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
	resp, r, err := apiClient.HelmActionsAPI.StopHelm(context.Background(), helmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmActionsAPI.StopHelm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopHelm`: Status
	fmt.Fprintf(os.Stdout, "Response from `HelmActionsAPI.StopHelm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopHelmRequest struct via the builder pattern


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


## UninstallHelm

> map[string]interface{} UninstallHelm(ctx, helmId).Body(body).Execute()

Uninstall helm



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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmActionsAPI.UninstallHelm(context.Background(), helmId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmActionsAPI.UninstallHelm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UninstallHelm`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `HelmActionsAPI.UninstallHelm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallHelmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

