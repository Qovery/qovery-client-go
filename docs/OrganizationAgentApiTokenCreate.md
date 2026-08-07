# OrganizationAgentApiTokenCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | the generated token to send in &#39;Authorization&#39; header prefixed by &#39;Token &#39;. It is returned only here and cannot be retrieved afterwards. | [optional] 
**CreatorName** | Pointer to **NullableString** |  | [optional] 
**UserSub** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewOrganizationAgentApiTokenCreate

`func NewOrganizationAgentApiTokenCreate(id string, createdAt time.Time, ) *OrganizationAgentApiTokenCreate`

NewOrganizationAgentApiTokenCreate instantiates a new OrganizationAgentApiTokenCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationAgentApiTokenCreateWithDefaults

`func NewOrganizationAgentApiTokenCreateWithDefaults() *OrganizationAgentApiTokenCreate`

NewOrganizationAgentApiTokenCreateWithDefaults instantiates a new OrganizationAgentApiTokenCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationAgentApiTokenCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationAgentApiTokenCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationAgentApiTokenCreate) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrganizationAgentApiTokenCreate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationAgentApiTokenCreate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationAgentApiTokenCreate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationAgentApiTokenCreate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationAgentApiTokenCreate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationAgentApiTokenCreate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationAgentApiTokenCreate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *OrganizationAgentApiTokenCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationAgentApiTokenCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationAgentApiTokenCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationAgentApiTokenCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationAgentApiTokenCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationAgentApiTokenCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationAgentApiTokenCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationAgentApiTokenCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetToken

`func (o *OrganizationAgentApiTokenCreate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OrganizationAgentApiTokenCreate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OrganizationAgentApiTokenCreate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OrganizationAgentApiTokenCreate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetCreatorName

`func (o *OrganizationAgentApiTokenCreate) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *OrganizationAgentApiTokenCreate) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *OrganizationAgentApiTokenCreate) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *OrganizationAgentApiTokenCreate) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### SetCreatorNameNil

`func (o *OrganizationAgentApiTokenCreate) SetCreatorNameNil(b bool)`

 SetCreatorNameNil sets the value for CreatorName to be an explicit nil

### UnsetCreatorName
`func (o *OrganizationAgentApiTokenCreate) UnsetCreatorName()`

UnsetCreatorName ensures that no value is present for CreatorName, not even an explicit nil
### GetUserSub

`func (o *OrganizationAgentApiTokenCreate) GetUserSub() string`

GetUserSub returns the UserSub field if non-nil, zero value otherwise.

### GetUserSubOk

`func (o *OrganizationAgentApiTokenCreate) GetUserSubOk() (*string, bool)`

GetUserSubOk returns a tuple with the UserSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSub

`func (o *OrganizationAgentApiTokenCreate) SetUserSub(v string)`

SetUserSub sets UserSub field to given value.

### HasUserSub

`func (o *OrganizationAgentApiTokenCreate) HasUserSub() bool`

HasUserSub returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrganizationAgentApiTokenCreate) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrganizationAgentApiTokenCreate) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrganizationAgentApiTokenCreate) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrganizationAgentApiTokenCreate) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *OrganizationAgentApiTokenCreate) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *OrganizationAgentApiTokenCreate) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


