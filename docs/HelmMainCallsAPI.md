# \HelmMainCallsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHelm**](HelmMainCallsAPI.md#DeleteHelm) | **Delete** /helm/{helmId} | Delete helm
[**EditHelm**](HelmMainCallsAPI.md#EditHelm) | **Put** /helm/{helmId} | Edit helm
[**GetHelm**](HelmMainCallsAPI.md#GetHelm) | **Get** /helm/{helmId} | Get helm by ID
[**GetHelmStatus**](HelmMainCallsAPI.md#GetHelmStatus) | **Get** /helm/{helmId}/status | Get helm status



## DeleteHelm

> DeleteHelm(ctx, helmId).Execute()

Delete helm



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
    helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HelmMainCallsAPI.DeleteHelm(context.Background(), helmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmMainCallsAPI.DeleteHelm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHelmRequest struct via the builder pattern


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


## EditHelm

> HelmResponse EditHelm(ctx, helmId).HelmRequest(helmRequest).Execute()

Edit helm



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
    helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID
    helmRequest := *openapiclient.NewHelmRequest("Name_example") // HelmRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmMainCallsAPI.EditHelm(context.Background(), helmId).HelmRequest(helmRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmMainCallsAPI.EditHelm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditHelm`: HelmResponse
    fmt.Fprintf(os.Stdout, "Response from `HelmMainCallsAPI.EditHelm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditHelmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **helmRequest** | [**HelmRequest**](HelmRequest.md) |  | 

### Return type

[**HelmResponse**](HelmResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHelm

> HelmResponse GetHelm(ctx, helmId).Execute()

Get helm by ID

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
    helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmMainCallsAPI.GetHelm(context.Background(), helmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmMainCallsAPI.GetHelm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHelm`: HelmResponse
    fmt.Fprintf(os.Stdout, "Response from `HelmMainCallsAPI.GetHelm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHelmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HelmResponse**](HelmResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHelmStatus

> Status GetHelmStatus(ctx, helmId).Execute()

Get helm status

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
    helmId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Helm ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmMainCallsAPI.GetHelmStatus(context.Background(), helmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmMainCallsAPI.GetHelmStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHelmStatus`: Status
    fmt.Fprintf(os.Stdout, "Response from `HelmMainCallsAPI.GetHelmStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHelmStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

