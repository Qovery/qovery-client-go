# MetricMemoryDatapointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**RequestedInMb** | **int32** |  | 
**ConsumedInMb** | **int32** |  | 
**ConsumedInPercent** | **float32** |  | 

## Methods

### NewMetricMemoryDatapointResponse

`func NewMetricMemoryDatapointResponse(createdAt time.Time, requestedInMb int32, consumedInMb int32, consumedInPercent float32, ) *MetricMemoryDatapointResponse`

NewMetricMemoryDatapointResponse instantiates a new MetricMemoryDatapointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricMemoryDatapointResponseWithDefaults

`func NewMetricMemoryDatapointResponseWithDefaults() *MetricMemoryDatapointResponse`

NewMetricMemoryDatapointResponseWithDefaults instantiates a new MetricMemoryDatapointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MetricMemoryDatapointResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetricMemoryDatapointResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetricMemoryDatapointResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetRequestedInMb

`func (o *MetricMemoryDatapointResponse) GetRequestedInMb() int32`

GetRequestedInMb returns the RequestedInMb field if non-nil, zero value otherwise.

### GetRequestedInMbOk

`func (o *MetricMemoryDatapointResponse) GetRequestedInMbOk() (*int32, bool)`

GetRequestedInMbOk returns a tuple with the RequestedInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedInMb

`func (o *MetricMemoryDatapointResponse) SetRequestedInMb(v int32)`

SetRequestedInMb sets RequestedInMb field to given value.


### GetConsumedInMb

`func (o *MetricMemoryDatapointResponse) GetConsumedInMb() int32`

GetConsumedInMb returns the ConsumedInMb field if non-nil, zero value otherwise.

### GetConsumedInMbOk

`func (o *MetricMemoryDatapointResponse) GetConsumedInMbOk() (*int32, bool)`

GetConsumedInMbOk returns a tuple with the ConsumedInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInMb

`func (o *MetricMemoryDatapointResponse) SetConsumedInMb(v int32)`

SetConsumedInMb sets ConsumedInMb field to given value.


### GetConsumedInPercent

`func (o *MetricMemoryDatapointResponse) GetConsumedInPercent() float32`

GetConsumedInPercent returns the ConsumedInPercent field if non-nil, zero value otherwise.

### GetConsumedInPercentOk

`func (o *MetricMemoryDatapointResponse) GetConsumedInPercentOk() (*float32, bool)`

GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInPercent

`func (o *MetricMemoryDatapointResponse) SetConsumedInPercent(v float32)`

SetConsumedInPercent sets ConsumedInPercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


