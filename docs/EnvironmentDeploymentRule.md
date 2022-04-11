# EnvironmentDeploymentRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**AutoDeploy** | Pointer to **bool** |  | [optional] [default to true]
**AutoStop** | Pointer to **bool** |  | [optional] [default to false]
**AutoDelete** | Pointer to **bool** |  | [optional] [default to false]
**AutoPreview** | Pointer to **bool** |  | [optional] [default to false]
**Timezone** | **string** |  | 
**StartTime** | **time.Time** |  | 
**StopTime** | **time.Time** |  | 
**Weekdays** | [**[]WeekdayEnum**](WeekdayEnum.md) |  | 

## Methods

### NewEnvironmentDeploymentRule

`func NewEnvironmentDeploymentRule(id string, createdAt time.Time, timezone string, startTime time.Time, stopTime time.Time, weekdays []WeekdayEnum, ) *EnvironmentDeploymentRule`

NewEnvironmentDeploymentRule instantiates a new EnvironmentDeploymentRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDeploymentRuleWithDefaults

`func NewEnvironmentDeploymentRuleWithDefaults() *EnvironmentDeploymentRule`

NewEnvironmentDeploymentRuleWithDefaults instantiates a new EnvironmentDeploymentRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentDeploymentRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentDeploymentRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentDeploymentRule) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EnvironmentDeploymentRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentDeploymentRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentDeploymentRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EnvironmentDeploymentRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentDeploymentRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentDeploymentRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentDeploymentRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAutoDeploy

`func (o *EnvironmentDeploymentRule) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *EnvironmentDeploymentRule) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *EnvironmentDeploymentRule) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *EnvironmentDeploymentRule) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetAutoStop

`func (o *EnvironmentDeploymentRule) GetAutoStop() bool`

GetAutoStop returns the AutoStop field if non-nil, zero value otherwise.

### GetAutoStopOk

`func (o *EnvironmentDeploymentRule) GetAutoStopOk() (*bool, bool)`

GetAutoStopOk returns a tuple with the AutoStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStop

`func (o *EnvironmentDeploymentRule) SetAutoStop(v bool)`

SetAutoStop sets AutoStop field to given value.

### HasAutoStop

`func (o *EnvironmentDeploymentRule) HasAutoStop() bool`

HasAutoStop returns a boolean if a field has been set.

### GetAutoDelete

`func (o *EnvironmentDeploymentRule) GetAutoDelete() bool`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *EnvironmentDeploymentRule) GetAutoDeleteOk() (*bool, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *EnvironmentDeploymentRule) SetAutoDelete(v bool)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *EnvironmentDeploymentRule) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetAutoPreview

`func (o *EnvironmentDeploymentRule) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *EnvironmentDeploymentRule) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *EnvironmentDeploymentRule) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.

### HasAutoPreview

`func (o *EnvironmentDeploymentRule) HasAutoPreview() bool`

HasAutoPreview returns a boolean if a field has been set.

### GetTimezone

`func (o *EnvironmentDeploymentRule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *EnvironmentDeploymentRule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *EnvironmentDeploymentRule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStartTime

`func (o *EnvironmentDeploymentRule) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EnvironmentDeploymentRule) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EnvironmentDeploymentRule) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetStopTime

`func (o *EnvironmentDeploymentRule) GetStopTime() time.Time`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *EnvironmentDeploymentRule) GetStopTimeOk() (*time.Time, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *EnvironmentDeploymentRule) SetStopTime(v time.Time)`

SetStopTime sets StopTime field to given value.


### GetWeekdays

`func (o *EnvironmentDeploymentRule) GetWeekdays() []WeekdayEnum`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *EnvironmentDeploymentRule) GetWeekdaysOk() (*[]WeekdayEnum, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *EnvironmentDeploymentRule) SetWeekdays(v []WeekdayEnum)`

SetWeekdays sets Weekdays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


