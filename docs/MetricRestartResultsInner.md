# MetricRestartResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **time.Time** |  | 
**Message** | **string** |  | 

## Methods

### NewMetricRestartResultsInner

`func NewMetricRestartResultsInner(datetime time.Time, message string, ) *MetricRestartResultsInner`

NewMetricRestartResultsInner instantiates a new MetricRestartResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricRestartResultsInnerWithDefaults

`func NewMetricRestartResultsInnerWithDefaults() *MetricRestartResultsInner`

NewMetricRestartResultsInnerWithDefaults instantiates a new MetricRestartResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *MetricRestartResultsInner) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *MetricRestartResultsInner) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *MetricRestartResultsInner) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.


### GetMessage

`func (o *MetricRestartResultsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MetricRestartResultsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MetricRestartResultsInner) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


