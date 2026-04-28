# \ArgoCDAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckArgoCdConnection**](ArgoCDAPI.md#CheckArgoCdConnection) | **Post** /cluster/{clusterId}/argoCdConfig/check | Check ArgoCD connection
[**DeleteArgoCdCredentials**](ArgoCDAPI.md#DeleteArgoCdCredentials) | **Delete** /cluster/{clusterId}/argoCdConfig | Delete ArgoCD credentials for a cluster
[**GetArgoCdApp**](ArgoCDAPI.md#GetArgoCdApp) | **Get** /argocdApp/{argocdAppId} | Get ArgoCD app by ID
[**GetArgoCdCredentials**](ArgoCDAPI.md#GetArgoCdCredentials) | **Get** /cluster/{clusterId}/argoCdConfig | Get ArgoCD credentials for a cluster
[**SaveArgoCdCredentials**](ArgoCDAPI.md#SaveArgoCdCredentials) | **Post** /cluster/{clusterId}/argoCdConfig | Save ArgoCD credentials for a cluster



## CheckArgoCdConnection

> ArgoCdConnectionCheckResponse CheckArgoCdConnection(ctx, clusterId).ArgoCdCredentialsRequest(argoCdCredentialsRequest).Execute()

Check ArgoCD connection



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	argoCdCredentialsRequest := *openapiclient.NewArgoCdCredentialsRequest("https://argocd.example.com", "ArgocdToken_example") // ArgoCdCredentialsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArgoCDAPI.CheckArgoCdConnection(context.Background(), clusterId).ArgoCdCredentialsRequest(argoCdCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArgoCDAPI.CheckArgoCdConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckArgoCdConnection`: ArgoCdConnectionCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `ArgoCDAPI.CheckArgoCdConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckArgoCdConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **argoCdCredentialsRequest** | [**ArgoCdCredentialsRequest**](ArgoCdCredentialsRequest.md) |  | 

### Return type

[**ArgoCdConnectionCheckResponse**](ArgoCdConnectionCheckResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArgoCdCredentials

> DeleteArgoCdCredentials(ctx, clusterId).Execute()

Delete ArgoCD credentials for a cluster



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArgoCDAPI.DeleteArgoCdCredentials(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArgoCDAPI.DeleteArgoCdCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArgoCdCredentialsRequest struct via the builder pattern


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


## GetArgoCdApp

> ArgocdAppResponse GetArgoCdApp(ctx, argocdAppId).Execute()

Get ArgoCD app by ID

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
	argocdAppId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ArgoCD App ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArgoCDAPI.GetArgoCdApp(context.Background(), argocdAppId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArgoCDAPI.GetArgoCdApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArgoCdApp`: ArgocdAppResponse
	fmt.Fprintf(os.Stdout, "Response from `ArgoCDAPI.GetArgoCdApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**argocdAppId** | **string** | ArgoCD App ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArgoCdAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArgocdAppResponse**](ArgocdAppResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArgoCdCredentials

> ArgoCdCredentialsResponse GetArgoCdCredentials(ctx, clusterId).Execute()

Get ArgoCD credentials for a cluster



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArgoCDAPI.GetArgoCdCredentials(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArgoCDAPI.GetArgoCdCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArgoCdCredentials`: ArgoCdCredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `ArgoCDAPI.GetArgoCdCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArgoCdCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArgoCdCredentialsResponse**](ArgoCdCredentialsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveArgoCdCredentials

> ArgoCdCredentialsResponse SaveArgoCdCredentials(ctx, clusterId).ArgoCdCredentialsRequest(argoCdCredentialsRequest).Execute()

Save ArgoCD credentials for a cluster



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	argoCdCredentialsRequest := *openapiclient.NewArgoCdCredentialsRequest("https://argocd.example.com", "ArgocdToken_example") // ArgoCdCredentialsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArgoCDAPI.SaveArgoCdCredentials(context.Background(), clusterId).ArgoCdCredentialsRequest(argoCdCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArgoCDAPI.SaveArgoCdCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveArgoCdCredentials`: ArgoCdCredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `ArgoCDAPI.SaveArgoCdCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveArgoCdCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **argoCdCredentialsRequest** | [**ArgoCdCredentialsRequest**](ArgoCdCredentialsRequest.md) |  | 

### Return type

[**ArgoCdCredentialsResponse**](ArgoCdCredentialsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

