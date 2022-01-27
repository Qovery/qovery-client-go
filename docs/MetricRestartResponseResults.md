# MetricRestartResponseResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **time.Time** |  | 
**Message** | **string** |  | 

## Methods

### NewMetricRestartResponseResults

`func NewMetricRestartResponseResults(datetime time.Time, message string, ) *MetricRestartResponseResults`

NewMetricRestartResponseResults instantiates a new MetricRestartResponseResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricRestartResponseResultsWithDefaults

`func NewMetricRestartResponseResultsWithDefaults() *MetricRestartResponseResults`

NewMetricRestartResponseResultsWithDefaults instantiates a new MetricRestartResponseResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *MetricRestartResponseResults) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *MetricRestartResponseResults) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *MetricRestartResponseResults) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.


### GetMessage

`func (o *MetricRestartResponseResults) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MetricRestartResponseResults) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MetricRestartResponseResults) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


