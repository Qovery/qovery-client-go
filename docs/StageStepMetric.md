# StageStepMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepName** | Pointer to [**StageStepMetricNameEnum**](StageStepMetricNameEnum.md) |  | [optional] 
**Status** | Pointer to [**StepMetricStatusEnum**](StepMetricStatusEnum.md) |  | [optional] 
**DurationSec** | Pointer to **int32** | The duration of the step in seconds. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


