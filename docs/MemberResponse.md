# MemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**ProfilePictureUrl** | Pointer to **string** |  | [optional] 
**LastActivityAt** | Pointer to **time.Time** | last time the user was connected | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewMemberResponse

`func NewMemberResponse(email string, id string, createdAt time.Time, ) *MemberResponse`

NewMemberResponse instantiates a new MemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberResponseWithDefaults

`func NewMemberResponseWithDefaults() *MemberResponse`

NewMemberResponseWithDefaults instantiates a new MemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MemberResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemberResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *MemberResponse) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *MemberResponse) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *MemberResponse) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *MemberResponse) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetEmail

`func (o *MemberResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MemberResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MemberResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProfilePictureUrl

`func (o *MemberResponse) GetProfilePictureUrl() string`

GetProfilePictureUrl returns the ProfilePictureUrl field if non-nil, zero value otherwise.

### GetProfilePictureUrlOk

`func (o *MemberResponse) GetProfilePictureUrlOk() (*string, bool)`

GetProfilePictureUrlOk returns a tuple with the ProfilePictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePictureUrl

`func (o *MemberResponse) SetProfilePictureUrl(v string)`

SetProfilePictureUrl sets ProfilePictureUrl field to given value.

### HasProfilePictureUrl

`func (o *MemberResponse) HasProfilePictureUrl() bool`

HasProfilePictureUrl returns a boolean if a field has been set.

### GetLastActivityAt

`func (o *MemberResponse) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *MemberResponse) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *MemberResponse) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.

### HasLastActivityAt

`func (o *MemberResponse) HasLastActivityAt() bool`

HasLastActivityAt returns a boolean if a field has been set.

### GetRole

`func (o *MemberResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MemberResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetId

`func (o *MemberResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *MemberResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MemberResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MemberResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MemberResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MemberResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MemberResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MemberResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


