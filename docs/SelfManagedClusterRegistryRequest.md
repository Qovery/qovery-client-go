# SelfManagedClusterRegistryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL of the registry: &#x60;http&#x60; or &#x60;https&#x60; scheme, host and optional port only, without user info, path, query or fragment, and not localhost. | 
**Credentials** | Pointer to [**SelfManagedClusterRegistryCredentialsRequest**](SelfManagedClusterRegistryCredentialsRequest.md) |  | [optional] 

## Methods

### NewSelfManagedClusterRegistryRequest

`func NewSelfManagedClusterRegistryRequest(url string, ) *SelfManagedClusterRegistryRequest`

NewSelfManagedClusterRegistryRequest instantiates a new SelfManagedClusterRegistryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfManagedClusterRegistryRequestWithDefaults

`func NewSelfManagedClusterRegistryRequestWithDefaults() *SelfManagedClusterRegistryRequest`

NewSelfManagedClusterRegistryRequestWithDefaults instantiates a new SelfManagedClusterRegistryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *SelfManagedClusterRegistryRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SelfManagedClusterRegistryRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SelfManagedClusterRegistryRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCredentials

`func (o *SelfManagedClusterRegistryRequest) GetCredentials() SelfManagedClusterRegistryCredentialsRequest`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SelfManagedClusterRegistryRequest) GetCredentialsOk() (*SelfManagedClusterRegistryCredentialsRequest, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SelfManagedClusterRegistryRequest) SetCredentials(v SelfManagedClusterRegistryCredentialsRequest)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SelfManagedClusterRegistryRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


