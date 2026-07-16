# \PlatformConfigurationAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterPlatformBinding**](PlatformConfigurationAPI.md#GetClusterPlatformBinding) | **Get** /organization/{organizationId}/cluster/{clusterId}/platformBinding | Get the cluster platform binding
[**ListPlatformTemplates**](PlatformConfigurationAPI.md#ListPlatformTemplates) | **Get** /organization/{organizationId}/platformTemplate | List platform templates
[**ResolvePlatformComponentConfiguration**](PlatformConfigurationAPI.md#ResolvePlatformComponentConfiguration) | **Post** /organization/{organizationId}/cluster/{clusterId}/platformBinding/component/{componentKey}/resolve | Resolve a platform component configuration
[**ResolvePlatformTemplateComponentConfiguration**](PlatformConfigurationAPI.md#ResolvePlatformTemplateComponentConfiguration) | **Post** /organization/{organizationId}/platformTemplate/{templateKey}/{templateVersion}/component/{componentKey}/resolve | Resolve a platform component configuration before cluster creation
[**UpdateClusterPlatformBinding**](PlatformConfigurationAPI.md#UpdateClusterPlatformBinding) | **Put** /organization/{organizationId}/cluster/{clusterId}/platformBinding | Update the cluster platform binding



## GetClusterPlatformBinding

> ClusterPlatformBindingResponse GetClusterPlatformBinding(ctx, organizationId, clusterId).Execute()

Get the cluster platform binding



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
	resp, r, err := apiClient.PlatformConfigurationAPI.GetClusterPlatformBinding(context.Background(), organizationId, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformConfigurationAPI.GetClusterPlatformBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterPlatformBinding`: ClusterPlatformBindingResponse
	fmt.Fprintf(os.Stdout, "Response from `PlatformConfigurationAPI.GetClusterPlatformBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterPlatformBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterPlatformBindingResponse**](ClusterPlatformBindingResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlatformTemplates

> PlatformTemplateCatalogResponse ListPlatformTemplates(ctx, organizationId).ClusterMode(clusterMode).CloudProvider(cloudProvider).Execute()

List platform templates



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
	clusterMode := openapiclient.PlatformClusterMode("QOVERY_MANAGED") // PlatformClusterMode | Cluster management mode. Must be supplied together with cloudProvider. (optional)
	cloudProvider := openapiclient.PlatformCloudVendor("AWS") // PlatformCloudVendor | Cluster cloud provider. Must be supplied together with clusterMode. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformConfigurationAPI.ListPlatformTemplates(context.Background(), organizationId).ClusterMode(clusterMode).CloudProvider(cloudProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformConfigurationAPI.ListPlatformTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlatformTemplates`: PlatformTemplateCatalogResponse
	fmt.Fprintf(os.Stdout, "Response from `PlatformConfigurationAPI.ListPlatformTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPlatformTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterMode** | [**PlatformClusterMode**](PlatformClusterMode.md) | Cluster management mode. Must be supplied together with cloudProvider. | 
 **cloudProvider** | [**PlatformCloudVendor**](PlatformCloudVendor.md) | Cluster cloud provider. Must be supplied together with clusterMode. | 

### Return type

[**PlatformTemplateCatalogResponse**](PlatformTemplateCatalogResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolvePlatformComponentConfiguration

> PlatformComponentConfigurationPreviewResponse ResolvePlatformComponentConfiguration(ctx, organizationId, clusterId, componentKey).PlatformComponentConfigurationPreviewRequest(platformComponentConfigurationPreviewRequest).Execute()

Resolve a platform component configuration



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
	componentKey := "componentKey_example" // string | Platform component key
	platformComponentConfigurationPreviewRequest := *openapiclient.NewPlatformComponentConfigurationPreviewRequest() // PlatformComponentConfigurationPreviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformConfigurationAPI.ResolvePlatformComponentConfiguration(context.Background(), organizationId, clusterId, componentKey).PlatformComponentConfigurationPreviewRequest(platformComponentConfigurationPreviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformConfigurationAPI.ResolvePlatformComponentConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolvePlatformComponentConfiguration`: PlatformComponentConfigurationPreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `PlatformConfigurationAPI.ResolvePlatformComponentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 
**componentKey** | **string** | Platform component key | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolvePlatformComponentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **platformComponentConfigurationPreviewRequest** | [**PlatformComponentConfigurationPreviewRequest**](PlatformComponentConfigurationPreviewRequest.md) |  | 

### Return type

[**PlatformComponentConfigurationPreviewResponse**](PlatformComponentConfigurationPreviewResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolvePlatformTemplateComponentConfiguration

> PlatformComponentConfigurationResolutionResponse ResolvePlatformTemplateComponentConfiguration(ctx, organizationId, templateKey, templateVersion, componentKey).ClusterMode(clusterMode).CloudProvider(cloudProvider).PlatformComponentConfigurationPreviewRequest(platformComponentConfigurationPreviewRequest).Execute()

Resolve a platform component configuration before cluster creation



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
	templateKey := "templateKey_example" // string | Platform template key
	templateVersion := "templateVersion_example" // string | Platform template version
	componentKey := "componentKey_example" // string | Platform component key
	clusterMode := openapiclient.PlatformClusterMode("QOVERY_MANAGED") // PlatformClusterMode | Cluster management mode used to resolve component applicability
	cloudProvider := openapiclient.PlatformCloudVendor("AWS") // PlatformCloudVendor | Cluster cloud provider used to resolve component applicability
	platformComponentConfigurationPreviewRequest := *openapiclient.NewPlatformComponentConfigurationPreviewRequest() // PlatformComponentConfigurationPreviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformConfigurationAPI.ResolvePlatformTemplateComponentConfiguration(context.Background(), organizationId, templateKey, templateVersion, componentKey).ClusterMode(clusterMode).CloudProvider(cloudProvider).PlatformComponentConfigurationPreviewRequest(platformComponentConfigurationPreviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformConfigurationAPI.ResolvePlatformTemplateComponentConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolvePlatformTemplateComponentConfiguration`: PlatformComponentConfigurationResolutionResponse
	fmt.Fprintf(os.Stdout, "Response from `PlatformConfigurationAPI.ResolvePlatformTemplateComponentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**templateKey** | **string** | Platform template key | 
**templateVersion** | **string** | Platform template version | 
**componentKey** | **string** | Platform component key | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolvePlatformTemplateComponentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **clusterMode** | [**PlatformClusterMode**](PlatformClusterMode.md) | Cluster management mode used to resolve component applicability | 
 **cloudProvider** | [**PlatformCloudVendor**](PlatformCloudVendor.md) | Cluster cloud provider used to resolve component applicability | 
 **platformComponentConfigurationPreviewRequest** | [**PlatformComponentConfigurationPreviewRequest**](PlatformComponentConfigurationPreviewRequest.md) |  | 

### Return type

[**PlatformComponentConfigurationResolutionResponse**](PlatformComponentConfigurationResolutionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClusterPlatformBinding

> ClusterPlatformBindingResponse UpdateClusterPlatformBinding(ctx, organizationId, clusterId).ClusterPlatformBindingRequest(clusterPlatformBindingRequest).Execute()

Update the cluster platform binding



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
	clusterPlatformBindingRequest := *openapiclient.NewClusterPlatformBindingRequest("TemplateKey_example", "TemplateVersion_example") // ClusterPlatformBindingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformConfigurationAPI.UpdateClusterPlatformBinding(context.Background(), organizationId, clusterId).ClusterPlatformBindingRequest(clusterPlatformBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformConfigurationAPI.UpdateClusterPlatformBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClusterPlatformBinding`: ClusterPlatformBindingResponse
	fmt.Fprintf(os.Stdout, "Response from `PlatformConfigurationAPI.UpdateClusterPlatformBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterPlatformBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterPlatformBindingRequest** | [**ClusterPlatformBindingRequest**](ClusterPlatformBindingRequest.md) |  | 

### Return type

[**ClusterPlatformBindingResponse**](ClusterPlatformBindingResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

