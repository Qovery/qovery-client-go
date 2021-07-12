# MetricGenericDatapointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | [readonly] 
**Value** | **float32** |  | 

## Methods

### NewMetricGenericDatapointResponse

`func NewMetricGenericDatapointResponse(createdAt time.Time, value float32, ) *MetricGenericDatapointResponse`

NewMetricGenericDatapointResponse instantiates a new MetricGenericDatapointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricGenericDatapointResponseWithDefaults

`func NewMetricGenericDatapointResponseWithDefaults() *MetricGenericDatapointResponse`

NewMetricGenericDatapointResponseWithDefaults instantiates a new MetricGenericDatapointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MetricGenericDatapointResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetricGenericDatapointResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetricGenericDatapointResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetValue

`func (o *MetricGenericDatapointResponse) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricGenericDatapointResponse) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricGenericDatapointResponse) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


