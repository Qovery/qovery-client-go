# InstanceMemory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedInMb** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB. | [optional] 
**ConsumedInMb** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB. | [optional] 
**ConsumedInPercent** | Pointer to **float32** |  | [optional] 
**WarningThresholdInPercent** | Pointer to **float32** |  | [optional] 
**AlertThresholdInPercent** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to [**ThresholdMetricStatusEnum**](ThresholdMetricStatusEnum.md) |  | [optional] 

## Methods

### NewInstanceMemory

`func NewInstanceMemory() *InstanceMemory`

NewInstanceMemory instantiates a new InstanceMemory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceMemoryWithDefaults

`func NewInstanceMemoryWithDefaults() *InstanceMemory`

NewInstanceMemoryWithDefaults instantiates a new InstanceMemory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedInMb

`func (o *InstanceMemory) GetRequestedInMb() int32`

GetRequestedInMb returns the RequestedInMb field if non-nil, zero value otherwise.

### GetRequestedInMbOk

`func (o *InstanceMemory) GetRequestedInMbOk() (*int32, bool)`

GetRequestedInMbOk returns a tuple with the RequestedInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedInMb

`func (o *InstanceMemory) SetRequestedInMb(v int32)`

SetRequestedInMb sets RequestedInMb field to given value.

### HasRequestedInMb

`func (o *InstanceMemory) HasRequestedInMb() bool`

HasRequestedInMb returns a boolean if a field has been set.

### GetConsumedInMb

`func (o *InstanceMemory) GetConsumedInMb() int32`

GetConsumedInMb returns the ConsumedInMb field if non-nil, zero value otherwise.

### GetConsumedInMbOk

`func (o *InstanceMemory) GetConsumedInMbOk() (*int32, bool)`

GetConsumedInMbOk returns a tuple with the ConsumedInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInMb

`func (o *InstanceMemory) SetConsumedInMb(v int32)`

SetConsumedInMb sets ConsumedInMb field to given value.

### HasConsumedInMb

`func (o *InstanceMemory) HasConsumedInMb() bool`

HasConsumedInMb returns a boolean if a field has been set.

### GetConsumedInPercent

`func (o *InstanceMemory) GetConsumedInPercent() float32`

GetConsumedInPercent returns the ConsumedInPercent field if non-nil, zero value otherwise.

### GetConsumedInPercentOk

`func (o *InstanceMemory) GetConsumedInPercentOk() (*float32, bool)`

GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInPercent

`func (o *InstanceMemory) SetConsumedInPercent(v float32)`

SetConsumedInPercent sets ConsumedInPercent field to given value.

### HasConsumedInPercent

`func (o *InstanceMemory) HasConsumedInPercent() bool`

HasConsumedInPercent returns a boolean if a field has been set.

### GetWarningThresholdInPercent

`func (o *InstanceMemory) GetWarningThresholdInPercent() float32`

GetWarningThresholdInPercent returns the WarningThresholdInPercent field if non-nil, zero value otherwise.

### GetWarningThresholdInPercentOk

`func (o *InstanceMemory) GetWarningThresholdInPercentOk() (*float32, bool)`

GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThresholdInPercent

`func (o *InstanceMemory) SetWarningThresholdInPercent(v float32)`

SetWarningThresholdInPercent sets WarningThresholdInPercent field to given value.

### HasWarningThresholdInPercent

`func (o *InstanceMemory) HasWarningThresholdInPercent() bool`

HasWarningThresholdInPercent returns a boolean if a field has been set.

### GetAlertThresholdInPercent

`func (o *InstanceMemory) GetAlertThresholdInPercent() float32`

GetAlertThresholdInPercent returns the AlertThresholdInPercent field if non-nil, zero value otherwise.

### GetAlertThresholdInPercentOk

`func (o *InstanceMemory) GetAlertThresholdInPercentOk() (*float32, bool)`

GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholdInPercent

`func (o *InstanceMemory) SetAlertThresholdInPercent(v float32)`

SetAlertThresholdInPercent sets AlertThresholdInPercent field to given value.

### HasAlertThresholdInPercent

`func (o *InstanceMemory) HasAlertThresholdInPercent() bool`

HasAlertThresholdInPercent returns a boolean if a field has been set.

### GetStatus

`func (o *InstanceMemory) GetStatus() ThresholdMetricStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceMemory) GetStatusOk() (*ThresholdMetricStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceMemory) SetStatus(v ThresholdMetricStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstanceMemory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


