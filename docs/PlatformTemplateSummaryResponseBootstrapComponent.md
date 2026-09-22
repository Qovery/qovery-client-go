# PlatformTemplateSummaryResponseBootstrapComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Kind** | [**PlatformTemplateComponentKind**](PlatformTemplateComponentKind.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Fields** | [**[]FieldSchemaResponse**](FieldSchemaResponse.md) |  | 
**ConfigurationSections** | Pointer to [**[]PlatformConfigurationSectionResponse**](PlatformConfigurationSectionResponse.md) | Additional configuration sections displayed under this component. Its own configuration and cluster inputs remain available. Omitted or empty keeps the existing editor behavior; fields always remain declared on their owner. | [optional] 

## Methods

### NewPlatformTemplateSummaryResponseBootstrapComponent

`func NewPlatformTemplateSummaryResponseBootstrapComponent(key string, kind PlatformTemplateComponentKind, fields []FieldSchemaResponse, ) *PlatformTemplateSummaryResponseBootstrapComponent`

NewPlatformTemplateSummaryResponseBootstrapComponent instantiates a new PlatformTemplateSummaryResponseBootstrapComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformTemplateSummaryResponseBootstrapComponentWithDefaults

`func NewPlatformTemplateSummaryResponseBootstrapComponentWithDefaults() *PlatformTemplateSummaryResponseBootstrapComponent`

NewPlatformTemplateSummaryResponseBootstrapComponentWithDefaults instantiates a new PlatformTemplateSummaryResponseBootstrapComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) SetKey(v string)`

SetKey sets Key field to given value.


### GetKind

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) GetKind() PlatformTemplateComponentKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) GetKindOk() (*PlatformTemplateComponentKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) SetKind(v PlatformTemplateComponentKind)`

SetKind sets Kind field to given value.


### GetDescription

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PlatformTemplateSummaryResponseBootstrapComponent) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFields

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) GetFields() []FieldSchemaResponse`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) GetFieldsOk() (*[]FieldSchemaResponse, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) SetFields(v []FieldSchemaResponse)`

SetFields sets Fields field to given value.


### GetConfigurationSections

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) GetConfigurationSections() []PlatformConfigurationSectionResponse`

GetConfigurationSections returns the ConfigurationSections field if non-nil, zero value otherwise.

### GetConfigurationSectionsOk

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) GetConfigurationSectionsOk() (*[]PlatformConfigurationSectionResponse, bool)`

GetConfigurationSectionsOk returns a tuple with the ConfigurationSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationSections

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) SetConfigurationSections(v []PlatformConfigurationSectionResponse)`

SetConfigurationSections sets ConfigurationSections field to given value.

### HasConfigurationSections

`func (o *PlatformTemplateSummaryResponseBootstrapComponent) HasConfigurationSections() bool`

HasConfigurationSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


