# EnvironmentVariableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Key** | **string** | key is case sensitive | 
**Value** | **string** | value of the env variable. | 
**OverriddenVariable** | Pointer to [**EnvironmentVariableResponseAllOfOverriddenVariable**](EnvironmentVariableResponseAllOfOverriddenVariable.md) |  | [optional] 
**AliasedVariable** | Pointer to [**EnvironmentVariableResponseAllOfAliasedVariable**](EnvironmentVariableResponseAllOfAliasedVariable.md) |  | [optional] 
**Scope** | [**EnvironmentVariableScopeEnum**](EnvironmentVariableScopeEnum.md) |  | 
**ServiceName** | Pointer to **string** |  | [optional] 

## Methods

### NewEnvironmentVariableResponse

`func NewEnvironmentVariableResponse(id string, createdAt time.Time, key string, value string, scope EnvironmentVariableScopeEnum, ) *EnvironmentVariableResponse`

NewEnvironmentVariableResponse instantiates a new EnvironmentVariableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariableResponseWithDefaults

`func NewEnvironmentVariableResponseWithDefaults() *EnvironmentVariableResponse`

NewEnvironmentVariableResponseWithDefaults instantiates a new EnvironmentVariableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentVariableResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentVariableResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentVariableResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EnvironmentVariableResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentVariableResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentVariableResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EnvironmentVariableResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentVariableResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentVariableResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentVariableResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetKey

`func (o *EnvironmentVariableResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnvironmentVariableResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnvironmentVariableResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *EnvironmentVariableResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvironmentVariableResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvironmentVariableResponse) SetValue(v string)`

SetValue sets Value field to given value.


### GetOverriddenVariable

`func (o *EnvironmentVariableResponse) GetOverriddenVariable() EnvironmentVariableResponseAllOfOverriddenVariable`

GetOverriddenVariable returns the OverriddenVariable field if non-nil, zero value otherwise.

### GetOverriddenVariableOk

`func (o *EnvironmentVariableResponse) GetOverriddenVariableOk() (*EnvironmentVariableResponseAllOfOverriddenVariable, bool)`

GetOverriddenVariableOk returns a tuple with the OverriddenVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenVariable

`func (o *EnvironmentVariableResponse) SetOverriddenVariable(v EnvironmentVariableResponseAllOfOverriddenVariable)`

SetOverriddenVariable sets OverriddenVariable field to given value.

### HasOverriddenVariable

`func (o *EnvironmentVariableResponse) HasOverriddenVariable() bool`

HasOverriddenVariable returns a boolean if a field has been set.

### GetAliasedVariable

`func (o *EnvironmentVariableResponse) GetAliasedVariable() EnvironmentVariableResponseAllOfAliasedVariable`

GetAliasedVariable returns the AliasedVariable field if non-nil, zero value otherwise.

### GetAliasedVariableOk

`func (o *EnvironmentVariableResponse) GetAliasedVariableOk() (*EnvironmentVariableResponseAllOfAliasedVariable, bool)`

GetAliasedVariableOk returns a tuple with the AliasedVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasedVariable

`func (o *EnvironmentVariableResponse) SetAliasedVariable(v EnvironmentVariableResponseAllOfAliasedVariable)`

SetAliasedVariable sets AliasedVariable field to given value.

### HasAliasedVariable

`func (o *EnvironmentVariableResponse) HasAliasedVariable() bool`

HasAliasedVariable returns a boolean if a field has been set.

### GetScope

`func (o *EnvironmentVariableResponse) GetScope() EnvironmentVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EnvironmentVariableResponse) GetScopeOk() (*EnvironmentVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EnvironmentVariableResponse) SetScope(v EnvironmentVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetServiceName

`func (o *EnvironmentVariableResponse) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *EnvironmentVariableResponse) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *EnvironmentVariableResponse) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *EnvironmentVariableResponse) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


