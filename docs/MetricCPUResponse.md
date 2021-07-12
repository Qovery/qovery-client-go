# MetricCPUResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceName** | **string** |  | 
**Data** | [**[]MetricCPUDatapointResponse**](MetricCPUDatapointResponse.md) |  | 

## Methods

### NewMetricCPUResponse

`func NewMetricCPUResponse(instanceName string, data []MetricCPUDatapointResponse, ) *MetricCPUResponse`

NewMetricCPUResponse instantiates a new MetricCPUResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricCPUResponseWithDefaults

`func NewMetricCPUResponseWithDefaults() *MetricCPUResponse`

NewMetricCPUResponseWithDefaults instantiates a new MetricCPUResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceName

`func (o *MetricCPUResponse) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *MetricCPUResponse) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *MetricCPUResponse) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetData

`func (o *MetricCPUResponse) GetData() []MetricCPUDatapointResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetricCPUResponse) GetDataOk() (*[]MetricCPUDatapointResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetricCPUResponse) SetData(v []MetricCPUDatapointResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


