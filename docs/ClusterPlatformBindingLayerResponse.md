# ClusterPlatformBindingLayerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Status** | [**PlatformLayerResolutionStatus**](PlatformLayerResolutionStatus.md) |  | 
**Reason** | **string** |  | 
**ComponentKeys** | **[]string** |  | 

## Methods

### NewClusterPlatformBindingLayerResponse

`func NewClusterPlatformBindingLayerResponse(key string, status PlatformLayerResolutionStatus, reason string, componentKeys []string, ) *ClusterPlatformBindingLayerResponse`

NewClusterPlatformBindingLayerResponse instantiates a new ClusterPlatformBindingLayerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPlatformBindingLayerResponseWithDefaults

`func NewClusterPlatformBindingLayerResponseWithDefaults() *ClusterPlatformBindingLayerResponse`

NewClusterPlatformBindingLayerResponseWithDefaults instantiates a new ClusterPlatformBindingLayerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ClusterPlatformBindingLayerResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ClusterPlatformBindingLayerResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ClusterPlatformBindingLayerResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetStatus

`func (o *ClusterPlatformBindingLayerResponse) GetStatus() PlatformLayerResolutionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterPlatformBindingLayerResponse) GetStatusOk() (*PlatformLayerResolutionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterPlatformBindingLayerResponse) SetStatus(v PlatformLayerResolutionStatus)`

SetStatus sets Status field to given value.


### GetReason

`func (o *ClusterPlatformBindingLayerResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterPlatformBindingLayerResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterPlatformBindingLayerResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetComponentKeys

`func (o *ClusterPlatformBindingLayerResponse) GetComponentKeys() []string`

GetComponentKeys returns the ComponentKeys field if non-nil, zero value otherwise.

### GetComponentKeysOk

`func (o *ClusterPlatformBindingLayerResponse) GetComponentKeysOk() (*[]string, bool)`

GetComponentKeysOk returns a tuple with the ComponentKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentKeys

`func (o *ClusterPlatformBindingLayerResponse) SetComponentKeys(v []string)`

SetComponentKeys sets ComponentKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


