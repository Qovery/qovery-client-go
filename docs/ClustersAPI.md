# \ClustersAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](ClustersAPI.md#CreateCluster) | **Post** /organization/{organizationId}/cluster | Create a cluster
[**DeleteCluster**](ClustersAPI.md#DeleteCluster) | **Delete** /organization/{organizationId}/cluster/{clusterId} | Delete a cluster
[**DeployCluster**](ClustersAPI.md#DeployCluster) | **Post** /organization/{organizationId}/cluster/{clusterId}/deploy | Deploy a cluster
[**EditCluster**](ClustersAPI.md#EditCluster) | **Put** /organization/{organizationId}/cluster/{clusterId} | Edit a cluster
[**EditClusterAdvancedSettings**](ClustersAPI.md#EditClusterAdvancedSettings) | **Put** /organization/{organizationId}/cluster/{clusterId}/advancedSettings | Edit advanced settings
[**EditClusterKubeconfig**](ClustersAPI.md#EditClusterKubeconfig) | **Put** /organization/{organizationId}/cluster/{clusterId}/kubeconfig | Edit cluster kubeconfig
[**EditRoutingTable**](ClustersAPI.md#EditRoutingTable) | **Put** /organization/{organizationId}/cluster/{clusterId}/routingTable | Edit routing table
[**GetClusterAdvancedSettings**](ClustersAPI.md#GetClusterAdvancedSettings) | **Get** /organization/{organizationId}/cluster/{clusterId}/advancedSettings | Get advanced settings
[**GetClusterKubeconfig**](ClustersAPI.md#GetClusterKubeconfig) | **Get** /organization/{organizationId}/cluster/{clusterId}/kubeconfig | Get cluster kubeconfig
[**GetClusterKubernetesEvents**](ClustersAPI.md#GetClusterKubernetesEvents) | **Get** /cluster/{clusterId}/events | List Cluster Kubernetes Events
[**GetClusterMetrics**](ClustersAPI.md#GetClusterMetrics) | **Get** /cluster/{clusterId}/metrics | Fetch cluster metrics
[**GetClusterReadinessStatus**](ClustersAPI.md#GetClusterReadinessStatus) | **Get** /organization/{organizationId}/cluster/{clusterId}/isReady | Know if a cluster is ready to be deployed or not
[**GetClusterStatus**](ClustersAPI.md#GetClusterStatus) | **Get** /organization/{organizationId}/cluster/{clusterId}/status | Get cluster status
[**GetDefaultClusterAdvancedSettings**](ClustersAPI.md#GetDefaultClusterAdvancedSettings) | **Get** /defaultClusterAdvancedSettings | List default cluster advanced settings
[**GetInstallationHelmValues**](ClustersAPI.md#GetInstallationHelmValues) | **Get** /organization/{organizationId}/cluster/{clusterId}/installationHelmValues | Get cluster helm values for self managed installation
[**GetOrganizationCloudProviderInfo**](ClustersAPI.md#GetOrganizationCloudProviderInfo) | **Get** /organization/{organizationId}/cluster/{clusterId}/cloudProviderInfo | Get cluster cloud provider info and credentials
[**GetOrganizationClusterStatus**](ClustersAPI.md#GetOrganizationClusterStatus) | **Get** /organization/{organizationId}/cluster/status | List all clusters statuses
[**GetRoutingTable**](ClustersAPI.md#GetRoutingTable) | **Get** /organization/{organizationId}/cluster/{clusterId}/routingTable | Get routing table
[**ListClusterLogs**](ClustersAPI.md#ListClusterLogs) | **Get** /organization/{organizationId}/cluster/{clusterId}/logs | List Cluster Logs
[**ListOrganizationCluster**](ClustersAPI.md#ListOrganizationCluster) | **Get** /organization/{organizationId}/cluster | List organization clusters
[**LockCluster**](ClustersAPI.md#LockCluster) | **Post** /cluster/{clusterId}/lock | Lock Cluster
[**SpecifyClusterCloudProviderInfo**](ClustersAPI.md#SpecifyClusterCloudProviderInfo) | **Post** /organization/{organizationId}/cluster/{clusterId}/cloudProviderInfo | Specify cluster cloud provider info and credentials
[**StopCluster**](ClustersAPI.md#StopCluster) | **Post** /organization/{organizationId}/cluster/{clusterId}/stop | Stop cluster
[**UnlockCluster**](ClustersAPI.md#UnlockCluster) | **Delete** /cluster/{clusterId}/lock | Unlock Cluster
[**UpdateKarpenterPrivateFargateSubnetIds**](ClustersAPI.md#UpdateKarpenterPrivateFargateSubnetIds) | **Put** /organization/{organizationId}/cluster/{clusterId}/karpenterPrivateSubnetIds | Update karpenter private fargate subnet ids
[**UpgradeCluster**](ClustersAPI.md#UpgradeCluster) | **Post** /cluster/{clusterId}/upgrade | Upgrade a cluster



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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	clusterRequest := *openapiclient.NewClusterRequest("Name_example", "Region_example", openapiclient.CloudVendorEnum("AWS")) // ClusterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.CreateCluster(context.Background(), organizationId).ClusterRequest(clusterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.CreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCluster`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.CreateCluster`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCluster

> DeleteCluster(ctx, organizationId, clusterId).DeleteMode(deleteMode).Execute()

Delete a cluster

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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	deleteMode := openapiclient.ClusterDeleteMode("DEFAULT") // ClusterDeleteMode |  (optional) (default to "DEFAULT")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClustersAPI.DeleteCluster(context.Background(), organizationId, clusterId).DeleteMode(deleteMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.DeleteCluster``: %v\n", err)
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


 **deleteMode** | [**ClusterDeleteMode**](ClusterDeleteMode.md) |  | [default to &quot;DEFAULT&quot;]

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


## DeployCluster

> ClusterStatus DeployCluster(ctx, organizationId, clusterId).DryRun(dryRun).Execute()

Deploy a cluster



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	dryRun := true // bool | default: false. Decide if the deployment of the cluster should be done in dry_run mode. Avoiding any changes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.DeployCluster(context.Background(), organizationId, clusterId).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.DeployCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeployCluster`: ClusterStatus
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.DeployCluster`: %v\n", resp)
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


 **dryRun** | **bool** | default: false. Decide if the deployment of the cluster should be done in dry_run mode. Avoiding any changes | 

### Return type

[**ClusterStatus**](ClusterStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	clusterRequest := *openapiclient.NewClusterRequest("Name_example", "Region_example", openapiclient.CloudVendorEnum("AWS")) // ClusterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.EditCluster(context.Background(), organizationId, clusterId).ClusterRequest(clusterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.EditCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditCluster`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.EditCluster`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditClusterAdvancedSettings

> ClusterAdvancedSettings EditClusterAdvancedSettings(ctx, organizationId, clusterId).ClusterAdvancedSettings(clusterAdvancedSettings).Execute()

Edit advanced settings



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	clusterAdvancedSettings := *openapiclient.NewClusterAdvancedSettings() // ClusterAdvancedSettings |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.EditClusterAdvancedSettings(context.Background(), organizationId, clusterId).ClusterAdvancedSettings(clusterAdvancedSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.EditClusterAdvancedSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditClusterAdvancedSettings`: ClusterAdvancedSettings
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.EditClusterAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditClusterAdvancedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterAdvancedSettings** | [**ClusterAdvancedSettings**](ClusterAdvancedSettings.md) |  | 

### Return type

[**ClusterAdvancedSettings**](ClusterAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditClusterKubeconfig

> EditClusterKubeconfig(ctx, organizationId, clusterId).Body(body).Execute()

Edit cluster kubeconfig

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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClustersAPI.EditClusterKubeconfig(context.Background(), organizationId, clusterId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.EditClusterKubeconfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEditClusterKubeconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-yaml
- **Accept**: Not defined

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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	clusterRoutingTableRequest := *openapiclient.NewClusterRoutingTableRequest([]openapiclient.ClusterRoutingTableResultsInner{*openapiclient.NewClusterRoutingTableResultsInner("Destination_example", "Target_example", "Description_example")}) // ClusterRoutingTableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.EditRoutingTable(context.Background(), organizationId, clusterId).ClusterRoutingTableRequest(clusterRoutingTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.EditRoutingTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditRoutingTable`: ClusterRoutingTable
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.EditRoutingTable`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterAdvancedSettings

> ClusterAdvancedSettings GetClusterAdvancedSettings(ctx, organizationId, clusterId).Execute()

Get advanced settings



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClusterAdvancedSettings(context.Background(), organizationId, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClusterAdvancedSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterAdvancedSettings`: ClusterAdvancedSettings
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClusterAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterAdvancedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterAdvancedSettings**](ClusterAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterKubeconfig

> string GetClusterKubeconfig(ctx, organizationId, clusterId).WithTokenFromCli(withTokenFromCli).Execute()

Get cluster kubeconfig

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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	withTokenFromCli := true // bool | If true, the user auth part will have an exec command with qovery cli (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClusterKubeconfig(context.Background(), organizationId, clusterId).WithTokenFromCli(withTokenFromCli).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClusterKubeconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterKubeconfig`: string
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClusterKubeconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterKubeconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withTokenFromCli** | **bool** | If true, the user auth part will have an exec command with qovery cli | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterKubernetesEvents

> GetClusterKubernetesEvents200Response GetClusterKubernetesEvents(ctx, clusterId).FromDateTime(fromDateTime).ToDateTime(toDateTime).NodeName(nodeName).PodName(podName).ReportingComponent(reportingComponent).Execute()

List Cluster Kubernetes Events



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	fromDateTime := "2025-03-03T09:00:00+00:00" // string | The start date time to fetch events from, following ISO-8601 format.   The `+` character must be escaped (`%2B`) 
	toDateTime := "2025-03-03T03:00:00+00:00" // string | The end date time to fetch events from, following ISO-8601 format.   The `+` character must be escaped (`%2B`) 
	nodeName := "nodeName_example" // string | The name of the node to fetch event from (optional)
	podName := "podName_example" // string | The name of the pod to fetch event from (optional)
	reportingComponent := "reportingComponent_example" // string | The name of the reporting component used to filter events. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClusterKubernetesEvents(context.Background(), clusterId).FromDateTime(fromDateTime).ToDateTime(toDateTime).NodeName(nodeName).PodName(podName).ReportingComponent(reportingComponent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClusterKubernetesEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterKubernetesEvents`: GetClusterKubernetesEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClusterKubernetesEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterKubernetesEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDateTime** | **string** | The start date time to fetch events from, following ISO-8601 format.   The &#x60;+&#x60; character must be escaped (&#x60;%2B&#x60;)  | 
 **toDateTime** | **string** | The end date time to fetch events from, following ISO-8601 format.   The &#x60;+&#x60; character must be escaped (&#x60;%2B&#x60;)  | 
 **nodeName** | **string** | The name of the node to fetch event from | 
 **podName** | **string** | The name of the pod to fetch event from | 
 **reportingComponent** | **string** | The name of the reporting component used to filter events. | 

### Return type

[**GetClusterKubernetesEvents200Response**](GetClusterKubernetesEvents200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterMetrics

> ClusterMetricsResponse GetClusterMetrics(ctx, clusterId).Endpoint(endpoint).Query(query).Start(start).End(end).Step(step).Time(time).Timeout(timeout).Dedup(dedup).PartialResponse(partialResponse).MaxSourceResolution(maxSourceResolution).Execute()

Fetch cluster metrics



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	endpoint := "endpoint_example" // string | 
	query := "query_example" // string | 
	start := "start_example" // string |  (optional)
	end := "end_example" // string |  (optional)
	step := "step_example" // string |  (optional)
	time := "time_example" // string |  (optional)
	timeout := "timeout_example" // string |  (optional)
	dedup := "dedup_example" // string |  (optional)
	partialResponse := "partialResponse_example" // string |  (optional)
	maxSourceResolution := "maxSourceResolution_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClusterMetrics(context.Background(), clusterId).Endpoint(endpoint).Query(query).Start(start).End(end).Step(step).Time(time).Timeout(timeout).Dedup(dedup).PartialResponse(partialResponse).MaxSourceResolution(maxSourceResolution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClusterMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterMetrics`: ClusterMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClusterMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpoint** | **string** |  | 
 **query** | **string** |  | 
 **start** | **string** |  | 
 **end** | **string** |  | 
 **step** | **string** |  | 
 **time** | **string** |  | 
 **timeout** | **string** |  | 
 **dedup** | **string** |  | 
 **partialResponse** | **string** |  | 
 **maxSourceResolution** | **string** |  | 

### Return type

[**ClusterMetricsResponse**](ClusterMetricsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClusterReadinessStatus(context.Background(), organizationId, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClusterReadinessStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterReadinessStatus`: ClusterReadinessStatus
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClusterReadinessStatus`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterStatus

> ClusterStatus GetClusterStatus(ctx, organizationId, clusterId).Execute()

Get cluster status

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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClusterStatus(context.Background(), organizationId, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClusterStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterStatus`: ClusterStatus
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClusterStatus`: %v\n", resp)
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

[**ClusterStatus**](ClusterStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultClusterAdvancedSettings

> ClusterAdvancedSettings GetDefaultClusterAdvancedSettings(ctx).Execute()

List default cluster advanced settings



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
	resp, r, err := apiClient.ClustersAPI.GetDefaultClusterAdvancedSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetDefaultClusterAdvancedSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultClusterAdvancedSettings`: ClusterAdvancedSettings
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetDefaultClusterAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultClusterAdvancedSettingsRequest struct via the builder pattern


### Return type

[**ClusterAdvancedSettings**](ClusterAdvancedSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstallationHelmValues

> string GetInstallationHelmValues(ctx, organizationId, clusterId).Execute()

Get cluster helm values for self managed installation

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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetInstallationHelmValues(context.Background(), organizationId, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetInstallationHelmValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstallationHelmValues`: string
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetInstallationHelmValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstallationHelmValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-yaml

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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetOrganizationCloudProviderInfo(context.Background(), organizationId, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetOrganizationCloudProviderInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCloudProviderInfo`: ClusterCloudProviderInfo
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetOrganizationCloudProviderInfo`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetOrganizationClusterStatus(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetOrganizationClusterStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationClusterStatus`: ClusterStatusResponseList
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetOrganizationClusterStatus`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetRoutingTable(context.Background(), organizationId, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetRoutingTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoutingTable`: ClusterRoutingTable
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetRoutingTable`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterLogs

> ClusterLogsResponseList ListClusterLogs(ctx, organizationId, clusterId).Execute()

List Cluster Logs



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.ListClusterLogs(context.Background(), organizationId, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ListClusterLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterLogs`: ClusterLogsResponseList
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ListClusterLogs`: %v\n", resp)
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

[**ClusterLogsResponseList**](ClusterLogsResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.ListOrganizationCluster(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ListOrganizationCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationCluster`: ClusterResponseList
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ListOrganizationCluster`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockCluster

> ClusterLock LockCluster(ctx, clusterId).ClusterLockRequest(clusterLockRequest).Execute()

Lock Cluster



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
	clusterId := "clusterId_example" // string | 
	clusterLockRequest := *openapiclient.NewClusterLockRequest("Reason_example") // ClusterLockRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.LockCluster(context.Background(), clusterId).ClusterLockRequest(clusterLockRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.LockCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LockCluster`: ClusterLock
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.LockCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterLockRequest** | [**ClusterLockRequest**](ClusterLockRequest.md) |  | 

### Return type

[**ClusterLock**](ClusterLock.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	clusterCloudProviderInfoRequest := *openapiclient.NewClusterCloudProviderInfoRequest() // ClusterCloudProviderInfoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.SpecifyClusterCloudProviderInfo(context.Background(), organizationId, clusterId).ClusterCloudProviderInfoRequest(clusterCloudProviderInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.SpecifyClusterCloudProviderInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpecifyClusterCloudProviderInfo`: ClusterCloudProviderInfo
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.SpecifyClusterCloudProviderInfo`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

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
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.StopCluster(context.Background(), organizationId, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.StopCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopCluster`: ClusterStatus
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.StopCluster`: %v\n", resp)
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlockCluster

> UnlockCluster(ctx, clusterId).Execute()

Unlock Cluster



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
	clusterId := "clusterId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClustersAPI.UnlockCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.UnlockCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockClusterRequest struct via the builder pattern


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


## UpdateKarpenterPrivateFargateSubnetIds

> UpdateKarpenterPrivateFargateSubnetIds(ctx, organizationId, clusterId).ClusterKarpenterPrivateSubnetIdsPutRequest(clusterKarpenterPrivateSubnetIdsPutRequest).Execute()

Update karpenter private fargate subnet ids



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID
	clusterKarpenterPrivateSubnetIdsPutRequest := *openapiclient.NewClusterKarpenterPrivateSubnetIdsPutRequest() // ClusterKarpenterPrivateSubnetIdsPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClustersAPI.UpdateKarpenterPrivateFargateSubnetIds(context.Background(), organizationId, clusterId).ClusterKarpenterPrivateSubnetIdsPutRequest(clusterKarpenterPrivateSubnetIdsPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.UpdateKarpenterPrivateFargateSubnetIds``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateKarpenterPrivateFargateSubnetIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterKarpenterPrivateSubnetIdsPutRequest** | [**ClusterKarpenterPrivateSubnetIdsPutRequest**](ClusterKarpenterPrivateSubnetIdsPutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeCluster

> ClusterStatus UpgradeCluster(ctx, clusterId).Execute()

Upgrade a cluster



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
	clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.UpgradeCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.UpgradeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeCluster`: ClusterStatus
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.UpgradeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterStatus**](ClusterStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

