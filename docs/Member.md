# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**ProfilePictureUrl** | Pointer to **string** |  | [optional] 
**LastActivityAt** | Pointer to **time.Time** | last time the user was connected | [optional] 
**Role** | Pointer to [**InviteMemberRoleEnum**](InviteMemberRoleEnum.md) |  | [optional] 
**RoleName** | Pointer to **string** | the role linked to the user | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 

## Methods

### NewMember

`func NewMember(id string, createdAt time.Time, email string, ) *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Member) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Member) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Member) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Member) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Member) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Member) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Member) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Member) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Member) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Member) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *Member) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Member) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Member) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Member) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *Member) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *Member) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *Member) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *Member) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetEmail

`func (o *Member) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Member) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Member) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProfilePictureUrl

`func (o *Member) GetProfilePictureUrl() string`

GetProfilePictureUrl returns the ProfilePictureUrl field if non-nil, zero value otherwise.

### GetProfilePictureUrlOk

`func (o *Member) GetProfilePictureUrlOk() (*string, bool)`

GetProfilePictureUrlOk returns a tuple with the ProfilePictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePictureUrl

`func (o *Member) SetProfilePictureUrl(v string)`

SetProfilePictureUrl sets ProfilePictureUrl field to given value.

### HasProfilePictureUrl

`func (o *Member) HasProfilePictureUrl() bool`

HasProfilePictureUrl returns a boolean if a field has been set.

### GetLastActivityAt

`func (o *Member) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *Member) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *Member) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.

### HasLastActivityAt

`func (o *Member) HasLastActivityAt() bool`

HasLastActivityAt returns a boolean if a field has been set.

### GetRole

`func (o *Member) GetRole() InviteMemberRoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Member) GetRoleOk() (*InviteMemberRoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Member) SetRole(v InviteMemberRoleEnum)`

SetRole sets Role field to given value.

### HasRole

`func (o *Member) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoleName

`func (o *Member) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *Member) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *Member) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *Member) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetRoleId

`func (o *Member) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *Member) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *Member) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *Member) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


