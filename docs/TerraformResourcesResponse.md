# TerraformResourcesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | [**[]TerraformResourceResponse**](TerraformResourceResponse.md) | Array of Terraform resources | 

## Methods

### NewTerraformResourcesResponse

`func NewTerraformResourcesResponse(resources []TerraformResourceResponse, ) *TerraformResourcesResponse`

NewTerraformResourcesResponse instantiates a new TerraformResourcesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformResourcesResponseWithDefaults

`func NewTerraformResourcesResponseWithDefaults() *TerraformResourcesResponse`

NewTerraformResourcesResponseWithDefaults instantiates a new TerraformResourcesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *TerraformResourcesResponse) GetResources() []TerraformResourceResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *TerraformResourcesResponse) GetResourcesOk() (*[]TerraformResourceResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *TerraformResourcesResponse) SetResources(v []TerraformResourceResponse)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


