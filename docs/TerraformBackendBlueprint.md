# TerraformBackendBlueprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Terraform backend type (e.g. s3, gcs, azurerm) | 
**Config** | Pointer to **map[string]string** | Static backend configuration (bucket, region, etc.). Credentials should be provided via environment variables, not here. | [optional] 

## Methods

### NewTerraformBackendBlueprint

`func NewTerraformBackendBlueprint(type_ string, ) *TerraformBackendBlueprint`

NewTerraformBackendBlueprint instantiates a new TerraformBackendBlueprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformBackendBlueprintWithDefaults

`func NewTerraformBackendBlueprintWithDefaults() *TerraformBackendBlueprint`

NewTerraformBackendBlueprintWithDefaults instantiates a new TerraformBackendBlueprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TerraformBackendBlueprint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TerraformBackendBlueprint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TerraformBackendBlueprint) SetType(v string)`

SetType sets Type field to given value.


### GetConfig

`func (o *TerraformBackendBlueprint) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TerraformBackendBlueprint) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TerraformBackendBlueprint) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *TerraformBackendBlueprint) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


