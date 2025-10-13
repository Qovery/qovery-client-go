# \AlertReceiversAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertReceiver**](AlertReceiversAPI.md#CreateAlertReceiver) | **Post** /api/alert-receivers | Create alert receiver
[**DeleteAlertReceiver**](AlertReceiversAPI.md#DeleteAlertReceiver) | **Delete** /api/alert-receivers/{alertReceiverId} | Delete alert receiver
[**EditAlertReceiver**](AlertReceiversAPI.md#EditAlertReceiver) | **Put** /api/alert-receivers/{alertReceiverId} | Update alert receiver
[**GetAlertReceiver**](AlertReceiversAPI.md#GetAlertReceiver) | **Get** /api/alert-receivers/{alertReceiverId} | Get alert receiver
[**GetAlertReceivers**](AlertReceiversAPI.md#GetAlertReceivers) | **Get** /api/organization/{organizationId}/alert-receivers | List alert receivers



## CreateAlertReceiver

> AlertReceiverResponse CreateAlertReceiver(ctx).AlertReceiverCreationRequest(alertReceiverCreationRequest).Execute()

Create alert receiver



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
	alertReceiverCreationRequest := *openapiclient.NewAlertReceiverCreationRequest("OrganizationId_example", "Name_example", "Description_example", openapiclient.AlertReceiverType("SLACK"), false) // AlertReceiverCreationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertReceiversAPI.CreateAlertReceiver(context.Background()).AlertReceiverCreationRequest(alertReceiverCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertReceiversAPI.CreateAlertReceiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlertReceiver`: AlertReceiverResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertReceiversAPI.CreateAlertReceiver`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertReceiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertReceiverCreationRequest** | [**AlertReceiverCreationRequest**](AlertReceiverCreationRequest.md) |  | 

### Return type

[**AlertReceiverResponse**](AlertReceiverResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertReceiver

> DeleteAlertReceiver(ctx, alertReceiverId).Execute()

Delete alert receiver



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
	alertReceiverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Alert Receiver ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertReceiversAPI.DeleteAlertReceiver(context.Background(), alertReceiverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertReceiversAPI.DeleteAlertReceiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertReceiverId** | **string** | Alert Receiver ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertReceiverRequest struct via the builder pattern


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


## EditAlertReceiver

> AlertReceiverResponse EditAlertReceiver(ctx, alertReceiverId).AlertReceiverEditRequest(alertReceiverEditRequest).Execute()

Update alert receiver



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
	alertReceiverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Alert Receiver ID
	alertReceiverEditRequest := *openapiclient.NewAlertReceiverEditRequest("Name_example", "Description_example", openapiclient.AlertReceiverType("SLACK"), false) // AlertReceiverEditRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertReceiversAPI.EditAlertReceiver(context.Background(), alertReceiverId).AlertReceiverEditRequest(alertReceiverEditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertReceiversAPI.EditAlertReceiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAlertReceiver`: AlertReceiverResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertReceiversAPI.EditAlertReceiver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertReceiverId** | **string** | Alert Receiver ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAlertReceiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertReceiverEditRequest** | [**AlertReceiverEditRequest**](AlertReceiverEditRequest.md) |  | 

### Return type

[**AlertReceiverResponse**](AlertReceiverResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertReceiver

> AlertReceiverResponse GetAlertReceiver(ctx, alertReceiverId).Execute()

Get alert receiver



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
	alertReceiverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Alert Receiver ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertReceiversAPI.GetAlertReceiver(context.Background(), alertReceiverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertReceiversAPI.GetAlertReceiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertReceiver`: AlertReceiverResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertReceiversAPI.GetAlertReceiver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertReceiverId** | **string** | Alert Receiver ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertReceiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertReceiverResponse**](AlertReceiverResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertReceivers

> AlertReceiverList GetAlertReceivers(ctx, organizationId).Execute()

List alert receivers



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
	resp, r, err := apiClient.AlertReceiversAPI.GetAlertReceivers(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertReceiversAPI.GetAlertReceivers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertReceivers`: AlertReceiverList
	fmt.Fprintf(os.Stdout, "Response from `AlertReceiversAPI.GetAlertReceivers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertReceiversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertReceiverList**](AlertReceiverList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

