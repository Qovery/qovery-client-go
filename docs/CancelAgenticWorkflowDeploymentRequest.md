# CancelAgenticWorkflowDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceCancel** | Pointer to **bool** | When false, the engine stops the deployment at the next safe point, letting the workflow&#39;s job terminate on its own. When true it stops waiting for that job and abandons it immediately. | [optional] [default to false]

## Methods

### NewCancelAgenticWorkflowDeploymentRequest

`func NewCancelAgenticWorkflowDeploymentRequest() *CancelAgenticWorkflowDeploymentRequest`

NewCancelAgenticWorkflowDeploymentRequest instantiates a new CancelAgenticWorkflowDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelAgenticWorkflowDeploymentRequestWithDefaults

`func NewCancelAgenticWorkflowDeploymentRequestWithDefaults() *CancelAgenticWorkflowDeploymentRequest`

NewCancelAgenticWorkflowDeploymentRequestWithDefaults instantiates a new CancelAgenticWorkflowDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceCancel

`func (o *CancelAgenticWorkflowDeploymentRequest) GetForceCancel() bool`

GetForceCancel returns the ForceCancel field if non-nil, zero value otherwise.

### GetForceCancelOk

`func (o *CancelAgenticWorkflowDeploymentRequest) GetForceCancelOk() (*bool, bool)`

GetForceCancelOk returns a tuple with the ForceCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceCancel

`func (o *CancelAgenticWorkflowDeploymentRequest) SetForceCancel(v bool)`

SetForceCancel sets ForceCancel field to given value.

### HasForceCancel

`func (o *CancelAgenticWorkflowDeploymentRequest) HasForceCancel() bool`

HasForceCancel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


