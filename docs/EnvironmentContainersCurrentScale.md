# EnvironmentContainersCurrentScale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** |  | [optional] 
**Min** | Pointer to **int32** |  | [optional] 
**Max** | Pointer to **int32** |  | [optional] 
**Running** | Pointer to **int32** |  | [optional] 
**RunningInPercent** | Pointer to **float32** |  | [optional] 
**WarningThresholdInPercent** | Pointer to **float32** |  | [optional] 
**AlertThresholdInPercent** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to [**ThresholdMetricStatusEnum**](ThresholdMetricStatusEnum.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewEnvironmentContainersCurrentScale

`func NewEnvironmentContainersCurrentScale() *EnvironmentContainersCurrentScale`

NewEnvironmentContainersCurrentScale instantiates a new EnvironmentContainersCurrentScale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentContainersCurrentScaleWithDefaults

`func NewEnvironmentContainersCurrentScaleWithDefaults() *EnvironmentContainersCurrentScale`

NewEnvironmentContainersCurrentScaleWithDefaults instantiates a new EnvironmentContainersCurrentScale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *EnvironmentContainersCurrentScale) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *EnvironmentContainersCurrentScale) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *EnvironmentContainersCurrentScale) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *EnvironmentContainersCurrentScale) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetMin

`func (o *EnvironmentContainersCurrentScale) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *EnvironmentContainersCurrentScale) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *EnvironmentContainersCurrentScale) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *EnvironmentContainersCurrentScale) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *EnvironmentContainersCurrentScale) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *EnvironmentContainersCurrentScale) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *EnvironmentContainersCurrentScale) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *EnvironmentContainersCurrentScale) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetRunning

`func (o *EnvironmentContainersCurrentScale) GetRunning() int32`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *EnvironmentContainersCurrentScale) GetRunningOk() (*int32, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *EnvironmentContainersCurrentScale) SetRunning(v int32)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *EnvironmentContainersCurrentScale) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetRunningInPercent

`func (o *EnvironmentContainersCurrentScale) GetRunningInPercent() float32`

GetRunningInPercent returns the RunningInPercent field if non-nil, zero value otherwise.

### GetRunningInPercentOk

`func (o *EnvironmentContainersCurrentScale) GetRunningInPercentOk() (*float32, bool)`

GetRunningInPercentOk returns a tuple with the RunningInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningInPercent

`func (o *EnvironmentContainersCurrentScale) SetRunningInPercent(v float32)`

SetRunningInPercent sets RunningInPercent field to given value.

### HasRunningInPercent

`func (o *EnvironmentContainersCurrentScale) HasRunningInPercent() bool`

HasRunningInPercent returns a boolean if a field has been set.

### GetWarningThresholdInPercent

`func (o *EnvironmentContainersCurrentScale) GetWarningThresholdInPercent() float32`

GetWarningThresholdInPercent returns the WarningThresholdInPercent field if non-nil, zero value otherwise.

### GetWarningThresholdInPercentOk

`func (o *EnvironmentContainersCurrentScale) GetWarningThresholdInPercentOk() (*float32, bool)`

GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThresholdInPercent

`func (o *EnvironmentContainersCurrentScale) SetWarningThresholdInPercent(v float32)`

SetWarningThresholdInPercent sets WarningThresholdInPercent field to given value.

### HasWarningThresholdInPercent

`func (o *EnvironmentContainersCurrentScale) HasWarningThresholdInPercent() bool`

HasWarningThresholdInPercent returns a boolean if a field has been set.

### GetAlertThresholdInPercent

`func (o *EnvironmentContainersCurrentScale) GetAlertThresholdInPercent() float32`

GetAlertThresholdInPercent returns the AlertThresholdInPercent field if non-nil, zero value otherwise.

### GetAlertThresholdInPercentOk

`func (o *EnvironmentContainersCurrentScale) GetAlertThresholdInPercentOk() (*float32, bool)`

GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholdInPercent

`func (o *EnvironmentContainersCurrentScale) SetAlertThresholdInPercent(v float32)`

SetAlertThresholdInPercent sets AlertThresholdInPercent field to given value.

### HasAlertThresholdInPercent

`func (o *EnvironmentContainersCurrentScale) HasAlertThresholdInPercent() bool`

HasAlertThresholdInPercent returns a boolean if a field has been set.

### GetStatus

`func (o *EnvironmentContainersCurrentScale) GetStatus() ThresholdMetricStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentContainersCurrentScale) GetStatusOk() (*ThresholdMetricStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentContainersCurrentScale) SetStatus(v ThresholdMetricStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnvironmentContainersCurrentScale) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentContainersCurrentScale) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentContainersCurrentScale) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentContainersCurrentScale) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentContainersCurrentScale) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


