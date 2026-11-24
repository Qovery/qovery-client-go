# SkillPresetSkillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillName** | **string** | Lowercase letters, digits and hyphens, at most 64 characters. A Qovery skill name (&#x60;qovery&#x60; or &#x60;qovery-*&#x60;) for DISABLE, OVERLAY and REPLACE; any other name for ADD. | 
**Mode** | [**SkillPresetSkillMode**](SkillPresetSkillMode.md) |  | 
**Description** | Pointer to **NullableString** | When the agent should use the skill, at most 1024 characters. Required for ADD and REPLACE, since agents pick a skill from its description. DISABLE and OVERLAY take none: a non-blank value is refused with 400. | [optional] 
**Content** | Pointer to **NullableString** | Markdown delivered to the agent, at most 64 KB. Required for every mode except DISABLE, which takes none: a non-blank value is refused with 400. | [optional] 

## Methods

### NewSkillPresetSkillRequest

`func NewSkillPresetSkillRequest(skillName string, mode SkillPresetSkillMode, ) *SkillPresetSkillRequest`

NewSkillPresetSkillRequest instantiates a new SkillPresetSkillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetSkillRequestWithDefaults

`func NewSkillPresetSkillRequestWithDefaults() *SkillPresetSkillRequest`

NewSkillPresetSkillRequestWithDefaults instantiates a new SkillPresetSkillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillName

`func (o *SkillPresetSkillRequest) GetSkillName() string`

GetSkillName returns the SkillName field if non-nil, zero value otherwise.

### GetSkillNameOk

`func (o *SkillPresetSkillRequest) GetSkillNameOk() (*string, bool)`

GetSkillNameOk returns a tuple with the SkillName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillName

`func (o *SkillPresetSkillRequest) SetSkillName(v string)`

SetSkillName sets SkillName field to given value.


### GetMode

`func (o *SkillPresetSkillRequest) GetMode() SkillPresetSkillMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SkillPresetSkillRequest) GetModeOk() (*SkillPresetSkillMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SkillPresetSkillRequest) SetMode(v SkillPresetSkillMode)`

SetMode sets Mode field to given value.


### GetDescription

`func (o *SkillPresetSkillRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SkillPresetSkillRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SkillPresetSkillRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SkillPresetSkillRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SkillPresetSkillRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SkillPresetSkillRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *SkillPresetSkillRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SkillPresetSkillRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SkillPresetSkillRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SkillPresetSkillRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *SkillPresetSkillRequest) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *SkillPresetSkillRequest) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


