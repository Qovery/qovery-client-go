# ObjectFieldSchemaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Type** | **string** |  | 
**Required** | **bool** |  | 
**Label** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Sensitive** | **bool** |  | 
**Fields** | [**[]FieldSchemaResponse**](FieldSchemaResponse.md) | Nested field descriptors for this object&#39;s properties. | 

## Methods

### NewObjectFieldSchemaResponse

`func NewObjectFieldSchemaResponse(key string, type_ string, required bool, label string, sensitive bool, fields []FieldSchemaResponse, ) *ObjectFieldSchemaResponse`

NewObjectFieldSchemaResponse instantiates a new ObjectFieldSchemaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectFieldSchemaResponseWithDefaults

`func NewObjectFieldSchemaResponseWithDefaults() *ObjectFieldSchemaResponse`

NewObjectFieldSchemaResponseWithDefaults instantiates a new ObjectFieldSchemaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ObjectFieldSchemaResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ObjectFieldSchemaResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ObjectFieldSchemaResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *ObjectFieldSchemaResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObjectFieldSchemaResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObjectFieldSchemaResponse) SetType(v string)`

SetType sets Type field to given value.


### GetRequired

`func (o *ObjectFieldSchemaResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ObjectFieldSchemaResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ObjectFieldSchemaResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetLabel

`func (o *ObjectFieldSchemaResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ObjectFieldSchemaResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ObjectFieldSchemaResponse) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *ObjectFieldSchemaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectFieldSchemaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectFieldSchemaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectFieldSchemaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ObjectFieldSchemaResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ObjectFieldSchemaResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSensitive

`func (o *ObjectFieldSchemaResponse) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *ObjectFieldSchemaResponse) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *ObjectFieldSchemaResponse) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.


### GetFields

`func (o *ObjectFieldSchemaResponse) GetFields() []FieldSchemaResponse`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ObjectFieldSchemaResponse) GetFieldsOk() (*[]FieldSchemaResponse, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ObjectFieldSchemaResponse) SetFields(v []FieldSchemaResponse)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


