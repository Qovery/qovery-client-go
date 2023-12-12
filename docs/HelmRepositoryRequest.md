# HelmRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Kind** | [**HelmRepositoryKindEnum**](HelmRepositoryKindEnum.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL of the helm chart repository: * For &#x60;OCI&#x60;: it must start by oci:// * For &#x60;HTTPS&#x60;: it must be start by https://  | [optional] 
**SkipTlsVerification** | **bool** | Bypass tls certificate verification when connecting to repository | 
**Config** | [**HelmRepositoryRequestConfig**](HelmRepositoryRequestConfig.md) |  | 

## Methods

### NewHelmRepositoryRequest

`func NewHelmRepositoryRequest(name string, kind HelmRepositoryKindEnum, skipTlsVerification bool, config HelmRepositoryRequestConfig, ) *HelmRepositoryRequest`

NewHelmRepositoryRequest instantiates a new HelmRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRepositoryRequestWithDefaults

`func NewHelmRepositoryRequestWithDefaults() *HelmRepositoryRequest`

NewHelmRepositoryRequestWithDefaults instantiates a new HelmRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HelmRepositoryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmRepositoryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmRepositoryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *HelmRepositoryRequest) GetKind() HelmRepositoryKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HelmRepositoryRequest) GetKindOk() (*HelmRepositoryKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HelmRepositoryRequest) SetKind(v HelmRepositoryKindEnum)`

SetKind sets Kind field to given value.


### GetDescription

`func (o *HelmRepositoryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HelmRepositoryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HelmRepositoryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HelmRepositoryRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *HelmRepositoryRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HelmRepositoryRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HelmRepositoryRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HelmRepositoryRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSkipTlsVerification

`func (o *HelmRepositoryRequest) GetSkipTlsVerification() bool`

GetSkipTlsVerification returns the SkipTlsVerification field if non-nil, zero value otherwise.

### GetSkipTlsVerificationOk

`func (o *HelmRepositoryRequest) GetSkipTlsVerificationOk() (*bool, bool)`

GetSkipTlsVerificationOk returns a tuple with the SkipTlsVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTlsVerification

`func (o *HelmRepositoryRequest) SetSkipTlsVerification(v bool)`

SetSkipTlsVerification sets SkipTlsVerification field to given value.


### GetConfig

`func (o *HelmRepositoryRequest) GetConfig() HelmRepositoryRequestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *HelmRepositoryRequest) GetConfigOk() (*HelmRepositoryRequestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *HelmRepositoryRequest) SetConfig(v HelmRepositoryRequestConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


