# KubernetesServiceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to [**[]KubernetesServicePort**](KubernetesServicePort.md) |  | [optional] 
**Selectors** | Pointer to [**[]KubernetesSelector**](KubernetesSelector.md) |  | [optional] 

## Methods

### NewKubernetesServiceSpec

`func NewKubernetesServiceSpec() *KubernetesServiceSpec`

NewKubernetesServiceSpec instantiates a new KubernetesServiceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesServiceSpecWithDefaults

`func NewKubernetesServiceSpecWithDefaults() *KubernetesServiceSpec`

NewKubernetesServiceSpecWithDefaults instantiates a new KubernetesServiceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *KubernetesServiceSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesServiceSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesServiceSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubernetesServiceSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPorts

`func (o *KubernetesServiceSpec) GetPorts() []KubernetesServicePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *KubernetesServiceSpec) GetPortsOk() (*[]KubernetesServicePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *KubernetesServiceSpec) SetPorts(v []KubernetesServicePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *KubernetesServiceSpec) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetSelectors

`func (o *KubernetesServiceSpec) GetSelectors() []KubernetesSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *KubernetesServiceSpec) GetSelectorsOk() (*[]KubernetesSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *KubernetesServiceSpec) SetSelectors(v []KubernetesSelector)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *KubernetesServiceSpec) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


