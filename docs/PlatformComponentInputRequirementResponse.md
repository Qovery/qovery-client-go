# PlatformComponentInputRequirementResponse

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
**Scope** | [**PlatformComponentConfigurationInputScope**](PlatformComponentConfigurationInputScope.md) |  | 
**Status** | [**PlatformComponentConfigurationRequirementStatus**](PlatformComponentConfigurationRequirementStatus.md) |  | 

## Methods

### NewPlatformComponentInputRequirementResponse

`func NewPlatformComponentInputRequirementResponse(key string, type_ string, required bool, label string, sensitive bool, constraints FieldSchemaConstraintsResponse, scope PlatformComponentConfigurationInputScope, status PlatformComponentConfigurationRequirementStatus, ) *PlatformComponentInputRequirementResponse`

NewPlatformComponentInputRequirementResponse instantiates a new PlatformComponentInputRequirementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformComponentInputRequirementResponseWithDefaults

`func NewPlatformComponentInputRequirementResponseWithDefaults() *PlatformComponentInputRequirementResponse`

NewPlatformComponentInputRequirementResponseWithDefaults instantiates a new PlatformComponentInputRequirementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PlatformComponentInputRequirementResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PlatformComponentInputRequirementResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PlatformComponentInputRequirementResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *PlatformComponentInputRequirementResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlatformComponentInputRequirementResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlatformComponentInputRequirementResponse) SetType(v string)`

SetType sets Type field to given value.


### GetRequired

`func (o *PlatformComponentInputRequirementResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PlatformComponentInputRequirementResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PlatformComponentInputRequirementResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDefaultValue

`func (o *PlatformComponentInputRequirementResponse) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *PlatformComponentInputRequirementResponse) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *PlatformComponentInputRequirementResponse) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *PlatformComponentInputRequirementResponse) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *PlatformComponentInputRequirementResponse) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *PlatformComponentInputRequirementResponse) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetLabel

`func (o *PlatformComponentInputRequirementResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PlatformComponentInputRequirementResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PlatformComponentInputRequirementResponse) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *PlatformComponentInputRequirementResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformComponentInputRequirementResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformComponentInputRequirementResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformComponentInputRequirementResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PlatformComponentInputRequirementResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PlatformComponentInputRequirementResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSensitive

`func (o *PlatformComponentInputRequirementResponse) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *PlatformComponentInputRequirementResponse) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *PlatformComponentInputRequirementResponse) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.


### GetConstraints

`func (o *PlatformComponentInputRequirementResponse) GetConstraints() FieldSchemaConstraintsResponse`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *PlatformComponentInputRequirementResponse) GetConstraintsOk() (*FieldSchemaConstraintsResponse, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *PlatformComponentInputRequirementResponse) SetConstraints(v FieldSchemaConstraintsResponse)`

SetConstraints sets Constraints field to given value.


### GetScope

`func (o *PlatformComponentInputRequirementResponse) GetScope() PlatformComponentConfigurationInputScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PlatformComponentInputRequirementResponse) GetScopeOk() (*PlatformComponentConfigurationInputScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PlatformComponentInputRequirementResponse) SetScope(v PlatformComponentConfigurationInputScope)`

SetScope sets Scope field to given value.


### GetStatus

`func (o *PlatformComponentInputRequirementResponse) GetStatus() PlatformComponentConfigurationRequirementStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformComponentInputRequirementResponse) GetStatusOk() (*PlatformComponentConfigurationRequirementStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformComponentInputRequirementResponse) SetStatus(v PlatformComponentConfigurationRequirementStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


