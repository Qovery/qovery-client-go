# SkillPresetResolvedSkillResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**SkillPresetResolutionSource**](SkillPresetResolutionSource.md) |  | 
**SkillName** | **string** |  | 
**SkillPreset** | Pointer to [**NullableSkillPresetResolvedSkillResponseSkillPreset**](SkillPresetResolvedSkillResponseSkillPreset.md) |  | [optional] 
**Mode** | Pointer to [**NullableSkillPresetSkillMode**](SkillPresetSkillMode.md) | Null when the preset leaves this Qovery skill alone, or when no preset applies: load the Qovery skill as published. DISABLE means the agent must not load the skill. | [optional] 
**Description** | Pointer to **NullableString** | The entry&#39;s description. Null when the preset leaves the skill alone or the entry has none. | [optional] 
**Content** | Pointer to **NullableString** | Markdown to load: team instructions after the Qovery skill for OVERLAY, the whole skill for REPLACE and ADD. Null for DISABLE and when the preset leaves the skill alone. | [optional] 
**ContentSha256** | Pointer to **NullableString** | Hex-encoded SHA-256 of &#x60;content&#x60;. Null when &#x60;content&#x60; is null. | [optional] 
**TeamContext** | Pointer to **NullableString** | The preset&#39;s team context, delivered with every skill. Null when there is none or no preset applies. | [optional] 
**TeamContextSha256** | Pointer to **NullableString** | Hex-encoded SHA-256 of &#x60;team_context&#x60;. Null when &#x60;team_context&#x60; is null. | [optional] 

## Methods

### NewSkillPresetResolvedSkillResponse

`func NewSkillPresetResolvedSkillResponse(source SkillPresetResolutionSource, skillName string, ) *SkillPresetResolvedSkillResponse`

NewSkillPresetResolvedSkillResponse instantiates a new SkillPresetResolvedSkillResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetResolvedSkillResponseWithDefaults

`func NewSkillPresetResolvedSkillResponseWithDefaults() *SkillPresetResolvedSkillResponse`

NewSkillPresetResolvedSkillResponseWithDefaults instantiates a new SkillPresetResolvedSkillResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *SkillPresetResolvedSkillResponse) GetSource() SkillPresetResolutionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SkillPresetResolvedSkillResponse) GetSourceOk() (*SkillPresetResolutionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SkillPresetResolvedSkillResponse) SetSource(v SkillPresetResolutionSource)`

SetSource sets Source field to given value.


### GetSkillName

`func (o *SkillPresetResolvedSkillResponse) GetSkillName() string`

GetSkillName returns the SkillName field if non-nil, zero value otherwise.

### GetSkillNameOk

`func (o *SkillPresetResolvedSkillResponse) GetSkillNameOk() (*string, bool)`

GetSkillNameOk returns a tuple with the SkillName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillName

`func (o *SkillPresetResolvedSkillResponse) SetSkillName(v string)`

SetSkillName sets SkillName field to given value.


### GetSkillPreset

`func (o *SkillPresetResolvedSkillResponse) GetSkillPreset() SkillPresetResolvedSkillResponseSkillPreset`

GetSkillPreset returns the SkillPreset field if non-nil, zero value otherwise.

### GetSkillPresetOk

`func (o *SkillPresetResolvedSkillResponse) GetSkillPresetOk() (*SkillPresetResolvedSkillResponseSkillPreset, bool)`

GetSkillPresetOk returns a tuple with the SkillPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillPreset

`func (o *SkillPresetResolvedSkillResponse) SetSkillPreset(v SkillPresetResolvedSkillResponseSkillPreset)`

SetSkillPreset sets SkillPreset field to given value.

### HasSkillPreset

`func (o *SkillPresetResolvedSkillResponse) HasSkillPreset() bool`

HasSkillPreset returns a boolean if a field has been set.

### SetSkillPresetNil

`func (o *SkillPresetResolvedSkillResponse) SetSkillPresetNil(b bool)`

 SetSkillPresetNil sets the value for SkillPreset to be an explicit nil

### UnsetSkillPreset
`func (o *SkillPresetResolvedSkillResponse) UnsetSkillPreset()`

UnsetSkillPreset ensures that no value is present for SkillPreset, not even an explicit nil
### GetMode

`func (o *SkillPresetResolvedSkillResponse) GetMode() SkillPresetSkillMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SkillPresetResolvedSkillResponse) GetModeOk() (*SkillPresetSkillMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SkillPresetResolvedSkillResponse) SetMode(v SkillPresetSkillMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SkillPresetResolvedSkillResponse) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *SkillPresetResolvedSkillResponse) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *SkillPresetResolvedSkillResponse) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetDescription

`func (o *SkillPresetResolvedSkillResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SkillPresetResolvedSkillResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SkillPresetResolvedSkillResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SkillPresetResolvedSkillResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SkillPresetResolvedSkillResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SkillPresetResolvedSkillResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *SkillPresetResolvedSkillResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SkillPresetResolvedSkillResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SkillPresetResolvedSkillResponse) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SkillPresetResolvedSkillResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *SkillPresetResolvedSkillResponse) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *SkillPresetResolvedSkillResponse) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetContentSha256

`func (o *SkillPresetResolvedSkillResponse) GetContentSha256() string`

GetContentSha256 returns the ContentSha256 field if non-nil, zero value otherwise.

### GetContentSha256Ok

`func (o *SkillPresetResolvedSkillResponse) GetContentSha256Ok() (*string, bool)`

GetContentSha256Ok returns a tuple with the ContentSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSha256

`func (o *SkillPresetResolvedSkillResponse) SetContentSha256(v string)`

SetContentSha256 sets ContentSha256 field to given value.

### HasContentSha256

`func (o *SkillPresetResolvedSkillResponse) HasContentSha256() bool`

HasContentSha256 returns a boolean if a field has been set.

### SetContentSha256Nil

`func (o *SkillPresetResolvedSkillResponse) SetContentSha256Nil(b bool)`

 SetContentSha256Nil sets the value for ContentSha256 to be an explicit nil

### UnsetContentSha256
`func (o *SkillPresetResolvedSkillResponse) UnsetContentSha256()`

UnsetContentSha256 ensures that no value is present for ContentSha256, not even an explicit nil
### GetTeamContext

`func (o *SkillPresetResolvedSkillResponse) GetTeamContext() string`

GetTeamContext returns the TeamContext field if non-nil, zero value otherwise.

### GetTeamContextOk

`func (o *SkillPresetResolvedSkillResponse) GetTeamContextOk() (*string, bool)`

GetTeamContextOk returns a tuple with the TeamContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamContext

`func (o *SkillPresetResolvedSkillResponse) SetTeamContext(v string)`

SetTeamContext sets TeamContext field to given value.

### HasTeamContext

`func (o *SkillPresetResolvedSkillResponse) HasTeamContext() bool`

HasTeamContext returns a boolean if a field has been set.

### SetTeamContextNil

`func (o *SkillPresetResolvedSkillResponse) SetTeamContextNil(b bool)`

 SetTeamContextNil sets the value for TeamContext to be an explicit nil

### UnsetTeamContext
`func (o *SkillPresetResolvedSkillResponse) UnsetTeamContext()`

UnsetTeamContext ensures that no value is present for TeamContext, not even an explicit nil
### GetTeamContextSha256

`func (o *SkillPresetResolvedSkillResponse) GetTeamContextSha256() string`

GetTeamContextSha256 returns the TeamContextSha256 field if non-nil, zero value otherwise.

### GetTeamContextSha256Ok

`func (o *SkillPresetResolvedSkillResponse) GetTeamContextSha256Ok() (*string, bool)`

GetTeamContextSha256Ok returns a tuple with the TeamContextSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamContextSha256

`func (o *SkillPresetResolvedSkillResponse) SetTeamContextSha256(v string)`

SetTeamContextSha256 sets TeamContextSha256 field to given value.

### HasTeamContextSha256

`func (o *SkillPresetResolvedSkillResponse) HasTeamContextSha256() bool`

HasTeamContextSha256 returns a boolean if a field has been set.

### SetTeamContextSha256Nil

`func (o *SkillPresetResolvedSkillResponse) SetTeamContextSha256Nil(b bool)`

 SetTeamContextSha256Nil sets the value for TeamContextSha256 to be an explicit nil

### UnsetTeamContextSha256
`func (o *SkillPresetResolvedSkillResponse) UnsetTeamContextSha256()`

UnsetTeamContextSha256 ensures that no value is present for TeamContextSha256, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


