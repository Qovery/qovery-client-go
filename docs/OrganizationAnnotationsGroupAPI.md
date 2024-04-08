# \OrganizationAnnotationsGroupAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationAnnotationsGroup**](OrganizationAnnotationsGroupAPI.md#CreateOrganizationAnnotationsGroup) | **Post** /organization/{organizationId}/annotationsGroups | Create an organization annotations group
[**DeleteOrganizationAnnotationsGroup**](OrganizationAnnotationsGroupAPI.md#DeleteOrganizationAnnotationsGroup) | **Delete** /organization/{organizationId}/annotationsGroups/{annotationsGroupId} | Delete organization annotations group
[**EditOrganizationAnnotationsGroup**](OrganizationAnnotationsGroupAPI.md#EditOrganizationAnnotationsGroup) | **Put** /organization/{organizationId}/annotationsGroups/{annotationsGroupId} | Edit organization annotations group
[**GetOrganizationAnnotationsGroup**](OrganizationAnnotationsGroupAPI.md#GetOrganizationAnnotationsGroup) | **Get** /organization/{organizationId}/annotationsGroups/{annotationsGroupId} | Get organization annotations group
[**GetOrganizationAnnotationsGroupAssociatedItems**](OrganizationAnnotationsGroupAPI.md#GetOrganizationAnnotationsGroupAssociatedItems) | **Get** /organization/{organizationId}/annotationsGroups/{annotationsGroupId}/associatedItems | Get organization annotations group associated items
[**ListOrganizationAnnotationsGroup**](OrganizationAnnotationsGroupAPI.md#ListOrganizationAnnotationsGroup) | **Get** /organization/{organizationId}/annotationsGroups | List organization annotations group



## CreateOrganizationAnnotationsGroup

> OrganizationAnnotationsGroupResponse CreateOrganizationAnnotationsGroup(ctx, organizationId).OrganizationAnnotationsGroupCreateRequest(organizationAnnotationsGroupCreateRequest).Execute()

Create an organization annotations group



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
    organizationAnnotationsGroupCreateRequest := *openapiclient.NewOrganizationAnnotationsGroupCreateRequest("Name_example", []openapiclient.Annotation{*openapiclient.NewAnnotation("Key_example", "Value_example")}, []openapiclient.OrganizationAnnotationsGroupScopeEnum{openapiclient.OrganizationAnnotationsGroupScopeEnum("PERSISTENT_VOLUME_CLAIMS")}) // OrganizationAnnotationsGroupCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAnnotationsGroupAPI.CreateOrganizationAnnotationsGroup(context.Background(), organizationId).OrganizationAnnotationsGroupCreateRequest(organizationAnnotationsGroupCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAnnotationsGroupAPI.CreateOrganizationAnnotationsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAnnotationsGroup`: OrganizationAnnotationsGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAnnotationsGroupAPI.CreateOrganizationAnnotationsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAnnotationsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationAnnotationsGroupCreateRequest** | [**OrganizationAnnotationsGroupCreateRequest**](OrganizationAnnotationsGroupCreateRequest.md) |  | 

### Return type

[**OrganizationAnnotationsGroupResponse**](OrganizationAnnotationsGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationAnnotationsGroup

> DeleteOrganizationAnnotationsGroup(ctx, organizationId, annotationsGroupId).Execute()

Delete organization annotations group



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
    annotationsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization annotations group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrganizationAnnotationsGroupAPI.DeleteOrganizationAnnotationsGroup(context.Background(), organizationId, annotationsGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAnnotationsGroupAPI.DeleteOrganizationAnnotationsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**annotationsGroupId** | **string** | Organization annotations group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAnnotationsGroupRequest struct via the builder pattern


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


## EditOrganizationAnnotationsGroup

> OrganizationAnnotationsGroupResponse EditOrganizationAnnotationsGroup(ctx, organizationId, annotationsGroupId).OrganizationAnnotationsGroupCreateRequest(organizationAnnotationsGroupCreateRequest).Execute()

Edit organization annotations group



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
    annotationsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization annotations group ID
    organizationAnnotationsGroupCreateRequest := *openapiclient.NewOrganizationAnnotationsGroupCreateRequest("Name_example", []openapiclient.Annotation{*openapiclient.NewAnnotation("Key_example", "Value_example")}, []openapiclient.OrganizationAnnotationsGroupScopeEnum{openapiclient.OrganizationAnnotationsGroupScopeEnum("PERSISTENT_VOLUME_CLAIMS")}) // OrganizationAnnotationsGroupCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAnnotationsGroupAPI.EditOrganizationAnnotationsGroup(context.Background(), organizationId, annotationsGroupId).OrganizationAnnotationsGroupCreateRequest(organizationAnnotationsGroupCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAnnotationsGroupAPI.EditOrganizationAnnotationsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditOrganizationAnnotationsGroup`: OrganizationAnnotationsGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAnnotationsGroupAPI.EditOrganizationAnnotationsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**annotationsGroupId** | **string** | Organization annotations group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganizationAnnotationsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **organizationAnnotationsGroupCreateRequest** | [**OrganizationAnnotationsGroupCreateRequest**](OrganizationAnnotationsGroupCreateRequest.md) |  | 

### Return type

[**OrganizationAnnotationsGroupResponse**](OrganizationAnnotationsGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAnnotationsGroup

> OrganizationAnnotationsGroupResponse GetOrganizationAnnotationsGroup(ctx, organizationId, annotationsGroupId).Execute()

Get organization annotations group



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
    annotationsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization annotations group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAnnotationsGroupAPI.GetOrganizationAnnotationsGroup(context.Background(), organizationId, annotationsGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAnnotationsGroupAPI.GetOrganizationAnnotationsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAnnotationsGroup`: OrganizationAnnotationsGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAnnotationsGroupAPI.GetOrganizationAnnotationsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**annotationsGroupId** | **string** | Organization annotations group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAnnotationsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationAnnotationsGroupResponse**](OrganizationAnnotationsGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAnnotationsGroupAssociatedItems

> OrganizationAnnotationsGroupAssociatedItemsResponseList GetOrganizationAnnotationsGroupAssociatedItems(ctx, organizationId, annotationsGroupId).Execute()

Get organization annotations group associated items



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
    annotationsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization annotations group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAnnotationsGroupAPI.GetOrganizationAnnotationsGroupAssociatedItems(context.Background(), organizationId, annotationsGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAnnotationsGroupAPI.GetOrganizationAnnotationsGroupAssociatedItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAnnotationsGroupAssociatedItems`: OrganizationAnnotationsGroupAssociatedItemsResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAnnotationsGroupAPI.GetOrganizationAnnotationsGroupAssociatedItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**annotationsGroupId** | **string** | Organization annotations group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAnnotationsGroupAssociatedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationAnnotationsGroupAssociatedItemsResponseList**](OrganizationAnnotationsGroupAssociatedItemsResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationAnnotationsGroup

> OrganizationAnnotationsGroupResponseList ListOrganizationAnnotationsGroup(ctx, organizationId).Execute()

List organization annotations group



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
    resp, r, err := apiClient.OrganizationAnnotationsGroupAPI.ListOrganizationAnnotationsGroup(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAnnotationsGroupAPI.ListOrganizationAnnotationsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizationAnnotationsGroup`: OrganizationAnnotationsGroupResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAnnotationsGroupAPI.ListOrganizationAnnotationsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationAnnotationsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationAnnotationsGroupResponseList**](OrganizationAnnotationsGroupResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

