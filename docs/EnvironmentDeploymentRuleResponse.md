# EnvironmentDeploymentRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeploy** | Pointer to **bool** |  | [optional] [default to true]
**AlwaysUp** | Pointer to **bool** |  | [optional] [default to false]
**StartTime** | Pointer to **NullableTime** | specify value only if always_up &#x3D; false | [optional] 
**StopTime** | Pointer to **NullableTime** | specify value only if always_up &#x3D; false | [optional] 
**Weekday** | Pointer to **NullableString** | specify value only if always_up &#x3D; false | [optional] 
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

### GetAlwaysUp

`func (o *EnvironmentDeploymentRuleResponse) GetAlwaysUp() bool`

GetAlwaysUp returns the AlwaysUp field if non-nil, zero value otherwise.

### GetAlwaysUpOk

`func (o *EnvironmentDeploymentRuleResponse) GetAlwaysUpOk() (*bool, bool)`

GetAlwaysUpOk returns a tuple with the AlwaysUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysUp

`func (o *EnvironmentDeploymentRuleResponse) SetAlwaysUp(v bool)`

SetAlwaysUp sets AlwaysUp field to given value.

### HasAlwaysUp

`func (o *EnvironmentDeploymentRuleResponse) HasAlwaysUp() bool`

HasAlwaysUp returns a boolean if a field has been set.

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
### GetWeekday

`func (o *EnvironmentDeploymentRuleResponse) GetWeekday() string`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *EnvironmentDeploymentRuleResponse) GetWeekdayOk() (*string, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *EnvironmentDeploymentRuleResponse) SetWeekday(v string)`

SetWeekday sets Weekday field to given value.

### HasWeekday

`func (o *EnvironmentDeploymentRuleResponse) HasWeekday() bool`

HasWeekday returns a boolean if a field has been set.

### SetWeekdayNil

`func (o *EnvironmentDeploymentRuleResponse) SetWeekdayNil(b bool)`

 SetWeekdayNil sets the value for Weekday to be an explicit nil

### UnsetWeekday
`func (o *EnvironmentDeploymentRuleResponse) UnsetWeekday()`

UnsetWeekday ensures that no value is present for Weekday, not even an explicit nil
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


