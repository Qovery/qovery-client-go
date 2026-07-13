# PlatformComponentConfigurationFieldResponse

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
**Constraints** | [**PlatformComponentConfigurationConstraintsResponse**](PlatformComponentConfigurationConstraintsResponse.md) |  | 

## Methods

### NewPlatformComponentConfigurationFieldResponse

`func NewPlatformComponentConfigurationFieldResponse(key string, type_ string, required bool, label string, sensitive bool, constraints PlatformComponentConfigurationConstraintsResponse, ) *PlatformComponentConfigurationFieldResponse`

NewPlatformComponentConfigurationFieldResponse instantiates a new PlatformComponentConfigurationFieldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformComponentConfigurationFieldResponseWithDefaults

`func NewPlatformComponentConfigurationFieldResponseWithDefaults() *PlatformComponentConfigurationFieldResponse`

NewPlatformComponentConfigurationFieldResponseWithDefaults instantiates a new PlatformComponentConfigurationFieldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PlatformComponentConfigurationFieldResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PlatformComponentConfigurationFieldResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PlatformComponentConfigurationFieldResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *PlatformComponentConfigurationFieldResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlatformComponentConfigurationFieldResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlatformComponentConfigurationFieldResponse) SetType(v string)`

SetType sets Type field to given value.


### GetRequired

`func (o *PlatformComponentConfigurationFieldResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PlatformComponentConfigurationFieldResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PlatformComponentConfigurationFieldResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDefaultValue

`func (o *PlatformComponentConfigurationFieldResponse) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *PlatformComponentConfigurationFieldResponse) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *PlatformComponentConfigurationFieldResponse) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *PlatformComponentConfigurationFieldResponse) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *PlatformComponentConfigurationFieldResponse) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *PlatformComponentConfigurationFieldResponse) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetLabel

`func (o *PlatformComponentConfigurationFieldResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PlatformComponentConfigurationFieldResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PlatformComponentConfigurationFieldResponse) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *PlatformComponentConfigurationFieldResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformComponentConfigurationFieldResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformComponentConfigurationFieldResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformComponentConfigurationFieldResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PlatformComponentConfigurationFieldResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PlatformComponentConfigurationFieldResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSensitive

`func (o *PlatformComponentConfigurationFieldResponse) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *PlatformComponentConfigurationFieldResponse) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *PlatformComponentConfigurationFieldResponse) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.


### GetConstraints

`func (o *PlatformComponentConfigurationFieldResponse) GetConstraints() PlatformComponentConfigurationConstraintsResponse`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *PlatformComponentConfigurationFieldResponse) GetConstraintsOk() (*PlatformComponentConfigurationConstraintsResponse, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *PlatformComponentConfigurationFieldResponse) SetConstraints(v PlatformComponentConfigurationConstraintsResponse)`

SetConstraints sets Constraints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


