# ReferenceObjectStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | **string** | Status is a state machine. It starts with &#x60;BUILDING&#x60; or &#x60;DEPLOYING&#x60; state (or &#x60;INITIALIZED&#x60;if auto-deploy is deactivated). Then finish with &#x60;*_ERROR&#x60; or any termination state.  | 
**Message** | Pointer to **NullableString** | message related to the state | [optional] 
**ServiceDeploymentStatus** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewReferenceObjectStatusResponse

`func NewReferenceObjectStatusResponse(id string, state string, ) *ReferenceObjectStatusResponse`

NewReferenceObjectStatusResponse instantiates a new ReferenceObjectStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceObjectStatusResponseWithDefaults

`func NewReferenceObjectStatusResponseWithDefaults() *ReferenceObjectStatusResponse`

NewReferenceObjectStatusResponseWithDefaults instantiates a new ReferenceObjectStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReferenceObjectStatusResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReferenceObjectStatusResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReferenceObjectStatusResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *ReferenceObjectStatusResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReferenceObjectStatusResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReferenceObjectStatusResponse) SetState(v string)`

SetState sets State field to given value.


### GetMessage

`func (o *ReferenceObjectStatusResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReferenceObjectStatusResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReferenceObjectStatusResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ReferenceObjectStatusResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ReferenceObjectStatusResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ReferenceObjectStatusResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetServiceDeploymentStatus

`func (o *ReferenceObjectStatusResponse) GetServiceDeploymentStatus() string`

GetServiceDeploymentStatus returns the ServiceDeploymentStatus field if non-nil, zero value otherwise.

### GetServiceDeploymentStatusOk

`func (o *ReferenceObjectStatusResponse) GetServiceDeploymentStatusOk() (*string, bool)`

GetServiceDeploymentStatusOk returns a tuple with the ServiceDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDeploymentStatus

`func (o *ReferenceObjectStatusResponse) SetServiceDeploymentStatus(v string)`

SetServiceDeploymentStatus sets ServiceDeploymentStatus field to given value.

### HasServiceDeploymentStatus

`func (o *ReferenceObjectStatusResponse) HasServiceDeploymentStatus() bool`

HasServiceDeploymentStatus returns a boolean if a field has been set.

### SetServiceDeploymentStatusNil

`func (o *ReferenceObjectStatusResponse) SetServiceDeploymentStatusNil(b bool)`

 SetServiceDeploymentStatusNil sets the value for ServiceDeploymentStatus to be an explicit nil

### UnsetServiceDeploymentStatus
`func (o *ReferenceObjectStatusResponse) UnsetServiceDeploymentStatus()`

UnsetServiceDeploymentStatus ensures that no value is present for ServiceDeploymentStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


