# ReferenceObjectStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**GlobalDeploymentStatus**](GlobalDeploymentStatus.md) |  | 
**Message** | Pointer to **NullableString** | message related to the state | [optional] 
**ServiceDeploymentStatus** | Pointer to [**NullableServiceDeploymentStatusEnum**](ServiceDeploymentStatusEnum.md) |  | [optional] 

## Methods

### NewReferenceObjectStatusResponse

`func NewReferenceObjectStatusResponse(id string, state GlobalDeploymentStatus, ) *ReferenceObjectStatusResponse`

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

`func (o *ReferenceObjectStatusResponse) GetState() GlobalDeploymentStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReferenceObjectStatusResponse) GetStateOk() (*GlobalDeploymentStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReferenceObjectStatusResponse) SetState(v GlobalDeploymentStatus)`

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

`func (o *ReferenceObjectStatusResponse) GetServiceDeploymentStatus() ServiceDeploymentStatusEnum`

GetServiceDeploymentStatus returns the ServiceDeploymentStatus field if non-nil, zero value otherwise.

### GetServiceDeploymentStatusOk

`func (o *ReferenceObjectStatusResponse) GetServiceDeploymentStatusOk() (*ServiceDeploymentStatusEnum, bool)`

GetServiceDeploymentStatusOk returns a tuple with the ServiceDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDeploymentStatus

`func (o *ReferenceObjectStatusResponse) SetServiceDeploymentStatus(v ServiceDeploymentStatusEnum)`

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


