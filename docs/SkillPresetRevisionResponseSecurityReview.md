# SkillPresetRevisionResponseSecurityReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verdict** | [**SkillPresetSecurityReviewVerdict**](SkillPresetSecurityReviewVerdict.md) |  | 
**Model** | **string** | The model that reviewed the content. | 
**PromptVersion** | **string** | Version of the review instructions the model used. | 
**Findings** | [**[]SkillPresetRejectionFinding**](SkillPresetRejectionFinding.md) | Notes the model raised without blocking. Usually empty. | 

## Methods

### NewSkillPresetRevisionResponseSecurityReview

`func NewSkillPresetRevisionResponseSecurityReview(verdict SkillPresetSecurityReviewVerdict, model string, promptVersion string, findings []SkillPresetRejectionFinding, ) *SkillPresetRevisionResponseSecurityReview`

NewSkillPresetRevisionResponseSecurityReview instantiates a new SkillPresetRevisionResponseSecurityReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetRevisionResponseSecurityReviewWithDefaults

`func NewSkillPresetRevisionResponseSecurityReviewWithDefaults() *SkillPresetRevisionResponseSecurityReview`

NewSkillPresetRevisionResponseSecurityReviewWithDefaults instantiates a new SkillPresetRevisionResponseSecurityReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerdict

`func (o *SkillPresetRevisionResponseSecurityReview) GetVerdict() SkillPresetSecurityReviewVerdict`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *SkillPresetRevisionResponseSecurityReview) GetVerdictOk() (*SkillPresetSecurityReviewVerdict, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *SkillPresetRevisionResponseSecurityReview) SetVerdict(v SkillPresetSecurityReviewVerdict)`

SetVerdict sets Verdict field to given value.


### GetModel

`func (o *SkillPresetRevisionResponseSecurityReview) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SkillPresetRevisionResponseSecurityReview) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SkillPresetRevisionResponseSecurityReview) SetModel(v string)`

SetModel sets Model field to given value.


### GetPromptVersion

`func (o *SkillPresetRevisionResponseSecurityReview) GetPromptVersion() string`

GetPromptVersion returns the PromptVersion field if non-nil, zero value otherwise.

### GetPromptVersionOk

`func (o *SkillPresetRevisionResponseSecurityReview) GetPromptVersionOk() (*string, bool)`

GetPromptVersionOk returns a tuple with the PromptVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptVersion

`func (o *SkillPresetRevisionResponseSecurityReview) SetPromptVersion(v string)`

SetPromptVersion sets PromptVersion field to given value.


### GetFindings

`func (o *SkillPresetRevisionResponseSecurityReview) GetFindings() []SkillPresetRejectionFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *SkillPresetRevisionResponseSecurityReview) GetFindingsOk() (*[]SkillPresetRejectionFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *SkillPresetRevisionResponseSecurityReview) SetFindings(v []SkillPresetRejectionFinding)`

SetFindings sets Findings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


