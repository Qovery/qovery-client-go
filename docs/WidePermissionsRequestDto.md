# WidePermissionsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsingRestrictivePermissions** | Pointer to **bool** |  | [optional] [default to false]
**WideScopeCloudAccountId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWidePermissionsRequestDto

`func NewWidePermissionsRequestDto() *WidePermissionsRequestDto`

NewWidePermissionsRequestDto instantiates a new WidePermissionsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidePermissionsRequestDtoWithDefaults

`func NewWidePermissionsRequestDtoWithDefaults() *WidePermissionsRequestDto`

NewWidePermissionsRequestDtoWithDefaults instantiates a new WidePermissionsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsingRestrictivePermissions

`func (o *WidePermissionsRequestDto) GetUsingRestrictivePermissions() bool`

GetUsingRestrictivePermissions returns the UsingRestrictivePermissions field if non-nil, zero value otherwise.

### GetUsingRestrictivePermissionsOk

`func (o *WidePermissionsRequestDto) GetUsingRestrictivePermissionsOk() (*bool, bool)`

GetUsingRestrictivePermissionsOk returns a tuple with the UsingRestrictivePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingRestrictivePermissions

`func (o *WidePermissionsRequestDto) SetUsingRestrictivePermissions(v bool)`

SetUsingRestrictivePermissions sets UsingRestrictivePermissions field to given value.

### HasUsingRestrictivePermissions

`func (o *WidePermissionsRequestDto) HasUsingRestrictivePermissions() bool`

HasUsingRestrictivePermissions returns a boolean if a field has been set.

### GetWideScopeCloudAccountId

`func (o *WidePermissionsRequestDto) GetWideScopeCloudAccountId() string`

GetWideScopeCloudAccountId returns the WideScopeCloudAccountId field if non-nil, zero value otherwise.

### GetWideScopeCloudAccountIdOk

`func (o *WidePermissionsRequestDto) GetWideScopeCloudAccountIdOk() (*string, bool)`

GetWideScopeCloudAccountIdOk returns a tuple with the WideScopeCloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWideScopeCloudAccountId

`func (o *WidePermissionsRequestDto) SetWideScopeCloudAccountId(v string)`

SetWideScopeCloudAccountId sets WideScopeCloudAccountId field to given value.

### HasWideScopeCloudAccountId

`func (o *WidePermissionsRequestDto) HasWideScopeCloudAccountId() bool`

HasWideScopeCloudAccountId returns a boolean if a field has been set.

### SetWideScopeCloudAccountIdNil

`func (o *WidePermissionsRequestDto) SetWideScopeCloudAccountIdNil(b bool)`

 SetWideScopeCloudAccountIdNil sets the value for WideScopeCloudAccountId to be an explicit nil

### UnsetWideScopeCloudAccountId
`func (o *WidePermissionsRequestDto) UnsetWideScopeCloudAccountId()`

UnsetWideScopeCloudAccountId ensures that no value is present for WideScopeCloudAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


