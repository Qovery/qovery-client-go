# MetricGeneric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceName** | **string** |  | 
**Data** | [**[]MetricGenericDatapoint**](MetricGenericDatapoint.md) |  | 

## Methods

### NewMetricGeneric

`func NewMetricGeneric(instanceName string, data []MetricGenericDatapoint, ) *MetricGeneric`

NewMetricGeneric instantiates a new MetricGeneric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricGenericWithDefaults

`func NewMetricGenericWithDefaults() *MetricGeneric`

NewMetricGenericWithDefaults instantiates a new MetricGeneric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceName

`func (o *MetricGeneric) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *MetricGeneric) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *MetricGeneric) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetData

`func (o *MetricGeneric) GetData() []MetricGenericDatapoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetricGeneric) GetDataOk() (*[]MetricGenericDatapoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetricGeneric) SetData(v []MetricGenericDatapoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


