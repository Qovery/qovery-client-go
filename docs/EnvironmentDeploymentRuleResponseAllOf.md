# EnvironmentDeploymentRuleResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeploy** | Pointer to **bool** |  | [optional] [default to true]
**AutoStop** | Pointer to **bool** |  | [optional] [default to false]
**AutoDelete** | Pointer to **bool** |  | [optional] [default to false]
**AutoPreview** | Pointer to **bool** |  | [optional] [default to false]
**Timezone** | **string** |  | 
**StartTime** | **time.Time** |  | 
**StopTime** | **time.Time** |  | 
**Weekdays** | **[]string** |  | 

## Methods

### NewEnvironmentDeploymentRuleResponseAllOf

`func NewEnvironmentDeploymentRuleResponseAllOf(timezone string, startTime time.Time, stopTime time.Time, weekdays []string, ) *EnvironmentDeploymentRuleResponseAllOf`

NewEnvironmentDeploymentRuleResponseAllOf instantiates a new EnvironmentDeploymentRuleResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDeploymentRuleResponseAllOfWithDefaults

`func NewEnvironmentDeploymentRuleResponseAllOfWithDefaults() *EnvironmentDeploymentRuleResponseAllOf`

NewEnvironmentDeploymentRuleResponseAllOfWithDefaults instantiates a new EnvironmentDeploymentRuleResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDeploy

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *EnvironmentDeploymentRuleResponseAllOf) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *EnvironmentDeploymentRuleResponseAllOf) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetAutoStop

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetAutoStop() bool`

GetAutoStop returns the AutoStop field if non-nil, zero value otherwise.

### GetAutoStopOk

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetAutoStopOk() (*bool, bool)`

GetAutoStopOk returns a tuple with the AutoStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStop

`func (o *EnvironmentDeploymentRuleResponseAllOf) SetAutoStop(v bool)`

SetAutoStop sets AutoStop field to given value.

### HasAutoStop

`func (o *EnvironmentDeploymentRuleResponseAllOf) HasAutoStop() bool`

HasAutoStop returns a boolean if a field has been set.

### GetAutoDelete

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetAutoDelete() bool`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetAutoDeleteOk() (*bool, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *EnvironmentDeploymentRuleResponseAllOf) SetAutoDelete(v bool)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *EnvironmentDeploymentRuleResponseAllOf) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetAutoPreview

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *EnvironmentDeploymentRuleResponseAllOf) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.

### HasAutoPreview

`func (o *EnvironmentDeploymentRuleResponseAllOf) HasAutoPreview() bool`

HasAutoPreview returns a boolean if a field has been set.

### GetTimezone

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *EnvironmentDeploymentRuleResponseAllOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStartTime

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EnvironmentDeploymentRuleResponseAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetStopTime

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetStopTime() time.Time`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetStopTimeOk() (*time.Time, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *EnvironmentDeploymentRuleResponseAllOf) SetStopTime(v time.Time)`

SetStopTime sets StopTime field to given value.


### GetWeekdays

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetWeekdays() []string`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *EnvironmentDeploymentRuleResponseAllOf) GetWeekdaysOk() (*[]string, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *EnvironmentDeploymentRuleResponseAllOf) SetWeekdays(v []string)`

SetWeekdays sets Weekdays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


