# \DefaultAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnvironmentLifecycleTemplate**](DefaultAPI.md#GetEnvironmentLifecycleTemplate) | **Get** /environment/{environmentId}/lifecycleTemplate/{lifecycleTemplateId} | Get specific lifecycle tempalte
[**GetEnvironmentLifecycleTemplates**](DefaultAPI.md#GetEnvironmentLifecycleTemplates) | **Get** /environment/{environmentId}/lifecycleTemplate | Your GET endpoint



## GetEnvironmentLifecycleTemplate

> LifecycleTemplateResponse GetEnvironmentLifecycleTemplate(ctx, environmentId, lifecycleTemplateId).Execute()

Get specific lifecycle tempalte

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
	resp, r, err := apiClient.DefaultAPI.GetEnvironmentLifecycleTemplate(context.Background(), environmentId, lifecycleTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEnvironmentLifecycleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentLifecycleTemplate`: LifecycleTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEnvironmentLifecycleTemplate`: %v\n", resp)
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


## GetEnvironmentLifecycleTemplates

> LifecycleTemplateListResponse GetEnvironmentLifecycleTemplates(ctx, environmentId).Execute()

Your GET endpoint

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
	resp, r, err := apiClient.DefaultAPI.GetEnvironmentLifecycleTemplates(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEnvironmentLifecycleTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentLifecycleTemplates`: LifecycleTemplateListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEnvironmentLifecycleTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentLifecycleTemplatesRequest struct via the builder pattern


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

