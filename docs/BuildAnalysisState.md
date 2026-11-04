# BuildAnalysisState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatestBuildExecutionId** | **NullableString** | Execution id of the latest build of the service, the only build that can be analysed. Null when no deployment in the kept history ran a build. | 
**Analysis** | [**NullableBuildAnalysisStateAnalysis**](BuildAnalysisStateAnalysis.md) |  | 

## Methods

### NewBuildAnalysisState

`func NewBuildAnalysisState(latestBuildExecutionId NullableString, analysis NullableBuildAnalysisStateAnalysis, ) *BuildAnalysisState`

NewBuildAnalysisState instantiates a new BuildAnalysisState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildAnalysisStateWithDefaults

`func NewBuildAnalysisStateWithDefaults() *BuildAnalysisState`

NewBuildAnalysisStateWithDefaults instantiates a new BuildAnalysisState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatestBuildExecutionId

`func (o *BuildAnalysisState) GetLatestBuildExecutionId() string`

GetLatestBuildExecutionId returns the LatestBuildExecutionId field if non-nil, zero value otherwise.

### GetLatestBuildExecutionIdOk

`func (o *BuildAnalysisState) GetLatestBuildExecutionIdOk() (*string, bool)`

GetLatestBuildExecutionIdOk returns a tuple with the LatestBuildExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestBuildExecutionId

`func (o *BuildAnalysisState) SetLatestBuildExecutionId(v string)`

SetLatestBuildExecutionId sets LatestBuildExecutionId field to given value.


### SetLatestBuildExecutionIdNil

`func (o *BuildAnalysisState) SetLatestBuildExecutionIdNil(b bool)`

 SetLatestBuildExecutionIdNil sets the value for LatestBuildExecutionId to be an explicit nil

### UnsetLatestBuildExecutionId
`func (o *BuildAnalysisState) UnsetLatestBuildExecutionId()`

UnsetLatestBuildExecutionId ensures that no value is present for LatestBuildExecutionId, not even an explicit nil
### GetAnalysis

`func (o *BuildAnalysisState) GetAnalysis() BuildAnalysisStateAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *BuildAnalysisState) GetAnalysisOk() (*BuildAnalysisStateAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *BuildAnalysisState) SetAnalysis(v BuildAnalysisStateAnalysis)`

SetAnalysis sets Analysis field to given value.


### SetAnalysisNil

`func (o *BuildAnalysisState) SetAnalysisNil(b bool)`

 SetAnalysisNil sets the value for Analysis to be an explicit nil

### UnsetAnalysis
`func (o *BuildAnalysisState) UnsetAnalysis()`

UnsetAnalysis ensures that no value is present for Analysis, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


