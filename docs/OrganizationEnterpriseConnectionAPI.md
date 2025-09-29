# \OrganizationEnterpriseConnectionAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationEnterpriseConnection**](OrganizationEnterpriseConnectionAPI.md#GetOrganizationEnterpriseConnection) | **Get** /organization/{organizationId}/enterpriseconnection/{connectionName} | Get enterprise connection
[**UpdateOrganizationEnterpriseConnection**](OrganizationEnterpriseConnectionAPI.md#UpdateOrganizationEnterpriseConnection) | **Put** /organization/{organizationId}/enterpriseconnection/{connectionName} | Update enterprise connection



## GetOrganizationEnterpriseConnection

> EnterpriseConnectionDto GetOrganizationEnterpriseConnection(ctx, organizationId, connectionName).Execute()

Get enterprise connection

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
	connectionName := "connectionName_example" // string | The name of the Organization's Enterprise Connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationEnterpriseConnectionAPI.GetOrganizationEnterpriseConnection(context.Background(), organizationId, connectionName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationEnterpriseConnectionAPI.GetOrganizationEnterpriseConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationEnterpriseConnection`: EnterpriseConnectionDto
	fmt.Fprintf(os.Stdout, "Response from `OrganizationEnterpriseConnectionAPI.GetOrganizationEnterpriseConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**connectionName** | **string** | The name of the Organization&#39;s Enterprise Connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationEnterpriseConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnterpriseConnectionDto**](EnterpriseConnectionDto.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationEnterpriseConnection

> EnterpriseConnectionDto UpdateOrganizationEnterpriseConnection(ctx, organizationId, connectionName).EnterpriseConnectionDto(enterpriseConnectionDto).Execute()

Update enterprise connection

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
	connectionName := "connectionName_example" // string | The name of the Organization's Enterprise Connection
	enterpriseConnectionDto := *openapiclient.NewEnterpriseConnectionDto("DefaultRole_example", false, map[string][]string{"key": []string{"Inner_example"}}) // EnterpriseConnectionDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationEnterpriseConnectionAPI.UpdateOrganizationEnterpriseConnection(context.Background(), organizationId, connectionName).EnterpriseConnectionDto(enterpriseConnectionDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationEnterpriseConnectionAPI.UpdateOrganizationEnterpriseConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationEnterpriseConnection`: EnterpriseConnectionDto
	fmt.Fprintf(os.Stdout, "Response from `OrganizationEnterpriseConnectionAPI.UpdateOrganizationEnterpriseConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**connectionName** | **string** | The name of the Organization&#39;s Enterprise Connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationEnterpriseConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enterpriseConnectionDto** | [**EnterpriseConnectionDto**](EnterpriseConnectionDto.md) |  | 

### Return type

[**EnterpriseConnectionDto**](EnterpriseConnectionDto.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

