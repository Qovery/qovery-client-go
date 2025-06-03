# StageStepMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalDurationSec** | Pointer to **int32** | The total duration in seconds of the stage deployment or null if the deployment is not completed | [optional] 
**Details** | Pointer to [**[]StageStepMetric**](StageStepMetric.md) | A list of metrics for deployment steps of the stage. | [optional] 

## Methods

### NewStageStepMetrics

`func NewStageStepMetrics() *StageStepMetrics`

NewStageStepMetrics instantiates a new StageStepMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageStepMetricsWithDefaults

`func NewStageStepMetricsWithDefaults() *StageStepMetrics`

NewStageStepMetricsWithDefaults instantiates a new StageStepMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalDurationSec

`func (o *StageStepMetrics) GetTotalDurationSec() int32`

GetTotalDurationSec returns the TotalDurationSec field if non-nil, zero value otherwise.

### GetTotalDurationSecOk

`func (o *StageStepMetrics) GetTotalDurationSecOk() (*int32, bool)`

GetTotalDurationSecOk returns a tuple with the TotalDurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationSec

`func (o *StageStepMetrics) SetTotalDurationSec(v int32)`

SetTotalDurationSec sets TotalDurationSec field to given value.

### HasTotalDurationSec

`func (o *StageStepMetrics) HasTotalDurationSec() bool`

HasTotalDurationSec returns a boolean if a field has been set.

### GetDetails

`func (o *StageStepMetrics) GetDetails() []StageStepMetric`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *StageStepMetrics) GetDetailsOk() (*[]StageStepMetric, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *StageStepMetrics) SetDetails(v []StageStepMetric)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *StageStepMetrics) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


