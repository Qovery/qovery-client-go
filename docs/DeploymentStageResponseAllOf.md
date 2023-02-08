# DeploymentStageResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | [**ReferenceObject**](ReferenceObject.md) |  | 
**Name** | Pointer to **string** | name is case insensitive | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeploymentOrder** | Pointer to **int32** | Position of the deployment stage within the environment | [optional] 
**Services** | Pointer to [**[]DeploymentStageServiceResponse**](DeploymentStageServiceResponse.md) |  | [optional] 

## Methods

### NewDeploymentStageResponseAllOf

`func NewDeploymentStageResponseAllOf(environment ReferenceObject, ) *DeploymentStageResponseAllOf`

NewDeploymentStageResponseAllOf instantiates a new DeploymentStageResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStageResponseAllOfWithDefaults

`func NewDeploymentStageResponseAllOfWithDefaults() *DeploymentStageResponseAllOf`

NewDeploymentStageResponseAllOfWithDefaults instantiates a new DeploymentStageResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *DeploymentStageResponseAllOf) GetEnvironment() ReferenceObject`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DeploymentStageResponseAllOf) GetEnvironmentOk() (*ReferenceObject, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DeploymentStageResponseAllOf) SetEnvironment(v ReferenceObject)`

SetEnvironment sets Environment field to given value.


### GetName

`func (o *DeploymentStageResponseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentStageResponseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentStageResponseAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentStageResponseAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DeploymentStageResponseAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentStageResponseAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentStageResponseAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeploymentStageResponseAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeploymentOrder

`func (o *DeploymentStageResponseAllOf) GetDeploymentOrder() int32`

GetDeploymentOrder returns the DeploymentOrder field if non-nil, zero value otherwise.

### GetDeploymentOrderOk

`func (o *DeploymentStageResponseAllOf) GetDeploymentOrderOk() (*int32, bool)`

GetDeploymentOrderOk returns a tuple with the DeploymentOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentOrder

`func (o *DeploymentStageResponseAllOf) SetDeploymentOrder(v int32)`

SetDeploymentOrder sets DeploymentOrder field to given value.

### HasDeploymentOrder

`func (o *DeploymentStageResponseAllOf) HasDeploymentOrder() bool`

HasDeploymentOrder returns a boolean if a field has been set.

### GetServices

`func (o *DeploymentStageResponseAllOf) GetServices() []DeploymentStageServiceResponse`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *DeploymentStageResponseAllOf) GetServicesOk() (*[]DeploymentStageServiceResponse, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *DeploymentStageResponseAllOf) SetServices(v []DeploymentStageServiceResponse)`

SetServices sets Services field to given value.

### HasServices

`func (o *DeploymentStageResponseAllOf) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


