# OrganizationApiTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to [**OrganizationApiTokenScope**](OrganizationApiTokenScope.md) |  | [optional] 

## Methods

### NewOrganizationApiTokenResponse

`func NewOrganizationApiTokenResponse(id string, createdAt time.Time, ) *OrganizationApiTokenResponse`

NewOrganizationApiTokenResponse instantiates a new OrganizationApiTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationApiTokenResponseWithDefaults

`func NewOrganizationApiTokenResponseWithDefaults() *OrganizationApiTokenResponse`

NewOrganizationApiTokenResponseWithDefaults instantiates a new OrganizationApiTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationApiTokenResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationApiTokenResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationApiTokenResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrganizationApiTokenResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationApiTokenResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationApiTokenResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationApiTokenResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationApiTokenResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationApiTokenResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationApiTokenResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *OrganizationApiTokenResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationApiTokenResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationApiTokenResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationApiTokenResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationApiTokenResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationApiTokenResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationApiTokenResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationApiTokenResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScope

`func (o *OrganizationApiTokenResponse) GetScope() OrganizationApiTokenScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OrganizationApiTokenResponse) GetScopeOk() (*OrganizationApiTokenScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OrganizationApiTokenResponse) SetScope(v OrganizationApiTokenScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OrganizationApiTokenResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


