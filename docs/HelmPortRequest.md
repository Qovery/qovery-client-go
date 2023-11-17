# HelmPortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | Pointer to [**[]HelmPortRequestPortsInner**](HelmPortRequestPortsInner.md) |  | [optional] 

## Methods

### NewHelmPortRequest

`func NewHelmPortRequest() *HelmPortRequest`

NewHelmPortRequest instantiates a new HelmPortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmPortRequestWithDefaults

`func NewHelmPortRequestWithDefaults() *HelmPortRequest`

NewHelmPortRequestWithDefaults instantiates a new HelmPortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *HelmPortRequest) GetPorts() []HelmPortRequestPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *HelmPortRequest) GetPortsOk() (*[]HelmPortRequestPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *HelmPortRequest) SetPorts(v []HelmPortRequestPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *HelmPortRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


