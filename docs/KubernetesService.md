# KubernetesService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**KubernetesMetadata**](KubernetesMetadata.md) |  | 
**ServiceSpec** | [**KubernetesServiceSpec**](KubernetesServiceSpec.md) |  | 

## Methods

### NewKubernetesService

`func NewKubernetesService(metadata KubernetesMetadata, serviceSpec KubernetesServiceSpec, ) *KubernetesService`

NewKubernetesService instantiates a new KubernetesService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesServiceWithDefaults

`func NewKubernetesServiceWithDefaults() *KubernetesService`

NewKubernetesServiceWithDefaults instantiates a new KubernetesService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *KubernetesService) GetMetadata() KubernetesMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubernetesService) GetMetadataOk() (*KubernetesMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubernetesService) SetMetadata(v KubernetesMetadata)`

SetMetadata sets Metadata field to given value.


### GetServiceSpec

`func (o *KubernetesService) GetServiceSpec() KubernetesServiceSpec`

GetServiceSpec returns the ServiceSpec field if non-nil, zero value otherwise.

### GetServiceSpecOk

`func (o *KubernetesService) GetServiceSpecOk() (*KubernetesServiceSpec, bool)`

GetServiceSpecOk returns a tuple with the ServiceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSpec

`func (o *KubernetesService) SetServiceSpec(v KubernetesServiceSpec)`

SetServiceSpec sets ServiceSpec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


