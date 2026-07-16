# \AgenticWorkflowsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgenticWorkflow**](AgenticWorkflowsAPI.md#CreateAgenticWorkflow) | **Post** /environment/{environmentId}/agenticWorkflow | Create an agentic workflow
[**DeleteAgenticWorkflow**](AgenticWorkflowsAPI.md#DeleteAgenticWorkflow) | **Delete** /agenticWorkflow/{agenticWorkflowId} | Delete an agentic workflow
[**EditAgenticWorkflow**](AgenticWorkflowsAPI.md#EditAgenticWorkflow) | **Put** /agenticWorkflow/{agenticWorkflowId} | Edit an agentic workflow
[**GetAgenticWorkflow**](AgenticWorkflowsAPI.md#GetAgenticWorkflow) | **Get** /agenticWorkflow/{agenticWorkflowId} | Get an agentic workflow
[**ListAgenticWorkflows**](AgenticWorkflowsAPI.md#ListAgenticWorkflows) | **Get** /environment/{environmentId}/agenticWorkflow | List agentic workflows



## CreateAgenticWorkflow

> AgenticWorkflowResponse CreateAgenticWorkflow(ctx, environmentId).AgenticWorkflowRequest(agenticWorkflowRequest).Execute()

Create an agentic workflow

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
	agenticWorkflowRequest := *openapiclient.NewAgenticWorkflowRequest("Name_example") // AgenticWorkflowRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgenticWorkflowsAPI.CreateAgenticWorkflow(context.Background(), environmentId).AgenticWorkflowRequest(agenticWorkflowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgenticWorkflowsAPI.CreateAgenticWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAgenticWorkflow`: AgenticWorkflowResponse
	fmt.Fprintf(os.Stdout, "Response from `AgenticWorkflowsAPI.CreateAgenticWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgenticWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agenticWorkflowRequest** | [**AgenticWorkflowRequest**](AgenticWorkflowRequest.md) |  | 

### Return type

[**AgenticWorkflowResponse**](AgenticWorkflowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgenticWorkflow

> DeleteAgenticWorkflow(ctx, agenticWorkflowId).Execute()

Delete an agentic workflow

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
	agenticWorkflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgenticWorkflowsAPI.DeleteAgenticWorkflow(context.Background(), agenticWorkflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgenticWorkflowsAPI.DeleteAgenticWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agenticWorkflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgenticWorkflowRequest struct via the builder pattern


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


## EditAgenticWorkflow

> AgenticWorkflowResponse EditAgenticWorkflow(ctx, agenticWorkflowId).AgenticWorkflowRequest(agenticWorkflowRequest).Execute()

Edit an agentic workflow

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
	agenticWorkflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	agenticWorkflowRequest := *openapiclient.NewAgenticWorkflowRequest("Name_example") // AgenticWorkflowRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgenticWorkflowsAPI.EditAgenticWorkflow(context.Background(), agenticWorkflowId).AgenticWorkflowRequest(agenticWorkflowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgenticWorkflowsAPI.EditAgenticWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAgenticWorkflow`: AgenticWorkflowResponse
	fmt.Fprintf(os.Stdout, "Response from `AgenticWorkflowsAPI.EditAgenticWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agenticWorkflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAgenticWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agenticWorkflowRequest** | [**AgenticWorkflowRequest**](AgenticWorkflowRequest.md) |  | 

### Return type

[**AgenticWorkflowResponse**](AgenticWorkflowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgenticWorkflow

> AgenticWorkflowResponse GetAgenticWorkflow(ctx, agenticWorkflowId).Execute()

Get an agentic workflow

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
	agenticWorkflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgenticWorkflowsAPI.GetAgenticWorkflow(context.Background(), agenticWorkflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgenticWorkflowsAPI.GetAgenticWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgenticWorkflow`: AgenticWorkflowResponse
	fmt.Fprintf(os.Stdout, "Response from `AgenticWorkflowsAPI.GetAgenticWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agenticWorkflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgenticWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgenticWorkflowResponse**](AgenticWorkflowResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgenticWorkflows

> AgenticWorkflowResponseList ListAgenticWorkflows(ctx, environmentId).Execute()

List agentic workflows

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
	resp, r, err := apiClient.AgenticWorkflowsAPI.ListAgenticWorkflows(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgenticWorkflowsAPI.ListAgenticWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgenticWorkflows`: AgenticWorkflowResponseList
	fmt.Fprintf(os.Stdout, "Response from `AgenticWorkflowsAPI.ListAgenticWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAgenticWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgenticWorkflowResponseList**](AgenticWorkflowResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

