# DeploymentStageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Environment** | [**ReferenceObject**](ReferenceObject.md) |  | 
**Name** | Pointer to **string** | name is case insensitive | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeploymentOrder** | Pointer to **int32** | Position of the deployment stage within the environment | [optional] 
**Services** | Pointer to [**[]DeploymentStageServiceResponse**](DeploymentStageServiceResponse.md) |  | [optional] 

## Methods

### NewDeploymentStageResponse

`func NewDeploymentStageResponse(id string, createdAt time.Time, environment ReferenceObject, ) *DeploymentStageResponse`

NewDeploymentStageResponse instantiates a new DeploymentStageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStageResponseWithDefaults

`func NewDeploymentStageResponseWithDefaults() *DeploymentStageResponse`

NewDeploymentStageResponseWithDefaults instantiates a new DeploymentStageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentStageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentStageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentStageResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DeploymentStageResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentStageResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentStageResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeploymentStageResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentStageResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentStageResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeploymentStageResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *DeploymentStageResponse) GetEnvironment() ReferenceObject`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DeploymentStageResponse) GetEnvironmentOk() (*ReferenceObject, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DeploymentStageResponse) SetEnvironment(v ReferenceObject)`

SetEnvironment sets Environment field to given value.


### GetName

`func (o *DeploymentStageResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentStageResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentStageResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentStageResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DeploymentStageResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentStageResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentStageResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeploymentStageResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeploymentOrder

`func (o *DeploymentStageResponse) GetDeploymentOrder() int32`

GetDeploymentOrder returns the DeploymentOrder field if non-nil, zero value otherwise.

### GetDeploymentOrderOk

`func (o *DeploymentStageResponse) GetDeploymentOrderOk() (*int32, bool)`

GetDeploymentOrderOk returns a tuple with the DeploymentOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentOrder

`func (o *DeploymentStageResponse) SetDeploymentOrder(v int32)`

SetDeploymentOrder sets DeploymentOrder field to given value.

### HasDeploymentOrder

`func (o *DeploymentStageResponse) HasDeploymentOrder() bool`

HasDeploymentOrder returns a boolean if a field has been set.

### GetServices

`func (o *DeploymentStageResponse) GetServices() []DeploymentStageServiceResponse`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *DeploymentStageResponse) GetServicesOk() (*[]DeploymentStageServiceResponse, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *DeploymentStageResponse) SetServices(v []DeploymentStageServiceResponse)`

SetServices sets Services field to given value.

### HasServices

`func (o *DeploymentStageResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


