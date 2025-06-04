# \TerraformMainCallsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTerraform**](TerraformMainCallsAPI.md#DeleteTerraform) | **Delete** /terraform/{terraformId} | Delete Terraform
[**EditTerraform**](TerraformMainCallsAPI.md#EditTerraform) | **Put** /terraform/{terraformId} | Edit Terraform
[**GetTerraform**](TerraformMainCallsAPI.md#GetTerraform) | **Get** /terraform/{terraformId} | Get terraform by ID
[**ListTerraformCommit**](TerraformMainCallsAPI.md#ListTerraformCommit) | **Get** /terraform/{terraformId}/commit | List last commits



## DeleteTerraform

> DeleteTerraform(ctx, terraformId).Execute()

Delete Terraform

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
	terraformId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Terraform ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TerraformMainCallsAPI.DeleteTerraform(context.Background(), terraformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformMainCallsAPI.DeleteTerraform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTerraformRequest struct via the builder pattern


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


## EditTerraform

> TerraformResponse EditTerraform(ctx, terraformId).TerraformRequest(terraformRequest).Execute()

Edit Terraform

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
	terraformId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Terraform ID
	terraformRequest := *openapiclient.NewTerraformRequest("Name_example", "Description_example", false, false, openapiclient.TerraformRequest_terraform_files_source{TerraformRequestTerraformFilesSourceOneOf: openapiclient.NewTerraformRequestTerraformFilesSourceOneOf()}, *openapiclient.NewTerraformVariablesSourceRequest([]string{"TfVarFilePaths_example"}, [][]string{[]string{"TfVars_example"}}), "Provider_example", *openapiclient.NewTerraformProviderVersion(false), *openapiclient.NewTerraformRequestJobResources(int32(123), int32(123), int32(123))) // TerraformRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerraformMainCallsAPI.EditTerraform(context.Background(), terraformId).TerraformRequest(terraformRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformMainCallsAPI.EditTerraform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditTerraform`: TerraformResponse
	fmt.Fprintf(os.Stdout, "Response from `TerraformMainCallsAPI.EditTerraform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTerraformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terraformRequest** | [**TerraformRequest**](TerraformRequest.md) |  | 

### Return type

[**TerraformResponse**](TerraformResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTerraform

> TerraformResponse GetTerraform(ctx, terraformId).Execute()

Get terraform by ID

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
	terraformId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Terraform ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerraformMainCallsAPI.GetTerraform(context.Background(), terraformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformMainCallsAPI.GetTerraform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTerraform`: TerraformResponse
	fmt.Fprintf(os.Stdout, "Response from `TerraformMainCallsAPI.GetTerraform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTerraformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TerraformResponse**](TerraformResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTerraformCommit

> CommitResponseList ListTerraformCommit(ctx, terraformId).Execute()

List last commits



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
	terraformId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Terraform ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerraformMainCallsAPI.ListTerraformCommit(context.Background(), terraformId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerraformMainCallsAPI.ListTerraformCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTerraformCommit`: CommitResponseList
	fmt.Fprintf(os.Stdout, "Response from `TerraformMainCallsAPI.ListTerraformCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terraformId** | **string** | Terraform ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTerraformCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommitResponseList**](CommitResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

