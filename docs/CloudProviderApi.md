# \CloudProviderApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAWSFeatures**](CloudProviderApi.md#ListAWSFeatures) | **Get** /aws/clusterFeature | List AWS features available
[**ListAWSRegions**](CloudProviderApi.md#ListAWSRegions) | **Get** /aws/region | List AWS regions
[**ListCloudProvider**](CloudProviderApi.md#ListCloudProvider) | **Get** /cloudProvider | List Cloud providers available
[**ListDOFeatures**](CloudProviderApi.md#ListDOFeatures) | **Get** /digitalOcean/clusterFeature | List DO features available
[**ListDORegions**](CloudProviderApi.md#ListDORegions) | **Get** /digitalOcean/region | List DO regions
[**ListScalewayFeatures**](CloudProviderApi.md#ListScalewayFeatures) | **Get** /scaleway/clusterFeature | List Scaleway features available
[**ListScalewayRegions**](CloudProviderApi.md#ListScalewayRegions) | **Get** /scaleway/region | List Scaleway regions



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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderApi.ListAWSFeatures(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderApi.ListAWSFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSFeatures`: ClusterFeatureResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderApi.ListAWSFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAWSFeaturesRequest struct via the builder pattern


### Return type

[**ClusterFeatureResponseList**](ClusterFeatureResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderApi.ListAWSRegions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderApi.ListAWSRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSRegions`: ClusterRegionResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderApi.ListAWSRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAWSRegionsRequest struct via the builder pattern


### Return type

[**ClusterRegionResponseList**](ClusterRegionResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderApi.ListCloudProvider(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderApi.ListCloudProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCloudProvider`: CloudProviderResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderApi.ListCloudProvider`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudProviderRequest struct via the builder pattern


### Return type

[**CloudProviderResponseList**](CloudProviderResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderApi.ListDOFeatures(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderApi.ListDOFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDOFeatures`: ClusterFeatureResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderApi.ListDOFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDOFeaturesRequest struct via the builder pattern


### Return type

[**ClusterFeatureResponseList**](ClusterFeatureResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderApi.ListDORegions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderApi.ListDORegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDORegions`: ClusterRegionResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderApi.ListDORegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDORegionsRequest struct via the builder pattern


### Return type

[**ClusterRegionResponseList**](ClusterRegionResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderApi.ListScalewayFeatures(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderApi.ListScalewayFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScalewayFeatures`: ClusterFeatureResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderApi.ListScalewayFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListScalewayFeaturesRequest struct via the builder pattern


### Return type

[**ClusterFeatureResponseList**](ClusterFeatureResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProviderApi.ListScalewayRegions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderApi.ListScalewayRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScalewayRegions`: ClusterRegionResponseList
    fmt.Fprintf(os.Stdout, "Response from `CloudProviderApi.ListScalewayRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListScalewayRegionsRequest struct via the builder pattern


### Return type

[**ClusterRegionResponseList**](ClusterRegionResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

