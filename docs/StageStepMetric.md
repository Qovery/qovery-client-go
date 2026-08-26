# StageStepMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StageId** | Pointer to **string** |  | [optional] 
**StepName** | Pointer to [**StageStepMetricNameEnum**](StageStepMetricNameEnum.md) |  | [optional] 
**Status** | Pointer to [**StepMetricStatusEnum**](StepMetricStatusEnum.md) |  | [optional] 
**DurationSec** | Pointer to **int32** | The duration of the step in seconds. | [optional] 
**StartedAt** | Pointer to **NullableTime** | The time at which the step started. Present while the step is ongoing and may be retained after completion. | [optional] 

## Methods

### NewStageStepMetric

`func NewStageStepMetric() *StageStepMetric`

NewStageStepMetric instantiates a new StageStepMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageStepMetricWithDefaults

`func NewStageStepMetricWithDefaults() *StageStepMetric`

NewStageStepMetricWithDefaults instantiates a new StageStepMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStageId

`func (o *StageStepMetric) GetStageId() string`

GetStageId returns the StageId field if non-nil, zero value otherwise.

### GetStageIdOk

`func (o *StageStepMetric) GetStageIdOk() (*string, bool)`

GetStageIdOk returns a tuple with the StageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageId

`func (o *StageStepMetric) SetStageId(v string)`

SetStageId sets StageId field to given value.

### HasStageId

`func (o *StageStepMetric) HasStageId() bool`

HasStageId returns a boolean if a field has been set.

### GetStepName

`func (o *StageStepMetric) GetStepName() StageStepMetricNameEnum`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *StageStepMetric) GetStepNameOk() (*StageStepMetricNameEnum, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *StageStepMetric) SetStepName(v StageStepMetricNameEnum)`

SetStepName sets StepName field to given value.

### HasStepName

`func (o *StageStepMetric) HasStepName() bool`

HasStepName returns a boolean if a field has been set.

### GetStatus

`func (o *StageStepMetric) GetStatus() StepMetricStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StageStepMetric) GetStatusOk() (*StepMetricStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StageStepMetric) SetStatus(v StepMetricStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StageStepMetric) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDurationSec

`func (o *StageStepMetric) GetDurationSec() int32`

GetDurationSec returns the DurationSec field if non-nil, zero value otherwise.

### GetDurationSecOk

`func (o *StageStepMetric) GetDurationSecOk() (*int32, bool)`

GetDurationSecOk returns a tuple with the DurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSec

`func (o *StageStepMetric) SetDurationSec(v int32)`

SetDurationSec sets DurationSec field to given value.

### HasDurationSec

`func (o *StageStepMetric) HasDurationSec() bool`

HasDurationSec returns a boolean if a field has been set.

### GetStartedAt

`func (o *StageStepMetric) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *StageStepMetric) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *StageStepMetric) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *StageStepMetric) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *StageStepMetric) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *StageStepMetric) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


