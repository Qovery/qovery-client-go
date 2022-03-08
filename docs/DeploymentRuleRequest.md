# DeploymentRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | [**EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | 
**Cluster** | **string** |  | 
**AutoDeploy** | Pointer to **bool** |  | [optional] [default to true]
**AutoStop** | **bool** |  | [default to false]
**Timezone** | Pointer to **string** | specify value only if auto_stop &#x3D; false | [optional] [default to "Europe/London"]
**StartTime** | Pointer to **NullableTime** | specify value only if auto_stop &#x3D; false | [optional] 
**StopTime** | Pointer to **NullableTime** | specify value only if auto_stop &#x3D; false | [optional] 
**Weekdays** | Pointer to [**[]WeekdayEnum**](WeekdayEnum.md) | specify value only if auto_stop &#x3D; false | [optional] 

## Methods

### NewDeploymentRuleRequest

`func NewDeploymentRuleRequest(name string, mode EnvironmentModeEnum, cluster string, autoStop bool, ) *DeploymentRuleRequest`

NewDeploymentRuleRequest instantiates a new DeploymentRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentRuleRequestWithDefaults

`func NewDeploymentRuleRequestWithDefaults() *DeploymentRuleRequest`

NewDeploymentRuleRequestWithDefaults instantiates a new DeploymentRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DeploymentRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeploymentRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *DeploymentRuleRequest) GetMode() EnvironmentModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DeploymentRuleRequest) GetModeOk() (*EnvironmentModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DeploymentRuleRequest) SetMode(v EnvironmentModeEnum)`

SetMode sets Mode field to given value.


### GetCluster

`func (o *DeploymentRuleRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DeploymentRuleRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DeploymentRuleRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetAutoDeploy

`func (o *DeploymentRuleRequest) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *DeploymentRuleRequest) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *DeploymentRuleRequest) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *DeploymentRuleRequest) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetAutoStop

`func (o *DeploymentRuleRequest) GetAutoStop() bool`

GetAutoStop returns the AutoStop field if non-nil, zero value otherwise.

### GetAutoStopOk

`func (o *DeploymentRuleRequest) GetAutoStopOk() (*bool, bool)`

GetAutoStopOk returns a tuple with the AutoStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStop

`func (o *DeploymentRuleRequest) SetAutoStop(v bool)`

SetAutoStop sets AutoStop field to given value.


### GetTimezone

`func (o *DeploymentRuleRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *DeploymentRuleRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *DeploymentRuleRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *DeploymentRuleRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetStartTime

`func (o *DeploymentRuleRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DeploymentRuleRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DeploymentRuleRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DeploymentRuleRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *DeploymentRuleRequest) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *DeploymentRuleRequest) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetStopTime

`func (o *DeploymentRuleRequest) GetStopTime() time.Time`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *DeploymentRuleRequest) GetStopTimeOk() (*time.Time, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *DeploymentRuleRequest) SetStopTime(v time.Time)`

SetStopTime sets StopTime field to given value.

### HasStopTime

`func (o *DeploymentRuleRequest) HasStopTime() bool`

HasStopTime returns a boolean if a field has been set.

### SetStopTimeNil

`func (o *DeploymentRuleRequest) SetStopTimeNil(b bool)`

 SetStopTimeNil sets the value for StopTime to be an explicit nil

### UnsetStopTime
`func (o *DeploymentRuleRequest) UnsetStopTime()`

UnsetStopTime ensures that no value is present for StopTime, not even an explicit nil
### GetWeekdays

`func (o *DeploymentRuleRequest) GetWeekdays() []WeekdayEnum`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *DeploymentRuleRequest) GetWeekdaysOk() (*[]WeekdayEnum, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *DeploymentRuleRequest) SetWeekdays(v []WeekdayEnum)`

SetWeekdays sets Weekdays field to given value.

### HasWeekdays

`func (o *DeploymentRuleRequest) HasWeekdays() bool`

HasWeekdays returns a boolean if a field has been set.

### SetWeekdaysNil

`func (o *DeploymentRuleRequest) SetWeekdaysNil(b bool)`

 SetWeekdaysNil sets the value for Weekdays to be an explicit nil

### UnsetWeekdays
`func (o *DeploymentRuleRequest) UnsetWeekdays()`

UnsetWeekdays ensures that no value is present for Weekdays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


