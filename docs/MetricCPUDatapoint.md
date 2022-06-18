# MetricCPUDatapoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**RequestedInNumber** | Pointer to **float32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] 
**ConsumedInNumber** | **float32** | unit is millicores (m). 1000m &#x3D; 1 cpu | 
**ConsumedInPercent** | **float32** |  | 

## Methods

### NewMetricCPUDatapoint

`func NewMetricCPUDatapoint(createdAt time.Time, consumedInNumber float32, consumedInPercent float32, ) *MetricCPUDatapoint`

NewMetricCPUDatapoint instantiates a new MetricCPUDatapoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricCPUDatapointWithDefaults

`func NewMetricCPUDatapointWithDefaults() *MetricCPUDatapoint`

NewMetricCPUDatapointWithDefaults instantiates a new MetricCPUDatapoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MetricCPUDatapoint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetricCPUDatapoint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetricCPUDatapoint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetRequestedInNumber

`func (o *MetricCPUDatapoint) GetRequestedInNumber() float32`

GetRequestedInNumber returns the RequestedInNumber field if non-nil, zero value otherwise.

### GetRequestedInNumberOk

`func (o *MetricCPUDatapoint) GetRequestedInNumberOk() (*float32, bool)`

GetRequestedInNumberOk returns a tuple with the RequestedInNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedInNumber

`func (o *MetricCPUDatapoint) SetRequestedInNumber(v float32)`

SetRequestedInNumber sets RequestedInNumber field to given value.

### HasRequestedInNumber

`func (o *MetricCPUDatapoint) HasRequestedInNumber() bool`

HasRequestedInNumber returns a boolean if a field has been set.

### GetConsumedInNumber

`func (o *MetricCPUDatapoint) GetConsumedInNumber() float32`

GetConsumedInNumber returns the ConsumedInNumber field if non-nil, zero value otherwise.

### GetConsumedInNumberOk

`func (o *MetricCPUDatapoint) GetConsumedInNumberOk() (*float32, bool)`

GetConsumedInNumberOk returns a tuple with the ConsumedInNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInNumber

`func (o *MetricCPUDatapoint) SetConsumedInNumber(v float32)`

SetConsumedInNumber sets ConsumedInNumber field to given value.


### GetConsumedInPercent

`func (o *MetricCPUDatapoint) GetConsumedInPercent() float32`

GetConsumedInPercent returns the ConsumedInPercent field if non-nil, zero value otherwise.

### GetConsumedInPercentOk

`func (o *MetricCPUDatapoint) GetConsumedInPercentOk() (*float32, bool)`

GetConsumedInPercentOk returns a tuple with the ConsumedInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedInPercent

`func (o *MetricCPUDatapoint) SetConsumedInPercent(v float32)`

SetConsumedInPercent sets ConsumedInPercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


