# \OrganizationCustomRoleAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationCustomRole**](OrganizationCustomRoleAPI.md#CreateOrganizationCustomRole) | **Post** /organization/{organizationId}/customRole | Create an organization custom role
[**DeleteOrganizationCustomRole**](OrganizationCustomRoleAPI.md#DeleteOrganizationCustomRole) | **Delete** /organization/{organizationId}/customRole/{customRoleId} | Delete organization custom role
[**EditOrganizationCustomRole**](OrganizationCustomRoleAPI.md#EditOrganizationCustomRole) | **Put** /organization/{organizationId}/customRole/{customRoleId} | Edit an organization custom role
[**GetOrganizationCustomRole**](OrganizationCustomRoleAPI.md#GetOrganizationCustomRole) | **Get** /organization/{organizationId}/customRole/{customRoleId} | Get an organization custom role 
[**ListOrganizationCustomRoles**](OrganizationCustomRoleAPI.md#ListOrganizationCustomRoles) | **Get** /organization/{organizationId}/customRole | List organization custom roles



## CreateOrganizationCustomRole

> OrganizationCustomRole CreateOrganizationCustomRole(ctx, organizationId).OrganizationCustomRoleCreateRequest(organizationCustomRoleCreateRequest).Execute()

Create an organization custom role



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
    organizationCustomRoleCreateRequest := *openapiclient.NewOrganizationCustomRoleCreateRequest("Name_example") // OrganizationCustomRoleCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationCustomRoleAPI.CreateOrganizationCustomRole(context.Background(), organizationId).OrganizationCustomRoleCreateRequest(organizationCustomRoleCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationCustomRoleAPI.CreateOrganizationCustomRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCustomRole`: OrganizationCustomRole
    fmt.Fprintf(os.Stdout, "Response from `OrganizationCustomRoleAPI.CreateOrganizationCustomRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationCustomRoleCreateRequest** | [**OrganizationCustomRoleCreateRequest**](OrganizationCustomRoleCreateRequest.md) |  | 

### Return type

[**OrganizationCustomRole**](OrganizationCustomRole.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationCustomRole

> DeleteOrganizationCustomRole(ctx, organizationId, customRoleId).Execute()

Delete organization custom role



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
    customRoleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Custom Role ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrganizationCustomRoleAPI.DeleteOrganizationCustomRole(context.Background(), organizationId, customRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationCustomRoleAPI.DeleteOrganizationCustomRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**customRoleId** | **string** | Custom Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationCustomRoleRequest struct via the builder pattern


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


## EditOrganizationCustomRole

> OrganizationCustomRole EditOrganizationCustomRole(ctx, organizationId, customRoleId).OrganizationCustomRoleUpdateRequest(organizationCustomRoleUpdateRequest).Execute()

Edit an organization custom role



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
    customRoleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Custom Role ID
    organizationCustomRoleUpdateRequest := *openapiclient.NewOrganizationCustomRoleUpdateRequest("Name_example", []openapiclient.OrganizationCustomRoleUpdateRequestClusterPermissionsInner{*openapiclient.NewOrganizationCustomRoleUpdateRequestClusterPermissionsInner()}, []openapiclient.OrganizationCustomRoleUpdateRequestProjectPermissionsInner{*openapiclient.NewOrganizationCustomRoleUpdateRequestProjectPermissionsInner()}) // OrganizationCustomRoleUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationCustomRoleAPI.EditOrganizationCustomRole(context.Background(), organizationId, customRoleId).OrganizationCustomRoleUpdateRequest(organizationCustomRoleUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationCustomRoleAPI.EditOrganizationCustomRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditOrganizationCustomRole`: OrganizationCustomRole
    fmt.Fprintf(os.Stdout, "Response from `OrganizationCustomRoleAPI.EditOrganizationCustomRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**customRoleId** | **string** | Custom Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganizationCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **organizationCustomRoleUpdateRequest** | [**OrganizationCustomRoleUpdateRequest**](OrganizationCustomRoleUpdateRequest.md) |  | 

### Return type

[**OrganizationCustomRole**](OrganizationCustomRole.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCustomRole

> OrganizationCustomRole GetOrganizationCustomRole(ctx, organizationId, customRoleId).Execute()

Get an organization custom role 



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
    customRoleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Custom Role ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationCustomRoleAPI.GetOrganizationCustomRole(context.Background(), organizationId, customRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationCustomRoleAPI.GetOrganizationCustomRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCustomRole`: OrganizationCustomRole
    fmt.Fprintf(os.Stdout, "Response from `OrganizationCustomRoleAPI.GetOrganizationCustomRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**customRoleId** | **string** | Custom Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationCustomRole**](OrganizationCustomRole.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationCustomRoles

> OrganizationCustomRoleList ListOrganizationCustomRoles(ctx, organizationId).Execute()

List organization custom roles



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
    resp, r, err := apiClient.OrganizationCustomRoleAPI.ListOrganizationCustomRoles(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationCustomRoleAPI.ListOrganizationCustomRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizationCustomRoles`: OrganizationCustomRoleList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationCustomRoleAPI.ListOrganizationCustomRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationCustomRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationCustomRoleList**](OrganizationCustomRoleList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

