# ScalarFieldSchemaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Type** | **string** | Field type understood by the Console. | 
**Required** | **bool** |  | 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Label** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Sensitive** | **bool** |  | 
**Constraints** | [**FieldSchemaConstraintsResponse**](FieldSchemaConstraintsResponse.md) |  | 
**Format** | Pointer to **string** | Optional editor format for a string field, independent of its scalar type. kubernetes-resource-yaml selects a single Kubernetes YAML object editor. Unknown formats should fall back to the ordinary string editor. | [optional] 
**Templates** | Pointer to [**[]FieldTemplateResponse**](FieldTemplateResponse.md) | Optional starting texts for an explicit user choice, with unique IDs within the field. Present only with format. Never apply as defaults or overwrite a saved value. The format selects the editor even when templates are absent. | [optional] 

## Methods

### NewScalarFieldSchemaResponse

`func NewScalarFieldSchemaResponse(key string, type_ string, required bool, label string, sensitive bool, constraints FieldSchemaConstraintsResponse, ) *ScalarFieldSchemaResponse`

NewScalarFieldSchemaResponse instantiates a new ScalarFieldSchemaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScalarFieldSchemaResponseWithDefaults

`func NewScalarFieldSchemaResponseWithDefaults() *ScalarFieldSchemaResponse`

NewScalarFieldSchemaResponseWithDefaults instantiates a new ScalarFieldSchemaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ScalarFieldSchemaResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ScalarFieldSchemaResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ScalarFieldSchemaResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *ScalarFieldSchemaResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScalarFieldSchemaResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScalarFieldSchemaResponse) SetType(v string)`

SetType sets Type field to given value.


### GetRequired

`func (o *ScalarFieldSchemaResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ScalarFieldSchemaResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ScalarFieldSchemaResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDefaultValue

`func (o *ScalarFieldSchemaResponse) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ScalarFieldSchemaResponse) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ScalarFieldSchemaResponse) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ScalarFieldSchemaResponse) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ScalarFieldSchemaResponse) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ScalarFieldSchemaResponse) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetLabel

`func (o *ScalarFieldSchemaResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ScalarFieldSchemaResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ScalarFieldSchemaResponse) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *ScalarFieldSchemaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScalarFieldSchemaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScalarFieldSchemaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScalarFieldSchemaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ScalarFieldSchemaResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ScalarFieldSchemaResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSensitive

`func (o *ScalarFieldSchemaResponse) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *ScalarFieldSchemaResponse) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *ScalarFieldSchemaResponse) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.


### GetConstraints

`func (o *ScalarFieldSchemaResponse) GetConstraints() FieldSchemaConstraintsResponse`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *ScalarFieldSchemaResponse) GetConstraintsOk() (*FieldSchemaConstraintsResponse, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *ScalarFieldSchemaResponse) SetConstraints(v FieldSchemaConstraintsResponse)`

SetConstraints sets Constraints field to given value.


### GetFormat

`func (o *ScalarFieldSchemaResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ScalarFieldSchemaResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ScalarFieldSchemaResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ScalarFieldSchemaResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetTemplates

`func (o *ScalarFieldSchemaResponse) GetTemplates() []FieldTemplateResponse`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *ScalarFieldSchemaResponse) GetTemplatesOk() (*[]FieldTemplateResponse, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *ScalarFieldSchemaResponse) SetTemplates(v []FieldTemplateResponse)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *ScalarFieldSchemaResponse) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


