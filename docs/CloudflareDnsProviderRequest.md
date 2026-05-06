# CloudflareDnsProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Domain** | **string** |  | 
**Email** | **string** |  | 
**ApiToken** | **string** | Cloudflare API token. Must be provided when editing Cloudflare DNS provider. | 
**Proxied** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewCloudflareDnsProviderRequest

`func NewCloudflareDnsProviderRequest(provider string, domain string, email string, apiToken string, ) *CloudflareDnsProviderRequest`

NewCloudflareDnsProviderRequest instantiates a new CloudflareDnsProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudflareDnsProviderRequestWithDefaults

`func NewCloudflareDnsProviderRequestWithDefaults() *CloudflareDnsProviderRequest`

NewCloudflareDnsProviderRequestWithDefaults instantiates a new CloudflareDnsProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *CloudflareDnsProviderRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudflareDnsProviderRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudflareDnsProviderRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetDomain

`func (o *CloudflareDnsProviderRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CloudflareDnsProviderRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CloudflareDnsProviderRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEmail

`func (o *CloudflareDnsProviderRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CloudflareDnsProviderRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CloudflareDnsProviderRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetApiToken

`func (o *CloudflareDnsProviderRequest) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *CloudflareDnsProviderRequest) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *CloudflareDnsProviderRequest) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetProxied

`func (o *CloudflareDnsProviderRequest) GetProxied() bool`

GetProxied returns the Proxied field if non-nil, zero value otherwise.

### GetProxiedOk

`func (o *CloudflareDnsProviderRequest) GetProxiedOk() (*bool, bool)`

GetProxiedOk returns a tuple with the Proxied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxied

`func (o *CloudflareDnsProviderRequest) SetProxied(v bool)`

SetProxied sets Proxied field to given value.

### HasProxied

`func (o *CloudflareDnsProviderRequest) HasProxied() bool`

HasProxied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


