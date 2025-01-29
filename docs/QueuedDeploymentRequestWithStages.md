# QueuedDeploymentRequestWithStages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | [**QueuedDeploymentRequestWithStagesIdentifier**](QueuedDeploymentRequestWithStagesIdentifier.md) |  | 
**AuditingData** | [**QueuedDeploymentRequestWithStagesAuditingData**](QueuedDeploymentRequestWithStagesAuditingData.md) |  | 
**TriggerAction** | [**DeploymentHistoryTriggerAction**](DeploymentHistoryTriggerAction.md) |  | 
**Stages** | [**[]QueuedDeploymentRequestWithStagesStagesInner**](QueuedDeploymentRequestWithStagesStagesInner.md) |  | 

## Methods

### NewQueuedDeploymentRequestWithStages

`func NewQueuedDeploymentRequestWithStages(identifier QueuedDeploymentRequestWithStagesIdentifier, auditingData QueuedDeploymentRequestWithStagesAuditingData, triggerAction DeploymentHistoryTriggerAction, stages []QueuedDeploymentRequestWithStagesStagesInner, ) *QueuedDeploymentRequestWithStages`

NewQueuedDeploymentRequestWithStages instantiates a new QueuedDeploymentRequestWithStages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuedDeploymentRequestWithStagesWithDefaults

`func NewQueuedDeploymentRequestWithStagesWithDefaults() *QueuedDeploymentRequestWithStages`

NewQueuedDeploymentRequestWithStagesWithDefaults instantiates a new QueuedDeploymentRequestWithStages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *QueuedDeploymentRequestWithStages) GetIdentifier() QueuedDeploymentRequestWithStagesIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *QueuedDeploymentRequestWithStages) GetIdentifierOk() (*QueuedDeploymentRequestWithStagesIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *QueuedDeploymentRequestWithStages) SetIdentifier(v QueuedDeploymentRequestWithStagesIdentifier)`

SetIdentifier sets Identifier field to given value.


### GetAuditingData

`func (o *QueuedDeploymentRequestWithStages) GetAuditingData() QueuedDeploymentRequestWithStagesAuditingData`

GetAuditingData returns the AuditingData field if non-nil, zero value otherwise.

### GetAuditingDataOk

`func (o *QueuedDeploymentRequestWithStages) GetAuditingDataOk() (*QueuedDeploymentRequestWithStagesAuditingData, bool)`

GetAuditingDataOk returns a tuple with the AuditingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditingData

`func (o *QueuedDeploymentRequestWithStages) SetAuditingData(v QueuedDeploymentRequestWithStagesAuditingData)`

SetAuditingData sets AuditingData field to given value.


### GetTriggerAction

`func (o *QueuedDeploymentRequestWithStages) GetTriggerAction() DeploymentHistoryTriggerAction`

GetTriggerAction returns the TriggerAction field if non-nil, zero value otherwise.

### GetTriggerActionOk

`func (o *QueuedDeploymentRequestWithStages) GetTriggerActionOk() (*DeploymentHistoryTriggerAction, bool)`

GetTriggerActionOk returns a tuple with the TriggerAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAction

`func (o *QueuedDeploymentRequestWithStages) SetTriggerAction(v DeploymentHistoryTriggerAction)`

SetTriggerAction sets TriggerAction field to given value.


### GetStages

`func (o *QueuedDeploymentRequestWithStages) GetStages() []QueuedDeploymentRequestWithStagesStagesInner`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *QueuedDeploymentRequestWithStages) GetStagesOk() (*[]QueuedDeploymentRequestWithStagesStagesInner, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *QueuedDeploymentRequestWithStages) SetStages(v []QueuedDeploymentRequestWithStagesStagesInner)`

SetStages sets Stages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


