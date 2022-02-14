# \OrganizationMainCallsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganization**](OrganizationMainCallsApi.md#CreateOrganization) | **Post** /organization | Create an organization
[**DeleteOrganization**](OrganizationMainCallsApi.md#DeleteOrganization) | **Delete** /organization/{organizationId} | Delete an organization
[**EditOrganization**](OrganizationMainCallsApi.md#EditOrganization) | **Put** /organization/{organizationId} | Edit an organization
[**GetOrganization**](OrganizationMainCallsApi.md#GetOrganization) | **Get** /organization/{organizationId} | Get organization by ID
[**ListOrganization**](OrganizationMainCallsApi.md#ListOrganization) | **Get** /organization | List user organizations



## CreateOrganization

> OrganizationResponse CreateOrganization(ctx).OrganizationRequest(organizationRequest).Execute()

Create an organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationRequest := *openapiclient.NewOrganizationRequest("Name_example", "Plan_example") // OrganizationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationMainCallsApi.CreateOrganization(context.Background()).OrganizationRequest(organizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsApi.CreateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganization`: OrganizationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsApi.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationRequest** | [**OrganizationRequest**](OrganizationRequest.md) |  | 

### Return type

[**OrganizationResponse**](OrganizationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> DeleteOrganization(ctx, organizationId).Execute()

Delete an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationMainCallsApi.DeleteOrganization(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsApi.DeleteOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOrganization

> OrganizationResponse EditOrganization(ctx, organizationId).OrganizationEditRequest(organizationEditRequest).Execute()

Edit an organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    organizationEditRequest := *openapiclient.NewOrganizationEditRequest("Name_example") // OrganizationEditRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationMainCallsApi.EditOrganization(context.Background(), organizationId).OrganizationEditRequest(organizationEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsApi.EditOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditOrganization`: OrganizationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsApi.EditOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationEditRequest** | [**OrganizationEditRequest**](OrganizationEditRequest.md) |  | 

### Return type

[**OrganizationResponse**](OrganizationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> OrganizationResponse GetOrganization(ctx, organizationId).Execute()

Get organization by ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationMainCallsApi.GetOrganization(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsApi.GetOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganization`: OrganizationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsApi.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationResponse**](OrganizationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganization

> OrganizationResponseList ListOrganization(ctx).Execute()

List user organizations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationMainCallsApi.ListOrganization(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsApi.ListOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganization`: OrganizationResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsApi.ListOrganization`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationRequest struct via the builder pattern


### Return type

[**OrganizationResponseList**](OrganizationResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

