# DatabaseCurrentMetricCpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedInFloat** | Pointer to **float32** |  | [optional] 
**ConsumedInNumber** | Pointer to **float32** |  | [optional] 
**ConsumedInPercent** | Pointer to **float32** |  | [optional] 
**WarningThresholdInPercent** | Pointer to **float32** |  | [optional] 
**AlertThresholdInPercent** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to [**ThresholdMetricStatusEnum**](ThresholdMetricStatusEnum.md) |  | [optional] 

## Methods

### NewDatabaseCurrentMetricCpu

`func NewDatabaseCurrentMetricCpu() *DatabaseCurrentMetricCpu`

NewDatabaseCurrentMetricCpu instantiates a new DatabaseCurrentMetricCpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseCurrentMetricCpuWithDefaults

`func NewDatabaseCurrentMetricCpuWithDefaults() *DatabaseCurrentMetricCpu`

NewDatabaseCurrentMetricCpuWithDefaults instantiates a new DatabaseCurrentMetricCpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedInFloat

`func (o *DatabaseCurrentMetricCpu) GetRequestedInFloat() float32`

GetRequestedInFloat returns the RequestedInFloat field if non-nil, zero value otherwise.

### GetRequestedInFloatOk

`func (o *DatabaseCurrentMetricCpu) GetRequestedInFloatOk() (*float32, bool)`

GetRequestedInFloatOk returns a tuple with the RequestedInFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedInFloat

`func (o *DatabaseCurrentMetricCpu) SetRequestedInFloat(v float32)`

SetRequestedInFloat sets RequestedInFloat field to given value.

### HasRequestedInFloat

`func (o *DatabaseCurrentMetricCpu) HasRequestedInFloat() bool`

HasRequestedInFloat returns a boolean if a field has been set.

### GetConsumedInNumber

`func (o *DatabaseCurrentMetricCpu) GetConsumedInNumber() float32`

GetConsumedInNumber returns the ConsumedInNumber field if non-nil, zero value otherwise.

### GetConsumedInNumberOk

`func (o *DatabaseCurrentMetricCpu) GetConsumedInNumberOk() (*float32, bool)`

GetConsumedInNumberOk returns a tuple with the ConsumedInNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInNumber

`func (o *DatabaseCurrentMetricCpu) SetConsumedInNumber(v float32)`

SetConsumedInNumber sets ConsumedInNumber field to given value.

### HasConsumedInNumber

`func (o *DatabaseCurrentMetricCpu) HasConsumedInNumber() bool`

HasConsumedInNumber returns a boolean if a field has been set.

### GetConsumedInPercent

`func (o *DatabaseCurrentMetricCpu) GetConsumedInPercent() float32`

GetConsumedInPercent returns the ConsumedInPercent field if non-nil, zero value otherwise.

### GetConsumedInPercentOk

`func (o *DatabaseCurrentMetricCpu) GetConsumedInPercentOk() (*float32, bool)`

GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInPercent

`func (o *DatabaseCurrentMetricCpu) SetConsumedInPercent(v float32)`

SetConsumedInPercent sets ConsumedInPercent field to given value.

### HasConsumedInPercent

`func (o *DatabaseCurrentMetricCpu) HasConsumedInPercent() bool`

HasConsumedInPercent returns a boolean if a field has been set.

### GetWarningThresholdInPercent

`func (o *DatabaseCurrentMetricCpu) GetWarningThresholdInPercent() float32`

GetWarningThresholdInPercent returns the WarningThresholdInPercent field if non-nil, zero value otherwise.

### GetWarningThresholdInPercentOk

`func (o *DatabaseCurrentMetricCpu) GetWarningThresholdInPercentOk() (*float32, bool)`

GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThresholdInPercent

`func (o *DatabaseCurrentMetricCpu) SetWarningThresholdInPercent(v float32)`

SetWarningThresholdInPercent sets WarningThresholdInPercent field to given value.

### HasWarningThresholdInPercent

`func (o *DatabaseCurrentMetricCpu) HasWarningThresholdInPercent() bool`

HasWarningThresholdInPercent returns a boolean if a field has been set.

### GetAlertThresholdInPercent

`func (o *DatabaseCurrentMetricCpu) GetAlertThresholdInPercent() float32`

GetAlertThresholdInPercent returns the AlertThresholdInPercent field if non-nil, zero value otherwise.

### GetAlertThresholdInPercentOk

`func (o *DatabaseCurrentMetricCpu) GetAlertThresholdInPercentOk() (*float32, bool)`

GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholdInPercent

`func (o *DatabaseCurrentMetricCpu) SetAlertThresholdInPercent(v float32)`

SetAlertThresholdInPercent sets AlertThresholdInPercent field to given value.

### HasAlertThresholdInPercent

`func (o *DatabaseCurrentMetricCpu) HasAlertThresholdInPercent() bool`

HasAlertThresholdInPercent returns a boolean if a field has been set.

### GetStatus

`func (o *DatabaseCurrentMetricCpu) GetStatus() ThresholdMetricStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseCurrentMetricCpu) GetStatusOk() (*ThresholdMetricStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseCurrentMetricCpu) SetStatus(v ThresholdMetricStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatabaseCurrentMetricCpu) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


