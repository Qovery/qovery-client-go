# \CloudProviderCredentialsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAWSCredentials**](CloudProviderCredentialsAPI.md#CreateAWSCredentials) | **Post** /organization/{organizationId}/aws/credentials | Create AWS credentials set
[**CreateDOCredentials**](CloudProviderCredentialsAPI.md#CreateDOCredentials) | **Post** /organization/{organizationId}/digitalOcean/credentials | Create Digital Ocean credentials set
[**CreateGCPCredentials**](CloudProviderCredentialsAPI.md#CreateGCPCredentials) | **Post** /organization/{organizationId}/gcp/credentials | Create GCP credentials set
[**CreateScalewayCredentials**](CloudProviderCredentialsAPI.md#CreateScalewayCredentials) | **Post** /organization/{organizationId}/scaleway/credentials | Create Scaleway credentials set
[**DeleteAWSCredentials**](CloudProviderCredentialsAPI.md#DeleteAWSCredentials) | **Delete** /organization/{organizationId}/aws/credentials/{credentialsId} | Delete a set of AWS credentials
[**DeleteDOCredentials**](CloudProviderCredentialsAPI.md#DeleteDOCredentials) | **Delete** /organization/{organizationId}/digitalOcean/credentials/{credentialsId} | Delete a set of Digital Ocean credentials
[**DeleteGcpCredentials**](CloudProviderCredentialsAPI.md#DeleteGcpCredentials) | **Delete** /organization/{organizationId}/gcp/credentials/{credentialsId} | Delete a set of GCP credentials
[**DeleteScalewayCredentials**](CloudProviderCredentialsAPI.md#DeleteScalewayCredentials) | **Delete** /organization/{organizationId}/scaleway/credentials/{credentialsId} | Delete a set of Scaleway credentials
[**EditAWSCredentials**](CloudProviderCredentialsAPI.md#EditAWSCredentials) | **Put** /organization/{organizationId}/aws/credentials/{credentialsId} | Edit a set of AWS credentials
[**EditDOCredentials**](CloudProviderCredentialsAPI.md#EditDOCredentials) | **Put** /organization/{organizationId}/digitalOcean/credentials/{credentialsId} | Edit a set of Digital Ocean credentials
[**EditGcpCredentials**](CloudProviderCredentialsAPI.md#EditGcpCredentials) | **Put** /organization/{organizationId}/gcp/credentials/{credentialsId} | Edit a set of GCP credentials
[**EditScalewayCredentials**](CloudProviderCredentialsAPI.md#EditScalewayCredentials) | **Put** /organization/{organizationId}/scaleway/credentials/{credentialsId} | Edit a set of Scaleway credentials
[**GetAWSCredentials**](CloudProviderCredentialsAPI.md#GetAWSCredentials) | **Get** /organization/{organizationId}/aws/credentials/{credentialsId} | Get a set of AWS credentials
[**GetDOCredentials**](CloudProviderCredentialsAPI.md#GetDOCredentials) | **Get** /organization/{organizationId}/digitalOcean/credentials/{credentialsId} | Get a set of Digital Ocean credentials
[**GetGcpCredentials**](CloudProviderCredentialsAPI.md#GetGcpCredentials) | **Get** /organization/{organizationId}/gcp/credentials/{credentialsId} | Get a set of GCP credentials
[**GetScalewayCredentials**](CloudProviderCredentialsAPI.md#GetScalewayCredentials) | **Get** /organization/{organizationId}/scaleway/credentials/{credentialsId} | Get a set of Scaleway credentials
[**ListAWSCredentials**](CloudProviderCredentialsAPI.md#ListAWSCredentials) | **Get** /organization/{organizationId}/aws/credentials | List AWS credentials
[**ListDOCredentials**](CloudProviderCredentialsAPI.md#ListDOCredentials) | **Get** /organization/{organizationId}/digitalOcean/credentials | List DO credentials
[**ListGcpCredentials**](CloudProviderCredentialsAPI.md#ListGcpCredentials) | **Get** /organization/{organizationId}/gcp/credentials | List GCP credentials
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
    awsCredentialsRequest := *openapiclient.NewAwsCredentialsRequest("Name_example", "AccessKeyId_example", "SecretAccessKey_example") // AwsCredentialsRequest |  (optional)

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


## CreateDOCredentials

> ClusterCredentials CreateDOCredentials(ctx, organizationId).DoCredentialsRequest(doCredentialsRequest).Execute()

Create Digital Ocean credentials set

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
    doCredentialsRequest := *openapiclient.NewDoCredentialsRequest("Name_example") // DoCredentialsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProviderCredentialsAPI.CreateDOCredentials(context.Background(), organizationId).DoCredentialsRequest(doCredentialsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.CreateDOCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDOCredentials`: ClusterCredentials
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.CreateDOCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDOCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **doCredentialsRequest** | [**DoCredentialsRequest**](DoCredentialsRequest.md) |  | 

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


## CreateGCPCredentials

> ClusterCredentials CreateGCPCredentials(ctx, organizationId).GcpCredentialsRequest(gcpCredentialsRequest).Execute()

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
    gcpCredentialsRequest := *openapiclient.NewGcpCredentialsRequest("Name_example", "JsonCredentials_example") // GcpCredentialsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProviderCredentialsAPI.CreateGCPCredentials(context.Background(), organizationId).GcpCredentialsRequest(gcpCredentialsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.CreateGCPCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGCPCredentials`: ClusterCredentials
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.CreateGCPCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGCPCredentialsRequest struct via the builder pattern


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


## DeleteDOCredentials

> DeleteDOCredentials(ctx, credentialsId, organizationId).Execute()

Delete a set of Digital Ocean credentials

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
    r, err := apiClient.CloudProviderCredentialsAPI.DeleteDOCredentials(context.Background(), credentialsId, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.DeleteDOCredentials``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDOCredentialsRequest struct via the builder pattern


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
    awsCredentialsRequest := *openapiclient.NewAwsCredentialsRequest("Name_example", "AccessKeyId_example", "SecretAccessKey_example") // AwsCredentialsRequest |  (optional)

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


## EditDOCredentials

> ClusterCredentials EditDOCredentials(ctx, organizationId, credentialsId).DoCredentialsRequest(doCredentialsRequest).Execute()

Edit a set of Digital Ocean credentials

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
    doCredentialsRequest := *openapiclient.NewDoCredentialsRequest("Name_example") // DoCredentialsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProviderCredentialsAPI.EditDOCredentials(context.Background(), organizationId, credentialsId).DoCredentialsRequest(doCredentialsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.EditDOCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditDOCredentials`: ClusterCredentials
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.EditDOCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditDOCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **doCredentialsRequest** | [**DoCredentialsRequest**](DoCredentialsRequest.md) |  | 

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
    gcpCredentialsRequest := *openapiclient.NewGcpCredentialsRequest("Name_example", "JsonCredentials_example") // GcpCredentialsRequest |  (optional)

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


## GetDOCredentials

> ClusterCredentials GetDOCredentials(ctx, organizationId, credentialsId).Execute()

Get a set of Digital Ocean credentials

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
    resp, r, err := apiClient.CloudProviderCredentialsAPI.GetDOCredentials(context.Background(), organizationId, credentialsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.GetDOCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDOCredentials`: ClusterCredentials
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.GetDOCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDOCredentialsRequest struct via the builder pattern


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


## ListDOCredentials

> ClusterCredentialsResponseList ListDOCredentials(ctx, organizationId).Execute()

List DO credentials

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
    resp, r, err := apiClient.CloudProviderCredentialsAPI.ListDOCredentials(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.ListDOCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDOCredentials`: ClusterCredentialsResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.ListDOCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDOCredentialsRequest struct via the builder pattern


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

