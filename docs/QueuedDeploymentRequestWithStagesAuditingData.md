# QueuedDeploymentRequestWithStagesAuditingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggeredBy** | **string** |  | 
**Origin** | Pointer to [**OrganizationEventOrigin**](OrganizationEventOrigin.md) |  | [optional] 

## Methods

### NewQueuedDeploymentRequestWithStagesAuditingData

`func NewQueuedDeploymentRequestWithStagesAuditingData(triggeredBy string, ) *QueuedDeploymentRequestWithStagesAuditingData`

NewQueuedDeploymentRequestWithStagesAuditingData instantiates a new QueuedDeploymentRequestWithStagesAuditingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuedDeploymentRequestWithStagesAuditingDataWithDefaults

`func NewQueuedDeploymentRequestWithStagesAuditingDataWithDefaults() *QueuedDeploymentRequestWithStagesAuditingData`

NewQueuedDeploymentRequestWithStagesAuditingDataWithDefaults instantiates a new QueuedDeploymentRequestWithStagesAuditingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggeredBy

`func (o *QueuedDeploymentRequestWithStagesAuditingData) GetTriggeredBy() string`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *QueuedDeploymentRequestWithStagesAuditingData) GetTriggeredByOk() (*string, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *QueuedDeploymentRequestWithStagesAuditingData) SetTriggeredBy(v string)`

SetTriggeredBy sets TriggeredBy field to given value.


### GetOrigin

`func (o *QueuedDeploymentRequestWithStagesAuditingData) GetOrigin() OrganizationEventOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *QueuedDeploymentRequestWithStagesAuditingData) GetOriginOk() (*OrganizationEventOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *QueuedDeploymentRequestWithStagesAuditingData) SetOrigin(v OrganizationEventOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *QueuedDeploymentRequestWithStagesAuditingData) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


