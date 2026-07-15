# \IPsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListQoveryIps**](IPsAPI.md#ListQoveryIps) | **Get** /ips | List Qovery NAT gateway IP addresses



## ListQoveryIps

> QoveryIpsResponse ListQoveryIps(ctx).Execute()

List Qovery NAT gateway IP addresses



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
	resp, r, err := apiClient.IPsAPI.ListQoveryIps(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPsAPI.ListQoveryIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQoveryIps`: QoveryIpsResponse
	fmt.Fprintf(os.Stdout, "Response from `IPsAPI.ListQoveryIps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListQoveryIpsRequest struct via the builder pattern


### Return type

[**QoveryIpsResponse**](QoveryIpsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

