# TerraformDeploymentRestrictionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Mode** | [**DeploymentRestrictionModeEnum**](DeploymentRestrictionModeEnum.md) |  | 
**Type** | [**DeploymentRestrictionTypeEnum**](DeploymentRestrictionTypeEnum.md) |  | 
**Value** | **string** | ‘For &#x60;PATH&#x60; restrictions, the value must not start with &#x60;/&#x60;’ | 

## Methods

### NewTerraformDeploymentRestrictionResponse

`func NewTerraformDeploymentRestrictionResponse(id string, createdAt time.Time, mode DeploymentRestrictionModeEnum, type_ DeploymentRestrictionTypeEnum, value string, ) *TerraformDeploymentRestrictionResponse`

NewTerraformDeploymentRestrictionResponse instantiates a new TerraformDeploymentRestrictionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformDeploymentRestrictionResponseWithDefaults

`func NewTerraformDeploymentRestrictionResponseWithDefaults() *TerraformDeploymentRestrictionResponse`

NewTerraformDeploymentRestrictionResponseWithDefaults instantiates a new TerraformDeploymentRestrictionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TerraformDeploymentRestrictionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TerraformDeploymentRestrictionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TerraformDeploymentRestrictionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *TerraformDeploymentRestrictionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TerraformDeploymentRestrictionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TerraformDeploymentRestrictionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TerraformDeploymentRestrictionResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TerraformDeploymentRestrictionResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TerraformDeploymentRestrictionResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TerraformDeploymentRestrictionResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetMode

`func (o *TerraformDeploymentRestrictionResponse) GetMode() DeploymentRestrictionModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TerraformDeploymentRestrictionResponse) GetModeOk() (*DeploymentRestrictionModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TerraformDeploymentRestrictionResponse) SetMode(v DeploymentRestrictionModeEnum)`

SetMode sets Mode field to given value.


### GetType

`func (o *TerraformDeploymentRestrictionResponse) GetType() DeploymentRestrictionTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TerraformDeploymentRestrictionResponse) GetTypeOk() (*DeploymentRestrictionTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TerraformDeploymentRestrictionResponse) SetType(v DeploymentRestrictionTypeEnum)`

SetType sets Type field to given value.


### GetValue

`func (o *TerraformDeploymentRestrictionResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TerraformDeploymentRestrictionResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TerraformDeploymentRestrictionResponse) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


