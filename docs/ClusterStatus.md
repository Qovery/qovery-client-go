# ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ClusterStateEnum**](ClusterStateEnum.md) |  | [optional] 
**IsDeployed** | Pointer to **bool** |  | [optional] 
**NextK8sAvailableVersion** | Pointer to **NullableString** |  | [optional] 
**LastExecutionId** | Pointer to **string** |  | [optional] 
**ClusterLock** | Pointer to [**ClusterLock**](ClusterLock.md) |  | [optional] 
**LastDeploymentDate** | Pointer to **time.Time** |  | [optional] 

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

### GetNextK8sAvailableVersion

`func (o *ClusterStatus) GetNextK8sAvailableVersion() string`

GetNextK8sAvailableVersion returns the NextK8sAvailableVersion field if non-nil, zero value otherwise.

### GetNextK8sAvailableVersionOk

`func (o *ClusterStatus) GetNextK8sAvailableVersionOk() (*string, bool)`

GetNextK8sAvailableVersionOk returns a tuple with the NextK8sAvailableVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextK8sAvailableVersion

`func (o *ClusterStatus) SetNextK8sAvailableVersion(v string)`

SetNextK8sAvailableVersion sets NextK8sAvailableVersion field to given value.

### HasNextK8sAvailableVersion

`func (o *ClusterStatus) HasNextK8sAvailableVersion() bool`

HasNextK8sAvailableVersion returns a boolean if a field has been set.

### SetNextK8sAvailableVersionNil

`func (o *ClusterStatus) SetNextK8sAvailableVersionNil(b bool)`

 SetNextK8sAvailableVersionNil sets the value for NextK8sAvailableVersion to be an explicit nil

### UnsetNextK8sAvailableVersion
`func (o *ClusterStatus) UnsetNextK8sAvailableVersion()`

UnsetNextK8sAvailableVersion ensures that no value is present for NextK8sAvailableVersion, not even an explicit nil
### GetLastExecutionId

`func (o *ClusterStatus) GetLastExecutionId() string`

GetLastExecutionId returns the LastExecutionId field if non-nil, zero value otherwise.

### GetLastExecutionIdOk

`func (o *ClusterStatus) GetLastExecutionIdOk() (*string, bool)`

GetLastExecutionIdOk returns a tuple with the LastExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionId

`func (o *ClusterStatus) SetLastExecutionId(v string)`

SetLastExecutionId sets LastExecutionId field to given value.

### HasLastExecutionId

`func (o *ClusterStatus) HasLastExecutionId() bool`

HasLastExecutionId returns a boolean if a field has been set.

### GetClusterLock

`func (o *ClusterStatus) GetClusterLock() ClusterLock`

GetClusterLock returns the ClusterLock field if non-nil, zero value otherwise.

### GetClusterLockOk

`func (o *ClusterStatus) GetClusterLockOk() (*ClusterLock, bool)`

GetClusterLockOk returns a tuple with the ClusterLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterLock

`func (o *ClusterStatus) SetClusterLock(v ClusterLock)`

SetClusterLock sets ClusterLock field to given value.

### HasClusterLock

`func (o *ClusterStatus) HasClusterLock() bool`

HasClusterLock returns a boolean if a field has been set.

### GetLastDeploymentDate

`func (o *ClusterStatus) GetLastDeploymentDate() time.Time`

GetLastDeploymentDate returns the LastDeploymentDate field if non-nil, zero value otherwise.

### GetLastDeploymentDateOk

`func (o *ClusterStatus) GetLastDeploymentDateOk() (*time.Time, bool)`

GetLastDeploymentDateOk returns a tuple with the LastDeploymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploymentDate

`func (o *ClusterStatus) SetLastDeploymentDate(v time.Time)`

SetLastDeploymentDate sets LastDeploymentDate field to given value.

### HasLastDeploymentDate

`func (o *ClusterStatus) HasLastDeploymentDate() bool`

HasLastDeploymentDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


