# VariableEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | the key of the environment variable | 
**Value** | Pointer to **NullableString** | the value of the environment variable | [optional] 
**Description** | Pointer to **NullableString** | optional variable description (255 characters maximum) | [optional] 
**EnableInterpolationInFile** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewVariableEditRequest

`func NewVariableEditRequest(key string, ) *VariableEditRequest`

NewVariableEditRequest instantiates a new VariableEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableEditRequestWithDefaults

`func NewVariableEditRequestWithDefaults() *VariableEditRequest`

NewVariableEditRequestWithDefaults instantiates a new VariableEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VariableEditRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VariableEditRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VariableEditRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *VariableEditRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableEditRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableEditRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *VariableEditRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *VariableEditRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *VariableEditRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetDescription

`func (o *VariableEditRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VariableEditRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VariableEditRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VariableEditRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VariableEditRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VariableEditRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnableInterpolationInFile

`func (o *VariableEditRequest) GetEnableInterpolationInFile() bool`

GetEnableInterpolationInFile returns the EnableInterpolationInFile field if non-nil, zero value otherwise.

### GetEnableInterpolationInFileOk

`func (o *VariableEditRequest) GetEnableInterpolationInFileOk() (*bool, bool)`

GetEnableInterpolationInFileOk returns a tuple with the EnableInterpolationInFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInterpolationInFile

`func (o *VariableEditRequest) SetEnableInterpolationInFile(v bool)`

SetEnableInterpolationInFile sets EnableInterpolationInFile field to given value.

### HasEnableInterpolationInFile

`func (o *VariableEditRequest) HasEnableInterpolationInFile() bool`

HasEnableInterpolationInFile returns a boolean if a field has been set.

### SetEnableInterpolationInFileNil

`func (o *VariableEditRequest) SetEnableInterpolationInFileNil(b bool)`

 SetEnableInterpolationInFileNil sets the value for EnableInterpolationInFile to be an explicit nil

### UnsetEnableInterpolationInFile
`func (o *VariableEditRequest) UnsetEnableInterpolationInFile()`

UnsetEnableInterpolationInFile ensures that no value is present for EnableInterpolationInFile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


