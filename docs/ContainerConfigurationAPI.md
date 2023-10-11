# \ContainerConfigurationAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditContainerAdvancedSettings**](ContainerConfigurationAPI.md#EditContainerAdvancedSettings) | **Put** /container/{containerId}/advancedSettings | Edit advanced settings
[**EditContainerNetwork**](ContainerConfigurationAPI.md#EditContainerNetwork) | **Put** /container/{containerId}/network | Edit Container Network
[**GetContainerAdvancedSettings**](ContainerConfigurationAPI.md#GetContainerAdvancedSettings) | **Get** /container/{containerId}/advancedSettings | Get advanced settings
[**GetContainerNetwork**](ContainerConfigurationAPI.md#GetContainerNetwork) | **Get** /container/{containerId}/network | Get Container Network information



## EditContainerAdvancedSettings

> ContainerAdvancedSettings EditContainerAdvancedSettings(ctx, containerId).ContainerAdvancedSettings(containerAdvancedSettings).Execute()

Edit advanced settings



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    containerAdvancedSettings := *openapiclient.NewContainerAdvancedSettings() // ContainerAdvancedSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerConfigurationAPI.EditContainerAdvancedSettings(context.Background(), containerId).ContainerAdvancedSettings(containerAdvancedSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerConfigurationAPI.EditContainerAdvancedSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditContainerAdvancedSettings`: ContainerAdvancedSettings
    fmt.Fprintf(os.Stdout, "Response from `ContainerConfigurationAPI.EditContainerAdvancedSettings`: %v\n", resp)
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

 **containerAdvancedSettings** | [**ContainerAdvancedSettings**](ContainerAdvancedSettings.md) |  | 

### Return type

[**ContainerAdvancedSettings**](ContainerAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditContainerNetwork

> ContainerNetwork EditContainerNetwork(ctx, containerId).ContainerNetworkRequest(containerNetworkRequest).Execute()

Edit Container Network



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID
    containerNetworkRequest := *openapiclient.NewContainerNetworkRequest() // ContainerNetworkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerConfigurationAPI.EditContainerNetwork(context.Background(), containerId).ContainerNetworkRequest(containerNetworkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerConfigurationAPI.EditContainerNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditContainerNetwork`: ContainerNetwork
    fmt.Fprintf(os.Stdout, "Response from `ContainerConfigurationAPI.EditContainerNetwork`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerAdvancedSettings

> ContainerAdvancedSettings GetContainerAdvancedSettings(ctx, containerId).Execute()

Get advanced settings



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerConfigurationAPI.GetContainerAdvancedSettings(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerConfigurationAPI.GetContainerAdvancedSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerAdvancedSettings`: ContainerAdvancedSettings
    fmt.Fprintf(os.Stdout, "Response from `ContainerConfigurationAPI.GetContainerAdvancedSettings`: %v\n", resp)
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

[**ContainerAdvancedSettings**](ContainerAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerNetwork

> ContainerNetwork GetContainerNetwork(ctx, containerId).Execute()

Get Container Network information



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
    containerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Container ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerConfigurationAPI.GetContainerNetwork(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerConfigurationAPI.GetContainerNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerNetwork`: ContainerNetwork
    fmt.Fprintf(os.Stdout, "Response from `ContainerConfigurationAPI.GetContainerNetwork`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

