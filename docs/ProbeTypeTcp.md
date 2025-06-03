# ProbeTypeTcp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProbeTypeTcp

`func NewProbeTypeTcp() *ProbeTypeTcp`

NewProbeTypeTcp instantiates a new ProbeTypeTcp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProbeTypeTcpWithDefaults

`func NewProbeTypeTcpWithDefaults() *ProbeTypeTcp`

NewProbeTypeTcpWithDefaults instantiates a new ProbeTypeTcp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ProbeTypeTcp) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ProbeTypeTcp) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ProbeTypeTcp) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ProbeTypeTcp) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetHost

`func (o *ProbeTypeTcp) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ProbeTypeTcp) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ProbeTypeTcp) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ProbeTypeTcp) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *ProbeTypeTcp) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *ProbeTypeTcp) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


