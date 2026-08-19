# OrganizationPolicyApiToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**RoleId** | Pointer to **string** | the organization role this token acts as once its policy has allowed a request. Effective access is the intersection of this role and &#x60;opa_policy&#x60;. | [optional] 
**OpaPolicy** | Pointer to **string** | the Open Policy Agent (rego) policy evaluated on every request made with this token | [optional] 
**CreatorName** | Pointer to **NullableString** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewOrganizationPolicyApiToken

`func NewOrganizationPolicyApiToken(id string, createdAt time.Time, ) *OrganizationPolicyApiToken`

NewOrganizationPolicyApiToken instantiates a new OrganizationPolicyApiToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationPolicyApiTokenWithDefaults

`func NewOrganizationPolicyApiTokenWithDefaults() *OrganizationPolicyApiToken`

NewOrganizationPolicyApiTokenWithDefaults instantiates a new OrganizationPolicyApiToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationPolicyApiToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationPolicyApiToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationPolicyApiToken) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrganizationPolicyApiToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationPolicyApiToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationPolicyApiToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationPolicyApiToken) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationPolicyApiToken) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationPolicyApiToken) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationPolicyApiToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *OrganizationPolicyApiToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationPolicyApiToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationPolicyApiToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationPolicyApiToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationPolicyApiToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationPolicyApiToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationPolicyApiToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationPolicyApiToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRoleName

`func (o *OrganizationPolicyApiToken) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *OrganizationPolicyApiToken) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *OrganizationPolicyApiToken) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *OrganizationPolicyApiToken) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetRoleId

`func (o *OrganizationPolicyApiToken) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *OrganizationPolicyApiToken) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *OrganizationPolicyApiToken) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *OrganizationPolicyApiToken) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetOpaPolicy

`func (o *OrganizationPolicyApiToken) GetOpaPolicy() string`

GetOpaPolicy returns the OpaPolicy field if non-nil, zero value otherwise.

### GetOpaPolicyOk

`func (o *OrganizationPolicyApiToken) GetOpaPolicyOk() (*string, bool)`

GetOpaPolicyOk returns a tuple with the OpaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaPolicy

`func (o *OrganizationPolicyApiToken) SetOpaPolicy(v string)`

SetOpaPolicy sets OpaPolicy field to given value.

### HasOpaPolicy

`func (o *OrganizationPolicyApiToken) HasOpaPolicy() bool`

HasOpaPolicy returns a boolean if a field has been set.

### GetCreatorName

`func (o *OrganizationPolicyApiToken) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *OrganizationPolicyApiToken) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *OrganizationPolicyApiToken) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *OrganizationPolicyApiToken) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### SetCreatorNameNil

`func (o *OrganizationPolicyApiToken) SetCreatorNameNil(b bool)`

 SetCreatorNameNil sets the value for CreatorName to be an explicit nil

### UnsetCreatorName
`func (o *OrganizationPolicyApiToken) UnsetCreatorName()`

UnsetCreatorName ensures that no value is present for CreatorName, not even an explicit nil
### GetExpiresAt

`func (o *OrganizationPolicyApiToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrganizationPolicyApiToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrganizationPolicyApiToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrganizationPolicyApiToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *OrganizationPolicyApiToken) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *OrganizationPolicyApiToken) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


