# \ContainerRegistriesAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContainerRegistry**](ContainerRegistriesAPI.md#CreateContainerRegistry) | **Post** /organization/{organizationId}/containerRegistry | Create a container registry
[**DeleteContainerRegistry**](ContainerRegistriesAPI.md#DeleteContainerRegistry) | **Delete** /organization/{organizationId}/containerRegistry/{containerRegistryId} | Delete a container registry
[**EditContainerRegistry**](ContainerRegistriesAPI.md#EditContainerRegistry) | **Put** /organization/{organizationId}/containerRegistry/{containerRegistryId} | Edit a container registry
[**GetContainerRegistry**](ContainerRegistriesAPI.md#GetContainerRegistry) | **Get** /organization/{organizationId}/containerRegistry/{containerRegistryId} | Get a container registry
[**ListAvailableContainerRegistry**](ContainerRegistriesAPI.md#ListAvailableContainerRegistry) | **Get** /availableContainerRegistry | List supported container registries
[**ListContainerRegistry**](ContainerRegistriesAPI.md#ListContainerRegistry) | **Get** /organization/{organizationId}/containerRegistry | List organization container registries



## CreateContainerRegistry

> ContainerRegistryResponse CreateContainerRegistry(ctx, organizationId).ContainerRegistryRequest(containerRegistryRequest).Execute()

Create a container registry

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
	containerRegistryRequest := *openapiclient.NewContainerRegistryRequest("Name_example", openapiclient.ContainerRegistryKindEnum("ECR"), *openapiclient.NewContainerRegistryRequestConfig()) // ContainerRegistryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerRegistriesAPI.CreateContainerRegistry(context.Background(), organizationId).ContainerRegistryRequest(containerRegistryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistriesAPI.CreateContainerRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContainerRegistry`: ContainerRegistryResponse
	fmt.Fprintf(os.Stdout, "Response from `ContainerRegistriesAPI.CreateContainerRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerRegistryRequest** | [**ContainerRegistryRequest**](ContainerRegistryRequest.md) |  | 

### Return type

[**ContainerRegistryResponse**](ContainerRegistryResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerRegistry

> DeleteContainerRegistry(ctx, organizationId, containerRegistryId).Execute()

Delete a container registry

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
	containerRegistryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container Registry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerRegistriesAPI.DeleteContainerRegistry(context.Background(), organizationId, containerRegistryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistriesAPI.DeleteContainerRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**containerRegistryId** | **string** | Container Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerRegistryRequest struct via the builder pattern


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


## EditContainerRegistry

> ContainerRegistryResponse EditContainerRegistry(ctx, organizationId, containerRegistryId).ContainerRegistryRequest(containerRegistryRequest).Execute()

Edit a container registry

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
	containerRegistryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container Registry ID
	containerRegistryRequest := *openapiclient.NewContainerRegistryRequest("Name_example", openapiclient.ContainerRegistryKindEnum("ECR"), *openapiclient.NewContainerRegistryRequestConfig()) // ContainerRegistryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerRegistriesAPI.EditContainerRegistry(context.Background(), organizationId, containerRegistryId).ContainerRegistryRequest(containerRegistryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistriesAPI.EditContainerRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditContainerRegistry`: ContainerRegistryResponse
	fmt.Fprintf(os.Stdout, "Response from `ContainerRegistriesAPI.EditContainerRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**containerRegistryId** | **string** | Container Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditContainerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **containerRegistryRequest** | [**ContainerRegistryRequest**](ContainerRegistryRequest.md) |  | 

### Return type

[**ContainerRegistryResponse**](ContainerRegistryResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerRegistry

> ContainerRegistryResponse GetContainerRegistry(ctx, organizationId, containerRegistryId).Execute()

Get a container registry

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
	containerRegistryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container Registry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerRegistriesAPI.GetContainerRegistry(context.Background(), organizationId, containerRegistryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistriesAPI.GetContainerRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerRegistry`: ContainerRegistryResponse
	fmt.Fprintf(os.Stdout, "Response from `ContainerRegistriesAPI.GetContainerRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**containerRegistryId** | **string** | Container Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContainerRegistryResponse**](ContainerRegistryResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableContainerRegistry

> AvailableContainerRegistryResponseList ListAvailableContainerRegistry(ctx).Execute()

List supported container registries



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
	resp, r, err := apiClient.ContainerRegistriesAPI.ListAvailableContainerRegistry(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistriesAPI.ListAvailableContainerRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAvailableContainerRegistry`: AvailableContainerRegistryResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContainerRegistriesAPI.ListAvailableContainerRegistry`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableContainerRegistryRequest struct via the builder pattern


### Return type

[**AvailableContainerRegistryResponseList**](AvailableContainerRegistryResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerRegistry

> ContainerRegistryResponseList ListContainerRegistry(ctx, organizationId).Execute()

List organization container registries

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
	resp, r, err := apiClient.ContainerRegistriesAPI.ListContainerRegistry(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerRegistriesAPI.ListContainerRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContainerRegistry`: ContainerRegistryResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContainerRegistriesAPI.ListContainerRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerRegistryResponseList**](ContainerRegistryResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

