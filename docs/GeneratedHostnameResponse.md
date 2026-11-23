# GeneratedHostnameResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** |  | 
**PortId** | **string** | Id of the port the hostname reaches. | 
**PortName** | **string** |  | 
**Path** | **string** | The public path of the hostname. | 
**Kind** | [**GeneratedHostnameKindEnum**](GeneratedHostnameKindEnum.md) |  | 

## Methods

### NewGeneratedHostnameResponse

`func NewGeneratedHostnameResponse(hostname string, portId string, portName string, path string, kind GeneratedHostnameKindEnum, ) *GeneratedHostnameResponse`

NewGeneratedHostnameResponse instantiates a new GeneratedHostnameResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratedHostnameResponseWithDefaults

`func NewGeneratedHostnameResponseWithDefaults() *GeneratedHostnameResponse`

NewGeneratedHostnameResponseWithDefaults instantiates a new GeneratedHostnameResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *GeneratedHostnameResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GeneratedHostnameResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GeneratedHostnameResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPortId

`func (o *GeneratedHostnameResponse) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *GeneratedHostnameResponse) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *GeneratedHostnameResponse) SetPortId(v string)`

SetPortId sets PortId field to given value.


### GetPortName

`func (o *GeneratedHostnameResponse) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *GeneratedHostnameResponse) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *GeneratedHostnameResponse) SetPortName(v string)`

SetPortName sets PortName field to given value.


### GetPath

`func (o *GeneratedHostnameResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GeneratedHostnameResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GeneratedHostnameResponse) SetPath(v string)`

SetPath sets Path field to given value.


### GetKind

`func (o *GeneratedHostnameResponse) GetKind() GeneratedHostnameKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GeneratedHostnameResponse) GetKindOk() (*GeneratedHostnameKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GeneratedHostnameResponse) SetKind(v GeneratedHostnameKindEnum)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


