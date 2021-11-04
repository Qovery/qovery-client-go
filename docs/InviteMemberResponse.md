# InviteMemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Role** | **string** |  | 
**InvitationLink** | **string** |  | 
**InvitationStatus** | **string** |  | 
**Inviter** | **string** |  | 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewInviteMemberResponse

`func NewInviteMemberResponse(email string, role string, invitationLink string, invitationStatus string, inviter string, id string, createdAt time.Time, ) *InviteMemberResponse`

NewInviteMemberResponse instantiates a new InviteMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteMemberResponseWithDefaults

`func NewInviteMemberResponseWithDefaults() *InviteMemberResponse`

NewInviteMemberResponseWithDefaults instantiates a new InviteMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteMemberResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteMemberResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteMemberResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *InviteMemberResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InviteMemberResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InviteMemberResponse) SetRole(v string)`

SetRole sets Role field to given value.


### GetInvitationLink

`func (o *InviteMemberResponse) GetInvitationLink() string`

GetInvitationLink returns the InvitationLink field if non-nil, zero value otherwise.

### GetInvitationLinkOk

`func (o *InviteMemberResponse) GetInvitationLinkOk() (*string, bool)`

GetInvitationLinkOk returns a tuple with the InvitationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationLink

`func (o *InviteMemberResponse) SetInvitationLink(v string)`

SetInvitationLink sets InvitationLink field to given value.


### GetInvitationStatus

`func (o *InviteMemberResponse) GetInvitationStatus() string`

GetInvitationStatus returns the InvitationStatus field if non-nil, zero value otherwise.

### GetInvitationStatusOk

`func (o *InviteMemberResponse) GetInvitationStatusOk() (*string, bool)`

GetInvitationStatusOk returns a tuple with the InvitationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationStatus

`func (o *InviteMemberResponse) SetInvitationStatus(v string)`

SetInvitationStatus sets InvitationStatus field to given value.


### GetInviter

`func (o *InviteMemberResponse) GetInviter() string`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *InviteMemberResponse) GetInviterOk() (*string, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *InviteMemberResponse) SetInviter(v string)`

SetInviter sets Inviter field to given value.


### GetLogoUrl

`func (o *InviteMemberResponse) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *InviteMemberResponse) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *InviteMemberResponse) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *InviteMemberResponse) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetId

`func (o *InviteMemberResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InviteMemberResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InviteMemberResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *InviteMemberResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InviteMemberResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InviteMemberResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InviteMemberResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InviteMemberResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InviteMemberResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InviteMemberResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


