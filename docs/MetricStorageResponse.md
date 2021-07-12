# MetricStorageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageId** | Pointer to **string** |  | [optional] 
**Data** | [**[]MetricStorageDatapointResponse**](MetricStorageDatapointResponse.md) |  | 

## Methods

### NewMetricStorageResponse

`func NewMetricStorageResponse(data []MetricStorageDatapointResponse, ) *MetricStorageResponse`

NewMetricStorageResponse instantiates a new MetricStorageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricStorageResponseWithDefaults

`func NewMetricStorageResponseWithDefaults() *MetricStorageResponse`

NewMetricStorageResponseWithDefaults instantiates a new MetricStorageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageId

`func (o *MetricStorageResponse) GetStorageId() string`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *MetricStorageResponse) GetStorageIdOk() (*string, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *MetricStorageResponse) SetStorageId(v string)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *MetricStorageResponse) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetData

`func (o *MetricStorageResponse) GetData() []MetricStorageDatapointResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetricStorageResponse) GetDataOk() (*[]MetricStorageDatapointResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetricStorageResponse) SetData(v []MetricStorageDatapointResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


