# ServicePortPortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**InternalPort** | **int32** | The listening port of your service. | 
**ExternalPort** | Pointer to **int32** | The exposed port for your service. This is optional. If not set a default port will be used. | [optional] 
**PubliclyAccessible** | **bool** | Expose the port to the world | 
**Protocol** | [**PortProtocolEnum**](PortProtocolEnum.md) |  | [default to PORTPROTOCOLENUM_HTTP]

## Methods

### NewServicePortPortsInner

`func NewServicePortPortsInner(id string, internalPort int32, publiclyAccessible bool, protocol PortProtocolEnum, ) *ServicePortPortsInner`

NewServicePortPortsInner instantiates a new ServicePortPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePortPortsInnerWithDefaults

`func NewServicePortPortsInnerWithDefaults() *ServicePortPortsInner`

NewServicePortPortsInnerWithDefaults instantiates a new ServicePortPortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServicePortPortsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServicePortPortsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServicePortPortsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ServicePortPortsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePortPortsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePortPortsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServicePortPortsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalPort

`func (o *ServicePortPortsInner) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *ServicePortPortsInner) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *ServicePortPortsInner) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.


### GetExternalPort

`func (o *ServicePortPortsInner) GetExternalPort() int32`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *ServicePortPortsInner) GetExternalPortOk() (*int32, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *ServicePortPortsInner) SetExternalPort(v int32)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *ServicePortPortsInner) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetPubliclyAccessible

`func (o *ServicePortPortsInner) GetPubliclyAccessible() bool`

GetPubliclyAccessible returns the PubliclyAccessible field if non-nil, zero value otherwise.

### GetPubliclyAccessibleOk

`func (o *ServicePortPortsInner) GetPubliclyAccessibleOk() (*bool, bool)`

GetPubliclyAccessibleOk returns a tuple with the PubliclyAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubliclyAccessible

`func (o *ServicePortPortsInner) SetPubliclyAccessible(v bool)`

SetPubliclyAccessible sets PubliclyAccessible field to given value.


### GetProtocol

`func (o *ServicePortPortsInner) GetProtocol() PortProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ServicePortPortsInner) GetProtocolOk() (*PortProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ServicePortPortsInner) SetProtocol(v PortProtocolEnum)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


