# \AutoscalingPolicyAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateServiceAutoscaling**](AutoscalingPolicyAPI.md#CreateOrUpdateServiceAutoscaling) | **Put** /service/{serviceId}/autoscalingPolicy | createOrUpdateServiceAutoscaling
[**DeleteServiceAutoscaling**](AutoscalingPolicyAPI.md#DeleteServiceAutoscaling) | **Delete** /service/{serviceId}/autoscalingPolicy | deleteServiceAutoscaling
[**GetServiceAutoscaling**](AutoscalingPolicyAPI.md#GetServiceAutoscaling) | **Get** /service/{serviceId}/autoscalingPolicy | getServiceAutoscaling



## CreateOrUpdateServiceAutoscaling

> AutoscalingPolicyResponse CreateOrUpdateServiceAutoscaling(ctx, serviceId).AutoscalingPolicyRequest(autoscalingPolicyRequest).Execute()

createOrUpdateServiceAutoscaling

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
	autoscalingPolicyRequest := openapiclient.AutoscalingPolicyRequest{KedaAutoscalingRequest: openapiclient.NewKedaAutoscalingRequest(openapiclient.AutoscalingMode("KEDA"), []openapiclient.KedaScalerRequest{*openapiclient.NewKedaScalerRequest("ScalerType_example", openapiclient.KedaScalerRole("PRIMARY"))})} // AutoscalingPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoscalingPolicyAPI.CreateOrUpdateServiceAutoscaling(context.Background(), serviceId).AutoscalingPolicyRequest(autoscalingPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoscalingPolicyAPI.CreateOrUpdateServiceAutoscaling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrUpdateServiceAutoscaling`: AutoscalingPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoscalingPolicyAPI.CreateOrUpdateServiceAutoscaling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID of an application/job/container/database | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateServiceAutoscalingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoscalingPolicyRequest** | [**AutoscalingPolicyRequest**](AutoscalingPolicyRequest.md) |  | 

### Return type

[**AutoscalingPolicyResponse**](AutoscalingPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAutoscaling

> DeleteServiceAutoscaling(ctx, serviceId).Execute()

deleteServiceAutoscaling

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
	r, err := apiClient.AutoscalingPolicyAPI.DeleteServiceAutoscaling(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoscalingPolicyAPI.DeleteServiceAutoscaling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID of an application/job/container/database | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAutoscalingRequest struct via the builder pattern


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


## GetServiceAutoscaling

> AutoscalingPolicyResponse GetServiceAutoscaling(ctx, serviceId).Execute()

getServiceAutoscaling

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
	resp, r, err := apiClient.AutoscalingPolicyAPI.GetServiceAutoscaling(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoscalingPolicyAPI.GetServiceAutoscaling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceAutoscaling`: AutoscalingPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoscalingPolicyAPI.GetServiceAutoscaling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service ID of an application/job/container/database | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAutoscalingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoscalingPolicyResponse**](AutoscalingPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

