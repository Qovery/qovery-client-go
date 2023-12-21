# \HelmsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneHelm**](HelmsAPI.md#CloneHelm) | **Post** /helm/{helmId}/clone | Clone helm
[**CreateHelm**](HelmsAPI.md#CreateHelm) | **Post** /environment/{environmentId}/helm | Create a helm
[**CreateHelmDefaultValues**](HelmsAPI.md#CreateHelmDefaultValues) | **Post** /environment/{environmentId}/helmDefaultValues | Get helm default values
[**GetDefaultHelmAdvancedSettings**](HelmsAPI.md#GetDefaultHelmAdvancedSettings) | **Get** /defaultHelmAdvancedSettings | List default helm advanced settings
[**GetEnvironmentHelmStatus**](HelmsAPI.md#GetEnvironmentHelmStatus) | **Get** /environment/{environmentId}/helm/status | List all environment helm statuses
[**ListHelms**](HelmsAPI.md#ListHelms) | **Get** /environment/{environmentId}/helm | List helms



## CloneHelm

> HelmResponse CloneHelm(ctx, helmId).CloneServiceRequest(cloneServiceRequest).Execute()

Clone helm



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
    cloneServiceRequest := *openapiclient.NewCloneServiceRequest("Name_example", "EnvironmentId_example") // CloneServiceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmsAPI.CloneHelm(context.Background(), helmId).CloneServiceRequest(cloneServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmsAPI.CloneHelm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneHelm`: HelmResponse
    fmt.Fprintf(os.Stdout, "Response from `HelmsAPI.CloneHelm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**helmId** | **string** | Helm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneHelmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneServiceRequest** | [**CloneServiceRequest**](CloneServiceRequest.md) |  | 

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


## CreateHelm

> HelmResponse CreateHelm(ctx, environmentId).HelmRequest(helmRequest).Execute()

Create a helm

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
    helmRequest := *openapiclient.NewHelmRequest("Name_example", false, openapiclient.HelmRequest_allOf_source{HelmRequestAllOfSourceOneOf: openapiclient.NewHelmRequestAllOfSourceOneOf()}, []string{"Arguments_example"}, *openapiclient.NewHelmRequestAllOfValuesOverride()) // HelmRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmsAPI.CreateHelm(context.Background(), environmentId).HelmRequest(helmRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmsAPI.CreateHelm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHelm`: HelmResponse
    fmt.Fprintf(os.Stdout, "Response from `HelmsAPI.CreateHelm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHelmRequest struct via the builder pattern


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


## CreateHelmDefaultValues

> string CreateHelmDefaultValues(ctx, environmentId).HelmDefaultValuesRequest(helmDefaultValuesRequest).Execute()

Get helm default values

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
    helmDefaultValuesRequest := *openapiclient.NewHelmDefaultValuesRequest(openapiclient.HelmRequest_allOf_source{HelmRequestAllOfSourceOneOf: openapiclient.NewHelmRequestAllOfSourceOneOf()}) // HelmDefaultValuesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmsAPI.CreateHelmDefaultValues(context.Background(), environmentId).HelmDefaultValuesRequest(helmDefaultValuesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmsAPI.CreateHelmDefaultValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHelmDefaultValues`: string
    fmt.Fprintf(os.Stdout, "Response from `HelmsAPI.CreateHelmDefaultValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHelmDefaultValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **helmDefaultValuesRequest** | [**HelmDefaultValuesRequest**](HelmDefaultValuesRequest.md) |  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultHelmAdvancedSettings

> map[string]interface{} GetDefaultHelmAdvancedSettings(ctx).Execute()

List default helm advanced settings

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmsAPI.GetDefaultHelmAdvancedSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmsAPI.GetDefaultHelmAdvancedSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultHelmAdvancedSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HelmsAPI.GetDefaultHelmAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultHelmAdvancedSettingsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentHelmStatus

> ReferenceObjectStatusResponseList GetEnvironmentHelmStatus(ctx, environmentId).Execute()

List all environment helm statuses



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmsAPI.GetEnvironmentHelmStatus(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmsAPI.GetEnvironmentHelmStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentHelmStatus`: ReferenceObjectStatusResponseList
    fmt.Fprintf(os.Stdout, "Response from `HelmsAPI.GetEnvironmentHelmStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentHelmStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReferenceObjectStatusResponseList**](ReferenceObjectStatusResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHelms

> HelmResponseList ListHelms(ctx, environmentId).Execute()

List helms

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HelmsAPI.ListHelms(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HelmsAPI.ListHelms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHelms`: HelmResponseList
    fmt.Fprintf(os.Stdout, "Response from `HelmsAPI.ListHelms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHelmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HelmResponseList**](HelmResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

