# SkillPresetManifest

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

### NewSkillPresetManifest

`func NewSkillPresetManifest(id string, name string, base SkillPresetBase, revision int32, updatedAt time.Time, skills []SkillPresetManifestSkill, ) *SkillPresetManifest`

NewSkillPresetManifest instantiates a new SkillPresetManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetManifestWithDefaults

`func NewSkillPresetManifestWithDefaults() *SkillPresetManifest`

NewSkillPresetManifestWithDefaults instantiates a new SkillPresetManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SkillPresetManifest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkillPresetManifest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkillPresetManifest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SkillPresetManifest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SkillPresetManifest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SkillPresetManifest) SetName(v string)`

SetName sets Name field to given value.


### GetBase

`func (o *SkillPresetManifest) GetBase() SkillPresetBase`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *SkillPresetManifest) GetBaseOk() (*SkillPresetBase, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *SkillPresetManifest) SetBase(v SkillPresetBase)`

SetBase sets Base field to given value.


### GetRevision

`func (o *SkillPresetManifest) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SkillPresetManifest) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SkillPresetManifest) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetUpdatedAt

`func (o *SkillPresetManifest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SkillPresetManifest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SkillPresetManifest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTeamContext

`func (o *SkillPresetManifest) GetTeamContext() SkillPresetManifestTeamContext`

GetTeamContext returns the TeamContext field if non-nil, zero value otherwise.

### GetTeamContextOk

`func (o *SkillPresetManifest) GetTeamContextOk() (*SkillPresetManifestTeamContext, bool)`

GetTeamContextOk returns a tuple with the TeamContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamContext

`func (o *SkillPresetManifest) SetTeamContext(v SkillPresetManifestTeamContext)`

SetTeamContext sets TeamContext field to given value.

### HasTeamContext

`func (o *SkillPresetManifest) HasTeamContext() bool`

HasTeamContext returns a boolean if a field has been set.

### SetTeamContextNil

`func (o *SkillPresetManifest) SetTeamContextNil(b bool)`

 SetTeamContextNil sets the value for TeamContext to be an explicit nil

### UnsetTeamContext
`func (o *SkillPresetManifest) UnsetTeamContext()`

UnsetTeamContext ensures that no value is present for TeamContext, not even an explicit nil
### GetSkills

`func (o *SkillPresetManifest) GetSkills() []SkillPresetManifestSkill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *SkillPresetManifest) GetSkillsOk() (*[]SkillPresetManifestSkill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *SkillPresetManifest) SetSkills(v []SkillPresetManifestSkill)`

SetSkills sets Skills field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


