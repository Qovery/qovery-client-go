# HealthcheckStartupProbe

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

### NewHealthcheckStartupProbe

`func NewHealthcheckStartupProbe() *HealthcheckStartupProbe`

NewHealthcheckStartupProbe instantiates a new HealthcheckStartupProbe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthcheckStartupProbeWithDefaults

`func NewHealthcheckStartupProbeWithDefaults() *HealthcheckStartupProbe`

NewHealthcheckStartupProbeWithDefaults instantiates a new HealthcheckStartupProbe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HealthcheckStartupProbe) GetType() ProbeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthcheckStartupProbe) GetTypeOk() (*ProbeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthcheckStartupProbe) SetType(v ProbeType)`

SetType sets Type field to given value.

### HasType

`func (o *HealthcheckStartupProbe) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInitialDelaySeconds

`func (o *HealthcheckStartupProbe) GetInitialDelaySeconds() int32`

GetInitialDelaySeconds returns the InitialDelaySeconds field if non-nil, zero value otherwise.

### GetInitialDelaySecondsOk

`func (o *HealthcheckStartupProbe) GetInitialDelaySecondsOk() (*int32, bool)`

GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDelaySeconds

`func (o *HealthcheckStartupProbe) SetInitialDelaySeconds(v int32)`

SetInitialDelaySeconds sets InitialDelaySeconds field to given value.

### HasInitialDelaySeconds

`func (o *HealthcheckStartupProbe) HasInitialDelaySeconds() bool`

HasInitialDelaySeconds returns a boolean if a field has been set.

### GetPeriodSeconds

`func (o *HealthcheckStartupProbe) GetPeriodSeconds() int32`

GetPeriodSeconds returns the PeriodSeconds field if non-nil, zero value otherwise.

### GetPeriodSecondsOk

`func (o *HealthcheckStartupProbe) GetPeriodSecondsOk() (*int32, bool)`

GetPeriodSecondsOk returns a tuple with the PeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSeconds

`func (o *HealthcheckStartupProbe) SetPeriodSeconds(v int32)`

SetPeriodSeconds sets PeriodSeconds field to given value.

### HasPeriodSeconds

`func (o *HealthcheckStartupProbe) HasPeriodSeconds() bool`

HasPeriodSeconds returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *HealthcheckStartupProbe) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *HealthcheckStartupProbe) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *HealthcheckStartupProbe) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *HealthcheckStartupProbe) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetSuccessThreshold

`func (o *HealthcheckStartupProbe) GetSuccessThreshold() int32`

GetSuccessThreshold returns the SuccessThreshold field if non-nil, zero value otherwise.

### GetSuccessThresholdOk

`func (o *HealthcheckStartupProbe) GetSuccessThresholdOk() (*int32, bool)`

GetSuccessThresholdOk returns a tuple with the SuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessThreshold

`func (o *HealthcheckStartupProbe) SetSuccessThreshold(v int32)`

SetSuccessThreshold sets SuccessThreshold field to given value.

### HasSuccessThreshold

`func (o *HealthcheckStartupProbe) HasSuccessThreshold() bool`

HasSuccessThreshold returns a boolean if a field has been set.

### GetFailureThreshold

`func (o *HealthcheckStartupProbe) GetFailureThreshold() int32`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *HealthcheckStartupProbe) GetFailureThresholdOk() (*int32, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *HealthcheckStartupProbe) SetFailureThreshold(v int32)`

SetFailureThreshold sets FailureThreshold field to given value.

### HasFailureThreshold

`func (o *HealthcheckStartupProbe) HasFailureThreshold() bool`

HasFailureThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


