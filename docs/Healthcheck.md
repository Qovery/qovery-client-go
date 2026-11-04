# Healthcheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadinessProbe** | Pointer to [**NullableProbe**](Probe.md) |  | [optional] 
**LivenessProbe** | Pointer to [**NullableProbe**](Probe.md) |  | [optional] 
**StartupProbe** | Pointer to [**NullableHealthcheckStartupProbe**](HealthcheckStartupProbe.md) |  | [optional] 

## Methods

### NewHealthcheck

`func NewHealthcheck() *Healthcheck`

NewHealthcheck instantiates a new Healthcheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthcheckWithDefaults

`func NewHealthcheckWithDefaults() *Healthcheck`

NewHealthcheckWithDefaults instantiates a new Healthcheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadinessProbe

`func (o *Healthcheck) GetReadinessProbe() Probe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *Healthcheck) GetReadinessProbeOk() (*Probe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *Healthcheck) SetReadinessProbe(v Probe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *Healthcheck) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### SetReadinessProbeNil

`func (o *Healthcheck) SetReadinessProbeNil(b bool)`

 SetReadinessProbeNil sets the value for ReadinessProbe to be an explicit nil

### UnsetReadinessProbe
`func (o *Healthcheck) UnsetReadinessProbe()`

UnsetReadinessProbe ensures that no value is present for ReadinessProbe, not even an explicit nil
### GetLivenessProbe

`func (o *Healthcheck) GetLivenessProbe() Probe`

GetLivenessProbe returns the LivenessProbe field if non-nil, zero value otherwise.

### GetLivenessProbeOk

`func (o *Healthcheck) GetLivenessProbeOk() (*Probe, bool)`

GetLivenessProbeOk returns a tuple with the LivenessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbe

`func (o *Healthcheck) SetLivenessProbe(v Probe)`

SetLivenessProbe sets LivenessProbe field to given value.

### HasLivenessProbe

`func (o *Healthcheck) HasLivenessProbe() bool`

HasLivenessProbe returns a boolean if a field has been set.

### SetLivenessProbeNil

`func (o *Healthcheck) SetLivenessProbeNil(b bool)`

 SetLivenessProbeNil sets the value for LivenessProbe to be an explicit nil

### UnsetLivenessProbe
`func (o *Healthcheck) UnsetLivenessProbe()`

UnsetLivenessProbe ensures that no value is present for LivenessProbe, not even an explicit nil
### GetStartupProbe

`func (o *Healthcheck) GetStartupProbe() HealthcheckStartupProbe`

GetStartupProbe returns the StartupProbe field if non-nil, zero value otherwise.

### GetStartupProbeOk

`func (o *Healthcheck) GetStartupProbeOk() (*HealthcheckStartupProbe, bool)`

GetStartupProbeOk returns a tuple with the StartupProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupProbe

`func (o *Healthcheck) SetStartupProbe(v HealthcheckStartupProbe)`

SetStartupProbe sets StartupProbe field to given value.

### HasStartupProbe

`func (o *Healthcheck) HasStartupProbe() bool`

HasStartupProbe returns a boolean if a field has been set.

### SetStartupProbeNil

`func (o *Healthcheck) SetStartupProbeNil(b bool)`

 SetStartupProbeNil sets the value for StartupProbe to be an explicit nil

### UnsetStartupProbe
`func (o *Healthcheck) UnsetStartupProbe()`

UnsetStartupProbe ensures that no value is present for StartupProbe, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


