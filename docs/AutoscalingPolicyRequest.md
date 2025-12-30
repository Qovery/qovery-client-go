# AutoscalingPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**AutoscalingMode**](AutoscalingMode.md) |  | 
**PollingIntervalSeconds** | Pointer to **int32** |  | [optional] 
**CooldownPeriodSeconds** | Pointer to **int32** |  | [optional] 
**Scalers** | [**[]KedaScalerRequest**](KedaScalerRequest.md) |  | 

## Methods

### NewAutoscalingPolicyRequest

`func NewAutoscalingPolicyRequest(mode AutoscalingMode, scalers []KedaScalerRequest, ) *AutoscalingPolicyRequest`

NewAutoscalingPolicyRequest instantiates a new AutoscalingPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalingPolicyRequestWithDefaults

`func NewAutoscalingPolicyRequestWithDefaults() *AutoscalingPolicyRequest`

NewAutoscalingPolicyRequestWithDefaults instantiates a new AutoscalingPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *AutoscalingPolicyRequest) GetMode() AutoscalingMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AutoscalingPolicyRequest) GetModeOk() (*AutoscalingMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AutoscalingPolicyRequest) SetMode(v AutoscalingMode)`

SetMode sets Mode field to given value.


### GetPollingIntervalSeconds

`func (o *AutoscalingPolicyRequest) GetPollingIntervalSeconds() int32`

GetPollingIntervalSeconds returns the PollingIntervalSeconds field if non-nil, zero value otherwise.

### GetPollingIntervalSecondsOk

`func (o *AutoscalingPolicyRequest) GetPollingIntervalSecondsOk() (*int32, bool)`

GetPollingIntervalSecondsOk returns a tuple with the PollingIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingIntervalSeconds

`func (o *AutoscalingPolicyRequest) SetPollingIntervalSeconds(v int32)`

SetPollingIntervalSeconds sets PollingIntervalSeconds field to given value.

### HasPollingIntervalSeconds

`func (o *AutoscalingPolicyRequest) HasPollingIntervalSeconds() bool`

HasPollingIntervalSeconds returns a boolean if a field has been set.

### GetCooldownPeriodSeconds

`func (o *AutoscalingPolicyRequest) GetCooldownPeriodSeconds() int32`

GetCooldownPeriodSeconds returns the CooldownPeriodSeconds field if non-nil, zero value otherwise.

### GetCooldownPeriodSecondsOk

`func (o *AutoscalingPolicyRequest) GetCooldownPeriodSecondsOk() (*int32, bool)`

GetCooldownPeriodSecondsOk returns a tuple with the CooldownPeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownPeriodSeconds

`func (o *AutoscalingPolicyRequest) SetCooldownPeriodSeconds(v int32)`

SetCooldownPeriodSeconds sets CooldownPeriodSeconds field to given value.

### HasCooldownPeriodSeconds

`func (o *AutoscalingPolicyRequest) HasCooldownPeriodSeconds() bool`

HasCooldownPeriodSeconds returns a boolean if a field has been set.

### GetScalers

`func (o *AutoscalingPolicyRequest) GetScalers() []KedaScalerRequest`

GetScalers returns the Scalers field if non-nil, zero value otherwise.

### GetScalersOk

`func (o *AutoscalingPolicyRequest) GetScalersOk() (*[]KedaScalerRequest, bool)`

GetScalersOk returns a tuple with the Scalers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalers

`func (o *AutoscalingPolicyRequest) SetScalers(v []KedaScalerRequest)`

SetScalers sets Scalers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


