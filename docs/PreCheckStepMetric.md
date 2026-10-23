# PreCheckStepMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | CLOUD_CREDENTIALS_CHECK (once per deployment) or GIT_CHECK (once per service built or fetched from Git). A string, not an enum, so a new check never breaks a client: accept values not listed here. Log lines of the check carry the same value in &#x60;step_name&#x60;.  | 
**Status** | [**StepMetricStatusEnum**](StepMetricStatusEnum.md) |  | 
**DurationSec** | **int32** | The duration of the check in seconds, rounded down (0 means under a second). | 
**StartedAt** | Pointer to **NullableTime** | The time at which the check started. | [optional] 
**ServiceId** | Pointer to **string** | The service whose Git configuration was checked. Only set on GIT_CHECK. | [optional] 

## Methods

### NewPreCheckStepMetric

`func NewPreCheckStepMetric(name string, status StepMetricStatusEnum, durationSec int32, ) *PreCheckStepMetric`

NewPreCheckStepMetric instantiates a new PreCheckStepMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreCheckStepMetricWithDefaults

`func NewPreCheckStepMetricWithDefaults() *PreCheckStepMetric`

NewPreCheckStepMetricWithDefaults instantiates a new PreCheckStepMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PreCheckStepMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PreCheckStepMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PreCheckStepMetric) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *PreCheckStepMetric) GetStatus() StepMetricStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PreCheckStepMetric) GetStatusOk() (*StepMetricStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PreCheckStepMetric) SetStatus(v StepMetricStatusEnum)`

SetStatus sets Status field to given value.


### GetDurationSec

`func (o *PreCheckStepMetric) GetDurationSec() int32`

GetDurationSec returns the DurationSec field if non-nil, zero value otherwise.

### GetDurationSecOk

`func (o *PreCheckStepMetric) GetDurationSecOk() (*int32, bool)`

GetDurationSecOk returns a tuple with the DurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSec

`func (o *PreCheckStepMetric) SetDurationSec(v int32)`

SetDurationSec sets DurationSec field to given value.


### GetStartedAt

`func (o *PreCheckStepMetric) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *PreCheckStepMetric) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *PreCheckStepMetric) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *PreCheckStepMetric) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *PreCheckStepMetric) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *PreCheckStepMetric) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetServiceId

`func (o *PreCheckStepMetric) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PreCheckStepMetric) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PreCheckStepMetric) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *PreCheckStepMetric) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


