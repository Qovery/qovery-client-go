# TerraformVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engine** | Pointer to **string** | Terraform engine | [optional] 
**Version** | Pointer to **string** | Terraform version string | [optional] 

## Methods

### NewTerraformVersionResponse

`func NewTerraformVersionResponse() *TerraformVersionResponse`

NewTerraformVersionResponse instantiates a new TerraformVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformVersionResponseWithDefaults

`func NewTerraformVersionResponseWithDefaults() *TerraformVersionResponse`

NewTerraformVersionResponseWithDefaults instantiates a new TerraformVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngine

`func (o *TerraformVersionResponse) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *TerraformVersionResponse) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *TerraformVersionResponse) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *TerraformVersionResponse) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetVersion

`func (o *TerraformVersionResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TerraformVersionResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TerraformVersionResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TerraformVersionResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


