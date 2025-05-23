# TerraformDeploymentRestrictionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**DeploymentRestrictionModeEnum**](DeploymentRestrictionModeEnum.md) |  | 
**Type** | [**DeploymentRestrictionTypeEnum**](DeploymentRestrictionTypeEnum.md) |  | 
**Value** | Pointer to **string** | ‘For &#x60;PATH&#x60; restrictions, the value must not start with &#x60;/&#x60;’ | [optional] 

## Methods

### NewTerraformDeploymentRestrictionRequest

`func NewTerraformDeploymentRestrictionRequest(mode DeploymentRestrictionModeEnum, type_ DeploymentRestrictionTypeEnum, ) *TerraformDeploymentRestrictionRequest`

NewTerraformDeploymentRestrictionRequest instantiates a new TerraformDeploymentRestrictionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformDeploymentRestrictionRequestWithDefaults

`func NewTerraformDeploymentRestrictionRequestWithDefaults() *TerraformDeploymentRestrictionRequest`

NewTerraformDeploymentRestrictionRequestWithDefaults instantiates a new TerraformDeploymentRestrictionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *TerraformDeploymentRestrictionRequest) GetMode() DeploymentRestrictionModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TerraformDeploymentRestrictionRequest) GetModeOk() (*DeploymentRestrictionModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TerraformDeploymentRestrictionRequest) SetMode(v DeploymentRestrictionModeEnum)`

SetMode sets Mode field to given value.


### GetType

`func (o *TerraformDeploymentRestrictionRequest) GetType() DeploymentRestrictionTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TerraformDeploymentRestrictionRequest) GetTypeOk() (*DeploymentRestrictionTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TerraformDeploymentRestrictionRequest) SetType(v DeploymentRestrictionTypeEnum)`

SetType sets Type field to given value.


### GetValue

`func (o *TerraformDeploymentRestrictionRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TerraformDeploymentRestrictionRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TerraformDeploymentRestrictionRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TerraformDeploymentRestrictionRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


