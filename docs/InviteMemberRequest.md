# InviteMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Role** | [**InviteMemberRoleEnum**](InviteMemberRoleEnum.md) |  | 

## Methods

### NewInviteMemberRequest

`func NewInviteMemberRequest(email string, role InviteMemberRoleEnum, ) *InviteMemberRequest`

NewInviteMemberRequest instantiates a new InviteMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteMemberRequestWithDefaults

`func NewInviteMemberRequestWithDefaults() *InviteMemberRequest`

NewInviteMemberRequestWithDefaults instantiates a new InviteMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteMemberRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteMemberRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteMemberRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *InviteMemberRequest) GetRole() InviteMemberRoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InviteMemberRequest) GetRoleOk() (*InviteMemberRoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InviteMemberRequest) SetRole(v InviteMemberRoleEnum)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


