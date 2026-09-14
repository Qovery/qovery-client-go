# ArrayFieldSchemaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Type** | **string** |  | 
**Required** | **bool** |  | 
**Label** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Sensitive** | **bool** |  | 
**Constraints** | [**CollectionConstraintsResponse**](CollectionConstraintsResponse.md) |  | 
**Items** | [**ArrayItemResponse**](ArrayItemResponse.md) |  | 
**ItemFields** | Pointer to [**[][]FieldSchemaResponse**]([]FieldSchemaResponse.md) | Evaluated field descriptors for each object item, in the same order as the configuration array. Use these row-specific descriptors when available; items.fields describes an object&#39;s fields before per-item evaluation. Omitted when unavailable, including scalar arrays. An evaluated empty object array has an empty itemFields array. | [optional] 

## Methods

### NewArrayFieldSchemaResponse

`func NewArrayFieldSchemaResponse(key string, type_ string, required bool, label string, sensitive bool, constraints CollectionConstraintsResponse, items ArrayItemResponse, ) *ArrayFieldSchemaResponse`

NewArrayFieldSchemaResponse instantiates a new ArrayFieldSchemaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArrayFieldSchemaResponseWithDefaults

`func NewArrayFieldSchemaResponseWithDefaults() *ArrayFieldSchemaResponse`

NewArrayFieldSchemaResponseWithDefaults instantiates a new ArrayFieldSchemaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ArrayFieldSchemaResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ArrayFieldSchemaResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ArrayFieldSchemaResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *ArrayFieldSchemaResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArrayFieldSchemaResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArrayFieldSchemaResponse) SetType(v string)`

SetType sets Type field to given value.


### GetRequired

`func (o *ArrayFieldSchemaResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ArrayFieldSchemaResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ArrayFieldSchemaResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetLabel

`func (o *ArrayFieldSchemaResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ArrayFieldSchemaResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ArrayFieldSchemaResponse) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *ArrayFieldSchemaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArrayFieldSchemaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArrayFieldSchemaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArrayFieldSchemaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ArrayFieldSchemaResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ArrayFieldSchemaResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSensitive

`func (o *ArrayFieldSchemaResponse) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *ArrayFieldSchemaResponse) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *ArrayFieldSchemaResponse) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.


### GetConstraints

`func (o *ArrayFieldSchemaResponse) GetConstraints() CollectionConstraintsResponse`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *ArrayFieldSchemaResponse) GetConstraintsOk() (*CollectionConstraintsResponse, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *ArrayFieldSchemaResponse) SetConstraints(v CollectionConstraintsResponse)`

SetConstraints sets Constraints field to given value.


### GetItems

`func (o *ArrayFieldSchemaResponse) GetItems() ArrayItemResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ArrayFieldSchemaResponse) GetItemsOk() (*ArrayItemResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ArrayFieldSchemaResponse) SetItems(v ArrayItemResponse)`

SetItems sets Items field to given value.


### GetItemFields

`func (o *ArrayFieldSchemaResponse) GetItemFields() [][]FieldSchemaResponse`

GetItemFields returns the ItemFields field if non-nil, zero value otherwise.

### GetItemFieldsOk

`func (o *ArrayFieldSchemaResponse) GetItemFieldsOk() (*[][]FieldSchemaResponse, bool)`

GetItemFieldsOk returns a tuple with the ItemFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemFields

`func (o *ArrayFieldSchemaResponse) SetItemFields(v [][]FieldSchemaResponse)`

SetItemFields sets ItemFields field to given value.

### HasItemFields

`func (o *ArrayFieldSchemaResponse) HasItemFields() bool`

HasItemFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


