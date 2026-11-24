# SkillPresetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Preset name, unique in the organization, 1 to 100 characters | 
**Description** | Pointer to **string** | What the preset is for, at most 1024 characters. Shown to the members it applies to. | [optional] [default to ""]
**Base** | Pointer to [**SkillPresetBase**](SkillPresetBase.md) | The Qovery skill set the preset starts from. Defaults to FULL. | [optional] [default to SKILLPRESETBASE_FULL]
**IsTemplate** | Pointer to **bool** | Whether this is a template: a starting point admins copy into new presets. A template is never resolved for a member and cannot have assignments (400 otherwise). Defaults to false. | [optional] [default to false]
**TeamContext** | Pointer to **NullableString** | Markdown on how the team uses Qovery, delivered with every skill the agent loads, at most 32 KB. Null or blank means none. | [optional] 
**Skills** | Pointer to [**[]SkillPresetSkillRequest**](SkillPresetSkillRequest.md) | Changes on top of the base, at most 50, each skill name at most once. | [optional] [default to []]
**AssignedTo** | Pointer to [**SkillPresetAssignedTo**](SkillPresetAssignedTo.md) |  | [optional] 

## Methods

### NewSkillPresetRequest

`func NewSkillPresetRequest(name string, ) *SkillPresetRequest`

NewSkillPresetRequest instantiates a new SkillPresetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetRequestWithDefaults

`func NewSkillPresetRequestWithDefaults() *SkillPresetRequest`

NewSkillPresetRequestWithDefaults instantiates a new SkillPresetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SkillPresetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SkillPresetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SkillPresetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SkillPresetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SkillPresetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SkillPresetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SkillPresetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBase

`func (o *SkillPresetRequest) GetBase() SkillPresetBase`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *SkillPresetRequest) GetBaseOk() (*SkillPresetBase, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *SkillPresetRequest) SetBase(v SkillPresetBase)`

SetBase sets Base field to given value.

### HasBase

`func (o *SkillPresetRequest) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetIsTemplate

`func (o *SkillPresetRequest) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *SkillPresetRequest) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *SkillPresetRequest) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *SkillPresetRequest) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetTeamContext

`func (o *SkillPresetRequest) GetTeamContext() string`

GetTeamContext returns the TeamContext field if non-nil, zero value otherwise.

### GetTeamContextOk

`func (o *SkillPresetRequest) GetTeamContextOk() (*string, bool)`

GetTeamContextOk returns a tuple with the TeamContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamContext

`func (o *SkillPresetRequest) SetTeamContext(v string)`

SetTeamContext sets TeamContext field to given value.

### HasTeamContext

`func (o *SkillPresetRequest) HasTeamContext() bool`

HasTeamContext returns a boolean if a field has been set.

### SetTeamContextNil

`func (o *SkillPresetRequest) SetTeamContextNil(b bool)`

 SetTeamContextNil sets the value for TeamContext to be an explicit nil

### UnsetTeamContext
`func (o *SkillPresetRequest) UnsetTeamContext()`

UnsetTeamContext ensures that no value is present for TeamContext, not even an explicit nil
### GetSkills

`func (o *SkillPresetRequest) GetSkills() []SkillPresetSkillRequest`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *SkillPresetRequest) GetSkillsOk() (*[]SkillPresetSkillRequest, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *SkillPresetRequest) SetSkills(v []SkillPresetSkillRequest)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *SkillPresetRequest) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetAssignedTo

`func (o *SkillPresetRequest) GetAssignedTo() SkillPresetAssignedTo`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *SkillPresetRequest) GetAssignedToOk() (*SkillPresetAssignedTo, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *SkillPresetRequest) SetAssignedTo(v SkillPresetAssignedTo)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *SkillPresetRequest) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


