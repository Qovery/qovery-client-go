# SkillPresetRevisionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillPresetId** | **string** | The preset&#39;s id. Still listed after the preset is deleted. | 
**Revision** | **int32** |  | 
**ChangeType** | [**SkillPresetChangeType**](SkillPresetChangeType.md) |  | 
**EditedBy** | **string** | Who saved this state. | 
**EditedAt** | **time.Time** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Base** | [**SkillPresetBase**](SkillPresetBase.md) |  | 
**IsTemplate** | **bool** | Whether the preset was a template at this revision. | 
**TeamContext** | Pointer to **NullableString** | Null when the preset had none at this revision. | [optional] 
**Skills** | [**[]SkillPresetSkillResponse**](SkillPresetSkillResponse.md) |  | 
**AssignedTo** | [**SkillPresetAssignedTo**](SkillPresetAssignedTo.md) |  | 
**SecurityReview** | Pointer to [**NullableSkillPresetRevisionResponseSecurityReview**](SkillPresetRevisionResponseSecurityReview.md) |  | [optional] 

## Methods

### NewSkillPresetRevisionResponse

`func NewSkillPresetRevisionResponse(skillPresetId string, revision int32, changeType SkillPresetChangeType, editedBy string, editedAt time.Time, name string, description string, base SkillPresetBase, isTemplate bool, skills []SkillPresetSkillResponse, assignedTo SkillPresetAssignedTo, ) *SkillPresetRevisionResponse`

NewSkillPresetRevisionResponse instantiates a new SkillPresetRevisionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetRevisionResponseWithDefaults

`func NewSkillPresetRevisionResponseWithDefaults() *SkillPresetRevisionResponse`

NewSkillPresetRevisionResponseWithDefaults instantiates a new SkillPresetRevisionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillPresetId

`func (o *SkillPresetRevisionResponse) GetSkillPresetId() string`

GetSkillPresetId returns the SkillPresetId field if non-nil, zero value otherwise.

### GetSkillPresetIdOk

`func (o *SkillPresetRevisionResponse) GetSkillPresetIdOk() (*string, bool)`

GetSkillPresetIdOk returns a tuple with the SkillPresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillPresetId

`func (o *SkillPresetRevisionResponse) SetSkillPresetId(v string)`

SetSkillPresetId sets SkillPresetId field to given value.


### GetRevision

`func (o *SkillPresetRevisionResponse) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SkillPresetRevisionResponse) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SkillPresetRevisionResponse) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetChangeType

`func (o *SkillPresetRevisionResponse) GetChangeType() SkillPresetChangeType`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *SkillPresetRevisionResponse) GetChangeTypeOk() (*SkillPresetChangeType, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *SkillPresetRevisionResponse) SetChangeType(v SkillPresetChangeType)`

SetChangeType sets ChangeType field to given value.


### GetEditedBy

`func (o *SkillPresetRevisionResponse) GetEditedBy() string`

GetEditedBy returns the EditedBy field if non-nil, zero value otherwise.

### GetEditedByOk

`func (o *SkillPresetRevisionResponse) GetEditedByOk() (*string, bool)`

GetEditedByOk returns a tuple with the EditedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedBy

`func (o *SkillPresetRevisionResponse) SetEditedBy(v string)`

SetEditedBy sets EditedBy field to given value.


### GetEditedAt

`func (o *SkillPresetRevisionResponse) GetEditedAt() time.Time`

GetEditedAt returns the EditedAt field if non-nil, zero value otherwise.

### GetEditedAtOk

`func (o *SkillPresetRevisionResponse) GetEditedAtOk() (*time.Time, bool)`

GetEditedAtOk returns a tuple with the EditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedAt

`func (o *SkillPresetRevisionResponse) SetEditedAt(v time.Time)`

SetEditedAt sets EditedAt field to given value.


### GetName

`func (o *SkillPresetRevisionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SkillPresetRevisionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SkillPresetRevisionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SkillPresetRevisionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SkillPresetRevisionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SkillPresetRevisionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetBase

`func (o *SkillPresetRevisionResponse) GetBase() SkillPresetBase`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *SkillPresetRevisionResponse) GetBaseOk() (*SkillPresetBase, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *SkillPresetRevisionResponse) SetBase(v SkillPresetBase)`

SetBase sets Base field to given value.


### GetIsTemplate

`func (o *SkillPresetRevisionResponse) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *SkillPresetRevisionResponse) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *SkillPresetRevisionResponse) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.


### GetTeamContext

`func (o *SkillPresetRevisionResponse) GetTeamContext() string`

GetTeamContext returns the TeamContext field if non-nil, zero value otherwise.

### GetTeamContextOk

`func (o *SkillPresetRevisionResponse) GetTeamContextOk() (*string, bool)`

GetTeamContextOk returns a tuple with the TeamContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamContext

`func (o *SkillPresetRevisionResponse) SetTeamContext(v string)`

SetTeamContext sets TeamContext field to given value.

### HasTeamContext

`func (o *SkillPresetRevisionResponse) HasTeamContext() bool`

HasTeamContext returns a boolean if a field has been set.

### SetTeamContextNil

`func (o *SkillPresetRevisionResponse) SetTeamContextNil(b bool)`

 SetTeamContextNil sets the value for TeamContext to be an explicit nil

### UnsetTeamContext
`func (o *SkillPresetRevisionResponse) UnsetTeamContext()`

UnsetTeamContext ensures that no value is present for TeamContext, not even an explicit nil
### GetSkills

`func (o *SkillPresetRevisionResponse) GetSkills() []SkillPresetSkillResponse`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *SkillPresetRevisionResponse) GetSkillsOk() (*[]SkillPresetSkillResponse, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *SkillPresetRevisionResponse) SetSkills(v []SkillPresetSkillResponse)`

SetSkills sets Skills field to given value.


### GetAssignedTo

`func (o *SkillPresetRevisionResponse) GetAssignedTo() SkillPresetAssignedTo`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *SkillPresetRevisionResponse) GetAssignedToOk() (*SkillPresetAssignedTo, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *SkillPresetRevisionResponse) SetAssignedTo(v SkillPresetAssignedTo)`

SetAssignedTo sets AssignedTo field to given value.


### GetSecurityReview

`func (o *SkillPresetRevisionResponse) GetSecurityReview() SkillPresetRevisionResponseSecurityReview`

GetSecurityReview returns the SecurityReview field if non-nil, zero value otherwise.

### GetSecurityReviewOk

`func (o *SkillPresetRevisionResponse) GetSecurityReviewOk() (*SkillPresetRevisionResponseSecurityReview, bool)`

GetSecurityReviewOk returns a tuple with the SecurityReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityReview

`func (o *SkillPresetRevisionResponse) SetSecurityReview(v SkillPresetRevisionResponseSecurityReview)`

SetSecurityReview sets SecurityReview field to given value.

### HasSecurityReview

`func (o *SkillPresetRevisionResponse) HasSecurityReview() bool`

HasSecurityReview returns a boolean if a field has been set.

### SetSecurityReviewNil

`func (o *SkillPresetRevisionResponse) SetSecurityReviewNil(b bool)`

 SetSecurityReviewNil sets the value for SecurityReview to be an explicit nil

### UnsetSecurityReview
`func (o *SkillPresetRevisionResponse) UnsetSecurityReview()`

UnsetSecurityReview ensures that no value is present for SecurityReview, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


