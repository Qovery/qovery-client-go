# SkillPresetManifestSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillName** | **string** |  | 
**Mode** | [**SkillPresetSkillMode**](SkillPresetSkillMode.md) |  | 
**Description** | Pointer to **NullableString** | When the agent should use the skill. Null when the entry has none. | [optional] 
**Content** | Pointer to [**NullableSkillPresetManifestSkillContent**](SkillPresetManifestSkillContent.md) |  | [optional] 

## Methods

### NewSkillPresetManifestSkill

`func NewSkillPresetManifestSkill(skillName string, mode SkillPresetSkillMode, ) *SkillPresetManifestSkill`

NewSkillPresetManifestSkill instantiates a new SkillPresetManifestSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetManifestSkillWithDefaults

`func NewSkillPresetManifestSkillWithDefaults() *SkillPresetManifestSkill`

NewSkillPresetManifestSkillWithDefaults instantiates a new SkillPresetManifestSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillName

`func (o *SkillPresetManifestSkill) GetSkillName() string`

GetSkillName returns the SkillName field if non-nil, zero value otherwise.

### GetSkillNameOk

`func (o *SkillPresetManifestSkill) GetSkillNameOk() (*string, bool)`

GetSkillNameOk returns a tuple with the SkillName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillName

`func (o *SkillPresetManifestSkill) SetSkillName(v string)`

SetSkillName sets SkillName field to given value.


### GetMode

`func (o *SkillPresetManifestSkill) GetMode() SkillPresetSkillMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SkillPresetManifestSkill) GetModeOk() (*SkillPresetSkillMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SkillPresetManifestSkill) SetMode(v SkillPresetSkillMode)`

SetMode sets Mode field to given value.


### GetDescription

`func (o *SkillPresetManifestSkill) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SkillPresetManifestSkill) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SkillPresetManifestSkill) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SkillPresetManifestSkill) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SkillPresetManifestSkill) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SkillPresetManifestSkill) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *SkillPresetManifestSkill) GetContent() SkillPresetManifestSkillContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SkillPresetManifestSkill) GetContentOk() (*SkillPresetManifestSkillContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SkillPresetManifestSkill) SetContent(v SkillPresetManifestSkillContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *SkillPresetManifestSkill) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *SkillPresetManifestSkill) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *SkillPresetManifestSkill) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


