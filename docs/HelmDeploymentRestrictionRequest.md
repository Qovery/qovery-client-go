# HelmDeploymentRestrictionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**DeploymentRestrictionModeEnum**](DeploymentRestrictionModeEnum.md) |  | 
**Type** | [**DeploymentRestrictionTypeEnum**](DeploymentRestrictionTypeEnum.md) |  | 
**Value** | **string** | For &#x60;PATH&#x60; restrictions, the value must not start with &#x60;/&#x60; | 

## Methods

### NewHelmDeploymentRestrictionRequest

`func NewHelmDeploymentRestrictionRequest(mode DeploymentRestrictionModeEnum, type_ DeploymentRestrictionTypeEnum, value string, ) *HelmDeploymentRestrictionRequest`

NewHelmDeploymentRestrictionRequest instantiates a new HelmDeploymentRestrictionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmDeploymentRestrictionRequestWithDefaults

`func NewHelmDeploymentRestrictionRequestWithDefaults() *HelmDeploymentRestrictionRequest`

NewHelmDeploymentRestrictionRequestWithDefaults instantiates a new HelmDeploymentRestrictionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *HelmDeploymentRestrictionRequest) GetMode() DeploymentRestrictionModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HelmDeploymentRestrictionRequest) GetModeOk() (*DeploymentRestrictionModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HelmDeploymentRestrictionRequest) SetMode(v DeploymentRestrictionModeEnum)`

SetMode sets Mode field to given value.


### GetType

`func (o *HelmDeploymentRestrictionRequest) GetType() DeploymentRestrictionTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HelmDeploymentRestrictionRequest) GetTypeOk() (*DeploymentRestrictionTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HelmDeploymentRestrictionRequest) SetType(v DeploymentRestrictionTypeEnum)`

SetType sets Type field to given value.


### GetValue

`func (o *HelmDeploymentRestrictionRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HelmDeploymentRestrictionRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HelmDeploymentRestrictionRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


