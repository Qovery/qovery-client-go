# OrganizationApiTokenCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | the generated token to send in &#39;Authorization&#39; header prefixed by &#39;Token &#39; | [optional] 
**Scope** | Pointer to [**OrganizationApiTokenScope**](OrganizationApiTokenScope.md) |  | [optional] 

## Methods

### NewOrganizationApiTokenCreateResponse

`func NewOrganizationApiTokenCreateResponse(id string, createdAt time.Time, ) *OrganizationApiTokenCreateResponse`

NewOrganizationApiTokenCreateResponse instantiates a new OrganizationApiTokenCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationApiTokenCreateResponseWithDefaults

`func NewOrganizationApiTokenCreateResponseWithDefaults() *OrganizationApiTokenCreateResponse`

NewOrganizationApiTokenCreateResponseWithDefaults instantiates a new OrganizationApiTokenCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationApiTokenCreateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationApiTokenCreateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationApiTokenCreateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrganizationApiTokenCreateResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationApiTokenCreateResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationApiTokenCreateResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationApiTokenCreateResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationApiTokenCreateResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationApiTokenCreateResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationApiTokenCreateResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *OrganizationApiTokenCreateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationApiTokenCreateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationApiTokenCreateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationApiTokenCreateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationApiTokenCreateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationApiTokenCreateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationApiTokenCreateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationApiTokenCreateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetToken

`func (o *OrganizationApiTokenCreateResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OrganizationApiTokenCreateResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OrganizationApiTokenCreateResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OrganizationApiTokenCreateResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetScope

`func (o *OrganizationApiTokenCreateResponse) GetScope() OrganizationApiTokenScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OrganizationApiTokenCreateResponse) GetScopeOk() (*OrganizationApiTokenScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OrganizationApiTokenCreateResponse) SetScope(v OrganizationApiTokenScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OrganizationApiTokenCreateResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


