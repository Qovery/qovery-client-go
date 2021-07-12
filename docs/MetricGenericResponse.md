# MetricGenericResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceName** | **string** |  | 
**Data** | [**[]MetricGenericDatapointResponse**](MetricGenericDatapointResponse.md) |  | 

## Methods

### NewMetricGenericResponse

`func NewMetricGenericResponse(instanceName string, data []MetricGenericDatapointResponse, ) *MetricGenericResponse`

NewMetricGenericResponse instantiates a new MetricGenericResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricGenericResponseWithDefaults

`func NewMetricGenericResponseWithDefaults() *MetricGenericResponse`

NewMetricGenericResponseWithDefaults instantiates a new MetricGenericResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceName

`func (o *MetricGenericResponse) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *MetricGenericResponse) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *MetricGenericResponse) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.


### GetData

`func (o *MetricGenericResponse) GetData() []MetricGenericDatapointResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetricGenericResponse) GetDataOk() (*[]MetricGenericDatapointResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetricGenericResponse) SetData(v []MetricGenericDatapointResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


