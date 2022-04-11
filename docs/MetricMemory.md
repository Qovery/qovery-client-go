# MetricMemory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceName** | **string** |  | 
**Data** | [**[]MetricMemoryDatapoint**](MetricMemoryDatapoint.md) |  | 

## Methods

### NewMetricMemory

`func NewMetricMemory(instanceName string, data []MetricMemoryDatapoint, ) *MetricMemory`

NewMetricMemory instantiates a new MetricMemory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricMemoryWithDefaults

`func NewMetricMemoryWithDefaults() *MetricMemory`

NewMetricMemoryWithDefaults instantiates a new MetricMemory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceName

`func (o *MetricMemory) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *MetricMemory) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *MetricMemory) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetData

`func (o *MetricMemory) GetData() []MetricMemoryDatapoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetricMemory) GetDataOk() (*[]MetricMemoryDatapoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetricMemory) SetData(v []MetricMemoryDatapoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


