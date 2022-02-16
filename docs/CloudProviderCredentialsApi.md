# \CloudProviderCredentialsApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAWSCredentials**](CloudProviderCredentialsApi.md#CreateAWSCredentials) | **Post** /organization/{organizationId}/aws/credentials | Create AWS credentials set
[**CreateDOCredentials**](CloudProviderCredentialsApi.md#CreateDOCredentials) | **Post** /organization/{organizationId}/digitalOcean/credentials | Create Digital Ocean credentials set
[**CreateScalewayCredentials**](CloudProviderCredentialsApi.md#CreateScalewayCredentials) | **Post** /organization/{organizationId}/scaleway/credentials | Create Scaleway credentials set
[**DeleteAWSCredentials**](CloudProviderCredentialsApi.md#DeleteAWSCredentials) | **Delete** /organization/{organizationId}/aws/credentials/{credentialsId} | Delete a set of AWS credentials
[**DeleteDOCredentials**](CloudProviderCredentialsApi.md#DeleteDOCredentials) | **Delete** /organization/{organizationId}/digitalOcean/credentials/{credentialsId} | Delete a set of Digital Ocean credentials
[**DeleteScalewayCredentials**](CloudProviderCredentialsApi.md#DeleteScalewayCredentials) | **Delete** /organization/{organizationId}/scaleway/credentials/{credentialsId} | Delete a set of Scaleway credentials
[**EditAWSCredentials**](CloudProviderCredentialsApi.md#EditAWSCredentials) | **Put** /organization/{organizationId}/aws/credentials/{credentialsId} | Edit a set of AWS credentials
[**EditDOCredentials**](CloudProviderCredentialsApi.md#EditDOCredentials) | **Put** /organization/{organizationId}/digitalOcean/credentials/{credentialsId} | Edit a set of Digital Ocean credentials
[**EditScalewayCredentials**](CloudProviderCredentialsApi.md#EditScalewayCredentials) | **Put** /organization/{organizationId}/scaleway/credentials/{credentialsId} | Edit a set of Scaleway credentials
[**ListAWSCredentials**](CloudProviderCredentialsApi.md#ListAWSCredentials) | **Get** /organization/{organizationId}/aws/credentials | List AWS credentials
[**ListDOCredentials**](CloudProviderCredentialsApi.md#ListDOCredentials) | **Get** /organization/{organizationId}/digitalOcean/credentials | List DO credentials
[**ListScalewayCredentials**](CloudProviderCredentialsApi.md#ListScalewayCredentials) | **Get** /organization/{organizationId}/scaleway/credentials | List Scaleway credentials



## CreateAWSCredentials

> ClusterCredentialsResponse CreateAWSCredentials(ctx, organizationId).AwsCredentialsRequest(awsCredentialsRequest).Execute()

Create AWS credentials set

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    awsCredentialsRequest := *openapiclient.NewAwsCredentialsRequest("Name_example") // AwsCredentialsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderCredentialsApi.CreateAWSCredentials(context.Background(), organizationId).AwsCredentialsRequest(awsCredentialsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsApi.CreateAWSCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAWSCredentials`: ClusterCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsApi.CreateAWSCredentials`: %v\n", resp)
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

[**ClusterCredentialsResponse**](ClusterCredentialsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDOCredentials

> ClusterCredentialsResponse CreateDOCredentials(ctx, organizationId).DoCredentialsRequest(doCredentialsRequest).Execute()

Create Digital Ocean credentials set

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    doCredentialsRequest := *openapiclient.NewDoCredentialsRequest("Name_example") // DoCredentialsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderCredentialsApi.CreateDOCredentials(context.Background(), organizationId).DoCredentialsRequest(doCredentialsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsApi.CreateDOCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDOCredentials`: ClusterCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsApi.CreateDOCredentials`: %v\n", resp)
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

[**ClusterCredentialsResponse**](ClusterCredentialsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScalewayCredentials

> ClusterCredentialsResponse CreateScalewayCredentials(ctx, organizationId).ScalewayCredentialsRequest(scalewayCredentialsRequest).Execute()

Create Scaleway credentials set

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    scalewayCredentialsRequest := *openapiclient.NewScalewayCredentialsRequest("Name_example") // ScalewayCredentialsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderCredentialsApi.CreateScalewayCredentials(context.Background(), organizationId).ScalewayCredentialsRequest(scalewayCredentialsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsApi.CreateScalewayCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateScalewayCredentials`: ClusterCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsApi.CreateScalewayCredentials`: %v\n", resp)
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

[**ClusterCredentialsResponse**](ClusterCredentialsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
    openapiclient "./openapi"
)

func main() {
    credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderCredentialsApi.DeleteAWSCredentials(context.Background(), credentialsId, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsApi.DeleteAWSCredentials``: %v\n", err)
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

[bearerAuth](../README.md#bearerAuth)

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
    openapiclient "./openapi"
)

func main() {
    credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderCredentialsApi.DeleteDOCredentials(context.Background(), credentialsId, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsApi.DeleteDOCredentials``: %v\n", err)
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

[bearerAuth](../README.md#bearerAuth)

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
    openapiclient "./openapi"
)

func main() {
    credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderCredentialsApi.DeleteScalewayCredentials(context.Background(), credentialsId, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsApi.DeleteScalewayCredentials``: %v\n", err)
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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditAWSCredentials

> ClusterCredentialsResponse EditAWSCredentials(ctx, organizationId, credentialsId).AwsCredentialsRequest(awsCredentialsRequest).Execute()

Edit a set of AWS credentials

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
    awsCredentialsRequest := *openapiclient.NewAwsCredentialsRequest("Name_example") // AwsCredentialsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderCredentialsApi.EditAWSCredentials(context.Background(), organizationId, credentialsId).AwsCredentialsRequest(awsCredentialsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsApi.EditAWSCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditAWSCredentials`: ClusterCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsApi.EditAWSCredentials`: %v\n", resp)
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

[**ClusterCredentialsResponse**](ClusterCredentialsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditDOCredentials

> ClusterCredentialsResponse EditDOCredentials(ctx, organizationId, credentialsId).DoCredentialsRequest(doCredentialsRequest).Execute()

Edit a set of Digital Ocean credentials

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
    doCredentialsRequest := *openapiclient.NewDoCredentialsRequest("Name_example") // DoCredentialsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderCredentialsApi.EditDOCredentials(context.Background(), organizationId, credentialsId).DoCredentialsRequest(doCredentialsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsApi.EditDOCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditDOCredentials`: ClusterCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsApi.EditDOCredentials`: %v\n", resp)
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

[**ClusterCredentialsResponse**](ClusterCredentialsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditScalewayCredentials

> ClusterCredentialsResponse EditScalewayCredentials(ctx, organizationId, credentialsId).ScalewayCredentialsRequest(scalewayCredentialsRequest).Execute()

Edit a set of Scaleway credentials

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credentials ID
    scalewayCredentialsRequest := *openapiclient.NewScalewayCredentialsRequest("Name_example") // ScalewayCredentialsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderCredentialsApi.EditScalewayCredentials(context.Background(), organizationId, credentialsId).ScalewayCredentialsRequest(scalewayCredentialsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsApi.EditScalewayCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditScalewayCredentials`: ClusterCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsApi.EditScalewayCredentials`: %v\n", resp)
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

[**ClusterCredentialsResponse**](ClusterCredentialsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderCredentialsApi.ListAWSCredentials(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsApi.ListAWSCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSCredentials`: ClusterCredentialsResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsApi.ListAWSCredentials`: %v\n", resp)
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

[bearerAuth](../README.md#bearerAuth)

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
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderCredentialsApi.ListDOCredentials(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsApi.ListDOCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDOCredentials`: ClusterCredentialsResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsApi.ListDOCredentials`: %v\n", resp)
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

[bearerAuth](../README.md#bearerAuth)

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
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderCredentialsApi.ListScalewayCredentials(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsApi.ListScalewayCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScalewayCredentials`: ClusterCredentialsResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsApi.ListScalewayCredentials`: %v\n", resp)
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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

