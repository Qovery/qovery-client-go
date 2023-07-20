# ProjectDeploymentRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Mode** | [**EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | 
**ClusterId** | **string** |  | 
**AutoDeploy** | Pointer to **bool** |  | [optional] [default to false]
**AutoStop** | Pointer to **bool** |  | [optional] [default to false]
**Timezone** | **string** |  | 
**StartTime** | **time.Time** |  | 
**StopTime** | **time.Time** |  | 
**Weekdays** | [**[]WeekdayEnum**](WeekdayEnum.md) |  | 
**Wildcard** | **string** | wildcard pattern composed of &#39;?&#39; and/or &#39;*&#39; used to target new created environments | [default to ""]
**PriorityIndex** | Pointer to **int32** | used to select the first deployment rule to match new created environments | [optional] 

## Methods

### NewProjectDeploymentRule

`func NewProjectDeploymentRule(id string, createdAt time.Time, name string, mode EnvironmentModeEnum, clusterId string, timezone string, startTime time.Time, stopTime time.Time, weekdays []WeekdayEnum, wildcard string, ) *ProjectDeploymentRule`

NewProjectDeploymentRule instantiates a new ProjectDeploymentRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDeploymentRuleWithDefaults

`func NewProjectDeploymentRuleWithDefaults() *ProjectDeploymentRule`

NewProjectDeploymentRuleWithDefaults instantiates a new ProjectDeploymentRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectDeploymentRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectDeploymentRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectDeploymentRule) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ProjectDeploymentRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectDeploymentRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectDeploymentRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ProjectDeploymentRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectDeploymentRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectDeploymentRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectDeploymentRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *ProjectDeploymentRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectDeploymentRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectDeploymentRule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProjectDeploymentRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectDeploymentRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectDeploymentRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectDeploymentRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProjectDeploymentRule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProjectDeploymentRule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMode

`func (o *ProjectDeploymentRule) GetMode() EnvironmentModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ProjectDeploymentRule) GetModeOk() (*EnvironmentModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ProjectDeploymentRule) SetMode(v EnvironmentModeEnum)`

SetMode sets Mode field to given value.


### GetClusterId

`func (o *ProjectDeploymentRule) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ProjectDeploymentRule) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ProjectDeploymentRule) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetAutoDeploy

`func (o *ProjectDeploymentRule) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *ProjectDeploymentRule) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *ProjectDeploymentRule) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *ProjectDeploymentRule) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetAutoStop

`func (o *ProjectDeploymentRule) GetAutoStop() bool`

GetAutoStop returns the AutoStop field if non-nil, zero value otherwise.

### GetAutoStopOk

`func (o *ProjectDeploymentRule) GetAutoStopOk() (*bool, bool)`

GetAutoStopOk returns a tuple with the AutoStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStop

`func (o *ProjectDeploymentRule) SetAutoStop(v bool)`

SetAutoStop sets AutoStop field to given value.

### HasAutoStop

`func (o *ProjectDeploymentRule) HasAutoStop() bool`

HasAutoStop returns a boolean if a field has been set.

### GetTimezone

`func (o *ProjectDeploymentRule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ProjectDeploymentRule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ProjectDeploymentRule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStartTime

`func (o *ProjectDeploymentRule) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProjectDeploymentRule) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ProjectDeploymentRule) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetStopTime

`func (o *ProjectDeploymentRule) GetStopTime() time.Time`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *ProjectDeploymentRule) GetStopTimeOk() (*time.Time, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *ProjectDeploymentRule) SetStopTime(v time.Time)`

SetStopTime sets StopTime field to given value.


### GetWeekdays

`func (o *ProjectDeploymentRule) GetWeekdays() []WeekdayEnum`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *ProjectDeploymentRule) GetWeekdaysOk() (*[]WeekdayEnum, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *ProjectDeploymentRule) SetWeekdays(v []WeekdayEnum)`

SetWeekdays sets Weekdays field to given value.


### GetWildcard

`func (o *ProjectDeploymentRule) GetWildcard() string`

GetWildcard returns the Wildcard field if non-nil, zero value otherwise.

### GetWildcardOk

`func (o *ProjectDeploymentRule) GetWildcardOk() (*string, bool)`

GetWildcardOk returns a tuple with the Wildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcard

`func (o *ProjectDeploymentRule) SetWildcard(v string)`

SetWildcard sets Wildcard field to given value.


### GetPriorityIndex

`func (o *ProjectDeploymentRule) GetPriorityIndex() int32`

GetPriorityIndex returns the PriorityIndex field if non-nil, zero value otherwise.

### GetPriorityIndexOk

`func (o *ProjectDeploymentRule) GetPriorityIndexOk() (*int32, bool)`

GetPriorityIndexOk returns a tuple with the PriorityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityIndex

`func (o *ProjectDeploymentRule) SetPriorityIndex(v int32)`

SetPriorityIndex sets PriorityIndex field to given value.

### HasPriorityIndex

`func (o *ProjectDeploymentRule) HasPriorityIndex() bool`

HasPriorityIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


