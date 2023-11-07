# HelmRepositoryRequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipTlsVerification** | Pointer to **bool** | Bypass tls certificate verification when connecting to repository | [optional] [default to false]
**Login** | Pointer to **string** | Required if the repository is private | [optional] 
**Password** | Pointer to **string** | Required if the repository is private | [optional] 

## Methods

### NewHelmRepositoryRequestConfig

`func NewHelmRepositoryRequestConfig() *HelmRepositoryRequestConfig`

NewHelmRepositoryRequestConfig instantiates a new HelmRepositoryRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRepositoryRequestConfigWithDefaults

`func NewHelmRepositoryRequestConfigWithDefaults() *HelmRepositoryRequestConfig`

NewHelmRepositoryRequestConfigWithDefaults instantiates a new HelmRepositoryRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipTlsVerification

`func (o *HelmRepositoryRequestConfig) GetSkipTlsVerification() bool`

GetSkipTlsVerification returns the SkipTlsVerification field if non-nil, zero value otherwise.

### GetSkipTlsVerificationOk

`func (o *HelmRepositoryRequestConfig) GetSkipTlsVerificationOk() (*bool, bool)`

GetSkipTlsVerificationOk returns a tuple with the SkipTlsVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTlsVerification

`func (o *HelmRepositoryRequestConfig) SetSkipTlsVerification(v bool)`

SetSkipTlsVerification sets SkipTlsVerification field to given value.

### HasSkipTlsVerification

`func (o *HelmRepositoryRequestConfig) HasSkipTlsVerification() bool`

HasSkipTlsVerification returns a boolean if a field has been set.

### GetLogin

`func (o *HelmRepositoryRequestConfig) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *HelmRepositoryRequestConfig) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *HelmRepositoryRequestConfig) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *HelmRepositoryRequestConfig) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetPassword

`func (o *HelmRepositoryRequestConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HelmRepositoryRequestConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HelmRepositoryRequestConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HelmRepositoryRequestConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


