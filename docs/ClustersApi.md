# \ClustersApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](ClustersApi.md#CreateCluster) | **Post** /organization/{organizationId}/cluster | Create a cluster
[**DeleteCluster**](ClustersApi.md#DeleteCluster) | **Delete** /organization/{organizationId}/cluster/{clusterId} | Delete a cluster
[**DeployCluster**](ClustersApi.md#DeployCluster) | **Post** /organization/{organizationId}/cluster/{clusterId}/deploy | Deploy a cluster
[**EditCluster**](ClustersApi.md#EditCluster) | **Put** /organization/{organizationId}/cluster/{clusterId} | Edit a cluster
[**EditRoutingTable**](ClustersApi.md#EditRoutingTable) | **Put** /organization/{organizationId}/cluster/{clusterId}/routingTable | Edit routing table
[**GetClusterReadinessStatus**](ClustersApi.md#GetClusterReadinessStatus) | **Get** /organization/{organizationId}/cluster/{clusterId}/isReady | Know if a cluster is ready to be deployed or not
[**GetClusterStatus**](ClustersApi.md#GetClusterStatus) | **Get** /organization/{organizationId}/cluster/{clusterId}/status | Get cluster status
[**GetOrganizationCloudProviderInfo**](ClustersApi.md#GetOrganizationCloudProviderInfo) | **Get** /organization/{organizationId}/cluster/{clusterId}/cloudProviderInfo | Get cluster cloud provider info and credentials
[**GetOrganizationClusterStatus**](ClustersApi.md#GetOrganizationClusterStatus) | **Get** /organization/{organizationId}/cluster/status | List all clusters statuses
[**GetRoutingTable**](ClustersApi.md#GetRoutingTable) | **Get** /organization/{organizationId}/cluster/{clusterId}/routingTable | Get routing table
[**ListClusterLogs**](ClustersApi.md#ListClusterLogs) | **Get** /organization/{organizationId}/cluster/{clusterId}/logs | List Cluster Logs
[**ListOrganizationCluster**](ClustersApi.md#ListOrganizationCluster) | **Get** /organization/{organizationId}/cluster | List organization clusters
[**SpecifyClusterCloudProviderInfo**](ClustersApi.md#SpecifyClusterCloudProviderInfo) | **Post** /organization/{organizationId}/cluster/{clusterId}/cloudProviderInfo | Specify cluster cloud provider info and credentials
[**StopCluster**](ClustersApi.md#StopCluster) | **Post** /organization/{organizationId}/cluster/{clusterId}/stop | Stop cluster
[**UpdateCluster**](ClustersApi.md#UpdateCluster) | **Post** /organization/{organizationId}/cluster/{clusterId}/update | Update a cluster Version



## CreateCluster

> Cluster CreateCluster(ctx, organizationId).ClusterRequest(clusterRequest).Execute()

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    clusterRequest := *openapiclient.NewClusterRequest("Name_example", "Region_example", openapiclient.CloudProviderEnum("AWS")) // ClusterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.CreateCluster(context.Background(), organizationId).ClusterRequest(clusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.CreateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.CreateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterRequest** | [**ClusterRequest**](ClusterRequest.md) |  | 

### Return type

[**Cluster**](Cluster.md)

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.DeleteCluster(context.Background(), organizationId, clusterId).Execute()
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
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

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

> ClusterStatus DeployCluster(ctx, organizationId, clusterId).Execute()

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.DeployCluster(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeployCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployCluster`: ClusterStatus
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.DeployCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterStatus**](ClusterStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditCluster

> Cluster EditCluster(ctx, organizationId, clusterId).ClusterRequest(clusterRequest).Execute()

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
    clusterRequest := *openapiclient.NewClusterRequest("Name_example", "Region_example", openapiclient.CloudProviderEnum("AWS")) // ClusterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.EditCluster(context.Background(), organizationId, clusterId).ClusterRequest(clusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.EditCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.EditCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterRequest** | [**ClusterRequest**](ClusterRequest.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditRoutingTable

> ClusterRoutingTable EditRoutingTable(ctx, organizationId, clusterId).ClusterRoutingTableRequest(clusterRoutingTableRequest).Execute()

Edit routing table



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
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
    clusterRoutingTableRequest := *openapiclient.NewClusterRoutingTableRequest([]openapiclient.ClusterRoutingTableRequestRoutesInner{*openapiclient.NewClusterRoutingTableRequestRoutesInner("Destination_example", "Target_example", "Description_example")}) // ClusterRoutingTableRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.EditRoutingTable(context.Background(), organizationId, clusterId).ClusterRoutingTableRequest(clusterRoutingTableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.EditRoutingTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditRoutingTable`: ClusterRoutingTable
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.EditRoutingTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditRoutingTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterRoutingTableRequest** | [**ClusterRoutingTableRequest**](ClusterRoutingTableRequest.md) |  | 

### Return type

[**ClusterRoutingTable**](ClusterRoutingTable.md)

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.GetClusterReadinessStatus(context.Background(), organizationId, clusterId).Execute()
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
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

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

> ClusterStatusGet GetClusterStatus(ctx, organizationId, clusterId).Execute()

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.GetClusterStatus(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterStatus`: ClusterStatusGet
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterStatusGet**](ClusterStatusGet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCloudProviderInfo

> ClusterCloudProviderInfo GetOrganizationCloudProviderInfo(ctx, organizationId, clusterId).Execute()

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.GetOrganizationCloudProviderInfo(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetOrganizationCloudProviderInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCloudProviderInfo`: ClusterCloudProviderInfo
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetOrganizationCloudProviderInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCloudProviderInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterCloudProviderInfo**](ClusterCloudProviderInfo.md)

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.GetOrganizationClusterStatus(context.Background(), organizationId).Execute()
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
**organizationId** | **string** | Organization ID | 

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


## GetRoutingTable

> ClusterRoutingTable GetRoutingTable(ctx, organizationId, clusterId).Execute()

Get routing table



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
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.GetRoutingTable(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetRoutingTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoutingTable`: ClusterRoutingTable
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetRoutingTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutingTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterRoutingTable**](ClusterRoutingTable.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterLogs

> ListClusterLogs200Response ListClusterLogs(ctx, organizationId, clusterId).Execute()

List Cluster Logs



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
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ListClusterLogs(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusterLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterLogs`: ListClusterLogs200Response
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusterLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListClusterLogs200Response**](ListClusterLogs200Response.md)

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ListOrganizationCluster(context.Background(), organizationId).Execute()
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
**organizationId** | **string** | Organization ID | 

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

> ClusterCloudProviderInfo SpecifyClusterCloudProviderInfo(ctx, organizationId, clusterId).ClusterCloudProviderInfoRequest(clusterCloudProviderInfoRequest).Execute()

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
    clusterCloudProviderInfoRequest := *openapiclient.NewClusterCloudProviderInfoRequest() // ClusterCloudProviderInfoRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.SpecifyClusterCloudProviderInfo(context.Background(), organizationId, clusterId).ClusterCloudProviderInfoRequest(clusterCloudProviderInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.SpecifyClusterCloudProviderInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpecifyClusterCloudProviderInfo`: ClusterCloudProviderInfo
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.SpecifyClusterCloudProviderInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpecifyClusterCloudProviderInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterCloudProviderInfoRequest** | [**ClusterCloudProviderInfoRequest**](ClusterCloudProviderInfoRequest.md) |  | 

### Return type

[**ClusterCloudProviderInfo**](ClusterCloudProviderInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopCluster

> ClusterStatus StopCluster(ctx, organizationId, clusterId).Execute()

Stop cluster



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
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.StopCluster(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.StopCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopCluster`: ClusterStatus
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.StopCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterStatus**](ClusterStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> ClusterStatus UpdateCluster(ctx, organizationId, clusterId).Execute()

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.UpdateCluster(context.Background(), organizationId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCluster`: ClusterStatus
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterStatus**](ClusterStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

