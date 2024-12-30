# KarpenterNodePoolConsolidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | [default to false]
**Days** | [**[]WeekdayEnum**](WeekdayEnum.md) |  | 
**StartTime** | **string** | The start date of the consolidation. The format should follow ISO-8601 convention: \&quot;PThh:mm\&quot;  | 
**Duration** | **string** | The duration during the consolidation will be active. The format should follow ISO-8601 convention: \&quot;PThhHmmM\&quot;  | 

## Methods

### NewKarpenterNodePoolConsolidation

`func NewKarpenterNodePoolConsolidation(enabled bool, days []WeekdayEnum, startTime string, duration string, ) *KarpenterNodePoolConsolidation`

NewKarpenterNodePoolConsolidation instantiates a new KarpenterNodePoolConsolidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKarpenterNodePoolConsolidationWithDefaults

`func NewKarpenterNodePoolConsolidationWithDefaults() *KarpenterNodePoolConsolidation`

NewKarpenterNodePoolConsolidationWithDefaults instantiates a new KarpenterNodePoolConsolidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *KarpenterNodePoolConsolidation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KarpenterNodePoolConsolidation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KarpenterNodePoolConsolidation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDays

`func (o *KarpenterNodePoolConsolidation) GetDays() []WeekdayEnum`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *KarpenterNodePoolConsolidation) GetDaysOk() (*[]WeekdayEnum, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *KarpenterNodePoolConsolidation) SetDays(v []WeekdayEnum)`

SetDays sets Days field to given value.


### GetStartTime

`func (o *KarpenterNodePoolConsolidation) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *KarpenterNodePoolConsolidation) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *KarpenterNodePoolConsolidation) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetDuration

`func (o *KarpenterNodePoolConsolidation) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *KarpenterNodePoolConsolidation) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *KarpenterNodePoolConsolidation) SetDuration(v string)`

SetDuration sets Duration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


