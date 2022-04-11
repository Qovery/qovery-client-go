# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | uuid of the associated service (application, database, job, gateway...) | 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Type** | Pointer to [**ServiceTypeEnum**](ServiceTypeEnum.md) |  | [optional] 
**Name** | Pointer to **string** | name of the service | [optional] 
**DeployedCommitId** | Pointer to **string** | Git commit ID corresponding to the deployed version of the application | [optional] 
**LastUpdatedBy** | Pointer to **string** | uuid of the user that made the last update | [optional] 
**ConsumedResourcesInPercent** | Pointer to **float32** | global overview of resources consumption of the service | [optional] 
**ServiceTypology** | Pointer to **string** | describes the typology of service (container, postgresl, redis...) | [optional] 
**ServiceVersion** | Pointer to **string** | for databases this field exposes the database version | [optional] 
**ToUpdate** | Pointer to **bool** |  | [optional] 

## Methods

### NewService

`func NewService(id string, createdAt time.Time, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Service) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Service) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Service) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Service) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Service) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Service) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Service) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *Service) GetType() ServiceTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Service) GetTypeOk() (*ServiceTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Service) SetType(v ServiceTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *Service) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Service) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeployedCommitId

`func (o *Service) GetDeployedCommitId() string`

GetDeployedCommitId returns the DeployedCommitId field if non-nil, zero value otherwise.

### GetDeployedCommitIdOk

`func (o *Service) GetDeployedCommitIdOk() (*string, bool)`

GetDeployedCommitIdOk returns a tuple with the DeployedCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitId

`func (o *Service) SetDeployedCommitId(v string)`

SetDeployedCommitId sets DeployedCommitId field to given value.

### HasDeployedCommitId

`func (o *Service) HasDeployedCommitId() bool`

HasDeployedCommitId returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *Service) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *Service) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *Service) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *Service) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetConsumedResourcesInPercent

`func (o *Service) GetConsumedResourcesInPercent() float32`

GetConsumedResourcesInPercent returns the ConsumedResourcesInPercent field if non-nil, zero value otherwise.

### GetConsumedResourcesInPercentOk

`func (o *Service) GetConsumedResourcesInPercentOk() (*float32, bool)`

GetConsumedResourcesInPercentOk returns a tuple with the ConsumedResourcesInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedResourcesInPercent

`func (o *Service) SetConsumedResourcesInPercent(v float32)`

SetConsumedResourcesInPercent sets ConsumedResourcesInPercent field to given value.

### HasConsumedResourcesInPercent

`func (o *Service) HasConsumedResourcesInPercent() bool`

HasConsumedResourcesInPercent returns a boolean if a field has been set.

### GetServiceTypology

`func (o *Service) GetServiceTypology() string`

GetServiceTypology returns the ServiceTypology field if non-nil, zero value otherwise.

### GetServiceTypologyOk

`func (o *Service) GetServiceTypologyOk() (*string, bool)`

GetServiceTypologyOk returns a tuple with the ServiceTypology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypology

`func (o *Service) SetServiceTypology(v string)`

SetServiceTypology sets ServiceTypology field to given value.

### HasServiceTypology

`func (o *Service) HasServiceTypology() bool`

HasServiceTypology returns a boolean if a field has been set.

### GetServiceVersion

`func (o *Service) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *Service) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *Service) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *Service) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetToUpdate

`func (o *Service) GetToUpdate() bool`

GetToUpdate returns the ToUpdate field if non-nil, zero value otherwise.

### GetToUpdateOk

`func (o *Service) GetToUpdateOk() (*bool, bool)`

GetToUpdateOk returns a tuple with the ToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUpdate

`func (o *Service) SetToUpdate(v bool)`

SetToUpdate sets ToUpdate field to given value.

### HasToUpdate

`func (o *Service) HasToUpdate() bool`

HasToUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


