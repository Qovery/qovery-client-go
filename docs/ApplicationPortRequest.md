# ApplicationPortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | Pointer to [**[]ApplicationPortRequestPorts**](ApplicationPortRequestPorts.md) |  | [optional] 

## Methods

### NewApplicationPortRequest

`func NewApplicationPortRequest() *ApplicationPortRequest`

NewApplicationPortRequest instantiates a new ApplicationPortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPortRequestWithDefaults

`func NewApplicationPortRequestWithDefaults() *ApplicationPortRequest`

NewApplicationPortRequestWithDefaults instantiates a new ApplicationPortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *ApplicationPortRequest) GetPorts() []ApplicationPortRequestPorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ApplicationPortRequest) GetPortsOk() (*[]ApplicationPortRequestPorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ApplicationPortRequest) SetPorts(v []ApplicationPortRequestPorts)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ApplicationPortRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


