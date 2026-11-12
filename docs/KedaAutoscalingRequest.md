# KedaAutoscalingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**AutoscalingMode**](AutoscalingMode.md) |  | 
**PollingIntervalSeconds** | Pointer to **int32** |  | [optional] 
**CooldownPeriodSeconds** | Pointer to **int32** |  | [optional] 
**InitialCooldownPeriodSeconds** | Pointer to **NullableInt32** | KEDA ScaledObject &#x60;initialCooldownPeriod&#x60;: seconds to wait after the ScaledObject is created before the service can scale to zero. Omitted or null leaves KEDA&#39;s default (0). | [optional] 
**IdleReplicaCount** | Pointer to **NullableInt32** | KEDA ScaledObject &#x60;idleReplicaCount&#x60;: replicas to keep while no trigger is active. Only 0 is supported by KEDA, and it must be strictly less than the service&#39;s minimum running instances. Omitted or null disables the idle state. | [optional] 
**AdvancedYaml** | Pointer to **NullableString** | Raw YAML of the KEDA ScaledObject &#x60;advanced&#x60; field, e.g. &#x60;horizontalPodAutoscalerConfig.behavior&#x60; to tune how fast the HPA scales up and down. Must be a YAML mapping whose top-level keys are among &#x60;horizontalPodAutoscalerConfig&#x60;, &#x60;restoreToOriginalReplicaCount&#x60; and &#x60;scalingModifiers&#x60;. Omitted, null or blank means no advanced settings. | [optional] 
**Scalers** | [**[]KedaScalerRequest**](KedaScalerRequest.md) |  | 

## Methods

### NewKedaAutoscalingRequest

`func NewKedaAutoscalingRequest(mode AutoscalingMode, scalers []KedaScalerRequest, ) *KedaAutoscalingRequest`

NewKedaAutoscalingRequest instantiates a new KedaAutoscalingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKedaAutoscalingRequestWithDefaults

`func NewKedaAutoscalingRequestWithDefaults() *KedaAutoscalingRequest`

NewKedaAutoscalingRequestWithDefaults instantiates a new KedaAutoscalingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *KedaAutoscalingRequest) GetMode() AutoscalingMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *KedaAutoscalingRequest) GetModeOk() (*AutoscalingMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *KedaAutoscalingRequest) SetMode(v AutoscalingMode)`

SetMode sets Mode field to given value.


### GetPollingIntervalSeconds

`func (o *KedaAutoscalingRequest) GetPollingIntervalSeconds() int32`

GetPollingIntervalSeconds returns the PollingIntervalSeconds field if non-nil, zero value otherwise.

### GetPollingIntervalSecondsOk

`func (o *KedaAutoscalingRequest) GetPollingIntervalSecondsOk() (*int32, bool)`

GetPollingIntervalSecondsOk returns a tuple with the PollingIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingIntervalSeconds

`func (o *KedaAutoscalingRequest) SetPollingIntervalSeconds(v int32)`

SetPollingIntervalSeconds sets PollingIntervalSeconds field to given value.

### HasPollingIntervalSeconds

`func (o *KedaAutoscalingRequest) HasPollingIntervalSeconds() bool`

HasPollingIntervalSeconds returns a boolean if a field has been set.

### GetCooldownPeriodSeconds

`func (o *KedaAutoscalingRequest) GetCooldownPeriodSeconds() int32`

GetCooldownPeriodSeconds returns the CooldownPeriodSeconds field if non-nil, zero value otherwise.

### GetCooldownPeriodSecondsOk

`func (o *KedaAutoscalingRequest) GetCooldownPeriodSecondsOk() (*int32, bool)`

GetCooldownPeriodSecondsOk returns a tuple with the CooldownPeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownPeriodSeconds

`func (o *KedaAutoscalingRequest) SetCooldownPeriodSeconds(v int32)`

SetCooldownPeriodSeconds sets CooldownPeriodSeconds field to given value.

### HasCooldownPeriodSeconds

`func (o *KedaAutoscalingRequest) HasCooldownPeriodSeconds() bool`

HasCooldownPeriodSeconds returns a boolean if a field has been set.

### GetInitialCooldownPeriodSeconds

`func (o *KedaAutoscalingRequest) GetInitialCooldownPeriodSeconds() int32`

GetInitialCooldownPeriodSeconds returns the InitialCooldownPeriodSeconds field if non-nil, zero value otherwise.

### GetInitialCooldownPeriodSecondsOk

`func (o *KedaAutoscalingRequest) GetInitialCooldownPeriodSecondsOk() (*int32, bool)`

GetInitialCooldownPeriodSecondsOk returns a tuple with the InitialCooldownPeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCooldownPeriodSeconds

`func (o *KedaAutoscalingRequest) SetInitialCooldownPeriodSeconds(v int32)`

SetInitialCooldownPeriodSeconds sets InitialCooldownPeriodSeconds field to given value.

### HasInitialCooldownPeriodSeconds

`func (o *KedaAutoscalingRequest) HasInitialCooldownPeriodSeconds() bool`

HasInitialCooldownPeriodSeconds returns a boolean if a field has been set.

### SetInitialCooldownPeriodSecondsNil

`func (o *KedaAutoscalingRequest) SetInitialCooldownPeriodSecondsNil(b bool)`

 SetInitialCooldownPeriodSecondsNil sets the value for InitialCooldownPeriodSeconds to be an explicit nil

### UnsetInitialCooldownPeriodSeconds
`func (o *KedaAutoscalingRequest) UnsetInitialCooldownPeriodSeconds()`

UnsetInitialCooldownPeriodSeconds ensures that no value is present for InitialCooldownPeriodSeconds, not even an explicit nil
### GetIdleReplicaCount

`func (o *KedaAutoscalingRequest) GetIdleReplicaCount() int32`

GetIdleReplicaCount returns the IdleReplicaCount field if non-nil, zero value otherwise.

### GetIdleReplicaCountOk

`func (o *KedaAutoscalingRequest) GetIdleReplicaCountOk() (*int32, bool)`

GetIdleReplicaCountOk returns a tuple with the IdleReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleReplicaCount

`func (o *KedaAutoscalingRequest) SetIdleReplicaCount(v int32)`

SetIdleReplicaCount sets IdleReplicaCount field to given value.

### HasIdleReplicaCount

`func (o *KedaAutoscalingRequest) HasIdleReplicaCount() bool`

HasIdleReplicaCount returns a boolean if a field has been set.

### SetIdleReplicaCountNil

`func (o *KedaAutoscalingRequest) SetIdleReplicaCountNil(b bool)`

 SetIdleReplicaCountNil sets the value for IdleReplicaCount to be an explicit nil

### UnsetIdleReplicaCount
`func (o *KedaAutoscalingRequest) UnsetIdleReplicaCount()`

UnsetIdleReplicaCount ensures that no value is present for IdleReplicaCount, not even an explicit nil
### GetAdvancedYaml

`func (o *KedaAutoscalingRequest) GetAdvancedYaml() string`

GetAdvancedYaml returns the AdvancedYaml field if non-nil, zero value otherwise.

### GetAdvancedYamlOk

`func (o *KedaAutoscalingRequest) GetAdvancedYamlOk() (*string, bool)`

GetAdvancedYamlOk returns a tuple with the AdvancedYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedYaml

`func (o *KedaAutoscalingRequest) SetAdvancedYaml(v string)`

SetAdvancedYaml sets AdvancedYaml field to given value.

### HasAdvancedYaml

`func (o *KedaAutoscalingRequest) HasAdvancedYaml() bool`

HasAdvancedYaml returns a boolean if a field has been set.

### SetAdvancedYamlNil

`func (o *KedaAutoscalingRequest) SetAdvancedYamlNil(b bool)`

 SetAdvancedYamlNil sets the value for AdvancedYaml to be an explicit nil

### UnsetAdvancedYaml
`func (o *KedaAutoscalingRequest) UnsetAdvancedYaml()`

UnsetAdvancedYaml ensures that no value is present for AdvancedYaml, not even an explicit nil
### GetScalers

`func (o *KedaAutoscalingRequest) GetScalers() []KedaScalerRequest`

GetScalers returns the Scalers field if non-nil, zero value otherwise.

### GetScalersOk

`func (o *KedaAutoscalingRequest) GetScalersOk() (*[]KedaScalerRequest, bool)`

GetScalersOk returns a tuple with the Scalers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalers

`func (o *KedaAutoscalingRequest) SetScalers(v []KedaScalerRequest)`

SetScalers sets Scalers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


