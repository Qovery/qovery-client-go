# ClusterDnsProviderRequestProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Domain** | **string** |  | 
**Email** | **string** |  | 
**ApiToken** | **string** | Cloudflare API token. Must be provided when editing Cloudflare DNS provider. | 
**Proxied** | Pointer to **bool** |  | [optional] [default to false]
**Credentials** | [**Route53DnsProviderRequestCredentials**](Route53DnsProviderRequestCredentials.md) |  | 
**AwsRegion** | **string** |  | 
**HostedZoneId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewClusterDnsProviderRequestProvider

`func NewClusterDnsProviderRequestProvider(provider string, domain string, email string, apiToken string, credentials Route53DnsProviderRequestCredentials, awsRegion string, ) *ClusterDnsProviderRequestProvider`

NewClusterDnsProviderRequestProvider instantiates a new ClusterDnsProviderRequestProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDnsProviderRequestProviderWithDefaults

`func NewClusterDnsProviderRequestProviderWithDefaults() *ClusterDnsProviderRequestProvider`

NewClusterDnsProviderRequestProviderWithDefaults instantiates a new ClusterDnsProviderRequestProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ClusterDnsProviderRequestProvider) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ClusterDnsProviderRequestProvider) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ClusterDnsProviderRequestProvider) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetDomain

`func (o *ClusterDnsProviderRequestProvider) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ClusterDnsProviderRequestProvider) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ClusterDnsProviderRequestProvider) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEmail

`func (o *ClusterDnsProviderRequestProvider) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ClusterDnsProviderRequestProvider) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ClusterDnsProviderRequestProvider) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetApiToken

`func (o *ClusterDnsProviderRequestProvider) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *ClusterDnsProviderRequestProvider) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *ClusterDnsProviderRequestProvider) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetProxied

`func (o *ClusterDnsProviderRequestProvider) GetProxied() bool`

GetProxied returns the Proxied field if non-nil, zero value otherwise.

### GetProxiedOk

`func (o *ClusterDnsProviderRequestProvider) GetProxiedOk() (*bool, bool)`

GetProxiedOk returns a tuple with the Proxied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxied

`func (o *ClusterDnsProviderRequestProvider) SetProxied(v bool)`

SetProxied sets Proxied field to given value.

### HasProxied

`func (o *ClusterDnsProviderRequestProvider) HasProxied() bool`

HasProxied returns a boolean if a field has been set.

### GetCredentials

`func (o *ClusterDnsProviderRequestProvider) GetCredentials() Route53DnsProviderRequestCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ClusterDnsProviderRequestProvider) GetCredentialsOk() (*Route53DnsProviderRequestCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ClusterDnsProviderRequestProvider) SetCredentials(v Route53DnsProviderRequestCredentials)`

SetCredentials sets Credentials field to given value.


### GetAwsRegion

`func (o *ClusterDnsProviderRequestProvider) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *ClusterDnsProviderRequestProvider) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *ClusterDnsProviderRequestProvider) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### GetHostedZoneId

`func (o *ClusterDnsProviderRequestProvider) GetHostedZoneId() string`

GetHostedZoneId returns the HostedZoneId field if non-nil, zero value otherwise.

### GetHostedZoneIdOk

`func (o *ClusterDnsProviderRequestProvider) GetHostedZoneIdOk() (*string, bool)`

GetHostedZoneIdOk returns a tuple with the HostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneId

`func (o *ClusterDnsProviderRequestProvider) SetHostedZoneId(v string)`

SetHostedZoneId sets HostedZoneId field to given value.

### HasHostedZoneId

`func (o *ClusterDnsProviderRequestProvider) HasHostedZoneId() bool`

HasHostedZoneId returns a boolean if a field has been set.

### SetHostedZoneIdNil

`func (o *ClusterDnsProviderRequestProvider) SetHostedZoneIdNil(b bool)`

 SetHostedZoneIdNil sets the value for HostedZoneId to be an explicit nil

### UnsetHostedZoneId
`func (o *ClusterDnsProviderRequestProvider) UnsetHostedZoneId()`

UnsetHostedZoneId ensures that no value is present for HostedZoneId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


