# MetricRestartResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **time.Time** |  | 
**Message** | **string** |  | 

## Methods

### NewMetricRestartResults

`func NewMetricRestartResults(datetime time.Time, message string, ) *MetricRestartResults`

NewMetricRestartResults instantiates a new MetricRestartResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricRestartResultsWithDefaults

`func NewMetricRestartResultsWithDefaults() *MetricRestartResults`

NewMetricRestartResultsWithDefaults instantiates a new MetricRestartResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *MetricRestartResults) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *MetricRestartResults) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *MetricRestartResults) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.


### GetMessage

`func (o *MetricRestartResults) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MetricRestartResults) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MetricRestartResults) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


