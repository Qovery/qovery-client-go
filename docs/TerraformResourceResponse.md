# TerraformResourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this resource record | 
**ResourceType** | **string** | Type of the Terraform resource (e.g., aws_instance, aws_s3_bucket) | 
**Name** | **string** | Name of the resource as defined in Terraform configuration | 
**Address** | **string** | Full address of the resource (e.g., aws_instance.web_server) | 
**Provider** | **string** | Terraform provider name (e.g., aws, google, azurerm) | 
**Mode** | **string** | Resource mode (managed or data source) | 
**Attributes** | **map[string]interface{}** | All resource attributes as key-value pairs | 
**ExtractedAt** | **time.Time** | Timestamp when the resource was extracted from Terraform state | 

## Methods

### NewTerraformResourceResponse

`func NewTerraformResourceResponse(id string, resourceType string, name string, address string, provider string, mode string, attributes map[string]interface{}, extractedAt time.Time, ) *TerraformResourceResponse`

NewTerraformResourceResponse instantiates a new TerraformResourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformResourceResponseWithDefaults

`func NewTerraformResourceResponseWithDefaults() *TerraformResourceResponse`

NewTerraformResourceResponseWithDefaults instantiates a new TerraformResourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TerraformResourceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TerraformResourceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TerraformResourceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetResourceType

`func (o *TerraformResourceResponse) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *TerraformResourceResponse) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *TerraformResourceResponse) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetName

`func (o *TerraformResourceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TerraformResourceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TerraformResourceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *TerraformResourceResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TerraformResourceResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TerraformResourceResponse) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetProvider

`func (o *TerraformResourceResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TerraformResourceResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TerraformResourceResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetMode

`func (o *TerraformResourceResponse) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TerraformResourceResponse) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TerraformResourceResponse) SetMode(v string)`

SetMode sets Mode field to given value.


### GetAttributes

`func (o *TerraformResourceResponse) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TerraformResourceResponse) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TerraformResourceResponse) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetExtractedAt

`func (o *TerraformResourceResponse) GetExtractedAt() time.Time`

GetExtractedAt returns the ExtractedAt field if non-nil, zero value otherwise.

### GetExtractedAtOk

`func (o *TerraformResourceResponse) GetExtractedAtOk() (*time.Time, bool)`

GetExtractedAtOk returns a tuple with the ExtractedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedAt

`func (o *TerraformResourceResponse) SetExtractedAt(v time.Time)`

SetExtractedAt sets ExtractedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


