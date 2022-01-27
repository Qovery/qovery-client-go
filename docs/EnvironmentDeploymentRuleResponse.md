# EnvironmentDeploymentRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeploy** | Pointer to **bool** |  | [optional] [default to true]
**AutoStop** | Pointer to **bool** |  | [optional] [default to false]
**Timezone** | Pointer to **string** | specify value only if auto_stop &#x3D; false | [optional] [default to "Europe/London"]
**StartTime** | Pointer to **NullableTime** | specify value only if auto_stop &#x3D; false | [optional] 
**StopTime** | Pointer to **NullableTime** | specify value only if auto_stop &#x3D; false | [optional] 
**Weekdays** | Pointer to **[]string** | specify value only if auto_stop &#x3D; false | [optional] 
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewEnvironmentDeploymentRuleResponse

`func NewEnvironmentDeploymentRuleResponse(id string, createdAt time.Time, ) *EnvironmentDeploymentRuleResponse`

NewEnvironmentDeploymentRuleResponse instantiates a new EnvironmentDeploymentRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDeploymentRuleResponseWithDefaults

`func NewEnvironmentDeploymentRuleResponseWithDefaults() *EnvironmentDeploymentRuleResponse`

NewEnvironmentDeploymentRuleResponseWithDefaults instantiates a new EnvironmentDeploymentRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDeploy

`func (o *EnvironmentDeploymentRuleResponse) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *EnvironmentDeploymentRuleResponse) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *EnvironmentDeploymentRuleResponse) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *EnvironmentDeploymentRuleResponse) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetAutoStop

`func (o *EnvironmentDeploymentRuleResponse) GetAutoStop() bool`

GetAutoStop returns the AutoStop field if non-nil, zero value otherwise.

### GetAutoStopOk

`func (o *EnvironmentDeploymentRuleResponse) GetAutoStopOk() (*bool, bool)`

GetAutoStopOk returns a tuple with the AutoStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStop

`func (o *EnvironmentDeploymentRuleResponse) SetAutoStop(v bool)`

SetAutoStop sets AutoStop field to given value.

### HasAutoStop

`func (o *EnvironmentDeploymentRuleResponse) HasAutoStop() bool`

HasAutoStop returns a boolean if a field has been set.

### GetTimezone

`func (o *EnvironmentDeploymentRuleResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *EnvironmentDeploymentRuleResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *EnvironmentDeploymentRuleResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *EnvironmentDeploymentRuleResponse) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetStartTime

`func (o *EnvironmentDeploymentRuleResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EnvironmentDeploymentRuleResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EnvironmentDeploymentRuleResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *EnvironmentDeploymentRuleResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *EnvironmentDeploymentRuleResponse) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *EnvironmentDeploymentRuleResponse) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetStopTime

`func (o *EnvironmentDeploymentRuleResponse) GetStopTime() time.Time`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *EnvironmentDeploymentRuleResponse) GetStopTimeOk() (*time.Time, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *EnvironmentDeploymentRuleResponse) SetStopTime(v time.Time)`

SetStopTime sets StopTime field to given value.

### HasStopTime

`func (o *EnvironmentDeploymentRuleResponse) HasStopTime() bool`

HasStopTime returns a boolean if a field has been set.

### SetStopTimeNil

`func (o *EnvironmentDeploymentRuleResponse) SetStopTimeNil(b bool)`

 SetStopTimeNil sets the value for StopTime to be an explicit nil

### UnsetStopTime
`func (o *EnvironmentDeploymentRuleResponse) UnsetStopTime()`

UnsetStopTime ensures that no value is present for StopTime, not even an explicit nil
### GetWeekdays

`func (o *EnvironmentDeploymentRuleResponse) GetWeekdays() []string`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *EnvironmentDeploymentRuleResponse) GetWeekdaysOk() (*[]string, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *EnvironmentDeploymentRuleResponse) SetWeekdays(v []string)`

SetWeekdays sets Weekdays field to given value.

### HasWeekdays

`func (o *EnvironmentDeploymentRuleResponse) HasWeekdays() bool`

HasWeekdays returns a boolean if a field has been set.

### SetWeekdaysNil

`func (o *EnvironmentDeploymentRuleResponse) SetWeekdaysNil(b bool)`

 SetWeekdaysNil sets the value for Weekdays to be an explicit nil

### UnsetWeekdays
`func (o *EnvironmentDeploymentRuleResponse) UnsetWeekdays()`

UnsetWeekdays ensures that no value is present for Weekdays, not even an explicit nil
### GetId

`func (o *EnvironmentDeploymentRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentDeploymentRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentDeploymentRuleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EnvironmentDeploymentRuleResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentDeploymentRuleResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentDeploymentRuleResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EnvironmentDeploymentRuleResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentDeploymentRuleResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentDeploymentRuleResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentDeploymentRuleResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


