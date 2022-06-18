# EnvironmentDatabasesCurrentMetricMemory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedInMb** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] 
**ConsumedInMb** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] 
**ConsumedInPercent** | Pointer to **float32** |  | [optional] 
**WarningThresholdInPercent** | Pointer to **float32** |  | [optional] 
**AlertThresholdInPercent** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to [**ThresholdMetricStatusEnum**](ThresholdMetricStatusEnum.md) |  | [optional] 

## Methods

### NewEnvironmentDatabasesCurrentMetricMemory

`func NewEnvironmentDatabasesCurrentMetricMemory() *EnvironmentDatabasesCurrentMetricMemory`

NewEnvironmentDatabasesCurrentMetricMemory instantiates a new EnvironmentDatabasesCurrentMetricMemory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDatabasesCurrentMetricMemoryWithDefaults

`func NewEnvironmentDatabasesCurrentMetricMemoryWithDefaults() *EnvironmentDatabasesCurrentMetricMemory`

NewEnvironmentDatabasesCurrentMetricMemoryWithDefaults instantiates a new EnvironmentDatabasesCurrentMetricMemory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedInMb

`func (o *EnvironmentDatabasesCurrentMetricMemory) GetRequestedInMb() int32`

GetRequestedInMb returns the RequestedInMb field if non-nil, zero value otherwise.

### GetRequestedInMbOk

`func (o *EnvironmentDatabasesCurrentMetricMemory) GetRequestedInMbOk() (*int32, bool)`

GetRequestedInMbOk returns a tuple with the RequestedInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedInMb

`func (o *EnvironmentDatabasesCurrentMetricMemory) SetRequestedInMb(v int32)`

SetRequestedInMb sets RequestedInMb field to given value.

### HasRequestedInMb

`func (o *EnvironmentDatabasesCurrentMetricMemory) HasRequestedInMb() bool`

HasRequestedInMb returns a boolean if a field has been set.

### GetConsumedInMb

`func (o *EnvironmentDatabasesCurrentMetricMemory) GetConsumedInMb() int32`

GetConsumedInMb returns the ConsumedInMb field if non-nil, zero value otherwise.

### GetConsumedInMbOk

`func (o *EnvironmentDatabasesCurrentMetricMemory) GetConsumedInMbOk() (*int32, bool)`

GetConsumedInMbOk returns a tuple with the ConsumedInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInMb

`func (o *EnvironmentDatabasesCurrentMetricMemory) SetConsumedInMb(v int32)`

SetConsumedInMb sets ConsumedInMb field to given value.

### HasConsumedInMb

`func (o *EnvironmentDatabasesCurrentMetricMemory) HasConsumedInMb() bool`

HasConsumedInMb returns a boolean if a field has been set.

### GetConsumedInPercent

`func (o *EnvironmentDatabasesCurrentMetricMemory) GetConsumedInPercent() float32`

GetConsumedInPercent returns the ConsumedInPercent field if non-nil, zero value otherwise.

### GetConsumedInPercentOk

`func (o *EnvironmentDatabasesCurrentMetricMemory) GetConsumedInPercentOk() (*float32, bool)`

GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInPercent

`func (o *EnvironmentDatabasesCurrentMetricMemory) SetConsumedInPercent(v float32)`

SetConsumedInPercent sets ConsumedInPercent field to given value.

### HasConsumedInPercent

`func (o *EnvironmentDatabasesCurrentMetricMemory) HasConsumedInPercent() bool`

HasConsumedInPercent returns a boolean if a field has been set.

### GetWarningThresholdInPercent

`func (o *EnvironmentDatabasesCurrentMetricMemory) GetWarningThresholdInPercent() float32`

GetWarningThresholdInPercent returns the WarningThresholdInPercent field if non-nil, zero value otherwise.

### GetWarningThresholdInPercentOk

`func (o *EnvironmentDatabasesCurrentMetricMemory) GetWarningThresholdInPercentOk() (*float32, bool)`

GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThresholdInPercent

`func (o *EnvironmentDatabasesCurrentMetricMemory) SetWarningThresholdInPercent(v float32)`

SetWarningThresholdInPercent sets WarningThresholdInPercent field to given value.

### HasWarningThresholdInPercent

`func (o *EnvironmentDatabasesCurrentMetricMemory) HasWarningThresholdInPercent() bool`

HasWarningThresholdInPercent returns a boolean if a field has been set.

### GetAlertThresholdInPercent

`func (o *EnvironmentDatabasesCurrentMetricMemory) GetAlertThresholdInPercent() float32`

GetAlertThresholdInPercent returns the AlertThresholdInPercent field if non-nil, zero value otherwise.

### GetAlertThresholdInPercentOk

`func (o *EnvironmentDatabasesCurrentMetricMemory) GetAlertThresholdInPercentOk() (*float32, bool)`

GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholdInPercent

`func (o *EnvironmentDatabasesCurrentMetricMemory) SetAlertThresholdInPercent(v float32)`

SetAlertThresholdInPercent sets AlertThresholdInPercent field to given value.

### HasAlertThresholdInPercent

`func (o *EnvironmentDatabasesCurrentMetricMemory) HasAlertThresholdInPercent() bool`

HasAlertThresholdInPercent returns a boolean if a field has been set.

### GetStatus

`func (o *EnvironmentDatabasesCurrentMetricMemory) GetStatus() ThresholdMetricStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentDatabasesCurrentMetricMemory) GetStatusOk() (*ThresholdMetricStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentDatabasesCurrentMetricMemory) SetStatus(v ThresholdMetricStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnvironmentDatabasesCurrentMetricMemory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


