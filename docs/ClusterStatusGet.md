# ClusterStatusGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StateEnum**](StateEnum.md) |  | [optional] 
**IsDeployed** | Pointer to **bool** |  | [optional] 
**LastExecutionId** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterStatusGet

`func NewClusterStatusGet() *ClusterStatusGet`

NewClusterStatusGet instantiates a new ClusterStatusGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusGetWithDefaults

`func NewClusterStatusGetWithDefaults() *ClusterStatusGet`

NewClusterStatusGetWithDefaults instantiates a new ClusterStatusGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ClusterStatusGet) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterStatusGet) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterStatusGet) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ClusterStatusGet) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterStatusGet) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterStatusGet) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterStatusGet) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterStatusGet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsDeployed

`func (o *ClusterStatusGet) GetIsDeployed() bool`

GetIsDeployed returns the IsDeployed field if non-nil, zero value otherwise.

### GetIsDeployedOk

`func (o *ClusterStatusGet) GetIsDeployedOk() (*bool, bool)`

GetIsDeployedOk returns a tuple with the IsDeployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeployed

`func (o *ClusterStatusGet) SetIsDeployed(v bool)`

SetIsDeployed sets IsDeployed field to given value.

### HasIsDeployed

`func (o *ClusterStatusGet) HasIsDeployed() bool`

HasIsDeployed returns a boolean if a field has been set.

### GetLastExecutionId

`func (o *ClusterStatusGet) GetLastExecutionId() string`

GetLastExecutionId returns the LastExecutionId field if non-nil, zero value otherwise.

### GetLastExecutionIdOk

`func (o *ClusterStatusGet) GetLastExecutionIdOk() (*string, bool)`

GetLastExecutionIdOk returns a tuple with the LastExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionId

`func (o *ClusterStatusGet) SetLastExecutionId(v string)`

SetLastExecutionId sets LastExecutionId field to given value.

### HasLastExecutionId

`func (o *ClusterStatusGet) HasLastExecutionId() bool`

HasLastExecutionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


