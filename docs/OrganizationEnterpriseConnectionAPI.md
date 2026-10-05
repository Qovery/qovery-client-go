# \OrganizationEnterpriseConnectionAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnterpriseConnectionRoles**](OrganizationEnterpriseConnectionAPI.md#GetEnterpriseConnectionRoles) | **Get** /account/enterpriseconnection/roles | Resolve enterprise connection roles
[**GetOrganizationEnterpriseConnection**](OrganizationEnterpriseConnectionAPI.md#GetOrganizationEnterpriseConnection) | **Get** /organization/{organizationId}/enterpriseconnection/{connectionName} | Get enterprise connection
[**ListOrganizationEnterpriseConnections**](OrganizationEnterpriseConnectionAPI.md#ListOrganizationEnterpriseConnections) | **Get** /organization/{organizationId}/enterpriseconnection | List enterprise connections
[**NotifyEnterpriseMemberAccessUpdated**](OrganizationEnterpriseConnectionAPI.md#NotifyEnterpriseMemberAccessUpdated) | **Post** /account/enterpriseconnection/notifyMemberAccessUpdated | Notify enterprise member access changes
[**UpdateOrganizationEnterpriseConnection**](OrganizationEnterpriseConnectionAPI.md#UpdateOrganizationEnterpriseConnection) | **Put** /organization/{organizationId}/enterpriseconnection/{connectionName} | Update enterprise connection



## GetEnterpriseConnectionRoles

> EnterpriseConnectionAccessList GetEnterpriseConnectionRoles(ctx).XQoveryAuth0PostLoginToken(xQoveryAuth0PostLoginToken).ConnectionName(connectionName).FederatedGroups(federatedGroups).UserSub(userSub).Execute()

Resolve enterprise connection roles



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
	xQoveryAuth0PostLoginToken := "xQoveryAuth0PostLoginToken_example" // string | 
	connectionName := "connectionName_example" // string | 
	federatedGroups := "federatedGroups_example" // string | 
	userSub := "userSub_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationEnterpriseConnectionAPI.GetEnterpriseConnectionRoles(context.Background()).XQoveryAuth0PostLoginToken(xQoveryAuth0PostLoginToken).ConnectionName(connectionName).FederatedGroups(federatedGroups).UserSub(userSub).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationEnterpriseConnectionAPI.GetEnterpriseConnectionRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnterpriseConnectionRoles`: EnterpriseConnectionAccessList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationEnterpriseConnectionAPI.GetEnterpriseConnectionRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnterpriseConnectionRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xQoveryAuth0PostLoginToken** | **string** |  | 
 **connectionName** | **string** |  | 
 **federatedGroups** | **string** |  | 
 **userSub** | **string** |  | 

### Return type

[**EnterpriseConnectionAccessList**](EnterpriseConnectionAccessList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## ListOrganizationEnterpriseConnections

> EnterpriseConnectionResponseList ListOrganizationEnterpriseConnections(ctx, organizationId).Execute()

List enterprise connections

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
	resp, r, err := apiClient.OrganizationEnterpriseConnectionAPI.ListOrganizationEnterpriseConnections(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationEnterpriseConnectionAPI.ListOrganizationEnterpriseConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationEnterpriseConnections`: EnterpriseConnectionResponseList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationEnterpriseConnectionAPI.ListOrganizationEnterpriseConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationEnterpriseConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnterpriseConnectionResponseList**](EnterpriseConnectionResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotifyEnterpriseMemberAccessUpdated

> NotifyEnterpriseMemberAccessUpdated(ctx).XQoveryAuth0PostLoginToken(xQoveryAuth0PostLoginToken).EnterpriseConnectionMemberAccessUpdateRequest(enterpriseConnectionMemberAccessUpdateRequest).Execute()

Notify enterprise member access changes



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
	xQoveryAuth0PostLoginToken := "xQoveryAuth0PostLoginToken_example" // string | 
	enterpriseConnectionMemberAccessUpdateRequest := *openapiclient.NewEnterpriseConnectionMemberAccessUpdateRequest("UserId_example", []string{"AddedOrganizationIds_example"}, []string{"RemovedOrganizationIds_example"}, []openapiclient.MemberAccessRoleUpdated{*openapiclient.NewMemberAccessRoleUpdated("OrganizationId_example", "Role_example")}) // EnterpriseConnectionMemberAccessUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationEnterpriseConnectionAPI.NotifyEnterpriseMemberAccessUpdated(context.Background()).XQoveryAuth0PostLoginToken(xQoveryAuth0PostLoginToken).EnterpriseConnectionMemberAccessUpdateRequest(enterpriseConnectionMemberAccessUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationEnterpriseConnectionAPI.NotifyEnterpriseMemberAccessUpdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotifyEnterpriseMemberAccessUpdatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xQoveryAuth0PostLoginToken** | **string** |  | 
 **enterpriseConnectionMemberAccessUpdateRequest** | [**EnterpriseConnectionMemberAccessUpdateRequest**](EnterpriseConnectionMemberAccessUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

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
	enterpriseConnectionDto := *openapiclient.NewEnterpriseConnectionDto("ConnectionName_example", "DefaultRole_example", false, map[string][]string{"key": []string{"Inner_example"}}) // EnterpriseConnectionDto |  (optional)

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

