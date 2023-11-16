# \VariableMainCallsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVariable**](VariableMainCallsAPI.md#CreateVariable) | **Post** /variable | Create a variable
[**CreateVariableAlias**](VariableMainCallsAPI.md#CreateVariableAlias) | **Post** /variable/{variableId}/alias | Create a variable alias
[**CreateVariableOverride**](VariableMainCallsAPI.md#CreateVariableOverride) | **Post** /variable/{variableId}/override | Create a variable override
[**DeleteVariable**](VariableMainCallsAPI.md#DeleteVariable) | **Delete** /variable/{variableId} | Delete a variable
[**EditVariable**](VariableMainCallsAPI.md#EditVariable) | **Put** /variable/{variableId} | Edit a variable
[**ImportEnvironmentVariables**](VariableMainCallsAPI.md#ImportEnvironmentVariables) | **Post** /variable/import | Import variables
[**ListVariables**](VariableMainCallsAPI.md#ListVariables) | **Get** /variable | List variables



## CreateVariable

> VariableResponse CreateVariable(ctx).VariableRequest(variableRequest).Execute()

Create a variable



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
    variableRequest := *openapiclient.NewVariableRequest("Key_example", "Value_example", false, openapiclient.APIVariableScopeEnum("APPLICATION"), "VariableParentId_example") // VariableRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariableMainCallsAPI.CreateVariable(context.Background()).VariableRequest(variableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariableMainCallsAPI.CreateVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVariable`: VariableResponse
    fmt.Fprintf(os.Stdout, "Response from `VariableMainCallsAPI.CreateVariable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableRequest** | [**VariableRequest**](VariableRequest.md) |  | 

### Return type

[**VariableResponse**](VariableResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVariableAlias

> VariableResponse CreateVariableAlias(ctx, variableId).VariableAliasRequest(variableAliasRequest).Execute()

Create a variable alias



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
    variableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Variable ID
    variableAliasRequest := *openapiclient.NewVariableAliasRequest("Key_example", openapiclient.APIVariableScopeEnum("APPLICATION"), "AliasParentId_example") // VariableAliasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariableMainCallsAPI.CreateVariableAlias(context.Background(), variableId).VariableAliasRequest(variableAliasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariableMainCallsAPI.CreateVariableAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVariableAlias`: VariableResponse
    fmt.Fprintf(os.Stdout, "Response from `VariableMainCallsAPI.CreateVariableAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** | Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVariableAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variableAliasRequest** | [**VariableAliasRequest**](VariableAliasRequest.md) |  | 

### Return type

[**VariableResponse**](VariableResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVariableOverride

> VariableResponse CreateVariableOverride(ctx, variableId).VariableOverrideRequest(variableOverrideRequest).Execute()

Create a variable override



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
    variableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Variable ID
    variableOverrideRequest := *openapiclient.NewVariableOverrideRequest("Value_example", openapiclient.APIVariableScopeEnum("APPLICATION"), "OverrideParentId_example") // VariableOverrideRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariableMainCallsAPI.CreateVariableOverride(context.Background(), variableId).VariableOverrideRequest(variableOverrideRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariableMainCallsAPI.CreateVariableOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVariableOverride`: VariableResponse
    fmt.Fprintf(os.Stdout, "Response from `VariableMainCallsAPI.CreateVariableOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** | Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVariableOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variableOverrideRequest** | [**VariableOverrideRequest**](VariableOverrideRequest.md) |  | 

### Return type

[**VariableResponse**](VariableResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVariable

> DeleteVariable(ctx, variableId).Execute()

Delete a variable



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
    variableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Variable ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VariableMainCallsAPI.DeleteVariable(context.Background(), variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariableMainCallsAPI.DeleteVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** | Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVariableRequest struct via the builder pattern


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


## EditVariable

> VariableResponse EditVariable(ctx, variableId).VariableEditRequest(variableEditRequest).Execute()

Edit a variable



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
    variableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Variable ID
    variableEditRequest := *openapiclient.NewVariableEditRequest("Key_example", "Value_example") // VariableEditRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariableMainCallsAPI.EditVariable(context.Background(), variableId).VariableEditRequest(variableEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariableMainCallsAPI.EditVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditVariable`: VariableResponse
    fmt.Fprintf(os.Stdout, "Response from `VariableMainCallsAPI.EditVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** | Variable ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variableEditRequest** | [**VariableEditRequest**](VariableEditRequest.md) |  | 

### Return type

[**VariableResponse**](VariableResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportEnvironmentVariables

> VariableImport ImportEnvironmentVariables(ctx).ServiceId(serviceId).Scope(scope).VariableImportRequest(variableImportRequest).Execute()

Import variables



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
    serviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | service id
    scope := openapiclient.APIVariableScopeEnum("APPLICATION") // APIVariableScopeEnum | scope
    variableImportRequest := *openapiclient.NewVariableImportRequest(false, []openapiclient.VariableImportRequestVarsInner{*openapiclient.NewVariableImportRequestVarsInner("Name_example", "Value_example", openapiclient.APIVariableScopeEnum("APPLICATION"), false)}) // VariableImportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariableMainCallsAPI.ImportEnvironmentVariables(context.Background()).ServiceId(serviceId).Scope(scope).VariableImportRequest(variableImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariableMainCallsAPI.ImportEnvironmentVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportEnvironmentVariables`: VariableImport
    fmt.Fprintf(os.Stdout, "Response from `VariableMainCallsAPI.ImportEnvironmentVariables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportEnvironmentVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceId** | **string** | service id | 
 **scope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) | scope | 
 **variableImportRequest** | [**VariableImportRequest**](VariableImportRequest.md) |  | 

### Return type

[**VariableImport**](VariableImport.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVariables

> VariableResponseList ListVariables(ctx).ParentId(parentId).Scope(scope).IsSecret(isSecret).Execute()

List variables



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
    parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | it filters the list by returning only the variables accessible by the selected parent_id. This field shall contain the id of a project, environment or service depending on the selected scope. Example, if scope = APPLICATION and parent_id=<application_id>, the result will contain any variable accessible by the application. The result will contain also any variable declared at an higher scope.
    scope := openapiclient.APIVariableScopeEnum("APPLICATION") // APIVariableScopeEnum | the type of the parent_id (application, project, environment etc..).
    isSecret := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariableMainCallsAPI.ListVariables(context.Background()).ParentId(parentId).Scope(scope).IsSecret(isSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariableMainCallsAPI.ListVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVariables`: VariableResponseList
    fmt.Fprintf(os.Stdout, "Response from `VariableMainCallsAPI.ListVariables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentId** | **string** | it filters the list by returning only the variables accessible by the selected parent_id. This field shall contain the id of a project, environment or service depending on the selected scope. Example, if scope &#x3D; APPLICATION and parent_id&#x3D;&lt;application_id&gt;, the result will contain any variable accessible by the application. The result will contain also any variable declared at an higher scope. | 
 **scope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) | the type of the parent_id (application, project, environment etc..). | 
 **isSecret** | **bool** |  | 

### Return type

[**VariableResponseList**](VariableResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

