# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**Message** | Pointer to **NullableString** | message related to the state | [optional] 
**ServiceDeploymentStatus** | [**NullableServiceDeploymentStatusEnum**](ServiceDeploymentStatusEnum.md) |  | 
**LastDeploymentDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStatus

`func NewStatus(id string, state StateEnum, serviceDeploymentStatus NullableServiceDeploymentStatusEnum, ) *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Status) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Status) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Status) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *Status) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Status) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Status) SetState(v StateEnum)`

SetState sets State field to given value.


### GetMessage

`func (o *Status) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Status) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Status) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Status) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *Status) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *Status) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetServiceDeploymentStatus

`func (o *Status) GetServiceDeploymentStatus() ServiceDeploymentStatusEnum`

GetServiceDeploymentStatus returns the ServiceDeploymentStatus field if non-nil, zero value otherwise.

### GetServiceDeploymentStatusOk

`func (o *Status) GetServiceDeploymentStatusOk() (*ServiceDeploymentStatusEnum, bool)`

GetServiceDeploymentStatusOk returns a tuple with the ServiceDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDeploymentStatus

`func (o *Status) SetServiceDeploymentStatus(v ServiceDeploymentStatusEnum)`

SetServiceDeploymentStatus sets ServiceDeploymentStatus field to given value.


### SetServiceDeploymentStatusNil

`func (o *Status) SetServiceDeploymentStatusNil(b bool)`

 SetServiceDeploymentStatusNil sets the value for ServiceDeploymentStatus to be an explicit nil

### UnsetServiceDeploymentStatus
`func (o *Status) UnsetServiceDeploymentStatus()`

UnsetServiceDeploymentStatus ensures that no value is present for ServiceDeploymentStatus, not even an explicit nil
### GetLastDeploymentDate

`func (o *Status) GetLastDeploymentDate() time.Time`

GetLastDeploymentDate returns the LastDeploymentDate field if non-nil, zero value otherwise.

### GetLastDeploymentDateOk

`func (o *Status) GetLastDeploymentDateOk() (*time.Time, bool)`

GetLastDeploymentDateOk returns a tuple with the LastDeploymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploymentDate

`func (o *Status) SetLastDeploymentDate(v time.Time)`

SetLastDeploymentDate sets LastDeploymentDate field to given value.

### HasLastDeploymentDate

`func (o *Status) HasLastDeploymentDate() bool`

HasLastDeploymentDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


