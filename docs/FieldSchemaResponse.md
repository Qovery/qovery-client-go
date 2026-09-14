# FieldSchemaResponse

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
**Constraints** | [**CollectionConstraintsResponse**](CollectionConstraintsResponse.md) |  | 
**Format** | Pointer to **string** | Optional editor format for a string field, independent of its scalar type. kubernetes-resource-yaml selects a single Kubernetes YAML object editor. Unknown formats should fall back to the ordinary string editor. | [optional] 
**Templates** | Pointer to [**[]FieldTemplateResponse**](FieldTemplateResponse.md) | Optional starting texts for an explicit user choice, with unique IDs within the field. Present only with format. Never apply as defaults or overwrite a saved value. The format selects the editor even when templates are absent. | [optional] 
**Fields** | [**[]FieldSchemaResponse**](FieldSchemaResponse.md) | Nested field descriptors for this object&#39;s properties. | 
**Items** | [**ArrayItemResponse**](ArrayItemResponse.md) |  | 
**ItemFields** | Pointer to [**[][]FieldSchemaResponse**]([]FieldSchemaResponse.md) | Evaluated field descriptors for each object item, in the same order as the configuration array. Use these row-specific descriptors when available; items.fields describes an object&#39;s fields before per-item evaluation. Omitted when unavailable, including scalar arrays. An evaluated empty object array has an empty itemFields array. | [optional] 

## Methods

### NewFieldSchemaResponse

`func NewFieldSchemaResponse(key string, type_ string, required bool, label string, sensitive bool, constraints CollectionConstraintsResponse, fields []FieldSchemaResponse, items ArrayItemResponse, ) *FieldSchemaResponse`

NewFieldSchemaResponse instantiates a new FieldSchemaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldSchemaResponseWithDefaults

`func NewFieldSchemaResponseWithDefaults() *FieldSchemaResponse`

NewFieldSchemaResponseWithDefaults instantiates a new FieldSchemaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FieldSchemaResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FieldSchemaResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FieldSchemaResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *FieldSchemaResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldSchemaResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldSchemaResponse) SetType(v string)`

SetType sets Type field to given value.


### GetRequired

`func (o *FieldSchemaResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldSchemaResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldSchemaResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDefaultValue

`func (o *FieldSchemaResponse) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *FieldSchemaResponse) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *FieldSchemaResponse) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *FieldSchemaResponse) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *FieldSchemaResponse) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *FieldSchemaResponse) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetLabel

`func (o *FieldSchemaResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FieldSchemaResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FieldSchemaResponse) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *FieldSchemaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldSchemaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldSchemaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FieldSchemaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FieldSchemaResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FieldSchemaResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSensitive

`func (o *FieldSchemaResponse) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *FieldSchemaResponse) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *FieldSchemaResponse) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.


### GetConstraints

`func (o *FieldSchemaResponse) GetConstraints() CollectionConstraintsResponse`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *FieldSchemaResponse) GetConstraintsOk() (*CollectionConstraintsResponse, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *FieldSchemaResponse) SetConstraints(v CollectionConstraintsResponse)`

SetConstraints sets Constraints field to given value.


### GetFormat

`func (o *FieldSchemaResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *FieldSchemaResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *FieldSchemaResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *FieldSchemaResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetTemplates

`func (o *FieldSchemaResponse) GetTemplates() []FieldTemplateResponse`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *FieldSchemaResponse) GetTemplatesOk() (*[]FieldTemplateResponse, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *FieldSchemaResponse) SetTemplates(v []FieldTemplateResponse)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *FieldSchemaResponse) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetFields

`func (o *FieldSchemaResponse) GetFields() []FieldSchemaResponse`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FieldSchemaResponse) GetFieldsOk() (*[]FieldSchemaResponse, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FieldSchemaResponse) SetFields(v []FieldSchemaResponse)`

SetFields sets Fields field to given value.


### GetItems

`func (o *FieldSchemaResponse) GetItems() ArrayItemResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FieldSchemaResponse) GetItemsOk() (*ArrayItemResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FieldSchemaResponse) SetItems(v ArrayItemResponse)`

SetItems sets Items field to given value.


### GetItemFields

`func (o *FieldSchemaResponse) GetItemFields() [][]FieldSchemaResponse`

GetItemFields returns the ItemFields field if non-nil, zero value otherwise.

### GetItemFieldsOk

`func (o *FieldSchemaResponse) GetItemFieldsOk() (*[][]FieldSchemaResponse, bool)`

GetItemFieldsOk returns a tuple with the ItemFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemFields

`func (o *FieldSchemaResponse) SetItemFields(v [][]FieldSchemaResponse)`

SetItemFields sets ItemFields field to given value.

### HasItemFields

`func (o *FieldSchemaResponse) HasItemFields() bool`

HasItemFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


