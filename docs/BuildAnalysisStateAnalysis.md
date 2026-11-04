# BuildAnalysisStateAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** |  | 
**ExecutionId** | **string** | Execution id of the analysed build. | 
**Status** | [**BuildAnalysisStatus**](BuildAnalysisStatus.md) |  | 
**Stale** | **bool** | True when the latest build of the service is not the analysed one. The analysis stays readable, but only the latest build can be analysed. | 
**DockerfileSource** | [**BuildAnalysisDockerfileSource**](BuildAnalysisDockerfileSource.md) |  | 
**FailureCode** | [**NullableBuildAnalysisFailureCode**](BuildAnalysisFailureCode.md) | Why the analysis failed, when &#x60;status&#x60; is FAILED. | 
**FailureReason** | **NullableString** | Human-readable detail of the failure, when &#x60;status&#x60; is FAILED. | 
**Result** | [**NullableBuildAnalysisResponseResult**](BuildAnalysisResponseResult.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewBuildAnalysisStateAnalysis

`func NewBuildAnalysisStateAnalysis(serviceId string, executionId string, status BuildAnalysisStatus, stale bool, dockerfileSource BuildAnalysisDockerfileSource, failureCode NullableBuildAnalysisFailureCode, failureReason NullableString, result NullableBuildAnalysisResponseResult, createdAt time.Time, updatedAt time.Time, ) *BuildAnalysisStateAnalysis`

NewBuildAnalysisStateAnalysis instantiates a new BuildAnalysisStateAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildAnalysisStateAnalysisWithDefaults

`func NewBuildAnalysisStateAnalysisWithDefaults() *BuildAnalysisStateAnalysis`

NewBuildAnalysisStateAnalysisWithDefaults instantiates a new BuildAnalysisStateAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *BuildAnalysisStateAnalysis) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *BuildAnalysisStateAnalysis) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *BuildAnalysisStateAnalysis) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetExecutionId

`func (o *BuildAnalysisStateAnalysis) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BuildAnalysisStateAnalysis) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BuildAnalysisStateAnalysis) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetStatus

`func (o *BuildAnalysisStateAnalysis) GetStatus() BuildAnalysisStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BuildAnalysisStateAnalysis) GetStatusOk() (*BuildAnalysisStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BuildAnalysisStateAnalysis) SetStatus(v BuildAnalysisStatus)`

SetStatus sets Status field to given value.


### GetStale

`func (o *BuildAnalysisStateAnalysis) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *BuildAnalysisStateAnalysis) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *BuildAnalysisStateAnalysis) SetStale(v bool)`

SetStale sets Stale field to given value.


### GetDockerfileSource

`func (o *BuildAnalysisStateAnalysis) GetDockerfileSource() BuildAnalysisDockerfileSource`

GetDockerfileSource returns the DockerfileSource field if non-nil, zero value otherwise.

### GetDockerfileSourceOk

`func (o *BuildAnalysisStateAnalysis) GetDockerfileSourceOk() (*BuildAnalysisDockerfileSource, bool)`

GetDockerfileSourceOk returns a tuple with the DockerfileSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfileSource

`func (o *BuildAnalysisStateAnalysis) SetDockerfileSource(v BuildAnalysisDockerfileSource)`

SetDockerfileSource sets DockerfileSource field to given value.


### GetFailureCode

`func (o *BuildAnalysisStateAnalysis) GetFailureCode() BuildAnalysisFailureCode`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *BuildAnalysisStateAnalysis) GetFailureCodeOk() (*BuildAnalysisFailureCode, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *BuildAnalysisStateAnalysis) SetFailureCode(v BuildAnalysisFailureCode)`

SetFailureCode sets FailureCode field to given value.


### SetFailureCodeNil

`func (o *BuildAnalysisStateAnalysis) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *BuildAnalysisStateAnalysis) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetFailureReason

`func (o *BuildAnalysisStateAnalysis) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *BuildAnalysisStateAnalysis) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *BuildAnalysisStateAnalysis) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.


### SetFailureReasonNil

`func (o *BuildAnalysisStateAnalysis) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *BuildAnalysisStateAnalysis) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetResult

`func (o *BuildAnalysisStateAnalysis) GetResult() BuildAnalysisResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BuildAnalysisStateAnalysis) GetResultOk() (*BuildAnalysisResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BuildAnalysisStateAnalysis) SetResult(v BuildAnalysisResponseResult)`

SetResult sets Result field to given value.


### SetResultNil

`func (o *BuildAnalysisStateAnalysis) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *BuildAnalysisStateAnalysis) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetCreatedAt

`func (o *BuildAnalysisStateAnalysis) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BuildAnalysisStateAnalysis) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BuildAnalysisStateAnalysis) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BuildAnalysisStateAnalysis) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BuildAnalysisStateAnalysis) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BuildAnalysisStateAnalysis) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


