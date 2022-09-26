# MemberRoleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | specify the git provider user id | [optional] 
**CustomRoleId** | Pointer to **string** | used to specify an organization custom role, otherwise &#x60;null&#x60; | [optional] 
**DefaultRoleName** | Pointer to [**DefaultMemberRole**](DefaultMemberRole.md) |  | [optional] 

## Methods

### NewMemberRoleUpdateRequest

`func NewMemberRoleUpdateRequest() *MemberRoleUpdateRequest`

NewMemberRoleUpdateRequest instantiates a new MemberRoleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberRoleUpdateRequestWithDefaults

`func NewMemberRoleUpdateRequestWithDefaults() *MemberRoleUpdateRequest`

NewMemberRoleUpdateRequestWithDefaults instantiates a new MemberRoleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *MemberRoleUpdateRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MemberRoleUpdateRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MemberRoleUpdateRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MemberRoleUpdateRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCustomRoleId

`func (o *MemberRoleUpdateRequest) GetCustomRoleId() string`

GetCustomRoleId returns the CustomRoleId field if non-nil, zero value otherwise.

### GetCustomRoleIdOk

`func (o *MemberRoleUpdateRequest) GetCustomRoleIdOk() (*string, bool)`

GetCustomRoleIdOk returns a tuple with the CustomRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleId

`func (o *MemberRoleUpdateRequest) SetCustomRoleId(v string)`

SetCustomRoleId sets CustomRoleId field to given value.

### HasCustomRoleId

`func (o *MemberRoleUpdateRequest) HasCustomRoleId() bool`

HasCustomRoleId returns a boolean if a field has been set.

### GetDefaultRoleName

`func (o *MemberRoleUpdateRequest) GetDefaultRoleName() DefaultMemberRole`

GetDefaultRoleName returns the DefaultRoleName field if non-nil, zero value otherwise.

### GetDefaultRoleNameOk

`func (o *MemberRoleUpdateRequest) GetDefaultRoleNameOk() (*DefaultMemberRole, bool)`

GetDefaultRoleNameOk returns a tuple with the DefaultRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleName

`func (o *MemberRoleUpdateRequest) SetDefaultRoleName(v DefaultMemberRole)`

SetDefaultRoleName sets DefaultRoleName field to given value.

### HasDefaultRoleName

`func (o *MemberRoleUpdateRequest) HasDefaultRoleName() bool`

HasDefaultRoleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


