# MetricStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageId** | Pointer to **string** |  | [optional] 
**Data** | [**[]MetricStorageDatapoint**](MetricStorageDatapoint.md) |  | 

## Methods

### NewMetricStorage

`func NewMetricStorage(data []MetricStorageDatapoint, ) *MetricStorage`

NewMetricStorage instantiates a new MetricStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricStorageWithDefaults

`func NewMetricStorageWithDefaults() *MetricStorage`

NewMetricStorageWithDefaults instantiates a new MetricStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageId

`func (o *MetricStorage) GetStorageId() string`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *MetricStorage) GetStorageIdOk() (*string, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *MetricStorage) SetStorageId(v string)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *MetricStorage) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetData

`func (o *MetricStorage) GetData() []MetricStorageDatapoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetricStorage) GetDataOk() (*[]MetricStorageDatapoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetricStorage) SetData(v []MetricStorageDatapoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


