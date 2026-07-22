# BlueprintDeploymentAckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Identifier of the blueprint deployment, used to track its status | 
**ExecutionId** | **string** | Engine execution identifier for this deployment | 

## Methods

### NewBlueprintDeploymentAckResponse

`func NewBlueprintDeploymentAckResponse(deploymentId string, executionId string, ) *BlueprintDeploymentAckResponse`

NewBlueprintDeploymentAckResponse instantiates a new BlueprintDeploymentAckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintDeploymentAckResponseWithDefaults

`func NewBlueprintDeploymentAckResponseWithDefaults() *BlueprintDeploymentAckResponse`

NewBlueprintDeploymentAckResponseWithDefaults instantiates a new BlueprintDeploymentAckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *BlueprintDeploymentAckResponse) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *BlueprintDeploymentAckResponse) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *BlueprintDeploymentAckResponse) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetExecutionId

`func (o *BlueprintDeploymentAckResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BlueprintDeploymentAckResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BlueprintDeploymentAckResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


