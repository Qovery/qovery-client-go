# \DeploymentStageMainCallsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachServiceToDeploymentStage**](DeploymentStageMainCallsAPI.md#AttachServiceToDeploymentStage) | **Put** /deploymentStage/{deploymentStageId}/service/{serviceId} | Attach service to deployment stage
[**CreateEnvironmentDeploymentStage**](DeploymentStageMainCallsAPI.md#CreateEnvironmentDeploymentStage) | **Post** /environment/{environmentId}/deploymentStage | Create environment deployment stage
[**DeleteDeploymentStage**](DeploymentStageMainCallsAPI.md#DeleteDeploymentStage) | **Delete** /deploymentStage/{deploymentStageId} | Delete deployment stage
[**EditDeploymentStage**](DeploymentStageMainCallsAPI.md#EditDeploymentStage) | **Put** /deploymentStage/{deploymentStageId} | Edit deployment stage
[**GetDeploymentStage**](DeploymentStageMainCallsAPI.md#GetDeploymentStage) | **Get** /deploymentStage/{deploymentStageId} | Get Deployment Stage
[**GetServiceDeploymentStage**](DeploymentStageMainCallsAPI.md#GetServiceDeploymentStage) | **Get** /service/{serviceId}/deploymentStage | Get Service Deployment Stage
[**ListEnvironmentDeploymentStage**](DeploymentStageMainCallsAPI.md#ListEnvironmentDeploymentStage) | **Get** /environment/{environmentId}/deploymentStage | List environment deployment stage
[**MoveAfterDeploymentStage**](DeploymentStageMainCallsAPI.md#MoveAfterDeploymentStage) | **Put** /deploymentStage/{deploymentStageId}/moveAfter/{stageId} | Move deployment stage after requested stage
[**MoveBeforeDeploymentStage**](DeploymentStageMainCallsAPI.md#MoveBeforeDeploymentStage) | **Put** /deploymentStage/{deploymentStageId}/moveBefore/{stageId} | Move deployment stage before requested stage



## AttachServiceToDeploymentStage

> DeploymentStageResponseList AttachServiceToDeploymentStage(ctx, deploymentStageId, serviceId).Execute()

Attach service to deployment stage

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
	deploymentStageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Stage ID
	serviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service ID of an application/job/container/database

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentStageMainCallsAPI.AttachServiceToDeploymentStage(context.Background(), deploymentStageId, serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentStageMainCallsAPI.AttachServiceToDeploymentStage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachServiceToDeploymentStage`: DeploymentStageResponseList
	fmt.Fprintf(os.Stdout, "Response from `DeploymentStageMainCallsAPI.AttachServiceToDeploymentStage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentStageId** | **string** | Deployment Stage ID | 
**serviceId** | **string** | Service ID of an application/job/container/database | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachServiceToDeploymentStageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeploymentStageResponseList**](DeploymentStageResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvironmentDeploymentStage

> DeploymentStageResponse CreateEnvironmentDeploymentStage(ctx, environmentId).DeploymentStageRequest(deploymentStageRequest).Execute()

Create environment deployment stage

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
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
	deploymentStageRequest := *openapiclient.NewDeploymentStageRequest("Name_example") // DeploymentStageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentStageMainCallsAPI.CreateEnvironmentDeploymentStage(context.Background(), environmentId).DeploymentStageRequest(deploymentStageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentStageMainCallsAPI.CreateEnvironmentDeploymentStage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnvironmentDeploymentStage`: DeploymentStageResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploymentStageMainCallsAPI.CreateEnvironmentDeploymentStage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentDeploymentStageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentStageRequest** | [**DeploymentStageRequest**](DeploymentStageRequest.md) |  | 

### Return type

[**DeploymentStageResponse**](DeploymentStageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeploymentStage

> DeleteDeploymentStage(ctx, deploymentStageId).Execute()

Delete deployment stage

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
	deploymentStageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Stage ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentStageMainCallsAPI.DeleteDeploymentStage(context.Background(), deploymentStageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentStageMainCallsAPI.DeleteDeploymentStage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentStageId** | **string** | Deployment Stage ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentStageRequest struct via the builder pattern


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


## EditDeploymentStage

> DeploymentStageResponse EditDeploymentStage(ctx, deploymentStageId).DeploymentStageRequest(deploymentStageRequest).Execute()

Edit deployment stage

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
	deploymentStageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Stage ID
	deploymentStageRequest := *openapiclient.NewDeploymentStageRequest("Name_example") // DeploymentStageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentStageMainCallsAPI.EditDeploymentStage(context.Background(), deploymentStageId).DeploymentStageRequest(deploymentStageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentStageMainCallsAPI.EditDeploymentStage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditDeploymentStage`: DeploymentStageResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploymentStageMainCallsAPI.EditDeploymentStage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentStageId** | **string** | Deployment Stage ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditDeploymentStageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentStageRequest** | [**DeploymentStageRequest**](DeploymentStageRequest.md) |  | 

### Return type

[**DeploymentStageResponse**](DeploymentStageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentStage

> DeploymentStageResponse GetDeploymentStage(ctx, deploymentStageId).Execute()

Get Deployment Stage

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
	deploymentStageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Stage ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentStageMainCallsAPI.GetDeploymentStage(context.Background(), deploymentStageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentStageMainCallsAPI.GetDeploymentStage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentStage`: DeploymentStageResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploymentStageMainCallsAPI.GetDeploymentStage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentStageId** | **string** | Deployment Stage ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentStageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeploymentStageResponse**](DeploymentStageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceDeploymentStage

> DeploymentStageResponse GetServiceDeploymentStage(ctx, serviceId).Execute()

Get Service Deployment Stage

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
	serviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service ID of an application/job/container/database

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentStageMainCallsAPI.GetServiceDeploymentStage(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentStageMainCallsAPI.GetServiceDeploymentStage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceDeploymentStage`: DeploymentStageResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploymentStageMainCallsAPI.GetServiceDeploymentStage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID of an application/job/container/database | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceDeploymentStageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeploymentStageResponse**](DeploymentStageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironmentDeploymentStage

> DeploymentStageResponseList ListEnvironmentDeploymentStage(ctx, environmentId).Execute()

List environment deployment stage

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
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentStageMainCallsAPI.ListEnvironmentDeploymentStage(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentStageMainCallsAPI.ListEnvironmentDeploymentStage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnvironmentDeploymentStage`: DeploymentStageResponseList
	fmt.Fprintf(os.Stdout, "Response from `DeploymentStageMainCallsAPI.ListEnvironmentDeploymentStage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentDeploymentStageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeploymentStageResponseList**](DeploymentStageResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveAfterDeploymentStage

> DeploymentStageResponseList MoveAfterDeploymentStage(ctx, deploymentStageId, stageId).Execute()

Move deployment stage after requested stage

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
	deploymentStageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Stage ID
	stageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Stage ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentStageMainCallsAPI.MoveAfterDeploymentStage(context.Background(), deploymentStageId, stageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentStageMainCallsAPI.MoveAfterDeploymentStage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveAfterDeploymentStage`: DeploymentStageResponseList
	fmt.Fprintf(os.Stdout, "Response from `DeploymentStageMainCallsAPI.MoveAfterDeploymentStage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentStageId** | **string** | Deployment Stage ID | 
**stageId** | **string** | Deployment Stage ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveAfterDeploymentStageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeploymentStageResponseList**](DeploymentStageResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveBeforeDeploymentStage

> DeploymentStageResponseList MoveBeforeDeploymentStage(ctx, deploymentStageId, stageId).Execute()

Move deployment stage before requested stage

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
	deploymentStageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Stage ID
	stageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Stage ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentStageMainCallsAPI.MoveBeforeDeploymentStage(context.Background(), deploymentStageId, stageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentStageMainCallsAPI.MoveBeforeDeploymentStage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveBeforeDeploymentStage`: DeploymentStageResponseList
	fmt.Fprintf(os.Stdout, "Response from `DeploymentStageMainCallsAPI.MoveBeforeDeploymentStage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentStageId** | **string** | Deployment Stage ID | 
**stageId** | **string** | Deployment Stage ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveBeforeDeploymentStageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeploymentStageResponseList**](DeploymentStageResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

