# MemberAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**ProfilePictureUrl** | Pointer to **string** |  | [optional] 
**LastActivityAt** | Pointer to **time.Time** | last time the user was connected | [optional] 
**Role** | Pointer to [**InviteMemberRoleEnum**](InviteMemberRoleEnum.md) |  | [optional] 
**RoleName** | Pointer to **string** | the role linked to the user | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 

## Methods

### NewMemberAllOf

`func NewMemberAllOf(email string, ) *MemberAllOf`

NewMemberAllOf instantiates a new MemberAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberAllOfWithDefaults

`func NewMemberAllOfWithDefaults() *MemberAllOf`

NewMemberAllOfWithDefaults instantiates a new MemberAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MemberAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemberAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *MemberAllOf) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *MemberAllOf) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *MemberAllOf) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *MemberAllOf) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetEmail

`func (o *MemberAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MemberAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MemberAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProfilePictureUrl

`func (o *MemberAllOf) GetProfilePictureUrl() string`

GetProfilePictureUrl returns the ProfilePictureUrl field if non-nil, zero value otherwise.

### GetProfilePictureUrlOk

`func (o *MemberAllOf) GetProfilePictureUrlOk() (*string, bool)`

GetProfilePictureUrlOk returns a tuple with the ProfilePictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePictureUrl

`func (o *MemberAllOf) SetProfilePictureUrl(v string)`

SetProfilePictureUrl sets ProfilePictureUrl field to given value.

### HasProfilePictureUrl

`func (o *MemberAllOf) HasProfilePictureUrl() bool`

HasProfilePictureUrl returns a boolean if a field has been set.

### GetLastActivityAt

`func (o *MemberAllOf) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *MemberAllOf) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *MemberAllOf) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.

### HasLastActivityAt

`func (o *MemberAllOf) HasLastActivityAt() bool`

HasLastActivityAt returns a boolean if a field has been set.

### GetRole

`func (o *MemberAllOf) GetRole() InviteMemberRoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberAllOf) GetRoleOk() (*InviteMemberRoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberAllOf) SetRole(v InviteMemberRoleEnum)`

SetRole sets Role field to given value.

### HasRole

`func (o *MemberAllOf) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoleName

`func (o *MemberAllOf) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *MemberAllOf) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *MemberAllOf) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *MemberAllOf) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetRoleId

`func (o *MemberAllOf) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *MemberAllOf) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *MemberAllOf) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *MemberAllOf) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


