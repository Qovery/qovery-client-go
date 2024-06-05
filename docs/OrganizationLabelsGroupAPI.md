# \OrganizationLabelsGroupAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationLabelsGroup**](OrganizationLabelsGroupAPI.md#CreateOrganizationLabelsGroup) | **Post** /organization/{organizationId}/labelsGroups | Create an organization labels group
[**DeleteOrganizationLabelsGroup**](OrganizationLabelsGroupAPI.md#DeleteOrganizationLabelsGroup) | **Delete** /organization/{organizationId}/labelsGroups/{labelsGroupId} | Delete organization labels group
[**EditOrganizationLabelsGroup**](OrganizationLabelsGroupAPI.md#EditOrganizationLabelsGroup) | **Put** /organization/{organizationId}/labelsGroups/{labelsGroupId} | Edit organization labels group
[**GetOrganizationLabelsGroupAssociatedItems**](OrganizationLabelsGroupAPI.md#GetOrganizationLabelsGroupAssociatedItems) | **Get** /organization/{organizationId}/labelsGroups/{labelsGroupId}/associatedItems | Get organization labels group associated items
[**GetOrganizationLabelssGroup**](OrganizationLabelsGroupAPI.md#GetOrganizationLabelssGroup) | **Get** /organization/{organizationId}/labelsGroups/{labelsGroupId} | Get organization labels group
[**ListOrganizationLabelsGroup**](OrganizationLabelsGroupAPI.md#ListOrganizationLabelsGroup) | **Get** /organization/{organizationId}/labelsGroups | List organization labels group



## CreateOrganizationLabelsGroup

> OrganizationLabelsGroupResponse CreateOrganizationLabelsGroup(ctx, organizationId).OrganizationLabelsGroupCreateRequest(organizationLabelsGroupCreateRequest).Execute()

Create an organization labels group



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
	organizationLabelsGroupCreateRequest := *openapiclient.NewOrganizationLabelsGroupCreateRequest("Name_example", []openapiclient.Label{*openapiclient.NewLabel("Key_example", "Value_example", false)}) // OrganizationLabelsGroupCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationLabelsGroupAPI.CreateOrganizationLabelsGroup(context.Background(), organizationId).OrganizationLabelsGroupCreateRequest(organizationLabelsGroupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationLabelsGroupAPI.CreateOrganizationLabelsGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationLabelsGroup`: OrganizationLabelsGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationLabelsGroupAPI.CreateOrganizationLabelsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationLabelsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationLabelsGroupCreateRequest** | [**OrganizationLabelsGroupCreateRequest**](OrganizationLabelsGroupCreateRequest.md) |  | 

### Return type

[**OrganizationLabelsGroupResponse**](OrganizationLabelsGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationLabelsGroup

> DeleteOrganizationLabelsGroup(ctx, organizationId, labelsGroupId).Execute()

Delete organization labels group



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
	labelsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization labels group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationLabelsGroupAPI.DeleteOrganizationLabelsGroup(context.Background(), organizationId, labelsGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationLabelsGroupAPI.DeleteOrganizationLabelsGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**labelsGroupId** | **string** | Organization labels group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationLabelsGroupRequest struct via the builder pattern


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


## EditOrganizationLabelsGroup

> OrganizationLabelsGroupResponse EditOrganizationLabelsGroup(ctx, organizationId, labelsGroupId).OrganizationLabelsGroupCreateRequest(organizationLabelsGroupCreateRequest).Execute()

Edit organization labels group



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
	labelsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization labels group ID
	organizationLabelsGroupCreateRequest := *openapiclient.NewOrganizationLabelsGroupCreateRequest("Name_example", []openapiclient.Label{*openapiclient.NewLabel("Key_example", "Value_example", false)}) // OrganizationLabelsGroupCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationLabelsGroupAPI.EditOrganizationLabelsGroup(context.Background(), organizationId, labelsGroupId).OrganizationLabelsGroupCreateRequest(organizationLabelsGroupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationLabelsGroupAPI.EditOrganizationLabelsGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOrganizationLabelsGroup`: OrganizationLabelsGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationLabelsGroupAPI.EditOrganizationLabelsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**labelsGroupId** | **string** | Organization labels group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganizationLabelsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **organizationLabelsGroupCreateRequest** | [**OrganizationLabelsGroupCreateRequest**](OrganizationLabelsGroupCreateRequest.md) |  | 

### Return type

[**OrganizationLabelsGroupResponse**](OrganizationLabelsGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLabelsGroupAssociatedItems

> OrganizationLabelsGroupAssociatedItemsResponseList GetOrganizationLabelsGroupAssociatedItems(ctx, organizationId, labelsGroupId).Execute()

Get organization labels group associated items



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
	labelsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization labels group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationLabelsGroupAPI.GetOrganizationLabelsGroupAssociatedItems(context.Background(), organizationId, labelsGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationLabelsGroupAPI.GetOrganizationLabelsGroupAssociatedItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationLabelsGroupAssociatedItems`: OrganizationLabelsGroupAssociatedItemsResponseList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationLabelsGroupAPI.GetOrganizationLabelsGroupAssociatedItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**labelsGroupId** | **string** | Organization labels group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLabelsGroupAssociatedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationLabelsGroupAssociatedItemsResponseList**](OrganizationLabelsGroupAssociatedItemsResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLabelssGroup

> OrganizationLabelsGroupResponse GetOrganizationLabelssGroup(ctx, organizationId, labelsGroupId).Execute()

Get organization labels group



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
	labelsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization labels group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationLabelsGroupAPI.GetOrganizationLabelssGroup(context.Background(), organizationId, labelsGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationLabelsGroupAPI.GetOrganizationLabelssGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationLabelssGroup`: OrganizationLabelsGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationLabelsGroupAPI.GetOrganizationLabelssGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**labelsGroupId** | **string** | Organization labels group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLabelssGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationLabelsGroupResponse**](OrganizationLabelsGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationLabelsGroup

> ListOrganizationLabelsGroup200Response ListOrganizationLabelsGroup(ctx, organizationId).Execute()

List organization labels group



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
	resp, r, err := apiClient.OrganizationLabelsGroupAPI.ListOrganizationLabelsGroup(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationLabelsGroupAPI.ListOrganizationLabelsGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationLabelsGroup`: ListOrganizationLabelsGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationLabelsGroupAPI.ListOrganizationLabelsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationLabelsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListOrganizationLabelsGroup200Response**](ListOrganizationLabelsGroup200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

