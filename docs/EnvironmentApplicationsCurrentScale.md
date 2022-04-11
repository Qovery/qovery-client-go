# EnvironmentApplicationsCurrentScale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** |  | [optional] 
**Min** | Pointer to **int32** |  | [optional] 
**Max** | Pointer to **int32** |  | [optional] 
**Running** | Pointer to **int32** |  | [optional] 
**RunningInPercent** | Pointer to **float32** |  | [optional] 
**WarningThresholdInPercent** | Pointer to **float32** |  | [optional] 
**AlertThresholdInPercent** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to [**ThresholdMetricStatusEnum**](ThresholdMetricStatusEnum.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewEnvironmentApplicationsCurrentScale

`func NewEnvironmentApplicationsCurrentScale() *EnvironmentApplicationsCurrentScale`

NewEnvironmentApplicationsCurrentScale instantiates a new EnvironmentApplicationsCurrentScale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentApplicationsCurrentScaleWithDefaults

`func NewEnvironmentApplicationsCurrentScaleWithDefaults() *EnvironmentApplicationsCurrentScale`

NewEnvironmentApplicationsCurrentScaleWithDefaults instantiates a new EnvironmentApplicationsCurrentScale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *EnvironmentApplicationsCurrentScale) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *EnvironmentApplicationsCurrentScale) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *EnvironmentApplicationsCurrentScale) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *EnvironmentApplicationsCurrentScale) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetMin

`func (o *EnvironmentApplicationsCurrentScale) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *EnvironmentApplicationsCurrentScale) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *EnvironmentApplicationsCurrentScale) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *EnvironmentApplicationsCurrentScale) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *EnvironmentApplicationsCurrentScale) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *EnvironmentApplicationsCurrentScale) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *EnvironmentApplicationsCurrentScale) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *EnvironmentApplicationsCurrentScale) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetRunning

`func (o *EnvironmentApplicationsCurrentScale) GetRunning() int32`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *EnvironmentApplicationsCurrentScale) GetRunningOk() (*int32, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *EnvironmentApplicationsCurrentScale) SetRunning(v int32)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *EnvironmentApplicationsCurrentScale) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetRunningInPercent

`func (o *EnvironmentApplicationsCurrentScale) GetRunningInPercent() float32`

GetRunningInPercent returns the RunningInPercent field if non-nil, zero value otherwise.

### GetRunningInPercentOk

`func (o *EnvironmentApplicationsCurrentScale) GetRunningInPercentOk() (*float32, bool)`

GetRunningInPercentOk returns a tuple with the RunningInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningInPercent

`func (o *EnvironmentApplicationsCurrentScale) SetRunningInPercent(v float32)`

SetRunningInPercent sets RunningInPercent field to given value.

### HasRunningInPercent

`func (o *EnvironmentApplicationsCurrentScale) HasRunningInPercent() bool`

HasRunningInPercent returns a boolean if a field has been set.

### GetWarningThresholdInPercent

`func (o *EnvironmentApplicationsCurrentScale) GetWarningThresholdInPercent() float32`

GetWarningThresholdInPercent returns the WarningThresholdInPercent field if non-nil, zero value otherwise.

### GetWarningThresholdInPercentOk

`func (o *EnvironmentApplicationsCurrentScale) GetWarningThresholdInPercentOk() (*float32, bool)`

GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThresholdInPercent

`func (o *EnvironmentApplicationsCurrentScale) SetWarningThresholdInPercent(v float32)`

SetWarningThresholdInPercent sets WarningThresholdInPercent field to given value.

### HasWarningThresholdInPercent

`func (o *EnvironmentApplicationsCurrentScale) HasWarningThresholdInPercent() bool`

HasWarningThresholdInPercent returns a boolean if a field has been set.

### GetAlertThresholdInPercent

`func (o *EnvironmentApplicationsCurrentScale) GetAlertThresholdInPercent() float32`

GetAlertThresholdInPercent returns the AlertThresholdInPercent field if non-nil, zero value otherwise.

### GetAlertThresholdInPercentOk

`func (o *EnvironmentApplicationsCurrentScale) GetAlertThresholdInPercentOk() (*float32, bool)`

GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholdInPercent

`func (o *EnvironmentApplicationsCurrentScale) SetAlertThresholdInPercent(v float32)`

SetAlertThresholdInPercent sets AlertThresholdInPercent field to given value.

### HasAlertThresholdInPercent

`func (o *EnvironmentApplicationsCurrentScale) HasAlertThresholdInPercent() bool`

HasAlertThresholdInPercent returns a boolean if a field has been set.

### GetStatus

`func (o *EnvironmentApplicationsCurrentScale) GetStatus() ThresholdMetricStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentApplicationsCurrentScale) GetStatusOk() (*ThresholdMetricStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentApplicationsCurrentScale) SetStatus(v ThresholdMetricStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnvironmentApplicationsCurrentScale) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentApplicationsCurrentScale) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentApplicationsCurrentScale) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentApplicationsCurrentScale) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentApplicationsCurrentScale) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


