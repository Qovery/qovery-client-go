# \EnvironmentAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckContainerImage**](EnvironmentAPI.md#CheckContainerImage) | **Post** /environment/{environmentId}/checkContainerImage | Check container image configuration is correct
[**CheckDockerfile**](EnvironmentAPI.md#CheckDockerfile) | **Post** /environment/{environmentId}/checkDockerfile | Check dockerfile configuration is correct
[**CheckGitFile**](EnvironmentAPI.md#CheckGitFile) | **Post** /environment/{environmentId}/checkGitFile | Check git file configuration is correct
[**CheckHelmRepository**](EnvironmentAPI.md#CheckHelmRepository) | **Post** /environment/{environmentId}/checkHelmRepository | Check helm repository configuration is correct
[**DeployAllApplications**](EnvironmentAPI.md#DeployAllApplications) | **Post** /environment/{environmentId}/application/deploy | Deploy applications



## CheckContainerImage

> map[string]interface{} CheckContainerImage(ctx, environmentId).ContainerImageCheckRequest(containerImageCheckRequest).Execute()

Check container image configuration is correct

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
	containerImageCheckRequest := *openapiclient.NewContainerImageCheckRequest("ImageName_example", "Tag_example") // ContainerImageCheckRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.CheckContainerImage(context.Background(), environmentId).ContainerImageCheckRequest(containerImageCheckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.CheckContainerImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckContainerImage`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.CheckContainerImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckContainerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerImageCheckRequest** | [**ContainerImageCheckRequest**](ContainerImageCheckRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## CheckGitFile

> map[string]interface{} CheckGitFile(ctx, environmentId).GitFileCheckRequest(gitFileCheckRequest).Execute()

Check git file configuration is correct

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
	gitFileCheckRequest := *openapiclient.NewGitFileCheckRequest(*openapiclient.NewHelmGitRepositoryRequest("https://github.com/Qovery/simple-node-app"), []string{"Files_example"}) // GitFileCheckRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.CheckGitFile(context.Background(), environmentId).GitFileCheckRequest(gitFileCheckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.CheckGitFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckGitFile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.CheckGitFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckGitFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gitFileCheckRequest** | [**GitFileCheckRequest**](GitFileCheckRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckHelmRepository

> map[string]interface{} CheckHelmRepository(ctx, environmentId).HelmCheckRequest(helmCheckRequest).Execute()

Check helm repository configuration is correct

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
	helmCheckRequest := *openapiclient.NewHelmCheckRequest(*openapiclient.NewHelmGitRepositoryRequest("https://github.com/Qovery/simple-node-app")) // HelmCheckRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.CheckHelmRepository(context.Background(), environmentId).HelmCheckRequest(helmCheckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.CheckHelmRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckHelmRepository`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.CheckHelmRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckHelmRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **helmCheckRequest** | [**HelmCheckRequest**](HelmCheckRequest.md) |  | 

### Return type

**map[string]interface{}**

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

