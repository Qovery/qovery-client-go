# ProbeType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**None** | Pointer to **map[string]interface{}** |  | [optional] 
**Tcp** | Pointer to [**NullableProbeTypeTcp**](ProbeTypeTcp.md) |  | [optional] 
**Http** | Pointer to [**NullableProbeTypeHttp**](ProbeTypeHttp.md) |  | [optional] 
**Exec** | Pointer to [**NullableProbeTypeExec**](ProbeTypeExec.md) |  | [optional] 
**Grpc** | Pointer to [**NullableProbeTypeGrpc**](ProbeTypeGrpc.md) |  | [optional] 

## Methods

### NewProbeType

`func NewProbeType() *ProbeType`

NewProbeType instantiates a new ProbeType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProbeTypeWithDefaults

`func NewProbeTypeWithDefaults() *ProbeType`

NewProbeTypeWithDefaults instantiates a new ProbeType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNone

`func (o *ProbeType) GetNone() map[string]interface{}`

GetNone returns the None field if non-nil, zero value otherwise.

### GetNoneOk

`func (o *ProbeType) GetNoneOk() (*map[string]interface{}, bool)`

GetNoneOk returns a tuple with the None field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNone

`func (o *ProbeType) SetNone(v map[string]interface{})`

SetNone sets None field to given value.

### HasNone

`func (o *ProbeType) HasNone() bool`

HasNone returns a boolean if a field has been set.

### SetNoneNil

`func (o *ProbeType) SetNoneNil(b bool)`

 SetNoneNil sets the value for None to be an explicit nil

### UnsetNone
`func (o *ProbeType) UnsetNone()`

UnsetNone ensures that no value is present for None, not even an explicit nil
### GetTcp

`func (o *ProbeType) GetTcp() ProbeTypeTcp`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *ProbeType) GetTcpOk() (*ProbeTypeTcp, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *ProbeType) SetTcp(v ProbeTypeTcp)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *ProbeType) HasTcp() bool`

HasTcp returns a boolean if a field has been set.

### SetTcpNil

`func (o *ProbeType) SetTcpNil(b bool)`

 SetTcpNil sets the value for Tcp to be an explicit nil

### UnsetTcp
`func (o *ProbeType) UnsetTcp()`

UnsetTcp ensures that no value is present for Tcp, not even an explicit nil
### GetHttp

`func (o *ProbeType) GetHttp() ProbeTypeHttp`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *ProbeType) GetHttpOk() (*ProbeTypeHttp, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *ProbeType) SetHttp(v ProbeTypeHttp)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *ProbeType) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### SetHttpNil

`func (o *ProbeType) SetHttpNil(b bool)`

 SetHttpNil sets the value for Http to be an explicit nil

### UnsetHttp
`func (o *ProbeType) UnsetHttp()`

UnsetHttp ensures that no value is present for Http, not even an explicit nil
### GetExec

`func (o *ProbeType) GetExec() ProbeTypeExec`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *ProbeType) GetExecOk() (*ProbeTypeExec, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *ProbeType) SetExec(v ProbeTypeExec)`

SetExec sets Exec field to given value.

### HasExec

`func (o *ProbeType) HasExec() bool`

HasExec returns a boolean if a field has been set.

### SetExecNil

`func (o *ProbeType) SetExecNil(b bool)`

 SetExecNil sets the value for Exec to be an explicit nil

### UnsetExec
`func (o *ProbeType) UnsetExec()`

UnsetExec ensures that no value is present for Exec, not even an explicit nil
### GetGrpc

`func (o *ProbeType) GetGrpc() ProbeTypeGrpc`

GetGrpc returns the Grpc field if non-nil, zero value otherwise.

### GetGrpcOk

`func (o *ProbeType) GetGrpcOk() (*ProbeTypeGrpc, bool)`

GetGrpcOk returns a tuple with the Grpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpc

`func (o *ProbeType) SetGrpc(v ProbeTypeGrpc)`

SetGrpc sets Grpc field to given value.

### HasGrpc

`func (o *ProbeType) HasGrpc() bool`

HasGrpc returns a boolean if a field has been set.

### SetGrpcNil

`func (o *ProbeType) SetGrpcNil(b bool)`

 SetGrpcNil sets the value for Grpc to be an explicit nil

### UnsetGrpc
`func (o *ProbeType) UnsetGrpc()`

UnsetGrpc ensures that no value is present for Grpc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


