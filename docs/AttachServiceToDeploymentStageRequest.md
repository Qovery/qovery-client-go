# AttachServiceToDeploymentStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSkipped** | Pointer to **bool** | When true, marks the service as skipped (excluded from environment-level bulk deployments) while keeping it in the specified stage. When false or omitted, the service is moved to the specified stage and un-skipped. | [optional] 

## Methods

### NewAttachServiceToDeploymentStageRequest

`func NewAttachServiceToDeploymentStageRequest() *AttachServiceToDeploymentStageRequest`

NewAttachServiceToDeploymentStageRequest instantiates a new AttachServiceToDeploymentStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachServiceToDeploymentStageRequestWithDefaults

`func NewAttachServiceToDeploymentStageRequestWithDefaults() *AttachServiceToDeploymentStageRequest`

NewAttachServiceToDeploymentStageRequestWithDefaults instantiates a new AttachServiceToDeploymentStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSkipped

`func (o *AttachServiceToDeploymentStageRequest) GetIsSkipped() bool`

GetIsSkipped returns the IsSkipped field if non-nil, zero value otherwise.

### GetIsSkippedOk

`func (o *AttachServiceToDeploymentStageRequest) GetIsSkippedOk() (*bool, bool)`

GetIsSkippedOk returns a tuple with the IsSkipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSkipped

`func (o *AttachServiceToDeploymentStageRequest) SetIsSkipped(v bool)`

SetIsSkipped sets IsSkipped field to given value.

### HasIsSkipped

`func (o *AttachServiceToDeploymentStageRequest) HasIsSkipped() bool`

HasIsSkipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


