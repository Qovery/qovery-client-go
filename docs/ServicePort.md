# ServicePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**InternalPort** | **int32** | The listening port of your service. | 
**ExternalPort** | Pointer to **int32** | The exposed port for your service. This is optional. If not set a default port will be used. | [optional] 
**PubliclyAccessible** | **bool** | Expose the port to the world | 
**IsDefault** | Pointer to **bool** | is the default port to use for domain | [optional] 
**HasReadinessProbe** | Pointer to **bool** | is port to use for probe readiness check | [optional] 
**HasLivenessProbe** | Pointer to **bool** | is port to use for probe liveness check | [optional] 
**Protocol** | [**PortProtocolEnum**](PortProtocolEnum.md) |  | [default to PORTPROTOCOLENUM_HTTP]

## Methods

### NewServicePort

`func NewServicePort(id string, internalPort int32, publiclyAccessible bool, protocol PortProtocolEnum, ) *ServicePort`

NewServicePort instantiates a new ServicePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePortWithDefaults

`func NewServicePortWithDefaults() *ServicePort`

NewServicePortWithDefaults instantiates a new ServicePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServicePort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServicePort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServicePort) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ServicePort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServicePort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalPort

`func (o *ServicePort) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *ServicePort) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *ServicePort) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.


### GetExternalPort

`func (o *ServicePort) GetExternalPort() int32`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *ServicePort) GetExternalPortOk() (*int32, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *ServicePort) SetExternalPort(v int32)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *ServicePort) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetPubliclyAccessible

`func (o *ServicePort) GetPubliclyAccessible() bool`

GetPubliclyAccessible returns the PubliclyAccessible field if non-nil, zero value otherwise.

### GetPubliclyAccessibleOk

`func (o *ServicePort) GetPubliclyAccessibleOk() (*bool, bool)`

GetPubliclyAccessibleOk returns a tuple with the PubliclyAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubliclyAccessible

`func (o *ServicePort) SetPubliclyAccessible(v bool)`

SetPubliclyAccessible sets PubliclyAccessible field to given value.


### GetIsDefault

`func (o *ServicePort) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ServicePort) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ServicePort) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ServicePort) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetHasReadinessProbe

`func (o *ServicePort) GetHasReadinessProbe() bool`

GetHasReadinessProbe returns the HasReadinessProbe field if non-nil, zero value otherwise.

### GetHasReadinessProbeOk

`func (o *ServicePort) GetHasReadinessProbeOk() (*bool, bool)`

GetHasReadinessProbeOk returns a tuple with the HasReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReadinessProbe

`func (o *ServicePort) SetHasReadinessProbe(v bool)`

SetHasReadinessProbe sets HasReadinessProbe field to given value.

### HasHasReadinessProbe

`func (o *ServicePort) HasHasReadinessProbe() bool`

HasHasReadinessProbe returns a boolean if a field has been set.

### GetHasLivenessProbe

`func (o *ServicePort) GetHasLivenessProbe() bool`

GetHasLivenessProbe returns the HasLivenessProbe field if non-nil, zero value otherwise.

### GetHasLivenessProbeOk

`func (o *ServicePort) GetHasLivenessProbeOk() (*bool, bool)`

GetHasLivenessProbeOk returns a tuple with the HasLivenessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLivenessProbe

`func (o *ServicePort) SetHasLivenessProbe(v bool)`

SetHasLivenessProbe sets HasLivenessProbe field to given value.

### HasHasLivenessProbe

`func (o *ServicePort) HasHasLivenessProbe() bool`

HasHasLivenessProbe returns a boolean if a field has been set.

### GetProtocol

`func (o *ServicePort) GetProtocol() PortProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ServicePort) GetProtocolOk() (*PortProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ServicePort) SetProtocol(v PortProtocolEnum)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


