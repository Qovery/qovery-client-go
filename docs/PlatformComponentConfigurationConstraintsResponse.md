# PlatformComponentConfigurationConstraintsResponse

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

### NewPlatformComponentConfigurationConstraintsResponse

`func NewPlatformComponentConfigurationConstraintsResponse() *PlatformComponentConfigurationConstraintsResponse`

NewPlatformComponentConfigurationConstraintsResponse instantiates a new PlatformComponentConfigurationConstraintsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformComponentConfigurationConstraintsResponseWithDefaults

`func NewPlatformComponentConfigurationConstraintsResponseWithDefaults() *PlatformComponentConfigurationConstraintsResponse`

NewPlatformComponentConfigurationConstraintsResponseWithDefaults instantiates a new PlatformComponentConfigurationConstraintsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedValues

`func (o *PlatformComponentConfigurationConstraintsResponse) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *PlatformComponentConfigurationConstraintsResponse) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *PlatformComponentConfigurationConstraintsResponse) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *PlatformComponentConfigurationConstraintsResponse) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *PlatformComponentConfigurationConstraintsResponse) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *PlatformComponentConfigurationConstraintsResponse) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetMin

`func (o *PlatformComponentConfigurationConstraintsResponse) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *PlatformComponentConfigurationConstraintsResponse) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *PlatformComponentConfigurationConstraintsResponse) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *PlatformComponentConfigurationConstraintsResponse) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *PlatformComponentConfigurationConstraintsResponse) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *PlatformComponentConfigurationConstraintsResponse) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *PlatformComponentConfigurationConstraintsResponse) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *PlatformComponentConfigurationConstraintsResponse) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *PlatformComponentConfigurationConstraintsResponse) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *PlatformComponentConfigurationConstraintsResponse) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *PlatformComponentConfigurationConstraintsResponse) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *PlatformComponentConfigurationConstraintsResponse) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetMinLength

`func (o *PlatformComponentConfigurationConstraintsResponse) GetMinLength() int64`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *PlatformComponentConfigurationConstraintsResponse) GetMinLengthOk() (*int64, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *PlatformComponentConfigurationConstraintsResponse) SetMinLength(v int64)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *PlatformComponentConfigurationConstraintsResponse) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *PlatformComponentConfigurationConstraintsResponse) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *PlatformComponentConfigurationConstraintsResponse) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetMaxLength

`func (o *PlatformComponentConfigurationConstraintsResponse) GetMaxLength() int64`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *PlatformComponentConfigurationConstraintsResponse) GetMaxLengthOk() (*int64, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *PlatformComponentConfigurationConstraintsResponse) SetMaxLength(v int64)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *PlatformComponentConfigurationConstraintsResponse) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *PlatformComponentConfigurationConstraintsResponse) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *PlatformComponentConfigurationConstraintsResponse) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetPattern

`func (o *PlatformComponentConfigurationConstraintsResponse) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *PlatformComponentConfigurationConstraintsResponse) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *PlatformComponentConfigurationConstraintsResponse) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *PlatformComponentConfigurationConstraintsResponse) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *PlatformComponentConfigurationConstraintsResponse) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *PlatformComponentConfigurationConstraintsResponse) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


