# \JobSecretAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJobSecret**](JobSecretAPI.md#CreateJobSecret) | **Post** /job/{jobId}/secret | Add a secret to the job
[**CreateJobSecretAlias**](JobSecretAPI.md#CreateJobSecretAlias) | **Post** /job/{jobId}/secret/{secretId}/alias | Create a secret alias at the job level
[**CreateJobSecretOverride**](JobSecretAPI.md#CreateJobSecretOverride) | **Post** /job/{jobId}/secret/{secretId}/override | Create a secret override at the job level
[**DeleteJobSecret**](JobSecretAPI.md#DeleteJobSecret) | **Delete** /job/{jobId}/secret/{secretId} | Delete a secret from an job
[**EditJobSecret**](JobSecretAPI.md#EditJobSecret) | **Put** /job/{jobId}/secret/{secretId} | Edit a secret belonging to the job
[**ListJobSecrets**](JobSecretAPI.md#ListJobSecrets) | **Get** /job/{jobId}/secret | List job secrets



## CreateJobSecret

> Secret CreateJobSecret(ctx, jobId).SecretRequest(secretRequest).Execute()

Add a secret to the job



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
	secretRequest := *openapiclient.NewSecretRequest("Key_example") // SecretRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobSecretAPI.CreateJobSecret(context.Background(), jobId).SecretRequest(secretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobSecretAPI.CreateJobSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJobSecret`: Secret
	fmt.Fprintf(os.Stdout, "Response from `JobSecretAPI.CreateJobSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secretRequest** | [**SecretRequest**](SecretRequest.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJobSecretAlias

> Secret CreateJobSecretAlias(ctx, jobId, secretId).Key(key).Execute()

Create a secret alias at the job level



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
	secretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Secret ID
	key := *openapiclient.NewKey("Key_example") // Key |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobSecretAPI.CreateJobSecretAlias(context.Background(), jobId, secretId).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobSecretAPI.CreateJobSecretAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJobSecretAlias`: Secret
	fmt.Fprintf(os.Stdout, "Response from `JobSecretAPI.CreateJobSecretAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 
**secretId** | **string** | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobSecretAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | [**Key**](Key.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJobSecretOverride

> Secret CreateJobSecretOverride(ctx, jobId, secretId).Value(value).Execute()

Create a secret override at the job level



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
	secretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Secret ID
	value := *openapiclient.NewValue() // Value |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobSecretAPI.CreateJobSecretOverride(context.Background(), jobId, secretId).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobSecretAPI.CreateJobSecretOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJobSecretOverride`: Secret
	fmt.Fprintf(os.Stdout, "Response from `JobSecretAPI.CreateJobSecretOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 
**secretId** | **string** | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobSecretOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | [**Value**](Value.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJobSecret

> DeleteJobSecret(ctx, jobId, secretId).Execute()

Delete a secret from an job



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
	secretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Secret ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobSecretAPI.DeleteJobSecret(context.Background(), jobId, secretId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobSecretAPI.DeleteJobSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 
**secretId** | **string** | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobSecretRequest struct via the builder pattern


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


## EditJobSecret

> Secret EditJobSecret(ctx, jobId, secretId).SecretEditRequest(secretEditRequest).Execute()

Edit a secret belonging to the job



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
	secretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Secret ID
	secretEditRequest := *openapiclient.NewSecretEditRequest("Key_example") // SecretEditRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobSecretAPI.EditJobSecret(context.Background(), jobId, secretId).SecretEditRequest(secretEditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobSecretAPI.EditJobSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditJobSecret`: Secret
	fmt.Fprintf(os.Stdout, "Response from `JobSecretAPI.EditJobSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 
**secretId** | **string** | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditJobSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **secretEditRequest** | [**SecretEditRequest**](SecretEditRequest.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJobSecrets

> SecretResponseList ListJobSecrets(ctx, jobId).Execute()

List job secrets



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobSecretAPI.ListJobSecrets(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobSecretAPI.ListJobSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListJobSecrets`: SecretResponseList
	fmt.Fprintf(os.Stdout, "Response from `JobSecretAPI.ListJobSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListJobSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretResponseList**](SecretResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

