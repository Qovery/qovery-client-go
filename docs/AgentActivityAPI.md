# \AgentActivityAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentActivity**](AgentActivityAPI.md#GetAgentActivity) | **Get** /organization/{organizationId}/agentActivity | Get agent activity
[**GetMyAgentActivity**](AgentActivityAPI.md#GetMyAgentActivity) | **Get** /organization/{organizationId}/agentActivity/me | Get my agent activity



## GetAgentActivity

> AgentActivityResponse GetAgentActivity(ctx, organizationId).From(from).To(to).Bucket(bucket).UserSub(userSub).SkillName(skillName).ActorKind(actorKind).PolicyTokenId(policyTokenId).Outcome(outcome).Source(source).Execute()

Get agent activity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	from := time.Now() // time.Time | Start of the range, inclusive. Defaults to 30 days before `to`. (optional)
	to := time.Now() // time.Time | End of the range, exclusive. Defaults to now. Must be after `from`, and the range must not exceed 90 days, otherwise the request is refused with 400. (optional)
	bucket := openapiclient.SkillUsageBucket("HOUR") // SkillUsageBucket | Width of each point of the timeline. Defaults to DAY. (optional) (default to "DAY")
	userSub := "userSub_example" // string | Only show this member, by the id returned by \"List organization members\". The member is matched on the username of their most recent skill load, so a member with no skill load in the last 90 days shows no action. (optional)
	skillName := "skillName_example" // string | Only count the activity attributed to this skill. Lowercase letters, digits, `.`, `_` or `-`, at most 64 characters, otherwise the request is refused with 400. (optional)
	actorKind := openapiclient.AgentActorKind("MEMBER") // AgentActorKind | Only count members (MEMBER) or policy tokens (POLICY_TOKEN). Cannot be combined with `userSub` when POLICY_TOKEN, otherwise the request is refused with 400. (optional)
	policyTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only count this policy token, by its id. Implies `actorKind` POLICY_TOKEN. Cannot be combined with `actorKind` MEMBER or with `userSub`, otherwise the request is refused with 400. With `bucket` HOUR, it is required to fill the token's `timeline` in `by_policy_token`. (optional)
	outcome := openapiclient.AgentActionOutcome("ALLOWED") // AgentActionOutcome | DENIED lists refusals only, and leaves skill loads out. ALLOWED lists everything else, skill loads included. (optional)
	source := openapiclient.AgentActivitySource("SKILL_LOAD") // AgentActivitySource | Only count the activity from this source. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentActivityAPI.GetAgentActivity(context.Background(), organizationId).From(from).To(to).Bucket(bucket).UserSub(userSub).SkillName(skillName).ActorKind(actorKind).PolicyTokenId(policyTokenId).Outcome(outcome).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentActivityAPI.GetAgentActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentActivity`: AgentActivityResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentActivityAPI.GetAgentActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Start of the range, inclusive. Defaults to 30 days before &#x60;to&#x60;. | 
 **to** | **time.Time** | End of the range, exclusive. Defaults to now. Must be after &#x60;from&#x60;, and the range must not exceed 90 days, otherwise the request is refused with 400. | 
 **bucket** | [**SkillUsageBucket**](SkillUsageBucket.md) | Width of each point of the timeline. Defaults to DAY. | [default to &quot;DAY&quot;]
 **userSub** | **string** | Only show this member, by the id returned by \&quot;List organization members\&quot;. The member is matched on the username of their most recent skill load, so a member with no skill load in the last 90 days shows no action. | 
 **skillName** | **string** | Only count the activity attributed to this skill. Lowercase letters, digits, &#x60;.&#x60;, &#x60;_&#x60; or &#x60;-&#x60;, at most 64 characters, otherwise the request is refused with 400. | 
 **actorKind** | [**AgentActorKind**](AgentActorKind.md) | Only count members (MEMBER) or policy tokens (POLICY_TOKEN). Cannot be combined with &#x60;userSub&#x60; when POLICY_TOKEN, otherwise the request is refused with 400. | 
 **policyTokenId** | **string** | Only count this policy token, by its id. Implies &#x60;actorKind&#x60; POLICY_TOKEN. Cannot be combined with &#x60;actorKind&#x60; MEMBER or with &#x60;userSub&#x60;, otherwise the request is refused with 400. With &#x60;bucket&#x60; HOUR, it is required to fill the token&#39;s &#x60;timeline&#x60; in &#x60;by_policy_token&#x60;. | 
 **outcome** | [**AgentActionOutcome**](AgentActionOutcome.md) | DENIED lists refusals only, and leaves skill loads out. ALLOWED lists everything else, skill loads included. | 
 **source** | [**AgentActivitySource**](AgentActivitySource.md) | Only count the activity from this source. | 

### Return type

[**AgentActivityResponse**](AgentActivityResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyAgentActivity

> AgentActivityResponse GetMyAgentActivity(ctx, organizationId).From(from).To(to).Bucket(bucket).SkillName(skillName).Outcome(outcome).Source(source).Execute()

Get my agent activity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/qovery/qovery-client-go"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization ID
	from := time.Now() // time.Time | Start of the range, inclusive. Defaults to 30 days before `to`. (optional)
	to := time.Now() // time.Time | End of the range, exclusive. Defaults to now. Must be after `from`, and the range must not exceed 90 days, otherwise the request is refused with 400. (optional)
	bucket := openapiclient.SkillUsageBucket("HOUR") // SkillUsageBucket | Width of each point of the timeline. Defaults to DAY. (optional) (default to "DAY")
	skillName := "skillName_example" // string | Only count the activity attributed to this skill. Lowercase letters, digits, `.`, `_` or `-`, at most 64 characters, otherwise the request is refused with 400. (optional)
	outcome := openapiclient.AgentActionOutcome("ALLOWED") // AgentActionOutcome | DENIED lists refusals only, and leaves skill loads out. ALLOWED lists everything else, skill loads included. (optional)
	source := openapiclient.AgentActivitySource("SKILL_LOAD") // AgentActivitySource | Only count the activity from this source. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentActivityAPI.GetMyAgentActivity(context.Background(), organizationId).From(from).To(to).Bucket(bucket).SkillName(skillName).Outcome(outcome).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentActivityAPI.GetMyAgentActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyAgentActivity`: AgentActivityResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentActivityAPI.GetMyAgentActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyAgentActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Start of the range, inclusive. Defaults to 30 days before &#x60;to&#x60;. | 
 **to** | **time.Time** | End of the range, exclusive. Defaults to now. Must be after &#x60;from&#x60;, and the range must not exceed 90 days, otherwise the request is refused with 400. | 
 **bucket** | [**SkillUsageBucket**](SkillUsageBucket.md) | Width of each point of the timeline. Defaults to DAY. | [default to &quot;DAY&quot;]
 **skillName** | **string** | Only count the activity attributed to this skill. Lowercase letters, digits, &#x60;.&#x60;, &#x60;_&#x60; or &#x60;-&#x60;, at most 64 characters, otherwise the request is refused with 400. | 
 **outcome** | [**AgentActionOutcome**](AgentActionOutcome.md) | DENIED lists refusals only, and leaves skill loads out. ALLOWED lists everything else, skill loads included. | 
 **source** | [**AgentActivitySource**](AgentActivitySource.md) | Only count the activity from this source. | 

### Return type

[**AgentActivityResponse**](AgentActivityResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

