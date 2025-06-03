# ProjectDeploymentRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Mode** | [**EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | 
**ClusterId** | **string** |  | 
**AutoStop** | Pointer to **bool** |  | [optional] [default to false]
**Timezone** | **string** |  | 
**StartTime** | **time.Time** |  | 
**StopTime** | **time.Time** |  | 
**Weekdays** | [**[]WeekdayEnum**](WeekdayEnum.md) |  | 
**Wildcard** | **string** | wildcard pattern composed of &#39;?&#39; and/or &#39;*&#39; used to target new created environments | [default to ""]

## Methods

### NewProjectDeploymentRuleRequest

`func NewProjectDeploymentRuleRequest(name string, mode EnvironmentModeEnum, clusterId string, timezone string, startTime time.Time, stopTime time.Time, weekdays []WeekdayEnum, wildcard string, ) *ProjectDeploymentRuleRequest`

NewProjectDeploymentRuleRequest instantiates a new ProjectDeploymentRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDeploymentRuleRequestWithDefaults

`func NewProjectDeploymentRuleRequestWithDefaults() *ProjectDeploymentRuleRequest`

NewProjectDeploymentRuleRequestWithDefaults instantiates a new ProjectDeploymentRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectDeploymentRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectDeploymentRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectDeploymentRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProjectDeploymentRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectDeploymentRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectDeploymentRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectDeploymentRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProjectDeploymentRuleRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProjectDeploymentRuleRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMode

`func (o *ProjectDeploymentRuleRequest) GetMode() EnvironmentModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ProjectDeploymentRuleRequest) GetModeOk() (*EnvironmentModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ProjectDeploymentRuleRequest) SetMode(v EnvironmentModeEnum)`

SetMode sets Mode field to given value.


### GetClusterId

`func (o *ProjectDeploymentRuleRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ProjectDeploymentRuleRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ProjectDeploymentRuleRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetAutoStop

`func (o *ProjectDeploymentRuleRequest) GetAutoStop() bool`

GetAutoStop returns the AutoStop field if non-nil, zero value otherwise.

### GetAutoStopOk

`func (o *ProjectDeploymentRuleRequest) GetAutoStopOk() (*bool, bool)`

GetAutoStopOk returns a tuple with the AutoStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStop

`func (o *ProjectDeploymentRuleRequest) SetAutoStop(v bool)`

SetAutoStop sets AutoStop field to given value.

### HasAutoStop

`func (o *ProjectDeploymentRuleRequest) HasAutoStop() bool`

HasAutoStop returns a boolean if a field has been set.

### GetTimezone

`func (o *ProjectDeploymentRuleRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ProjectDeploymentRuleRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ProjectDeploymentRuleRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStartTime

`func (o *ProjectDeploymentRuleRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProjectDeploymentRuleRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ProjectDeploymentRuleRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetStopTime

`func (o *ProjectDeploymentRuleRequest) GetStopTime() time.Time`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *ProjectDeploymentRuleRequest) GetStopTimeOk() (*time.Time, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *ProjectDeploymentRuleRequest) SetStopTime(v time.Time)`

SetStopTime sets StopTime field to given value.


### GetWeekdays

`func (o *ProjectDeploymentRuleRequest) GetWeekdays() []WeekdayEnum`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *ProjectDeploymentRuleRequest) GetWeekdaysOk() (*[]WeekdayEnum, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *ProjectDeploymentRuleRequest) SetWeekdays(v []WeekdayEnum)`

SetWeekdays sets Weekdays field to given value.


### GetWildcard

`func (o *ProjectDeploymentRuleRequest) GetWildcard() string`

GetWildcard returns the Wildcard field if non-nil, zero value otherwise.

### GetWildcardOk

`func (o *ProjectDeploymentRuleRequest) GetWildcardOk() (*string, bool)`

GetWildcardOk returns a tuple with the Wildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcard

`func (o *ProjectDeploymentRuleRequest) SetWildcard(v string)`

SetWildcard sets Wildcard field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


