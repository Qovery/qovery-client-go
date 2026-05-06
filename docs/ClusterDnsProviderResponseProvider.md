# ClusterDnsProviderResponseProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Domain** | **string** |  | 
**Email** | **string** |  | 
**Proxied** | **bool** |  | 
**Credentials** | [**Route53DnsProviderResponseCredentials**](Route53DnsProviderResponseCredentials.md) |  | 
**AwsRegion** | **string** |  | 
**HostedZoneId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewClusterDnsProviderResponseProvider

`func NewClusterDnsProviderResponseProvider(provider string, domain string, email string, proxied bool, credentials Route53DnsProviderResponseCredentials, awsRegion string, ) *ClusterDnsProviderResponseProvider`

NewClusterDnsProviderResponseProvider instantiates a new ClusterDnsProviderResponseProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDnsProviderResponseProviderWithDefaults

`func NewClusterDnsProviderResponseProviderWithDefaults() *ClusterDnsProviderResponseProvider`

NewClusterDnsProviderResponseProviderWithDefaults instantiates a new ClusterDnsProviderResponseProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ClusterDnsProviderResponseProvider) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ClusterDnsProviderResponseProvider) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ClusterDnsProviderResponseProvider) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetDomain

`func (o *ClusterDnsProviderResponseProvider) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ClusterDnsProviderResponseProvider) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ClusterDnsProviderResponseProvider) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEmail

`func (o *ClusterDnsProviderResponseProvider) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ClusterDnsProviderResponseProvider) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ClusterDnsProviderResponseProvider) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProxied

`func (o *ClusterDnsProviderResponseProvider) GetProxied() bool`

GetProxied returns the Proxied field if non-nil, zero value otherwise.

### GetProxiedOk

`func (o *ClusterDnsProviderResponseProvider) GetProxiedOk() (*bool, bool)`

GetProxiedOk returns a tuple with the Proxied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxied

`func (o *ClusterDnsProviderResponseProvider) SetProxied(v bool)`

SetProxied sets Proxied field to given value.


### GetCredentials

`func (o *ClusterDnsProviderResponseProvider) GetCredentials() Route53DnsProviderResponseCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ClusterDnsProviderResponseProvider) GetCredentialsOk() (*Route53DnsProviderResponseCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ClusterDnsProviderResponseProvider) SetCredentials(v Route53DnsProviderResponseCredentials)`

SetCredentials sets Credentials field to given value.


### GetAwsRegion

`func (o *ClusterDnsProviderResponseProvider) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *ClusterDnsProviderResponseProvider) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *ClusterDnsProviderResponseProvider) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### GetHostedZoneId

`func (o *ClusterDnsProviderResponseProvider) GetHostedZoneId() string`

GetHostedZoneId returns the HostedZoneId field if non-nil, zero value otherwise.

### GetHostedZoneIdOk

`func (o *ClusterDnsProviderResponseProvider) GetHostedZoneIdOk() (*string, bool)`

GetHostedZoneIdOk returns a tuple with the HostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneId

`func (o *ClusterDnsProviderResponseProvider) SetHostedZoneId(v string)`

SetHostedZoneId sets HostedZoneId field to given value.

### HasHostedZoneId

`func (o *ClusterDnsProviderResponseProvider) HasHostedZoneId() bool`

HasHostedZoneId returns a boolean if a field has been set.

### SetHostedZoneIdNil

`func (o *ClusterDnsProviderResponseProvider) SetHostedZoneIdNil(b bool)`

 SetHostedZoneIdNil sets the value for HostedZoneId to be an explicit nil

### UnsetHostedZoneId
`func (o *ClusterDnsProviderResponseProvider) UnsetHostedZoneId()`

UnsetHostedZoneId ensures that no value is present for HostedZoneId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


