# TerraformVariableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The variable key/name | 
**Value** | Pointer to **NullableString** | The variable value (null if secret and not exposed) | [optional] 
**Secret** | **bool** | Whether the variable is secret | 

## Methods

### NewTerraformVariableResponse

`func NewTerraformVariableResponse(key string, secret bool, ) *TerraformVariableResponse`

NewTerraformVariableResponse instantiates a new TerraformVariableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformVariableResponseWithDefaults

`func NewTerraformVariableResponseWithDefaults() *TerraformVariableResponse`

NewTerraformVariableResponseWithDefaults instantiates a new TerraformVariableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TerraformVariableResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TerraformVariableResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TerraformVariableResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *TerraformVariableResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TerraformVariableResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TerraformVariableResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TerraformVariableResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TerraformVariableResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TerraformVariableResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecret

`func (o *TerraformVariableResponse) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TerraformVariableResponse) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TerraformVariableResponse) SetSecret(v bool)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


