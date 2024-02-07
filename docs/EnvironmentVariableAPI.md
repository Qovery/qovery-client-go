# \EnvironmentVariableAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironmentEnvironmentVariable**](EnvironmentVariableAPI.md#CreateEnvironmentEnvironmentVariable) | **Post** /environment/{environmentId}/environmentVariable | Add an environment variable to the environment
[**CreateEnvironmentEnvironmentVariableAlias**](EnvironmentVariableAPI.md#CreateEnvironmentEnvironmentVariableAlias) | **Post** /environment/{environmentId}/environmentVariable/{environmentVariableId}/alias | Create an environment variable alias at the environment level
[**CreateEnvironmentEnvironmentVariableOverride**](EnvironmentVariableAPI.md#CreateEnvironmentEnvironmentVariableOverride) | **Post** /environment/{environmentId}/environmentVariable/{environmentVariableId}/override | Create an environment variable override at the environment level
[**DeleteEnvironmentEnvironmentVariable**](EnvironmentVariableAPI.md#DeleteEnvironmentEnvironmentVariable) | **Delete** /environment/{environmentId}/environmentVariable/{environmentVariableId} | Delete an environment variable from an environment
[**EditEnvironmentEnvironmentVariable**](EnvironmentVariableAPI.md#EditEnvironmentEnvironmentVariable) | **Put** /environment/{environmentId}/environmentVariable/{environmentVariableId} | Edit an environment variable belonging to the environment
[**ListEnvironmentEnvironmentVariable**](EnvironmentVariableAPI.md#ListEnvironmentEnvironmentVariable) | **Get** /environment/{environmentId}/environmentVariable | List environment variables



## CreateEnvironmentEnvironmentVariable

> EnvironmentVariable CreateEnvironmentEnvironmentVariable(ctx, environmentId).EnvironmentVariableRequest(environmentVariableRequest).Execute()

Add an environment variable to the environment



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
	environmentVariableRequest := *openapiclient.NewEnvironmentVariableRequest("Key_example") // EnvironmentVariableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentVariableAPI.CreateEnvironmentEnvironmentVariable(context.Background(), environmentId).EnvironmentVariableRequest(environmentVariableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableAPI.CreateEnvironmentEnvironmentVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnvironmentEnvironmentVariable`: EnvironmentVariable
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableAPI.CreateEnvironmentEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentVariableRequest** | [**EnvironmentVariableRequest**](EnvironmentVariableRequest.md) |  | 

### Return type

[**EnvironmentVariable**](EnvironmentVariable.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvironmentEnvironmentVariableAlias

> EnvironmentVariable CreateEnvironmentEnvironmentVariableAlias(ctx, environmentId, environmentVariableId).Key(key).Execute()

Create an environment variable alias at the environment level



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
	environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
	key := *openapiclient.NewKey("Key_example") // Key |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentVariableAPI.CreateEnvironmentEnvironmentVariableAlias(context.Background(), environmentId, environmentVariableId).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableAPI.CreateEnvironmentEnvironmentVariableAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnvironmentEnvironmentVariableAlias`: EnvironmentVariable
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableAPI.CreateEnvironmentEnvironmentVariableAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentEnvironmentVariableAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | [**Key**](Key.md) |  | 

### Return type

[**EnvironmentVariable**](EnvironmentVariable.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvironmentEnvironmentVariableOverride

> EnvironmentVariable CreateEnvironmentEnvironmentVariableOverride(ctx, environmentId, environmentVariableId).Value(value).Execute()

Create an environment variable override at the environment level



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
	environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
	value := *openapiclient.NewValue() // Value |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentVariableAPI.CreateEnvironmentEnvironmentVariableOverride(context.Background(), environmentId, environmentVariableId).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableAPI.CreateEnvironmentEnvironmentVariableOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnvironmentEnvironmentVariableOverride`: EnvironmentVariable
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableAPI.CreateEnvironmentEnvironmentVariableOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentEnvironmentVariableOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | [**Value**](Value.md) |  | 

### Return type

[**EnvironmentVariable**](EnvironmentVariable.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentEnvironmentVariable

> DeleteEnvironmentEnvironmentVariable(ctx, environmentId, environmentVariableId).Execute()

Delete an environment variable from an environment



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
	environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentVariableAPI.DeleteEnvironmentEnvironmentVariable(context.Background(), environmentId, environmentVariableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableAPI.DeleteEnvironmentEnvironmentVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentEnvironmentVariableRequest struct via the builder pattern


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


## EditEnvironmentEnvironmentVariable

> EnvironmentVariable EditEnvironmentEnvironmentVariable(ctx, environmentId, environmentVariableId).EnvironmentVariableEditRequest(environmentVariableEditRequest).Execute()

Edit an environment variable belonging to the environment



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
	environmentVariableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment Variable ID
	environmentVariableEditRequest := *openapiclient.NewEnvironmentVariableEditRequest("Key_example") // EnvironmentVariableEditRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentVariableAPI.EditEnvironmentEnvironmentVariable(context.Background(), environmentId, environmentVariableId).EnvironmentVariableEditRequest(environmentVariableEditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableAPI.EditEnvironmentEnvironmentVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditEnvironmentEnvironmentVariable`: EnvironmentVariable
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableAPI.EditEnvironmentEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 
**environmentVariableId** | **string** | Environment Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditEnvironmentEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environmentVariableEditRequest** | [**EnvironmentVariableEditRequest**](EnvironmentVariableEditRequest.md) |  | 

### Return type

[**EnvironmentVariable**](EnvironmentVariable.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironmentEnvironmentVariable

> EnvironmentVariableResponseList ListEnvironmentEnvironmentVariable(ctx, environmentId).Execute()

List environment variables

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
	resp, r, err := apiClient.EnvironmentVariableAPI.ListEnvironmentEnvironmentVariable(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableAPI.ListEnvironmentEnvironmentVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnvironmentEnvironmentVariable`: EnvironmentVariableResponseList
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableAPI.ListEnvironmentEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentVariableResponseList**](EnvironmentVariableResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

