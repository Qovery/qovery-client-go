# TerraformResourcesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TerraformId** | **string** | ID of the Terraform service | 
**ExecutionId** | **string** | Execution ID in format \&quot;UUID-version\&quot; or \&quot;UUID-version-timestamp\&quot; Example: 550e8400-e29b-41d4-a716-446655440000-1  | 
**ResourcesJson** | **string** | JSON array of terraform resources extracted from terraform show output. Each resource contains: resource_type, name, provider, address, mode, attributes  | 

## Methods

### NewTerraformResourcesRequest

`func NewTerraformResourcesRequest(terraformId string, executionId string, resourcesJson string, ) *TerraformResourcesRequest`

NewTerraformResourcesRequest instantiates a new TerraformResourcesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformResourcesRequestWithDefaults

`func NewTerraformResourcesRequestWithDefaults() *TerraformResourcesRequest`

NewTerraformResourcesRequestWithDefaults instantiates a new TerraformResourcesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerraformId

`func (o *TerraformResourcesRequest) GetTerraformId() string`

GetTerraformId returns the TerraformId field if non-nil, zero value otherwise.

### GetTerraformIdOk

`func (o *TerraformResourcesRequest) GetTerraformIdOk() (*string, bool)`

GetTerraformIdOk returns a tuple with the TerraformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformId

`func (o *TerraformResourcesRequest) SetTerraformId(v string)`

SetTerraformId sets TerraformId field to given value.


### GetExecutionId

`func (o *TerraformResourcesRequest) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *TerraformResourcesRequest) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *TerraformResourcesRequest) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetResourcesJson

`func (o *TerraformResourcesRequest) GetResourcesJson() string`

GetResourcesJson returns the ResourcesJson field if non-nil, zero value otherwise.

### GetResourcesJsonOk

`func (o *TerraformResourcesRequest) GetResourcesJsonOk() (*string, bool)`

GetResourcesJsonOk returns a tuple with the ResourcesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesJson

`func (o *TerraformResourcesRequest) SetResourcesJson(v string)`

SetResourcesJson sets ResourcesJson field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


