# WidePermissionsResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsingRestrictivePermissions** | Pointer to **bool** |  | [optional] [default to false]
**WideScopeCloudAccountId** | Pointer to **NullableString** |  | [optional] 
**RequireWideAccessToDeploy** | Pointer to **bool** |  | [optional] [default to false]
**SlackWebhookUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWidePermissionsResponseDto

`func NewWidePermissionsResponseDto() *WidePermissionsResponseDto`

NewWidePermissionsResponseDto instantiates a new WidePermissionsResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidePermissionsResponseDtoWithDefaults

`func NewWidePermissionsResponseDtoWithDefaults() *WidePermissionsResponseDto`

NewWidePermissionsResponseDtoWithDefaults instantiates a new WidePermissionsResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsingRestrictivePermissions

`func (o *WidePermissionsResponseDto) GetUsingRestrictivePermissions() bool`

GetUsingRestrictivePermissions returns the UsingRestrictivePermissions field if non-nil, zero value otherwise.

### GetUsingRestrictivePermissionsOk

`func (o *WidePermissionsResponseDto) GetUsingRestrictivePermissionsOk() (*bool, bool)`

GetUsingRestrictivePermissionsOk returns a tuple with the UsingRestrictivePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingRestrictivePermissions

`func (o *WidePermissionsResponseDto) SetUsingRestrictivePermissions(v bool)`

SetUsingRestrictivePermissions sets UsingRestrictivePermissions field to given value.

### HasUsingRestrictivePermissions

`func (o *WidePermissionsResponseDto) HasUsingRestrictivePermissions() bool`

HasUsingRestrictivePermissions returns a boolean if a field has been set.

### GetWideScopeCloudAccountId

`func (o *WidePermissionsResponseDto) GetWideScopeCloudAccountId() string`

GetWideScopeCloudAccountId returns the WideScopeCloudAccountId field if non-nil, zero value otherwise.

### GetWideScopeCloudAccountIdOk

`func (o *WidePermissionsResponseDto) GetWideScopeCloudAccountIdOk() (*string, bool)`

GetWideScopeCloudAccountIdOk returns a tuple with the WideScopeCloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWideScopeCloudAccountId

`func (o *WidePermissionsResponseDto) SetWideScopeCloudAccountId(v string)`

SetWideScopeCloudAccountId sets WideScopeCloudAccountId field to given value.

### HasWideScopeCloudAccountId

`func (o *WidePermissionsResponseDto) HasWideScopeCloudAccountId() bool`

HasWideScopeCloudAccountId returns a boolean if a field has been set.

### SetWideScopeCloudAccountIdNil

`func (o *WidePermissionsResponseDto) SetWideScopeCloudAccountIdNil(b bool)`

 SetWideScopeCloudAccountIdNil sets the value for WideScopeCloudAccountId to be an explicit nil

### UnsetWideScopeCloudAccountId
`func (o *WidePermissionsResponseDto) UnsetWideScopeCloudAccountId()`

UnsetWideScopeCloudAccountId ensures that no value is present for WideScopeCloudAccountId, not even an explicit nil
### GetRequireWideAccessToDeploy

`func (o *WidePermissionsResponseDto) GetRequireWideAccessToDeploy() bool`

GetRequireWideAccessToDeploy returns the RequireWideAccessToDeploy field if non-nil, zero value otherwise.

### GetRequireWideAccessToDeployOk

`func (o *WidePermissionsResponseDto) GetRequireWideAccessToDeployOk() (*bool, bool)`

GetRequireWideAccessToDeployOk returns a tuple with the RequireWideAccessToDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireWideAccessToDeploy

`func (o *WidePermissionsResponseDto) SetRequireWideAccessToDeploy(v bool)`

SetRequireWideAccessToDeploy sets RequireWideAccessToDeploy field to given value.

### HasRequireWideAccessToDeploy

`func (o *WidePermissionsResponseDto) HasRequireWideAccessToDeploy() bool`

HasRequireWideAccessToDeploy returns a boolean if a field has been set.

### GetSlackWebhookUrl

`func (o *WidePermissionsResponseDto) GetSlackWebhookUrl() string`

GetSlackWebhookUrl returns the SlackWebhookUrl field if non-nil, zero value otherwise.

### GetSlackWebhookUrlOk

`func (o *WidePermissionsResponseDto) GetSlackWebhookUrlOk() (*string, bool)`

GetSlackWebhookUrlOk returns a tuple with the SlackWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWebhookUrl

`func (o *WidePermissionsResponseDto) SetSlackWebhookUrl(v string)`

SetSlackWebhookUrl sets SlackWebhookUrl field to given value.

### HasSlackWebhookUrl

`func (o *WidePermissionsResponseDto) HasSlackWebhookUrl() bool`

HasSlackWebhookUrl returns a boolean if a field has been set.

### SetSlackWebhookUrlNil

`func (o *WidePermissionsResponseDto) SetSlackWebhookUrlNil(b bool)`

 SetSlackWebhookUrlNil sets the value for SlackWebhookUrl to be an explicit nil

### UnsetSlackWebhookUrl
`func (o *WidePermissionsResponseDto) UnsetSlackWebhookUrl()`

UnsetSlackWebhookUrl ensures that no value is present for SlackWebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


