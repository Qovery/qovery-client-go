# KubernetesServicePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Port** | **int32** |  | 

## Methods

### NewKubernetesServicePort

`func NewKubernetesServicePort(port int32, ) *KubernetesServicePort`

NewKubernetesServicePort instantiates a new KubernetesServicePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesServicePortWithDefaults

`func NewKubernetesServicePortWithDefaults() *KubernetesServicePort`

NewKubernetesServicePortWithDefaults instantiates a new KubernetesServicePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesServicePort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesServicePort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesServicePort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesServicePort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *KubernetesServicePort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KubernetesServicePort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KubernetesServicePort) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


