# OrganizationApiTokenCreateResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | the generated token to send in &#39;Authorization&#39; header prefixed by &#39;Token &#39; | [optional] 
**Scope** | Pointer to [**OrganizationApiTokenScope**](OrganizationApiTokenScope.md) |  | [optional] 

## Methods

### NewOrganizationApiTokenCreateResponseAllOf

`func NewOrganizationApiTokenCreateResponseAllOf() *OrganizationApiTokenCreateResponseAllOf`

NewOrganizationApiTokenCreateResponseAllOf instantiates a new OrganizationApiTokenCreateResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationApiTokenCreateResponseAllOfWithDefaults

`func NewOrganizationApiTokenCreateResponseAllOfWithDefaults() *OrganizationApiTokenCreateResponseAllOf`

NewOrganizationApiTokenCreateResponseAllOfWithDefaults instantiates a new OrganizationApiTokenCreateResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationApiTokenCreateResponseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationApiTokenCreateResponseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationApiTokenCreateResponseAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationApiTokenCreateResponseAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationApiTokenCreateResponseAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationApiTokenCreateResponseAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationApiTokenCreateResponseAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationApiTokenCreateResponseAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetToken

`func (o *OrganizationApiTokenCreateResponseAllOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OrganizationApiTokenCreateResponseAllOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OrganizationApiTokenCreateResponseAllOf) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OrganizationApiTokenCreateResponseAllOf) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetScope

`func (o *OrganizationApiTokenCreateResponseAllOf) GetScope() OrganizationApiTokenScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OrganizationApiTokenCreateResponseAllOf) GetScopeOk() (*OrganizationApiTokenScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OrganizationApiTokenCreateResponseAllOf) SetScope(v OrganizationApiTokenScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OrganizationApiTokenCreateResponseAllOf) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


