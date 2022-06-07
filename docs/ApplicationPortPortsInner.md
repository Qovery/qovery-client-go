# ApplicationPortPortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**InternalPort** | **int32** | The listening port of your application | 
**ExternalPort** | Pointer to **int32** | The exposed port for your application. This is optional. If not set a default port will be used. | [optional] 
**PubliclyAccessible** | **bool** | Expose the port to the world | 
**Protocol** | Pointer to [**PortProtocolEnum**](PortProtocolEnum.md) |  | [optional] [default to PORTPROTOCOLENUM_HTTP]

## Methods

### NewApplicationPortPortsInner

`func NewApplicationPortPortsInner(internalPort int32, publiclyAccessible bool, ) *ApplicationPortPortsInner`

NewApplicationPortPortsInner instantiates a new ApplicationPortPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPortPortsInnerWithDefaults

`func NewApplicationPortPortsInnerWithDefaults() *ApplicationPortPortsInner`

NewApplicationPortPortsInnerWithDefaults instantiates a new ApplicationPortPortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationPortPortsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationPortPortsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationPortPortsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationPortPortsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApplicationPortPortsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationPortPortsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationPortPortsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationPortPortsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApplicationPortPortsInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApplicationPortPortsInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetInternalPort

`func (o *ApplicationPortPortsInner) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *ApplicationPortPortsInner) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *ApplicationPortPortsInner) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.


### GetExternalPort

`func (o *ApplicationPortPortsInner) GetExternalPort() int32`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *ApplicationPortPortsInner) GetExternalPortOk() (*int32, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *ApplicationPortPortsInner) SetExternalPort(v int32)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *ApplicationPortPortsInner) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetPubliclyAccessible

`func (o *ApplicationPortPortsInner) GetPubliclyAccessible() bool`

GetPubliclyAccessible returns the PubliclyAccessible field if non-nil, zero value otherwise.

### GetPubliclyAccessibleOk

`func (o *ApplicationPortPortsInner) GetPubliclyAccessibleOk() (*bool, bool)`

GetPubliclyAccessibleOk returns a tuple with the PubliclyAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubliclyAccessible

`func (o *ApplicationPortPortsInner) SetPubliclyAccessible(v bool)`

SetPubliclyAccessible sets PubliclyAccessible field to given value.


### GetProtocol

`func (o *ApplicationPortPortsInner) GetProtocol() PortProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApplicationPortPortsInner) GetProtocolOk() (*PortProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApplicationPortPortsInner) SetProtocol(v PortProtocolEnum)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApplicationPortPortsInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


