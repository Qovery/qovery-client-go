# \OrganizationMainCallsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGitToken**](OrganizationMainCallsAPI.md#CreateGitToken) | **Post** /organization/{organizationId}/gitToken | Create a git token
[**CreateOrganization**](OrganizationMainCallsAPI.md#CreateOrganization) | **Post** /organization | Create an organization
[**DeleteGitToken**](OrganizationMainCallsAPI.md#DeleteGitToken) | **Delete** /organization/{organizationId}/gitToken/{gitTokenId} | Delete a git token
[**DeleteOrganization**](OrganizationMainCallsAPI.md#DeleteOrganization) | **Delete** /organization/{organizationId} | Delete an organization
[**EditGitToken**](OrganizationMainCallsAPI.md#EditGitToken) | **Put** /organization/{organizationId}/gitToken/{gitTokenId} | Edit a git token
[**EditOrganization**](OrganizationMainCallsAPI.md#EditOrganization) | **Put** /organization/{organizationId} | Edit an organization
[**GetContainerRegistryAssociatedServices**](OrganizationMainCallsAPI.md#GetContainerRegistryAssociatedServices) | **Get** /organization/{organizationId}/containerRegistry/{containerRegistryId}/associatedServices | Get organization container registry associated services
[**GetGitTokenAssociatedServices**](OrganizationMainCallsAPI.md#GetGitTokenAssociatedServices) | **Get** /organization/{organizationId}/gitToken/{gitTokenId}/associatedServices | Get organization git token associated services
[**GetHelmRepositoryAssociatedServices**](OrganizationMainCallsAPI.md#GetHelmRepositoryAssociatedServices) | **Get** /organization/{organizationId}/helmRepository/{helmRepositoryId}/associatedServices | Get organization helm repository associated services
[**GetOrganization**](OrganizationMainCallsAPI.md#GetOrganization) | **Get** /organization/{organizationId} | Get organization by ID
[**GetOrganizationGitToken**](OrganizationMainCallsAPI.md#GetOrganizationGitToken) | **Get** /organization/{organizationId}/gitToken/{gitTokenId} | Get organization git token
[**ListOrganization**](OrganizationMainCallsAPI.md#ListOrganization) | **Get** /organization | List user organizations
[**ListOrganizationAvailableRoles**](OrganizationMainCallsAPI.md#ListOrganizationAvailableRoles) | **Get** /organization/{organizationId}/availableRole | List organization available roles
[**ListOrganizationCredentials**](OrganizationMainCallsAPI.md#ListOrganizationCredentials) | **Get** /organization/{organizationId}/credentials | List credentials of an organization and their associated clusters
[**ListOrganizationGitTokens**](OrganizationMainCallsAPI.md#ListOrganizationGitTokens) | **Get** /organization/{organizationId}/gitToken | List organization git tokens
[**ListServicesByOrganizationId**](OrganizationMainCallsAPI.md#ListServicesByOrganizationId) | **Get** /organization/{organizationId}/services | List Services By OrganizationId
[**ListTfVarsFilesFromGitRepo**](OrganizationMainCallsAPI.md#ListTfVarsFilesFromGitRepo) | **Post** /organization/{organizationId}/listTfVarsFilesFromGitRepo | List Terraform tfvars files from Git repository
[**ParseTerraformVariablesFromGitRepo**](OrganizationMainCallsAPI.md#ParseTerraformVariablesFromGitRepo) | **Post** /organization/{organizationId}/parseTerraformVariablesFromGitRepo | Parse Terraform variables from Git repository



## CreateGitToken

> GitTokenResponse CreateGitToken(ctx, organizationId).GitTokenRequest(gitTokenRequest).Execute()

Create a git token



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
	gitTokenRequest := *openapiclient.NewGitTokenRequest("Name_example", openapiclient.GitProviderEnum("BITBUCKET"), "Token_example") // GitTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationMainCallsAPI.CreateGitToken(context.Background(), organizationId).GitTokenRequest(gitTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.CreateGitToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGitToken`: GitTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.CreateGitToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGitTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gitTokenRequest** | [**GitTokenRequest**](GitTokenRequest.md) |  | 

### Return type

[**GitTokenResponse**](GitTokenResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganization

> Organization CreateOrganization(ctx).OrganizationRequest(organizationRequest).Execute()

Create an organization

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
	organizationRequest := *openapiclient.NewOrganizationRequest("Name_example", openapiclient.PlanEnum("FREE")) // OrganizationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationMainCallsAPI.CreateOrganization(context.Background()).OrganizationRequest(organizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.CreateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationRequest** | [**OrganizationRequest**](OrganizationRequest.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGitToken

> DeleteGitToken(ctx, organizationId, gitTokenId).Execute()

Delete a git token

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
	gitTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Git Token ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationMainCallsAPI.DeleteGitToken(context.Background(), organizationId, gitTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.DeleteGitToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**gitTokenId** | **string** | Git Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## DeleteOrganization

> DeleteOrganization(ctx, organizationId).Execute()

Delete an organization



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
	r, err := apiClient.OrganizationMainCallsAPI.DeleteOrganization(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.DeleteOrganization``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## EditGitToken

> GitTokenResponse EditGitToken(ctx, organizationId, gitTokenId).GitTokenRequest(gitTokenRequest).Execute()

Edit a git token

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
	gitTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Git Token ID
	gitTokenRequest := *openapiclient.NewGitTokenRequest("Name_example", openapiclient.GitProviderEnum("BITBUCKET"), "Token_example") // GitTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationMainCallsAPI.EditGitToken(context.Background(), organizationId, gitTokenId).GitTokenRequest(gitTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.EditGitToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditGitToken`: GitTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.EditGitToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**gitTokenId** | **string** | Git Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditGitTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gitTokenRequest** | [**GitTokenRequest**](GitTokenRequest.md) |  | 

### Return type

[**GitTokenResponse**](GitTokenResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOrganization

> Organization EditOrganization(ctx, organizationId).OrganizationEditRequest(organizationEditRequest).Execute()

Edit an organization



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
	organizationEditRequest := *openapiclient.NewOrganizationEditRequest("Name_example") // OrganizationEditRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationMainCallsAPI.EditOrganization(context.Background(), organizationId).OrganizationEditRequest(organizationEditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.EditOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.EditOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationEditRequest** | [**OrganizationEditRequest**](OrganizationEditRequest.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerRegistryAssociatedServices

> ContainerRegistryAssociatedServicesResponseList GetContainerRegistryAssociatedServices(ctx, organizationId, containerRegistryId).Execute()

Get organization container registry associated services



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
	organizationId := "organizationId_example" // string | 
	containerRegistryId := "containerRegistryId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationMainCallsAPI.GetContainerRegistryAssociatedServices(context.Background(), organizationId, containerRegistryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.GetContainerRegistryAssociatedServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerRegistryAssociatedServices`: ContainerRegistryAssociatedServicesResponseList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.GetContainerRegistryAssociatedServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**containerRegistryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerRegistryAssociatedServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContainerRegistryAssociatedServicesResponseList**](ContainerRegistryAssociatedServicesResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitTokenAssociatedServices

> GitTokenAssociatedServicesResponseList GetGitTokenAssociatedServices(ctx, organizationId, gitTokenId).Execute()

Get organization git token associated services



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
	gitTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Git Token ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationMainCallsAPI.GetGitTokenAssociatedServices(context.Background(), organizationId, gitTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.GetGitTokenAssociatedServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitTokenAssociatedServices`: GitTokenAssociatedServicesResponseList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.GetGitTokenAssociatedServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**gitTokenId** | **string** | Git Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitTokenAssociatedServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GitTokenAssociatedServicesResponseList**](GitTokenAssociatedServicesResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHelmRepositoryAssociatedServices

> HelmRepositoryAssociatedServicesResponseList GetHelmRepositoryAssociatedServices(ctx, organizationId, helmRepositoryId).Execute()

Get organization helm repository associated services



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
	organizationId := "organizationId_example" // string | 
	helmRepositoryId := "helmRepositoryId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationMainCallsAPI.GetHelmRepositoryAssociatedServices(context.Background(), organizationId, helmRepositoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.GetHelmRepositoryAssociatedServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHelmRepositoryAssociatedServices`: HelmRepositoryAssociatedServicesResponseList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.GetHelmRepositoryAssociatedServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**helmRepositoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHelmRepositoryAssociatedServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HelmRepositoryAssociatedServicesResponseList**](HelmRepositoryAssociatedServicesResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> Organization GetOrganization(ctx, organizationId).Execute()

Get organization by ID

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
	resp, r, err := apiClient.OrganizationMainCallsAPI.GetOrganization(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.GetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organization**](Organization.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationGitToken

> GitTokenResponse GetOrganizationGitToken(ctx, organizationId, gitTokenId).Execute()

Get organization git token



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
	gitTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Git Token ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationMainCallsAPI.GetOrganizationGitToken(context.Background(), organizationId, gitTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.GetOrganizationGitToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationGitToken`: GitTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.GetOrganizationGitToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**gitTokenId** | **string** | Git Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationGitTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GitTokenResponse**](GitTokenResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganization

> OrganizationResponseList ListOrganization(ctx).Execute()

List user organizations

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
	resp, r, err := apiClient.OrganizationMainCallsAPI.ListOrganization(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.ListOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganization`: OrganizationResponseList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.ListOrganization`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationRequest struct via the builder pattern


### Return type

[**OrganizationResponseList**](OrganizationResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationAvailableRoles

> OrganizationAvailableRoleList ListOrganizationAvailableRoles(ctx, organizationId).Execute()

List organization available roles



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
	resp, r, err := apiClient.OrganizationMainCallsAPI.ListOrganizationAvailableRoles(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.ListOrganizationAvailableRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationAvailableRoles`: OrganizationAvailableRoleList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.ListOrganizationAvailableRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationAvailableRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationAvailableRoleList**](OrganizationAvailableRoleList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationCredentials

> OrganizationCrendentialsResponseList ListOrganizationCredentials(ctx, organizationId).Execute()

List credentials of an organization and their associated clusters



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
	organizationId := "organizationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationMainCallsAPI.ListOrganizationCredentials(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.ListOrganizationCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationCredentials`: OrganizationCrendentialsResponseList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.ListOrganizationCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationCrendentialsResponseList**](OrganizationCrendentialsResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationGitTokens

> GitTokenResponseList ListOrganizationGitTokens(ctx, organizationId).Execute()

List organization git tokens



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
	resp, r, err := apiClient.OrganizationMainCallsAPI.ListOrganizationGitTokens(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.ListOrganizationGitTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationGitTokens`: GitTokenResponseList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.ListOrganizationGitTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationGitTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitTokenResponseList**](GitTokenResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServicesByOrganizationId

> ListServicesByOrganizationId200Response ListServicesByOrganizationId(ctx, organizationId).ProjectId(projectId).EnvironmentId(environmentId).ClusterId(clusterId).Execute()

List Services By OrganizationId

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
	organizationId := "organizationId_example" // string | 
	projectId := "projectId_example" // string |  (optional)
	environmentId := "environmentId_example" // string |  (optional)
	clusterId := "clusterId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationMainCallsAPI.ListServicesByOrganizationId(context.Background(), organizationId).ProjectId(projectId).EnvironmentId(environmentId).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.ListServicesByOrganizationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServicesByOrganizationId`: ListServicesByOrganizationId200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.ListServicesByOrganizationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServicesByOrganizationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectId** | **string** |  | 
 **environmentId** | **string** |  | 
 **clusterId** | **string** |  | 

### Return type

[**ListServicesByOrganizationId200Response**](ListServicesByOrganizationId200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTfVarsFilesFromGitRepo

> ListTfVarsFilesFromGitRepo200Response ListTfVarsFilesFromGitRepo(ctx, organizationId).TfVarsListRequest(tfVarsListRequest).Execute()

List Terraform tfvars files from Git repository

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
	tfVarsListRequest := *openapiclient.NewTfVarsListRequest(*openapiclient.NewGitRepositoryRequest("https://github.com/Qovery/terraform-examples", *openapiclient.NewGitProvider("Kind_example"))) // TfVarsListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationMainCallsAPI.ListTfVarsFilesFromGitRepo(context.Background(), organizationId).TfVarsListRequest(tfVarsListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.ListTfVarsFilesFromGitRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTfVarsFilesFromGitRepo`: ListTfVarsFilesFromGitRepo200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.ListTfVarsFilesFromGitRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTfVarsFilesFromGitRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tfVarsListRequest** | [**TfVarsListRequest**](TfVarsListRequest.md) |  | 

### Return type

[**ListTfVarsFilesFromGitRepo200Response**](ListTfVarsFilesFromGitRepo200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParseTerraformVariablesFromGitRepo

> ParseTerraformVariablesFromGitRepo200Response ParseTerraformVariablesFromGitRepo(ctx, organizationId).TerraformVariableParsingRequest(terraformVariableParsingRequest).Execute()

Parse Terraform variables from Git repository

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
	terraformVariableParsingRequest := *openapiclient.NewTerraformVariableParsingRequest(*openapiclient.NewApplicationGitRepositoryRequest("https://github.com/Qovery/simple-node-app", openapiclient.GitProviderEnum("BITBUCKET"))) // TerraformVariableParsingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationMainCallsAPI.ParseTerraformVariablesFromGitRepo(context.Background(), organizationId).TerraformVariableParsingRequest(terraformVariableParsingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationMainCallsAPI.ParseTerraformVariablesFromGitRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParseTerraformVariablesFromGitRepo`: ParseTerraformVariablesFromGitRepo200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationMainCallsAPI.ParseTerraformVariablesFromGitRepo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiParseTerraformVariablesFromGitRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **terraformVariableParsingRequest** | [**TerraformVariableParsingRequest**](TerraformVariableParsingRequest.md) |  | 

### Return type

[**ParseTerraformVariablesFromGitRepo200Response**](ParseTerraformVariablesFromGitRepo200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

