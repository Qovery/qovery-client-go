# \KedaTriggerAuthenticationAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKedaTriggerAuthentication**](KedaTriggerAuthenticationAPI.md#CreateKedaTriggerAuthentication) | **Post** /organization/{organizationId}/keda-trigger-authentications | create Keda Trigger Authentications
[**DeleteKedaTriggerAuthentication**](KedaTriggerAuthenticationAPI.md#DeleteKedaTriggerAuthentication) | **Delete** /organization/{organizationId}/keda-trigger-authentications/{triggerAuthenticationId} | Delete a KEDA trigger authentication
[**GetKedaTriggerAuthentication**](KedaTriggerAuthenticationAPI.md#GetKedaTriggerAuthentication) | **Get** /organization/{organizationId}/keda-trigger-authentications/{triggerAuthenticationId} | Get a KEDA trigger authentication
[**ListKedaTriggerAuthentications**](KedaTriggerAuthenticationAPI.md#ListKedaTriggerAuthentications) | **Get** /organization/{organizationId}/keda-trigger-authentications | list Keda TriggerAuthentications
[**UpdateKedaTriggerAuthentication**](KedaTriggerAuthenticationAPI.md#UpdateKedaTriggerAuthentication) | **Put** /organization/{organizationId}/keda-trigger-authentications/{triggerAuthenticationId} | Update a KEDA trigger authentication



## CreateKedaTriggerAuthentication

> KedaTriggerAuthenticationResponse CreateKedaTriggerAuthentication(ctx, organizationId).KedaTriggerAuthenticationRequest(kedaTriggerAuthenticationRequest).Execute()

create Keda Trigger Authentications

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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	kedaTriggerAuthenticationRequest := *openapiclient.NewKedaTriggerAuthenticationRequest("Name_example") // KedaTriggerAuthenticationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KedaTriggerAuthenticationAPI.CreateKedaTriggerAuthentication(context.Background(), organizationId).KedaTriggerAuthenticationRequest(kedaTriggerAuthenticationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KedaTriggerAuthenticationAPI.CreateKedaTriggerAuthentication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKedaTriggerAuthentication`: KedaTriggerAuthenticationResponse
	fmt.Fprintf(os.Stdout, "Response from `KedaTriggerAuthenticationAPI.CreateKedaTriggerAuthentication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKedaTriggerAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kedaTriggerAuthenticationRequest** | [**KedaTriggerAuthenticationRequest**](KedaTriggerAuthenticationRequest.md) |  | 

### Return type

[**KedaTriggerAuthenticationResponse**](KedaTriggerAuthenticationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKedaTriggerAuthentication

> DeleteKedaTriggerAuthentication(ctx, organizationId, triggerAuthenticationId).Execute()

Delete a KEDA trigger authentication

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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	triggerAuthenticationId := "triggerAuthenticationId_example" // string | KEDA triggerAuthentication ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KedaTriggerAuthenticationAPI.DeleteKedaTriggerAuthentication(context.Background(), organizationId, triggerAuthenticationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KedaTriggerAuthenticationAPI.DeleteKedaTriggerAuthentication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**triggerAuthenticationId** | **string** | KEDA triggerAuthentication ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKedaTriggerAuthenticationRequest struct via the builder pattern


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


## GetKedaTriggerAuthentication

> KedaTriggerAuthenticationResponse GetKedaTriggerAuthentication(ctx, organizationId, triggerAuthenticationId).Execute()

Get a KEDA trigger authentication

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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	triggerAuthenticationId := "triggerAuthenticationId_example" // string | KEDA triggerAuthentication ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KedaTriggerAuthenticationAPI.GetKedaTriggerAuthentication(context.Background(), organizationId, triggerAuthenticationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KedaTriggerAuthenticationAPI.GetKedaTriggerAuthentication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKedaTriggerAuthentication`: KedaTriggerAuthenticationResponse
	fmt.Fprintf(os.Stdout, "Response from `KedaTriggerAuthenticationAPI.GetKedaTriggerAuthentication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**triggerAuthenticationId** | **string** | KEDA triggerAuthentication ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKedaTriggerAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KedaTriggerAuthenticationResponse**](KedaTriggerAuthenticationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKedaTriggerAuthentications

> KedaTriggerAuthenticationResponseList ListKedaTriggerAuthentications(ctx, organizationId).Execute()

list Keda TriggerAuthentications

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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KedaTriggerAuthenticationAPI.ListKedaTriggerAuthentications(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KedaTriggerAuthenticationAPI.ListKedaTriggerAuthentications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKedaTriggerAuthentications`: KedaTriggerAuthenticationResponseList
	fmt.Fprintf(os.Stdout, "Response from `KedaTriggerAuthenticationAPI.ListKedaTriggerAuthentications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKedaTriggerAuthenticationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KedaTriggerAuthenticationResponseList**](KedaTriggerAuthenticationResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKedaTriggerAuthentication

> KedaTriggerAuthenticationResponse UpdateKedaTriggerAuthentication(ctx, organizationId, triggerAuthenticationId).KedaTriggerAuthenticationRequest(kedaTriggerAuthenticationRequest).Execute()

Update a KEDA trigger authentication

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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	triggerAuthenticationId := "triggerAuthenticationId_example" // string | KEDA triggerAuthentication ID
	kedaTriggerAuthenticationRequest := *openapiclient.NewKedaTriggerAuthenticationRequest("Name_example") // KedaTriggerAuthenticationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KedaTriggerAuthenticationAPI.UpdateKedaTriggerAuthentication(context.Background(), organizationId, triggerAuthenticationId).KedaTriggerAuthenticationRequest(kedaTriggerAuthenticationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KedaTriggerAuthenticationAPI.UpdateKedaTriggerAuthentication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKedaTriggerAuthentication`: KedaTriggerAuthenticationResponse
	fmt.Fprintf(os.Stdout, "Response from `KedaTriggerAuthenticationAPI.UpdateKedaTriggerAuthentication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**triggerAuthenticationId** | **string** | KEDA triggerAuthentication ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKedaTriggerAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kedaTriggerAuthenticationRequest** | [**KedaTriggerAuthenticationRequest**](KedaTriggerAuthenticationRequest.md) |  | 

### Return type

[**KedaTriggerAuthenticationResponse**](KedaTriggerAuthenticationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

