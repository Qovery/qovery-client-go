# DeploymentStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the deployment stage | 
**Description** | Pointer to **NullableString** | free test describing this stage | [optional] 

## Methods

### NewDeploymentStageRequest

`func NewDeploymentStageRequest(name string, ) *DeploymentStageRequest`

NewDeploymentStageRequest instantiates a new DeploymentStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStageRequestWithDefaults

`func NewDeploymentStageRequestWithDefaults() *DeploymentStageRequest`

NewDeploymentStageRequestWithDefaults instantiates a new DeploymentStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DeploymentStageRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentStageRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentStageRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeploymentStageRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DeploymentStageRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DeploymentStageRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


