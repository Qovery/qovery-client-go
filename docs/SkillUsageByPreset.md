# SkillUsageByPreset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillPresetId** | Pointer to **NullableString** | Null groups the loads made when no preset applied. | [optional] 
**SkillPresetName** | Pointer to **NullableString** | The preset&#39;s name at its most recent load in the range. Null when no preset applied. | [optional] 
**Loads** | **int64** |  | 

## Methods

### NewSkillUsageByPreset

`func NewSkillUsageByPreset(loads int64, ) *SkillUsageByPreset`

NewSkillUsageByPreset instantiates a new SkillUsageByPreset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillUsageByPresetWithDefaults

`func NewSkillUsageByPresetWithDefaults() *SkillUsageByPreset`

NewSkillUsageByPresetWithDefaults instantiates a new SkillUsageByPreset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillPresetId

`func (o *SkillUsageByPreset) GetSkillPresetId() string`

GetSkillPresetId returns the SkillPresetId field if non-nil, zero value otherwise.

### GetSkillPresetIdOk

`func (o *SkillUsageByPreset) GetSkillPresetIdOk() (*string, bool)`

GetSkillPresetIdOk returns a tuple with the SkillPresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillPresetId

`func (o *SkillUsageByPreset) SetSkillPresetId(v string)`

SetSkillPresetId sets SkillPresetId field to given value.

### HasSkillPresetId

`func (o *SkillUsageByPreset) HasSkillPresetId() bool`

HasSkillPresetId returns a boolean if a field has been set.

### SetSkillPresetIdNil

`func (o *SkillUsageByPreset) SetSkillPresetIdNil(b bool)`

 SetSkillPresetIdNil sets the value for SkillPresetId to be an explicit nil

### UnsetSkillPresetId
`func (o *SkillUsageByPreset) UnsetSkillPresetId()`

UnsetSkillPresetId ensures that no value is present for SkillPresetId, not even an explicit nil
### GetSkillPresetName

`func (o *SkillUsageByPreset) GetSkillPresetName() string`

GetSkillPresetName returns the SkillPresetName field if non-nil, zero value otherwise.

### GetSkillPresetNameOk

`func (o *SkillUsageByPreset) GetSkillPresetNameOk() (*string, bool)`

GetSkillPresetNameOk returns a tuple with the SkillPresetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillPresetName

`func (o *SkillUsageByPreset) SetSkillPresetName(v string)`

SetSkillPresetName sets SkillPresetName field to given value.

### HasSkillPresetName

`func (o *SkillUsageByPreset) HasSkillPresetName() bool`

HasSkillPresetName returns a boolean if a field has been set.

### SetSkillPresetNameNil

`func (o *SkillUsageByPreset) SetSkillPresetNameNil(b bool)`

 SetSkillPresetNameNil sets the value for SkillPresetName to be an explicit nil

### UnsetSkillPresetName
`func (o *SkillUsageByPreset) UnsetSkillPresetName()`

UnsetSkillPresetName ensures that no value is present for SkillPresetName, not even an explicit nil
### GetLoads

`func (o *SkillUsageByPreset) GetLoads() int64`

GetLoads returns the Loads field if non-nil, zero value otherwise.

### GetLoadsOk

`func (o *SkillUsageByPreset) GetLoadsOk() (*int64, bool)`

GetLoadsOk returns a tuple with the Loads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoads

`func (o *SkillUsageByPreset) SetLoads(v int64)`

SetLoads sets Loads field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


