# InviteMemberAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Role** | [**InviteMemberRoleEnum**](InviteMemberRoleEnum.md) |  | 
**InvitationLink** | **string** |  | 
**InvitationStatus** | [**InviteStatusEnum**](InviteStatusEnum.md) |  | 
**Inviter** | **string** |  | 
**LogoUrl** | Pointer to **string** |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 

## Methods

### NewInviteMemberAllOf

`func NewInviteMemberAllOf(email string, role InviteMemberRoleEnum, invitationLink string, invitationStatus InviteStatusEnum, inviter string, ) *InviteMemberAllOf`

NewInviteMemberAllOf instantiates a new InviteMemberAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteMemberAllOfWithDefaults

`func NewInviteMemberAllOfWithDefaults() *InviteMemberAllOf`

NewInviteMemberAllOfWithDefaults instantiates a new InviteMemberAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteMemberAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteMemberAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteMemberAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *InviteMemberAllOf) GetRole() InviteMemberRoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InviteMemberAllOf) GetRoleOk() (*InviteMemberRoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InviteMemberAllOf) SetRole(v InviteMemberRoleEnum)`

SetRole sets Role field to given value.


### GetInvitationLink

`func (o *InviteMemberAllOf) GetInvitationLink() string`

GetInvitationLink returns the InvitationLink field if non-nil, zero value otherwise.

### GetInvitationLinkOk

`func (o *InviteMemberAllOf) GetInvitationLinkOk() (*string, bool)`

GetInvitationLinkOk returns a tuple with the InvitationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationLink

`func (o *InviteMemberAllOf) SetInvitationLink(v string)`

SetInvitationLink sets InvitationLink field to given value.


### GetInvitationStatus

`func (o *InviteMemberAllOf) GetInvitationStatus() InviteStatusEnum`

GetInvitationStatus returns the InvitationStatus field if non-nil, zero value otherwise.

### GetInvitationStatusOk

`func (o *InviteMemberAllOf) GetInvitationStatusOk() (*InviteStatusEnum, bool)`

GetInvitationStatusOk returns a tuple with the InvitationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationStatus

`func (o *InviteMemberAllOf) SetInvitationStatus(v InviteStatusEnum)`

SetInvitationStatus sets InvitationStatus field to given value.


### GetInviter

`func (o *InviteMemberAllOf) GetInviter() string`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *InviteMemberAllOf) GetInviterOk() (*string, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *InviteMemberAllOf) SetInviter(v string)`

SetInviter sets Inviter field to given value.


### GetLogoUrl

`func (o *InviteMemberAllOf) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *InviteMemberAllOf) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *InviteMemberAllOf) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *InviteMemberAllOf) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetRoleId

`func (o *InviteMemberAllOf) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *InviteMemberAllOf) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *InviteMemberAllOf) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *InviteMemberAllOf) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleName

`func (o *InviteMemberAllOf) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *InviteMemberAllOf) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *InviteMemberAllOf) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *InviteMemberAllOf) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


