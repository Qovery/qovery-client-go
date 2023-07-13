# \ApplicationsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneApplication**](ApplicationsApi.md#CloneApplication) | **Post** /application/{applicationId}/clone | Clone application
[**CreateApplication**](ApplicationsApi.md#CreateApplication) | **Post** /environment/{environmentId}/application | Create an application
[**GetDefaultApplicationAdvancedSettings**](ApplicationsApi.md#GetDefaultApplicationAdvancedSettings) | **Get** /defaultApplicationAdvancedSettings | List default application advanced settings
[**GetEnvironmentApplicationCurrentScale**](ApplicationsApi.md#GetEnvironmentApplicationCurrentScale) | **Get** /environment/{environmentId}/application/currentScale | List current scaling information for each application
[**GetEnvironmentApplicationCurrentStorage**](ApplicationsApi.md#GetEnvironmentApplicationCurrentStorage) | **Get** /environment/{environmentId}/application/currentStorage | List current storage disk usage for each application
[**GetEnvironmentApplicationStatus**](ApplicationsApi.md#GetEnvironmentApplicationStatus) | **Get** /environment/{environmentId}/application/status | List all environment applications statuses
[**GetEnvironmentApplicationSupportedLanguages**](ApplicationsApi.md#GetEnvironmentApplicationSupportedLanguages) | **Get** /environment/{environmentId}/application/supportedLanguage | List supported languages
[**ListApplication**](ApplicationsApi.md#ListApplication) | **Get** /environment/{environmentId}/application | List applications



## CloneApplication

> Application CloneApplication(ctx, applicationId).CloneApplicationRequest(cloneApplicationRequest).Execute()

Clone application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    cloneApplicationRequest := *openapiclient.NewCloneApplicationRequest("Name_example", "EnvironmentId_example") // CloneApplicationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.CloneApplication(context.Background(), applicationId).CloneApplicationRequest(cloneApplicationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.CloneApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.CloneApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneApplicationRequest** | [**CloneApplicationRequest**](CloneApplicationRequest.md) |  | 

### Return type

[**Application**](Application.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplication

> Application CreateApplication(ctx, environmentId).ApplicationRequest(applicationRequest).Execute()

Create an application

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
    applicationRequest := *openapiclient.NewApplicationRequest("Name_example", *openapiclient.NewApplicationGitRepositoryRequest("https://github.com/Qovery/simple-node-app")) // ApplicationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.CreateApplication(context.Background(), environmentId).ApplicationRequest(applicationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.CreateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplication`: Application
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.CreateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationRequest** | [**ApplicationRequest**](ApplicationRequest.md) |  | 

### Return type

[**Application**](Application.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultApplicationAdvancedSettings

> ApplicationAdvancedSettings GetDefaultApplicationAdvancedSettings(ctx).Execute()

List default application advanced settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.GetDefaultApplicationAdvancedSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetDefaultApplicationAdvancedSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultApplicationAdvancedSettings`: ApplicationAdvancedSettings
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetDefaultApplicationAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultApplicationAdvancedSettingsRequest struct via the builder pattern


### Return type

[**ApplicationAdvancedSettings**](ApplicationAdvancedSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentApplicationCurrentScale

> EnvironmentApplicationsCurrentScaleResponseList GetEnvironmentApplicationCurrentScale(ctx, environmentId).Execute()

List current scaling information for each application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.GetEnvironmentApplicationCurrentScale(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetEnvironmentApplicationCurrentScale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentApplicationCurrentScale`: EnvironmentApplicationsCurrentScaleResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetEnvironmentApplicationCurrentScale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentApplicationCurrentScaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentApplicationsCurrentScaleResponseList**](EnvironmentApplicationsCurrentScaleResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentApplicationCurrentStorage

> EnvironmentApplicationsStorageResponseList GetEnvironmentApplicationCurrentStorage(ctx, environmentId).Execute()

List current storage disk usage for each application

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.GetEnvironmentApplicationCurrentStorage(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetEnvironmentApplicationCurrentStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentApplicationCurrentStorage`: EnvironmentApplicationsStorageResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetEnvironmentApplicationCurrentStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentApplicationCurrentStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentApplicationsStorageResponseList**](EnvironmentApplicationsStorageResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentApplicationStatus

> ReferenceObjectStatusResponseList GetEnvironmentApplicationStatus(ctx, environmentId).Execute()

List all environment applications statuses



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.GetEnvironmentApplicationStatus(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetEnvironmentApplicationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentApplicationStatus`: ReferenceObjectStatusResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetEnvironmentApplicationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentApplicationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReferenceObjectStatusResponseList**](ReferenceObjectStatusResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentApplicationSupportedLanguages

> EnvironmentApplicationsSupportedLanguageList GetEnvironmentApplicationSupportedLanguages(ctx, environmentId).Execute()

List supported languages



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.GetEnvironmentApplicationSupportedLanguages(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetEnvironmentApplicationSupportedLanguages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentApplicationSupportedLanguages`: EnvironmentApplicationsSupportedLanguageList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetEnvironmentApplicationSupportedLanguages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentApplicationSupportedLanguagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentApplicationsSupportedLanguageList**](EnvironmentApplicationsSupportedLanguageList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplication

> ApplicationResponseList ListApplication(ctx, environmentId).Execute()

List applications

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.ListApplication(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.ListApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplication`: ApplicationResponseList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.ListApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationResponseList**](ApplicationResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

