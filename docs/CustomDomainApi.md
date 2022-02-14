# \CustomDomainApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationCustomDomain**](CustomDomainApi.md#CreateApplicationCustomDomain) | **Post** /application/{applicationId}/customDomain | Add custom domain to the application.
[**DeleteCustomDomain**](CustomDomainApi.md#DeleteCustomDomain) | **Delete** /application/{applicationId}/customDomain/{customDomainId} | Delete a Custom Domain
[**EditCustomDomain**](CustomDomainApi.md#EditCustomDomain) | **Put** /application/{applicationId}/customDomain/{customDomainId} | Edit a Custom Domain
[**GetCustomDomainStatus**](CustomDomainApi.md#GetCustomDomainStatus) | **Get** /application/{applicationId}/customDomain/{customDomainId}/status | Get Custom Domain status
[**ListApplicationCustomDomain**](CustomDomainApi.md#ListApplicationCustomDomain) | **Get** /application/{applicationId}/customDomain | List application custom domains



## CreateApplicationCustomDomain

> CustomDomainResponse CreateApplicationCustomDomain(ctx, applicationId).CustomDomainRequest(customDomainRequest).Execute()

Add custom domain to the application.



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    customDomainRequest := *openapiclient.NewCustomDomainRequest("my.domain.tld") // CustomDomainRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainApi.CreateApplicationCustomDomain(context.Background(), applicationId).CustomDomainRequest(customDomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainApi.CreateApplicationCustomDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationCustomDomain`: CustomDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomDomainApi.CreateApplicationCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customDomainRequest** | [**CustomDomainRequest**](CustomDomainRequest.md) |  | 

### Return type

[**CustomDomainResponse**](CustomDomainResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomDomain

> DeleteCustomDomain(ctx, applicationId, customDomainId).Execute()

Delete a Custom Domain



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    customDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Custom Domain ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainApi.DeleteCustomDomain(context.Background(), applicationId, customDomainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainApi.DeleteCustomDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**customDomainId** | **string** | Custom Domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomDomainRequest struct via the builder pattern


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


## EditCustomDomain

> CustomDomainResponse EditCustomDomain(ctx, applicationId, customDomainId).CustomDomainRequest(customDomainRequest).Execute()

Edit a Custom Domain



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    customDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Custom Domain ID
    customDomainRequest := *openapiclient.NewCustomDomainRequest("my.domain.tld") // CustomDomainRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainApi.EditCustomDomain(context.Background(), applicationId, customDomainId).CustomDomainRequest(customDomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainApi.EditCustomDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditCustomDomain`: CustomDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomDomainApi.EditCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**customDomainId** | **string** | Custom Domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customDomainRequest** | [**CustomDomainRequest**](CustomDomainRequest.md) |  | 

### Return type

[**CustomDomainResponse**](CustomDomainResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomDomainStatus

> CustomDomainResponse GetCustomDomainStatus(ctx, applicationId, customDomainId).Execute()

Get Custom Domain status

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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    customDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Custom Domain ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainApi.GetCustomDomainStatus(context.Background(), applicationId, customDomainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainApi.GetCustomDomainStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomDomainStatus`: CustomDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomDomainApi.GetCustomDomainStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 
**customDomainId** | **string** | Custom Domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomDomainStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomDomainResponse**](CustomDomainResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationCustomDomain

> CustomDomainResponseList ListApplicationCustomDomain(ctx, applicationId).Execute()

List application custom domains

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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainApi.ListApplicationCustomDomain(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainApi.ListApplicationCustomDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationCustomDomain`: CustomDomainResponseList
    fmt.Fprintf(os.Stdout, "Response from `CustomDomainApi.ListApplicationCustomDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationCustomDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomDomainResponseList**](CustomDomainResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

