# \ClustersApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](ClustersApi.md#CreateCluster) | **Post** /organization/{organizationId}/cluster | Create a cluster
[**DeleteCluster**](ClustersApi.md#DeleteCluster) | **Delete** /organization/{organizationId}/cluster/{clusterId} | Delete a cluster
[**DeployCluster**](ClustersApi.md#DeployCluster) | **Post** /organization/{organizationId}/cluster/{clusterId}/deploy | Deploy a cluster
[**EditCluster**](ClustersApi.md#EditCluster) | **Put** /organization/{organizationId}/cluster/{clusterId} | Edit a cluster
[**GetClusterReadinessStatus**](ClustersApi.md#GetClusterReadinessStatus) | **Get** /organization/{organizationId}/cluster/{clusterId}/isReady | Know if a cluster is ready to be deployed or not
[**GetClusterStatus**](ClustersApi.md#GetClusterStatus) | **Get** /organization/{organizationId}/cluster/{clusterId}/status | Get cluster status
[**GetOrganizationCloudProviderInfo**](ClustersApi.md#GetOrganizationCloudProviderInfo) | **Get** /organization/{organizationId}/cluster/{clusterId}/cloudProviderInfo | Get cluster cloud provider info and credentials
[**GetOrganizationClusterStatus**](ClustersApi.md#GetOrganizationClusterStatus) | **Get** /organization/{organizationId}/cluster/status | List all clusters statuses
[**ListOrganizationCluster**](ClustersApi.md#ListOrganizationCluster) | **Get** /organization/{organizationId}/cluster | List organization clusters
[**SpecifyClusterCloudProviderInfo**](ClustersApi.md#SpecifyClusterCloudProviderInfo) | **Post** /organization/{organizationId}/cluster/{clusterId}/cloudProviderInfo | Specify cluster cloud provider info and credentials
[**UpdateCluster**](ClustersApi.md#UpdateCluster) | **Post** /organization/{organizationId}/cluster/{clusterId}/update | Update a cluster Version



## CreateCluster

> ClusterResponse CreateCluster(ctx, organizationId).ClusterRequest(clusterRequest).Execute()

Create a cluster

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
    organizationId := TODO // string | Organization ID
    clusterRequest := *openapiclient.NewClusterRequest("Name_example") // ClusterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.CreateCluster(context.Background(), organizationId).ClusterRequest(clusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.CreateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCluster`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.CreateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterRequest** | [**ClusterRequest**](ClusterRequest.md) |  | 

### Return type

[**ClusterResponse**](ClusterResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCluster

> DeleteCluster(ctx, organizationId, clusterId).Execute()

Delete a cluster

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
    organizationId := TODO // string | Organization ID
    clusterId := TODO // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteCluster(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 
**clusterId** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterRequest struct via the builder pattern


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


## DeployCluster

> ClusterStatusResponse DeployCluster(ctx, organizationId, clusterId).Execute()

Deploy a cluster



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
    organizationId := TODO // string | Organization ID
    clusterId := TODO // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.DeployCluster(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeployCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployCluster`: ClusterStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.DeployCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 
**clusterId** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterStatusResponse**](ClusterStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditCluster

> ClusterResponse EditCluster(ctx, organizationId, clusterId).ClusterRequest(clusterRequest).Execute()

Edit a cluster

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
    organizationId := TODO // string | Organization ID
    clusterId := TODO // string | Cluster ID
    clusterRequest := *openapiclient.NewClusterRequest("Name_example") // ClusterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.EditCluster(context.Background(), organizationId, clusterId).ClusterRequest(clusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.EditCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditCluster`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.EditCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 
**clusterId** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterRequest** | [**ClusterRequest**](ClusterRequest.md) |  | 

### Return type

[**ClusterResponse**](ClusterResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterReadinessStatus

> ClusterReadinessStatus GetClusterReadinessStatus(ctx, organizationId, clusterId).Execute()

Know if a cluster is ready to be deployed or not

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
    organizationId := TODO // string | Organization ID
    clusterId := TODO // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetClusterReadinessStatus(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterReadinessStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterReadinessStatus`: ClusterReadinessStatus
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterReadinessStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 
**clusterId** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterReadinessStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterReadinessStatus**](ClusterReadinessStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterStatus

> ClusterStatusResponse GetClusterStatus(ctx, organizationId, clusterId).Execute()

Get cluster status

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
    organizationId := TODO // string | Organization ID
    clusterId := TODO // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetClusterStatus(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterStatus`: ClusterStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 
**clusterId** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterStatusResponse**](ClusterStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCloudProviderInfo

> ClusterCloudProviderInfoResponse GetOrganizationCloudProviderInfo(ctx, organizationId, clusterId).Execute()

Get cluster cloud provider info and credentials

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
    organizationId := TODO // string | Organization ID
    clusterId := TODO // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetOrganizationCloudProviderInfo(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetOrganizationCloudProviderInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCloudProviderInfo`: ClusterCloudProviderInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetOrganizationCloudProviderInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 
**clusterId** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCloudProviderInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterCloudProviderInfoResponse**](ClusterCloudProviderInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationClusterStatus

> ClusterStatusResponseList GetOrganizationClusterStatus(ctx, organizationId).Execute()

List all clusters statuses



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
    organizationId := TODO // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.GetOrganizationClusterStatus(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetOrganizationClusterStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationClusterStatus`: ClusterStatusResponseList
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetOrganizationClusterStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationClusterStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterStatusResponseList**](ClusterStatusResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationCluster

> ClusterResponseList ListOrganizationCluster(ctx, organizationId).Execute()

List organization clusters

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
    organizationId := TODO // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.ListOrganizationCluster(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListOrganizationCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizationCluster`: ClusterResponseList
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListOrganizationCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterResponseList**](ClusterResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpecifyClusterCloudProviderInfo

> ClusterCloudProviderInfoResponse SpecifyClusterCloudProviderInfo(ctx, organizationId, clusterId).ClusterCloudProviderInfoRequest(clusterCloudProviderInfoRequest).Execute()

Specify cluster cloud provider info and credentials

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
    organizationId := TODO // string | Organization ID
    clusterId := TODO // string | Cluster ID
    clusterCloudProviderInfoRequest := *openapiclient.NewClusterCloudProviderInfoRequest() // ClusterCloudProviderInfoRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.SpecifyClusterCloudProviderInfo(context.Background(), organizationId, clusterId).ClusterCloudProviderInfoRequest(clusterCloudProviderInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.SpecifyClusterCloudProviderInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpecifyClusterCloudProviderInfo`: ClusterCloudProviderInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.SpecifyClusterCloudProviderInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 
**clusterId** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpecifyClusterCloudProviderInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterCloudProviderInfoRequest** | [**ClusterCloudProviderInfoRequest**](ClusterCloudProviderInfoRequest.md) |  | 

### Return type

[**ClusterCloudProviderInfoResponse**](ClusterCloudProviderInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> ClusterStatusResponse UpdateCluster(ctx, organizationId, clusterId).Execute()

Update a cluster Version



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
    organizationId := TODO // string | Organization ID
    clusterId := TODO // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClustersApi.UpdateCluster(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCluster`: ClusterStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) | Organization ID | 
**clusterId** | [**string**](.md) | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterStatusResponse**](ClusterStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

