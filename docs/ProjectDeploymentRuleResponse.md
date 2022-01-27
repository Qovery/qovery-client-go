# ProjectDeploymentRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | **string** |  | 
**Cluster** | **string** |  | 
**AutoDeploy** | Pointer to **bool** |  | [optional] [default to true]
**AutoStop** | **bool** |  | [default to false]
**Timezone** | Pointer to **string** | specify value only if auto_stop &#x3D; false | [optional] [default to "Europe/London"]
**StartTime** | Pointer to **NullableTime** | specify value only if auto_stop &#x3D; false | [optional] 
**StopTime** | Pointer to **NullableTime** | specify value only if auto_stop &#x3D; false | [optional] 
**Weekdays** | Pointer to **[]string** | specify value only if auto_stop &#x3D; false | [optional] 

## Methods

### NewProjectDeploymentRuleResponse

`func NewProjectDeploymentRuleResponse(id string, createdAt time.Time, name string, mode string, cluster string, autoStop bool, ) *ProjectDeploymentRuleResponse`

NewProjectDeploymentRuleResponse instantiates a new ProjectDeploymentRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDeploymentRuleResponseWithDefaults

`func NewProjectDeploymentRuleResponseWithDefaults() *ProjectDeploymentRuleResponse`

NewProjectDeploymentRuleResponseWithDefaults instantiates a new ProjectDeploymentRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectDeploymentRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectDeploymentRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectDeploymentRuleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ProjectDeploymentRuleResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectDeploymentRuleResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectDeploymentRuleResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ProjectDeploymentRuleResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectDeploymentRuleResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectDeploymentRuleResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectDeploymentRuleResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *ProjectDeploymentRuleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectDeploymentRuleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectDeploymentRuleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProjectDeploymentRuleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectDeploymentRuleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectDeploymentRuleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectDeploymentRuleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *ProjectDeploymentRuleResponse) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ProjectDeploymentRuleResponse) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ProjectDeploymentRuleResponse) SetMode(v string)`

SetMode sets Mode field to given value.


### GetCluster

`func (o *ProjectDeploymentRuleResponse) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ProjectDeploymentRuleResponse) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ProjectDeploymentRuleResponse) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetAutoDeploy

`func (o *ProjectDeploymentRuleResponse) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *ProjectDeploymentRuleResponse) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *ProjectDeploymentRuleResponse) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *ProjectDeploymentRuleResponse) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetAutoStop

`func (o *ProjectDeploymentRuleResponse) GetAutoStop() bool`

GetAutoStop returns the AutoStop field if non-nil, zero value otherwise.

### GetAutoStopOk

`func (o *ProjectDeploymentRuleResponse) GetAutoStopOk() (*bool, bool)`

GetAutoStopOk returns a tuple with the AutoStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStop

`func (o *ProjectDeploymentRuleResponse) SetAutoStop(v bool)`

SetAutoStop sets AutoStop field to given value.


### GetTimezone

`func (o *ProjectDeploymentRuleResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ProjectDeploymentRuleResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ProjectDeploymentRuleResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ProjectDeploymentRuleResponse) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetStartTime

`func (o *ProjectDeploymentRuleResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProjectDeploymentRuleResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ProjectDeploymentRuleResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ProjectDeploymentRuleResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ProjectDeploymentRuleResponse) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *ProjectDeploymentRuleResponse) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetStopTime

`func (o *ProjectDeploymentRuleResponse) GetStopTime() time.Time`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *ProjectDeploymentRuleResponse) GetStopTimeOk() (*time.Time, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *ProjectDeploymentRuleResponse) SetStopTime(v time.Time)`

SetStopTime sets StopTime field to given value.

### HasStopTime

`func (o *ProjectDeploymentRuleResponse) HasStopTime() bool`

HasStopTime returns a boolean if a field has been set.

### SetStopTimeNil

`func (o *ProjectDeploymentRuleResponse) SetStopTimeNil(b bool)`

 SetStopTimeNil sets the value for StopTime to be an explicit nil

### UnsetStopTime
`func (o *ProjectDeploymentRuleResponse) UnsetStopTime()`

UnsetStopTime ensures that no value is present for StopTime, not even an explicit nil
### GetWeekdays

`func (o *ProjectDeploymentRuleResponse) GetWeekdays() []string`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *ProjectDeploymentRuleResponse) GetWeekdaysOk() (*[]string, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *ProjectDeploymentRuleResponse) SetWeekdays(v []string)`

SetWeekdays sets Weekdays field to given value.

### HasWeekdays

`func (o *ProjectDeploymentRuleResponse) HasWeekdays() bool`

HasWeekdays returns a boolean if a field has been set.

### SetWeekdaysNil

`func (o *ProjectDeploymentRuleResponse) SetWeekdaysNil(b bool)`

 SetWeekdaysNil sets the value for Weekdays to be an explicit nil

### UnsetWeekdays
`func (o *ProjectDeploymentRuleResponse) UnsetWeekdays()`

UnsetWeekdays ensures that no value is present for Weekdays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


