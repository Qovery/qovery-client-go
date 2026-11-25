# \ServiceMainCallsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceDeploymentRestriction**](ServiceMainCallsAPI.md#CreateServiceDeploymentRestriction) | **Post** /service/{serviceId}/deploymentRestriction | Create a service deployment restriction
[**DeleteServiceDeploymentRestriction**](ServiceMainCallsAPI.md#DeleteServiceDeploymentRestriction) | **Delete** /service/{serviceId}/deploymentRestriction/{deploymentRestrictionId} | Delete a service deployment restriction
[**EditServiceDeploymentRestriction**](ServiceMainCallsAPI.md#EditServiceDeploymentRestriction) | **Put** /service/{serviceId}/deploymentRestriction/{deploymentRestrictionId} | Edit a service deployment restriction
[**GetServiceGitWebhookStatus**](ServiceMainCallsAPI.md#GetServiceGitWebhookStatus) | **Get** /service/{serviceId}/gitWebhookStatus | Get git webhook status for a service
[**ListServiceDeploymentRestrictions**](ServiceMainCallsAPI.md#ListServiceDeploymentRestrictions) | **Get** /service/{serviceId}/deploymentRestriction | List service deployment restrictions
[**SyncServiceGitWebhook**](ServiceMainCallsAPI.md#SyncServiceGitWebhook) | **Post** /service/{serviceId}/gitWebhook/sync | Synchronize git webhook for a service



## CreateServiceDeploymentRestriction

> ServiceDeploymentRestrictionResponse CreateServiceDeploymentRestriction(ctx, serviceId).ServiceDeploymentRestrictionRequest(serviceDeploymentRestrictionRequest).Execute()

Create a service deployment restriction



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
	serviceDeploymentRestrictionRequest := *openapiclient.NewServiceDeploymentRestrictionRequest(openapiclient.DeploymentRestrictionModeEnum("EXCLUDE"), openapiclient.DeploymentRestrictionTypeEnum("PATH"), "service1/src/") // ServiceDeploymentRestrictionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceMainCallsAPI.CreateServiceDeploymentRestriction(context.Background(), serviceId).ServiceDeploymentRestrictionRequest(serviceDeploymentRestrictionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMainCallsAPI.CreateServiceDeploymentRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceDeploymentRestriction`: ServiceDeploymentRestrictionResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceMainCallsAPI.CreateServiceDeploymentRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID of an application/job/container/database | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceDeploymentRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceDeploymentRestrictionRequest** | [**ServiceDeploymentRestrictionRequest**](ServiceDeploymentRestrictionRequest.md) |  | 

### Return type

[**ServiceDeploymentRestrictionResponse**](ServiceDeploymentRestrictionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceDeploymentRestriction

> DeleteServiceDeploymentRestriction(ctx, serviceId, deploymentRestrictionId).Execute()

Delete a service deployment restriction



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
	deploymentRestrictionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Restriction ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceMainCallsAPI.DeleteServiceDeploymentRestriction(context.Background(), serviceId, deploymentRestrictionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMainCallsAPI.DeleteServiceDeploymentRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID of an application/job/container/database | 
**deploymentRestrictionId** | **string** | Deployment Restriction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceDeploymentRestrictionRequest struct via the builder pattern


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


## EditServiceDeploymentRestriction

> ServiceDeploymentRestrictionResponse EditServiceDeploymentRestriction(ctx, serviceId, deploymentRestrictionId).ServiceDeploymentRestrictionRequest(serviceDeploymentRestrictionRequest).Execute()

Edit a service deployment restriction



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
	deploymentRestrictionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment Restriction ID
	serviceDeploymentRestrictionRequest := *openapiclient.NewServiceDeploymentRestrictionRequest(openapiclient.DeploymentRestrictionModeEnum("EXCLUDE"), openapiclient.DeploymentRestrictionTypeEnum("PATH"), "service1/src/") // ServiceDeploymentRestrictionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceMainCallsAPI.EditServiceDeploymentRestriction(context.Background(), serviceId, deploymentRestrictionId).ServiceDeploymentRestrictionRequest(serviceDeploymentRestrictionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMainCallsAPI.EditServiceDeploymentRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditServiceDeploymentRestriction`: ServiceDeploymentRestrictionResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceMainCallsAPI.EditServiceDeploymentRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID of an application/job/container/database | 
**deploymentRestrictionId** | **string** | Deployment Restriction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditServiceDeploymentRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceDeploymentRestrictionRequest** | [**ServiceDeploymentRestrictionRequest**](ServiceDeploymentRestrictionRequest.md) |  | 

### Return type

[**ServiceDeploymentRestrictionResponse**](ServiceDeploymentRestrictionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceGitWebhookStatus

> GitWebhookStatusResponse GetServiceGitWebhookStatus(ctx, serviceId).Execute()

Get git webhook status for a service



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
	resp, r, err := apiClient.ServiceMainCallsAPI.GetServiceGitWebhookStatus(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMainCallsAPI.GetServiceGitWebhookStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceGitWebhookStatus`: GitWebhookStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceMainCallsAPI.GetServiceGitWebhookStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID of an application/job/container/database | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceGitWebhookStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitWebhookStatusResponse**](GitWebhookStatusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceDeploymentRestrictions

> ServiceDeploymentRestrictionResponseList ListServiceDeploymentRestrictions(ctx, serviceId).Execute()

List service deployment restrictions



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
	resp, r, err := apiClient.ServiceMainCallsAPI.ListServiceDeploymentRestrictions(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMainCallsAPI.ListServiceDeploymentRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceDeploymentRestrictions`: ServiceDeploymentRestrictionResponseList
	fmt.Fprintf(os.Stdout, "Response from `ServiceMainCallsAPI.ListServiceDeploymentRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID of an application/job/container/database | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceDeploymentRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDeploymentRestrictionResponseList**](ServiceDeploymentRestrictionResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncServiceGitWebhook

> GitWebhookStatusResponse SyncServiceGitWebhook(ctx, serviceId).Execute()

Synchronize git webhook for a service



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
	resp, r, err := apiClient.ServiceMainCallsAPI.SyncServiceGitWebhook(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceMainCallsAPI.SyncServiceGitWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncServiceGitWebhook`: GitWebhookStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceMainCallsAPI.SyncServiceGitWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID of an application/job/container/database | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncServiceGitWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitWebhookStatusResponse**](GitWebhookStatusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

