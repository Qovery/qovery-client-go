# ClusterDeploymentHistoryAuditingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Origin** | Pointer to [**NullableOrganizationEventOrigin**](OrganizationEventOrigin.md) |  | [optional] 
**TriggeredBy** | Pointer to **NullableString** | Who triggered the deployment. Null for deployments created before this metadata was recorded | [optional] 

## Methods

### NewClusterDeploymentHistoryAuditingData

`func NewClusterDeploymentHistoryAuditingData(createdAt time.Time, updatedAt time.Time, ) *ClusterDeploymentHistoryAuditingData`

NewClusterDeploymentHistoryAuditingData instantiates a new ClusterDeploymentHistoryAuditingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDeploymentHistoryAuditingDataWithDefaults

`func NewClusterDeploymentHistoryAuditingDataWithDefaults() *ClusterDeploymentHistoryAuditingData`

NewClusterDeploymentHistoryAuditingDataWithDefaults instantiates a new ClusterDeploymentHistoryAuditingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ClusterDeploymentHistoryAuditingData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterDeploymentHistoryAuditingData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterDeploymentHistoryAuditingData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ClusterDeploymentHistoryAuditingData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClusterDeploymentHistoryAuditingData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClusterDeploymentHistoryAuditingData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOrigin

`func (o *ClusterDeploymentHistoryAuditingData) GetOrigin() OrganizationEventOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ClusterDeploymentHistoryAuditingData) GetOriginOk() (*OrganizationEventOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ClusterDeploymentHistoryAuditingData) SetOrigin(v OrganizationEventOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ClusterDeploymentHistoryAuditingData) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *ClusterDeploymentHistoryAuditingData) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *ClusterDeploymentHistoryAuditingData) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetTriggeredBy

`func (o *ClusterDeploymentHistoryAuditingData) GetTriggeredBy() string`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *ClusterDeploymentHistoryAuditingData) GetTriggeredByOk() (*string, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *ClusterDeploymentHistoryAuditingData) SetTriggeredBy(v string)`

SetTriggeredBy sets TriggeredBy field to given value.

### HasTriggeredBy

`func (o *ClusterDeploymentHistoryAuditingData) HasTriggeredBy() bool`

HasTriggeredBy returns a boolean if a field has been set.

### SetTriggeredByNil

`func (o *ClusterDeploymentHistoryAuditingData) SetTriggeredByNil(b bool)`

 SetTriggeredByNil sets the value for TriggeredBy to be an explicit nil

### UnsetTriggeredBy
`func (o *ClusterDeploymentHistoryAuditingData) UnsetTriggeredBy()`

UnsetTriggeredBy ensures that no value is present for TriggeredBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


