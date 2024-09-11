# HelmPortRequestPortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**InternalPort** | **int32** | The listening port of your service. | 
**ExternalPort** | Pointer to **int32** | The exposed port for your service. This is optional. If not set a default port will be used. | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to [**HelmPortProtocolEnum**](HelmPortProtocolEnum.md) |  | [optional] [default to HELMPORTPROTOCOLENUM_HTTP]
**IsDefault** | Pointer to **bool** | is the default port to use for domain | [optional] 
**ServiceSelectors** | Pointer to [**[]KubernetesSelector**](KubernetesSelector.md) |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 

## Methods

### NewHelmPortRequestPortsInner

`func NewHelmPortRequestPortsInner(internalPort int32, ) *HelmPortRequestPortsInner`

NewHelmPortRequestPortsInner instantiates a new HelmPortRequestPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmPortRequestPortsInnerWithDefaults

`func NewHelmPortRequestPortsInnerWithDefaults() *HelmPortRequestPortsInner`

NewHelmPortRequestPortsInnerWithDefaults instantiates a new HelmPortRequestPortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HelmPortRequestPortsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmPortRequestPortsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmPortRequestPortsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HelmPortRequestPortsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalPort

`func (o *HelmPortRequestPortsInner) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *HelmPortRequestPortsInner) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *HelmPortRequestPortsInner) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.


### GetExternalPort

`func (o *HelmPortRequestPortsInner) GetExternalPort() int32`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *HelmPortRequestPortsInner) GetExternalPortOk() (*int32, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *HelmPortRequestPortsInner) SetExternalPort(v int32)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *HelmPortRequestPortsInner) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetNamespace

`func (o *HelmPortRequestPortsInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HelmPortRequestPortsInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HelmPortRequestPortsInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *HelmPortRequestPortsInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetProtocol

`func (o *HelmPortRequestPortsInner) GetProtocol() HelmPortProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HelmPortRequestPortsInner) GetProtocolOk() (*HelmPortProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HelmPortRequestPortsInner) SetProtocol(v HelmPortProtocolEnum)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *HelmPortRequestPortsInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetIsDefault

`func (o *HelmPortRequestPortsInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *HelmPortRequestPortsInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *HelmPortRequestPortsInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *HelmPortRequestPortsInner) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetServiceSelectors

`func (o *HelmPortRequestPortsInner) GetServiceSelectors() []KubernetesSelector`

GetServiceSelectors returns the ServiceSelectors field if non-nil, zero value otherwise.

### GetServiceSelectorsOk

`func (o *HelmPortRequestPortsInner) GetServiceSelectorsOk() (*[]KubernetesSelector, bool)`

GetServiceSelectorsOk returns a tuple with the ServiceSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSelectors

`func (o *HelmPortRequestPortsInner) SetServiceSelectors(v []KubernetesSelector)`

SetServiceSelectors sets ServiceSelectors field to given value.

### HasServiceSelectors

`func (o *HelmPortRequestPortsInner) HasServiceSelectors() bool`

HasServiceSelectors returns a boolean if a field has been set.

### GetServiceName

`func (o *HelmPortRequestPortsInner) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *HelmPortRequestPortsInner) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *HelmPortRequestPortsInner) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *HelmPortRequestPortsInner) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


