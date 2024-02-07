# \ApplicationConfigurationAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditAdvancedSettings**](ApplicationConfigurationAPI.md#EditAdvancedSettings) | **Put** /application/{applicationId}/advancedSettings | Edit advanced settings
[**EditApplicationNetwork**](ApplicationConfigurationAPI.md#EditApplicationNetwork) | **Put** /application/{applicationId}/network | Edit Application Network
[**GetAdvancedSettings**](ApplicationConfigurationAPI.md#GetAdvancedSettings) | **Get** /application/{applicationId}/advancedSettings | Get advanced settings
[**GetApplicationNetwork**](ApplicationConfigurationAPI.md#GetApplicationNetwork) | **Get** /application/{applicationId}/network | Get Application Network information



## EditAdvancedSettings

> ApplicationAdvancedSettings EditAdvancedSettings(ctx, applicationId).ApplicationAdvancedSettings(applicationAdvancedSettings).Execute()

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
	applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
	applicationAdvancedSettings := *openapiclient.NewApplicationAdvancedSettings() // ApplicationAdvancedSettings |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationConfigurationAPI.EditAdvancedSettings(context.Background(), applicationId).ApplicationAdvancedSettings(applicationAdvancedSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConfigurationAPI.EditAdvancedSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAdvancedSettings`: ApplicationAdvancedSettings
	fmt.Fprintf(os.Stdout, "Response from `ApplicationConfigurationAPI.EditAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAdvancedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationAdvancedSettings** | [**ApplicationAdvancedSettings**](ApplicationAdvancedSettings.md) |  | 

### Return type

[**ApplicationAdvancedSettings**](ApplicationAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditApplicationNetwork

> ApplicationNetwork EditApplicationNetwork(ctx, applicationId).ApplicationNetworkRequest(applicationNetworkRequest).Execute()

Edit Application Network



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
	applicationNetworkRequest := *openapiclient.NewApplicationNetworkRequest() // ApplicationNetworkRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationConfigurationAPI.EditApplicationNetwork(context.Background(), applicationId).ApplicationNetworkRequest(applicationNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConfigurationAPI.EditApplicationNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditApplicationNetwork`: ApplicationNetwork
	fmt.Fprintf(os.Stdout, "Response from `ApplicationConfigurationAPI.EditApplicationNetwork`: %v\n", resp)
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

[**ApplicationNetwork**](ApplicationNetwork.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdvancedSettings

> ApplicationAdvancedSettings GetAdvancedSettings(ctx, applicationId).Execute()

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
	applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationConfigurationAPI.GetAdvancedSettings(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConfigurationAPI.GetAdvancedSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdvancedSettings`: ApplicationAdvancedSettings
	fmt.Fprintf(os.Stdout, "Response from `ApplicationConfigurationAPI.GetAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdvancedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationAdvancedSettings**](ApplicationAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationNetwork

> ApplicationNetwork GetApplicationNetwork(ctx, applicationId).Execute()

Get Application Network information



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
	resp, r, err := apiClient.ApplicationConfigurationAPI.GetApplicationNetwork(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConfigurationAPI.GetApplicationNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationNetwork`: ApplicationNetwork
	fmt.Fprintf(os.Stdout, "Response from `ApplicationConfigurationAPI.GetApplicationNetwork`: %v\n", resp)
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

[**ApplicationNetwork**](ApplicationNetwork.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

