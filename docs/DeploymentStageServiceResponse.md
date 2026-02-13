# DeploymentStageServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ServiceId** | Pointer to **string** | id of the service attached to the stage | [optional] 
**ServiceType** | Pointer to **string** | type of the service (i.e APPLICATION, JOB, DATABASE, ...) | [optional] 
**IsSkipped** | Pointer to **bool** | whether the service is excluded from environment-level deployments | [optional] 

## Methods

### NewDeploymentStageServiceResponse

`func NewDeploymentStageServiceResponse(id string, createdAt time.Time, ) *DeploymentStageServiceResponse`

NewDeploymentStageServiceResponse instantiates a new DeploymentStageServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStageServiceResponseWithDefaults

`func NewDeploymentStageServiceResponseWithDefaults() *DeploymentStageServiceResponse`

NewDeploymentStageServiceResponseWithDefaults instantiates a new DeploymentStageServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentStageServiceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentStageServiceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentStageServiceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DeploymentStageServiceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentStageServiceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentStageServiceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeploymentStageServiceResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentStageServiceResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentStageServiceResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeploymentStageServiceResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetServiceId

`func (o *DeploymentStageServiceResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeploymentStageServiceResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeploymentStageServiceResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DeploymentStageServiceResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceType

`func (o *DeploymentStageServiceResponse) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *DeploymentStageServiceResponse) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *DeploymentStageServiceResponse) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *DeploymentStageServiceResponse) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetIsSkipped

`func (o *DeploymentStageServiceResponse) GetIsSkipped() bool`

GetIsSkipped returns the IsSkipped field if non-nil, zero value otherwise.

### GetIsSkippedOk

`func (o *DeploymentStageServiceResponse) GetIsSkippedOk() (*bool, bool)`

GetIsSkippedOk returns a tuple with the IsSkipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSkipped

`func (o *DeploymentStageServiceResponse) SetIsSkipped(v bool)`

SetIsSkipped sets IsSkipped field to given value.

### HasIsSkipped

`func (o *DeploymentStageServiceResponse) HasIsSkipped() bool`

HasIsSkipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


