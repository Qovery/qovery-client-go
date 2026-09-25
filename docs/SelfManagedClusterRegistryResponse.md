# SelfManagedClusterRegistryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | [**ContainerRegistryKindEnum**](ContainerRegistryKindEnum.md) |  | 

## Methods

### NewSelfManagedClusterRegistryResponse

`func NewSelfManagedClusterRegistryResponse(id string, kind ContainerRegistryKindEnum, ) *SelfManagedClusterRegistryResponse`

NewSelfManagedClusterRegistryResponse instantiates a new SelfManagedClusterRegistryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfManagedClusterRegistryResponseWithDefaults

`func NewSelfManagedClusterRegistryResponseWithDefaults() *SelfManagedClusterRegistryResponse`

NewSelfManagedClusterRegistryResponseWithDefaults instantiates a new SelfManagedClusterRegistryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SelfManagedClusterRegistryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelfManagedClusterRegistryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelfManagedClusterRegistryResponse) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *SelfManagedClusterRegistryResponse) GetKind() ContainerRegistryKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SelfManagedClusterRegistryResponse) GetKindOk() (*ContainerRegistryKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SelfManagedClusterRegistryResponse) SetKind(v ContainerRegistryKindEnum)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


