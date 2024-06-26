# \UserSignUpAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserSignUp**](UserSignUpAPI.md#CreateUserSignUp) | **Post** /admin/userSignUp | Send Sign Up request
[**GetUserSignUp**](UserSignUpAPI.md#GetUserSignUp) | **Get** /admin/userSignUp | Get Sign up information



## CreateUserSignUp

> CreateUserSignUp(ctx).SignUpRequest(signUpRequest).Execute()

Send Sign Up request



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
	signUpRequest := *openapiclient.NewSignUpRequest("FirstName_example", "LastName_example", "UserEmail_example", openapiclient.TypeOfUseEnum("PERSONAL"), "QoveryUsage_example") // SignUpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserSignUpAPI.CreateUserSignUp(context.Background()).SignUpRequest(signUpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSignUpAPI.CreateUserSignUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserSignUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signUpRequest** | [**SignUpRequest**](SignUpRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSignUp

> SignUp GetUserSignUp(ctx).Execute()

Get Sign up information



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSignUpAPI.GetUserSignUp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSignUpAPI.GetUserSignUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSignUp`: SignUp
	fmt.Fprintf(os.Stdout, "Response from `UserSignUpAPI.GetUserSignUp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSignUpRequest struct via the builder pattern


### Return type

[**SignUp**](SignUp.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

