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


