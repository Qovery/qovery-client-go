# HelmPortResponseBase

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

## Methods

### NewHelmPortResponseBase

`func NewHelmPortResponseBase(id string, portType string, internalPort int32, protocol HelmPortProtocolEnum, ) *HelmPortResponseBase`

NewHelmPortResponseBase instantiates a new HelmPortResponseBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmPortResponseBaseWithDefaults

`func NewHelmPortResponseBaseWithDefaults() *HelmPortResponseBase`

NewHelmPortResponseBaseWithDefaults instantiates a new HelmPortResponseBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HelmPortResponseBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HelmPortResponseBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HelmPortResponseBase) SetId(v string)`

SetId sets Id field to given value.


### GetPortType

`func (o *HelmPortResponseBase) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *HelmPortResponseBase) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *HelmPortResponseBase) SetPortType(v string)`

SetPortType sets PortType field to given value.


### GetName

`func (o *HelmPortResponseBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmPortResponseBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmPortResponseBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HelmPortResponseBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalPort

`func (o *HelmPortResponseBase) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *HelmPortResponseBase) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *HelmPortResponseBase) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.


### GetExternalPort

`func (o *HelmPortResponseBase) GetExternalPort() int32`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *HelmPortResponseBase) GetExternalPortOk() (*int32, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *HelmPortResponseBase) SetExternalPort(v int32)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *HelmPortResponseBase) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetNamespace

`func (o *HelmPortResponseBase) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HelmPortResponseBase) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HelmPortResponseBase) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *HelmPortResponseBase) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetProtocol

`func (o *HelmPortResponseBase) GetProtocol() HelmPortProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HelmPortResponseBase) GetProtocolOk() (*HelmPortProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HelmPortResponseBase) SetProtocol(v HelmPortProtocolEnum)`

SetProtocol sets Protocol field to given value.


### GetIsDefault

`func (o *HelmPortResponseBase) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *HelmPortResponseBase) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *HelmPortResponseBase) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *HelmPortResponseBase) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


