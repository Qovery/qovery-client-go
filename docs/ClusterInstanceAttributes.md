# ClusterInstanceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceCategory** | Pointer to **string** | The category of the instance. | [optional] 
**InstanceGeneration** | Pointer to **int32** | The generation of the instance. | [optional] 
**InstanceFamily** | Pointer to **string** | The family or series of the instance. | [optional] 
**InstanceSize** | Pointer to **string** | Specifies the size of the instance within its family. | [optional] 
**WithGpu** | Pointer to **bool** | The instance has gpu. | [optional] 
**MeetsResourceReqs** | Pointer to **bool** | The instance has sufficient resources to be chosen as a standalone instance in a cluster. | [optional] 

## Methods

### NewClusterInstanceAttributes

`func NewClusterInstanceAttributes() *ClusterInstanceAttributes`

NewClusterInstanceAttributes instantiates a new ClusterInstanceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInstanceAttributesWithDefaults

`func NewClusterInstanceAttributesWithDefaults() *ClusterInstanceAttributes`

NewClusterInstanceAttributesWithDefaults instantiates a new ClusterInstanceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceCategory

`func (o *ClusterInstanceAttributes) GetInstanceCategory() string`

GetInstanceCategory returns the InstanceCategory field if non-nil, zero value otherwise.

### GetInstanceCategoryOk

`func (o *ClusterInstanceAttributes) GetInstanceCategoryOk() (*string, bool)`

GetInstanceCategoryOk returns a tuple with the InstanceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCategory

`func (o *ClusterInstanceAttributes) SetInstanceCategory(v string)`

SetInstanceCategory sets InstanceCategory field to given value.

### HasInstanceCategory

`func (o *ClusterInstanceAttributes) HasInstanceCategory() bool`

HasInstanceCategory returns a boolean if a field has been set.

### GetInstanceGeneration

`func (o *ClusterInstanceAttributes) GetInstanceGeneration() int32`

GetInstanceGeneration returns the InstanceGeneration field if non-nil, zero value otherwise.

### GetInstanceGenerationOk

`func (o *ClusterInstanceAttributes) GetInstanceGenerationOk() (*int32, bool)`

GetInstanceGenerationOk returns a tuple with the InstanceGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGeneration

`func (o *ClusterInstanceAttributes) SetInstanceGeneration(v int32)`

SetInstanceGeneration sets InstanceGeneration field to given value.

### HasInstanceGeneration

`func (o *ClusterInstanceAttributes) HasInstanceGeneration() bool`

HasInstanceGeneration returns a boolean if a field has been set.

### GetInstanceFamily

`func (o *ClusterInstanceAttributes) GetInstanceFamily() string`

GetInstanceFamily returns the InstanceFamily field if non-nil, zero value otherwise.

### GetInstanceFamilyOk

`func (o *ClusterInstanceAttributes) GetInstanceFamilyOk() (*string, bool)`

GetInstanceFamilyOk returns a tuple with the InstanceFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceFamily

`func (o *ClusterInstanceAttributes) SetInstanceFamily(v string)`

SetInstanceFamily sets InstanceFamily field to given value.

### HasInstanceFamily

`func (o *ClusterInstanceAttributes) HasInstanceFamily() bool`

HasInstanceFamily returns a boolean if a field has been set.

### GetInstanceSize

`func (o *ClusterInstanceAttributes) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *ClusterInstanceAttributes) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *ClusterInstanceAttributes) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### HasInstanceSize

`func (o *ClusterInstanceAttributes) HasInstanceSize() bool`

HasInstanceSize returns a boolean if a field has been set.

### GetWithGpu

`func (o *ClusterInstanceAttributes) GetWithGpu() bool`

GetWithGpu returns the WithGpu field if non-nil, zero value otherwise.

### GetWithGpuOk

`func (o *ClusterInstanceAttributes) GetWithGpuOk() (*bool, bool)`

GetWithGpuOk returns a tuple with the WithGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithGpu

`func (o *ClusterInstanceAttributes) SetWithGpu(v bool)`

SetWithGpu sets WithGpu field to given value.

### HasWithGpu

`func (o *ClusterInstanceAttributes) HasWithGpu() bool`

HasWithGpu returns a boolean if a field has been set.

### GetMeetsResourceReqs

`func (o *ClusterInstanceAttributes) GetMeetsResourceReqs() bool`

GetMeetsResourceReqs returns the MeetsResourceReqs field if non-nil, zero value otherwise.

### GetMeetsResourceReqsOk

`func (o *ClusterInstanceAttributes) GetMeetsResourceReqsOk() (*bool, bool)`

GetMeetsResourceReqsOk returns a tuple with the MeetsResourceReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetsResourceReqs

`func (o *ClusterInstanceAttributes) SetMeetsResourceReqs(v bool)`

SetMeetsResourceReqs sets MeetsResourceReqs field to given value.

### HasMeetsResourceReqs

`func (o *ClusterInstanceAttributes) HasMeetsResourceReqs() bool`

HasMeetsResourceReqs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


