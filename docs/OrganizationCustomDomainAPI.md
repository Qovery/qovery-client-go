# \OrganizationCustomDomainAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteOrganizationCustomDomain**](OrganizationCustomDomainAPI.md#BulkDeleteOrganizationCustomDomain) | **Post** /organization/{organizationId}/customDomain/bulkDelete | Delete custom domains of several environments
[**ListOrganizationCustomDomain**](OrganizationCustomDomainAPI.md#ListOrganizationCustomDomain) | **Get** /organization/{organizationId}/customDomain | List organization custom domains



## BulkDeleteOrganizationCustomDomain

> CustomDomainDeletionResponse BulkDeleteOrganizationCustomDomain(ctx, organizationId).CustomDomainBulkDeleteItemRequest(customDomainBulkDeleteItemRequest).Execute()

Delete custom domains of several environments



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
	customDomainBulkDeleteItemRequest := []openapiclient.CustomDomainBulkDeleteItemRequest{*openapiclient.NewCustomDomainBulkDeleteItemRequest("EnvironmentId_example", "my.domain.tld")} // []CustomDomainBulkDeleteItemRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationCustomDomainAPI.BulkDeleteOrganizationCustomDomain(context.Background(), organizationId).CustomDomainBulkDeleteItemRequest(customDomainBulkDeleteItemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationCustomDomainAPI.BulkDeleteOrganizationCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkDeleteOrganizationCustomDomain`: CustomDomainDeletionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationCustomDomainAPI.BulkDeleteOrganizationCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteOrganizationCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customDomainBulkDeleteItemRequest** | [**[]CustomDomainBulkDeleteItemRequest**](CustomDomainBulkDeleteItemRequest.md) |  | 

### Return type

[**CustomDomainDeletionResponse**](CustomDomainDeletionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationCustomDomain

> OrganizationCustomDomainPaginatedResponseList ListOrganizationCustomDomain(ctx, organizationId).ProjectId(projectId).EnvironmentId(environmentId).Search(search).Page(page).PageSize(pageSize).Execute()

List organization custom domains



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only list the domains of this project. A project of another organization gives an empty list. (optional)
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only list the domains of this environment. An environment of another organization gives an empty list. (optional)
	search := "search_example" // string | Only list the domains whose name contains this text, ignoring case. The text is trimmed and matched literally, with no wildcard. A blank value does not filter. (optional)
	page := int32(56) // int32 | The page to return, starting at 1. A page after the last one has no results. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of domains per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationCustomDomainAPI.ListOrganizationCustomDomain(context.Background(), organizationId).ProjectId(projectId).EnvironmentId(environmentId).Search(search).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationCustomDomainAPI.ListOrganizationCustomDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationCustomDomain`: OrganizationCustomDomainPaginatedResponseList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationCustomDomainAPI.ListOrganizationCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **string** | Only list the domains of this project. A project of another organization gives an empty list. | 
 **environmentId** | **string** | Only list the domains of this environment. An environment of another organization gives an empty list. | 
 **search** | **string** | Only list the domains whose name contains this text, ignoring case. The text is trimmed and matched literally, with no wildcard. A blank value does not filter. | 
 **page** | **int32** | The page to return, starting at 1. A page after the last one has no results. | [default to 1]
 **pageSize** | **int32** | The number of domains per page. | [default to 20]

### Return type

[**OrganizationCustomDomainPaginatedResponseList**](OrganizationCustomDomainPaginatedResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

