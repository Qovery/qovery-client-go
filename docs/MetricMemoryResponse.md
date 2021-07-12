# MetricMemoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceName** | **string** |  | 
**Data** | [**[]MetricMemoryDatapointResponse**](MetricMemoryDatapointResponse.md) |  | 

## Methods

### NewMetricMemoryResponse

`func NewMetricMemoryResponse(instanceName string, data []MetricMemoryDatapointResponse, ) *MetricMemoryResponse`

NewMetricMemoryResponse instantiates a new MetricMemoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricMemoryResponseWithDefaults

`func NewMetricMemoryResponseWithDefaults() *MetricMemoryResponse`

NewMetricMemoryResponseWithDefaults instantiates a new MetricMemoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceName

`func (o *MetricMemoryResponse) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *MetricMemoryResponse) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *MetricMemoryResponse) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetData

`func (o *MetricMemoryResponse) GetData() []MetricMemoryDatapointResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetricMemoryResponse) GetDataOk() (*[]MetricMemoryDatapointResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetricMemoryResponse) SetData(v []MetricMemoryDatapointResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


