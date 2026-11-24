# SkillPresetManifestResponseSkillPreset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Base** | [**SkillPresetBase**](SkillPresetBase.md) |  | 
**Revision** | **int32** |  | 
**UpdatedAt** | **time.Time** |  | 
**TeamContext** | Pointer to [**NullableSkillPresetManifestTeamContext**](SkillPresetManifestTeamContext.md) |  | [optional] 
**Skills** | [**[]SkillPresetManifestSkill**](SkillPresetManifestSkill.md) |  | 

## Methods

### NewSkillPresetManifestResponseSkillPreset

`func NewSkillPresetManifestResponseSkillPreset(id string, name string, base SkillPresetBase, revision int32, updatedAt time.Time, skills []SkillPresetManifestSkill, ) *SkillPresetManifestResponseSkillPreset`

NewSkillPresetManifestResponseSkillPreset instantiates a new SkillPresetManifestResponseSkillPreset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetManifestResponseSkillPresetWithDefaults

`func NewSkillPresetManifestResponseSkillPresetWithDefaults() *SkillPresetManifestResponseSkillPreset`

NewSkillPresetManifestResponseSkillPresetWithDefaults instantiates a new SkillPresetManifestResponseSkillPreset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SkillPresetManifestResponseSkillPreset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkillPresetManifestResponseSkillPreset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkillPresetManifestResponseSkillPreset) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SkillPresetManifestResponseSkillPreset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SkillPresetManifestResponseSkillPreset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SkillPresetManifestResponseSkillPreset) SetName(v string)`

SetName sets Name field to given value.


### GetBase

`func (o *SkillPresetManifestResponseSkillPreset) GetBase() SkillPresetBase`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *SkillPresetManifestResponseSkillPreset) GetBaseOk() (*SkillPresetBase, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *SkillPresetManifestResponseSkillPreset) SetBase(v SkillPresetBase)`

SetBase sets Base field to given value.


### GetRevision

`func (o *SkillPresetManifestResponseSkillPreset) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SkillPresetManifestResponseSkillPreset) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SkillPresetManifestResponseSkillPreset) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetUpdatedAt

`func (o *SkillPresetManifestResponseSkillPreset) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SkillPresetManifestResponseSkillPreset) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SkillPresetManifestResponseSkillPreset) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTeamContext

`func (o *SkillPresetManifestResponseSkillPreset) GetTeamContext() SkillPresetManifestTeamContext`

GetTeamContext returns the TeamContext field if non-nil, zero value otherwise.

### GetTeamContextOk

`func (o *SkillPresetManifestResponseSkillPreset) GetTeamContextOk() (*SkillPresetManifestTeamContext, bool)`

GetTeamContextOk returns a tuple with the TeamContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamContext

`func (o *SkillPresetManifestResponseSkillPreset) SetTeamContext(v SkillPresetManifestTeamContext)`

SetTeamContext sets TeamContext field to given value.

### HasTeamContext

`func (o *SkillPresetManifestResponseSkillPreset) HasTeamContext() bool`

HasTeamContext returns a boolean if a field has been set.

### SetTeamContextNil

`func (o *SkillPresetManifestResponseSkillPreset) SetTeamContextNil(b bool)`

 SetTeamContextNil sets the value for TeamContext to be an explicit nil

### UnsetTeamContext
`func (o *SkillPresetManifestResponseSkillPreset) UnsetTeamContext()`

UnsetTeamContext ensures that no value is present for TeamContext, not even an explicit nil
### GetSkills

`func (o *SkillPresetManifestResponseSkillPreset) GetSkills() []SkillPresetManifestSkill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *SkillPresetManifestResponseSkillPreset) GetSkillsOk() (*[]SkillPresetManifestSkill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *SkillPresetManifestResponseSkillPreset) SetSkills(v []SkillPresetManifestSkill)`

SetSkills sets Skills field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


