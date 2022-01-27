# ServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | type of the service (application, database, job, gateway...) | [optional] 
**Name** | Pointer to **string** | name of the service | [optional] 
**Id** | **string** | uuid of the associated service (application, database, job, gateway...) | 
**DeployedCommitId** | Pointer to **string** | Git commit ID corresponding to the deployed version of the application | [optional] 
**LastUpdatedBy** | Pointer to **string** | uuid of the user that made the last update | [optional] 
**ConsumedResourcesInPercent** | Pointer to **float32** | global overview of resources consumption of the service | [optional] 
**ServiceTypology** | Pointer to **string** | describes the typology of service (container, postgresl, redis...) | [optional] 
**ServiceVersion** | Pointer to **string** | for databases this field exposes the database version | [optional] 
**ToUpdate** | Pointer to **bool** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewServiceResponse

`func NewServiceResponse(id string, createdAt time.Time, ) *ServiceResponse`

NewServiceResponse instantiates a new ServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceResponseWithDefaults

`func NewServiceResponseWithDefaults() *ServiceResponse`

NewServiceResponseWithDefaults instantiates a new ServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ServiceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDeployedCommitId

`func (o *ServiceResponse) GetDeployedCommitId() string`

GetDeployedCommitId returns the DeployedCommitId field if non-nil, zero value otherwise.

### GetDeployedCommitIdOk

`func (o *ServiceResponse) GetDeployedCommitIdOk() (*string, bool)`

GetDeployedCommitIdOk returns a tuple with the DeployedCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitId

`func (o *ServiceResponse) SetDeployedCommitId(v string)`

SetDeployedCommitId sets DeployedCommitId field to given value.

### HasDeployedCommitId

`func (o *ServiceResponse) HasDeployedCommitId() bool`

HasDeployedCommitId returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *ServiceResponse) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *ServiceResponse) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *ServiceResponse) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *ServiceResponse) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetConsumedResourcesInPercent

`func (o *ServiceResponse) GetConsumedResourcesInPercent() float32`

GetConsumedResourcesInPercent returns the ConsumedResourcesInPercent field if non-nil, zero value otherwise.

### GetConsumedResourcesInPercentOk

`func (o *ServiceResponse) GetConsumedResourcesInPercentOk() (*float32, bool)`

GetConsumedResourcesInPercentOk returns a tuple with the ConsumedResourcesInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedResourcesInPercent

`func (o *ServiceResponse) SetConsumedResourcesInPercent(v float32)`

SetConsumedResourcesInPercent sets ConsumedResourcesInPercent field to given value.

### HasConsumedResourcesInPercent

`func (o *ServiceResponse) HasConsumedResourcesInPercent() bool`

HasConsumedResourcesInPercent returns a boolean if a field has been set.

### GetServiceTypology

`func (o *ServiceResponse) GetServiceTypology() string`

GetServiceTypology returns the ServiceTypology field if non-nil, zero value otherwise.

### GetServiceTypologyOk

`func (o *ServiceResponse) GetServiceTypologyOk() (*string, bool)`

GetServiceTypologyOk returns a tuple with the ServiceTypology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypology

`func (o *ServiceResponse) SetServiceTypology(v string)`

SetServiceTypology sets ServiceTypology field to given value.

### HasServiceTypology

`func (o *ServiceResponse) HasServiceTypology() bool`

HasServiceTypology returns a boolean if a field has been set.

### GetServiceVersion

`func (o *ServiceResponse) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *ServiceResponse) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *ServiceResponse) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *ServiceResponse) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetToUpdate

`func (o *ServiceResponse) GetToUpdate() bool`

GetToUpdate returns the ToUpdate field if non-nil, zero value otherwise.

### GetToUpdateOk

`func (o *ServiceResponse) GetToUpdateOk() (*bool, bool)`

GetToUpdateOk returns a tuple with the ToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUpdate

`func (o *ServiceResponse) SetToUpdate(v bool)`

SetToUpdate sets ToUpdate field to given value.

### HasToUpdate

`func (o *ServiceResponse) HasToUpdate() bool`

HasToUpdate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ServiceResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


