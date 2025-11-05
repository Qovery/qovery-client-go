# \AlertRulesAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertRule**](AlertRulesAPI.md#CreateAlertRule) | **Post** /alert-rules | Create alert rule
[**DeleteAlertRule**](AlertRulesAPI.md#DeleteAlertRule) | **Delete** /alert-rules/{alertRuleId} | Delete alert rule
[**EditAlertRule**](AlertRulesAPI.md#EditAlertRule) | **Put** /alert-rules/{alertRuleId} | Update alert rule
[**GetAlertRule**](AlertRulesAPI.md#GetAlertRule) | **Get** /alert-rules/{alertRuleId} | Get alert rule
[**GetAlertRules**](AlertRulesAPI.md#GetAlertRules) | **Get** /organization/{organizationId}/alert-rules | List alert rules



## CreateAlertRule

> AlertRuleResponse CreateAlertRule(ctx).AlertRuleCreationRequest(alertRuleCreationRequest).Execute()

Create alert rule



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
	alertRuleCreationRequest := *openapiclient.NewAlertRuleCreationRequest("OrganizationId_example", "ClusterId_example", "Name_example", "Description_example", "PromqlExpr_example", "ForDuration_example", openapiclient.AlertSeverity("WARNING"), *openapiclient.NewAlertPresentation(), false, []string{"AlertReceiverIds_example"}, *openapiclient.NewAlertTarget(openapiclient.AlertTargetType(" CLUSTER"), "TargetId_example")) // AlertRuleCreationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.CreateAlertRule(context.Background()).AlertRuleCreationRequest(alertRuleCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.CreateAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlertRule`: AlertRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.CreateAlertRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertRuleCreationRequest** | [**AlertRuleCreationRequest**](AlertRuleCreationRequest.md) |  | 

### Return type

[**AlertRuleResponse**](AlertRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertRule

> DeleteAlertRule(ctx, alertRuleId).Execute()

Delete alert rule



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
	alertRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Alert Rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertRulesAPI.DeleteAlertRule(context.Background(), alertRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.DeleteAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertRuleId** | **string** | Alert Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertRuleRequest struct via the builder pattern


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


## EditAlertRule

> AlertRuleResponse EditAlertRule(ctx, alertRuleId).AlertRuleEditRequest(alertRuleEditRequest).Execute()

Update alert rule



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
	alertRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Alert Rule ID
	alertRuleEditRequest := *openapiclient.NewAlertRuleEditRequest("Name_example", "Description_example", "PromqlExpr_example", "ForDuration_example", openapiclient.AlertSeverity("WARNING"), false, []string{"AlertReceiverIds_example"}, *openapiclient.NewAlertPresentation()) // AlertRuleEditRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.EditAlertRule(context.Background(), alertRuleId).AlertRuleEditRequest(alertRuleEditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.EditAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAlertRule`: AlertRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.EditAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertRuleId** | **string** | Alert Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertRuleEditRequest** | [**AlertRuleEditRequest**](AlertRuleEditRequest.md) |  | 

### Return type

[**AlertRuleResponse**](AlertRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertRule

> AlertRuleResponse GetAlertRule(ctx, alertRuleId).Execute()

Get alert rule



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
	alertRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Alert Rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertRulesAPI.GetAlertRule(context.Background(), alertRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.GetAlertRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertRule`: AlertRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.GetAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertRuleId** | **string** | Alert Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertRuleResponse**](AlertRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertRules

> AlertRuleList GetAlertRules(ctx, organizationId).Execute()

List alert rules



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
	resp, r, err := apiClient.AlertRulesAPI.GetAlertRules(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesAPI.GetAlertRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertRules`: AlertRuleList
	fmt.Fprintf(os.Stdout, "Response from `AlertRulesAPI.GetAlertRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertRuleList**](AlertRuleList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

