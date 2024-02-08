# \OrganizationWebhookAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationWebhook**](OrganizationWebhookAPI.md#CreateOrganizationWebhook) | **Post** /organization/{organizationId}/webhook | Create an organization webhook
[**DeleteOrganizationWebhook**](OrganizationWebhookAPI.md#DeleteOrganizationWebhook) | **Delete** /organization/{organizationId}/webhook/{webhookId} | Delete organization webhook
[**EditOrganizationWebhook**](OrganizationWebhookAPI.md#EditOrganizationWebhook) | **Put** /organization/{organizationId}/webhook/{webhookId} | Edit an organization webhook
[**GetOrganizationWebhook**](OrganizationWebhookAPI.md#GetOrganizationWebhook) | **Get** /organization/{organizationId}/webhook/{webhookId} | Get an Organization webhook
[**ListOrganizationWebHooks**](OrganizationWebhookAPI.md#ListOrganizationWebHooks) | **Get** /organization/{organizationId}/webhook | List organization webhooks



## CreateOrganizationWebhook

> OrganizationWebhookCreateResponse CreateOrganizationWebhook(ctx, organizationId).OrganizationWebhookCreateRequest(organizationWebhookCreateRequest).Execute()

Create an organization webhook



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
    organizationWebhookCreateRequest := *openapiclient.NewOrganizationWebhookCreateRequest(openapiclient.OrganizationWebhookKindEnum("STANDARD"), "TargetUrl_example", []openapiclient.OrganizationWebhookEventEnum{openapiclient.OrganizationWebhookEventEnum("DEPLOYMENT_STARTED")}) // OrganizationWebhookCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationWebhookAPI.CreateOrganizationWebhook(context.Background(), organizationId).OrganizationWebhookCreateRequest(organizationWebhookCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationWebhookAPI.CreateOrganizationWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationWebhook`: OrganizationWebhookCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationWebhookAPI.CreateOrganizationWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationWebhookCreateRequest** | [**OrganizationWebhookCreateRequest**](OrganizationWebhookCreateRequest.md) |  | 

### Return type

[**OrganizationWebhookCreateResponse**](OrganizationWebhookCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationWebhook

> DeleteOrganizationWebhook(ctx, organizationId, webhookId).Execute()

Delete organization webhook



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
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrganizationWebhookAPI.DeleteOrganizationWebhook(context.Background(), organizationId, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationWebhookAPI.DeleteOrganizationWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**webhookId** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationWebhookRequest struct via the builder pattern


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


## EditOrganizationWebhook

> OrganizationWebhookCreateResponse EditOrganizationWebhook(ctx, organizationId, webhookId).OrganizationWebhookCreateRequest(organizationWebhookCreateRequest).Execute()

Edit an organization webhook



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
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook ID
    organizationWebhookCreateRequest := *openapiclient.NewOrganizationWebhookCreateRequest(openapiclient.OrganizationWebhookKindEnum("STANDARD"), "TargetUrl_example", []openapiclient.OrganizationWebhookEventEnum{openapiclient.OrganizationWebhookEventEnum("DEPLOYMENT_STARTED")}) // OrganizationWebhookCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationWebhookAPI.EditOrganizationWebhook(context.Background(), organizationId, webhookId).OrganizationWebhookCreateRequest(organizationWebhookCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationWebhookAPI.EditOrganizationWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditOrganizationWebhook`: OrganizationWebhookCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationWebhookAPI.EditOrganizationWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**webhookId** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganizationWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **organizationWebhookCreateRequest** | [**OrganizationWebhookCreateRequest**](OrganizationWebhookCreateRequest.md) |  | 

### Return type

[**OrganizationWebhookCreateResponse**](OrganizationWebhookCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWebhook

> OrganizationWebhookResponse GetOrganizationWebhook(ctx, organizationId, webhookId).Execute()

Get an Organization webhook



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
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationWebhookAPI.GetOrganizationWebhook(context.Background(), organizationId, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationWebhookAPI.GetOrganizationWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWebhook`: OrganizationWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationWebhookAPI.GetOrganizationWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**webhookId** | **string** | Webhook ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationWebhookResponse**](OrganizationWebhookResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationWebHooks

> OrganizationWebhookResponseList ListOrganizationWebHooks(ctx, organizationId).Execute()

List organization webhooks



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
    resp, r, err := apiClient.OrganizationWebhookAPI.ListOrganizationWebHooks(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationWebhookAPI.ListOrganizationWebHooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizationWebHooks`: OrganizationWebhookResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationWebhookAPI.ListOrganizationWebHooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationWebHooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationWebhookResponseList**](OrganizationWebhookResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

