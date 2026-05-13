# \SecretManagerAccessAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSecretManagerAccessExternalSecrets**](SecretManagerAccessAPI.md#ListSecretManagerAccessExternalSecrets) | **Get** /secretManagerAccess/{secretManagerAccessId}/associatedServices | List external secrets used by a secret manager access
[**ListUpstreamSecretsFromSecretProvider**](SecretManagerAccessAPI.md#ListUpstreamSecretsFromSecretProvider) | **Get** /secretManagerAccess/{secretManagerAccessId}/secrets | List upstream secrets from secret provider



## ListSecretManagerAccessExternalSecrets

> ExternalSecretAssociatedServiceResponseList ListSecretManagerAccessExternalSecrets(ctx, secretManagerAccessId).Execute()

List external secrets used by a secret manager access



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
	secretManagerAccessId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Secret Manager Access ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretManagerAccessAPI.ListSecretManagerAccessExternalSecrets(context.Background(), secretManagerAccessId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretManagerAccessAPI.ListSecretManagerAccessExternalSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecretManagerAccessExternalSecrets`: ExternalSecretAssociatedServiceResponseList
	fmt.Fprintf(os.Stdout, "Response from `SecretManagerAccessAPI.ListSecretManagerAccessExternalSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretManagerAccessId** | **string** | Secret Manager Access ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecretManagerAccessExternalSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalSecretAssociatedServiceResponseList**](ExternalSecretAssociatedServiceResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUpstreamSecretsFromSecretProvider

> ProviderSecrets ListUpstreamSecretsFromSecretProvider(ctx, secretManagerAccessId).NamePrefix(namePrefix).Execute()

List upstream secrets from secret provider



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
	secretManagerAccessId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Secret Manager Access ID
	namePrefix := "namePrefix_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretManagerAccessAPI.ListUpstreamSecretsFromSecretProvider(context.Background(), secretManagerAccessId).NamePrefix(namePrefix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretManagerAccessAPI.ListUpstreamSecretsFromSecretProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUpstreamSecretsFromSecretProvider`: ProviderSecrets
	fmt.Fprintf(os.Stdout, "Response from `SecretManagerAccessAPI.ListUpstreamSecretsFromSecretProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretManagerAccessId** | **string** | Secret Manager Access ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUpstreamSecretsFromSecretProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namePrefix** | **string** |  | 

### Return type

[**ProviderSecrets**](ProviderSecrets.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

