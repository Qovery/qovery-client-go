# VariableOverrideRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | the value to be used as Override of the targeted environment variable. | 
**OverrideScope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 
**OverrideParentId** | **string** | the id of the variable that is aliased. | 
**Description** | Pointer to **NullableString** | optional variable description (255 characters maximum) | [optional] 

## Methods

### NewVariableOverrideRequest

`func NewVariableOverrideRequest(value string, overrideScope APIVariableScopeEnum, overrideParentId string, ) *VariableOverrideRequest`

NewVariableOverrideRequest instantiates a new VariableOverrideRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableOverrideRequestWithDefaults

`func NewVariableOverrideRequestWithDefaults() *VariableOverrideRequest`

NewVariableOverrideRequestWithDefaults instantiates a new VariableOverrideRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *VariableOverrideRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableOverrideRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableOverrideRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetOverrideScope

`func (o *VariableOverrideRequest) GetOverrideScope() APIVariableScopeEnum`

GetOverrideScope returns the OverrideScope field if non-nil, zero value otherwise.

### GetOverrideScopeOk

`func (o *VariableOverrideRequest) GetOverrideScopeOk() (*APIVariableScopeEnum, bool)`

GetOverrideScopeOk returns a tuple with the OverrideScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideScope

`func (o *VariableOverrideRequest) SetOverrideScope(v APIVariableScopeEnum)`

SetOverrideScope sets OverrideScope field to given value.


### GetOverrideParentId

`func (o *VariableOverrideRequest) GetOverrideParentId() string`

GetOverrideParentId returns the OverrideParentId field if non-nil, zero value otherwise.

### GetOverrideParentIdOk

`func (o *VariableOverrideRequest) GetOverrideParentIdOk() (*string, bool)`

GetOverrideParentIdOk returns a tuple with the OverrideParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideParentId

`func (o *VariableOverrideRequest) SetOverrideParentId(v string)`

SetOverrideParentId sets OverrideParentId field to given value.


### GetDescription

`func (o *VariableOverrideRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VariableOverrideRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VariableOverrideRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VariableOverrideRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VariableOverrideRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VariableOverrideRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


