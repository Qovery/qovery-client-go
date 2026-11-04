# BuildAnalysisResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confidence** | [**BuildAnalysisConfidence**](BuildAnalysisConfidence.md) |  | 
**EstimatedTimeSaved** | [**NullableBuildAnalysisResultEstimatedTimeSaved**](BuildAnalysisResultEstimatedTimeSaved.md) |  | 
**AssetsReviewed** | [**[]BuildAnalysisAsset**](BuildAnalysisAsset.md) |  | 
**Summary** | **string** | What slows the build and why, with the evidence from the build. | 
**Improvements** | **[]string** | Changes to make, one sentence each, ordered by impact on build time, highest first. Empty when the build is already efficient. | 
**Dockerfile** | **NullableString** | Reworked Dockerfile to copy. Null when no improvement is a Dockerfile change, for example a variable value. Values hidden from the analysis appear as &#x60;&lt;redacted&gt;&#x60; and must be put back before use. | 
**DockerfileRationale** | **NullableString** | Why the reworked Dockerfile is faster. | 
**AgentPrompt** | **string** | Prompt the customer pastes into their own AI agent to apply the improvements in their repository. | 

## Methods

### NewBuildAnalysisResponseResult

`func NewBuildAnalysisResponseResult(confidence BuildAnalysisConfidence, estimatedTimeSaved NullableBuildAnalysisResultEstimatedTimeSaved, assetsReviewed []BuildAnalysisAsset, summary string, improvements []string, dockerfile NullableString, dockerfileRationale NullableString, agentPrompt string, ) *BuildAnalysisResponseResult`

NewBuildAnalysisResponseResult instantiates a new BuildAnalysisResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildAnalysisResponseResultWithDefaults

`func NewBuildAnalysisResponseResultWithDefaults() *BuildAnalysisResponseResult`

NewBuildAnalysisResponseResultWithDefaults instantiates a new BuildAnalysisResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidence

`func (o *BuildAnalysisResponseResult) GetConfidence() BuildAnalysisConfidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *BuildAnalysisResponseResult) GetConfidenceOk() (*BuildAnalysisConfidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *BuildAnalysisResponseResult) SetConfidence(v BuildAnalysisConfidence)`

SetConfidence sets Confidence field to given value.


### GetEstimatedTimeSaved

`func (o *BuildAnalysisResponseResult) GetEstimatedTimeSaved() BuildAnalysisResultEstimatedTimeSaved`

GetEstimatedTimeSaved returns the EstimatedTimeSaved field if non-nil, zero value otherwise.

### GetEstimatedTimeSavedOk

`func (o *BuildAnalysisResponseResult) GetEstimatedTimeSavedOk() (*BuildAnalysisResultEstimatedTimeSaved, bool)`

GetEstimatedTimeSavedOk returns a tuple with the EstimatedTimeSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeSaved

`func (o *BuildAnalysisResponseResult) SetEstimatedTimeSaved(v BuildAnalysisResultEstimatedTimeSaved)`

SetEstimatedTimeSaved sets EstimatedTimeSaved field to given value.


### SetEstimatedTimeSavedNil

`func (o *BuildAnalysisResponseResult) SetEstimatedTimeSavedNil(b bool)`

 SetEstimatedTimeSavedNil sets the value for EstimatedTimeSaved to be an explicit nil

### UnsetEstimatedTimeSaved
`func (o *BuildAnalysisResponseResult) UnsetEstimatedTimeSaved()`

UnsetEstimatedTimeSaved ensures that no value is present for EstimatedTimeSaved, not even an explicit nil
### GetAssetsReviewed

`func (o *BuildAnalysisResponseResult) GetAssetsReviewed() []BuildAnalysisAsset`

GetAssetsReviewed returns the AssetsReviewed field if non-nil, zero value otherwise.

### GetAssetsReviewedOk

`func (o *BuildAnalysisResponseResult) GetAssetsReviewedOk() (*[]BuildAnalysisAsset, bool)`

GetAssetsReviewedOk returns a tuple with the AssetsReviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetsReviewed

`func (o *BuildAnalysisResponseResult) SetAssetsReviewed(v []BuildAnalysisAsset)`

SetAssetsReviewed sets AssetsReviewed field to given value.


### GetSummary

`func (o *BuildAnalysisResponseResult) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *BuildAnalysisResponseResult) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *BuildAnalysisResponseResult) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetImprovements

`func (o *BuildAnalysisResponseResult) GetImprovements() []string`

GetImprovements returns the Improvements field if non-nil, zero value otherwise.

### GetImprovementsOk

`func (o *BuildAnalysisResponseResult) GetImprovementsOk() (*[]string, bool)`

GetImprovementsOk returns a tuple with the Improvements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImprovements

`func (o *BuildAnalysisResponseResult) SetImprovements(v []string)`

SetImprovements sets Improvements field to given value.


### GetDockerfile

`func (o *BuildAnalysisResponseResult) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *BuildAnalysisResponseResult) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *BuildAnalysisResponseResult) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.


### SetDockerfileNil

`func (o *BuildAnalysisResponseResult) SetDockerfileNil(b bool)`

 SetDockerfileNil sets the value for Dockerfile to be an explicit nil

### UnsetDockerfile
`func (o *BuildAnalysisResponseResult) UnsetDockerfile()`

UnsetDockerfile ensures that no value is present for Dockerfile, not even an explicit nil
### GetDockerfileRationale

`func (o *BuildAnalysisResponseResult) GetDockerfileRationale() string`

GetDockerfileRationale returns the DockerfileRationale field if non-nil, zero value otherwise.

### GetDockerfileRationaleOk

`func (o *BuildAnalysisResponseResult) GetDockerfileRationaleOk() (*string, bool)`

GetDockerfileRationaleOk returns a tuple with the DockerfileRationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfileRationale

`func (o *BuildAnalysisResponseResult) SetDockerfileRationale(v string)`

SetDockerfileRationale sets DockerfileRationale field to given value.


### SetDockerfileRationaleNil

`func (o *BuildAnalysisResponseResult) SetDockerfileRationaleNil(b bool)`

 SetDockerfileRationaleNil sets the value for DockerfileRationale to be an explicit nil

### UnsetDockerfileRationale
`func (o *BuildAnalysisResponseResult) UnsetDockerfileRationale()`

UnsetDockerfileRationale ensures that no value is present for DockerfileRationale, not even an explicit nil
### GetAgentPrompt

`func (o *BuildAnalysisResponseResult) GetAgentPrompt() string`

GetAgentPrompt returns the AgentPrompt field if non-nil, zero value otherwise.

### GetAgentPromptOk

`func (o *BuildAnalysisResponseResult) GetAgentPromptOk() (*string, bool)`

GetAgentPromptOk returns a tuple with the AgentPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPrompt

`func (o *BuildAnalysisResponseResult) SetAgentPrompt(v string)`

SetAgentPrompt sets AgentPrompt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


