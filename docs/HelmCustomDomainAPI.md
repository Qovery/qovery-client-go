# \HelmCustomDomainAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckHelmCustomDomain**](HelmCustomDomainAPI.md#CheckHelmCustomDomain) | **Get** /helm/{helmId}/checkCustomDomain | Check Helm Custom Domain
[**CreateHelmCustomDomain**](HelmCustomDomainAPI.md#CreateHelmCustomDomain) | **Post** /helm/{helmId}/customDomain | Add custom domain to the helm.
[**DeleteHelmCustomDomain**](HelmCustomDomainAPI.md#DeleteHelmCustomDomain) | **Delete** /helm/{helmId}/customDomain/{customDomainId} | Delete a Custom Domain
[**EditHelmCustomDomain**](HelmCustomDomainAPI.md#EditHelmCustomDomain) | **Put** /helm/{helmId}/customDomain/{customDomainId} | Edit a Custom Domain
[**GetHelmCustomDomain**](HelmCustomDomainAPI.md#GetHelmCustomDomain) | **Get** /helm/{helmId}/customDomain/{customDomainId} | Get a Custom Domain
[**ListHelmCustomDomain**](HelmCustomDomainAPI.md#ListHelmCustomDomain) | **Get** /helm/{helmId}/customDomain | List helm custom domains



## CheckHelmCustomDomain

> CheckedCustomDomainsResponse CheckHelmCustomDomain(ctx, helmId).Execute()

Check Helm Custom Domain

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
	resp, r, err := apiClient.HelmCustomDomainAPI.CheckHelmCustomDomain(context.Background(), helmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmCustomDomainAPI.CheckHelmCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckHelmCustomDomain`: CheckedCustomDomainsResponse
	fmt.Fprintf(os.Stdout, "Response from `HelmCustomDomainAPI.CheckHelmCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckHelmCustomDomainRequest struct via the builder pattern


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


## CreateHelmCustomDomain

> CustomDomain CreateHelmCustomDomain(ctx, helmId).CustomDomainRequest(customDomainRequest).Execute()

Add custom domain to the helm.



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
	customDomainRequest := *openapiclient.NewCustomDomainRequest("my.domain.tld", false) // CustomDomainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmCustomDomainAPI.CreateHelmCustomDomain(context.Background(), helmId).CustomDomainRequest(customDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmCustomDomainAPI.CreateHelmCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHelmCustomDomain`: CustomDomain
	fmt.Fprintf(os.Stdout, "Response from `HelmCustomDomainAPI.CreateHelmCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHelmCustomDomainRequest struct via the builder pattern


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


## DeleteHelmCustomDomain

> DeleteHelmCustomDomain(ctx, helmId, customDomainId).Execute()

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
	helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID
	customDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Custom Domain ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HelmCustomDomainAPI.DeleteHelmCustomDomain(context.Background(), helmId, customDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmCustomDomainAPI.DeleteHelmCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 
**customDomainId** | **string** | Custom Domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHelmCustomDomainRequest struct via the builder pattern


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


## EditHelmCustomDomain

> CustomDomain EditHelmCustomDomain(ctx, helmId, customDomainId).CustomDomainRequest(customDomainRequest).Execute()

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
	helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID
	customDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Custom Domain ID
	customDomainRequest := *openapiclient.NewCustomDomainRequest("my.domain.tld", false) // CustomDomainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmCustomDomainAPI.EditHelmCustomDomain(context.Background(), helmId, customDomainId).CustomDomainRequest(customDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmCustomDomainAPI.EditHelmCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditHelmCustomDomain`: CustomDomain
	fmt.Fprintf(os.Stdout, "Response from `HelmCustomDomainAPI.EditHelmCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 
**customDomainId** | **string** | Custom Domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditHelmCustomDomainRequest struct via the builder pattern


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


## GetHelmCustomDomain

> CustomDomain GetHelmCustomDomain(ctx, helmId, customDomainId).Execute()

Get a Custom Domain



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
	customDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Custom Domain ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmCustomDomainAPI.GetHelmCustomDomain(context.Background(), helmId, customDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmCustomDomainAPI.GetHelmCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHelmCustomDomain`: CustomDomain
	fmt.Fprintf(os.Stdout, "Response from `HelmCustomDomainAPI.GetHelmCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 
**customDomainId** | **string** | Custom Domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHelmCustomDomainRequest struct via the builder pattern


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


## ListHelmCustomDomain

> CustomDomainResponseList ListHelmCustomDomain(ctx, helmId).Execute()

List helm custom domains



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
	resp, r, err := apiClient.HelmCustomDomainAPI.ListHelmCustomDomain(context.Background(), helmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmCustomDomainAPI.ListHelmCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHelmCustomDomain`: CustomDomainResponseList
	fmt.Fprintf(os.Stdout, "Response from `HelmCustomDomainAPI.ListHelmCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHelmCustomDomainRequest struct via the builder pattern


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

