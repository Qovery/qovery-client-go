# \SkillPresetsAPI

All URIs are relative to *https://api.qovery.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSkillPreset**](SkillPresetsAPI.md#CreateSkillPreset) | **Post** /organization/{organizationId}/skillPreset | Create a skill preset
[**DeleteSkillPreset**](SkillPresetsAPI.md#DeleteSkillPreset) | **Delete** /skillPreset/{skillPresetId} | Delete a skill preset
[**EditSkillPreset**](SkillPresetsAPI.md#EditSkillPreset) | **Put** /skillPreset/{skillPresetId} | Edit a skill preset
[**GetResolvedSkill**](SkillPresetsAPI.md#GetResolvedSkill) | **Get** /organization/{organizationId}/skillPreset/resolved/skill/{skillName} | Get what to load for one skill
[**GetResolvedSkillPreset**](SkillPresetsAPI.md#GetResolvedSkillPreset) | **Get** /organization/{organizationId}/skillPreset/resolved | Get the skill preset that applies to the caller
[**GetSkillPreset**](SkillPresetsAPI.md#GetSkillPreset) | **Get** /skillPreset/{skillPresetId} | Get a skill preset
[**GetSkillUsage**](SkillPresetsAPI.md#GetSkillUsage) | **Get** /organization/{organizationId}/skillPreset/usage | Get skill usage
[**ListSkillPresetRevisions**](SkillPresetsAPI.md#ListSkillPresetRevisions) | **Get** /organization/{organizationId}/skillPreset/revision | List skill preset revisions
[**ListSkillPresets**](SkillPresetsAPI.md#ListSkillPresets) | **Get** /organization/{organizationId}/skillPreset | List organization skill presets



## CreateSkillPreset

> SkillPresetResponse CreateSkillPreset(ctx, organizationId).SkillPresetRequest(skillPresetRequest).Execute()

Create a skill preset



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
	skillPresetRequest := *openapiclient.NewSkillPresetRequest("backend") // SkillPresetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillPresetsAPI.CreateSkillPreset(context.Background(), organizationId).SkillPresetRequest(skillPresetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillPresetsAPI.CreateSkillPreset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSkillPreset`: SkillPresetResponse
	fmt.Fprintf(os.Stdout, "Response from `SkillPresetsAPI.CreateSkillPreset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSkillPresetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skillPresetRequest** | [**SkillPresetRequest**](SkillPresetRequest.md) |  | 

### Return type

[**SkillPresetResponse**](SkillPresetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSkillPreset

> DeleteSkillPreset(ctx, skillPresetId).Execute()

Delete a skill preset



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
	skillPresetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Skill Preset ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SkillPresetsAPI.DeleteSkillPreset(context.Background(), skillPresetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillPresetsAPI.DeleteSkillPreset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillPresetId** | **string** | Skill Preset ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSkillPresetRequest struct via the builder pattern


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


## EditSkillPreset

> SkillPresetResponse EditSkillPreset(ctx, skillPresetId).SkillPresetRequest(skillPresetRequest).Execute()

Edit a skill preset



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
	skillPresetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Skill Preset ID
	skillPresetRequest := *openapiclient.NewSkillPresetRequest("backend") // SkillPresetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillPresetsAPI.EditSkillPreset(context.Background(), skillPresetId).SkillPresetRequest(skillPresetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillPresetsAPI.EditSkillPreset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditSkillPreset`: SkillPresetResponse
	fmt.Fprintf(os.Stdout, "Response from `SkillPresetsAPI.EditSkillPreset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillPresetId** | **string** | Skill Preset ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditSkillPresetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skillPresetRequest** | [**SkillPresetRequest**](SkillPresetRequest.md) |  | 

### Return type

[**SkillPresetResponse**](SkillPresetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResolvedSkill

> SkillPresetResolvedSkillResponse GetResolvedSkill(ctx, organizationId, skillName).Purpose(purpose).IfNoneMatch(ifNoneMatch).Execute()

Get what to load for one skill



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
	skillName := "qovery-deploy" // string | Name of the skill to load, for example `qovery-deploy` or a skill the preset adds
	purpose := openapiclient.SkillLoadPurpose("LOAD") // SkillLoadPurpose | Who reads the skill. Defaults to LOAD. Any other value than LOAD or VIEW is refused with 400. (optional) (default to "LOAD")
	ifNoneMatch := ""role-0193f1c2-6b1e-7c3a-9a4f-2d5e8b7c1a90-3"" // string | The ETag of the caller's cached copy. When it matches the current ETag, the response is a 304 with no body. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillPresetsAPI.GetResolvedSkill(context.Background(), organizationId, skillName).Purpose(purpose).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillPresetsAPI.GetResolvedSkill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResolvedSkill`: SkillPresetResolvedSkillResponse
	fmt.Fprintf(os.Stdout, "Response from `SkillPresetsAPI.GetResolvedSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**skillName** | **string** | Name of the skill to load, for example &#x60;qovery-deploy&#x60; or a skill the preset adds | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResolvedSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **purpose** | [**SkillLoadPurpose**](SkillLoadPurpose.md) | Who reads the skill. Defaults to LOAD. Any other value than LOAD or VIEW is refused with 400. | [default to &quot;LOAD&quot;]
 **ifNoneMatch** | **string** | The ETag of the caller&#39;s cached copy. When it matches the current ETag, the response is a 304 with no body. | 

### Return type

[**SkillPresetResolvedSkillResponse**](SkillPresetResolvedSkillResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResolvedSkillPreset

> SkillPresetManifestResponse GetResolvedSkillPreset(ctx, organizationId).IfNoneMatch(ifNoneMatch).Execute()

Get the skill preset that applies to the caller



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
	ifNoneMatch := ""role-0193f1c2-6b1e-7c3a-9a4f-2d5e8b7c1a90-3"" // string | The ETag of the caller's cached copy. When it matches the current ETag, the response is a 304 with no body. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillPresetsAPI.GetResolvedSkillPreset(context.Background(), organizationId).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillPresetsAPI.GetResolvedSkillPreset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResolvedSkillPreset`: SkillPresetManifestResponse
	fmt.Fprintf(os.Stdout, "Response from `SkillPresetsAPI.GetResolvedSkillPreset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResolvedSkillPresetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifNoneMatch** | **string** | The ETag of the caller&#39;s cached copy. When it matches the current ETag, the response is a 304 with no body. | 

### Return type

[**SkillPresetManifestResponse**](SkillPresetManifestResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSkillPreset

> SkillPresetResponse GetSkillPreset(ctx, skillPresetId).Execute()

Get a skill preset



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
	skillPresetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Skill Preset ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillPresetsAPI.GetSkillPreset(context.Background(), skillPresetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillPresetsAPI.GetSkillPreset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSkillPreset`: SkillPresetResponse
	fmt.Fprintf(os.Stdout, "Response from `SkillPresetsAPI.GetSkillPreset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillPresetId** | **string** | Skill Preset ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillPresetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SkillPresetResponse**](SkillPresetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSkillUsage

> SkillUsageResponse GetSkillUsage(ctx, organizationId).From(from).To(to).Bucket(bucket).SkillName(skillName).UserSub(userSub).SkillPresetId(skillPresetId).Execute()

Get skill usage



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
	skillName := "skillName_example" // string | Only count the loads of this skill. (optional)
	userSub := "userSub_example" // string | Only count the loads of this member, by the id returned by \"List organization members\". (optional)
	skillPresetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only count the loads made under this preset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillPresetsAPI.GetSkillUsage(context.Background(), organizationId).From(from).To(to).Bucket(bucket).SkillName(skillName).UserSub(userSub).SkillPresetId(skillPresetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillPresetsAPI.GetSkillUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSkillUsage`: SkillUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `SkillPresetsAPI.GetSkillUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Start of the range, inclusive. Defaults to 30 days before &#x60;to&#x60;. | 
 **to** | **time.Time** | End of the range, exclusive. Defaults to now. Must be after &#x60;from&#x60;, and the range must not exceed 90 days, otherwise the request is refused with 400. | 
 **bucket** | [**SkillUsageBucket**](SkillUsageBucket.md) | Width of each point of the timeline. Defaults to DAY. | [default to &quot;DAY&quot;]
 **skillName** | **string** | Only count the loads of this skill. | 
 **userSub** | **string** | Only count the loads of this member, by the id returned by \&quot;List organization members\&quot;. | 
 **skillPresetId** | **string** | Only count the loads made under this preset. | 

### Return type

[**SkillUsageResponse**](SkillUsageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSkillPresetRevisions

> SkillPresetRevisionResponseList ListSkillPresetRevisions(ctx, organizationId).SkillPresetId(skillPresetId).Limit(limit).Execute()

List skill preset revisions



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
	skillPresetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only list the revisions of this preset, deleted or not. Omit it to list the revisions of every preset of the organization. (optional)
	limit := int32(56) // int32 | Maximum number of revisions to return. A value outside 1 to 100 is brought back into that range. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillPresetsAPI.ListSkillPresetRevisions(context.Background(), organizationId).SkillPresetId(skillPresetId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillPresetsAPI.ListSkillPresetRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSkillPresetRevisions`: SkillPresetRevisionResponseList
	fmt.Fprintf(os.Stdout, "Response from `SkillPresetsAPI.ListSkillPresetRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSkillPresetRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skillPresetId** | **string** | Only list the revisions of this preset, deleted or not. Omit it to list the revisions of every preset of the organization. | 
 **limit** | **int32** | Maximum number of revisions to return. A value outside 1 to 100 is brought back into that range. | [default to 50]

### Return type

[**SkillPresetRevisionResponseList**](SkillPresetRevisionResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSkillPresets

> SkillPresetResponseList ListSkillPresets(ctx, organizationId).IsTemplate(isTemplate).Execute()

List organization skill presets



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
	isTemplate := true // bool | `true` lists templates only, `false` lists assignable presets only. Omit it to list both. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillPresetsAPI.ListSkillPresets(context.Background(), organizationId).IsTemplate(isTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillPresetsAPI.ListSkillPresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSkillPresets`: SkillPresetResponseList
	fmt.Fprintf(os.Stdout, "Response from `SkillPresetsAPI.ListSkillPresets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSkillPresetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isTemplate** | **bool** | &#x60;true&#x60; lists templates only, &#x60;false&#x60; lists assignable presets only. Omit it to list both. | 

### Return type

[**SkillPresetResponseList**](SkillPresetResponseList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

