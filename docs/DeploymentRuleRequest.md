# DeploymentRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewDeploymentRuleRequest

`func NewDeploymentRuleRequest(name string, mode string, cluster string, alwaysUp bool, ) *DeploymentRuleRequest`

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

`func (o *DeploymentRuleRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DeploymentRuleRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DeploymentRuleRequest) SetMode(v string)`

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

### GetAlwaysUp

`func (o *DeploymentRuleRequest) GetAlwaysUp() bool`

GetAlwaysUp returns the AlwaysUp field if non-nil, zero value otherwise.

### GetAlwaysUpOk

`func (o *DeploymentRuleRequest) GetAlwaysUpOk() (*bool, bool)`

GetAlwaysUpOk returns a tuple with the AlwaysUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysUp

`func (o *DeploymentRuleRequest) SetAlwaysUp(v bool)`

SetAlwaysUp sets AlwaysUp field to given value.


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
### GetWeekday

`func (o *DeploymentRuleRequest) GetWeekday() string`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *DeploymentRuleRequest) GetWeekdayOk() (*string, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *DeploymentRuleRequest) SetWeekday(v string)`

SetWeekday sets Weekday field to given value.

### HasWeekday

`func (o *DeploymentRuleRequest) HasWeekday() bool`

HasWeekday returns a boolean if a field has been set.

### SetWeekdayNil

`func (o *DeploymentRuleRequest) SetWeekdayNil(b bool)`

 SetWeekdayNil sets the value for Weekday to be an explicit nil

### UnsetWeekday
`func (o *DeploymentRuleRequest) UnsetWeekday()`

UnsetWeekday ensures that no value is present for Weekday, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


