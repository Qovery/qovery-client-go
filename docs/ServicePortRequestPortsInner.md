# ServicePortRequestPortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**InternalPort** | **int32** | The listening port of your service. | 
**ExternalPort** | Pointer to **int32** | The exposed port for your service. This is optional. If not set a default port will be used. | [optional] 
**PubliclyAccessible** | **bool** | Expose the port to the world | 
**IsDefault** | Pointer to **bool** | is the default port to use for domain | [optional] 
**HasReadinessProbe** | Pointer to **bool** | is port to use for probe readiness check | [optional] 
**HasLivenessProbe** | Pointer to **bool** | is port to use for probe liveness check | [optional] 
**Protocol** | Pointer to [**PortProtocolEnum**](PortProtocolEnum.md) |  | [optional] [default to PORTPROTOCOLENUM_HTTP]

## Methods

### NewServicePortRequestPortsInner

`func NewServicePortRequestPortsInner(internalPort int32, publiclyAccessible bool, ) *ServicePortRequestPortsInner`

NewServicePortRequestPortsInner instantiates a new ServicePortRequestPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePortRequestPortsInnerWithDefaults

`func NewServicePortRequestPortsInnerWithDefaults() *ServicePortRequestPortsInner`

NewServicePortRequestPortsInnerWithDefaults instantiates a new ServicePortRequestPortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServicePortRequestPortsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePortRequestPortsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePortRequestPortsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServicePortRequestPortsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalPort

`func (o *ServicePortRequestPortsInner) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *ServicePortRequestPortsInner) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *ServicePortRequestPortsInner) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.


### GetExternalPort

`func (o *ServicePortRequestPortsInner) GetExternalPort() int32`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *ServicePortRequestPortsInner) GetExternalPortOk() (*int32, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *ServicePortRequestPortsInner) SetExternalPort(v int32)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *ServicePortRequestPortsInner) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetPubliclyAccessible

`func (o *ServicePortRequestPortsInner) GetPubliclyAccessible() bool`

GetPubliclyAccessible returns the PubliclyAccessible field if non-nil, zero value otherwise.

### GetPubliclyAccessibleOk

`func (o *ServicePortRequestPortsInner) GetPubliclyAccessibleOk() (*bool, bool)`

GetPubliclyAccessibleOk returns a tuple with the PubliclyAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubliclyAccessible

`func (o *ServicePortRequestPortsInner) SetPubliclyAccessible(v bool)`

SetPubliclyAccessible sets PubliclyAccessible field to given value.


### GetIsDefault

`func (o *ServicePortRequestPortsInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ServicePortRequestPortsInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ServicePortRequestPortsInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ServicePortRequestPortsInner) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetHasReadinessProbe

`func (o *ServicePortRequestPortsInner) GetHasReadinessProbe() bool`

GetHasReadinessProbe returns the HasReadinessProbe field if non-nil, zero value otherwise.

### GetHasReadinessProbeOk

`func (o *ServicePortRequestPortsInner) GetHasReadinessProbeOk() (*bool, bool)`

GetHasReadinessProbeOk returns a tuple with the HasReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReadinessProbe

`func (o *ServicePortRequestPortsInner) SetHasReadinessProbe(v bool)`

SetHasReadinessProbe sets HasReadinessProbe field to given value.

### HasHasReadinessProbe

`func (o *ServicePortRequestPortsInner) HasHasReadinessProbe() bool`

HasHasReadinessProbe returns a boolean if a field has been set.

### GetHasLivenessProbe

`func (o *ServicePortRequestPortsInner) GetHasLivenessProbe() bool`

GetHasLivenessProbe returns the HasLivenessProbe field if non-nil, zero value otherwise.

### GetHasLivenessProbeOk

`func (o *ServicePortRequestPortsInner) GetHasLivenessProbeOk() (*bool, bool)`

GetHasLivenessProbeOk returns a tuple with the HasLivenessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLivenessProbe

`func (o *ServicePortRequestPortsInner) SetHasLivenessProbe(v bool)`

SetHasLivenessProbe sets HasLivenessProbe field to given value.

### HasHasLivenessProbe

`func (o *ServicePortRequestPortsInner) HasHasLivenessProbe() bool`

HasHasLivenessProbe returns a boolean if a field has been set.

### GetProtocol

`func (o *ServicePortRequestPortsInner) GetProtocol() PortProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ServicePortRequestPortsInner) GetProtocolOk() (*PortProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ServicePortRequestPortsInner) SetProtocol(v PortProtocolEnum)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ServicePortRequestPortsInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


