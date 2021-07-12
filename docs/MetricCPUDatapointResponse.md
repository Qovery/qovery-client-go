# MetricCPUDatapointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**RequestedInNumber** | Pointer to **float32** |  | [optional] 
**ConsumedInNumber** | **float32** |  | 
**ConsumedInPercent** | **float32** |  | 

## Methods

### NewMetricCPUDatapointResponse

`func NewMetricCPUDatapointResponse(createdAt time.Time, consumedInNumber float32, consumedInPercent float32, ) *MetricCPUDatapointResponse`

NewMetricCPUDatapointResponse instantiates a new MetricCPUDatapointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricCPUDatapointResponseWithDefaults

`func NewMetricCPUDatapointResponseWithDefaults() *MetricCPUDatapointResponse`

NewMetricCPUDatapointResponseWithDefaults instantiates a new MetricCPUDatapointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MetricCPUDatapointResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetricCPUDatapointResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetricCPUDatapointResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetRequestedInNumber

`func (o *MetricCPUDatapointResponse) GetRequestedInNumber() float32`

GetRequestedInNumber returns the RequestedInNumber field if non-nil, zero value otherwise.

### GetRequestedInNumberOk

`func (o *MetricCPUDatapointResponse) GetRequestedInNumberOk() (*float32, bool)`

GetRequestedInNumberOk returns a tuple with the RequestedInNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedInNumber

`func (o *MetricCPUDatapointResponse) SetRequestedInNumber(v float32)`

SetRequestedInNumber sets RequestedInNumber field to given value.

### HasRequestedInNumber

`func (o *MetricCPUDatapointResponse) HasRequestedInNumber() bool`

HasRequestedInNumber returns a boolean if a field has been set.

### GetConsumedInNumber

`func (o *MetricCPUDatapointResponse) GetConsumedInNumber() float32`

GetConsumedInNumber returns the ConsumedInNumber field if non-nil, zero value otherwise.

### GetConsumedInNumberOk

`func (o *MetricCPUDatapointResponse) GetConsumedInNumberOk() (*float32, bool)`

GetConsumedInNumberOk returns a tuple with the ConsumedInNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInNumber

`func (o *MetricCPUDatapointResponse) SetConsumedInNumber(v float32)`

SetConsumedInNumber sets ConsumedInNumber field to given value.


### GetConsumedInPercent

`func (o *MetricCPUDatapointResponse) GetConsumedInPercent() float32`

GetConsumedInPercent returns the ConsumedInPercent field if non-nil, zero value otherwise.

### GetConsumedInPercentOk

`func (o *MetricCPUDatapointResponse) GetConsumedInPercentOk() (*float32, bool)`

GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInPercent

`func (o *MetricCPUDatapointResponse) SetConsumedInPercent(v float32)`

SetConsumedInPercent sets ConsumedInPercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


