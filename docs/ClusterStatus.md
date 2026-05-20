# ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**Status** | [**ClusterStateEnum**](ClusterStateEnum.md) |  | 
**IsDeployed** | **bool** |  | 
**NextK8sAvailableVersion** | Pointer to **NullableString** |  | [optional] 
**LastExecutionId** | Pointer to **NullableString** |  | [optional] 
**ClusterLock** | Pointer to [**NullableClusterLock**](ClusterLock.md) |  | [optional] 
**LastDeploymentDate** | Pointer to **NullableTime** |  | [optional] 
**Reason** | [**DeploymentInfraReason**](DeploymentInfraReason.md) |  | 

## Methods

### NewClusterStatus

`func NewClusterStatus(clusterId string, status ClusterStateEnum, isDeployed bool, reason DeploymentInfraReason, ) *ClusterStatus`

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

### SetLastExecutionIdNil

`func (o *ClusterStatus) SetLastExecutionIdNil(b bool)`

 SetLastExecutionIdNil sets the value for LastExecutionId to be an explicit nil

### UnsetLastExecutionId
`func (o *ClusterStatus) UnsetLastExecutionId()`

UnsetLastExecutionId ensures that no value is present for LastExecutionId, not even an explicit nil
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

### SetClusterLockNil

`func (o *ClusterStatus) SetClusterLockNil(b bool)`

 SetClusterLockNil sets the value for ClusterLock to be an explicit nil

### UnsetClusterLock
`func (o *ClusterStatus) UnsetClusterLock()`

UnsetClusterLock ensures that no value is present for ClusterLock, not even an explicit nil
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

### SetLastDeploymentDateNil

`func (o *ClusterStatus) SetLastDeploymentDateNil(b bool)`

 SetLastDeploymentDateNil sets the value for LastDeploymentDate to be an explicit nil

### UnsetLastDeploymentDate
`func (o *ClusterStatus) UnsetLastDeploymentDate()`

UnsetLastDeploymentDate ensures that no value is present for LastDeploymentDate, not even an explicit nil
### GetReason

`func (o *ClusterStatus) GetReason() DeploymentInfraReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterStatus) GetReasonOk() (*DeploymentInfraReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterStatus) SetReason(v DeploymentInfraReason)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


