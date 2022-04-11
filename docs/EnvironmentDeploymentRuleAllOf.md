# EnvironmentDeploymentRuleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeploy** | Pointer to **bool** |  | [optional] [default to true]
**AutoStop** | Pointer to **bool** |  | [optional] [default to false]
**AutoDelete** | Pointer to **bool** |  | [optional] [default to false]
**AutoPreview** | Pointer to **bool** |  | [optional] [default to false]
**Timezone** | **string** |  | 
**StartTime** | **time.Time** |  | 
**StopTime** | **time.Time** |  | 
**Weekdays** | [**[]WeekdayEnum**](WeekdayEnum.md) |  | 

## Methods

### NewEnvironmentDeploymentRuleAllOf

`func NewEnvironmentDeploymentRuleAllOf(timezone string, startTime time.Time, stopTime time.Time, weekdays []WeekdayEnum, ) *EnvironmentDeploymentRuleAllOf`

NewEnvironmentDeploymentRuleAllOf instantiates a new EnvironmentDeploymentRuleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDeploymentRuleAllOfWithDefaults

`func NewEnvironmentDeploymentRuleAllOfWithDefaults() *EnvironmentDeploymentRuleAllOf`

NewEnvironmentDeploymentRuleAllOfWithDefaults instantiates a new EnvironmentDeploymentRuleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDeploy

`func (o *EnvironmentDeploymentRuleAllOf) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *EnvironmentDeploymentRuleAllOf) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *EnvironmentDeploymentRuleAllOf) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *EnvironmentDeploymentRuleAllOf) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetAutoStop

`func (o *EnvironmentDeploymentRuleAllOf) GetAutoStop() bool`

GetAutoStop returns the AutoStop field if non-nil, zero value otherwise.

### GetAutoStopOk

`func (o *EnvironmentDeploymentRuleAllOf) GetAutoStopOk() (*bool, bool)`

GetAutoStopOk returns a tuple with the AutoStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStop

`func (o *EnvironmentDeploymentRuleAllOf) SetAutoStop(v bool)`

SetAutoStop sets AutoStop field to given value.

### HasAutoStop

`func (o *EnvironmentDeploymentRuleAllOf) HasAutoStop() bool`

HasAutoStop returns a boolean if a field has been set.

### GetAutoDelete

`func (o *EnvironmentDeploymentRuleAllOf) GetAutoDelete() bool`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *EnvironmentDeploymentRuleAllOf) GetAutoDeleteOk() (*bool, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *EnvironmentDeploymentRuleAllOf) SetAutoDelete(v bool)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *EnvironmentDeploymentRuleAllOf) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetAutoPreview

`func (o *EnvironmentDeploymentRuleAllOf) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *EnvironmentDeploymentRuleAllOf) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *EnvironmentDeploymentRuleAllOf) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.

### HasAutoPreview

`func (o *EnvironmentDeploymentRuleAllOf) HasAutoPreview() bool`

HasAutoPreview returns a boolean if a field has been set.

### GetTimezone

`func (o *EnvironmentDeploymentRuleAllOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *EnvironmentDeploymentRuleAllOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *EnvironmentDeploymentRuleAllOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStartTime

`func (o *EnvironmentDeploymentRuleAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EnvironmentDeploymentRuleAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EnvironmentDeploymentRuleAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetStopTime

`func (o *EnvironmentDeploymentRuleAllOf) GetStopTime() time.Time`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *EnvironmentDeploymentRuleAllOf) GetStopTimeOk() (*time.Time, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *EnvironmentDeploymentRuleAllOf) SetStopTime(v time.Time)`

SetStopTime sets StopTime field to given value.


### GetWeekdays

`func (o *EnvironmentDeploymentRuleAllOf) GetWeekdays() []WeekdayEnum`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *EnvironmentDeploymentRuleAllOf) GetWeekdaysOk() (*[]WeekdayEnum, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *EnvironmentDeploymentRuleAllOf) SetWeekdays(v []WeekdayEnum)`

SetWeekdays sets Weekdays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


