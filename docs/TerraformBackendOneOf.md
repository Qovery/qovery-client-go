# TerraformBackendOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kubernetes** | **map[string]interface{}** | Kubernetes-specific backend configuration | 

## Methods

### NewTerraformBackendOneOf

`func NewTerraformBackendOneOf(kubernetes map[string]interface{}, ) *TerraformBackendOneOf`

NewTerraformBackendOneOf instantiates a new TerraformBackendOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformBackendOneOfWithDefaults

`func NewTerraformBackendOneOfWithDefaults() *TerraformBackendOneOf`

NewTerraformBackendOneOfWithDefaults instantiates a new TerraformBackendOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubernetes

`func (o *TerraformBackendOneOf) GetKubernetes() map[string]interface{}`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *TerraformBackendOneOf) GetKubernetesOk() (*map[string]interface{}, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *TerraformBackendOneOf) SetKubernetes(v map[string]interface{})`

SetKubernetes sets Kubernetes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


