# AutoscalingPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ServiceId** | **string** |  | 
**Mode** | [**AutoscalingMode**](AutoscalingMode.md) |  | 
**PollingIntervalSeconds** | **int32** |  | 
**CooldownPeriodSeconds** | **int32** |  | 
**InitialCooldownPeriodSeconds** | Pointer to **NullableInt32** | KEDA ScaledObject &#x60;initialCooldownPeriod&#x60;: seconds to wait after the ScaledObject is created before the service can scale to zero. Omitted or null leaves KEDA&#39;s default (0). | [optional] 
**IdleReplicaCount** | Pointer to **NullableInt32** | KEDA ScaledObject &#x60;idleReplicaCount&#x60;: replicas to keep while no trigger is active. Only 0 is supported by KEDA, and it must be strictly less than the service&#39;s minimum running instances. Omitted or null disables the idle state. | [optional] 
**AdvancedYaml** | Pointer to **NullableString** | Raw YAML of the KEDA ScaledObject &#x60;advanced&#x60; field, e.g. &#x60;horizontalPodAutoscalerConfig.behavior&#x60; to tune how fast the HPA scales up and down. Must be a YAML mapping whose top-level keys are among &#x60;horizontalPodAutoscalerConfig&#x60;, &#x60;restoreToOriginalReplicaCount&#x60; and &#x60;scalingModifiers&#x60;. Omitted, null or blank means no advanced settings. | [optional] 
**Scalers** | [**[]KedaScalerResponse**](KedaScalerResponse.md) |  | 

## Methods

### NewAutoscalingPolicyResponse

`func NewAutoscalingPolicyResponse(id string, createdAt time.Time, serviceId string, mode AutoscalingMode, pollingIntervalSeconds int32, cooldownPeriodSeconds int32, scalers []KedaScalerResponse, ) *AutoscalingPolicyResponse`

NewAutoscalingPolicyResponse instantiates a new AutoscalingPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalingPolicyResponseWithDefaults

`func NewAutoscalingPolicyResponseWithDefaults() *AutoscalingPolicyResponse`

NewAutoscalingPolicyResponseWithDefaults instantiates a new AutoscalingPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoscalingPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoscalingPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoscalingPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *AutoscalingPolicyResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutoscalingPolicyResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutoscalingPolicyResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AutoscalingPolicyResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AutoscalingPolicyResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AutoscalingPolicyResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AutoscalingPolicyResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetServiceId

`func (o *AutoscalingPolicyResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AutoscalingPolicyResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AutoscalingPolicyResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetMode

`func (o *AutoscalingPolicyResponse) GetMode() AutoscalingMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AutoscalingPolicyResponse) GetModeOk() (*AutoscalingMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AutoscalingPolicyResponse) SetMode(v AutoscalingMode)`

SetMode sets Mode field to given value.


### GetPollingIntervalSeconds

`func (o *AutoscalingPolicyResponse) GetPollingIntervalSeconds() int32`

GetPollingIntervalSeconds returns the PollingIntervalSeconds field if non-nil, zero value otherwise.

### GetPollingIntervalSecondsOk

`func (o *AutoscalingPolicyResponse) GetPollingIntervalSecondsOk() (*int32, bool)`

GetPollingIntervalSecondsOk returns a tuple with the PollingIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingIntervalSeconds

`func (o *AutoscalingPolicyResponse) SetPollingIntervalSeconds(v int32)`

SetPollingIntervalSeconds sets PollingIntervalSeconds field to given value.


### GetCooldownPeriodSeconds

`func (o *AutoscalingPolicyResponse) GetCooldownPeriodSeconds() int32`

GetCooldownPeriodSeconds returns the CooldownPeriodSeconds field if non-nil, zero value otherwise.

### GetCooldownPeriodSecondsOk

`func (o *AutoscalingPolicyResponse) GetCooldownPeriodSecondsOk() (*int32, bool)`

GetCooldownPeriodSecondsOk returns a tuple with the CooldownPeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownPeriodSeconds

`func (o *AutoscalingPolicyResponse) SetCooldownPeriodSeconds(v int32)`

SetCooldownPeriodSeconds sets CooldownPeriodSeconds field to given value.


### GetInitialCooldownPeriodSeconds

`func (o *AutoscalingPolicyResponse) GetInitialCooldownPeriodSeconds() int32`

GetInitialCooldownPeriodSeconds returns the InitialCooldownPeriodSeconds field if non-nil, zero value otherwise.

### GetInitialCooldownPeriodSecondsOk

`func (o *AutoscalingPolicyResponse) GetInitialCooldownPeriodSecondsOk() (*int32, bool)`

GetInitialCooldownPeriodSecondsOk returns a tuple with the InitialCooldownPeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCooldownPeriodSeconds

`func (o *AutoscalingPolicyResponse) SetInitialCooldownPeriodSeconds(v int32)`

SetInitialCooldownPeriodSeconds sets InitialCooldownPeriodSeconds field to given value.

### HasInitialCooldownPeriodSeconds

`func (o *AutoscalingPolicyResponse) HasInitialCooldownPeriodSeconds() bool`

HasInitialCooldownPeriodSeconds returns a boolean if a field has been set.

### SetInitialCooldownPeriodSecondsNil

`func (o *AutoscalingPolicyResponse) SetInitialCooldownPeriodSecondsNil(b bool)`

 SetInitialCooldownPeriodSecondsNil sets the value for InitialCooldownPeriodSeconds to be an explicit nil

### UnsetInitialCooldownPeriodSeconds
`func (o *AutoscalingPolicyResponse) UnsetInitialCooldownPeriodSeconds()`

UnsetInitialCooldownPeriodSeconds ensures that no value is present for InitialCooldownPeriodSeconds, not even an explicit nil
### GetIdleReplicaCount

`func (o *AutoscalingPolicyResponse) GetIdleReplicaCount() int32`

GetIdleReplicaCount returns the IdleReplicaCount field if non-nil, zero value otherwise.

### GetIdleReplicaCountOk

`func (o *AutoscalingPolicyResponse) GetIdleReplicaCountOk() (*int32, bool)`

GetIdleReplicaCountOk returns a tuple with the IdleReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleReplicaCount

`func (o *AutoscalingPolicyResponse) SetIdleReplicaCount(v int32)`

SetIdleReplicaCount sets IdleReplicaCount field to given value.

### HasIdleReplicaCount

`func (o *AutoscalingPolicyResponse) HasIdleReplicaCount() bool`

HasIdleReplicaCount returns a boolean if a field has been set.

### SetIdleReplicaCountNil

`func (o *AutoscalingPolicyResponse) SetIdleReplicaCountNil(b bool)`

 SetIdleReplicaCountNil sets the value for IdleReplicaCount to be an explicit nil

### UnsetIdleReplicaCount
`func (o *AutoscalingPolicyResponse) UnsetIdleReplicaCount()`

UnsetIdleReplicaCount ensures that no value is present for IdleReplicaCount, not even an explicit nil
### GetAdvancedYaml

`func (o *AutoscalingPolicyResponse) GetAdvancedYaml() string`

GetAdvancedYaml returns the AdvancedYaml field if non-nil, zero value otherwise.

### GetAdvancedYamlOk

`func (o *AutoscalingPolicyResponse) GetAdvancedYamlOk() (*string, bool)`

GetAdvancedYamlOk returns a tuple with the AdvancedYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedYaml

`func (o *AutoscalingPolicyResponse) SetAdvancedYaml(v string)`

SetAdvancedYaml sets AdvancedYaml field to given value.

### HasAdvancedYaml

`func (o *AutoscalingPolicyResponse) HasAdvancedYaml() bool`

HasAdvancedYaml returns a boolean if a field has been set.

### SetAdvancedYamlNil

`func (o *AutoscalingPolicyResponse) SetAdvancedYamlNil(b bool)`

 SetAdvancedYamlNil sets the value for AdvancedYaml to be an explicit nil

### UnsetAdvancedYaml
`func (o *AutoscalingPolicyResponse) UnsetAdvancedYaml()`

UnsetAdvancedYaml ensures that no value is present for AdvancedYaml, not even an explicit nil
### GetScalers

`func (o *AutoscalingPolicyResponse) GetScalers() []KedaScalerResponse`

GetScalers returns the Scalers field if non-nil, zero value otherwise.

### GetScalersOk

`func (o *AutoscalingPolicyResponse) GetScalersOk() (*[]KedaScalerResponse, bool)`

GetScalersOk returns a tuple with the Scalers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalers

`func (o *AutoscalingPolicyResponse) SetScalers(v []KedaScalerResponse)`

SetScalers sets Scalers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


