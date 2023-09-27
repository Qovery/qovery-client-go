# OrganizationApiTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to [**NullableOrganizationApiTokenScope**](OrganizationApiTokenScope.md) |  | [optional] 
**RoleId** | **NullableString** | the roleId provided by the \&quot;List organization custom roles\&quot; endpoint. | 

## Methods

### NewOrganizationApiTokenCreateRequest

`func NewOrganizationApiTokenCreateRequest(name string, roleId NullableString, ) *OrganizationApiTokenCreateRequest`

NewOrganizationApiTokenCreateRequest instantiates a new OrganizationApiTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationApiTokenCreateRequestWithDefaults

`func NewOrganizationApiTokenCreateRequestWithDefaults() *OrganizationApiTokenCreateRequest`

NewOrganizationApiTokenCreateRequestWithDefaults instantiates a new OrganizationApiTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationApiTokenCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationApiTokenCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationApiTokenCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OrganizationApiTokenCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationApiTokenCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationApiTokenCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationApiTokenCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScope

`func (o *OrganizationApiTokenCreateRequest) GetScope() OrganizationApiTokenScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OrganizationApiTokenCreateRequest) GetScopeOk() (*OrganizationApiTokenScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OrganizationApiTokenCreateRequest) SetScope(v OrganizationApiTokenScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OrganizationApiTokenCreateRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *OrganizationApiTokenCreateRequest) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *OrganizationApiTokenCreateRequest) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetRoleId

`func (o *OrganizationApiTokenCreateRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *OrganizationApiTokenCreateRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *OrganizationApiTokenCreateRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### SetRoleIdNil

`func (o *OrganizationApiTokenCreateRequest) SetRoleIdNil(b bool)`

 SetRoleIdNil sets the value for RoleId to be an explicit nil

### UnsetRoleId
`func (o *OrganizationApiTokenCreateRequest) UnsetRoleId()`

UnsetRoleId ensures that no value is present for RoleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


