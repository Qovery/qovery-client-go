# InviteMemberResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Role** | [**InviteMemberRoleEnum**](InviteMemberRoleEnum.md) |  | 
**InvitationLink** | **string** |  | 
**InvitationStatus** | [**InviteStatusEnum**](InviteStatusEnum.md) |  | 
**Inviter** | **string** |  | 
**LogoUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewInviteMemberResponseAllOf

`func NewInviteMemberResponseAllOf(email string, role InviteMemberRoleEnum, invitationLink string, invitationStatus InviteStatusEnum, inviter string, ) *InviteMemberResponseAllOf`

NewInviteMemberResponseAllOf instantiates a new InviteMemberResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteMemberResponseAllOfWithDefaults

`func NewInviteMemberResponseAllOfWithDefaults() *InviteMemberResponseAllOf`

NewInviteMemberResponseAllOfWithDefaults instantiates a new InviteMemberResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteMemberResponseAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteMemberResponseAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteMemberResponseAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *InviteMemberResponseAllOf) GetRole() InviteMemberRoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InviteMemberResponseAllOf) GetRoleOk() (*InviteMemberRoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InviteMemberResponseAllOf) SetRole(v InviteMemberRoleEnum)`

SetRole sets Role field to given value.


### GetInvitationLink

`func (o *InviteMemberResponseAllOf) GetInvitationLink() string`

GetInvitationLink returns the InvitationLink field if non-nil, zero value otherwise.

### GetInvitationLinkOk

`func (o *InviteMemberResponseAllOf) GetInvitationLinkOk() (*string, bool)`

GetInvitationLinkOk returns a tuple with the InvitationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationLink

`func (o *InviteMemberResponseAllOf) SetInvitationLink(v string)`

SetInvitationLink sets InvitationLink field to given value.


### GetInvitationStatus

`func (o *InviteMemberResponseAllOf) GetInvitationStatus() InviteStatusEnum`

GetInvitationStatus returns the InvitationStatus field if non-nil, zero value otherwise.

### GetInvitationStatusOk

`func (o *InviteMemberResponseAllOf) GetInvitationStatusOk() (*InviteStatusEnum, bool)`

GetInvitationStatusOk returns a tuple with the InvitationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationStatus

`func (o *InviteMemberResponseAllOf) SetInvitationStatus(v InviteStatusEnum)`

SetInvitationStatus sets InvitationStatus field to given value.


### GetInviter

`func (o *InviteMemberResponseAllOf) GetInviter() string`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *InviteMemberResponseAllOf) GetInviterOk() (*string, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *InviteMemberResponseAllOf) SetInviter(v string)`

SetInviter sets Inviter field to given value.


### GetLogoUrl

`func (o *InviteMemberResponseAllOf) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *InviteMemberResponseAllOf) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *InviteMemberResponseAllOf) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *InviteMemberResponseAllOf) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


