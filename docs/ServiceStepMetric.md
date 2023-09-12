# ServiceStepMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepName** | Pointer to [**ServiceStepMetricNameEnum**](ServiceStepMetricNameEnum.md) |  | [optional] 
**Status** | Pointer to [**StepMetricStatusEnum**](StepMetricStatusEnum.md) |  | [optional] 
**DurationSec** | Pointer to **int32** | The duration of the step in seconds. | [optional] 

## Methods

### NewServiceStepMetric

`func NewServiceStepMetric() *ServiceStepMetric`

NewServiceStepMetric instantiates a new ServiceStepMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStepMetricWithDefaults

`func NewServiceStepMetricWithDefaults() *ServiceStepMetric`

NewServiceStepMetricWithDefaults instantiates a new ServiceStepMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepName

`func (o *ServiceStepMetric) GetStepName() ServiceStepMetricNameEnum`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *ServiceStepMetric) GetStepNameOk() (*ServiceStepMetricNameEnum, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *ServiceStepMetric) SetStepName(v ServiceStepMetricNameEnum)`

SetStepName sets StepName field to given value.

### HasStepName

`func (o *ServiceStepMetric) HasStepName() bool`

HasStepName returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceStepMetric) GetStatus() StepMetricStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceStepMetric) GetStatusOk() (*StepMetricStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceStepMetric) SetStatus(v StepMetricStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceStepMetric) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDurationSec

`func (o *ServiceStepMetric) GetDurationSec() int32`

GetDurationSec returns the DurationSec field if non-nil, zero value otherwise.

### GetDurationSecOk

`func (o *ServiceStepMetric) GetDurationSecOk() (*int32, bool)`

GetDurationSecOk returns a tuple with the DurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSec

`func (o *ServiceStepMetric) SetDurationSec(v int32)`

SetDurationSec sets DurationSec field to given value.

### HasDurationSec

`func (o *ServiceStepMetric) HasDurationSec() bool`

HasDurationSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


