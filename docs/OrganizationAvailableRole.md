# OrganizationAvailableRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Filled only for an organization custom role | [optional] 
**Name** | Pointer to **string** | It can be either a custom role name or a default role name | [optional] 
**IsDefault** | Pointer to **bool** | - &#x60;true&#x60; if it is a Qovery role - &#x60;false&#x60; if it is a custom role  | [optional] 

## Methods

### NewOrganizationAvailableRole

`func NewOrganizationAvailableRole() *OrganizationAvailableRole`

NewOrganizationAvailableRole instantiates a new OrganizationAvailableRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationAvailableRoleWithDefaults

`func NewOrganizationAvailableRoleWithDefaults() *OrganizationAvailableRole`

NewOrganizationAvailableRoleWithDefaults instantiates a new OrganizationAvailableRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationAvailableRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationAvailableRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationAvailableRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationAvailableRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationAvailableRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationAvailableRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationAvailableRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationAvailableRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsDefault

`func (o *OrganizationAvailableRole) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *OrganizationAvailableRole) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *OrganizationAvailableRole) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *OrganizationAvailableRole) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


