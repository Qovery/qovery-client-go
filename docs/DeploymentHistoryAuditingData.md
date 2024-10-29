# DeploymentHistoryAuditingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**TriggeredBy** | **string** |  | 
**Origin** | Pointer to [**OrganizationEventOrigin**](OrganizationEventOrigin.md) |  | [optional] 

## Methods

### NewDeploymentHistoryAuditingData

`func NewDeploymentHistoryAuditingData(createdAt time.Time, updatedAt time.Time, triggeredBy string, ) *DeploymentHistoryAuditingData`

NewDeploymentHistoryAuditingData instantiates a new DeploymentHistoryAuditingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryAuditingDataWithDefaults

`func NewDeploymentHistoryAuditingDataWithDefaults() *DeploymentHistoryAuditingData`

NewDeploymentHistoryAuditingDataWithDefaults instantiates a new DeploymentHistoryAuditingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DeploymentHistoryAuditingData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentHistoryAuditingData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentHistoryAuditingData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeploymentHistoryAuditingData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentHistoryAuditingData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentHistoryAuditingData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTriggeredBy

`func (o *DeploymentHistoryAuditingData) GetTriggeredBy() string`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *DeploymentHistoryAuditingData) GetTriggeredByOk() (*string, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *DeploymentHistoryAuditingData) SetTriggeredBy(v string)`

SetTriggeredBy sets TriggeredBy field to given value.


### GetOrigin

`func (o *DeploymentHistoryAuditingData) GetOrigin() OrganizationEventOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DeploymentHistoryAuditingData) GetOriginOk() (*OrganizationEventOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DeploymentHistoryAuditingData) SetOrigin(v OrganizationEventOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *DeploymentHistoryAuditingData) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


