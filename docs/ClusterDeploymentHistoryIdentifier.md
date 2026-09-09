# ClusterDeploymentHistoryIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** |  | 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ClusterId** | **string** |  | 

## Methods

### NewClusterDeploymentHistoryIdentifier

`func NewClusterDeploymentHistoryIdentifier(deploymentId string, clusterId string, ) *ClusterDeploymentHistoryIdentifier`

NewClusterDeploymentHistoryIdentifier instantiates a new ClusterDeploymentHistoryIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDeploymentHistoryIdentifierWithDefaults

`func NewClusterDeploymentHistoryIdentifierWithDefaults() *ClusterDeploymentHistoryIdentifier`

NewClusterDeploymentHistoryIdentifierWithDefaults instantiates a new ClusterDeploymentHistoryIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *ClusterDeploymentHistoryIdentifier) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ClusterDeploymentHistoryIdentifier) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ClusterDeploymentHistoryIdentifier) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetExecutionId

`func (o *ClusterDeploymentHistoryIdentifier) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ClusterDeploymentHistoryIdentifier) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ClusterDeploymentHistoryIdentifier) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ClusterDeploymentHistoryIdentifier) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *ClusterDeploymentHistoryIdentifier) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *ClusterDeploymentHistoryIdentifier) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetClusterId

`func (o *ClusterDeploymentHistoryIdentifier) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterDeploymentHistoryIdentifier) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterDeploymentHistoryIdentifier) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


