# SecretEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewSecretEditRequest

`func NewSecretEditRequest() *SecretEditRequest`

NewSecretEditRequest instantiates a new SecretEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretEditRequestWithDefaults

`func NewSecretEditRequestWithDefaults() *SecretEditRequest`

NewSecretEditRequestWithDefaults instantiates a new SecretEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SecretEditRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SecretEditRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SecretEditRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SecretEditRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetKey

`func (o *SecretEditRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SecretEditRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SecretEditRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SecretEditRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


