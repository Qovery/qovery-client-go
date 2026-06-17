# \BlueprintMainCallsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckBlueprintUpdate**](BlueprintMainCallsAPI.md#CheckBlueprintUpdate) | **Get** /blueprint/{blueprintId}/update | Check if a blueprint service has an available update
[**CreateBlueprint**](BlueprintMainCallsAPI.md#CreateBlueprint) | **Post** /environment/{environmentId}/blueprint | Create a blueprint service in an environment
[**GetBlueprintCatalog**](BlueprintMainCallsAPI.md#GetBlueprintCatalog) | **Get** /organization/{organizationId}/blueprint/catalog | Get the blueprint service catalog
[**PreviewBlueprintUpdate**](BlueprintMainCallsAPI.md#PreviewBlueprintUpdate) | **Post** /blueprint/{blueprintId}/update/preview | Preview a blueprint update
[**UpdateBlueprint**](BlueprintMainCallsAPI.md#UpdateBlueprint) | **Patch** /blueprint/{blueprintId} | Update a blueprint service



## CheckBlueprintUpdate

> BlueprintUpdateResponse CheckBlueprintUpdate(ctx, blueprintId).Execute()

Check if a blueprint service has an available update



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
	blueprintId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Blueprint ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintMainCallsAPI.CheckBlueprintUpdate(context.Background(), blueprintId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintMainCallsAPI.CheckBlueprintUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckBlueprintUpdate`: BlueprintUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `BlueprintMainCallsAPI.CheckBlueprintUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** | Blueprint ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckBlueprintUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlueprintUpdateResponse**](BlueprintUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBlueprint

> BlueprintResponse CreateBlueprint(ctx, environmentId).BlueprintCreateRequest(blueprintCreateRequest).Deploy(deploy).Execute()

Create a blueprint service in an environment



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
	blueprintCreateRequest := *openapiclient.NewBlueprintCreateRequest("my-postgres", "aws/postgres/17/1.0.1", "https://cdn.qovery.com/icons/postgresql.svg") // BlueprintCreateRequest | 
	deploy := true // bool | Trigger a deployment immediately after creation (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintMainCallsAPI.CreateBlueprint(context.Background(), environmentId).BlueprintCreateRequest(blueprintCreateRequest).Deploy(deploy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintMainCallsAPI.CreateBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlueprint`: BlueprintResponse
	fmt.Fprintf(os.Stdout, "Response from `BlueprintMainCallsAPI.CreateBlueprint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlueprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blueprintCreateRequest** | [**BlueprintCreateRequest**](BlueprintCreateRequest.md) |  | 
 **deploy** | **bool** | Trigger a deployment immediately after creation | [default to false]

### Return type

[**BlueprintResponse**](BlueprintResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlueprintCatalog

> BlueprintCatalogResponse GetBlueprintCatalog(ctx, organizationId).Execute()

Get the blueprint service catalog



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
	resp, r, err := apiClient.BlueprintMainCallsAPI.GetBlueprintCatalog(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintMainCallsAPI.GetBlueprintCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlueprintCatalog`: BlueprintCatalogResponse
	fmt.Fprintf(os.Stdout, "Response from `BlueprintMainCallsAPI.GetBlueprintCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlueprintCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlueprintCatalogResponse**](BlueprintCatalogResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewBlueprintUpdate

> BlueprintUpdatePreviewResponse PreviewBlueprintUpdate(ctx, blueprintId).BlueprintUpdatePreviewRequest(blueprintUpdatePreviewRequest).Execute()

Preview a blueprint update



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
	blueprintId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Blueprint ID
	blueprintUpdatePreviewRequest := *openapiclient.NewBlueprintUpdatePreviewRequest("my-postgres", "aws/postgres/17/1.1.0", "https://cdn.qovery.com/icons/postgresql.svg") // BlueprintUpdatePreviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintMainCallsAPI.PreviewBlueprintUpdate(context.Background(), blueprintId).BlueprintUpdatePreviewRequest(blueprintUpdatePreviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintMainCallsAPI.PreviewBlueprintUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewBlueprintUpdate`: BlueprintUpdatePreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `BlueprintMainCallsAPI.PreviewBlueprintUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** | Blueprint ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewBlueprintUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blueprintUpdatePreviewRequest** | [**BlueprintUpdatePreviewRequest**](BlueprintUpdatePreviewRequest.md) |  | 

### Return type

[**BlueprintUpdatePreviewResponse**](BlueprintUpdatePreviewResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlueprint

> BlueprintResponse UpdateBlueprint(ctx, blueprintId).BlueprintUpdateRequest(blueprintUpdateRequest).Execute()

Update a blueprint service



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
	blueprintId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Blueprint ID
	blueprintUpdateRequest := *openapiclient.NewBlueprintUpdateRequest("my-postgres", "aws/postgres/17/1.1.0", "https://cdn.qovery.com/icons/postgresql.svg") // BlueprintUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintMainCallsAPI.UpdateBlueprint(context.Background(), blueprintId).BlueprintUpdateRequest(blueprintUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintMainCallsAPI.UpdateBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBlueprint`: BlueprintResponse
	fmt.Fprintf(os.Stdout, "Response from `BlueprintMainCallsAPI.UpdateBlueprint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** | Blueprint ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlueprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blueprintUpdateRequest** | [**BlueprintUpdateRequest**](BlueprintUpdateRequest.md) |  | 

### Return type

[**BlueprintResponse**](BlueprintResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

