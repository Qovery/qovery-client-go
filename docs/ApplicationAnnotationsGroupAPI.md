# \ApplicationAnnotationsGroupAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAnnotationsGroupToApplication**](ApplicationAnnotationsGroupAPI.md#AddAnnotationsGroupToApplication) | **Post** /application/{applicationId}/annotationsGroup/{annotationsGroupId} | Add annotations group to application
[**DeleteAnnotationsGroupToApplication**](ApplicationAnnotationsGroupAPI.md#DeleteAnnotationsGroupToApplication) | **Delete** /application/{applicationId}/annotationsGroup/{annotationsGroupId} | Delete annotations group to application
[**ListApplicationAnnotationsGroup**](ApplicationAnnotationsGroupAPI.md#ListApplicationAnnotationsGroup) | **Get** /application/{applicationId}/annotationsGroup | List application annotations group



## AddAnnotationsGroupToApplication

> AddAnnotationsGroupToApplication(ctx, applicationId, annotationsGroupId).Execute()

Add annotations group to application



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    annotationsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization annotations group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationAnnotationsGroupAPI.AddAnnotationsGroupToApplication(context.Background(), applicationId, annotationsGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAnnotationsGroupAPI.AddAnnotationsGroupToApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**annotationsGroupId** | **string** | Organization annotations group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAnnotationsGroupToApplicationRequest struct via the builder pattern


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


## DeleteAnnotationsGroupToApplication

> DeleteAnnotationsGroupToApplication(ctx, applicationId, annotationsGroupId).Execute()

Delete annotations group to application



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    annotationsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization annotations group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationAnnotationsGroupAPI.DeleteAnnotationsGroupToApplication(context.Background(), applicationId, annotationsGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAnnotationsGroupAPI.DeleteAnnotationsGroupToApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**annotationsGroupId** | **string** | Organization annotations group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnnotationsGroupToApplicationRequest struct via the builder pattern


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


## ListApplicationAnnotationsGroup

> []OrganizationAnnotationsGroupResponse ListApplicationAnnotationsGroup(ctx, applicationId).Execute()

List application annotations group



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAnnotationsGroupAPI.ListApplicationAnnotationsGroup(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAnnotationsGroupAPI.ListApplicationAnnotationsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationAnnotationsGroup`: []OrganizationAnnotationsGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAnnotationsGroupAPI.ListApplicationAnnotationsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationAnnotationsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OrganizationAnnotationsGroupResponse**](OrganizationAnnotationsGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

