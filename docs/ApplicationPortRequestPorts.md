# ApplicationPortRequestPorts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**InternalPort** | **int32** | The listening port of your application | 
**ExternalPort** | Pointer to **int32** | The exposed port for your application. This is optional. If not set a default port will be used. | [optional] 
**PubliclyAccessible** | Pointer to **bool** | Expose the port to the world | [optional] 
**Protocol** | **string** |  | [default to "HTTP"]

## Methods

### NewApplicationPortRequestPorts

`func NewApplicationPortRequestPorts(internalPort int32, protocol string, ) *ApplicationPortRequestPorts`

NewApplicationPortRequestPorts instantiates a new ApplicationPortRequestPorts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPortRequestPortsWithDefaults

`func NewApplicationPortRequestPortsWithDefaults() *ApplicationPortRequestPorts`

NewApplicationPortRequestPortsWithDefaults instantiates a new ApplicationPortRequestPorts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationPortRequestPorts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationPortRequestPorts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationPortRequestPorts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationPortRequestPorts) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalPort

`func (o *ApplicationPortRequestPorts) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *ApplicationPortRequestPorts) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *ApplicationPortRequestPorts) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.


### GetExternalPort

`func (o *ApplicationPortRequestPorts) GetExternalPort() int32`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *ApplicationPortRequestPorts) GetExternalPortOk() (*int32, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *ApplicationPortRequestPorts) SetExternalPort(v int32)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *ApplicationPortRequestPorts) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetPubliclyAccessible

`func (o *ApplicationPortRequestPorts) GetPubliclyAccessible() bool`

GetPubliclyAccessible returns the PubliclyAccessible field if non-nil, zero value otherwise.

### GetPubliclyAccessibleOk

`func (o *ApplicationPortRequestPorts) GetPubliclyAccessibleOk() (*bool, bool)`

GetPubliclyAccessibleOk returns a tuple with the PubliclyAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubliclyAccessible

`func (o *ApplicationPortRequestPorts) SetPubliclyAccessible(v bool)`

SetPubliclyAccessible sets PubliclyAccessible field to given value.

### HasPubliclyAccessible

`func (o *ApplicationPortRequestPorts) HasPubliclyAccessible() bool`

HasPubliclyAccessible returns a boolean if a field has been set.

### GetProtocol

`func (o *ApplicationPortRequestPorts) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApplicationPortRequestPorts) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApplicationPortRequestPorts) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


