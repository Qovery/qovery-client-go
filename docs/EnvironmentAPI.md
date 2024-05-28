# \EnvironmentAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckDockerfile**](EnvironmentAPI.md#CheckDockerfile) | **Post** /environment/{environmentId}/checkDockerfile | Check dockerfile configuration is correct
[**DeployAllApplications**](EnvironmentAPI.md#DeployAllApplications) | **Post** /environment/{environmentId}/application/deploy | Deploy applications



## CheckDockerfile

> DockerfileCheckResponse CheckDockerfile(ctx, environmentId).DockerfileCheckRequest(dockerfileCheckRequest).Execute()

Check dockerfile configuration is correct

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
    dockerfileCheckRequest := *openapiclient.NewDockerfileCheckRequest(*openapiclient.NewApplicationGitRepositoryRequest("https://github.com/Qovery/simple-node-app"), "DockerfilePath_example") // DockerfileCheckRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAPI.CheckDockerfile(context.Background(), environmentId).DockerfileCheckRequest(dockerfileCheckRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.CheckDockerfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckDockerfile`: DockerfileCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.CheckDockerfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckDockerfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dockerfileCheckRequest** | [**DockerfileCheckRequest**](DockerfileCheckRequest.md) |  | 

### Return type

[**DockerfileCheckResponse**](DockerfileCheckResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeployAllApplications

> Status DeployAllApplications(ctx, environmentId).DeployAllRequest(deployAllRequest).Execute()

Deploy applications



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Environment ID
    deployAllRequest := *openapiclient.NewDeployAllRequest() // DeployAllRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAPI.DeployAllApplications(context.Background(), environmentId).DeployAllRequest(deployAllRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.DeployAllApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployAllApplications`: Status
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.DeployAllApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployAllApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deployAllRequest** | [**DeployAllRequest**](DeployAllRequest.md) |  | 

### Return type

[**Status**](Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

