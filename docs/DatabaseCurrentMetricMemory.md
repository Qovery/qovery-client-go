# DatabaseCurrentMetricMemory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedInMb** | Pointer to **int32** |  | [optional] 
**ConsumedInMb** | Pointer to **int32** |  | [optional] 
**ConsumedInPercent** | Pointer to **float32** |  | [optional] 
**WarningThresholdInPercent** | Pointer to **float32** |  | [optional] 
**AlertThresholdInPercent** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to [**ThresholdMetricStatusEnum**](ThresholdMetricStatusEnum.md) |  | [optional] 

## Methods

### NewDatabaseCurrentMetricMemory

`func NewDatabaseCurrentMetricMemory() *DatabaseCurrentMetricMemory`

NewDatabaseCurrentMetricMemory instantiates a new DatabaseCurrentMetricMemory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseCurrentMetricMemoryWithDefaults

`func NewDatabaseCurrentMetricMemoryWithDefaults() *DatabaseCurrentMetricMemory`

NewDatabaseCurrentMetricMemoryWithDefaults instantiates a new DatabaseCurrentMetricMemory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedInMb

`func (o *DatabaseCurrentMetricMemory) GetRequestedInMb() int32`

GetRequestedInMb returns the RequestedInMb field if non-nil, zero value otherwise.

### GetRequestedInMbOk

`func (o *DatabaseCurrentMetricMemory) GetRequestedInMbOk() (*int32, bool)`

GetRequestedInMbOk returns a tuple with the RequestedInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedInMb

`func (o *DatabaseCurrentMetricMemory) SetRequestedInMb(v int32)`

SetRequestedInMb sets RequestedInMb field to given value.

### HasRequestedInMb

`func (o *DatabaseCurrentMetricMemory) HasRequestedInMb() bool`

HasRequestedInMb returns a boolean if a field has been set.

### GetConsumedInMb

`func (o *DatabaseCurrentMetricMemory) GetConsumedInMb() int32`

GetConsumedInMb returns the ConsumedInMb field if non-nil, zero value otherwise.

### GetConsumedInMbOk

`func (o *DatabaseCurrentMetricMemory) GetConsumedInMbOk() (*int32, bool)`

GetConsumedInMbOk returns a tuple with the ConsumedInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInMb

`func (o *DatabaseCurrentMetricMemory) SetConsumedInMb(v int32)`

SetConsumedInMb sets ConsumedInMb field to given value.

### HasConsumedInMb

`func (o *DatabaseCurrentMetricMemory) HasConsumedInMb() bool`

HasConsumedInMb returns a boolean if a field has been set.

### GetConsumedInPercent

`func (o *DatabaseCurrentMetricMemory) GetConsumedInPercent() float32`

GetConsumedInPercent returns the ConsumedInPercent field if non-nil, zero value otherwise.

### GetConsumedInPercentOk

`func (o *DatabaseCurrentMetricMemory) GetConsumedInPercentOk() (*float32, bool)`

GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInPercent

`func (o *DatabaseCurrentMetricMemory) SetConsumedInPercent(v float32)`

SetConsumedInPercent sets ConsumedInPercent field to given value.

### HasConsumedInPercent

`func (o *DatabaseCurrentMetricMemory) HasConsumedInPercent() bool`

HasConsumedInPercent returns a boolean if a field has been set.

### GetWarningThresholdInPercent

`func (o *DatabaseCurrentMetricMemory) GetWarningThresholdInPercent() float32`

GetWarningThresholdInPercent returns the WarningThresholdInPercent field if non-nil, zero value otherwise.

### GetWarningThresholdInPercentOk

`func (o *DatabaseCurrentMetricMemory) GetWarningThresholdInPercentOk() (*float32, bool)`

GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThresholdInPercent

`func (o *DatabaseCurrentMetricMemory) SetWarningThresholdInPercent(v float32)`

SetWarningThresholdInPercent sets WarningThresholdInPercent field to given value.

### HasWarningThresholdInPercent

`func (o *DatabaseCurrentMetricMemory) HasWarningThresholdInPercent() bool`

HasWarningThresholdInPercent returns a boolean if a field has been set.

### GetAlertThresholdInPercent

`func (o *DatabaseCurrentMetricMemory) GetAlertThresholdInPercent() float32`

GetAlertThresholdInPercent returns the AlertThresholdInPercent field if non-nil, zero value otherwise.

### GetAlertThresholdInPercentOk

`func (o *DatabaseCurrentMetricMemory) GetAlertThresholdInPercentOk() (*float32, bool)`

GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholdInPercent

`func (o *DatabaseCurrentMetricMemory) SetAlertThresholdInPercent(v float32)`

SetAlertThresholdInPercent sets AlertThresholdInPercent field to given value.

### HasAlertThresholdInPercent

`func (o *DatabaseCurrentMetricMemory) HasAlertThresholdInPercent() bool`

HasAlertThresholdInPercent returns a boolean if a field has been set.

### GetStatus

`func (o *DatabaseCurrentMetricMemory) GetStatus() ThresholdMetricStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseCurrentMetricMemory) GetStatusOk() (*ThresholdMetricStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseCurrentMetricMemory) SetStatus(v ThresholdMetricStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatabaseCurrentMetricMemory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


