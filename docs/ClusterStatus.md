# ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ClusterStateEnum**](ClusterStateEnum.md) |  | [optional] 
**IsDeployed** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterStatus

`func NewClusterStatus() *ClusterStatus`

NewClusterStatus instantiates a new ClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusWithDefaults

`func NewClusterStatusWithDefaults() *ClusterStatus`

NewClusterStatusWithDefaults instantiates a new ClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ClusterStatus) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterStatus) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterStatus) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ClusterStatus) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterStatus) GetStatus() ClusterStateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterStatus) GetStatusOk() (*ClusterStateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterStatus) SetStatus(v ClusterStateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsDeployed

`func (o *ClusterStatus) GetIsDeployed() bool`

GetIsDeployed returns the IsDeployed field if non-nil, zero value otherwise.

### GetIsDeployedOk

`func (o *ClusterStatus) GetIsDeployedOk() (*bool, bool)`

GetIsDeployedOk returns a tuple with the IsDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeployed

`func (o *ClusterStatus) SetIsDeployed(v bool)`

SetIsDeployed sets IsDeployed field to given value.

### HasIsDeployed

`func (o *ClusterStatus) HasIsDeployed() bool`

HasIsDeployed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


