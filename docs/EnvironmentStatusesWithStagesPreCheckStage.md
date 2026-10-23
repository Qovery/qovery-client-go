# EnvironmentStatusesWithStagesPreCheckStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**StepMetricStatusEnum**](StepMetricStatusEnum.md) |  | [optional] 
**TotalDurationSec** | Pointer to **int32** |  | [optional] 
**Steps** | Pointer to [**[]PreCheckStepMetric**](PreCheckStepMetric.md) | One entry per check the pre-check ran, in order. Their time is already part of &#x60;total_duration_sec&#x60;: never add it to a total.  | [optional] 

## Methods

### NewEnvironmentStatusesWithStagesPreCheckStage

`func NewEnvironmentStatusesWithStagesPreCheckStage() *EnvironmentStatusesWithStagesPreCheckStage`

NewEnvironmentStatusesWithStagesPreCheckStage instantiates a new EnvironmentStatusesWithStagesPreCheckStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentStatusesWithStagesPreCheckStageWithDefaults

`func NewEnvironmentStatusesWithStagesPreCheckStageWithDefaults() *EnvironmentStatusesWithStagesPreCheckStage`

NewEnvironmentStatusesWithStagesPreCheckStageWithDefaults instantiates a new EnvironmentStatusesWithStagesPreCheckStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EnvironmentStatusesWithStagesPreCheckStage) GetStatus() StepMetricStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentStatusesWithStagesPreCheckStage) GetStatusOk() (*StepMetricStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentStatusesWithStagesPreCheckStage) SetStatus(v StepMetricStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnvironmentStatusesWithStagesPreCheckStage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalDurationSec

`func (o *EnvironmentStatusesWithStagesPreCheckStage) GetTotalDurationSec() int32`

GetTotalDurationSec returns the TotalDurationSec field if non-nil, zero value otherwise.

### GetTotalDurationSecOk

`func (o *EnvironmentStatusesWithStagesPreCheckStage) GetTotalDurationSecOk() (*int32, bool)`

GetTotalDurationSecOk returns a tuple with the TotalDurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationSec

`func (o *EnvironmentStatusesWithStagesPreCheckStage) SetTotalDurationSec(v int32)`

SetTotalDurationSec sets TotalDurationSec field to given value.

### HasTotalDurationSec

`func (o *EnvironmentStatusesWithStagesPreCheckStage) HasTotalDurationSec() bool`

HasTotalDurationSec returns a boolean if a field has been set.

### GetSteps

`func (o *EnvironmentStatusesWithStagesPreCheckStage) GetSteps() []PreCheckStepMetric`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *EnvironmentStatusesWithStagesPreCheckStage) GetStepsOk() (*[]PreCheckStepMetric, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *EnvironmentStatusesWithStagesPreCheckStage) SetSteps(v []PreCheckStepMetric)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *EnvironmentStatusesWithStagesPreCheckStage) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


