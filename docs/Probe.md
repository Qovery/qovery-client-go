# Probe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ProbeType**](ProbeType.md) |  | [optional] 
**InitialDelaySeconds** | Pointer to **int32** |  | [optional] [default to 30]
**PeriodSeconds** | Pointer to **int32** |  | [optional] [default to 10]
**TimeoutSeconds** | Pointer to **int32** |  | [optional] [default to 5]
**SuccessThreshold** | Pointer to **int32** |  | [optional] [default to 1]
**FailureThreshold** | Pointer to **int32** |  | [optional] [default to 9]

## Methods

### NewProbe

`func NewProbe() *Probe`

NewProbe instantiates a new Probe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProbeWithDefaults

`func NewProbeWithDefaults() *Probe`

NewProbeWithDefaults instantiates a new Probe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Probe) GetType() ProbeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Probe) GetTypeOk() (*ProbeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Probe) SetType(v ProbeType)`

SetType sets Type field to given value.

### HasType

`func (o *Probe) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInitialDelaySeconds

`func (o *Probe) GetInitialDelaySeconds() int32`

GetInitialDelaySeconds returns the InitialDelaySeconds field if non-nil, zero value otherwise.

### GetInitialDelaySecondsOk

`func (o *Probe) GetInitialDelaySecondsOk() (*int32, bool)`

GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDelaySeconds

`func (o *Probe) SetInitialDelaySeconds(v int32)`

SetInitialDelaySeconds sets InitialDelaySeconds field to given value.

### HasInitialDelaySeconds

`func (o *Probe) HasInitialDelaySeconds() bool`

HasInitialDelaySeconds returns a boolean if a field has been set.

### GetPeriodSeconds

`func (o *Probe) GetPeriodSeconds() int32`

GetPeriodSeconds returns the PeriodSeconds field if non-nil, zero value otherwise.

### GetPeriodSecondsOk

`func (o *Probe) GetPeriodSecondsOk() (*int32, bool)`

GetPeriodSecondsOk returns a tuple with the PeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSeconds

`func (o *Probe) SetPeriodSeconds(v int32)`

SetPeriodSeconds sets PeriodSeconds field to given value.

### HasPeriodSeconds

`func (o *Probe) HasPeriodSeconds() bool`

HasPeriodSeconds returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *Probe) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *Probe) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *Probe) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *Probe) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetSuccessThreshold

`func (o *Probe) GetSuccessThreshold() int32`

GetSuccessThreshold returns the SuccessThreshold field if non-nil, zero value otherwise.

### GetSuccessThresholdOk

`func (o *Probe) GetSuccessThresholdOk() (*int32, bool)`

GetSuccessThresholdOk returns a tuple with the SuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessThreshold

`func (o *Probe) SetSuccessThreshold(v int32)`

SetSuccessThreshold sets SuccessThreshold field to given value.

### HasSuccessThreshold

`func (o *Probe) HasSuccessThreshold() bool`

HasSuccessThreshold returns a boolean if a field has been set.

### GetFailureThreshold

`func (o *Probe) GetFailureThreshold() int32`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *Probe) GetFailureThresholdOk() (*int32, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *Probe) SetFailureThreshold(v int32)`

SetFailureThreshold sets FailureThreshold field to given value.

### HasFailureThreshold

`func (o *Probe) HasFailureThreshold() bool`

HasFailureThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


