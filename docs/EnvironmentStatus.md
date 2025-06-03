# EnvironmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**LastDeploymentDate** | Pointer to **NullableTime** |  | [optional] 
**LastDeploymentState** | [**StateEnum**](StateEnum.md) |  | 
**LastDeploymentId** | Pointer to **NullableString** |  | [optional] 
**TotalDeploymentDurationInSeconds** | Pointer to **NullableInt32** |  | [optional] 
**Origin** | Pointer to [**NullableEnvironmentStatusEventOriginEnum**](EnvironmentStatusEventOriginEnum.md) |  | [optional] 
**TriggeredBy** | Pointer to **NullableString** |  | [optional] 
**DeploymentStatus** | Pointer to [**EnvironmentDeploymentStatusEnum**](EnvironmentDeploymentStatusEnum.md) |  | [optional] 

## Methods

### NewEnvironmentStatus

`func NewEnvironmentStatus(id string, state StateEnum, lastDeploymentState StateEnum, ) *EnvironmentStatus`

NewEnvironmentStatus instantiates a new EnvironmentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentStatusWithDefaults

`func NewEnvironmentStatusWithDefaults() *EnvironmentStatus`

NewEnvironmentStatusWithDefaults instantiates a new EnvironmentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentStatus) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *EnvironmentStatus) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EnvironmentStatus) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EnvironmentStatus) SetState(v StateEnum)`

SetState sets State field to given value.


### GetLastDeploymentDate

`func (o *EnvironmentStatus) GetLastDeploymentDate() time.Time`

GetLastDeploymentDate returns the LastDeploymentDate field if non-nil, zero value otherwise.

### GetLastDeploymentDateOk

`func (o *EnvironmentStatus) GetLastDeploymentDateOk() (*time.Time, bool)`

GetLastDeploymentDateOk returns a tuple with the LastDeploymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploymentDate

`func (o *EnvironmentStatus) SetLastDeploymentDate(v time.Time)`

SetLastDeploymentDate sets LastDeploymentDate field to given value.

### HasLastDeploymentDate

`func (o *EnvironmentStatus) HasLastDeploymentDate() bool`

HasLastDeploymentDate returns a boolean if a field has been set.

### SetLastDeploymentDateNil

`func (o *EnvironmentStatus) SetLastDeploymentDateNil(b bool)`

 SetLastDeploymentDateNil sets the value for LastDeploymentDate to be an explicit nil

### UnsetLastDeploymentDate
`func (o *EnvironmentStatus) UnsetLastDeploymentDate()`

UnsetLastDeploymentDate ensures that no value is present for LastDeploymentDate, not even an explicit nil
### GetLastDeploymentState

`func (o *EnvironmentStatus) GetLastDeploymentState() StateEnum`

GetLastDeploymentState returns the LastDeploymentState field if non-nil, zero value otherwise.

### GetLastDeploymentStateOk

`func (o *EnvironmentStatus) GetLastDeploymentStateOk() (*StateEnum, bool)`

GetLastDeploymentStateOk returns a tuple with the LastDeploymentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploymentState

`func (o *EnvironmentStatus) SetLastDeploymentState(v StateEnum)`

SetLastDeploymentState sets LastDeploymentState field to given value.


### GetLastDeploymentId

`func (o *EnvironmentStatus) GetLastDeploymentId() string`

GetLastDeploymentId returns the LastDeploymentId field if non-nil, zero value otherwise.

### GetLastDeploymentIdOk

`func (o *EnvironmentStatus) GetLastDeploymentIdOk() (*string, bool)`

GetLastDeploymentIdOk returns a tuple with the LastDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploymentId

`func (o *EnvironmentStatus) SetLastDeploymentId(v string)`

SetLastDeploymentId sets LastDeploymentId field to given value.

### HasLastDeploymentId

`func (o *EnvironmentStatus) HasLastDeploymentId() bool`

HasLastDeploymentId returns a boolean if a field has been set.

### SetLastDeploymentIdNil

`func (o *EnvironmentStatus) SetLastDeploymentIdNil(b bool)`

 SetLastDeploymentIdNil sets the value for LastDeploymentId to be an explicit nil

### UnsetLastDeploymentId
`func (o *EnvironmentStatus) UnsetLastDeploymentId()`

UnsetLastDeploymentId ensures that no value is present for LastDeploymentId, not even an explicit nil
### GetTotalDeploymentDurationInSeconds

`func (o *EnvironmentStatus) GetTotalDeploymentDurationInSeconds() int32`

GetTotalDeploymentDurationInSeconds returns the TotalDeploymentDurationInSeconds field if non-nil, zero value otherwise.

### GetTotalDeploymentDurationInSecondsOk

`func (o *EnvironmentStatus) GetTotalDeploymentDurationInSecondsOk() (*int32, bool)`

GetTotalDeploymentDurationInSecondsOk returns a tuple with the TotalDeploymentDurationInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeploymentDurationInSeconds

`func (o *EnvironmentStatus) SetTotalDeploymentDurationInSeconds(v int32)`

SetTotalDeploymentDurationInSeconds sets TotalDeploymentDurationInSeconds field to given value.

### HasTotalDeploymentDurationInSeconds

`func (o *EnvironmentStatus) HasTotalDeploymentDurationInSeconds() bool`

HasTotalDeploymentDurationInSeconds returns a boolean if a field has been set.

### SetTotalDeploymentDurationInSecondsNil

`func (o *EnvironmentStatus) SetTotalDeploymentDurationInSecondsNil(b bool)`

 SetTotalDeploymentDurationInSecondsNil sets the value for TotalDeploymentDurationInSeconds to be an explicit nil

### UnsetTotalDeploymentDurationInSeconds
`func (o *EnvironmentStatus) UnsetTotalDeploymentDurationInSeconds()`

UnsetTotalDeploymentDurationInSeconds ensures that no value is present for TotalDeploymentDurationInSeconds, not even an explicit nil
### GetOrigin

`func (o *EnvironmentStatus) GetOrigin() EnvironmentStatusEventOriginEnum`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *EnvironmentStatus) GetOriginOk() (*EnvironmentStatusEventOriginEnum, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *EnvironmentStatus) SetOrigin(v EnvironmentStatusEventOriginEnum)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *EnvironmentStatus) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *EnvironmentStatus) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *EnvironmentStatus) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetTriggeredBy

`func (o *EnvironmentStatus) GetTriggeredBy() string`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *EnvironmentStatus) GetTriggeredByOk() (*string, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *EnvironmentStatus) SetTriggeredBy(v string)`

SetTriggeredBy sets TriggeredBy field to given value.

### HasTriggeredBy

`func (o *EnvironmentStatus) HasTriggeredBy() bool`

HasTriggeredBy returns a boolean if a field has been set.

### SetTriggeredByNil

`func (o *EnvironmentStatus) SetTriggeredByNil(b bool)`

 SetTriggeredByNil sets the value for TriggeredBy to be an explicit nil

### UnsetTriggeredBy
`func (o *EnvironmentStatus) UnsetTriggeredBy()`

UnsetTriggeredBy ensures that no value is present for TriggeredBy, not even an explicit nil
### GetDeploymentStatus

`func (o *EnvironmentStatus) GetDeploymentStatus() EnvironmentDeploymentStatusEnum`

GetDeploymentStatus returns the DeploymentStatus field if non-nil, zero value otherwise.

### GetDeploymentStatusOk

`func (o *EnvironmentStatus) GetDeploymentStatusOk() (*EnvironmentDeploymentStatusEnum, bool)`

GetDeploymentStatusOk returns a tuple with the DeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentStatus

`func (o *EnvironmentStatus) SetDeploymentStatus(v EnvironmentDeploymentStatusEnum)`

SetDeploymentStatus sets DeploymentStatus field to given value.

### HasDeploymentStatus

`func (o *EnvironmentStatus) HasDeploymentStatus() bool`

HasDeploymentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


