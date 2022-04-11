# ServiceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ServiceTypeEnum**](ServiceTypeEnum.md) |  | [optional] 
**Name** | Pointer to **string** | name of the service | [optional] 
**Id** | **string** | uuid of the associated service (application, database, job, gateway...) | 
**DeployedCommitId** | Pointer to **string** | Git commit ID corresponding to the deployed version of the application | [optional] 
**LastUpdatedBy** | Pointer to **string** | uuid of the user that made the last update | [optional] 
**ConsumedResourcesInPercent** | Pointer to **float32** | global overview of resources consumption of the service | [optional] 
**ServiceTypology** | Pointer to **string** | describes the typology of service (container, postgresl, redis...) | [optional] 
**ServiceVersion** | Pointer to **string** | for databases this field exposes the database version | [optional] 
**ToUpdate** | Pointer to **bool** |  | [optional] 

## Methods

### NewServiceAllOf

`func NewServiceAllOf(id string, ) *ServiceAllOf`

NewServiceAllOf instantiates a new ServiceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAllOfWithDefaults

`func NewServiceAllOfWithDefaults() *ServiceAllOf`

NewServiceAllOfWithDefaults instantiates a new ServiceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceAllOf) GetType() ServiceTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceAllOf) GetTypeOk() (*ServiceTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceAllOf) SetType(v ServiceTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ServiceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ServiceAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetDeployedCommitId

`func (o *ServiceAllOf) GetDeployedCommitId() string`

GetDeployedCommitId returns the DeployedCommitId field if non-nil, zero value otherwise.

### GetDeployedCommitIdOk

`func (o *ServiceAllOf) GetDeployedCommitIdOk() (*string, bool)`

GetDeployedCommitIdOk returns a tuple with the DeployedCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitId

`func (o *ServiceAllOf) SetDeployedCommitId(v string)`

SetDeployedCommitId sets DeployedCommitId field to given value.

### HasDeployedCommitId

`func (o *ServiceAllOf) HasDeployedCommitId() bool`

HasDeployedCommitId returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *ServiceAllOf) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *ServiceAllOf) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *ServiceAllOf) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *ServiceAllOf) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetConsumedResourcesInPercent

`func (o *ServiceAllOf) GetConsumedResourcesInPercent() float32`

GetConsumedResourcesInPercent returns the ConsumedResourcesInPercent field if non-nil, zero value otherwise.

### GetConsumedResourcesInPercentOk

`func (o *ServiceAllOf) GetConsumedResourcesInPercentOk() (*float32, bool)`

GetConsumedResourcesInPercentOk returns a tuple with the ConsumedResourcesInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedResourcesInPercent

`func (o *ServiceAllOf) SetConsumedResourcesInPercent(v float32)`

SetConsumedResourcesInPercent sets ConsumedResourcesInPercent field to given value.

### HasConsumedResourcesInPercent

`func (o *ServiceAllOf) HasConsumedResourcesInPercent() bool`

HasConsumedResourcesInPercent returns a boolean if a field has been set.

### GetServiceTypology

`func (o *ServiceAllOf) GetServiceTypology() string`

GetServiceTypology returns the ServiceTypology field if non-nil, zero value otherwise.

### GetServiceTypologyOk

`func (o *ServiceAllOf) GetServiceTypologyOk() (*string, bool)`

GetServiceTypologyOk returns a tuple with the ServiceTypology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypology

`func (o *ServiceAllOf) SetServiceTypology(v string)`

SetServiceTypology sets ServiceTypology field to given value.

### HasServiceTypology

`func (o *ServiceAllOf) HasServiceTypology() bool`

HasServiceTypology returns a boolean if a field has been set.

### GetServiceVersion

`func (o *ServiceAllOf) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *ServiceAllOf) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *ServiceAllOf) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *ServiceAllOf) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetToUpdate

`func (o *ServiceAllOf) GetToUpdate() bool`

GetToUpdate returns the ToUpdate field if non-nil, zero value otherwise.

### GetToUpdateOk

`func (o *ServiceAllOf) GetToUpdateOk() (*bool, bool)`

GetToUpdateOk returns a tuple with the ToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUpdate

`func (o *ServiceAllOf) SetToUpdate(v bool)`

SetToUpdate sets ToUpdate field to given value.

### HasToUpdate

`func (o *ServiceAllOf) HasToUpdate() bool`

HasToUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


