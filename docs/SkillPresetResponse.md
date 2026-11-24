# SkillPresetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Base** | [**SkillPresetBase**](SkillPresetBase.md) |  | 
**IsTemplate** | **bool** | Whether this is a template, never resolved for a member and never assigned. | 
**TeamContext** | Pointer to **NullableString** | Markdown delivered with every skill the agent loads. Null when the preset has none. | [optional] 
**Skills** | [**[]SkillPresetSkillResponse**](SkillPresetSkillResponse.md) |  | 
**Revision** | **int32** | Starts at 1 and increases by one on every edit. | 
**CreatorName** | Pointer to **NullableString** | Name of the member who created the preset. Null when unknown. | [optional] 
**AssignedTo** | [**SkillPresetAssignedTo**](SkillPresetAssignedTo.md) |  | 

## Methods

### NewSkillPresetResponse

`func NewSkillPresetResponse(id string, createdAt time.Time, updatedAt time.Time, name string, description string, base SkillPresetBase, isTemplate bool, skills []SkillPresetSkillResponse, revision int32, assignedTo SkillPresetAssignedTo, ) *SkillPresetResponse`

NewSkillPresetResponse instantiates a new SkillPresetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetResponseWithDefaults

`func NewSkillPresetResponseWithDefaults() *SkillPresetResponse`

NewSkillPresetResponseWithDefaults instantiates a new SkillPresetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SkillPresetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkillPresetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkillPresetResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *SkillPresetResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SkillPresetResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SkillPresetResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SkillPresetResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SkillPresetResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SkillPresetResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *SkillPresetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SkillPresetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SkillPresetResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SkillPresetResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SkillPresetResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SkillPresetResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetBase

`func (o *SkillPresetResponse) GetBase() SkillPresetBase`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *SkillPresetResponse) GetBaseOk() (*SkillPresetBase, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *SkillPresetResponse) SetBase(v SkillPresetBase)`

SetBase sets Base field to given value.


### GetIsTemplate

`func (o *SkillPresetResponse) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *SkillPresetResponse) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *SkillPresetResponse) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.


### GetTeamContext

`func (o *SkillPresetResponse) GetTeamContext() string`

GetTeamContext returns the TeamContext field if non-nil, zero value otherwise.

### GetTeamContextOk

`func (o *SkillPresetResponse) GetTeamContextOk() (*string, bool)`

GetTeamContextOk returns a tuple with the TeamContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamContext

`func (o *SkillPresetResponse) SetTeamContext(v string)`

SetTeamContext sets TeamContext field to given value.

### HasTeamContext

`func (o *SkillPresetResponse) HasTeamContext() bool`

HasTeamContext returns a boolean if a field has been set.

### SetTeamContextNil

`func (o *SkillPresetResponse) SetTeamContextNil(b bool)`

 SetTeamContextNil sets the value for TeamContext to be an explicit nil

### UnsetTeamContext
`func (o *SkillPresetResponse) UnsetTeamContext()`

UnsetTeamContext ensures that no value is present for TeamContext, not even an explicit nil
### GetSkills

`func (o *SkillPresetResponse) GetSkills() []SkillPresetSkillResponse`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *SkillPresetResponse) GetSkillsOk() (*[]SkillPresetSkillResponse, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *SkillPresetResponse) SetSkills(v []SkillPresetSkillResponse)`

SetSkills sets Skills field to given value.


### GetRevision

`func (o *SkillPresetResponse) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SkillPresetResponse) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SkillPresetResponse) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetCreatorName

`func (o *SkillPresetResponse) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *SkillPresetResponse) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *SkillPresetResponse) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *SkillPresetResponse) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### SetCreatorNameNil

`func (o *SkillPresetResponse) SetCreatorNameNil(b bool)`

 SetCreatorNameNil sets the value for CreatorName to be an explicit nil

### UnsetCreatorName
`func (o *SkillPresetResponse) UnsetCreatorName()`

UnsetCreatorName ensures that no value is present for CreatorName, not even an explicit nil
### GetAssignedTo

`func (o *SkillPresetResponse) GetAssignedTo() SkillPresetAssignedTo`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *SkillPresetResponse) GetAssignedToOk() (*SkillPresetAssignedTo, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *SkillPresetResponse) SetAssignedTo(v SkillPresetAssignedTo)`

SetAssignedTo sets AssignedTo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


