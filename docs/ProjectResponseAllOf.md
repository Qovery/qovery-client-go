# ProjectResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**ReferenceObject**](ReferenceObject.md) |  | [optional] 

## Methods

### NewProjectResponseAllOf

`func NewProjectResponseAllOf(name string, ) *ProjectResponseAllOf`

NewProjectResponseAllOf instantiates a new ProjectResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectResponseAllOfWithDefaults

`func NewProjectResponseAllOfWithDefaults() *ProjectResponseAllOf`

NewProjectResponseAllOfWithDefaults instantiates a new ProjectResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectResponseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectResponseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectResponseAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProjectResponseAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectResponseAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectResponseAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectResponseAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrganization

`func (o *ProjectResponseAllOf) GetOrganization() ReferenceObject`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProjectResponseAllOf) GetOrganizationOk() (*ReferenceObject, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProjectResponseAllOf) SetOrganization(v ReferenceObject)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ProjectResponseAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


