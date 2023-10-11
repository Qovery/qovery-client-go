# \OrganizationAccountGitRepositoriesAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationBitbucketRepositories**](OrganizationAccountGitRepositoriesAPI.md#GetOrganizationBitbucketRepositories) | **Get** /organization/{organizationId}/account/bitbucket/repository | Get bitbucket repositories of the connected user
[**GetOrganizationBitbucketRepositoryBranches**](OrganizationAccountGitRepositoriesAPI.md#GetOrganizationBitbucketRepositoryBranches) | **Get** /organization/{organizationId}/account/bitbucket/repository/branch | Get bitbucket branches of the specified repository
[**GetOrganizationGitProviderAccount**](OrganizationAccountGitRepositoriesAPI.md#GetOrganizationGitProviderAccount) | **Get** /organization/{organizationId}/account/gitAuthProvider | Get git provider accounts
[**GetOrganizationGithubRepositories**](OrganizationAccountGitRepositoriesAPI.md#GetOrganizationGithubRepositories) | **Get** /organization/{organizationId}/account/github/repository | Get github repositories of the connected user
[**GetOrganizationGithubRepositoryBranches**](OrganizationAccountGitRepositoriesAPI.md#GetOrganizationGithubRepositoryBranches) | **Get** /organization/{organizationId}/account/github/repository/branch | Get github branches of the specified repository
[**GetOrganizationGitlabRepositories**](OrganizationAccountGitRepositoriesAPI.md#GetOrganizationGitlabRepositories) | **Get** /organization/{organizationId}/account/gitlab/repository | Get gitlab repositories of the connected user
[**GetOrganizationGitlabRepositoryBranches**](OrganizationAccountGitRepositoriesAPI.md#GetOrganizationGitlabRepositoryBranches) | **Get** /organization/{organizationId}/account/gitlab/repository/branch | Get gitlab branches of the specified repository



## GetOrganizationBitbucketRepositories

> GitRepositoryResponseList GetOrganizationBitbucketRepositories(ctx, organizationId).Execute()

Get bitbucket repositories of the connected user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAccountGitRepositoriesAPI.GetOrganizationBitbucketRepositories(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAccountGitRepositoriesAPI.GetOrganizationBitbucketRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationBitbucketRepositories`: GitRepositoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAccountGitRepositoriesAPI.GetOrganizationBitbucketRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBitbucketRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitRepositoryResponseList**](GitRepositoryResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationBitbucketRepositoryBranches

> GitRepositoryBranchResponseList GetOrganizationBitbucketRepositoryBranches(ctx, organizationId).Name(name).Execute()

Get bitbucket branches of the specified repository

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
    name := "name_example" // string | The name of the repository where to retrieve the branches (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAccountGitRepositoriesAPI.GetOrganizationBitbucketRepositoryBranches(context.Background(), organizationId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAccountGitRepositoriesAPI.GetOrganizationBitbucketRepositoryBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationBitbucketRepositoryBranches`: GitRepositoryBranchResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAccountGitRepositoriesAPI.GetOrganizationBitbucketRepositoryBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBitbucketRepositoryBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the repository where to retrieve the branches | 

### Return type

[**GitRepositoryBranchResponseList**](GitRepositoryBranchResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationGitProviderAccount

> GitAuthProviderResponseList GetOrganizationGitProviderAccount(ctx, organizationId).Execute()

Get git provider accounts

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAccountGitRepositoriesAPI.GetOrganizationGitProviderAccount(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAccountGitRepositoriesAPI.GetOrganizationGitProviderAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationGitProviderAccount`: GitAuthProviderResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAccountGitRepositoriesAPI.GetOrganizationGitProviderAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationGitProviderAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitAuthProviderResponseList**](GitAuthProviderResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationGithubRepositories

> GitRepositoryResponseList GetOrganizationGithubRepositories(ctx, organizationId).Execute()

Get github repositories of the connected user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAccountGitRepositoriesAPI.GetOrganizationGithubRepositories(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAccountGitRepositoriesAPI.GetOrganizationGithubRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationGithubRepositories`: GitRepositoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAccountGitRepositoriesAPI.GetOrganizationGithubRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationGithubRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitRepositoryResponseList**](GitRepositoryResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationGithubRepositoryBranches

> GitRepositoryBranchResponseList GetOrganizationGithubRepositoryBranches(ctx, organizationId).Name(name).Execute()

Get github branches of the specified repository

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
    name := "name_example" // string | The name of the repository where to retrieve the branches (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAccountGitRepositoriesAPI.GetOrganizationGithubRepositoryBranches(context.Background(), organizationId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAccountGitRepositoriesAPI.GetOrganizationGithubRepositoryBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationGithubRepositoryBranches`: GitRepositoryBranchResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAccountGitRepositoriesAPI.GetOrganizationGithubRepositoryBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationGithubRepositoryBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the repository where to retrieve the branches | 

### Return type

[**GitRepositoryBranchResponseList**](GitRepositoryBranchResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationGitlabRepositories

> GitRepositoryResponseList GetOrganizationGitlabRepositories(ctx, organizationId).Execute()

Get gitlab repositories of the connected user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAccountGitRepositoriesAPI.GetOrganizationGitlabRepositories(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAccountGitRepositoriesAPI.GetOrganizationGitlabRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationGitlabRepositories`: GitRepositoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAccountGitRepositoriesAPI.GetOrganizationGitlabRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationGitlabRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitRepositoryResponseList**](GitRepositoryResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationGitlabRepositoryBranches

> GitRepositoryBranchResponseList GetOrganizationGitlabRepositoryBranches(ctx, organizationId).Name(name).Execute()

Get gitlab branches of the specified repository

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
    name := "name_example" // string | The name of the repository to retrieve the branches (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAccountGitRepositoriesAPI.GetOrganizationGitlabRepositoryBranches(context.Background(), organizationId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAccountGitRepositoriesAPI.GetOrganizationGitlabRepositoryBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationGitlabRepositoryBranches`: GitRepositoryBranchResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAccountGitRepositoriesAPI.GetOrganizationGitlabRepositoryBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationGitlabRepositoryBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the repository to retrieve the branches | 

### Return type

[**GitRepositoryBranchResponseList**](GitRepositoryBranchResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

