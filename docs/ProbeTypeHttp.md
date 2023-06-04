# ProbeTypeHttp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] [default to "/"]
**Scheme** | Pointer to **string** |  | [optional] [default to "HTTP"]
**Port** | Pointer to **int32** |  | [optional] 

## Methods

### NewProbeTypeHttp

`func NewProbeTypeHttp() *ProbeTypeHttp`

NewProbeTypeHttp instantiates a new ProbeTypeHttp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProbeTypeHttpWithDefaults

`func NewProbeTypeHttpWithDefaults() *ProbeTypeHttp`

NewProbeTypeHttpWithDefaults instantiates a new ProbeTypeHttp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ProbeTypeHttp) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ProbeTypeHttp) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ProbeTypeHttp) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ProbeTypeHttp) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetScheme

`func (o *ProbeTypeHttp) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ProbeTypeHttp) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ProbeTypeHttp) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *ProbeTypeHttp) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetPort

`func (o *ProbeTypeHttp) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ProbeTypeHttp) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ProbeTypeHttp) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ProbeTypeHttp) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


