# HelmPortResponseWithServiceName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PortType** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**InternalPort** | **int32** | The listening port of your service. | 
**ExternalPort** | Pointer to **int32** | The exposed port for your service. This is optional. If not set a default port will be used. | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Protocol** | [**HelmPortProtocolEnum**](HelmPortProtocolEnum.md) |  | [default to HELMPORTPROTOCOLENUM_HTTP]
**IsDefault** | Pointer to **bool** | is the default port to use for domain | [optional] 
**ServiceName** | **string** |  | 

## Methods

### NewHelmPortResponseWithServiceName

`func NewHelmPortResponseWithServiceName(id string, portType string, internalPort int32, protocol HelmPortProtocolEnum, serviceName string, ) *HelmPortResponseWithServiceName`

NewHelmPortResponseWithServiceName instantiates a new HelmPortResponseWithServiceName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmPortResponseWithServiceNameWithDefaults

`func NewHelmPortResponseWithServiceNameWithDefaults() *HelmPortResponseWithServiceName`

NewHelmPortResponseWithServiceNameWithDefaults instantiates a new HelmPortResponseWithServiceName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HelmPortResponseWithServiceName) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HelmPortResponseWithServiceName) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HelmPortResponseWithServiceName) SetId(v string)`

SetId sets Id field to given value.


### GetPortType

`func (o *HelmPortResponseWithServiceName) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *HelmPortResponseWithServiceName) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *HelmPortResponseWithServiceName) SetPortType(v string)`

SetPortType sets PortType field to given value.


### GetName

`func (o *HelmPortResponseWithServiceName) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmPortResponseWithServiceName) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmPortResponseWithServiceName) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HelmPortResponseWithServiceName) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalPort

`func (o *HelmPortResponseWithServiceName) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *HelmPortResponseWithServiceName) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *HelmPortResponseWithServiceName) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.


### GetExternalPort

`func (o *HelmPortResponseWithServiceName) GetExternalPort() int32`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *HelmPortResponseWithServiceName) GetExternalPortOk() (*int32, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *HelmPortResponseWithServiceName) SetExternalPort(v int32)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *HelmPortResponseWithServiceName) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetNamespace

`func (o *HelmPortResponseWithServiceName) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HelmPortResponseWithServiceName) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HelmPortResponseWithServiceName) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *HelmPortResponseWithServiceName) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetProtocol

`func (o *HelmPortResponseWithServiceName) GetProtocol() HelmPortProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HelmPortResponseWithServiceName) GetProtocolOk() (*HelmPortProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HelmPortResponseWithServiceName) SetProtocol(v HelmPortProtocolEnum)`

SetProtocol sets Protocol field to given value.


### GetIsDefault

`func (o *HelmPortResponseWithServiceName) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *HelmPortResponseWithServiceName) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *HelmPortResponseWithServiceName) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *HelmPortResponseWithServiceName) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetServiceName

`func (o *HelmPortResponseWithServiceName) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *HelmPortResponseWithServiceName) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *HelmPortResponseWithServiceName) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


