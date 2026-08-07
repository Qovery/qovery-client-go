# \OrganizationPolicyApiTokenAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationPolicyApiToken**](OrganizationPolicyApiTokenAPI.md#CreateOrganizationPolicyApiToken) | **Post** /organization/{organizationId}/policyApiToken | Create an organization policy api token
[**DeleteOrganizationPolicyApiToken**](OrganizationPolicyApiTokenAPI.md#DeleteOrganizationPolicyApiToken) | **Delete** /organization/{organizationId}/policyApiToken/{policyApiTokenId} | Delete organization policy api token
[**ListOrganizationPolicyApiTokens**](OrganizationPolicyApiTokenAPI.md#ListOrganizationPolicyApiTokens) | **Get** /organization/{organizationId}/policyApiToken | List organization policy api tokens



## CreateOrganizationPolicyApiToken

> OrganizationPolicyApiTokenCreate CreateOrganizationPolicyApiToken(ctx, organizationId).OrganizationPolicyApiTokenCreateRequest(organizationPolicyApiTokenCreateRequest).Execute()

Create an organization policy api token



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
	organizationPolicyApiTokenCreateRequest := *openapiclient.NewOrganizationPolicyApiTokenCreateRequest("Name_example", "OpaPolicy_example") // OrganizationPolicyApiTokenCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationPolicyApiTokenAPI.CreateOrganizationPolicyApiToken(context.Background(), organizationId).OrganizationPolicyApiTokenCreateRequest(organizationPolicyApiTokenCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationPolicyApiTokenAPI.CreateOrganizationPolicyApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationPolicyApiToken`: OrganizationPolicyApiTokenCreate
	fmt.Fprintf(os.Stdout, "Response from `OrganizationPolicyApiTokenAPI.CreateOrganizationPolicyApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPolicyApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationPolicyApiTokenCreateRequest** | [**OrganizationPolicyApiTokenCreateRequest**](OrganizationPolicyApiTokenCreateRequest.md) |  | 

### Return type

[**OrganizationPolicyApiTokenCreate**](OrganizationPolicyApiTokenCreate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationPolicyApiToken

> DeleteOrganizationPolicyApiToken(ctx, organizationId, policyApiTokenId).Execute()

Delete organization policy api token



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
	policyApiTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization Policy Api Token ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationPolicyApiTokenAPI.DeleteOrganizationPolicyApiToken(context.Background(), organizationId, policyApiTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationPolicyApiTokenAPI.DeleteOrganizationPolicyApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyApiTokenId** | **string** | Organization Policy Api Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPolicyApiTokenRequest struct via the builder pattern


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


## ListOrganizationPolicyApiTokens

> OrganizationPolicyApiTokenResponseList ListOrganizationPolicyApiTokens(ctx, organizationId).Execute()

List organization policy api tokens



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
	resp, r, err := apiClient.OrganizationPolicyApiTokenAPI.ListOrganizationPolicyApiTokens(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationPolicyApiTokenAPI.ListOrganizationPolicyApiTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationPolicyApiTokens`: OrganizationPolicyApiTokenResponseList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationPolicyApiTokenAPI.ListOrganizationPolicyApiTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationPolicyApiTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationPolicyApiTokenResponseList**](OrganizationPolicyApiTokenResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

