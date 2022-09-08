# ReferenceObjectStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**Message** | Pointer to **NullableString** | message related to the state | [optional] 
**ServiceDeploymentStatus** | [**ServiceDeploymentStatusEnum**](ServiceDeploymentStatusEnum.md) |  | 
**LastDeploymentDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewReferenceObjectStatus

`func NewReferenceObjectStatus(id string, state StateEnum, serviceDeploymentStatus ServiceDeploymentStatusEnum, ) *ReferenceObjectStatus`

NewReferenceObjectStatus instantiates a new ReferenceObjectStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceObjectStatusWithDefaults

`func NewReferenceObjectStatusWithDefaults() *ReferenceObjectStatus`

NewReferenceObjectStatusWithDefaults instantiates a new ReferenceObjectStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReferenceObjectStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReferenceObjectStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReferenceObjectStatus) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *ReferenceObjectStatus) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReferenceObjectStatus) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReferenceObjectStatus) SetState(v StateEnum)`

SetState sets State field to given value.


### GetMessage

`func (o *ReferenceObjectStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReferenceObjectStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReferenceObjectStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ReferenceObjectStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ReferenceObjectStatus) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ReferenceObjectStatus) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetServiceDeploymentStatus

`func (o *ReferenceObjectStatus) GetServiceDeploymentStatus() ServiceDeploymentStatusEnum`

GetServiceDeploymentStatus returns the ServiceDeploymentStatus field if non-nil, zero value otherwise.

### GetServiceDeploymentStatusOk

`func (o *ReferenceObjectStatus) GetServiceDeploymentStatusOk() (*ServiceDeploymentStatusEnum, bool)`

GetServiceDeploymentStatusOk returns a tuple with the ServiceDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDeploymentStatus

`func (o *ReferenceObjectStatus) SetServiceDeploymentStatus(v ServiceDeploymentStatusEnum)`

SetServiceDeploymentStatus sets ServiceDeploymentStatus field to given value.


### GetLastDeploymentDate

`func (o *ReferenceObjectStatus) GetLastDeploymentDate() time.Time`

GetLastDeploymentDate returns the LastDeploymentDate field if non-nil, zero value otherwise.

### GetLastDeploymentDateOk

`func (o *ReferenceObjectStatus) GetLastDeploymentDateOk() (*time.Time, bool)`

GetLastDeploymentDateOk returns a tuple with the LastDeploymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploymentDate

`func (o *ReferenceObjectStatus) SetLastDeploymentDate(v time.Time)`

SetLastDeploymentDate sets LastDeploymentDate field to given value.

### HasLastDeploymentDate

`func (o *ReferenceObjectStatus) HasLastDeploymentDate() bool`

HasLastDeploymentDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


