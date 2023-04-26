# \OrganizationEventApi

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationEvents**](OrganizationEventApi.md#GetOrganizationEvents) | **Get** /organization/{organizationId}/events | Get all events inside the organization



## GetOrganizationEvents

> OrganizationEventResponseList GetOrganizationEvents(ctx, organizationId).PageSize(pageSize).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).EventType(eventType).TargetType(targetType).TargetId(targetId).SubTargetType(subTargetType).TriggeredBy(triggeredBy).Origin(origin).Execute()

Get all events inside the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
    pageSize := float32(8.14) // float32 | The number of events to display in the current page (optional) (default to 10)
    fromTimestamp := time.Now() // time.Time | Display events triggered since this timestamp.   A range of date can be specified by using `from-timestamp` with `to-timestamp`  (optional)
    toTimestamp := time.Now() // time.Time | Display events triggered before this timestamp.   A range of date can be specified by using `to-timestamp` with `from-timestamp`  (optional)
    eventType := openapiclient.OrganizationEventType("CREATE") // OrganizationEventType |  (optional)
    targetType := openapiclient.OrganizationEventTargetType("APPLICATION") // OrganizationEventTargetType |  (optional)
    targetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The target resource id to search.   Must be specified with the corresponding `target_type`  (optional)
    subTargetType := openapiclient.OrganizationEventSubTargetType("ADVANCED_SETTINGS") // OrganizationEventSubTargetType |  (optional)
    triggeredBy := "triggeredBy_example" // string | Information about the owner of the event (user name / apitoken / automatic action) (optional)
    origin := openapiclient.OrganizationEventOrigin("API") // OrganizationEventOrigin |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationEventApi.GetOrganizationEvents(context.Background(), organizationId).PageSize(pageSize).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).EventType(eventType).TargetType(targetType).TargetId(targetId).SubTargetType(subTargetType).TriggeredBy(triggeredBy).Origin(origin).Execute()
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
 **fromTimestamp** | **time.Time** | Display events triggered since this timestamp.   A range of date can be specified by using &#x60;from-timestamp&#x60; with &#x60;to-timestamp&#x60;  | 
 **toTimestamp** | **time.Time** | Display events triggered before this timestamp.   A range of date can be specified by using &#x60;to-timestamp&#x60; with &#x60;from-timestamp&#x60;  | 
 **eventType** | [**OrganizationEventType**](OrganizationEventType.md) |  | 
 **targetType** | [**OrganizationEventTargetType**](OrganizationEventTargetType.md) |  | 
 **targetId** | **string** | The target resource id to search.   Must be specified with the corresponding &#x60;target_type&#x60;  | 
 **subTargetType** | [**OrganizationEventSubTargetType**](OrganizationEventSubTargetType.md) |  | 
 **triggeredBy** | **string** | Information about the owner of the event (user name / apitoken / automatic action) | 
 **origin** | [**OrganizationEventOrigin**](OrganizationEventOrigin.md) |  | 

### Return type

[**OrganizationEventResponseList**](OrganizationEventResponseList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

