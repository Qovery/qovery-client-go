# \ContainerCustomDomainAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckContainerCustomDomain**](ContainerCustomDomainAPI.md#CheckContainerCustomDomain) | **Get** /container/{containerId}/checkCustomDomain | Check Container Custom Domain
[**CreateContainerCustomDomain**](ContainerCustomDomainAPI.md#CreateContainerCustomDomain) | **Post** /container/{containerId}/customDomain | Add custom domain to the container.
[**DeleteContainerCustomDomain**](ContainerCustomDomainAPI.md#DeleteContainerCustomDomain) | **Delete** /container/{containerId}/customDomain/{customDomainId} | Delete a Custom Domain
[**EditContainerCustomDomain**](ContainerCustomDomainAPI.md#EditContainerCustomDomain) | **Put** /container/{containerId}/customDomain/{customDomainId} | Edit a Custom Domain
[**GetContainerCustomDomainStatus**](ContainerCustomDomainAPI.md#GetContainerCustomDomainStatus) | **Get** /container/{containerId}/customDomain/{customDomainId}/status | Get Custom Domain status
[**ListContainerCustomDomain**](ContainerCustomDomainAPI.md#ListContainerCustomDomain) | **Get** /container/{containerId}/customDomain | List container custom domains



## CheckContainerCustomDomain

> CheckedCustomDomainsResponse CheckContainerCustomDomain(ctx, containerId).Execute()

Check Container Custom Domain

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
	resp, r, err := apiClient.ContainerCustomDomainAPI.CheckContainerCustomDomain(context.Background(), containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerCustomDomainAPI.CheckContainerCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckContainerCustomDomain`: CheckedCustomDomainsResponse
	fmt.Fprintf(os.Stdout, "Response from `ContainerCustomDomainAPI.CheckContainerCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckContainerCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckedCustomDomainsResponse**](CheckedCustomDomainsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContainerCustomDomain

> CustomDomain CreateContainerCustomDomain(ctx, containerId).CustomDomainRequest(customDomainRequest).Execute()

Add custom domain to the container.



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
	customDomainRequest := *openapiclient.NewCustomDomainRequest("my.domain.tld", false) // CustomDomainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerCustomDomainAPI.CreateContainerCustomDomain(context.Background(), containerId).CustomDomainRequest(customDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerCustomDomainAPI.CreateContainerCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContainerCustomDomain`: CustomDomain
	fmt.Fprintf(os.Stdout, "Response from `ContainerCustomDomainAPI.CreateContainerCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customDomainRequest** | [**CustomDomainRequest**](CustomDomainRequest.md) |  | 

### Return type

[**CustomDomain**](CustomDomain.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerCustomDomain

> DeleteContainerCustomDomain(ctx, containerId, customDomainId).Execute()

Delete a Custom Domain



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
	customDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Custom Domain ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerCustomDomainAPI.DeleteContainerCustomDomain(context.Background(), containerId, customDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerCustomDomainAPI.DeleteContainerCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**customDomainId** | **string** | Custom Domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerCustomDomainRequest struct via the builder pattern


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


## EditContainerCustomDomain

> CustomDomain EditContainerCustomDomain(ctx, containerId, customDomainId).CustomDomainRequest(customDomainRequest).Execute()

Edit a Custom Domain



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
	customDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Custom Domain ID
	customDomainRequest := *openapiclient.NewCustomDomainRequest("my.domain.tld", false) // CustomDomainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerCustomDomainAPI.EditContainerCustomDomain(context.Background(), containerId, customDomainId).CustomDomainRequest(customDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerCustomDomainAPI.EditContainerCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditContainerCustomDomain`: CustomDomain
	fmt.Fprintf(os.Stdout, "Response from `ContainerCustomDomainAPI.EditContainerCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**customDomainId** | **string** | Custom Domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditContainerCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customDomainRequest** | [**CustomDomainRequest**](CustomDomainRequest.md) |  | 

### Return type

[**CustomDomain**](CustomDomain.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerCustomDomainStatus

> CustomDomain GetContainerCustomDomainStatus(ctx, containerId, customDomainId).Execute()

Get Custom Domain status

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
	customDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Custom Domain ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerCustomDomainAPI.GetContainerCustomDomainStatus(context.Background(), containerId, customDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerCustomDomainAPI.GetContainerCustomDomainStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerCustomDomainStatus`: CustomDomain
	fmt.Fprintf(os.Stdout, "Response from `ContainerCustomDomainAPI.GetContainerCustomDomainStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 
**customDomainId** | **string** | Custom Domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerCustomDomainStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomDomain**](CustomDomain.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerCustomDomain

> CustomDomainResponseList ListContainerCustomDomain(ctx, containerId).Execute()

List container custom domains

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
	resp, r, err := apiClient.ContainerCustomDomainAPI.ListContainerCustomDomain(context.Background(), containerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerCustomDomainAPI.ListContainerCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContainerCustomDomain`: CustomDomainResponseList
	fmt.Fprintf(os.Stdout, "Response from `ContainerCustomDomainAPI.ListContainerCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomDomainResponseList**](CustomDomainResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

