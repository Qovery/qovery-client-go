# VariableOverrideRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**AliasScope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 
**AliasParentId** | **string** |  | 

## Methods

### NewVariableOverrideRequest

`func NewVariableOverrideRequest(value string, aliasScope APIVariableScopeEnum, aliasParentId string, ) *VariableOverrideRequest`

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


### GetAliasScope

`func (o *VariableOverrideRequest) GetAliasScope() APIVariableScopeEnum`

GetAliasScope returns the AliasScope field if non-nil, zero value otherwise.

### GetAliasScopeOk

`func (o *VariableOverrideRequest) GetAliasScopeOk() (*APIVariableScopeEnum, bool)`

GetAliasScopeOk returns a tuple with the AliasScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasScope

`func (o *VariableOverrideRequest) SetAliasScope(v APIVariableScopeEnum)`

SetAliasScope sets AliasScope field to given value.


### GetAliasParentId

`func (o *VariableOverrideRequest) GetAliasParentId() string`

GetAliasParentId returns the AliasParentId field if non-nil, zero value otherwise.

### GetAliasParentIdOk

`func (o *VariableOverrideRequest) GetAliasParentIdOk() (*string, bool)`

GetAliasParentIdOk returns a tuple with the AliasParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasParentId

`func (o *VariableOverrideRequest) SetAliasParentId(v string)`

SetAliasParentId sets AliasParentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


