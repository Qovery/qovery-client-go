# VariableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**OverriddenVariable** | Pointer to [**VariableOverride**](VariableOverride.md) |  | [optional] 
**AliasedVariable** | Pointer to [**VariableAlias**](VariableAlias.md) |  | [optional] 
**Scope** | [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | 
**VariableType** | Pointer to [**APIVariableTypeEnum**](APIVariableTypeEnum.md) |  | [optional] 
**ServiceId** | Pointer to **string** | The id of the service referenced by this variable. | [optional] 
**ServiceName** | Pointer to **string** | The name of the service referenced by this variable. | [optional] 
**ServiceType** | Pointer to [**LinkedServiceTypeEnum**](LinkedServiceTypeEnum.md) |  | [optional] 
**OwnedBy** | Pointer to **string** | Entity that created/own the variable (i.e: Qovery, Doppler) | [optional] 

## Methods

### NewVariableResponse

`func NewVariableResponse(id string, createdAt time.Time, scope APIVariableScopeEnum, ) *VariableResponse`

NewVariableResponse instantiates a new VariableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableResponseWithDefaults

`func NewVariableResponseWithDefaults() *VariableResponse`

NewVariableResponseWithDefaults instantiates a new VariableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VariableResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariableResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariableResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *VariableResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VariableResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VariableResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VariableResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VariableResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VariableResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VariableResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOverriddenVariable

`func (o *VariableResponse) GetOverriddenVariable() VariableOverride`

GetOverriddenVariable returns the OverriddenVariable field if non-nil, zero value otherwise.

### GetOverriddenVariableOk

`func (o *VariableResponse) GetOverriddenVariableOk() (*VariableOverride, bool)`

GetOverriddenVariableOk returns a tuple with the OverriddenVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenVariable

`func (o *VariableResponse) SetOverriddenVariable(v VariableOverride)`

SetOverriddenVariable sets OverriddenVariable field to given value.

### HasOverriddenVariable

`func (o *VariableResponse) HasOverriddenVariable() bool`

HasOverriddenVariable returns a boolean if a field has been set.

### GetAliasedVariable

`func (o *VariableResponse) GetAliasedVariable() VariableAlias`

GetAliasedVariable returns the AliasedVariable field if non-nil, zero value otherwise.

### GetAliasedVariableOk

`func (o *VariableResponse) GetAliasedVariableOk() (*VariableAlias, bool)`

GetAliasedVariableOk returns a tuple with the AliasedVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasedVariable

`func (o *VariableResponse) SetAliasedVariable(v VariableAlias)`

SetAliasedVariable sets AliasedVariable field to given value.

### HasAliasedVariable

`func (o *VariableResponse) HasAliasedVariable() bool`

HasAliasedVariable returns a boolean if a field has been set.

### GetScope

`func (o *VariableResponse) GetScope() APIVariableScopeEnum`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VariableResponse) GetScopeOk() (*APIVariableScopeEnum, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VariableResponse) SetScope(v APIVariableScopeEnum)`

SetScope sets Scope field to given value.


### GetVariableType

`func (o *VariableResponse) GetVariableType() APIVariableTypeEnum`

GetVariableType returns the VariableType field if non-nil, zero value otherwise.

### GetVariableTypeOk

`func (o *VariableResponse) GetVariableTypeOk() (*APIVariableTypeEnum, bool)`

GetVariableTypeOk returns a tuple with the VariableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableType

`func (o *VariableResponse) SetVariableType(v APIVariableTypeEnum)`

SetVariableType sets VariableType field to given value.

### HasVariableType

`func (o *VariableResponse) HasVariableType() bool`

HasVariableType returns a boolean if a field has been set.

### GetServiceId

`func (o *VariableResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *VariableResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *VariableResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *VariableResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *VariableResponse) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *VariableResponse) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *VariableResponse) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *VariableResponse) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceType

`func (o *VariableResponse) GetServiceType() LinkedServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *VariableResponse) GetServiceTypeOk() (*LinkedServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *VariableResponse) SetServiceType(v LinkedServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *VariableResponse) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetOwnedBy

`func (o *VariableResponse) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *VariableResponse) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *VariableResponse) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *VariableResponse) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


