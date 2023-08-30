# \OrganizationEventApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationEventTargets**](OrganizationEventApi.md#GetOrganizationEventTargets) | **Get** /organization/{organizationId}/targets | Get available event targets to filter events
[**GetOrganizationEvents**](OrganizationEventApi.md#GetOrganizationEvents) | **Get** /organization/{organizationId}/events | Get all events inside the organization



## GetOrganizationEventTargets

> OrganizationEventTargetResponseList GetOrganizationEventTargets(ctx, organizationId).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).EventType(eventType).TargetType(targetType).TriggeredBy(triggeredBy).Origin(origin).ProjectId(projectId).EnvironmentId(environmentId).Execute()

Get available event targets to filter events



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
    fromTimestamp := "fromTimestamp_example" // string | Display targets available since this timestamp.   A range of date can be specified by using `from-timestamp` with `to-timestamp` The format is a timestamp with nano precision  (optional)
    toTimestamp := "toTimestamp_example" // string | Display targets triggered before this timestamp.   A range of date can be specified by using `to-timestamp` with `from-timestamp` The format is a timestamp with nano precision  (optional)
    eventType := openapiclient.OrganizationEventType("CREATE") // OrganizationEventType |  (optional)
    targetType := openapiclient.OrganizationEventTargetType("APPLICATION") // OrganizationEventTargetType |  (optional)
    triggeredBy := "triggeredBy_example" // string | Information about the owner of the event (user name / apitoken / automatic action) (optional)
    origin := openapiclient.OrganizationEventOrigin("API") // OrganizationEventOrigin |  (optional)
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mandatory when requesting an environment or a service (optional)
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mandatory when requesting a service (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationEventApi.GetOrganizationEventTargets(context.Background(), organizationId).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).EventType(eventType).TargetType(targetType).TriggeredBy(triggeredBy).Origin(origin).ProjectId(projectId).EnvironmentId(environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationEventApi.GetOrganizationEventTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationEventTargets`: OrganizationEventTargetResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationEventApi.GetOrganizationEventTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationEventTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromTimestamp** | **string** | Display targets available since this timestamp.   A range of date can be specified by using &#x60;from-timestamp&#x60; with &#x60;to-timestamp&#x60; The format is a timestamp with nano precision  | 
 **toTimestamp** | **string** | Display targets triggered before this timestamp.   A range of date can be specified by using &#x60;to-timestamp&#x60; with &#x60;from-timestamp&#x60; The format is a timestamp with nano precision  | 
 **eventType** | [**OrganizationEventType**](OrganizationEventType.md) |  | 
 **targetType** | [**OrganizationEventTargetType**](OrganizationEventTargetType.md) |  | 
 **triggeredBy** | **string** | Information about the owner of the event (user name / apitoken / automatic action) | 
 **origin** | [**OrganizationEventOrigin**](OrganizationEventOrigin.md) |  | 
 **projectId** | **string** | Mandatory when requesting an environment or a service | 
 **environmentId** | **string** | Mandatory when requesting a service | 

### Return type

[**OrganizationEventTargetResponseList**](OrganizationEventTargetResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationEvents

> OrganizationEventResponseList GetOrganizationEvents(ctx, organizationId).PageSize(pageSize).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).ContinueToken(continueToken).StepBackToken(stepBackToken).EventType(eventType).TargetType(targetType).TargetId(targetId).SubTargetType(subTargetType).TriggeredBy(triggeredBy).Origin(origin).Execute()

Get all events inside the organization



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
    pageSize := float32(8.14) // float32 | The number of events to display in the current page (optional) (default to 10)
    fromTimestamp := "fromTimestamp_example" // string | Display events triggered since this timestamp.   A range of date can be specified by using `from-timestamp` with `to-timestamp` The format is a timestamp with nano precision  (optional)
    toTimestamp := "toTimestamp_example" // string | Display events triggered before this timestamp.   A range of date can be specified by using `to-timestamp` with `from-timestamp` The format is a timestamp with nano precision  (optional)
    continueToken := "continueToken_example" // string | Token used to fetch the next page results The format is a timestamp with nano precision  (optional)
    stepBackToken := "stepBackToken_example" // string | Token used to fetch the previous page results The format is a timestamp with nano precision  (optional)
    eventType := openapiclient.OrganizationEventType("CREATE") // OrganizationEventType |  (optional)
    targetType := openapiclient.OrganizationEventTargetType("APPLICATION") // OrganizationEventTargetType |  (optional)
    targetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The target resource id to search.   Must be specified with the corresponding `target_type`  (optional)
    subTargetType := openapiclient.OrganizationEventSubTargetType("ADVANCED_SETTINGS") // OrganizationEventSubTargetType |  (optional)
    triggeredBy := "triggeredBy_example" // string | Information about the owner of the event (user name / apitoken / automatic action) (optional)
    origin := openapiclient.OrganizationEventOrigin("API") // OrganizationEventOrigin |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationEventApi.GetOrganizationEvents(context.Background(), organizationId).PageSize(pageSize).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).ContinueToken(continueToken).StepBackToken(stepBackToken).EventType(eventType).TargetType(targetType).TargetId(targetId).SubTargetType(subTargetType).TriggeredBy(triggeredBy).Origin(origin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationEventApi.GetOrganizationEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationEvents`: OrganizationEventResponseList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationEventApi.GetOrganizationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **float32** | The number of events to display in the current page | [default to 10]
 **fromTimestamp** | **string** | Display events triggered since this timestamp.   A range of date can be specified by using &#x60;from-timestamp&#x60; with &#x60;to-timestamp&#x60; The format is a timestamp with nano precision  | 
 **toTimestamp** | **string** | Display events triggered before this timestamp.   A range of date can be specified by using &#x60;to-timestamp&#x60; with &#x60;from-timestamp&#x60; The format is a timestamp with nano precision  | 
 **continueToken** | **string** | Token used to fetch the next page results The format is a timestamp with nano precision  | 
 **stepBackToken** | **string** | Token used to fetch the previous page results The format is a timestamp with nano precision  | 
 **eventType** | [**OrganizationEventType**](OrganizationEventType.md) |  | 
 **targetType** | [**OrganizationEventTargetType**](OrganizationEventTargetType.md) |  | 
 **targetId** | **string** | The target resource id to search.   Must be specified with the corresponding &#x60;target_type&#x60;  | 
 **subTargetType** | [**OrganizationEventSubTargetType**](OrganizationEventSubTargetType.md) |  | 
 **triggeredBy** | **string** | Information about the owner of the event (user name / apitoken / automatic action) | 
 **origin** | [**OrganizationEventOrigin**](OrganizationEventOrigin.md) |  | 

### Return type

[**OrganizationEventResponseList**](OrganizationEventResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

