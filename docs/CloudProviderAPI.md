# \CloudProviderAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAWSEKSInstanceType**](CloudProviderAPI.md#ListAWSEKSInstanceType) | **Get** /aws/eks/instanceType/{region} | List AWS EKS available instance types
[**ListAWSEc2InstanceType**](CloudProviderAPI.md#ListAWSEc2InstanceType) | **Get** /aws/ec2/instanceType/{region} | List AWS EC2 available instance types
[**ListAWSFeatures**](CloudProviderAPI.md#ListAWSFeatures) | **Get** /aws/clusterFeature | List AWS features available
[**ListAWSInstanceType**](CloudProviderAPI.md#ListAWSInstanceType) | **Get** /aws/instanceType | List AWS available instance types
[**ListAWSManagedDatabaseInstanceType**](CloudProviderAPI.md#ListAWSManagedDatabaseInstanceType) | **Get** /aws/managedDatabase/instanceType/{region}/{databaseType} | List AWS available managed database instance types
[**ListAWSManagedDatabaseType**](CloudProviderAPI.md#ListAWSManagedDatabaseType) | **Get** /aws/managedDatabase/type | List AWS available managed database types
[**ListAWSRegions**](CloudProviderAPI.md#ListAWSRegions) | **Get** /aws/region | List AWS regions
[**ListCloudProvider**](CloudProviderAPI.md#ListCloudProvider) | **Get** /cloudProvider | List Cloud providers available
[**ListDOFeatures**](CloudProviderAPI.md#ListDOFeatures) | **Get** /digitalOcean/clusterFeature | List DO features available
[**ListDOInstanceType**](CloudProviderAPI.md#ListDOInstanceType) | **Get** /digitalOcean/instanceType | List DO available instance types
[**ListDOManagedDatabaseInstanceType**](CloudProviderAPI.md#ListDOManagedDatabaseInstanceType) | **Get** /digitalOcean/managedDatabase/instanceType/{region}/{databaseType} | List Digital Ocean available managed database instance types
[**ListDOManagedDatabaseType**](CloudProviderAPI.md#ListDOManagedDatabaseType) | **Get** /digitalOcean/managedDatabase/type | List Digital Ocean available managed database types
[**ListDORegions**](CloudProviderAPI.md#ListDORegions) | **Get** /digitalOcean/region | List DO regions
[**ListSCWManagedDatabaseInstanceType**](CloudProviderAPI.md#ListSCWManagedDatabaseInstanceType) | **Get** /scaleway/managedDatabase/instanceType/{zone}/{databaseType} | List Scaleway available managed database instance types
[**ListSCWManagedDatabaseType**](CloudProviderAPI.md#ListSCWManagedDatabaseType) | **Get** /scaleway/managedDatabase/type | List Scaleway available managed database types
[**ListScalewayFeatures**](CloudProviderAPI.md#ListScalewayFeatures) | **Get** /scaleway/clusterFeature | List Scaleway features available
[**ListScalewayInstanceType**](CloudProviderAPI.md#ListScalewayInstanceType) | **Get** /scaleway/instanceType | List Scaleway available instance types
[**ListScalewayKapsuleInstanceType**](CloudProviderAPI.md#ListScalewayKapsuleInstanceType) | **Get** /scaleway/instanceType/{zone} | List Scaleway Kapsule available instance types
[**ListScalewayRegions**](CloudProviderAPI.md#ListScalewayRegions) | **Get** /scaleway/region | List Scaleway regions



## ListAWSEKSInstanceType

> ClusterInstanceTypeResponseList ListAWSEKSInstanceType(ctx, region).Execute()

List AWS EKS available instance types

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
    region := "us-east-2" // string | region name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProviderAPI.ListAWSEKSInstanceType(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListAWSEKSInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSEKSInstanceType`: ClusterInstanceTypeResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListAWSEKSInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAWSEKSInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterInstanceTypeResponseList**](ClusterInstanceTypeResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAWSEc2InstanceType

> ClusterInstanceTypeResponseList ListAWSEc2InstanceType(ctx, region).Execute()

List AWS EC2 available instance types

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
    region := "us-east-2" // string | region name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProviderAPI.ListAWSEc2InstanceType(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListAWSEc2InstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSEc2InstanceType`: ClusterInstanceTypeResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListAWSEc2InstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAWSEc2InstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterInstanceTypeResponseList**](ClusterInstanceTypeResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAWSFeatures

> ClusterFeatureResponseList ListAWSFeatures(ctx).Execute()

List AWS features available

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
    resp, r, err := apiClient.CloudProviderAPI.ListAWSFeatures(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListAWSFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSFeatures`: ClusterFeatureResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListAWSFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAWSFeaturesRequest struct via the builder pattern


### Return type

[**ClusterFeatureResponseList**](ClusterFeatureResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAWSInstanceType

> ClusterInstanceTypeResponseList ListAWSInstanceType(ctx).Execute()

List AWS available instance types

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
    resp, r, err := apiClient.CloudProviderAPI.ListAWSInstanceType(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListAWSInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSInstanceType`: ClusterInstanceTypeResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListAWSInstanceType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAWSInstanceTypeRequest struct via the builder pattern


### Return type

[**ClusterInstanceTypeResponseList**](ClusterInstanceTypeResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAWSManagedDatabaseInstanceType

> ManagedDatabaseInstanceTypeResponseList ListAWSManagedDatabaseInstanceType(ctx, region, databaseType).Execute()

List AWS available managed database instance types

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
    region := "us-east-2" // string | region name
    databaseType := "MYSQL" // string | Database type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProviderAPI.ListAWSManagedDatabaseInstanceType(context.Background(), region, databaseType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListAWSManagedDatabaseInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSManagedDatabaseInstanceType`: ManagedDatabaseInstanceTypeResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListAWSManagedDatabaseInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region name | 
**databaseType** | **string** | Database type | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAWSManagedDatabaseInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ManagedDatabaseInstanceTypeResponseList**](ManagedDatabaseInstanceTypeResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAWSManagedDatabaseType

> ManagedDatabaseTypeResponseList ListAWSManagedDatabaseType(ctx).Execute()

List AWS available managed database types

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
    resp, r, err := apiClient.CloudProviderAPI.ListAWSManagedDatabaseType(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListAWSManagedDatabaseType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSManagedDatabaseType`: ManagedDatabaseTypeResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListAWSManagedDatabaseType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAWSManagedDatabaseTypeRequest struct via the builder pattern


### Return type

[**ManagedDatabaseTypeResponseList**](ManagedDatabaseTypeResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAWSRegions

> ClusterRegionResponseList ListAWSRegions(ctx).Execute()

List AWS regions

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
    resp, r, err := apiClient.CloudProviderAPI.ListAWSRegions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListAWSRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSRegions`: ClusterRegionResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListAWSRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAWSRegionsRequest struct via the builder pattern


### Return type

[**ClusterRegionResponseList**](ClusterRegionResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudProvider

> CloudProviderResponseList ListCloudProvider(ctx).Execute()

List Cloud providers available

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
    resp, r, err := apiClient.CloudProviderAPI.ListCloudProvider(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListCloudProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCloudProvider`: CloudProviderResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListCloudProvider`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudProviderRequest struct via the builder pattern


### Return type

[**CloudProviderResponseList**](CloudProviderResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDOFeatures

> ClusterFeatureResponseList ListDOFeatures(ctx).Execute()

List DO features available

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
    resp, r, err := apiClient.CloudProviderAPI.ListDOFeatures(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListDOFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDOFeatures`: ClusterFeatureResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListDOFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDOFeaturesRequest struct via the builder pattern


### Return type

[**ClusterFeatureResponseList**](ClusterFeatureResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDOInstanceType

> ClusterInstanceTypeResponseList ListDOInstanceType(ctx).Execute()

List DO available instance types

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
    resp, r, err := apiClient.CloudProviderAPI.ListDOInstanceType(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListDOInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDOInstanceType`: ClusterInstanceTypeResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListDOInstanceType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDOInstanceTypeRequest struct via the builder pattern


### Return type

[**ClusterInstanceTypeResponseList**](ClusterInstanceTypeResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDOManagedDatabaseInstanceType

> ManagedDatabaseInstanceTypeResponseList ListDOManagedDatabaseInstanceType(ctx, region, databaseType).Execute()

List Digital Ocean available managed database instance types

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
    region := "us-east-2" // string | region name
    databaseType := "MYSQL" // string | Database type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProviderAPI.ListDOManagedDatabaseInstanceType(context.Background(), region, databaseType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListDOManagedDatabaseInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDOManagedDatabaseInstanceType`: ManagedDatabaseInstanceTypeResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListDOManagedDatabaseInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region name | 
**databaseType** | **string** | Database type | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDOManagedDatabaseInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ManagedDatabaseInstanceTypeResponseList**](ManagedDatabaseInstanceTypeResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDOManagedDatabaseType

> ManagedDatabaseTypeResponseList ListDOManagedDatabaseType(ctx).Execute()

List Digital Ocean available managed database types

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
    resp, r, err := apiClient.CloudProviderAPI.ListDOManagedDatabaseType(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListDOManagedDatabaseType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDOManagedDatabaseType`: ManagedDatabaseTypeResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListDOManagedDatabaseType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDOManagedDatabaseTypeRequest struct via the builder pattern


### Return type

[**ManagedDatabaseTypeResponseList**](ManagedDatabaseTypeResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDORegions

> ClusterRegionResponseList ListDORegions(ctx).Execute()

List DO regions

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
    resp, r, err := apiClient.CloudProviderAPI.ListDORegions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListDORegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDORegions`: ClusterRegionResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListDORegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDORegionsRequest struct via the builder pattern


### Return type

[**ClusterRegionResponseList**](ClusterRegionResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSCWManagedDatabaseInstanceType

> ManagedDatabaseInstanceTypeResponseList ListSCWManagedDatabaseInstanceType(ctx, databaseType).Execute()

List Scaleway available managed database instance types

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
    databaseType := "MYSQL" // string | Database type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProviderAPI.ListSCWManagedDatabaseInstanceType(context.Background(), databaseType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListSCWManagedDatabaseInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSCWManagedDatabaseInstanceType`: ManagedDatabaseInstanceTypeResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListSCWManagedDatabaseInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**databaseType** | **string** | Database type | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSCWManagedDatabaseInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedDatabaseInstanceTypeResponseList**](ManagedDatabaseInstanceTypeResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSCWManagedDatabaseType

> ManagedDatabaseTypeResponseList ListSCWManagedDatabaseType(ctx).Execute()

List Scaleway available managed database types

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
    resp, r, err := apiClient.CloudProviderAPI.ListSCWManagedDatabaseType(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListSCWManagedDatabaseType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSCWManagedDatabaseType`: ManagedDatabaseTypeResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListSCWManagedDatabaseType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSCWManagedDatabaseTypeRequest struct via the builder pattern


### Return type

[**ManagedDatabaseTypeResponseList**](ManagedDatabaseTypeResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScalewayFeatures

> ClusterFeatureResponseList ListScalewayFeatures(ctx).Execute()

List Scaleway features available

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
    resp, r, err := apiClient.CloudProviderAPI.ListScalewayFeatures(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListScalewayFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScalewayFeatures`: ClusterFeatureResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListScalewayFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListScalewayFeaturesRequest struct via the builder pattern


### Return type

[**ClusterFeatureResponseList**](ClusterFeatureResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScalewayInstanceType

> ClusterInstanceTypeResponseList ListScalewayInstanceType(ctx).Execute()

List Scaleway available instance types

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
    resp, r, err := apiClient.CloudProviderAPI.ListScalewayInstanceType(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListScalewayInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScalewayInstanceType`: ClusterInstanceTypeResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListScalewayInstanceType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListScalewayInstanceTypeRequest struct via the builder pattern


### Return type

[**ClusterInstanceTypeResponseList**](ClusterInstanceTypeResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScalewayKapsuleInstanceType

> ClusterInstanceTypeResponseList ListScalewayKapsuleInstanceType(ctx, zone).Execute()

List Scaleway Kapsule available instance types

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
    zone := "fr-par-1" // string | zone name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudProviderAPI.ListScalewayKapsuleInstanceType(context.Background(), zone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListScalewayKapsuleInstanceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScalewayKapsuleInstanceType`: ClusterInstanceTypeResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListScalewayKapsuleInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | zone name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListScalewayKapsuleInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterInstanceTypeResponseList**](ClusterInstanceTypeResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScalewayRegions

> ClusterRegionResponseList ListScalewayRegions(ctx).Execute()

List Scaleway regions

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
    resp, r, err := apiClient.CloudProviderAPI.ListScalewayRegions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderAPI.ListScalewayRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScalewayRegions`: ClusterRegionResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderAPI.ListScalewayRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListScalewayRegionsRequest struct via the builder pattern


### Return type

[**ClusterRegionResponseList**](ClusterRegionResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

