# ApplicationPortRequestPortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**InternalPort** | **int32** | The listening port of your application | 
**ExternalPort** | Pointer to **int32** | The exposed port for your application. This is optional. If not set a default port will be used. | [optional] 
**PubliclyAccessible** | **bool** | Expose the port to the world | 
**Protocol** | Pointer to [**PortProtocolEnum**](PortProtocolEnum.md) |  | [optional] [default to PORTPROTOCOLENUM_HTTP]

## Methods

### NewApplicationPortRequestPortsInner

`func NewApplicationPortRequestPortsInner(internalPort int32, publiclyAccessible bool, ) *ApplicationPortRequestPortsInner`

NewApplicationPortRequestPortsInner instantiates a new ApplicationPortRequestPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPortRequestPortsInnerWithDefaults

`func NewApplicationPortRequestPortsInnerWithDefaults() *ApplicationPortRequestPortsInner`

NewApplicationPortRequestPortsInnerWithDefaults instantiates a new ApplicationPortRequestPortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationPortRequestPortsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationPortRequestPortsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationPortRequestPortsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationPortRequestPortsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApplicationPortRequestPortsInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApplicationPortRequestPortsInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetInternalPort

`func (o *ApplicationPortRequestPortsInner) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *ApplicationPortRequestPortsInner) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *ApplicationPortRequestPortsInner) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.


### GetExternalPort

`func (o *ApplicationPortRequestPortsInner) GetExternalPort() int32`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *ApplicationPortRequestPortsInner) GetExternalPortOk() (*int32, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *ApplicationPortRequestPortsInner) SetExternalPort(v int32)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *ApplicationPortRequestPortsInner) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetPubliclyAccessible

`func (o *ApplicationPortRequestPortsInner) GetPubliclyAccessible() bool`

GetPubliclyAccessible returns the PubliclyAccessible field if non-nil, zero value otherwise.

### GetPubliclyAccessibleOk

`func (o *ApplicationPortRequestPortsInner) GetPubliclyAccessibleOk() (*bool, bool)`

GetPubliclyAccessibleOk returns a tuple with the PubliclyAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubliclyAccessible

`func (o *ApplicationPortRequestPortsInner) SetPubliclyAccessible(v bool)`

SetPubliclyAccessible sets PubliclyAccessible field to given value.


### GetProtocol

`func (o *ApplicationPortRequestPortsInner) GetProtocol() PortProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApplicationPortRequestPortsInner) GetProtocolOk() (*PortProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApplicationPortRequestPortsInner) SetProtocol(v PortProtocolEnum)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApplicationPortRequestPortsInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


