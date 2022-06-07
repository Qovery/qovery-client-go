# ApplicationPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | Pointer to [**[]ApplicationPortPortsInner**](ApplicationPortPortsInner.md) |  | [optional] 

## Methods

### NewApplicationPort

`func NewApplicationPort() *ApplicationPort`

NewApplicationPort instantiates a new ApplicationPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPortWithDefaults

`func NewApplicationPortWithDefaults() *ApplicationPort`

NewApplicationPortWithDefaults instantiates a new ApplicationPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *ApplicationPort) GetPorts() []ApplicationPortPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ApplicationPort) GetPortsOk() (*[]ApplicationPortPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ApplicationPort) SetPorts(v []ApplicationPortPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ApplicationPort) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


