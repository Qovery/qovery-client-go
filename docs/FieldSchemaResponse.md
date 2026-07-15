# FieldSchemaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Type** | **string** | Field type understood by the Console, such as string, number, or bool | 
**Required** | **bool** |  | 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Label** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Sensitive** | **bool** |  | 
**Constraints** | [**FieldSchemaConstraintsResponse**](FieldSchemaConstraintsResponse.md) |  | 

## Methods

### NewFieldSchemaResponse

`func NewFieldSchemaResponse(key string, type_ string, required bool, label string, sensitive bool, constraints FieldSchemaConstraintsResponse, ) *FieldSchemaResponse`

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

`func (o *FieldSchemaResponse) GetConstraints() FieldSchemaConstraintsResponse`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *FieldSchemaResponse) GetConstraintsOk() (*FieldSchemaConstraintsResponse, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *FieldSchemaResponse) SetConstraints(v FieldSchemaConstraintsResponse)`

SetConstraints sets Constraints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


