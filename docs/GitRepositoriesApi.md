# \GitRepositoriesApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGitProviderAccount**](GitRepositoriesApi.md#GetGitProviderAccount) | **Get** /account/gitAuthProvider | Get git provider accounts
[**GetGithubRepositories**](GitRepositoriesApi.md#GetGithubRepositories) | **Get** /account/github/repository | Get github repositories of the connected user
[**GetGitlabRepositories**](GitRepositoriesApi.md#GetGitlabRepositories) | **Get** /account/gitlab/repository | Get gitlab repositories of the connected user



## GetGitProviderAccount

> GitAuthProviderResponseList GetGitProviderAccount(ctx).Execute()

Get git provider accounts

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GitRepositoriesApi.GetGitProviderAccount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitRepositoriesApi.GetGitProviderAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGitProviderAccount`: GitAuthProviderResponseList
    fmt.Fprintf(os.Stdout, "Response from `GitRepositoriesApi.GetGitProviderAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitProviderAccountRequest struct via the builder pattern


### Return type

[**GitAuthProviderResponseList**](GitAuthProviderResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubRepositories

> GitRepositoryResponseList GetGithubRepositories(ctx).Execute()

Get github repositories of the connected user

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GitRepositoriesApi.GetGithubRepositories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitRepositoriesApi.GetGithubRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubRepositories`: GitRepositoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `GitRepositoriesApi.GetGithubRepositories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubRepositoriesRequest struct via the builder pattern


### Return type

[**GitRepositoryResponseList**](GitRepositoryResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitlabRepositories

> GitRepositoryResponseList GetGitlabRepositories(ctx).Execute()

Get gitlab repositories of the connected user

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GitRepositoriesApi.GetGitlabRepositories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitRepositoriesApi.GetGitlabRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGitlabRepositories`: GitRepositoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `GitRepositoriesApi.GetGitlabRepositories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitlabRepositoriesRequest struct via the builder pattern


### Return type

[**GitRepositoryResponseList**](GitRepositoryResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

