# MemberRoleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | specify the git provider user id | 
**RoleId** | **string** |  | 

## Methods

### NewMemberRoleUpdateRequest

`func NewMemberRoleUpdateRequest(userId string, roleId string, ) *MemberRoleUpdateRequest`

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


### GetRoleId

`func (o *MemberRoleUpdateRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *MemberRoleUpdateRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *MemberRoleUpdateRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


