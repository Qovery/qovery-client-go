# MetricCPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceName** | **string** |  | 
**Data** | [**[]MetricCPUDatapoint**](MetricCPUDatapoint.md) |  | 

## Methods

### NewMetricCPU

`func NewMetricCPU(instanceName string, data []MetricCPUDatapoint, ) *MetricCPU`

NewMetricCPU instantiates a new MetricCPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricCPUWithDefaults

`func NewMetricCPUWithDefaults() *MetricCPU`

NewMetricCPUWithDefaults instantiates a new MetricCPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceName

`func (o *MetricCPU) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *MetricCPU) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *MetricCPU) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetData

`func (o *MetricCPU) GetData() []MetricCPUDatapoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetricCPU) GetDataOk() (*[]MetricCPUDatapoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetricCPU) SetData(v []MetricCPUDatapoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


