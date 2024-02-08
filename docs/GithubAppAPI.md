# \GithubAppAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationGithubAppConnect**](GithubAppAPI.md#OrganizationGithubAppConnect) | **Post** /organization/{organizationId}/github/connect | Connect a github account to an organization
[**OrganizationGithubAppDisconnect**](GithubAppAPI.md#OrganizationGithubAppDisconnect) | **Delete** /organization/{organizationId}/github/disconnect | Disconnect a github account from an organization



## OrganizationGithubAppConnect

> OrganizationGithubAppConnect(ctx, organizationId).OrganizationGithubAppConnectRequest(organizationGithubAppConnectRequest).Execute()

Connect a github account to an organization

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
    organizationGithubAppConnectRequest := *openapiclient.NewOrganizationGithubAppConnectRequest("InstallationId_example", "Code_example") // OrganizationGithubAppConnectRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GithubAppAPI.OrganizationGithubAppConnect(context.Background(), organizationId).OrganizationGithubAppConnectRequest(organizationGithubAppConnectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GithubAppAPI.OrganizationGithubAppConnect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGithubAppConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationGithubAppConnectRequest** | [**OrganizationGithubAppConnectRequest**](OrganizationGithubAppConnectRequest.md) |  | 

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


## OrganizationGithubAppDisconnect

> OrganizationGithubAppDisconnect(ctx, organizationId).Force(force).Execute()

Disconnect a github account from an organization

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
    force := true // bool | Indicates if the github app should be disconnected despite github applications linked to organization (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GithubAppAPI.OrganizationGithubAppDisconnect(context.Background(), organizationId).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GithubAppAPI.OrganizationGithubAppDisconnect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGithubAppDisconnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Indicates if the github app should be disconnected despite github applications linked to organization | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

