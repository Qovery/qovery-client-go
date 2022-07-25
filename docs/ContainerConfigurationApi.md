# \ContainerConfigurationApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditContainerAdvancedSettings**](ContainerConfigurationApi.md#EditContainerAdvancedSettings) | **Put** /container/{containerId}/advancedSettings | NOT YET IMPLEMENTED - Edit advanced settings
[**EditContainerNetwork**](ContainerConfigurationApi.md#EditContainerNetwork) | **Put** /container/{containerId}/network | NOT YET IMPLEMENTED - Edit Container Network
[**GetContainerAdvancedSettings**](ContainerConfigurationApi.md#GetContainerAdvancedSettings) | **Get** /container/{containerId}/advancedSettings | NOT YET IMPLEMENTED - Get advanced settings
[**GetContainerNetwork**](ContainerConfigurationApi.md#GetContainerNetwork) | **Get** /container/{containerId}/network | NOT YET IMPLEMENTED - Get Container Network information



## EditContainerAdvancedSettings

> ContainerAdvancedSettingsResponse EditContainerAdvancedSettings(ctx, containerId).ContainerAdvancedSettings(containerAdvancedSettings).Execute()

NOT YET IMPLEMENTED - Edit advanced settings



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    containerAdvancedSettings := []openapiclient.ContainerAdvancedSettings{*openapiclient.NewContainerAdvancedSettings()} // []ContainerAdvancedSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerConfigurationApi.EditContainerAdvancedSettings(context.Background(), containerId).ContainerAdvancedSettings(containerAdvancedSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerConfigurationApi.EditContainerAdvancedSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditContainerAdvancedSettings`: ContainerAdvancedSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `ContainerConfigurationApi.EditContainerAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditContainerAdvancedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerAdvancedSettings** | [**[]ContainerAdvancedSettings**](ContainerAdvancedSettings.md) |  | 

### Return type

[**ContainerAdvancedSettingsResponse**](ContainerAdvancedSettingsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditContainerNetwork

> ContainerNetwork EditContainerNetwork(ctx, containerId).ContainerNetworkRequest(containerNetworkRequest).Execute()

NOT YET IMPLEMENTED - Edit Container Network



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    containerNetworkRequest := *openapiclient.NewContainerNetworkRequest() // ContainerNetworkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerConfigurationApi.EditContainerNetwork(context.Background(), containerId).ContainerNetworkRequest(containerNetworkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerConfigurationApi.EditContainerNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditContainerNetwork`: ContainerNetwork
    fmt.Fprintf(os.Stdout, "Response from `ContainerConfigurationApi.EditContainerNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditContainerNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerNetworkRequest** | [**ContainerNetworkRequest**](ContainerNetworkRequest.md) |  | 

### Return type

[**ContainerNetwork**](ContainerNetwork.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerAdvancedSettings

> ContainerAdvancedSettingsResponse GetContainerAdvancedSettings(ctx, containerId).Execute()

NOT YET IMPLEMENTED - Get advanced settings



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerConfigurationApi.GetContainerAdvancedSettings(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerConfigurationApi.GetContainerAdvancedSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerAdvancedSettings`: ContainerAdvancedSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `ContainerConfigurationApi.GetContainerAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerAdvancedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerAdvancedSettingsResponse**](ContainerAdvancedSettingsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerNetwork

> ContainerNetwork GetContainerNetwork(ctx, containerId).Execute()

NOT YET IMPLEMENTED - Get Container Network information



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerConfigurationApi.GetContainerNetwork(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerConfigurationApi.GetContainerNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerNetwork`: ContainerNetwork
    fmt.Fprintf(os.Stdout, "Response from `ContainerConfigurationApi.GetContainerNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** | Container ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerNetwork**](ContainerNetwork.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

