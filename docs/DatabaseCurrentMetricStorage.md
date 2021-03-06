# DatabaseCurrentMetricStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedInGb** | Pointer to **int32** |  | [optional] 
**ConsumedInGb** | Pointer to **int32** |  | [optional] 
**ConsumedInPercent** | Pointer to **float32** |  | [optional] 
**WarningThresholdInPercent** | Pointer to **float32** |  | [optional] 
**AlertThresholdInPercent** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to [**ThresholdMetricStatusEnum**](ThresholdMetricStatusEnum.md) |  | [optional] 

## Methods

### NewDatabaseCurrentMetricStorage

`func NewDatabaseCurrentMetricStorage() *DatabaseCurrentMetricStorage`

NewDatabaseCurrentMetricStorage instantiates a new DatabaseCurrentMetricStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseCurrentMetricStorageWithDefaults

`func NewDatabaseCurrentMetricStorageWithDefaults() *DatabaseCurrentMetricStorage`

NewDatabaseCurrentMetricStorageWithDefaults instantiates a new DatabaseCurrentMetricStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedInGb

`func (o *DatabaseCurrentMetricStorage) GetRequestedInGb() int32`

GetRequestedInGb returns the RequestedInGb field if non-nil, zero value otherwise.

### GetRequestedInGbOk

`func (o *DatabaseCurrentMetricStorage) GetRequestedInGbOk() (*int32, bool)`

GetRequestedInGbOk returns a tuple with the RequestedInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedInGb

`func (o *DatabaseCurrentMetricStorage) SetRequestedInGb(v int32)`

SetRequestedInGb sets RequestedInGb field to given value.

### HasRequestedInGb

`func (o *DatabaseCurrentMetricStorage) HasRequestedInGb() bool`

HasRequestedInGb returns a boolean if a field has been set.

### GetConsumedInGb

`func (o *DatabaseCurrentMetricStorage) GetConsumedInGb() int32`

GetConsumedInGb returns the ConsumedInGb field if non-nil, zero value otherwise.

### GetConsumedInGbOk

`func (o *DatabaseCurrentMetricStorage) GetConsumedInGbOk() (*int32, bool)`

GetConsumedInGbOk returns a tuple with the ConsumedInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInGb

`func (o *DatabaseCurrentMetricStorage) SetConsumedInGb(v int32)`

SetConsumedInGb sets ConsumedInGb field to given value.

### HasConsumedInGb

`func (o *DatabaseCurrentMetricStorage) HasConsumedInGb() bool`

HasConsumedInGb returns a boolean if a field has been set.

### GetConsumedInPercent

`func (o *DatabaseCurrentMetricStorage) GetConsumedInPercent() float32`

GetConsumedInPercent returns the ConsumedInPercent field if non-nil, zero value otherwise.

### GetConsumedInPercentOk

`func (o *DatabaseCurrentMetricStorage) GetConsumedInPercentOk() (*float32, bool)`

GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInPercent

`func (o *DatabaseCurrentMetricStorage) SetConsumedInPercent(v float32)`

SetConsumedInPercent sets ConsumedInPercent field to given value.

### HasConsumedInPercent

`func (o *DatabaseCurrentMetricStorage) HasConsumedInPercent() bool`

HasConsumedInPercent returns a boolean if a field has been set.

### GetWarningThresholdInPercent

`func (o *DatabaseCurrentMetricStorage) GetWarningThresholdInPercent() float32`

GetWarningThresholdInPercent returns the WarningThresholdInPercent field if non-nil, zero value otherwise.

### GetWarningThresholdInPercentOk

`func (o *DatabaseCurrentMetricStorage) GetWarningThresholdInPercentOk() (*float32, bool)`

GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThresholdInPercent

`func (o *DatabaseCurrentMetricStorage) SetWarningThresholdInPercent(v float32)`

SetWarningThresholdInPercent sets WarningThresholdInPercent field to given value.

### HasWarningThresholdInPercent

`func (o *DatabaseCurrentMetricStorage) HasWarningThresholdInPercent() bool`

HasWarningThresholdInPercent returns a boolean if a field has been set.

### GetAlertThresholdInPercent

`func (o *DatabaseCurrentMetricStorage) GetAlertThresholdInPercent() float32`

GetAlertThresholdInPercent returns the AlertThresholdInPercent field if non-nil, zero value otherwise.

### GetAlertThresholdInPercentOk

`func (o *DatabaseCurrentMetricStorage) GetAlertThresholdInPercentOk() (*float32, bool)`

GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholdInPercent

`func (o *DatabaseCurrentMetricStorage) SetAlertThresholdInPercent(v float32)`

SetAlertThresholdInPercent sets AlertThresholdInPercent field to given value.

### HasAlertThresholdInPercent

`func (o *DatabaseCurrentMetricStorage) HasAlertThresholdInPercent() bool`

HasAlertThresholdInPercent returns a boolean if a field has been set.

### GetStatus

`func (o *DatabaseCurrentMetricStorage) GetStatus() ThresholdMetricStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseCurrentMetricStorage) GetStatusOk() (*ThresholdMetricStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseCurrentMetricStorage) SetStatus(v ThresholdMetricStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatabaseCurrentMetricStorage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


