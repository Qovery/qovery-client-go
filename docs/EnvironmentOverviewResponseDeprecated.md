# EnvironmentOverviewResponseDeprecated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Name** | **string** |  | 
**Mode** | [**EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | 
**Cluster** | Pointer to [**ClusterOverviewResponse**](ClusterOverviewResponse.md) |  | [optional] 
**ServiceCount** | **int32** |  | 
**DeploymentStatus** | Pointer to [**EnvironmentStatus**](EnvironmentStatus.md) |  | [optional] 

## Methods

### NewEnvironmentOverviewResponseDeprecated

`func NewEnvironmentOverviewResponseDeprecated(id string, createdAt time.Time, updatedAt time.Time, name string, mode EnvironmentModeEnum, serviceCount int32, ) *EnvironmentOverviewResponseDeprecated`

NewEnvironmentOverviewResponseDeprecated instantiates a new EnvironmentOverviewResponseDeprecated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentOverviewResponseDeprecatedWithDefaults

`func NewEnvironmentOverviewResponseDeprecatedWithDefaults() *EnvironmentOverviewResponseDeprecated`

NewEnvironmentOverviewResponseDeprecatedWithDefaults instantiates a new EnvironmentOverviewResponseDeprecated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentOverviewResponseDeprecated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentOverviewResponseDeprecated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentOverviewResponseDeprecated) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EnvironmentOverviewResponseDeprecated) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentOverviewResponseDeprecated) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentOverviewResponseDeprecated) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EnvironmentOverviewResponseDeprecated) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentOverviewResponseDeprecated) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentOverviewResponseDeprecated) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *EnvironmentOverviewResponseDeprecated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentOverviewResponseDeprecated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentOverviewResponseDeprecated) SetName(v string)`

SetName sets Name field to given value.


### GetMode

`func (o *EnvironmentOverviewResponseDeprecated) GetMode() EnvironmentModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EnvironmentOverviewResponseDeprecated) GetModeOk() (*EnvironmentModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EnvironmentOverviewResponseDeprecated) SetMode(v EnvironmentModeEnum)`

SetMode sets Mode field to given value.


### GetCluster

`func (o *EnvironmentOverviewResponseDeprecated) GetCluster() ClusterOverviewResponse`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *EnvironmentOverviewResponseDeprecated) GetClusterOk() (*ClusterOverviewResponse, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *EnvironmentOverviewResponseDeprecated) SetCluster(v ClusterOverviewResponse)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *EnvironmentOverviewResponseDeprecated) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetServiceCount

`func (o *EnvironmentOverviewResponseDeprecated) GetServiceCount() int32`

GetServiceCount returns the ServiceCount field if non-nil, zero value otherwise.

### GetServiceCountOk

`func (o *EnvironmentOverviewResponseDeprecated) GetServiceCountOk() (*int32, bool)`

GetServiceCountOk returns a tuple with the ServiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCount

`func (o *EnvironmentOverviewResponseDeprecated) SetServiceCount(v int32)`

SetServiceCount sets ServiceCount field to given value.


### GetDeploymentStatus

`func (o *EnvironmentOverviewResponseDeprecated) GetDeploymentStatus() EnvironmentStatus`

GetDeploymentStatus returns the DeploymentStatus field if non-nil, zero value otherwise.

### GetDeploymentStatusOk

`func (o *EnvironmentOverviewResponseDeprecated) GetDeploymentStatusOk() (*EnvironmentStatus, bool)`

GetDeploymentStatusOk returns a tuple with the DeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentStatus

`func (o *EnvironmentOverviewResponseDeprecated) SetDeploymentStatus(v EnvironmentStatus)`

SetDeploymentStatus sets DeploymentStatus field to given value.

### HasDeploymentStatus

`func (o *EnvironmentOverviewResponseDeprecated) HasDeploymentStatus() bool`

HasDeploymentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


