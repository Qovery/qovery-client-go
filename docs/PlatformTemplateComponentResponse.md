# PlatformTemplateComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Kind** | [**PlatformTemplateComponentKind**](PlatformTemplateComponentKind.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Fields** | [**[]FieldSchemaResponse**](FieldSchemaResponse.md) |  | 

## Methods

### NewPlatformTemplateComponentResponse

`func NewPlatformTemplateComponentResponse(key string, kind PlatformTemplateComponentKind, fields []FieldSchemaResponse, ) *PlatformTemplateComponentResponse`

NewPlatformTemplateComponentResponse instantiates a new PlatformTemplateComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformTemplateComponentResponseWithDefaults

`func NewPlatformTemplateComponentResponseWithDefaults() *PlatformTemplateComponentResponse`

NewPlatformTemplateComponentResponseWithDefaults instantiates a new PlatformTemplateComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PlatformTemplateComponentResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PlatformTemplateComponentResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PlatformTemplateComponentResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetKind

`func (o *PlatformTemplateComponentResponse) GetKind() PlatformTemplateComponentKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PlatformTemplateComponentResponse) GetKindOk() (*PlatformTemplateComponentKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PlatformTemplateComponentResponse) SetKind(v PlatformTemplateComponentKind)`

SetKind sets Kind field to given value.


### GetDescription

`func (o *PlatformTemplateComponentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformTemplateComponentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformTemplateComponentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformTemplateComponentResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PlatformTemplateComponentResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PlatformTemplateComponentResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFields

`func (o *PlatformTemplateComponentResponse) GetFields() []FieldSchemaResponse`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PlatformTemplateComponentResponse) GetFieldsOk() (*[]FieldSchemaResponse, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PlatformTemplateComponentResponse) SetFields(v []FieldSchemaResponse)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


