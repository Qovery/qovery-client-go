# OrganizationAgentApiToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OpaPolicy** | Pointer to **string** | the Open Policy Agent (rego) policy evaluated on every request made with this token | [optional] 
**CreatorName** | Pointer to **NullableString** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewOrganizationAgentApiToken

`func NewOrganizationAgentApiToken(id string, createdAt time.Time, ) *OrganizationAgentApiToken`

NewOrganizationAgentApiToken instantiates a new OrganizationAgentApiToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationAgentApiTokenWithDefaults

`func NewOrganizationAgentApiTokenWithDefaults() *OrganizationAgentApiToken`

NewOrganizationAgentApiTokenWithDefaults instantiates a new OrganizationAgentApiToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationAgentApiToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationAgentApiToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationAgentApiToken) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrganizationAgentApiToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationAgentApiToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationAgentApiToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationAgentApiToken) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationAgentApiToken) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationAgentApiToken) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationAgentApiToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *OrganizationAgentApiToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationAgentApiToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationAgentApiToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationAgentApiToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationAgentApiToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationAgentApiToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationAgentApiToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationAgentApiToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOpaPolicy

`func (o *OrganizationAgentApiToken) GetOpaPolicy() string`

GetOpaPolicy returns the OpaPolicy field if non-nil, zero value otherwise.

### GetOpaPolicyOk

`func (o *OrganizationAgentApiToken) GetOpaPolicyOk() (*string, bool)`

GetOpaPolicyOk returns a tuple with the OpaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaPolicy

`func (o *OrganizationAgentApiToken) SetOpaPolicy(v string)`

SetOpaPolicy sets OpaPolicy field to given value.

### HasOpaPolicy

`func (o *OrganizationAgentApiToken) HasOpaPolicy() bool`

HasOpaPolicy returns a boolean if a field has been set.

### GetCreatorName

`func (o *OrganizationAgentApiToken) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *OrganizationAgentApiToken) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *OrganizationAgentApiToken) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *OrganizationAgentApiToken) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### SetCreatorNameNil

`func (o *OrganizationAgentApiToken) SetCreatorNameNil(b bool)`

 SetCreatorNameNil sets the value for CreatorName to be an explicit nil

### UnsetCreatorName
`func (o *OrganizationAgentApiToken) UnsetCreatorName()`

UnsetCreatorName ensures that no value is present for CreatorName, not even an explicit nil
### GetExpiresAt

`func (o *OrganizationAgentApiToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrganizationAgentApiToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrganizationAgentApiToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrganizationAgentApiToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *OrganizationAgentApiToken) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *OrganizationAgentApiToken) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


