# EnvironmentLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Scope** | Pointer to [**EnvironmentLogScope**](EnvironmentLogScope.md) |  | [optional] 
**State** | Pointer to [**GlobalDeploymentStatus**](GlobalDeploymentStatus.md) |  | [optional] 
**Message** | **NullableString** | Log message | 
**ExecutionId** | Pointer to **string** | Only for errors. Helps Qovery team to investigate. | [optional] 
**Hint** | Pointer to **string** |  | [optional] 

## Methods

### NewEnvironmentLog

`func NewEnvironmentLog(id string, createdAt time.Time, message NullableString, ) *EnvironmentLog`

NewEnvironmentLog instantiates a new EnvironmentLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentLogWithDefaults

`func NewEnvironmentLogWithDefaults() *EnvironmentLog`

NewEnvironmentLogWithDefaults instantiates a new EnvironmentLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentLog) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EnvironmentLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetScope

`func (o *EnvironmentLog) GetScope() EnvironmentLogScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EnvironmentLog) GetScopeOk() (*EnvironmentLogScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EnvironmentLog) SetScope(v EnvironmentLogScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EnvironmentLog) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetState

`func (o *EnvironmentLog) GetState() GlobalDeploymentStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EnvironmentLog) GetStateOk() (*GlobalDeploymentStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EnvironmentLog) SetState(v GlobalDeploymentStatus)`

SetState sets State field to given value.

### HasState

`func (o *EnvironmentLog) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessage

`func (o *EnvironmentLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EnvironmentLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EnvironmentLog) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *EnvironmentLog) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *EnvironmentLog) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetExecutionId

`func (o *EnvironmentLog) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *EnvironmentLog) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *EnvironmentLog) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *EnvironmentLog) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetHint

`func (o *EnvironmentLog) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *EnvironmentLog) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *EnvironmentLog) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *EnvironmentLog) HasHint() bool`

HasHint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


