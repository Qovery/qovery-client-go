# \EnvironmentsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironment**](EnvironmentsAPI.md#CreateEnvironment) | **Post** /project/{projectId}/environment | Create an environment
[**GetProjectEnvironmentServiceNumber**](EnvironmentsAPI.md#GetProjectEnvironmentServiceNumber) | **Get** /project/{projectId}/environment/stats | List total number of services for each environment of the project
[**GetProjectEnvironmentsOverview**](EnvironmentsAPI.md#GetProjectEnvironmentsOverview) | **Get** /project/{projectId}/environmentOverview | List environments overview
[**GetProjectEnvironmentsStatus**](EnvironmentsAPI.md#GetProjectEnvironmentsStatus) | **Get** /project/{projectId}/environment/status | List environments statuses
[**ListEnvironment**](EnvironmentsAPI.md#ListEnvironment) | **Get** /project/{projectId}/environment | List environments



## CreateEnvironment

> Environment CreateEnvironment(ctx, projectId).CreateEnvironmentRequest(createEnvironmentRequest).Execute()

Create an environment

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
	createEnvironmentRequest := *openapiclient.NewCreateEnvironmentRequest("Name_example") // CreateEnvironmentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.CreateEnvironment(context.Background(), projectId).CreateEnvironmentRequest(createEnvironmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.CreateEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnvironment`: Environment
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.CreateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createEnvironmentRequest** | [**CreateEnvironmentRequest**](CreateEnvironmentRequest.md) |  | 

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


## GetProjectEnvironmentServiceNumber

> EnvironmentStatsResponseList GetProjectEnvironmentServiceNumber(ctx, projectId).Execute()

List total number of services for each environment of the project



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
	resp, r, err := apiClient.EnvironmentsAPI.GetProjectEnvironmentServiceNumber(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.GetProjectEnvironmentServiceNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectEnvironmentServiceNumber`: EnvironmentStatsResponseList
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.GetProjectEnvironmentServiceNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectEnvironmentServiceNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentStatsResponseList**](EnvironmentStatsResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectEnvironmentsOverview

> EnvironmentOverviewResponseList GetProjectEnvironmentsOverview(ctx, projectId).Execute()

List environments overview



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
	resp, r, err := apiClient.EnvironmentsAPI.GetProjectEnvironmentsOverview(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.GetProjectEnvironmentsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectEnvironmentsOverview`: EnvironmentOverviewResponseList
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.GetProjectEnvironmentsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectEnvironmentsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentOverviewResponseList**](EnvironmentOverviewResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectEnvironmentsStatus

> EnvironmentStatusList GetProjectEnvironmentsStatus(ctx, projectId).Execute()

List environments statuses



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
	resp, r, err := apiClient.EnvironmentsAPI.GetProjectEnvironmentsStatus(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.GetProjectEnvironmentsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectEnvironmentsStatus`: EnvironmentStatusList
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.GetProjectEnvironmentsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectEnvironmentsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentStatusList**](EnvironmentStatusList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironment

> EnvironmentResponseList ListEnvironment(ctx, projectId).Execute()

List environments

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
	resp, r, err := apiClient.EnvironmentsAPI.ListEnvironment(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.ListEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnvironment`: EnvironmentResponseList
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.ListEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentResponseList**](EnvironmentResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

