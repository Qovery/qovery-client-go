# TerraformVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engine** | [**TerraformProviderEnum**](TerraformProviderEnum.md) |  | 
**Version** | **string** | Terraform version string | 

## Methods

### NewTerraformVersionResponse

`func NewTerraformVersionResponse(engine TerraformProviderEnum, version string, ) *TerraformVersionResponse`

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

`func (o *TerraformVersionResponse) GetEngine() TerraformProviderEnum`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *TerraformVersionResponse) GetEngineOk() (*TerraformProviderEnum, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *TerraformVersionResponse) SetEngine(v TerraformProviderEnum)`

SetEngine sets Engine field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


