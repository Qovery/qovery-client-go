# \AccountInfoAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditAccountInformation**](AccountInfoAPI.md#EditAccountInformation) | **Put** /account | Edit account information
[**GetAccountInformation**](AccountInfoAPI.md#GetAccountInformation) | **Get** /account | Get Account information



## EditAccountInformation

> AccountInfo EditAccountInformation(ctx).AccountInfoEditRequest(accountInfoEditRequest).Execute()

Edit account information

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
	accountInfoEditRequest := *openapiclient.NewAccountInfoEditRequest() // AccountInfoEditRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountInfoAPI.EditAccountInformation(context.Background()).AccountInfoEditRequest(accountInfoEditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountInfoAPI.EditAccountInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAccountInformation`: AccountInfo
	fmt.Fprintf(os.Stdout, "Response from `AccountInfoAPI.EditAccountInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditAccountInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountInfoEditRequest** | [**AccountInfoEditRequest**](AccountInfoEditRequest.md) |  | 

### Return type

[**AccountInfo**](AccountInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInformation

> AccountInfo GetAccountInformation(ctx).Execute()

Get Account information

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
	resp, r, err := apiClient.AccountInfoAPI.GetAccountInformation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountInfoAPI.GetAccountInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInformation`: AccountInfo
	fmt.Fprintf(os.Stdout, "Response from `AccountInfoAPI.GetAccountInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInformationRequest struct via the builder pattern


### Return type

[**AccountInfo**](AccountInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

