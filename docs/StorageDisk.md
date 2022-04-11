# StorageDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**StorageId** | Pointer to **string** |  | [optional] 
**RequestedInGb** | Pointer to **int32** |  | [optional] 
**ConsumedInGb** | Pointer to **float32** |  | [optional] 
**ConsumedInPercent** | Pointer to **float32** |  | [optional] 
**WarningThresholdInPercent** | Pointer to **float32** |  | [optional] 
**AlertThresholdInPercent** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to [**ThresholdMetricStatusEnum**](ThresholdMetricStatusEnum.md) |  | [optional] 

## Methods

### NewStorageDisk

`func NewStorageDisk() *StorageDisk`

NewStorageDisk instantiates a new StorageDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDiskWithDefaults

`func NewStorageDiskWithDefaults() *StorageDisk`

NewStorageDiskWithDefaults instantiates a new StorageDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *StorageDisk) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StorageDisk) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StorageDisk) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StorageDisk) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStorageId

`func (o *StorageDisk) GetStorageId() string`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *StorageDisk) GetStorageIdOk() (*string, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *StorageDisk) SetStorageId(v string)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *StorageDisk) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetRequestedInGb

`func (o *StorageDisk) GetRequestedInGb() int32`

GetRequestedInGb returns the RequestedInGb field if non-nil, zero value otherwise.

### GetRequestedInGbOk

`func (o *StorageDisk) GetRequestedInGbOk() (*int32, bool)`

GetRequestedInGbOk returns a tuple with the RequestedInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedInGb

`func (o *StorageDisk) SetRequestedInGb(v int32)`

SetRequestedInGb sets RequestedInGb field to given value.

### HasRequestedInGb

`func (o *StorageDisk) HasRequestedInGb() bool`

HasRequestedInGb returns a boolean if a field has been set.

### GetConsumedInGb

`func (o *StorageDisk) GetConsumedInGb() float32`

GetConsumedInGb returns the ConsumedInGb field if non-nil, zero value otherwise.

### GetConsumedInGbOk

`func (o *StorageDisk) GetConsumedInGbOk() (*float32, bool)`

GetConsumedInGbOk returns a tuple with the ConsumedInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInGb

`func (o *StorageDisk) SetConsumedInGb(v float32)`

SetConsumedInGb sets ConsumedInGb field to given value.

### HasConsumedInGb

`func (o *StorageDisk) HasConsumedInGb() bool`

HasConsumedInGb returns a boolean if a field has been set.

### GetConsumedInPercent

`func (o *StorageDisk) GetConsumedInPercent() float32`

GetConsumedInPercent returns the ConsumedInPercent field if non-nil, zero value otherwise.

### GetConsumedInPercentOk

`func (o *StorageDisk) GetConsumedInPercentOk() (*float32, bool)`

GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInPercent

`func (o *StorageDisk) SetConsumedInPercent(v float32)`

SetConsumedInPercent sets ConsumedInPercent field to given value.

### HasConsumedInPercent

`func (o *StorageDisk) HasConsumedInPercent() bool`

HasConsumedInPercent returns a boolean if a field has been set.

### GetWarningThresholdInPercent

`func (o *StorageDisk) GetWarningThresholdInPercent() float32`

GetWarningThresholdInPercent returns the WarningThresholdInPercent field if non-nil, zero value otherwise.

### GetWarningThresholdInPercentOk

`func (o *StorageDisk) GetWarningThresholdInPercentOk() (*float32, bool)`

GetWarningThresholdInPercentOk returns a tuple with the WarningThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThresholdInPercent

`func (o *StorageDisk) SetWarningThresholdInPercent(v float32)`

SetWarningThresholdInPercent sets WarningThresholdInPercent field to given value.

### HasWarningThresholdInPercent

`func (o *StorageDisk) HasWarningThresholdInPercent() bool`

HasWarningThresholdInPercent returns a boolean if a field has been set.

### GetAlertThresholdInPercent

`func (o *StorageDisk) GetAlertThresholdInPercent() float32`

GetAlertThresholdInPercent returns the AlertThresholdInPercent field if non-nil, zero value otherwise.

### GetAlertThresholdInPercentOk

`func (o *StorageDisk) GetAlertThresholdInPercentOk() (*float32, bool)`

GetAlertThresholdInPercentOk returns a tuple with the AlertThresholdInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholdInPercent

`func (o *StorageDisk) SetAlertThresholdInPercent(v float32)`

SetAlertThresholdInPercent sets AlertThresholdInPercent field to given value.

### HasAlertThresholdInPercent

`func (o *StorageDisk) HasAlertThresholdInPercent() bool`

HasAlertThresholdInPercent returns a boolean if a field has been set.

### GetStatus

`func (o *StorageDisk) GetStatus() ThresholdMetricStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageDisk) GetStatusOk() (*ThresholdMetricStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageDisk) SetStatus(v ThresholdMetricStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageDisk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


