# \EnvironmentActionsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelEnvironmentDeployment**](EnvironmentActionsAPI.md#CancelEnvironmentDeployment) | **Post** /environment/{environmentId}/cancelDeployment | Cancel environment deployment
[**CleanFailedJobs**](EnvironmentActionsAPI.md#CleanFailedJobs) | **Post** /environment/{environmentId}/cleanFailedJobs | Clean failed jobs within an environment
[**CloneEnvironment**](EnvironmentActionsAPI.md#CloneEnvironment) | **Post** /environment/{environmentId}/clone | Clone environment
[**DeleteSelectedServices**](EnvironmentActionsAPI.md#DeleteSelectedServices) | **Post** /environment/{environmentId}/service/delete | Delete services
[**DeployAllServices**](EnvironmentActionsAPI.md#DeployAllServices) | **Post** /environment/{environmentId}/service/deploy | Deploy services
[**DeployEnvironment**](EnvironmentActionsAPI.md#DeployEnvironment) | **Post** /environment/{environmentId}/deploy | Deploy environment
[**RebootServices**](EnvironmentActionsAPI.md#RebootServices) | **Post** /environment/{environmentId}/service/restart-service | Reboot services
[**RedeployEnvironment**](EnvironmentActionsAPI.md#RedeployEnvironment) | **Post** /environment/{environmentId}/redeploy | Redeploy environment
[**StopEnvironment**](EnvironmentActionsAPI.md#StopEnvironment) | **Post** /environment/{environmentId}/stop | Stop environment
[**StopSelectedServices**](EnvironmentActionsAPI.md#StopSelectedServices) | **Post** /environment/{environmentId}/service/stop | Stop services
[**UninstallSelectedServices**](EnvironmentActionsAPI.md#UninstallSelectedServices) | **Post** /environment/{environmentId}/service/uninstall | Uninstall services



## CancelEnvironmentDeployment

> EnvironmentStatus CancelEnvironmentDeployment(ctx, environmentId).CancelEnvironmentDeploymentRequest(cancelEnvironmentDeploymentRequest).Execute()

Cancel environment deployment



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
	cancelEnvironmentDeploymentRequest := *openapiclient.NewCancelEnvironmentDeploymentRequest() // CancelEnvironmentDeploymentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentActionsAPI.CancelEnvironmentDeployment(context.Background(), environmentId).CancelEnvironmentDeploymentRequest(cancelEnvironmentDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsAPI.CancelEnvironmentDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelEnvironmentDeployment`: EnvironmentStatus
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsAPI.CancelEnvironmentDeployment`: %v\n", resp)
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

 **cancelEnvironmentDeploymentRequest** | [**CancelEnvironmentDeploymentRequest**](CancelEnvironmentDeploymentRequest.md) |  | 

### Return type

[**EnvironmentStatus**](EnvironmentStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CleanFailedJobs

> CleanFailedJobs200Response CleanFailedJobs(ctx, environmentId).CleanFailedJobsRequest(cleanFailedJobsRequest).Execute()

Clean failed jobs within an environment

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
	environmentId := "environmentId_example" // string | 
	cleanFailedJobsRequest := *openapiclient.NewCleanFailedJobsRequest() // CleanFailedJobsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentActionsAPI.CleanFailedJobs(context.Background(), environmentId).CleanFailedJobsRequest(cleanFailedJobsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsAPI.CleanFailedJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CleanFailedJobs`: CleanFailedJobs200Response
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsAPI.CleanFailedJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCleanFailedJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cleanFailedJobsRequest** | [**CleanFailedJobsRequest**](CleanFailedJobsRequest.md) |  | 

### Return type

[**CleanFailedJobs200Response**](CleanFailedJobs200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneEnvironment

> Environment CloneEnvironment(ctx, environmentId).CloneEnvironmentRequest(cloneEnvironmentRequest).Execute()

Clone environment



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
	cloneEnvironmentRequest := *openapiclient.NewCloneEnvironmentRequest("Name_example") // CloneEnvironmentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentActionsAPI.CloneEnvironment(context.Background(), environmentId).CloneEnvironmentRequest(cloneEnvironmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsAPI.CloneEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneEnvironment`: Environment
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsAPI.CloneEnvironment`: %v\n", resp)
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

 **cloneEnvironmentRequest** | [**CloneEnvironmentRequest**](CloneEnvironmentRequest.md) |  | 

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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
	environmentServiceIdsAllRequest := *openapiclient.NewEnvironmentServiceIdsAllRequest() // EnvironmentServiceIdsAllRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentActionsAPI.DeleteSelectedServices(context.Background(), environmentId).EnvironmentServiceIdsAllRequest(environmentServiceIdsAllRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsAPI.DeleteSelectedServices``: %v\n", err)
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

> EnvironmentStatus DeployAllServices(ctx, environmentId).DeployAllRequest(deployAllRequest).Execute()

Deploy services



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
	deployAllRequest := *openapiclient.NewDeployAllRequest() // DeployAllRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentActionsAPI.DeployAllServices(context.Background(), environmentId).DeployAllRequest(deployAllRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsAPI.DeployAllServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployAllServices`: EnvironmentStatus
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsAPI.DeployAllServices`: %v\n", resp)
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

[**EnvironmentStatus**](EnvironmentStatus.md)

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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentActionsAPI.DeployEnvironment(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsAPI.DeployEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployEnvironment`: Status
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsAPI.DeployEnvironment`: %v\n", resp)
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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
	rebootServicesRequest := *openapiclient.NewRebootServicesRequest() // RebootServicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentActionsAPI.RebootServices(context.Background(), environmentId).RebootServicesRequest(rebootServicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsAPI.RebootServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootServices`: Status
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsAPI.RebootServices`: %v\n", resp)
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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentActionsAPI.RedeployEnvironment(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsAPI.RedeployEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeployEnvironment`: EnvironmentStatus
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsAPI.RedeployEnvironment`: %v\n", resp)
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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentActionsAPI.StopEnvironment(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsAPI.StopEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopEnvironment`: EnvironmentStatus
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentActionsAPI.StopEnvironment`: %v\n", resp)
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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
	environmentServiceIdsAllRequest := *openapiclient.NewEnvironmentServiceIdsAllRequest() // EnvironmentServiceIdsAllRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentActionsAPI.StopSelectedServices(context.Background(), environmentId).EnvironmentServiceIdsAllRequest(environmentServiceIdsAllRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsAPI.StopSelectedServices``: %v\n", err)
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


## UninstallSelectedServices

> UninstallSelectedServices(ctx, environmentId).EnvironmentServiceIdsAllRequest(environmentServiceIdsAllRequest).Execute()

Uninstall services



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
	environmentServiceIdsAllRequest := *openapiclient.NewEnvironmentServiceIdsAllRequest() // EnvironmentServiceIdsAllRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentActionsAPI.UninstallSelectedServices(context.Background(), environmentId).EnvironmentServiceIdsAllRequest(environmentServiceIdsAllRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentActionsAPI.UninstallSelectedServices``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUninstallSelectedServicesRequest struct via the builder pattern


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

