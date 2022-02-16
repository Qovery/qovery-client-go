# \ApplicationConfigurationApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditApplicationNetwork**](ApplicationConfigurationApi.md#EditApplicationNetwork) | **Put** /application/{applicationId}/network | Edit Application Network
[**GetApplicationNetwork**](ApplicationConfigurationApi.md#GetApplicationNetwork) | **Get** /application/{applicationId}/network | Get Application Network information



## EditApplicationNetwork

> ApplicationNetworkResponse EditApplicationNetwork(ctx, applicationId).ApplicationNetworkRequest(applicationNetworkRequest).Execute()

Edit Application Network



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
    applicationNetworkRequest := *openapiclient.NewApplicationNetworkRequest() // ApplicationNetworkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationConfigurationApi.EditApplicationNetwork(context.Background(), applicationId).ApplicationNetworkRequest(applicationNetworkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConfigurationApi.EditApplicationNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditApplicationNetwork`: ApplicationNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationConfigurationApi.EditApplicationNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditApplicationNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationNetworkRequest** | [**ApplicationNetworkRequest**](ApplicationNetworkRequest.md) |  | 

### Return type

[**ApplicationNetworkResponse**](ApplicationNetworkResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationNetwork

> ApplicationNetworkResponse GetApplicationNetwork(ctx, applicationId).Execute()

Get Application Network information



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationConfigurationApi.GetApplicationNetwork(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConfigurationApi.GetApplicationNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationNetwork`: ApplicationNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationConfigurationApi.GetApplicationNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationNetworkResponse**](ApplicationNetworkResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

