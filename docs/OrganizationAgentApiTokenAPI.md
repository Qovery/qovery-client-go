# \OrganizationAgentApiTokenAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationAgentApiToken**](OrganizationAgentApiTokenAPI.md#CreateOrganizationAgentApiToken) | **Post** /organization/{organizationId}/agentApiToken | Create an organization agent api token
[**DeleteOrganizationAgentApiToken**](OrganizationAgentApiTokenAPI.md#DeleteOrganizationAgentApiToken) | **Delete** /organization/{organizationId}/agentApiToken/{agentApiTokenId} | Delete organization agent api token
[**ListOrganizationAgentApiTokens**](OrganizationAgentApiTokenAPI.md#ListOrganizationAgentApiTokens) | **Get** /organization/{organizationId}/agentApiToken | List organization agent api tokens



## CreateOrganizationAgentApiToken

> OrganizationAgentApiTokenCreate CreateOrganizationAgentApiToken(ctx, organizationId).OrganizationAgentApiTokenCreateRequest(organizationAgentApiTokenCreateRequest).Execute()

Create an organization agent api token



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
	organizationAgentApiTokenCreateRequest := *openapiclient.NewOrganizationAgentApiTokenCreateRequest("Name_example", "OpaPolicy_example") // OrganizationAgentApiTokenCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAgentApiTokenAPI.CreateOrganizationAgentApiToken(context.Background(), organizationId).OrganizationAgentApiTokenCreateRequest(organizationAgentApiTokenCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAgentApiTokenAPI.CreateOrganizationAgentApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationAgentApiToken`: OrganizationAgentApiTokenCreate
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAgentApiTokenAPI.CreateOrganizationAgentApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAgentApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationAgentApiTokenCreateRequest** | [**OrganizationAgentApiTokenCreateRequest**](OrganizationAgentApiTokenCreateRequest.md) |  | 

### Return type

[**OrganizationAgentApiTokenCreate**](OrganizationAgentApiTokenCreate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationAgentApiToken

> DeleteOrganizationAgentApiToken(ctx, organizationId, agentApiTokenId).Execute()

Delete organization agent api token



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
	agentApiTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization Agent Api Token ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationAgentApiTokenAPI.DeleteOrganizationAgentApiToken(context.Background(), organizationId, agentApiTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAgentApiTokenAPI.DeleteOrganizationAgentApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**agentApiTokenId** | **string** | Organization Agent Api Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAgentApiTokenRequest struct via the builder pattern


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


## ListOrganizationAgentApiTokens

> OrganizationAgentApiTokenResponseList ListOrganizationAgentApiTokens(ctx, organizationId).Execute()

List organization agent api tokens



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
	resp, r, err := apiClient.OrganizationAgentApiTokenAPI.ListOrganizationAgentApiTokens(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAgentApiTokenAPI.ListOrganizationAgentApiTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationAgentApiTokens`: OrganizationAgentApiTokenResponseList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAgentApiTokenAPI.ListOrganizationAgentApiTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationAgentApiTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationAgentApiTokenResponseList**](OrganizationAgentApiTokenResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

