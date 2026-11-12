# AutoscalingPolicyRequest

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

### GetInitialCooldownPeriodSeconds

`func (o *AutoscalingPolicyRequest) GetInitialCooldownPeriodSeconds() int32`

GetInitialCooldownPeriodSeconds returns the InitialCooldownPeriodSeconds field if non-nil, zero value otherwise.

### GetInitialCooldownPeriodSecondsOk

`func (o *AutoscalingPolicyRequest) GetInitialCooldownPeriodSecondsOk() (*int32, bool)`

GetInitialCooldownPeriodSecondsOk returns a tuple with the InitialCooldownPeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCooldownPeriodSeconds

`func (o *AutoscalingPolicyRequest) SetInitialCooldownPeriodSeconds(v int32)`

SetInitialCooldownPeriodSeconds sets InitialCooldownPeriodSeconds field to given value.

### HasInitialCooldownPeriodSeconds

`func (o *AutoscalingPolicyRequest) HasInitialCooldownPeriodSeconds() bool`

HasInitialCooldownPeriodSeconds returns a boolean if a field has been set.

### SetInitialCooldownPeriodSecondsNil

`func (o *AutoscalingPolicyRequest) SetInitialCooldownPeriodSecondsNil(b bool)`

 SetInitialCooldownPeriodSecondsNil sets the value for InitialCooldownPeriodSeconds to be an explicit nil

### UnsetInitialCooldownPeriodSeconds
`func (o *AutoscalingPolicyRequest) UnsetInitialCooldownPeriodSeconds()`

UnsetInitialCooldownPeriodSeconds ensures that no value is present for InitialCooldownPeriodSeconds, not even an explicit nil
### GetIdleReplicaCount

`func (o *AutoscalingPolicyRequest) GetIdleReplicaCount() int32`

GetIdleReplicaCount returns the IdleReplicaCount field if non-nil, zero value otherwise.

### GetIdleReplicaCountOk

`func (o *AutoscalingPolicyRequest) GetIdleReplicaCountOk() (*int32, bool)`

GetIdleReplicaCountOk returns a tuple with the IdleReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleReplicaCount

`func (o *AutoscalingPolicyRequest) SetIdleReplicaCount(v int32)`

SetIdleReplicaCount sets IdleReplicaCount field to given value.

### HasIdleReplicaCount

`func (o *AutoscalingPolicyRequest) HasIdleReplicaCount() bool`

HasIdleReplicaCount returns a boolean if a field has been set.

### SetIdleReplicaCountNil

`func (o *AutoscalingPolicyRequest) SetIdleReplicaCountNil(b bool)`

 SetIdleReplicaCountNil sets the value for IdleReplicaCount to be an explicit nil

### UnsetIdleReplicaCount
`func (o *AutoscalingPolicyRequest) UnsetIdleReplicaCount()`

UnsetIdleReplicaCount ensures that no value is present for IdleReplicaCount, not even an explicit nil
### GetAdvancedYaml

`func (o *AutoscalingPolicyRequest) GetAdvancedYaml() string`

GetAdvancedYaml returns the AdvancedYaml field if non-nil, zero value otherwise.

### GetAdvancedYamlOk

`func (o *AutoscalingPolicyRequest) GetAdvancedYamlOk() (*string, bool)`

GetAdvancedYamlOk returns a tuple with the AdvancedYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedYaml

`func (o *AutoscalingPolicyRequest) SetAdvancedYaml(v string)`

SetAdvancedYaml sets AdvancedYaml field to given value.

### HasAdvancedYaml

`func (o *AutoscalingPolicyRequest) HasAdvancedYaml() bool`

HasAdvancedYaml returns a boolean if a field has been set.

### SetAdvancedYamlNil

`func (o *AutoscalingPolicyRequest) SetAdvancedYamlNil(b bool)`

 SetAdvancedYamlNil sets the value for AdvancedYaml to be an explicit nil

### UnsetAdvancedYaml
`func (o *AutoscalingPolicyRequest) UnsetAdvancedYaml()`

UnsetAdvancedYaml ensures that no value is present for AdvancedYaml, not even an explicit nil
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


