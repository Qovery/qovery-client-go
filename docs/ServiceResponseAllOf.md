# ServiceResponseAllOf

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

### NewServiceResponseAllOf

`func NewServiceResponseAllOf(id string, ) *ServiceResponseAllOf`

NewServiceResponseAllOf instantiates a new ServiceResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceResponseAllOfWithDefaults

`func NewServiceResponseAllOfWithDefaults() *ServiceResponseAllOf`

NewServiceResponseAllOfWithDefaults instantiates a new ServiceResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceResponseAllOf) GetType() ServiceTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceResponseAllOf) GetTypeOk() (*ServiceTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceResponseAllOf) SetType(v ServiceTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceResponseAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ServiceResponseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceResponseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceResponseAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceResponseAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ServiceResponseAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceResponseAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceResponseAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetDeployedCommitId

`func (o *ServiceResponseAllOf) GetDeployedCommitId() string`

GetDeployedCommitId returns the DeployedCommitId field if non-nil, zero value otherwise.

### GetDeployedCommitIdOk

`func (o *ServiceResponseAllOf) GetDeployedCommitIdOk() (*string, bool)`

GetDeployedCommitIdOk returns a tuple with the DeployedCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitId

`func (o *ServiceResponseAllOf) SetDeployedCommitId(v string)`

SetDeployedCommitId sets DeployedCommitId field to given value.

### HasDeployedCommitId

`func (o *ServiceResponseAllOf) HasDeployedCommitId() bool`

HasDeployedCommitId returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *ServiceResponseAllOf) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *ServiceResponseAllOf) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *ServiceResponseAllOf) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *ServiceResponseAllOf) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetConsumedResourcesInPercent

`func (o *ServiceResponseAllOf) GetConsumedResourcesInPercent() float32`

GetConsumedResourcesInPercent returns the ConsumedResourcesInPercent field if non-nil, zero value otherwise.

### GetConsumedResourcesInPercentOk

`func (o *ServiceResponseAllOf) GetConsumedResourcesInPercentOk() (*float32, bool)`

GetConsumedResourcesInPercentOk returns a tuple with the ConsumedResourcesInPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedResourcesInPercent

`func (o *ServiceResponseAllOf) SetConsumedResourcesInPercent(v float32)`

SetConsumedResourcesInPercent sets ConsumedResourcesInPercent field to given value.

### HasConsumedResourcesInPercent

`func (o *ServiceResponseAllOf) HasConsumedResourcesInPercent() bool`

HasConsumedResourcesInPercent returns a boolean if a field has been set.

### GetServiceTypology

`func (o *ServiceResponseAllOf) GetServiceTypology() string`

GetServiceTypology returns the ServiceTypology field if non-nil, zero value otherwise.

### GetServiceTypologyOk

`func (o *ServiceResponseAllOf) GetServiceTypologyOk() (*string, bool)`

GetServiceTypologyOk returns a tuple with the ServiceTypology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypology

`func (o *ServiceResponseAllOf) SetServiceTypology(v string)`

SetServiceTypology sets ServiceTypology field to given value.

### HasServiceTypology

`func (o *ServiceResponseAllOf) HasServiceTypology() bool`

HasServiceTypology returns a boolean if a field has been set.

### GetServiceVersion

`func (o *ServiceResponseAllOf) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *ServiceResponseAllOf) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *ServiceResponseAllOf) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *ServiceResponseAllOf) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetToUpdate

`func (o *ServiceResponseAllOf) GetToUpdate() bool`

GetToUpdate returns the ToUpdate field if non-nil, zero value otherwise.

### GetToUpdateOk

`func (o *ServiceResponseAllOf) GetToUpdateOk() (*bool, bool)`

GetToUpdateOk returns a tuple with the ToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUpdate

`func (o *ServiceResponseAllOf) SetToUpdate(v bool)`

SetToUpdate sets ToUpdate field to given value.

### HasToUpdate

`func (o *ServiceResponseAllOf) HasToUpdate() bool`

HasToUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


