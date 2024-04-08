# \JobAnnotationsGroupAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAnnotationsGroupToJob**](JobAnnotationsGroupAPI.md#AddAnnotationsGroupToJob) | **Post** /job/{jobId}/annotationsGroup/{annotationsGroupId} | Add annotations group to job
[**DeleteAnnotationsGroupToJob**](JobAnnotationsGroupAPI.md#DeleteAnnotationsGroupToJob) | **Delete** /job/{jobId}/annotationsGroup/{annotationsGroupId} | Delete annotations group to job
[**ListJobAnnotationsGroup**](JobAnnotationsGroupAPI.md#ListJobAnnotationsGroup) | **Get** /job/{jobId}/annotationsGroup | List job annotations group



## AddAnnotationsGroupToJob

> AddAnnotationsGroupToJob(ctx, jobId, annotationsGroupId).Execute()

Add annotations group to job



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
    annotationsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization annotations group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.JobAnnotationsGroupAPI.AddAnnotationsGroupToJob(context.Background(), jobId, annotationsGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobAnnotationsGroupAPI.AddAnnotationsGroupToJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 
**annotationsGroupId** | **string** | Organization annotations group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAnnotationsGroupToJobRequest struct via the builder pattern


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


## DeleteAnnotationsGroupToJob

> DeleteAnnotationsGroupToJob(ctx, jobId, annotationsGroupId).Execute()

Delete annotations group to job



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID
    annotationsGroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization annotations group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.JobAnnotationsGroupAPI.DeleteAnnotationsGroupToJob(context.Background(), jobId, annotationsGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobAnnotationsGroupAPI.DeleteAnnotationsGroupToJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 
**annotationsGroupId** | **string** | Organization annotations group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnnotationsGroupToJobRequest struct via the builder pattern


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


## ListJobAnnotationsGroup

> OrganizationAnnotationsGroupResponseList ListJobAnnotationsGroup(ctx, jobId).Execute()

List job annotations group



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobAnnotationsGroupAPI.ListJobAnnotationsGroup(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobAnnotationsGroupAPI.ListJobAnnotationsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJobAnnotationsGroup`: OrganizationAnnotationsGroupResponseList
    fmt.Fprintf(os.Stdout, "Response from `JobAnnotationsGroupAPI.ListJobAnnotationsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListJobAnnotationsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationAnnotationsGroupResponseList**](OrganizationAnnotationsGroupResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

