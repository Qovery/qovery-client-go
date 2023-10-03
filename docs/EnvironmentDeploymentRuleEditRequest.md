# EnvironmentDeploymentRuleEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnDemandPreview** | Pointer to **bool** |  | [optional] [default to false]
**AutoPreview** | Pointer to **bool** |  | [optional] [default to false]
**AutoStop** | Pointer to **bool** |  | [optional] [default to false]
**Timezone** | **string** |  | 
**StartTime** | **time.Time** |  | 
**StopTime** | **time.Time** |  | 
**Weekdays** | [**[]WeekdayEnum**](WeekdayEnum.md) |  | 

## Methods

### NewEnvironmentDeploymentRuleEditRequest

`func NewEnvironmentDeploymentRuleEditRequest(timezone string, startTime time.Time, stopTime time.Time, weekdays []WeekdayEnum, ) *EnvironmentDeploymentRuleEditRequest`

NewEnvironmentDeploymentRuleEditRequest instantiates a new EnvironmentDeploymentRuleEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDeploymentRuleEditRequestWithDefaults

`func NewEnvironmentDeploymentRuleEditRequestWithDefaults() *EnvironmentDeploymentRuleEditRequest`

NewEnvironmentDeploymentRuleEditRequestWithDefaults instantiates a new EnvironmentDeploymentRuleEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnDemandPreview

`func (o *EnvironmentDeploymentRuleEditRequest) GetOnDemandPreview() bool`

GetOnDemandPreview returns the OnDemandPreview field if non-nil, zero value otherwise.

### GetOnDemandPreviewOk

`func (o *EnvironmentDeploymentRuleEditRequest) GetOnDemandPreviewOk() (*bool, bool)`

GetOnDemandPreviewOk returns a tuple with the OnDemandPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandPreview

`func (o *EnvironmentDeploymentRuleEditRequest) SetOnDemandPreview(v bool)`

SetOnDemandPreview sets OnDemandPreview field to given value.

### HasOnDemandPreview

`func (o *EnvironmentDeploymentRuleEditRequest) HasOnDemandPreview() bool`

HasOnDemandPreview returns a boolean if a field has been set.

### GetAutoPreview

`func (o *EnvironmentDeploymentRuleEditRequest) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *EnvironmentDeploymentRuleEditRequest) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *EnvironmentDeploymentRuleEditRequest) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.

### HasAutoPreview

`func (o *EnvironmentDeploymentRuleEditRequest) HasAutoPreview() bool`

HasAutoPreview returns a boolean if a field has been set.

### GetAutoStop

`func (o *EnvironmentDeploymentRuleEditRequest) GetAutoStop() bool`

GetAutoStop returns the AutoStop field if non-nil, zero value otherwise.

### GetAutoStopOk

`func (o *EnvironmentDeploymentRuleEditRequest) GetAutoStopOk() (*bool, bool)`

GetAutoStopOk returns a tuple with the AutoStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStop

`func (o *EnvironmentDeploymentRuleEditRequest) SetAutoStop(v bool)`

SetAutoStop sets AutoStop field to given value.

### HasAutoStop

`func (o *EnvironmentDeploymentRuleEditRequest) HasAutoStop() bool`

HasAutoStop returns a boolean if a field has been set.

### GetTimezone

`func (o *EnvironmentDeploymentRuleEditRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *EnvironmentDeploymentRuleEditRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *EnvironmentDeploymentRuleEditRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStartTime

`func (o *EnvironmentDeploymentRuleEditRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EnvironmentDeploymentRuleEditRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EnvironmentDeploymentRuleEditRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetStopTime

`func (o *EnvironmentDeploymentRuleEditRequest) GetStopTime() time.Time`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *EnvironmentDeploymentRuleEditRequest) GetStopTimeOk() (*time.Time, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *EnvironmentDeploymentRuleEditRequest) SetStopTime(v time.Time)`

SetStopTime sets StopTime field to given value.


### GetWeekdays

`func (o *EnvironmentDeploymentRuleEditRequest) GetWeekdays() []WeekdayEnum`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *EnvironmentDeploymentRuleEditRequest) GetWeekdaysOk() (*[]WeekdayEnum, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *EnvironmentDeploymentRuleEditRequest) SetWeekdays(v []WeekdayEnum)`

SetWeekdays sets Weekdays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


