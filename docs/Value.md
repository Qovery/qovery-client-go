# Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** | optional variable description (255 characters maximum) | [optional] 
**EnableInterpolationInFile** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewValue

`func NewValue() *Value`

NewValue instantiates a new Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueWithDefaults

`func NewValueWithDefaults() *Value`

NewValueWithDefaults instantiates a new Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Value) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Value) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Value) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Value) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDescription

`func (o *Value) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Value) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Value) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Value) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Value) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Value) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnableInterpolationInFile

`func (o *Value) GetEnableInterpolationInFile() bool`

GetEnableInterpolationInFile returns the EnableInterpolationInFile field if non-nil, zero value otherwise.

### GetEnableInterpolationInFileOk

`func (o *Value) GetEnableInterpolationInFileOk() (*bool, bool)`

GetEnableInterpolationInFileOk returns a tuple with the EnableInterpolationInFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInterpolationInFile

`func (o *Value) SetEnableInterpolationInFile(v bool)`

SetEnableInterpolationInFile sets EnableInterpolationInFile field to given value.

### HasEnableInterpolationInFile

`func (o *Value) HasEnableInterpolationInFile() bool`

HasEnableInterpolationInFile returns a boolean if a field has been set.

### SetEnableInterpolationInFileNil

`func (o *Value) SetEnableInterpolationInFileNil(b bool)`

 SetEnableInterpolationInFileNil sets the value for EnableInterpolationInFile to be an explicit nil

### UnsetEnableInterpolationInFile
`func (o *Value) UnsetEnableInterpolationInFile()`

UnsetEnableInterpolationInFile ensures that no value is present for EnableInterpolationInFile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


