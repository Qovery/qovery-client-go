# \ServiceMainCallsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceGitWebhookStatus**](ServiceMainCallsAPI.md#GetServiceGitWebhookStatus) | **Get** /service/{serviceId}/gitWebhookStatus | Get git webhook status for a service
[**SyncServiceGitWebhook**](ServiceMainCallsAPI.md#SyncServiceGitWebhook) | **Post** /service/{serviceId}/gitWebhook/sync | Synchronize git webhook for a service



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

