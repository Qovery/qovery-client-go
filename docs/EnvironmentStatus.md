# EnvironmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**LastDeploymentDate** | Pointer to **NullableTime** |  | [optional] 
**LastDeploymentState** | [**StateEnum**](StateEnum.md) |  | 
**LastDeploymentId** | Pointer to **NullableString** |  | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


