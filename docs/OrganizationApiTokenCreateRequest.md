# OrganizationApiTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Scope** | [**OrganizationApiTokenScope**](OrganizationApiTokenScope.md) |  | 

## Methods

### NewOrganizationApiTokenCreateRequest

`func NewOrganizationApiTokenCreateRequest(name string, scope OrganizationApiTokenScope, ) *OrganizationApiTokenCreateRequest`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


