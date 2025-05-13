# \CloudProviderCredentialsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAWSCredentials**](CloudProviderCredentialsAPI.md#CreateAWSCredentials) | **Post** /organization/{organizationId}/aws/credentials | Create AWS credentials set
[**CreateAzureCredentials**](CloudProviderCredentialsAPI.md#CreateAzureCredentials) | **Post** /organization/{organizationId}/azure/credentials | Create Azure credentials set
[**CreateGcpCredentials**](CloudProviderCredentialsAPI.md#CreateGcpCredentials) | **Post** /organization/{organizationId}/gcp/credentials | Create GCP credentials set
[**CreateOnPremiseCredentials**](CloudProviderCredentialsAPI.md#CreateOnPremiseCredentials) | **Post** /organization/{organizationId}/onPremise/credentials | Create OnPremise credentials set
[**CreateScalewayCredentials**](CloudProviderCredentialsAPI.md#CreateScalewayCredentials) | **Post** /organization/{organizationId}/scaleway/credentials | Create Scaleway credentials set
[**DeleteAWSCredentials**](CloudProviderCredentialsAPI.md#DeleteAWSCredentials) | **Delete** /organization/{organizationId}/aws/credentials/{credentialsId} | Delete a set of AWS credentials
[**DeleteAzureCredentials**](CloudProviderCredentialsAPI.md#DeleteAzureCredentials) | **Delete** /organization/{organizationId}/azure/credentials/{credentialsId} | Delete a set of Azure credentials
[**DeleteGcpCredentials**](CloudProviderCredentialsAPI.md#DeleteGcpCredentials) | **Delete** /organization/{organizationId}/gcp/credentials/{credentialsId} | Delete a set of GCP credentials
[**DeleteOnPremiseCredentials**](CloudProviderCredentialsAPI.md#DeleteOnPremiseCredentials) | **Delete** /organization/{organizationId}/onPremise/credentials/{credentialsId} | Delete a set of OnPremise credentials
[**DeleteScalewayCredentials**](CloudProviderCredentialsAPI.md#DeleteScalewayCredentials) | **Delete** /organization/{organizationId}/scaleway/credentials/{credentialsId} | Delete a set of Scaleway credentials
[**EditAWSCredentials**](CloudProviderCredentialsAPI.md#EditAWSCredentials) | **Put** /organization/{organizationId}/aws/credentials/{credentialsId} | Edit a set of AWS credentials
[**EditAzureCredentials**](CloudProviderCredentialsAPI.md#EditAzureCredentials) | **Put** /organization/{organizationId}/azure/credentials/{credentialsId} | Edit a set of Azure credentials
[**EditGcpCredentials**](CloudProviderCredentialsAPI.md#EditGcpCredentials) | **Put** /organization/{organizationId}/gcp/credentials/{credentialsId} | Edit a set of GCP credentials
[**EditOnPremiseCredentials**](CloudProviderCredentialsAPI.md#EditOnPremiseCredentials) | **Put** /organization/{organizationId}/onPremise/credentials/{credentialsId} | Edit a set of OnPremise credentials
[**EditScalewayCredentials**](CloudProviderCredentialsAPI.md#EditScalewayCredentials) | **Put** /organization/{organizationId}/scaleway/credentials/{credentialsId} | Edit a set of Scaleway credentials
[**GetAWSCredentials**](CloudProviderCredentialsAPI.md#GetAWSCredentials) | **Get** /organization/{organizationId}/aws/credentials/{credentialsId} | Get a set of AWS credentials
[**GetAzureCredentials**](CloudProviderCredentialsAPI.md#GetAzureCredentials) | **Get** /organization/{organizationId}/azure/credentials/{credentialsId} | Get a set of Azure credentials
[**GetGcpCredentials**](CloudProviderCredentialsAPI.md#GetGcpCredentials) | **Get** /organization/{organizationId}/gcp/credentials/{credentialsId} | Get a set of GCP credentials
[**GetOnPremiseCredentials**](CloudProviderCredentialsAPI.md#GetOnPremiseCredentials) | **Get** /organization/{organizationId}/onPremise/credentials/{credentialsId} | Get a set of OnPremise credentials
[**GetScalewayCredentials**](CloudProviderCredentialsAPI.md#GetScalewayCredentials) | **Get** /organization/{organizationId}/scaleway/credentials/{credentialsId} | Get a set of Scaleway credentials
[**ListAWSCredentials**](CloudProviderCredentialsAPI.md#ListAWSCredentials) | **Get** /organization/{organizationId}/aws/credentials | List AWS credentials
[**ListAzureCredentials**](CloudProviderCredentialsAPI.md#ListAzureCredentials) | **Get** /organization/{organizationId}/azure/credentials | List Azure credentials
[**ListGcpCredentials**](CloudProviderCredentialsAPI.md#ListGcpCredentials) | **Get** /organization/{organizationId}/gcp/credentials | List GCP credentials
[**ListOnPremiseCredentials**](CloudProviderCredentialsAPI.md#ListOnPremiseCredentials) | **Get** /organization/{organizationId}/onPremise/credentials | List OnPremise credentials
[**ListScalewayCredentials**](CloudProviderCredentialsAPI.md#ListScalewayCredentials) | **Get** /organization/{organizationId}/scaleway/credentials | List Scaleway credentials



## CreateAWSCredentials

> ClusterCredentials CreateAWSCredentials(ctx, organizationId).AwsCredentialsRequest(awsCredentialsRequest).Execute()

Create AWS credentials set

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
	awsCredentialsRequest := openapiclient.AwsCredentialsRequest{AwsRoleCredentialsRequest: openapiclient.NewAwsRoleCredentialsRequest("Name_example", "RoleArn_example")} // AwsCredentialsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.CreateAWSCredentials(context.Background(), organizationId).AwsCredentialsRequest(awsCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.CreateAWSCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAWSCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.CreateAWSCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAWSCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsCredentialsRequest** | [**AwsCredentialsRequest**](AwsCredentialsRequest.md) |  | 

### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAzureCredentials

> ClusterCredentials CreateAzureCredentials(ctx, organizationId).AzureCredentialsRequest(azureCredentialsRequest).Execute()

Create Azure credentials set

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
	azureCredentialsRequest := *openapiclient.NewAzureCredentialsRequest("Name_example", "AzureSubscriptionId_example", "AzureTenantId_example", "AzureClientId_example", "AzureClientSecret_example", "AzureResourceGroupName_example") // AzureCredentialsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.CreateAzureCredentials(context.Background(), organizationId).AzureCredentialsRequest(azureCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.CreateAzureCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAzureCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.CreateAzureCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAzureCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureCredentialsRequest** | [**AzureCredentialsRequest**](AzureCredentialsRequest.md) |  | 

### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGcpCredentials

> ClusterCredentials CreateGcpCredentials(ctx, organizationId).GcpCredentialsRequest(gcpCredentialsRequest).Execute()

Create GCP credentials set

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
	gcpCredentialsRequest := *openapiclient.NewGcpCredentialsRequest("Name_example", "GcpCredentials_example") // GcpCredentialsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.CreateGcpCredentials(context.Background(), organizationId).GcpCredentialsRequest(gcpCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.CreateGcpCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGcpCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.CreateGcpCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGcpCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpCredentialsRequest** | [**GcpCredentialsRequest**](GcpCredentialsRequest.md) |  | 

### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOnPremiseCredentials

> ClusterCredentials CreateOnPremiseCredentials(ctx, organizationId).OnPremiseCredentialsRequest(onPremiseCredentialsRequest).Execute()

Create OnPremise credentials set

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
	onPremiseCredentialsRequest := *openapiclient.NewOnPremiseCredentialsRequest("Name_example") // OnPremiseCredentialsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.CreateOnPremiseCredentials(context.Background(), organizationId).OnPremiseCredentialsRequest(onPremiseCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.CreateOnPremiseCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOnPremiseCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.CreateOnPremiseCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOnPremiseCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onPremiseCredentialsRequest** | [**OnPremiseCredentialsRequest**](OnPremiseCredentialsRequest.md) |  | 

### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScalewayCredentials

> ClusterCredentials CreateScalewayCredentials(ctx, organizationId).ScalewayCredentialsRequest(scalewayCredentialsRequest).Execute()

Create Scaleway credentials set

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
	scalewayCredentialsRequest := *openapiclient.NewScalewayCredentialsRequest("Name_example", "ScalewayAccessKey_example", "ScalewaySecretKey_example", "ScalewayProjectId_example", "ScalewayOrganizationId_example") // ScalewayCredentialsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.CreateScalewayCredentials(context.Background(), organizationId).ScalewayCredentialsRequest(scalewayCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.CreateScalewayCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateScalewayCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.CreateScalewayCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScalewayCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scalewayCredentialsRequest** | [**ScalewayCredentialsRequest**](ScalewayCredentialsRequest.md) |  | 

### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAWSCredentials

> DeleteAWSCredentials(ctx, credentialsId, organizationId).Execute()

Delete a set of AWS credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudProviderCredentialsAPI.DeleteAWSCredentials(context.Background(), credentialsId, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.DeleteAWSCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** | Credentials ID | 
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAWSCredentialsRequest struct via the builder pattern


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


## DeleteAzureCredentials

> DeleteAzureCredentials(ctx, organizationId, credentialsId).Execute()

Delete a set of Azure credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudProviderCredentialsAPI.DeleteAzureCredentials(context.Background(), organizationId, credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.DeleteAzureCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAzureCredentialsRequest struct via the builder pattern


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


## DeleteGcpCredentials

> DeleteGcpCredentials(ctx, credentialsId, organizationId).Execute()

Delete a set of GCP credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudProviderCredentialsAPI.DeleteGcpCredentials(context.Background(), credentialsId, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.DeleteGcpCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** | Credentials ID | 
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGcpCredentialsRequest struct via the builder pattern


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


## DeleteOnPremiseCredentials

> DeleteOnPremiseCredentials(ctx, credentialsId, organizationId).Execute()

Delete a set of OnPremise credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudProviderCredentialsAPI.DeleteOnPremiseCredentials(context.Background(), credentialsId, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.DeleteOnPremiseCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** | Credentials ID | 
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOnPremiseCredentialsRequest struct via the builder pattern


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


## DeleteScalewayCredentials

> DeleteScalewayCredentials(ctx, credentialsId, organizationId).Execute()

Delete a set of Scaleway credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudProviderCredentialsAPI.DeleteScalewayCredentials(context.Background(), credentialsId, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.DeleteScalewayCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** | Credentials ID | 
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScalewayCredentialsRequest struct via the builder pattern


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


## EditAWSCredentials

> ClusterCredentials EditAWSCredentials(ctx, organizationId, credentialsId).AwsCredentialsRequest(awsCredentialsRequest).Execute()

Edit a set of AWS credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
	awsCredentialsRequest := openapiclient.AwsCredentialsRequest{AwsRoleCredentialsRequest: openapiclient.NewAwsRoleCredentialsRequest("Name_example", "RoleArn_example")} // AwsCredentialsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.EditAWSCredentials(context.Background(), organizationId, credentialsId).AwsCredentialsRequest(awsCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.EditAWSCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAWSCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.EditAWSCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAWSCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **awsCredentialsRequest** | [**AwsCredentialsRequest**](AwsCredentialsRequest.md) |  | 

### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditAzureCredentials

> ClusterCredentials EditAzureCredentials(ctx, organizationId, credentialsId).AzureCredentialsRequest(azureCredentialsRequest).Execute()

Edit a set of Azure credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
	azureCredentialsRequest := *openapiclient.NewAzureCredentialsRequest("Name_example", "AzureSubscriptionId_example", "AzureTenantId_example", "AzureClientId_example", "AzureClientSecret_example", "AzureResourceGroupName_example") // AzureCredentialsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.EditAzureCredentials(context.Background(), organizationId, credentialsId).AzureCredentialsRequest(azureCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.EditAzureCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAzureCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.EditAzureCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAzureCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **azureCredentialsRequest** | [**AzureCredentialsRequest**](AzureCredentialsRequest.md) |  | 

### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditGcpCredentials

> ClusterCredentials EditGcpCredentials(ctx, organizationId, credentialsId).GcpCredentialsRequest(gcpCredentialsRequest).Execute()

Edit a set of GCP credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
	gcpCredentialsRequest := *openapiclient.NewGcpCredentialsRequest("Name_example", "GcpCredentials_example") // GcpCredentialsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.EditGcpCredentials(context.Background(), organizationId, credentialsId).GcpCredentialsRequest(gcpCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.EditGcpCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditGcpCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.EditGcpCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditGcpCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gcpCredentialsRequest** | [**GcpCredentialsRequest**](GcpCredentialsRequest.md) |  | 

### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOnPremiseCredentials

> ClusterCredentials EditOnPremiseCredentials(ctx, organizationId, credentialsId).OnPremiseCredentialsRequest(onPremiseCredentialsRequest).Execute()

Edit a set of OnPremise credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
	onPremiseCredentialsRequest := *openapiclient.NewOnPremiseCredentialsRequest("Name_example") // OnPremiseCredentialsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.EditOnPremiseCredentials(context.Background(), organizationId, credentialsId).OnPremiseCredentialsRequest(onPremiseCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.EditOnPremiseCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOnPremiseCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.EditOnPremiseCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOnPremiseCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **onPremiseCredentialsRequest** | [**OnPremiseCredentialsRequest**](OnPremiseCredentialsRequest.md) |  | 

### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditScalewayCredentials

> ClusterCredentials EditScalewayCredentials(ctx, organizationId, credentialsId).ScalewayCredentialsRequest(scalewayCredentialsRequest).Execute()

Edit a set of Scaleway credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
	scalewayCredentialsRequest := *openapiclient.NewScalewayCredentialsRequest("Name_example", "ScalewayAccessKey_example", "ScalewaySecretKey_example", "ScalewayProjectId_example", "ScalewayOrganizationId_example") // ScalewayCredentialsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.EditScalewayCredentials(context.Background(), organizationId, credentialsId).ScalewayCredentialsRequest(scalewayCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.EditScalewayCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditScalewayCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.EditScalewayCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditScalewayCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scalewayCredentialsRequest** | [**ScalewayCredentialsRequest**](ScalewayCredentialsRequest.md) |  | 

### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAWSCredentials

> ClusterCredentials GetAWSCredentials(ctx, organizationId, credentialsId).Execute()

Get a set of AWS credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.GetAWSCredentials(context.Background(), organizationId, credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.GetAWSCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAWSCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.GetAWSCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAWSCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAzureCredentials

> ClusterCredentials GetAzureCredentials(ctx, organizationId, credentialsId).Execute()

Get a set of Azure credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.GetAzureCredentials(context.Background(), organizationId, credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.GetAzureCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAzureCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.GetAzureCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpCredentials

> ClusterCredentials GetGcpCredentials(ctx, organizationId, credentialsId).Execute()

Get a set of GCP credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.GetGcpCredentials(context.Background(), organizationId, credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.GetGcpCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGcpCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.GetGcpCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnPremiseCredentials

> ClusterCredentials GetOnPremiseCredentials(ctx, organizationId, credentialsId).Execute()

Get a set of OnPremise credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.GetOnPremiseCredentials(context.Background(), organizationId, credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.GetOnPremiseCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnPremiseCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.GetOnPremiseCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnPremiseCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScalewayCredentials

> ClusterCredentials GetScalewayCredentials(ctx, organizationId, credentialsId).Execute()

Get a set of Scaleway credentials

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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.GetScalewayCredentials(context.Background(), organizationId, credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.GetScalewayCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScalewayCredentials`: ClusterCredentials
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.GetScalewayCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScalewayCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterCredentials**](ClusterCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAWSCredentials

> ClusterCredentialsResponseList ListAWSCredentials(ctx, organizationId).Execute()

List AWS credentials

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
	resp, r, err := apiClient.CloudProviderCredentialsAPI.ListAWSCredentials(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.ListAWSCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAWSCredentials`: ClusterCredentialsResponseList
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.ListAWSCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAWSCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterCredentialsResponseList**](ClusterCredentialsResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAzureCredentials

> ClusterCredentialsResponseList ListAzureCredentials(ctx, organizationId).Execute()

List Azure credentials

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
	resp, r, err := apiClient.CloudProviderCredentialsAPI.ListAzureCredentials(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.ListAzureCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAzureCredentials`: ClusterCredentialsResponseList
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.ListAzureCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAzureCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterCredentialsResponseList**](ClusterCredentialsResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGcpCredentials

> ClusterCredentialsResponseList ListGcpCredentials(ctx, organizationId).Execute()

List GCP credentials

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
	resp, r, err := apiClient.CloudProviderCredentialsAPI.ListGcpCredentials(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.ListGcpCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGcpCredentials`: ClusterCredentialsResponseList
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.ListGcpCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGcpCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterCredentialsResponseList**](ClusterCredentialsResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOnPremiseCredentials

> ClusterCredentialsResponseList ListOnPremiseCredentials(ctx, organizationId).Execute()

List OnPremise credentials

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
	resp, r, err := apiClient.CloudProviderCredentialsAPI.ListOnPremiseCredentials(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.ListOnPremiseCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOnPremiseCredentials`: ClusterCredentialsResponseList
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.ListOnPremiseCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOnPremiseCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterCredentialsResponseList**](ClusterCredentialsResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScalewayCredentials

> ClusterCredentialsResponseList ListScalewayCredentials(ctx, organizationId).Execute()

List Scaleway credentials

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
	resp, r, err := apiClient.CloudProviderCredentialsAPI.ListScalewayCredentials(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.ListScalewayCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListScalewayCredentials`: ClusterCredentialsResponseList
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.ListScalewayCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListScalewayCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterCredentialsResponseList**](ClusterCredentialsResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

