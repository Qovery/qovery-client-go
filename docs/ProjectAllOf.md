# ProjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**ReferenceObject**](ReferenceObject.md) |  | [optional] 

## Methods

### NewProjectAllOf

`func NewProjectAllOf(name string, ) *ProjectAllOf`

NewProjectAllOf instantiates a new ProjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAllOfWithDefaults

`func NewProjectAllOfWithDefaults() *ProjectAllOf`

NewProjectAllOfWithDefaults instantiates a new ProjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProjectAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrganization

`func (o *ProjectAllOf) GetOrganization() ReferenceObject`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProjectAllOf) GetOrganizationOk() (*ReferenceObject, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProjectAllOf) SetOrganization(v ReferenceObject)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ProjectAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


