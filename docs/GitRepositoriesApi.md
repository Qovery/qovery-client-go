# \GitRepositoriesApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBitbucketRepositories**](GitRepositoriesApi.md#GetBitbucketRepositories) | **Get** /account/bitbucket/repository | Get bitbucket repositories of the connected user
[**GetBitbucketRepositoryBranches**](GitRepositoriesApi.md#GetBitbucketRepositoryBranches) | **Get** /account/bitbucket/repository/branch | Get bitbucket branches of the specified repository
[**GetGitProviderAccount**](GitRepositoriesApi.md#GetGitProviderAccount) | **Get** /account/gitAuthProvider | Get git provider accounts
[**GetGithubRepositories**](GitRepositoriesApi.md#GetGithubRepositories) | **Get** /account/github/repository | Get github repositories of the connected user
[**GetGithubRepositoryBranches**](GitRepositoriesApi.md#GetGithubRepositoryBranches) | **Get** /account/github/repository/branch | Get github branches of the specified repository
[**GetGitlabRepositories**](GitRepositoriesApi.md#GetGitlabRepositories) | **Get** /account/gitlab/repository | Get gitlab repositories of the connected user
[**GetGitlabRepositoryBranches**](GitRepositoriesApi.md#GetGitlabRepositoryBranches) | **Get** /account/gitlab/repository/branch | Get gitlab branches of the specified repository



## GetBitbucketRepositories

> GitRepositoryResponseList GetBitbucketRepositories(ctx).Execute()

Get bitbucket repositories of the connected user

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
    resp, r, err := apiClient.GitRepositoriesApi.GetBitbucketRepositories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitRepositoriesApi.GetBitbucketRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBitbucketRepositories`: GitRepositoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `GitRepositoriesApi.GetBitbucketRepositories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBitbucketRepositoriesRequest struct via the builder pattern


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


## GetBitbucketRepositoryBranches

> GitRepositoryBranchResponseList GetBitbucketRepositoryBranches(ctx).Name(name).Execute()

Get bitbucket branches of the specified repository

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
    name := "name_example" // string | The name of the repository where to retrieve the branches (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitRepositoriesApi.GetBitbucketRepositoryBranches(context.Background()).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitRepositoriesApi.GetBitbucketRepositoryBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBitbucketRepositoryBranches`: GitRepositoryBranchResponseList
    fmt.Fprintf(os.Stdout, "Response from `GitRepositoriesApi.GetBitbucketRepositoryBranches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBitbucketRepositoryBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the repository where to retrieve the branches | 

### Return type

[**GitRepositoryBranchResponseList**](GitRepositoryBranchResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitRepositoriesApi.GetGitProviderAccount(context.Background()).Execute()
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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitRepositoriesApi.GetGithubRepositories(context.Background()).Execute()
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


## GetGithubRepositoryBranches

> GitRepositoryBranchResponseList GetGithubRepositoryBranches(ctx).Name(name).Execute()

Get github branches of the specified repository

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
    name := "name_example" // string | The name of the repository where to retrieve the branches (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitRepositoriesApi.GetGithubRepositoryBranches(context.Background()).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitRepositoriesApi.GetGithubRepositoryBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubRepositoryBranches`: GitRepositoryBranchResponseList
    fmt.Fprintf(os.Stdout, "Response from `GitRepositoriesApi.GetGithubRepositoryBranches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubRepositoryBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the repository where to retrieve the branches | 

### Return type

[**GitRepositoryBranchResponseList**](GitRepositoryBranchResponseList.md)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitRepositoriesApi.GetGitlabRepositories(context.Background()).Execute()
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


## GetGitlabRepositoryBranches

> GitRepositoryBranchResponseList GetGitlabRepositoryBranches(ctx).Name(name).Execute()

Get gitlab branches of the specified repository

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
    name := "name_example" // string | The name of the repository to retrieve the branches (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GitRepositoriesApi.GetGitlabRepositoryBranches(context.Background()).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GitRepositoriesApi.GetGitlabRepositoryBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGitlabRepositoryBranches`: GitRepositoryBranchResponseList
    fmt.Fprintf(os.Stdout, "Response from `GitRepositoriesApi.GetGitlabRepositoryBranches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGitlabRepositoryBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the repository to retrieve the branches | 

### Return type

[**GitRepositoryBranchResponseList**](GitRepositoryBranchResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

