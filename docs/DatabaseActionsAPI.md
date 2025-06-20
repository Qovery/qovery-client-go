# \DatabaseActionsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployDatabase**](DatabaseActionsAPI.md#DeployDatabase) | **Post** /database/{databaseId}/deploy | Deploy database 
[**RebootDatabase**](DatabaseActionsAPI.md#RebootDatabase) | **Post** /database/{databaseId}/restart-service | Retart database
[**RedeployDatabase**](DatabaseActionsAPI.md#RedeployDatabase) | **Post** /database/{databaseId}/redeploy | Redeploy database
[**StopDatabase**](DatabaseActionsAPI.md#StopDatabase) | **Post** /database/{databaseId}/stop | Stop database
[**UninstallDatabase**](DatabaseActionsAPI.md#UninstallDatabase) | **Post** /database/{databaseId}/uninstall | Uninstall database



## DeployDatabase

> Status DeployDatabase(ctx, databaseId).Execute()

Deploy database 

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
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabaseActionsAPI.DeployDatabase(context.Background(), databaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseActionsAPI.DeployDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployDatabase`: Status
	fmt.Fprintf(os.Stdout, "Response from `DatabaseActionsAPI.DeployDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootDatabase

> Status RebootDatabase(ctx, databaseId).Execute()

Retart database

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
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabaseActionsAPI.RebootDatabase(context.Background(), databaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseActionsAPI.RebootDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootDatabase`: Status
	fmt.Fprintf(os.Stdout, "Response from `DatabaseActionsAPI.RebootDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeployDatabase

> Status RedeployDatabase(ctx, databaseId).Execute()

Redeploy database

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
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabaseActionsAPI.RedeployDatabase(context.Background(), databaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseActionsAPI.RedeployDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeployDatabase`: Status
	fmt.Fprintf(os.Stdout, "Response from `DatabaseActionsAPI.RedeployDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopDatabase

> Status StopDatabase(ctx, databaseId).Execute()

Stop database

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
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabaseActionsAPI.StopDatabase(context.Background(), databaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseActionsAPI.StopDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopDatabase`: Status
	fmt.Fprintf(os.Stdout, "Response from `DatabaseActionsAPI.StopDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallDatabase

> Status UninstallDatabase(ctx, databaseId).Body(body).Execute()

Uninstall database

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
	databaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Database ID
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabaseActionsAPI.UninstallDatabase(context.Background(), databaseId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseActionsAPI.UninstallDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UninstallDatabase`: Status
	fmt.Fprintf(os.Stdout, "Response from `DatabaseActionsAPI.UninstallDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseId** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

