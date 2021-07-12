# MetricStorageDatapointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**RequestedInGb** | Pointer to **int32** |  | [optional] 
**ConsumedInGb** | Pointer to **float32** |  | [optional] 
**ConsumedInPercent** | **float32** |  | 

## Methods

### NewMetricStorageDatapointResponse

`func NewMetricStorageDatapointResponse(createdAt time.Time, consumedInPercent float32, ) *MetricStorageDatapointResponse`

NewMetricStorageDatapointResponse instantiates a new MetricStorageDatapointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricStorageDatapointResponseWithDefaults

`func NewMetricStorageDatapointResponseWithDefaults() *MetricStorageDatapointResponse`

NewMetricStorageDatapointResponseWithDefaults instantiates a new MetricStorageDatapointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MetricStorageDatapointResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetricStorageDatapointResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetricStorageDatapointResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetRequestedInGb

`func (o *MetricStorageDatapointResponse) GetRequestedInGb() int32`

GetRequestedInGb returns the RequestedInGb field if non-nil, zero value otherwise.

### GetRequestedInGbOk

`func (o *MetricStorageDatapointResponse) GetRequestedInGbOk() (*int32, bool)`

GetRequestedInGbOk returns a tuple with the RequestedInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedInGb

`func (o *MetricStorageDatapointResponse) SetRequestedInGb(v int32)`

SetRequestedInGb sets RequestedInGb field to given value.

### HasRequestedInGb

`func (o *MetricStorageDatapointResponse) HasRequestedInGb() bool`

HasRequestedInGb returns a boolean if a field has been set.

### GetConsumedInGb

`func (o *MetricStorageDatapointResponse) GetConsumedInGb() float32`

GetConsumedInGb returns the ConsumedInGb field if non-nil, zero value otherwise.

### GetConsumedInGbOk

`func (o *MetricStorageDatapointResponse) GetConsumedInGbOk() (*float32, bool)`

GetConsumedInGbOk returns a tuple with the ConsumedInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInGb

`func (o *MetricStorageDatapointResponse) SetConsumedInGb(v float32)`

SetConsumedInGb sets ConsumedInGb field to given value.

### HasConsumedInGb

`func (o *MetricStorageDatapointResponse) HasConsumedInGb() bool`

HasConsumedInGb returns a boolean if a field has been set.

### GetConsumedInPercent

`func (o *MetricStorageDatapointResponse) GetConsumedInPercent() float32`

GetConsumedInPercent returns the ConsumedInPercent field if non-nil, zero value otherwise.

### GetConsumedInPercentOk

`func (o *MetricStorageDatapointResponse) GetConsumedInPercentOk() (*float32, bool)`

GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInPercent

`func (o *MetricStorageDatapointResponse) SetConsumedInPercent(v float32)`

SetConsumedInPercent sets ConsumedInPercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


