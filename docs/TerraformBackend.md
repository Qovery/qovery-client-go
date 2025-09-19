# TerraformBackend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kubernetes** | Pointer to **map[string]interface{}** | Kubernetes-specific backend configuration | [optional] 
**UserProvided** | Pointer to **map[string]interface{}** | User-provided backend configuration | [optional] 

## Methods

### NewTerraformBackend

`func NewTerraformBackend() *TerraformBackend`

NewTerraformBackend instantiates a new TerraformBackend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformBackendWithDefaults

`func NewTerraformBackendWithDefaults() *TerraformBackend`

NewTerraformBackendWithDefaults instantiates a new TerraformBackend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubernetes

`func (o *TerraformBackend) GetKubernetes() map[string]interface{}`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *TerraformBackend) GetKubernetesOk() (*map[string]interface{}, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *TerraformBackend) SetKubernetes(v map[string]interface{})`

SetKubernetes sets Kubernetes field to given value.

### HasKubernetes

`func (o *TerraformBackend) HasKubernetes() bool`

HasKubernetes returns a boolean if a field has been set.

### GetUserProvided

`func (o *TerraformBackend) GetUserProvided() map[string]interface{}`

GetUserProvided returns the UserProvided field if non-nil, zero value otherwise.

### GetUserProvidedOk

`func (o *TerraformBackend) GetUserProvidedOk() (*map[string]interface{}, bool)`

GetUserProvidedOk returns a tuple with the UserProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProvided

`func (o *TerraformBackend) SetUserProvided(v map[string]interface{})`

SetUserProvided sets UserProvided field to given value.

### HasUserProvided

`func (o *TerraformBackend) HasUserProvided() bool`

HasUserProvided returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


