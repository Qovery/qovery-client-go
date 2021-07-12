# ProjectDeploymentRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentTarget** | **string** | specify here a regex based on environment name that will target new environments after their creation | 
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | **string** |  | 
**Cluster** | **string** |  | 
**AutoDeploy** | Pointer to **bool** |  | [optional] [default to true]
**AlwaysUp** | **bool** |  | [default to false]
**StartTime** | Pointer to **NullableTime** | specify value only if always_up &#x3D; false | [optional] 
**StopTime** | Pointer to **NullableTime** | specify value only if always_up &#x3D; false | [optional] 
**Weekday** | Pointer to **NullableString** | specify value only if always_up &#x3D; false | [optional] 

## Methods

### NewProjectDeploymentRuleRequest

`func NewProjectDeploymentRuleRequest(environmentTarget string, name string, mode string, cluster string, alwaysUp bool, ) *ProjectDeploymentRuleRequest`

NewProjectDeploymentRuleRequest instantiates a new ProjectDeploymentRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDeploymentRuleRequestWithDefaults

`func NewProjectDeploymentRuleRequestWithDefaults() *ProjectDeploymentRuleRequest`

NewProjectDeploymentRuleRequestWithDefaults instantiates a new ProjectDeploymentRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentTarget

`func (o *ProjectDeploymentRuleRequest) GetEnvironmentTarget() string`

GetEnvironmentTarget returns the EnvironmentTarget field if non-nil, zero value otherwise.

### GetEnvironmentTargetOk

`func (o *ProjectDeploymentRuleRequest) GetEnvironmentTargetOk() (*string, bool)`

GetEnvironmentTargetOk returns a tuple with the EnvironmentTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentTarget

`func (o *ProjectDeploymentRuleRequest) SetEnvironmentTarget(v string)`

SetEnvironmentTarget sets EnvironmentTarget field to given value.


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

### GetMode

`func (o *ProjectDeploymentRuleRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ProjectDeploymentRuleRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ProjectDeploymentRuleRequest) SetMode(v string)`

SetMode sets Mode field to given value.


### GetCluster

`func (o *ProjectDeploymentRuleRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ProjectDeploymentRuleRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ProjectDeploymentRuleRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetAutoDeploy

`func (o *ProjectDeploymentRuleRequest) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *ProjectDeploymentRuleRequest) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *ProjectDeploymentRuleRequest) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *ProjectDeploymentRuleRequest) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetAlwaysUp

`func (o *ProjectDeploymentRuleRequest) GetAlwaysUp() bool`

GetAlwaysUp returns the AlwaysUp field if non-nil, zero value otherwise.

### GetAlwaysUpOk

`func (o *ProjectDeploymentRuleRequest) GetAlwaysUpOk() (*bool, bool)`

GetAlwaysUpOk returns a tuple with the AlwaysUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysUp

`func (o *ProjectDeploymentRuleRequest) SetAlwaysUp(v bool)`

SetAlwaysUp sets AlwaysUp field to given value.


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

### HasStartTime

`func (o *ProjectDeploymentRuleRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ProjectDeploymentRuleRequest) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *ProjectDeploymentRuleRequest) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
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

### HasStopTime

`func (o *ProjectDeploymentRuleRequest) HasStopTime() bool`

HasStopTime returns a boolean if a field has been set.

### SetStopTimeNil

`func (o *ProjectDeploymentRuleRequest) SetStopTimeNil(b bool)`

 SetStopTimeNil sets the value for StopTime to be an explicit nil

### UnsetStopTime
`func (o *ProjectDeploymentRuleRequest) UnsetStopTime()`

UnsetStopTime ensures that no value is present for StopTime, not even an explicit nil
### GetWeekday

`func (o *ProjectDeploymentRuleRequest) GetWeekday() string`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *ProjectDeploymentRuleRequest) GetWeekdayOk() (*string, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *ProjectDeploymentRuleRequest) SetWeekday(v string)`

SetWeekday sets Weekday field to given value.

### HasWeekday

`func (o *ProjectDeploymentRuleRequest) HasWeekday() bool`

HasWeekday returns a boolean if a field has been set.

### SetWeekdayNil

`func (o *ProjectDeploymentRuleRequest) SetWeekdayNil(b bool)`

 SetWeekdayNil sets the value for Weekday to be an explicit nil

### UnsetWeekday
`func (o *ProjectDeploymentRuleRequest) UnsetWeekday()`

UnsetWeekday ensures that no value is present for Weekday, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


