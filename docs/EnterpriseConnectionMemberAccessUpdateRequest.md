# EnterpriseConnectionMemberAccessUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**UserEmail** | Pointer to **NullableString** |  | [optional] 
**AddedOrganizationIds** | **[]string** |  | 
**RemovedOrganizationIds** | **[]string** |  | 
**RoleUpdatesByOrganizationId** | [**[]MemberAccessRoleUpdated**](MemberAccessRoleUpdated.md) |  | 

## Methods

### NewEnterpriseConnectionMemberAccessUpdateRequest

`func NewEnterpriseConnectionMemberAccessUpdateRequest(userId string, addedOrganizationIds []string, removedOrganizationIds []string, roleUpdatesByOrganizationId []MemberAccessRoleUpdated, ) *EnterpriseConnectionMemberAccessUpdateRequest`

NewEnterpriseConnectionMemberAccessUpdateRequest instantiates a new EnterpriseConnectionMemberAccessUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseConnectionMemberAccessUpdateRequestWithDefaults

`func NewEnterpriseConnectionMemberAccessUpdateRequestWithDefaults() *EnterpriseConnectionMemberAccessUpdateRequest`

NewEnterpriseConnectionMemberAccessUpdateRequestWithDefaults instantiates a new EnterpriseConnectionMemberAccessUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserEmail

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *EnterpriseConnectionMemberAccessUpdateRequest) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetAddedOrganizationIds

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) GetAddedOrganizationIds() []string`

GetAddedOrganizationIds returns the AddedOrganizationIds field if non-nil, zero value otherwise.

### GetAddedOrganizationIdsOk

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) GetAddedOrganizationIdsOk() (*[]string, bool)`

GetAddedOrganizationIdsOk returns a tuple with the AddedOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOrganizationIds

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) SetAddedOrganizationIds(v []string)`

SetAddedOrganizationIds sets AddedOrganizationIds field to given value.


### GetRemovedOrganizationIds

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) GetRemovedOrganizationIds() []string`

GetRemovedOrganizationIds returns the RemovedOrganizationIds field if non-nil, zero value otherwise.

### GetRemovedOrganizationIdsOk

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) GetRemovedOrganizationIdsOk() (*[]string, bool)`

GetRemovedOrganizationIdsOk returns a tuple with the RemovedOrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedOrganizationIds

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) SetRemovedOrganizationIds(v []string)`

SetRemovedOrganizationIds sets RemovedOrganizationIds field to given value.


### GetRoleUpdatesByOrganizationId

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) GetRoleUpdatesByOrganizationId() []MemberAccessRoleUpdated`

GetRoleUpdatesByOrganizationId returns the RoleUpdatesByOrganizationId field if non-nil, zero value otherwise.

### GetRoleUpdatesByOrganizationIdOk

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) GetRoleUpdatesByOrganizationIdOk() (*[]MemberAccessRoleUpdated, bool)`

GetRoleUpdatesByOrganizationIdOk returns a tuple with the RoleUpdatesByOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleUpdatesByOrganizationId

`func (o *EnterpriseConnectionMemberAccessUpdateRequest) SetRoleUpdatesByOrganizationId(v []MemberAccessRoleUpdated)`

SetRoleUpdatesByOrganizationId sets RoleUpdatesByOrganizationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


