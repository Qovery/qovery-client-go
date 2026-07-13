# PlatformTemplateLayerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Mandatory** | **bool** |  | 
**EnabledByDefault** | **bool** |  | 
**Modes** | [**[]PlatformClusterMode**](PlatformClusterMode.md) |  | 
**Providers** | Pointer to [**[]PlatformCloudVendor**](PlatformCloudVendor.md) |  | [optional] 
**ComponentKeys** | **[]string** |  | 
**Components** | [**[]PlatformTemplateComponentResponse**](PlatformTemplateComponentResponse.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPlatformTemplateLayerResponse

`func NewPlatformTemplateLayerResponse(key string, mandatory bool, enabledByDefault bool, modes []PlatformClusterMode, componentKeys []string, components []PlatformTemplateComponentResponse, ) *PlatformTemplateLayerResponse`

NewPlatformTemplateLayerResponse instantiates a new PlatformTemplateLayerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformTemplateLayerResponseWithDefaults

`func NewPlatformTemplateLayerResponseWithDefaults() *PlatformTemplateLayerResponse`

NewPlatformTemplateLayerResponseWithDefaults instantiates a new PlatformTemplateLayerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PlatformTemplateLayerResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PlatformTemplateLayerResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PlatformTemplateLayerResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetMandatory

`func (o *PlatformTemplateLayerResponse) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *PlatformTemplateLayerResponse) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *PlatformTemplateLayerResponse) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.


### GetEnabledByDefault

`func (o *PlatformTemplateLayerResponse) GetEnabledByDefault() bool`

GetEnabledByDefault returns the EnabledByDefault field if non-nil, zero value otherwise.

### GetEnabledByDefaultOk

`func (o *PlatformTemplateLayerResponse) GetEnabledByDefaultOk() (*bool, bool)`

GetEnabledByDefaultOk returns a tuple with the EnabledByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledByDefault

`func (o *PlatformTemplateLayerResponse) SetEnabledByDefault(v bool)`

SetEnabledByDefault sets EnabledByDefault field to given value.


### GetModes

`func (o *PlatformTemplateLayerResponse) GetModes() []PlatformClusterMode`

GetModes returns the Modes field if non-nil, zero value otherwise.

### GetModesOk

`func (o *PlatformTemplateLayerResponse) GetModesOk() (*[]PlatformClusterMode, bool)`

GetModesOk returns a tuple with the Modes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModes

`func (o *PlatformTemplateLayerResponse) SetModes(v []PlatformClusterMode)`

SetModes sets Modes field to given value.


### GetProviders

`func (o *PlatformTemplateLayerResponse) GetProviders() []PlatformCloudVendor`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *PlatformTemplateLayerResponse) GetProvidersOk() (*[]PlatformCloudVendor, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *PlatformTemplateLayerResponse) SetProviders(v []PlatformCloudVendor)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *PlatformTemplateLayerResponse) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### SetProvidersNil

`func (o *PlatformTemplateLayerResponse) SetProvidersNil(b bool)`

 SetProvidersNil sets the value for Providers to be an explicit nil

### UnsetProviders
`func (o *PlatformTemplateLayerResponse) UnsetProviders()`

UnsetProviders ensures that no value is present for Providers, not even an explicit nil
### GetComponentKeys

`func (o *PlatformTemplateLayerResponse) GetComponentKeys() []string`

GetComponentKeys returns the ComponentKeys field if non-nil, zero value otherwise.

### GetComponentKeysOk

`func (o *PlatformTemplateLayerResponse) GetComponentKeysOk() (*[]string, bool)`

GetComponentKeysOk returns a tuple with the ComponentKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentKeys

`func (o *PlatformTemplateLayerResponse) SetComponentKeys(v []string)`

SetComponentKeys sets ComponentKeys field to given value.


### GetComponents

`func (o *PlatformTemplateLayerResponse) GetComponents() []PlatformTemplateComponentResponse`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *PlatformTemplateLayerResponse) GetComponentsOk() (*[]PlatformTemplateComponentResponse, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *PlatformTemplateLayerResponse) SetComponents(v []PlatformTemplateComponentResponse)`

SetComponents sets Components field to given value.


### GetDescription

`func (o *PlatformTemplateLayerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformTemplateLayerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformTemplateLayerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformTemplateLayerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PlatformTemplateLayerResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PlatformTemplateLayerResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


