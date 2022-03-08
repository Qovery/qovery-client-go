# MemberResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**ProfilePictureUrl** | Pointer to **string** |  | [optional] 
**LastActivityAt** | Pointer to **time.Time** | last time the user was connected | [optional] 
**Role** | Pointer to [**InviteMemberRoleEnum**](InviteMemberRoleEnum.md) |  | [optional] 

## Methods

### NewMemberResponseAllOf

`func NewMemberResponseAllOf(email string, ) *MemberResponseAllOf`

NewMemberResponseAllOf instantiates a new MemberResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberResponseAllOfWithDefaults

`func NewMemberResponseAllOfWithDefaults() *MemberResponseAllOf`

NewMemberResponseAllOfWithDefaults instantiates a new MemberResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MemberResponseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberResponseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberResponseAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemberResponseAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *MemberResponseAllOf) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *MemberResponseAllOf) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *MemberResponseAllOf) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *MemberResponseAllOf) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetEmail

`func (o *MemberResponseAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MemberResponseAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MemberResponseAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProfilePictureUrl

`func (o *MemberResponseAllOf) GetProfilePictureUrl() string`

GetProfilePictureUrl returns the ProfilePictureUrl field if non-nil, zero value otherwise.

### GetProfilePictureUrlOk

`func (o *MemberResponseAllOf) GetProfilePictureUrlOk() (*string, bool)`

GetProfilePictureUrlOk returns a tuple with the ProfilePictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePictureUrl

`func (o *MemberResponseAllOf) SetProfilePictureUrl(v string)`

SetProfilePictureUrl sets ProfilePictureUrl field to given value.

### HasProfilePictureUrl

`func (o *MemberResponseAllOf) HasProfilePictureUrl() bool`

HasProfilePictureUrl returns a boolean if a field has been set.

### GetLastActivityAt

`func (o *MemberResponseAllOf) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *MemberResponseAllOf) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *MemberResponseAllOf) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.

### HasLastActivityAt

`func (o *MemberResponseAllOf) HasLastActivityAt() bool`

HasLastActivityAt returns a boolean if a field has been set.

### GetRole

`func (o *MemberResponseAllOf) GetRole() InviteMemberRoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberResponseAllOf) GetRoleOk() (*InviteMemberRoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberResponseAllOf) SetRole(v InviteMemberRoleEnum)`

SetRole sets Role field to given value.

### HasRole

`func (o *MemberResponseAllOf) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


