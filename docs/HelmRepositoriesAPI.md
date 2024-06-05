# \HelmRepositoriesAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHelmRepository**](HelmRepositoriesAPI.md#CreateHelmRepository) | **Post** /organization/{organizationId}/helmRepository | Create a helm repository
[**DeleteHelmRepository**](HelmRepositoriesAPI.md#DeleteHelmRepository) | **Delete** /organization/{organizationId}/helmRepository/{helmRepositoryId} | Delete a helm repository
[**EditHelmRepository**](HelmRepositoriesAPI.md#EditHelmRepository) | **Put** /organization/{organizationId}/helmRepository/{helmRepositoryId} | Edit a helm repository
[**GetHelmCharts**](HelmRepositoriesAPI.md#GetHelmCharts) | **Get** /organization/{organizationId}/helmRepository/{helmRepositoryId}/charts | List helm charts contained inside the repository
[**GetHelmRepository**](HelmRepositoriesAPI.md#GetHelmRepository) | **Get** /organization/{organizationId}/helmRepository/{helmRepositoryId} | Get a helm repository
[**ListAvailableHelmRepository**](HelmRepositoriesAPI.md#ListAvailableHelmRepository) | **Get** /availableHelmRepository | List supported helm repository
[**ListHelmRepository**](HelmRepositoriesAPI.md#ListHelmRepository) | **Get** /organization/{organizationId}/helmRepository | List organization helm repositories



## CreateHelmRepository

> HelmRepositoryResponse CreateHelmRepository(ctx, organizationId).HelmRepositoryRequest(helmRepositoryRequest).Execute()

Create a helm repository

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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	helmRepositoryRequest := *openapiclient.NewHelmRepositoryRequest("Name_example", openapiclient.HelmRepositoryKindEnum("HTTPS"), false, *openapiclient.NewHelmRepositoryRequestConfig()) // HelmRepositoryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmRepositoriesAPI.CreateHelmRepository(context.Background(), organizationId).HelmRepositoryRequest(helmRepositoryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmRepositoriesAPI.CreateHelmRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHelmRepository`: HelmRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `HelmRepositoriesAPI.CreateHelmRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHelmRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **helmRepositoryRequest** | [**HelmRepositoryRequest**](HelmRepositoryRequest.md) |  | 

### Return type

[**HelmRepositoryResponse**](HelmRepositoryResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHelmRepository

> DeleteHelmRepository(ctx, organizationId, helmRepositoryId).Execute()

Delete a helm repository

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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	helmRepositoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm chart repository ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HelmRepositoriesAPI.DeleteHelmRepository(context.Background(), organizationId, helmRepositoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmRepositoriesAPI.DeleteHelmRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**helmRepositoryId** | **string** | Helm chart repository ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHelmRepositoryRequest struct via the builder pattern


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


## EditHelmRepository

> HelmRepositoryResponse EditHelmRepository(ctx, organizationId, helmRepositoryId).HelmRepositoryRequest(helmRepositoryRequest).Execute()

Edit a helm repository

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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	helmRepositoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm chart repository ID
	helmRepositoryRequest := *openapiclient.NewHelmRepositoryRequest("Name_example", openapiclient.HelmRepositoryKindEnum("HTTPS"), false, *openapiclient.NewHelmRepositoryRequestConfig()) // HelmRepositoryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmRepositoriesAPI.EditHelmRepository(context.Background(), organizationId, helmRepositoryId).HelmRepositoryRequest(helmRepositoryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmRepositoriesAPI.EditHelmRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditHelmRepository`: HelmRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `HelmRepositoriesAPI.EditHelmRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**helmRepositoryId** | **string** | Helm chart repository ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditHelmRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **helmRepositoryRequest** | [**HelmRepositoryRequest**](HelmRepositoryRequest.md) |  | 

### Return type

[**HelmRepositoryResponse**](HelmRepositoryResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHelmCharts

> HelmVersionResponseList GetHelmCharts(ctx, organizationId, helmRepositoryId).ChartName(chartName).Execute()

List helm charts contained inside the repository

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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	helmRepositoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm chart repository ID
	chartName := "chartName_example" // string | Helm chart name to filter the result on (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmRepositoriesAPI.GetHelmCharts(context.Background(), organizationId, helmRepositoryId).ChartName(chartName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmRepositoriesAPI.GetHelmCharts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHelmCharts`: HelmVersionResponseList
	fmt.Fprintf(os.Stdout, "Response from `HelmRepositoriesAPI.GetHelmCharts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**helmRepositoryId** | **string** | Helm chart repository ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHelmChartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **chartName** | **string** | Helm chart name to filter the result on | 

### Return type

[**HelmVersionResponseList**](HelmVersionResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHelmRepository

> HelmRepositoryResponse GetHelmRepository(ctx, organizationId, helmRepositoryId).Execute()

Get a helm repository

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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	helmRepositoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm chart repository ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmRepositoriesAPI.GetHelmRepository(context.Background(), organizationId, helmRepositoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmRepositoriesAPI.GetHelmRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHelmRepository`: HelmRepositoryResponse
	fmt.Fprintf(os.Stdout, "Response from `HelmRepositoriesAPI.GetHelmRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**helmRepositoryId** | **string** | Helm chart repository ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHelmRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HelmRepositoryResponse**](HelmRepositoryResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableHelmRepository

> AvailableHelmRepositoryResponseList ListAvailableHelmRepository(ctx).Execute()

List supported helm repository



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmRepositoriesAPI.ListAvailableHelmRepository(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmRepositoriesAPI.ListAvailableHelmRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAvailableHelmRepository`: AvailableHelmRepositoryResponseList
	fmt.Fprintf(os.Stdout, "Response from `HelmRepositoriesAPI.ListAvailableHelmRepository`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableHelmRepositoryRequest struct via the builder pattern


### Return type

[**AvailableHelmRepositoryResponseList**](AvailableHelmRepositoryResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHelmRepository

> HelmRepositoryResponseList ListHelmRepository(ctx, organizationId).Execute()

List organization helm repositories

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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmRepositoriesAPI.ListHelmRepository(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmRepositoriesAPI.ListHelmRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHelmRepository`: HelmRepositoryResponseList
	fmt.Fprintf(os.Stdout, "Response from `HelmRepositoriesAPI.ListHelmRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHelmRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HelmRepositoryResponseList**](HelmRepositoryResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

