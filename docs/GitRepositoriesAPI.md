# \GitRepositoriesAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDirectoriesFromGitRepository**](GitRepositoriesAPI.md#ListDirectoriesFromGitRepository) | **Post** /organization/{organizationId}/listDirectoriesFromGitRepository | List directories from a git repository



## ListDirectoriesFromGitRepository

> ListDirectoriesFromGitRepository200Response ListDirectoriesFromGitRepository(ctx, organizationId).ApplicationGitRepositoryRequest(applicationGitRepositoryRequest).Execute()

List directories from a git repository



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
	applicationGitRepositoryRequest := *openapiclient.NewApplicationGitRepositoryRequest("https://github.com/Qovery/simple-node-app", openapiclient.GitProviderEnum("BITBUCKET")) // ApplicationGitRepositoryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitRepositoriesAPI.ListDirectoriesFromGitRepository(context.Background(), organizationId).ApplicationGitRepositoryRequest(applicationGitRepositoryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitRepositoriesAPI.ListDirectoriesFromGitRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDirectoriesFromGitRepository`: ListDirectoriesFromGitRepository200Response
	fmt.Fprintf(os.Stdout, "Response from `GitRepositoriesAPI.ListDirectoriesFromGitRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDirectoriesFromGitRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationGitRepositoryRequest** | [**ApplicationGitRepositoryRequest**](ApplicationGitRepositoryRequest.md) |  | 

### Return type

[**ListDirectoriesFromGitRepository200Response**](ListDirectoriesFromGitRepository200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

