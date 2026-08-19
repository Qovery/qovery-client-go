# OrganizationPolicyApiTokenCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**RoleId** | Pointer to **string** | the organization role this token acts as once its policy has allowed a request. Effective access is the intersection of this role and its policy. | [optional] 
**Token** | Pointer to **string** | the generated token to send in &#39;Authorization&#39; header prefixed by &#39;Token &#39;. It is returned only here and cannot be retrieved afterwards. | [optional] 
**CreatorName** | Pointer to **NullableString** |  | [optional] 
**UserSub** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewOrganizationPolicyApiTokenCreate

`func NewOrganizationPolicyApiTokenCreate(id string, createdAt time.Time, ) *OrganizationPolicyApiTokenCreate`

NewOrganizationPolicyApiTokenCreate instantiates a new OrganizationPolicyApiTokenCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationPolicyApiTokenCreateWithDefaults

`func NewOrganizationPolicyApiTokenCreateWithDefaults() *OrganizationPolicyApiTokenCreate`

NewOrganizationPolicyApiTokenCreateWithDefaults instantiates a new OrganizationPolicyApiTokenCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationPolicyApiTokenCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationPolicyApiTokenCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationPolicyApiTokenCreate) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrganizationPolicyApiTokenCreate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationPolicyApiTokenCreate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationPolicyApiTokenCreate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationPolicyApiTokenCreate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationPolicyApiTokenCreate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationPolicyApiTokenCreate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationPolicyApiTokenCreate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *OrganizationPolicyApiTokenCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationPolicyApiTokenCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationPolicyApiTokenCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationPolicyApiTokenCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationPolicyApiTokenCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationPolicyApiTokenCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationPolicyApiTokenCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationPolicyApiTokenCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRoleName

`func (o *OrganizationPolicyApiTokenCreate) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *OrganizationPolicyApiTokenCreate) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *OrganizationPolicyApiTokenCreate) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *OrganizationPolicyApiTokenCreate) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetRoleId

`func (o *OrganizationPolicyApiTokenCreate) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *OrganizationPolicyApiTokenCreate) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *OrganizationPolicyApiTokenCreate) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *OrganizationPolicyApiTokenCreate) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetToken

`func (o *OrganizationPolicyApiTokenCreate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OrganizationPolicyApiTokenCreate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OrganizationPolicyApiTokenCreate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OrganizationPolicyApiTokenCreate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCreatorName

`func (o *OrganizationPolicyApiTokenCreate) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *OrganizationPolicyApiTokenCreate) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *OrganizationPolicyApiTokenCreate) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *OrganizationPolicyApiTokenCreate) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### SetCreatorNameNil

`func (o *OrganizationPolicyApiTokenCreate) SetCreatorNameNil(b bool)`

 SetCreatorNameNil sets the value for CreatorName to be an explicit nil

### UnsetCreatorName
`func (o *OrganizationPolicyApiTokenCreate) UnsetCreatorName()`

UnsetCreatorName ensures that no value is present for CreatorName, not even an explicit nil
### GetUserSub

`func (o *OrganizationPolicyApiTokenCreate) GetUserSub() string`

GetUserSub returns the UserSub field if non-nil, zero value otherwise.

### GetUserSubOk

`func (o *OrganizationPolicyApiTokenCreate) GetUserSubOk() (*string, bool)`

GetUserSubOk returns a tuple with the UserSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSub

`func (o *OrganizationPolicyApiTokenCreate) SetUserSub(v string)`

SetUserSub sets UserSub field to given value.

### HasUserSub

`func (o *OrganizationPolicyApiTokenCreate) HasUserSub() bool`

HasUserSub returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrganizationPolicyApiTokenCreate) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrganizationPolicyApiTokenCreate) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrganizationPolicyApiTokenCreate) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrganizationPolicyApiTokenCreate) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *OrganizationPolicyApiTokenCreate) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *OrganizationPolicyApiTokenCreate) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


