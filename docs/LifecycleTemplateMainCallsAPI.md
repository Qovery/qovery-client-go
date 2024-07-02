# \LifecycleTemplateMainCallsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnvironmentLifecycleTemplate**](LifecycleTemplateMainCallsAPI.md#GetEnvironmentLifecycleTemplate) | **Get** /environment/{environmentId}/lifecycleTemplate/{lifecycleTemplateId} | Get specific lifecycle template
[**ListEnvironmentLifecycleTemplates**](LifecycleTemplateMainCallsAPI.md#ListEnvironmentLifecycleTemplates) | **Get** /environment/{environmentId}/lifecycleTemplate | List available lifecycle template for this environment



## GetEnvironmentLifecycleTemplate

> LifecycleTemplateResponse GetEnvironmentLifecycleTemplate(ctx, environmentId, lifecycleTemplateId).Execute()

Get specific lifecycle template

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
	environmentId := "environmentId_example" // string | 
	lifecycleTemplateId := "lifecycleTemplateId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LifecycleTemplateMainCallsAPI.GetEnvironmentLifecycleTemplate(context.Background(), environmentId, lifecycleTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleTemplateMainCallsAPI.GetEnvironmentLifecycleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentLifecycleTemplate`: LifecycleTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `LifecycleTemplateMainCallsAPI.GetEnvironmentLifecycleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 
**lifecycleTemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentLifecycleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LifecycleTemplateResponse**](LifecycleTemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironmentLifecycleTemplates

> LifecycleTemplateListResponse ListEnvironmentLifecycleTemplates(ctx, environmentId).Execute()

List available lifecycle template for this environment

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
	environmentId := "environmentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LifecycleTemplateMainCallsAPI.ListEnvironmentLifecycleTemplates(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleTemplateMainCallsAPI.ListEnvironmentLifecycleTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnvironmentLifecycleTemplates`: LifecycleTemplateListResponse
	fmt.Fprintf(os.Stdout, "Response from `LifecycleTemplateMainCallsAPI.ListEnvironmentLifecycleTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentLifecycleTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LifecycleTemplateListResponse**](LifecycleTemplateListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

