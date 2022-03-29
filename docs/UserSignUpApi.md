# \UserSignUpApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserSignUp**](UserSignUpApi.md#CreateUserSignUp) | **Post** /admin/userSignUp | Send Sign Up request
[**GetUserSignUp**](UserSignUpApi.md#GetUserSignUp) | **Get** /admin/userSignUp | Get Sign up information



## CreateUserSignUp

> CreateUserSignUp(ctx).InlineObject2(inlineObject2).Execute()

Send Sign Up request



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
    inlineObject2 := *openapiclient.NewInlineObject2("FirstName_example", "LastName_example", "UserEmail_example", "TypeOfUse_example", "QoveryUsage_example") // InlineObject2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSignUpApi.CreateUserSignUp(context.Background()).InlineObject2(inlineObject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSignUpApi.CreateUserSignUp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserSignUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSignUp

> BaseResponse GetUserSignUp(ctx).Execute()

Get Sign up information



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
    resp, r, err := apiClient.UserSignUpApi.GetUserSignUp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSignUpApi.GetUserSignUp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSignUp`: BaseResponse
    fmt.Fprintf(os.Stdout, "Response from `UserSignUpApi.GetUserSignUp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSignUpRequest struct via the builder pattern


### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

