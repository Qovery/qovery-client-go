# EnvironmentDeploymentRuleEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeploy** | Pointer to **bool** |  | [optional] [default to true]
**AlwaysUp** | Pointer to **bool** |  | [optional] [default to false]
**StartTime** | Pointer to **NullableTime** | specify value only if always_up &#x3D; false | [optional] 
**StopTime** | Pointer to **NullableTime** | specify value only if always_up &#x3D; false | [optional] 
**Weekday** | Pointer to **NullableString** | specify value only if always_up &#x3D; false | [optional] 

## Methods

### NewEnvironmentDeploymentRuleEditRequest

`func NewEnvironmentDeploymentRuleEditRequest() *EnvironmentDeploymentRuleEditRequest`

NewEnvironmentDeploymentRuleEditRequest instantiates a new EnvironmentDeploymentRuleEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDeploymentRuleEditRequestWithDefaults

`func NewEnvironmentDeploymentRuleEditRequestWithDefaults() *EnvironmentDeploymentRuleEditRequest`

NewEnvironmentDeploymentRuleEditRequestWithDefaults instantiates a new EnvironmentDeploymentRuleEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDeploy

`func (o *EnvironmentDeploymentRuleEditRequest) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *EnvironmentDeploymentRuleEditRequest) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *EnvironmentDeploymentRuleEditRequest) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *EnvironmentDeploymentRuleEditRequest) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetAlwaysUp

`func (o *EnvironmentDeploymentRuleEditRequest) GetAlwaysUp() bool`

GetAlwaysUp returns the AlwaysUp field if non-nil, zero value otherwise.

### GetAlwaysUpOk

`func (o *EnvironmentDeploymentRuleEditRequest) GetAlwaysUpOk() (*bool, bool)`

GetAlwaysUpOk returns a tuple with the AlwaysUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysUp

`func (o *EnvironmentDeploymentRuleEditRequest) SetAlwaysUp(v bool)`

SetAlwaysUp sets AlwaysUp field to given value.

### HasAlwaysUp

`func (o *EnvironmentDeploymentRuleEditRequest) HasAlwaysUp() bool`

HasAlwaysUp returns a boolean if a field has been set.

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

### HasStartTime

`func (o *EnvironmentDeploymentRuleEditRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *EnvironmentDeploymentRuleEditRequest) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *EnvironmentDeploymentRuleEditRequest) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
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

### HasStopTime

`func (o *EnvironmentDeploymentRuleEditRequest) HasStopTime() bool`

HasStopTime returns a boolean if a field has been set.

### SetStopTimeNil

`func (o *EnvironmentDeploymentRuleEditRequest) SetStopTimeNil(b bool)`

 SetStopTimeNil sets the value for StopTime to be an explicit nil

### UnsetStopTime
`func (o *EnvironmentDeploymentRuleEditRequest) UnsetStopTime()`

UnsetStopTime ensures that no value is present for StopTime, not even an explicit nil
### GetWeekday

`func (o *EnvironmentDeploymentRuleEditRequest) GetWeekday() string`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *EnvironmentDeploymentRuleEditRequest) GetWeekdayOk() (*string, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *EnvironmentDeploymentRuleEditRequest) SetWeekday(v string)`

SetWeekday sets Weekday field to given value.

### HasWeekday

`func (o *EnvironmentDeploymentRuleEditRequest) HasWeekday() bool`

HasWeekday returns a boolean if a field has been set.

### SetWeekdayNil

`func (o *EnvironmentDeploymentRuleEditRequest) SetWeekdayNil(b bool)`

 SetWeekdayNil sets the value for Weekday to be an explicit nil

### UnsetWeekday
`func (o *EnvironmentDeploymentRuleEditRequest) UnsetWeekday()`

UnsetWeekday ensures that no value is present for Weekday, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


