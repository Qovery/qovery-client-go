# \BlueprintCatalogAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlueprintCatalogServiceManifest**](BlueprintCatalogAPI.md#GetBlueprintCatalogServiceManifest) | **Get** /organization/{organizationId}/blueprint/catalog/{provider}/{serviceFamily}/{serviceVersion}/manifest | Get the input fields to display for a blueprint catalog service
[**GetBlueprintCatalogServiceReadme**](BlueprintCatalogAPI.md#GetBlueprintCatalogServiceReadme) | **Get** /organization/{organizationId}/blueprint/catalog/{provider}/{serviceFamily}/{serviceVersion}/readme | Get the README of a blueprint catalog service



## GetBlueprintCatalogServiceManifest

> GetBlueprintCatalogServiceManifest200Response GetBlueprintCatalogServiceManifest(ctx, organizationId, provider, serviceFamily, serviceVersion).Execute()

Get the input fields to display for a blueprint catalog service



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
	provider := "provider_example" // string | Cloud provider (e.g. aws, gcp, azure)
	serviceFamily := "serviceFamily_example" // string | Service family (e.g. mysql, postgresql)
	serviceVersion := "serviceVersion_example" // string | Service version (e.g. 8, 14)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintCatalogAPI.GetBlueprintCatalogServiceManifest(context.Background(), organizationId, provider, serviceFamily, serviceVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintCatalogAPI.GetBlueprintCatalogServiceManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlueprintCatalogServiceManifest`: GetBlueprintCatalogServiceManifest200Response
	fmt.Fprintf(os.Stdout, "Response from `BlueprintCatalogAPI.GetBlueprintCatalogServiceManifest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**provider** | **string** | Cloud provider (e.g. aws, gcp, azure) | 
**serviceFamily** | **string** | Service family (e.g. mysql, postgresql) | 
**serviceVersion** | **string** | Service version (e.g. 8, 14) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlueprintCatalogServiceManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetBlueprintCatalogServiceManifest200Response**](GetBlueprintCatalogServiceManifest200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlueprintCatalogServiceReadme

> string GetBlueprintCatalogServiceReadme(ctx, organizationId, provider, serviceFamily, serviceVersion).Execute()

Get the README of a blueprint catalog service

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
	provider := "provider_example" // string | Cloud provider (e.g. aws, gcp, azure)
	serviceFamily := "serviceFamily_example" // string | Service family (e.g. mysql, postgresql)
	serviceVersion := "serviceVersion_example" // string | Service version (e.g. 8, 14)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintCatalogAPI.GetBlueprintCatalogServiceReadme(context.Background(), organizationId, provider, serviceFamily, serviceVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintCatalogAPI.GetBlueprintCatalogServiceReadme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlueprintCatalogServiceReadme`: string
	fmt.Fprintf(os.Stdout, "Response from `BlueprintCatalogAPI.GetBlueprintCatalogServiceReadme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**provider** | **string** | Cloud provider (e.g. aws, gcp, azure) | 
**serviceFamily** | **string** | Service family (e.g. mysql, postgresql) | 
**serviceVersion** | **string** | Service version (e.g. 8, 14) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlueprintCatalogServiceReadmeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/markdown, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

