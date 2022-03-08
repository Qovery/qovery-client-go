# EnvironmentLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Scope** | Pointer to [**EnvironmentLogResponseScope**](EnvironmentLogResponseScope.md) |  | [optional] 
**State** | Pointer to [**GlobalDeploymentStatus**](GlobalDeploymentStatus.md) |  | [optional] 
**Message** | **NullableString** | Log message | 
**ExecutionId** | Pointer to **string** | Only for errors. Helps Qovery team to investigate. | [optional] 
**Hint** | Pointer to **string** |  | [optional] 

## Methods

### NewEnvironmentLogResponse

`func NewEnvironmentLogResponse(id string, createdAt time.Time, message NullableString, ) *EnvironmentLogResponse`

NewEnvironmentLogResponse instantiates a new EnvironmentLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentLogResponseWithDefaults

`func NewEnvironmentLogResponseWithDefaults() *EnvironmentLogResponse`

NewEnvironmentLogResponseWithDefaults instantiates a new EnvironmentLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentLogResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentLogResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentLogResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EnvironmentLogResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentLogResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentLogResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetScope

`func (o *EnvironmentLogResponse) GetScope() EnvironmentLogResponseScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EnvironmentLogResponse) GetScopeOk() (*EnvironmentLogResponseScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EnvironmentLogResponse) SetScope(v EnvironmentLogResponseScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EnvironmentLogResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetState

`func (o *EnvironmentLogResponse) GetState() GlobalDeploymentStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EnvironmentLogResponse) GetStateOk() (*GlobalDeploymentStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EnvironmentLogResponse) SetState(v GlobalDeploymentStatus)`

SetState sets State field to given value.

### HasState

`func (o *EnvironmentLogResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessage

`func (o *EnvironmentLogResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EnvironmentLogResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EnvironmentLogResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *EnvironmentLogResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *EnvironmentLogResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetExecutionId

`func (o *EnvironmentLogResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *EnvironmentLogResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *EnvironmentLogResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *EnvironmentLogResponse) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetHint

`func (o *EnvironmentLogResponse) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *EnvironmentLogResponse) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *EnvironmentLogResponse) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *EnvironmentLogResponse) HasHint() bool`

HasHint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


