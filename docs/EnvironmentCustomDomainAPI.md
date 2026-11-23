# \EnvironmentCustomDomainAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironmentCustomDomain**](EnvironmentCustomDomainAPI.md#CreateEnvironmentCustomDomain) | **Post** /environment/{environmentId}/customDomain | Create a custom domain for several services
[**DeleteEnvironmentCustomDomain**](EnvironmentCustomDomainAPI.md#DeleteEnvironmentCustomDomain) | **Delete** /environment/{environmentId}/customDomain/{domainName} | Delete a custom domain of several services
[**EditEnvironmentCustomDomain**](EnvironmentCustomDomainAPI.md#EditEnvironmentCustomDomain) | **Put** /environment/{environmentId}/customDomain/{domainName} | Edit a custom domain of several services
[**GetEnvironmentCustomDomainCnameTarget**](EnvironmentCustomDomainAPI.md#GetEnvironmentCustomDomainCnameTarget) | **Get** /environment/{environmentId}/customDomain/cnameTarget | Get the CNAME target of the environment
[**ListEnvironmentCustomDomain**](EnvironmentCustomDomainAPI.md#ListEnvironmentCustomDomain) | **Get** /environment/{environmentId}/customDomain | List environment custom domains
[**ValidateEnvironmentCustomDomain**](EnvironmentCustomDomainAPI.md#ValidateEnvironmentCustomDomain) | **Post** /environment/{environmentId}/customDomain/validate | Validate a custom domain for several services



## CreateEnvironmentCustomDomain

> EnvironmentCustomDomainResponse CreateEnvironmentCustomDomain(ctx, environmentId).EnvironmentCustomDomainRequest(environmentCustomDomainRequest).Execute()

Create a custom domain for several services



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
	environmentCustomDomainRequest := *openapiclient.NewEnvironmentCustomDomainRequest("my.domain.tld", []openapiclient.CustomDomainServiceRequest{*openapiclient.NewCustomDomainServiceRequest("Id_example", openapiclient.ServiceTypeEnum("APPLICATION"))}) // EnvironmentCustomDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentCustomDomainAPI.CreateEnvironmentCustomDomain(context.Background(), environmentId).EnvironmentCustomDomainRequest(environmentCustomDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentCustomDomainAPI.CreateEnvironmentCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnvironmentCustomDomain`: EnvironmentCustomDomainResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentCustomDomainAPI.CreateEnvironmentCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentCustomDomainRequest** | [**EnvironmentCustomDomainRequest**](EnvironmentCustomDomainRequest.md) |  | 

### Return type

[**EnvironmentCustomDomainResponse**](EnvironmentCustomDomainResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentCustomDomain

> CustomDomainDeletionResponse DeleteEnvironmentCustomDomain(ctx, environmentId, domainName).Execute()

Delete a custom domain of several services



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
	domainName := "my.domain.tld" // string | The current name of the custom domain.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentCustomDomainAPI.DeleteEnvironmentCustomDomain(context.Background(), environmentId, domainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentCustomDomainAPI.DeleteEnvironmentCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEnvironmentCustomDomain`: CustomDomainDeletionResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentCustomDomainAPI.DeleteEnvironmentCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 
**domainName** | **string** | The current name of the custom domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomDomainDeletionResponse**](CustomDomainDeletionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditEnvironmentCustomDomain

> EnvironmentCustomDomainResponse EditEnvironmentCustomDomain(ctx, environmentId, domainName).EnvironmentCustomDomainRequest(environmentCustomDomainRequest).Execute()

Edit a custom domain of several services



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
	domainName := "my.domain.tld" // string | The current name of the custom domain.
	environmentCustomDomainRequest := *openapiclient.NewEnvironmentCustomDomainRequest("my.domain.tld", []openapiclient.CustomDomainServiceRequest{*openapiclient.NewCustomDomainServiceRequest("Id_example", openapiclient.ServiceTypeEnum("APPLICATION"))}) // EnvironmentCustomDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentCustomDomainAPI.EditEnvironmentCustomDomain(context.Background(), environmentId, domainName).EnvironmentCustomDomainRequest(environmentCustomDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentCustomDomainAPI.EditEnvironmentCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditEnvironmentCustomDomain`: EnvironmentCustomDomainResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentCustomDomainAPI.EditEnvironmentCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 
**domainName** | **string** | The current name of the custom domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditEnvironmentCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environmentCustomDomainRequest** | [**EnvironmentCustomDomainRequest**](EnvironmentCustomDomainRequest.md) |  | 

### Return type

[**EnvironmentCustomDomainResponse**](EnvironmentCustomDomainResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentCustomDomainCnameTarget

> EnvironmentCnameTargetResponse GetEnvironmentCustomDomainCnameTarget(ctx, environmentId).Execute()

Get the CNAME target of the environment



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
	resp, r, err := apiClient.EnvironmentCustomDomainAPI.GetEnvironmentCustomDomainCnameTarget(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentCustomDomainAPI.GetEnvironmentCustomDomainCnameTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentCustomDomainCnameTarget`: EnvironmentCnameTargetResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentCustomDomainAPI.GetEnvironmentCustomDomainCnameTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentCustomDomainCnameTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentCnameTargetResponse**](EnvironmentCnameTargetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironmentCustomDomain

> OrganizationCustomDomainResponseList ListEnvironmentCustomDomain(ctx, environmentId).Execute()

List environment custom domains



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
	resp, r, err := apiClient.EnvironmentCustomDomainAPI.ListEnvironmentCustomDomain(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentCustomDomainAPI.ListEnvironmentCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnvironmentCustomDomain`: OrganizationCustomDomainResponseList
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentCustomDomainAPI.ListEnvironmentCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationCustomDomainResponseList**](OrganizationCustomDomainResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateEnvironmentCustomDomain

> CustomDomainErrorsResponse ValidateEnvironmentCustomDomain(ctx, environmentId).CustomDomainValidationRequest(customDomainValidationRequest).Execute()

Validate a custom domain for several services



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
	customDomainValidationRequest := *openapiclient.NewCustomDomainValidationRequest("my.domain.tld", []openapiclient.CustomDomainServiceRequest{*openapiclient.NewCustomDomainServiceRequest("Id_example", openapiclient.ServiceTypeEnum("APPLICATION"))}) // CustomDomainValidationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentCustomDomainAPI.ValidateEnvironmentCustomDomain(context.Background(), environmentId).CustomDomainValidationRequest(customDomainValidationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentCustomDomainAPI.ValidateEnvironmentCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateEnvironmentCustomDomain`: CustomDomainErrorsResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentCustomDomainAPI.ValidateEnvironmentCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateEnvironmentCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customDomainValidationRequest** | [**CustomDomainValidationRequest**](CustomDomainValidationRequest.md) |  | 

### Return type

[**CustomDomainErrorsResponse**](CustomDomainErrorsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

