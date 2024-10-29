# DeploymentHistoryService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | [**DeploymentHistoryServiceIdentifier**](DeploymentHistoryServiceIdentifier.md) |  | 
**Status** | [**StateEnum**](StateEnum.md) |  | 
**AuditingData** | [**DeploymentHistoryAuditingData**](DeploymentHistoryAuditingData.md) |  | 
**Details** | [**DeploymentHistoryServiceDetails**](DeploymentHistoryServiceDetails.md) |  | 

## Methods

### NewDeploymentHistoryService

`func NewDeploymentHistoryService(identifier DeploymentHistoryServiceIdentifier, status StateEnum, auditingData DeploymentHistoryAuditingData, details DeploymentHistoryServiceDetails, ) *DeploymentHistoryService`

NewDeploymentHistoryService instantiates a new DeploymentHistoryService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryServiceWithDefaults

`func NewDeploymentHistoryServiceWithDefaults() *DeploymentHistoryService`

NewDeploymentHistoryServiceWithDefaults instantiates a new DeploymentHistoryService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *DeploymentHistoryService) GetIdentifier() DeploymentHistoryServiceIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DeploymentHistoryService) GetIdentifierOk() (*DeploymentHistoryServiceIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DeploymentHistoryService) SetIdentifier(v DeploymentHistoryServiceIdentifier)`

SetIdentifier sets Identifier field to given value.


### GetStatus

`func (o *DeploymentHistoryService) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryService) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryService) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.


### GetAuditingData

`func (o *DeploymentHistoryService) GetAuditingData() DeploymentHistoryAuditingData`

GetAuditingData returns the AuditingData field if non-nil, zero value otherwise.

### GetAuditingDataOk

`func (o *DeploymentHistoryService) GetAuditingDataOk() (*DeploymentHistoryAuditingData, bool)`

GetAuditingDataOk returns a tuple with the AuditingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditingData

`func (o *DeploymentHistoryService) SetAuditingData(v DeploymentHistoryAuditingData)`

SetAuditingData sets AuditingData field to given value.


### GetDetails

`func (o *DeploymentHistoryService) GetDetails() DeploymentHistoryServiceDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DeploymentHistoryService) GetDetailsOk() (*DeploymentHistoryServiceDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DeploymentHistoryService) SetDetails(v DeploymentHistoryServiceDetails)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


