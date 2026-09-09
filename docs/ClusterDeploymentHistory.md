# ClusterDeploymentHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | [**ClusterDeploymentHistoryIdentifier**](ClusterDeploymentHistoryIdentifier.md) |  | 
**AuditingData** | [**ClusterDeploymentHistoryAuditingData**](ClusterDeploymentHistoryAuditingData.md) |  | 
**Status** | [**StateEnum**](StateEnum.md) |  | 
**ActionStatus** | [**DeploymentHistoryActionStatus**](DeploymentHistoryActionStatus.md) |  | 
**TriggerAction** | [**DeploymentHistoryTriggerAction**](DeploymentHistoryTriggerAction.md) |  | 
**Reason** | **string** |  | 
**TotalDuration** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewClusterDeploymentHistory

`func NewClusterDeploymentHistory(identifier ClusterDeploymentHistoryIdentifier, auditingData ClusterDeploymentHistoryAuditingData, status StateEnum, actionStatus DeploymentHistoryActionStatus, triggerAction DeploymentHistoryTriggerAction, reason string, ) *ClusterDeploymentHistory`

NewClusterDeploymentHistory instantiates a new ClusterDeploymentHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDeploymentHistoryWithDefaults

`func NewClusterDeploymentHistoryWithDefaults() *ClusterDeploymentHistory`

NewClusterDeploymentHistoryWithDefaults instantiates a new ClusterDeploymentHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *ClusterDeploymentHistory) GetIdentifier() ClusterDeploymentHistoryIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ClusterDeploymentHistory) GetIdentifierOk() (*ClusterDeploymentHistoryIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ClusterDeploymentHistory) SetIdentifier(v ClusterDeploymentHistoryIdentifier)`

SetIdentifier sets Identifier field to given value.


### GetAuditingData

`func (o *ClusterDeploymentHistory) GetAuditingData() ClusterDeploymentHistoryAuditingData`

GetAuditingData returns the AuditingData field if non-nil, zero value otherwise.

### GetAuditingDataOk

`func (o *ClusterDeploymentHistory) GetAuditingDataOk() (*ClusterDeploymentHistoryAuditingData, bool)`

GetAuditingDataOk returns a tuple with the AuditingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditingData

`func (o *ClusterDeploymentHistory) SetAuditingData(v ClusterDeploymentHistoryAuditingData)`

SetAuditingData sets AuditingData field to given value.


### GetStatus

`func (o *ClusterDeploymentHistory) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterDeploymentHistory) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterDeploymentHistory) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.


### GetActionStatus

`func (o *ClusterDeploymentHistory) GetActionStatus() DeploymentHistoryActionStatus`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *ClusterDeploymentHistory) GetActionStatusOk() (*DeploymentHistoryActionStatus, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *ClusterDeploymentHistory) SetActionStatus(v DeploymentHistoryActionStatus)`

SetActionStatus sets ActionStatus field to given value.


### GetTriggerAction

`func (o *ClusterDeploymentHistory) GetTriggerAction() DeploymentHistoryTriggerAction`

GetTriggerAction returns the TriggerAction field if non-nil, zero value otherwise.

### GetTriggerActionOk

`func (o *ClusterDeploymentHistory) GetTriggerActionOk() (*DeploymentHistoryTriggerAction, bool)`

GetTriggerActionOk returns a tuple with the TriggerAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAction

`func (o *ClusterDeploymentHistory) SetTriggerAction(v DeploymentHistoryTriggerAction)`

SetTriggerAction sets TriggerAction field to given value.


### GetReason

`func (o *ClusterDeploymentHistory) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterDeploymentHistory) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterDeploymentHistory) SetReason(v string)`

SetReason sets Reason field to given value.


### GetTotalDuration

`func (o *ClusterDeploymentHistory) GetTotalDuration() string`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *ClusterDeploymentHistory) GetTotalDurationOk() (*string, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *ClusterDeploymentHistory) SetTotalDuration(v string)`

SetTotalDuration sets TotalDuration field to given value.

### HasTotalDuration

`func (o *ClusterDeploymentHistory) HasTotalDuration() bool`

HasTotalDuration returns a boolean if a field has been set.

### SetTotalDurationNil

`func (o *ClusterDeploymentHistory) SetTotalDurationNil(b bool)`

 SetTotalDurationNil sets the value for TotalDuration to be an explicit nil

### UnsetTotalDuration
`func (o *ClusterDeploymentHistory) UnsetTotalDuration()`

UnsetTotalDuration ensures that no value is present for TotalDuration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


