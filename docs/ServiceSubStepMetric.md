# ServiceSubStepMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | PROVISION_BUILDER, IMAGE_PUSH or CACHE_EXPORT. A string, not an enum, so a new sub-step never breaks a client: accept values not listed here.  | 
**Status** | [**StepMetricStatusEnum**](StepMetricStatusEnum.md) |  | 
**DurationSec** | **int32** | The duration of the sub-step in seconds. | 
**StartedAt** | Pointer to **NullableTime** | The time at which the sub-step started. | [optional] 

## Methods

### NewServiceSubStepMetric

`func NewServiceSubStepMetric(name string, status StepMetricStatusEnum, durationSec int32, ) *ServiceSubStepMetric`

NewServiceSubStepMetric instantiates a new ServiceSubStepMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSubStepMetricWithDefaults

`func NewServiceSubStepMetricWithDefaults() *ServiceSubStepMetric`

NewServiceSubStepMetricWithDefaults instantiates a new ServiceSubStepMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceSubStepMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceSubStepMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceSubStepMetric) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ServiceSubStepMetric) GetStatus() StepMetricStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceSubStepMetric) GetStatusOk() (*StepMetricStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceSubStepMetric) SetStatus(v StepMetricStatusEnum)`

SetStatus sets Status field to given value.


### GetDurationSec

`func (o *ServiceSubStepMetric) GetDurationSec() int32`

GetDurationSec returns the DurationSec field if non-nil, zero value otherwise.

### GetDurationSecOk

`func (o *ServiceSubStepMetric) GetDurationSecOk() (*int32, bool)`

GetDurationSecOk returns a tuple with the DurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSec

`func (o *ServiceSubStepMetric) SetDurationSec(v int32)`

SetDurationSec sets DurationSec field to given value.


### GetStartedAt

`func (o *ServiceSubStepMetric) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ServiceSubStepMetric) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ServiceSubStepMetric) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ServiceSubStepMetric) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *ServiceSubStepMetric) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *ServiceSubStepMetric) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


