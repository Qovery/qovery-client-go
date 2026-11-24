# SkillPresetSecurityReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verdict** | [**SkillPresetSecurityReviewVerdict**](SkillPresetSecurityReviewVerdict.md) |  | 
**Model** | **string** | The model that reviewed the content. | 
**PromptVersion** | **string** | Version of the review instructions the model used. | 
**Findings** | [**[]SkillPresetRejectionFinding**](SkillPresetRejectionFinding.md) | Notes the model raised without blocking. Usually empty. | 

## Methods

### NewSkillPresetSecurityReview

`func NewSkillPresetSecurityReview(verdict SkillPresetSecurityReviewVerdict, model string, promptVersion string, findings []SkillPresetRejectionFinding, ) *SkillPresetSecurityReview`

NewSkillPresetSecurityReview instantiates a new SkillPresetSecurityReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetSecurityReviewWithDefaults

`func NewSkillPresetSecurityReviewWithDefaults() *SkillPresetSecurityReview`

NewSkillPresetSecurityReviewWithDefaults instantiates a new SkillPresetSecurityReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerdict

`func (o *SkillPresetSecurityReview) GetVerdict() SkillPresetSecurityReviewVerdict`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *SkillPresetSecurityReview) GetVerdictOk() (*SkillPresetSecurityReviewVerdict, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *SkillPresetSecurityReview) SetVerdict(v SkillPresetSecurityReviewVerdict)`

SetVerdict sets Verdict field to given value.


### GetModel

`func (o *SkillPresetSecurityReview) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SkillPresetSecurityReview) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SkillPresetSecurityReview) SetModel(v string)`

SetModel sets Model field to given value.


### GetPromptVersion

`func (o *SkillPresetSecurityReview) GetPromptVersion() string`

GetPromptVersion returns the PromptVersion field if non-nil, zero value otherwise.

### GetPromptVersionOk

`func (o *SkillPresetSecurityReview) GetPromptVersionOk() (*string, bool)`

GetPromptVersionOk returns a tuple with the PromptVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptVersion

`func (o *SkillPresetSecurityReview) SetPromptVersion(v string)`

SetPromptVersion sets PromptVersion field to given value.


### GetFindings

`func (o *SkillPresetSecurityReview) GetFindings() []SkillPresetRejectionFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *SkillPresetSecurityReview) GetFindingsOk() (*[]SkillPresetRejectionFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *SkillPresetSecurityReview) SetFindings(v []SkillPresetRejectionFinding)`

SetFindings sets Findings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


