# \VariableMainCallsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVariableAlias**](VariableMainCallsApi.md#CreateVariableAlias) | **Post** /variable/{variableId}/alias | Create a variable alias



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
    openapiclient "./openapi"
)

func main() {
    variableId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Variable ID
    variableAliasRequest := *openapiclient.NewVariableAliasRequest("Key_example", openapiclient.APIVariableScopeEnum("APPLICATION"), "AliasParentId_example") // VariableAliasRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariableMainCallsApi.CreateVariableAlias(context.Background(), variableId).VariableAliasRequest(variableAliasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariableMainCallsApi.CreateVariableAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVariableAlias`: VariableResponse
    fmt.Fprintf(os.Stdout, "Response from `VariableMainCallsApi.CreateVariableAlias`: %v\n", resp)
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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

