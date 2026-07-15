# FieldSchemaConstraintsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedValues** | Pointer to **[]string** |  | [optional] 
**Min** | Pointer to **NullableInt64** |  | [optional] 
**Max** | Pointer to **NullableInt64** |  | [optional] 
**MinLength** | Pointer to **NullableInt64** |  | [optional] 
**MaxLength** | Pointer to **NullableInt64** |  | [optional] 
**Pattern** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFieldSchemaConstraintsResponse

`func NewFieldSchemaConstraintsResponse() *FieldSchemaConstraintsResponse`

NewFieldSchemaConstraintsResponse instantiates a new FieldSchemaConstraintsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldSchemaConstraintsResponseWithDefaults

`func NewFieldSchemaConstraintsResponseWithDefaults() *FieldSchemaConstraintsResponse`

NewFieldSchemaConstraintsResponseWithDefaults instantiates a new FieldSchemaConstraintsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedValues

`func (o *FieldSchemaConstraintsResponse) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *FieldSchemaConstraintsResponse) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *FieldSchemaConstraintsResponse) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *FieldSchemaConstraintsResponse) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *FieldSchemaConstraintsResponse) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *FieldSchemaConstraintsResponse) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetMin

`func (o *FieldSchemaConstraintsResponse) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *FieldSchemaConstraintsResponse) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *FieldSchemaConstraintsResponse) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *FieldSchemaConstraintsResponse) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *FieldSchemaConstraintsResponse) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *FieldSchemaConstraintsResponse) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *FieldSchemaConstraintsResponse) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *FieldSchemaConstraintsResponse) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *FieldSchemaConstraintsResponse) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *FieldSchemaConstraintsResponse) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *FieldSchemaConstraintsResponse) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *FieldSchemaConstraintsResponse) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetMinLength

`func (o *FieldSchemaConstraintsResponse) GetMinLength() int64`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *FieldSchemaConstraintsResponse) GetMinLengthOk() (*int64, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *FieldSchemaConstraintsResponse) SetMinLength(v int64)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *FieldSchemaConstraintsResponse) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *FieldSchemaConstraintsResponse) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *FieldSchemaConstraintsResponse) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetMaxLength

`func (o *FieldSchemaConstraintsResponse) GetMaxLength() int64`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *FieldSchemaConstraintsResponse) GetMaxLengthOk() (*int64, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *FieldSchemaConstraintsResponse) SetMaxLength(v int64)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *FieldSchemaConstraintsResponse) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *FieldSchemaConstraintsResponse) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *FieldSchemaConstraintsResponse) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetPattern

`func (o *FieldSchemaConstraintsResponse) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *FieldSchemaConstraintsResponse) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *FieldSchemaConstraintsResponse) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *FieldSchemaConstraintsResponse) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *FieldSchemaConstraintsResponse) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *FieldSchemaConstraintsResponse) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


