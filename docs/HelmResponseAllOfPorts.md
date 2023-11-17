# HelmResponseAllOfPorts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**InternalPort** | **int32** | The listening port of your service. | 
**ExternalPort** | Pointer to **int32** | The exposed port for your service. This is optional. If not set a default port will be used. | [optional] 
**ServiceName** | **string** |  | 
**Namespace** | Pointer to **string** |  | [optional] 
**Protocol** | [**HelmPortProtocolEnum**](HelmPortProtocolEnum.md) |  | [default to HELMPORTPROTOCOLENUM_HTTP]

## Methods

### NewHelmResponseAllOfPorts

`func NewHelmResponseAllOfPorts(id string, internalPort int32, serviceName string, protocol HelmPortProtocolEnum, ) *HelmResponseAllOfPorts`

NewHelmResponseAllOfPorts instantiates a new HelmResponseAllOfPorts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmResponseAllOfPortsWithDefaults

`func NewHelmResponseAllOfPortsWithDefaults() *HelmResponseAllOfPorts`

NewHelmResponseAllOfPortsWithDefaults instantiates a new HelmResponseAllOfPorts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HelmResponseAllOfPorts) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HelmResponseAllOfPorts) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HelmResponseAllOfPorts) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *HelmResponseAllOfPorts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmResponseAllOfPorts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmResponseAllOfPorts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HelmResponseAllOfPorts) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalPort

`func (o *HelmResponseAllOfPorts) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *HelmResponseAllOfPorts) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *HelmResponseAllOfPorts) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.


### GetExternalPort

`func (o *HelmResponseAllOfPorts) GetExternalPort() int32`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *HelmResponseAllOfPorts) GetExternalPortOk() (*int32, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *HelmResponseAllOfPorts) SetExternalPort(v int32)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *HelmResponseAllOfPorts) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetServiceName

`func (o *HelmResponseAllOfPorts) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *HelmResponseAllOfPorts) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *HelmResponseAllOfPorts) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetNamespace

`func (o *HelmResponseAllOfPorts) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HelmResponseAllOfPorts) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HelmResponseAllOfPorts) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *HelmResponseAllOfPorts) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetProtocol

`func (o *HelmResponseAllOfPorts) GetProtocol() HelmPortProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HelmResponseAllOfPorts) GetProtocolOk() (*HelmPortProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HelmResponseAllOfPorts) SetProtocol(v HelmPortProtocolEnum)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


