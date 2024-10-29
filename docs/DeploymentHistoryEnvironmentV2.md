# DeploymentHistoryEnvironmentV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | [**DeploymentHistoryEnvironmentV2Identifier**](DeploymentHistoryEnvironmentV2Identifier.md) |  | 
**AuditingData** | [**DeploymentHistoryAuditingData**](DeploymentHistoryAuditingData.md) |  | 
**Status** | [**StateEnum**](StateEnum.md) |  | 
**TriggerAction** | [**DeploymentHistoryTriggerAction**](DeploymentHistoryTriggerAction.md) |  | 
**TotalDuration** | **string** |  | 
**Stages** | [**[]DeploymentHistoryStage**](DeploymentHistoryStage.md) |  | 

## Methods

### NewDeploymentHistoryEnvironmentV2

`func NewDeploymentHistoryEnvironmentV2(identifier DeploymentHistoryEnvironmentV2Identifier, auditingData DeploymentHistoryAuditingData, status StateEnum, triggerAction DeploymentHistoryTriggerAction, totalDuration string, stages []DeploymentHistoryStage, ) *DeploymentHistoryEnvironmentV2`

NewDeploymentHistoryEnvironmentV2 instantiates a new DeploymentHistoryEnvironmentV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryEnvironmentV2WithDefaults

`func NewDeploymentHistoryEnvironmentV2WithDefaults() *DeploymentHistoryEnvironmentV2`

NewDeploymentHistoryEnvironmentV2WithDefaults instantiates a new DeploymentHistoryEnvironmentV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *DeploymentHistoryEnvironmentV2) GetIdentifier() DeploymentHistoryEnvironmentV2Identifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DeploymentHistoryEnvironmentV2) GetIdentifierOk() (*DeploymentHistoryEnvironmentV2Identifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DeploymentHistoryEnvironmentV2) SetIdentifier(v DeploymentHistoryEnvironmentV2Identifier)`

SetIdentifier sets Identifier field to given value.


### GetAuditingData

`func (o *DeploymentHistoryEnvironmentV2) GetAuditingData() DeploymentHistoryAuditingData`

GetAuditingData returns the AuditingData field if non-nil, zero value otherwise.

### GetAuditingDataOk

`func (o *DeploymentHistoryEnvironmentV2) GetAuditingDataOk() (*DeploymentHistoryAuditingData, bool)`

GetAuditingDataOk returns a tuple with the AuditingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditingData

`func (o *DeploymentHistoryEnvironmentV2) SetAuditingData(v DeploymentHistoryAuditingData)`

SetAuditingData sets AuditingData field to given value.


### GetStatus

`func (o *DeploymentHistoryEnvironmentV2) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryEnvironmentV2) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryEnvironmentV2) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.


### GetTriggerAction

`func (o *DeploymentHistoryEnvironmentV2) GetTriggerAction() DeploymentHistoryTriggerAction`

GetTriggerAction returns the TriggerAction field if non-nil, zero value otherwise.

### GetTriggerActionOk

`func (o *DeploymentHistoryEnvironmentV2) GetTriggerActionOk() (*DeploymentHistoryTriggerAction, bool)`

GetTriggerActionOk returns a tuple with the TriggerAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAction

`func (o *DeploymentHistoryEnvironmentV2) SetTriggerAction(v DeploymentHistoryTriggerAction)`

SetTriggerAction sets TriggerAction field to given value.


### GetTotalDuration

`func (o *DeploymentHistoryEnvironmentV2) GetTotalDuration() string`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *DeploymentHistoryEnvironmentV2) GetTotalDurationOk() (*string, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *DeploymentHistoryEnvironmentV2) SetTotalDuration(v string)`

SetTotalDuration sets TotalDuration field to given value.


### GetStages

`func (o *DeploymentHistoryEnvironmentV2) GetStages() []DeploymentHistoryStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *DeploymentHistoryEnvironmentV2) GetStagesOk() (*[]DeploymentHistoryStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *DeploymentHistoryEnvironmentV2) SetStages(v []DeploymentHistoryStage)`

SetStages sets Stages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


