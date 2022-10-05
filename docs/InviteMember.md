# InviteMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Email** | **string** |  | 
**Role** | [**InviteMemberRoleEnum**](InviteMemberRoleEnum.md) |  | 
**InvitationLink** | **string** |  | 
**InvitationStatus** | [**InviteStatusEnum**](InviteStatusEnum.md) |  | 
**Inviter** | **string** |  | 
**LogoUrl** | Pointer to **string** |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 

## Methods

### NewInviteMember

`func NewInviteMember(id string, createdAt time.Time, email string, role InviteMemberRoleEnum, invitationLink string, invitationStatus InviteStatusEnum, inviter string, ) *InviteMember`

NewInviteMember instantiates a new InviteMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteMemberWithDefaults

`func NewInviteMemberWithDefaults() *InviteMember`

NewInviteMemberWithDefaults instantiates a new InviteMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InviteMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InviteMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InviteMember) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *InviteMember) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InviteMember) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InviteMember) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InviteMember) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InviteMember) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InviteMember) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InviteMember) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *InviteMember) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteMember) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteMember) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *InviteMember) GetRole() InviteMemberRoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InviteMember) GetRoleOk() (*InviteMemberRoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InviteMember) SetRole(v InviteMemberRoleEnum)`

SetRole sets Role field to given value.


### GetInvitationLink

`func (o *InviteMember) GetInvitationLink() string`

GetInvitationLink returns the InvitationLink field if non-nil, zero value otherwise.

### GetInvitationLinkOk

`func (o *InviteMember) GetInvitationLinkOk() (*string, bool)`

GetInvitationLinkOk returns a tuple with the InvitationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationLink

`func (o *InviteMember) SetInvitationLink(v string)`

SetInvitationLink sets InvitationLink field to given value.


### GetInvitationStatus

`func (o *InviteMember) GetInvitationStatus() InviteStatusEnum`

GetInvitationStatus returns the InvitationStatus field if non-nil, zero value otherwise.

### GetInvitationStatusOk

`func (o *InviteMember) GetInvitationStatusOk() (*InviteStatusEnum, bool)`

GetInvitationStatusOk returns a tuple with the InvitationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationStatus

`func (o *InviteMember) SetInvitationStatus(v InviteStatusEnum)`

SetInvitationStatus sets InvitationStatus field to given value.


### GetInviter

`func (o *InviteMember) GetInviter() string`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *InviteMember) GetInviterOk() (*string, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *InviteMember) SetInviter(v string)`

SetInviter sets Inviter field to given value.


### GetLogoUrl

`func (o *InviteMember) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *InviteMember) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *InviteMember) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *InviteMember) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetRoleId

`func (o *InviteMember) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *InviteMember) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *InviteMember) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *InviteMember) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *InviteMember) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *InviteMember) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *InviteMember) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *InviteMember) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


