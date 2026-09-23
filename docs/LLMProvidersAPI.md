# \LLMProvidersAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLlmProvider**](LLMProvidersAPI.md#CreateLlmProvider) | **Post** /organization/{organizationId}/llmProvider | Create an LLM provider
[**DeleteLlmProvider**](LLMProvidersAPI.md#DeleteLlmProvider) | **Delete** /llmProvider/{llmProviderId} | Delete an LLM provider
[**EditLlmProvider**](LLMProvidersAPI.md#EditLlmProvider) | **Put** /llmProvider/{llmProviderId} | Edit an LLM provider
[**GetLlmProvider**](LLMProvidersAPI.md#GetLlmProvider) | **Get** /llmProvider/{llmProviderId} | Get an LLM provider
[**ListLlmProviderModels**](LLMProvidersAPI.md#ListLlmProviderModels) | **Get** /llmProvider/{llmProviderId}/models | List the models of an LLM provider
[**ListLlmProviders**](LLMProvidersAPI.md#ListLlmProviders) | **Get** /organization/{organizationId}/llmProvider | List organization LLM providers



## CreateLlmProvider

> LlmProviderResponse CreateLlmProvider(ctx, organizationId).LlmProviderRequest(llmProviderRequest).Execute()

Create an LLM provider



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
	llmProviderRequest := *openapiclient.NewLlmProviderRequest("Name_example", openapiclient.LlmProviderType("CLAUDE")) // LlmProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LLMProvidersAPI.CreateLlmProvider(context.Background(), organizationId).LlmProviderRequest(llmProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMProvidersAPI.CreateLlmProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLlmProvider`: LlmProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `LLMProvidersAPI.CreateLlmProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLlmProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **llmProviderRequest** | [**LlmProviderRequest**](LlmProviderRequest.md) |  | 

### Return type

[**LlmProviderResponse**](LlmProviderResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLlmProvider

> DeleteLlmProvider(ctx, llmProviderId).Execute()

Delete an LLM provider



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
	llmProviderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | LLM Provider ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LLMProvidersAPI.DeleteLlmProvider(context.Background(), llmProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMProvidersAPI.DeleteLlmProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**llmProviderId** | **string** | LLM Provider ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLlmProviderRequest struct via the builder pattern


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


## EditLlmProvider

> LlmProviderResponse EditLlmProvider(ctx, llmProviderId).LlmProviderRequest(llmProviderRequest).Execute()

Edit an LLM provider



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
	llmProviderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | LLM Provider ID
	llmProviderRequest := *openapiclient.NewLlmProviderRequest("Name_example", openapiclient.LlmProviderType("CLAUDE")) // LlmProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LLMProvidersAPI.EditLlmProvider(context.Background(), llmProviderId).LlmProviderRequest(llmProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMProvidersAPI.EditLlmProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditLlmProvider`: LlmProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `LLMProvidersAPI.EditLlmProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**llmProviderId** | **string** | LLM Provider ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLlmProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **llmProviderRequest** | [**LlmProviderRequest**](LlmProviderRequest.md) |  | 

### Return type

[**LlmProviderResponse**](LlmProviderResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLlmProvider

> LlmProviderResponse GetLlmProvider(ctx, llmProviderId).Execute()

Get an LLM provider



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
	llmProviderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | LLM Provider ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LLMProvidersAPI.GetLlmProvider(context.Background(), llmProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMProvidersAPI.GetLlmProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLlmProvider`: LlmProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `LLMProvidersAPI.GetLlmProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**llmProviderId** | **string** | LLM Provider ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLlmProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LlmProviderResponse**](LlmProviderResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLlmProviderModels

> LlmProviderModelResponseList ListLlmProviderModels(ctx, llmProviderId).Execute()

List the models of an LLM provider



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
	llmProviderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | LLM Provider ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LLMProvidersAPI.ListLlmProviderModels(context.Background(), llmProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMProvidersAPI.ListLlmProviderModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLlmProviderModels`: LlmProviderModelResponseList
	fmt.Fprintf(os.Stdout, "Response from `LLMProvidersAPI.ListLlmProviderModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**llmProviderId** | **string** | LLM Provider ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLlmProviderModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LlmProviderModelResponseList**](LlmProviderModelResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLlmProviders

> LlmProviderResponseList ListLlmProviders(ctx, organizationId).Execute()

List organization LLM providers



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
	resp, r, err := apiClient.LLMProvidersAPI.ListLlmProviders(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMProvidersAPI.ListLlmProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLlmProviders`: LlmProviderResponseList
	fmt.Fprintf(os.Stdout, "Response from `LLMProvidersAPI.ListLlmProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLlmProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LlmProviderResponseList**](LlmProviderResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

