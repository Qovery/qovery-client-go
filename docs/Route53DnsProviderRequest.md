# Route53DnsProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Domain** | **string** |  | 
**Credentials** | [**Route53DnsProviderRequestCredentials**](Route53DnsProviderRequestCredentials.md) |  | 
**AwsRegion** | **string** |  | 
**HostedZoneId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRoute53DnsProviderRequest

`func NewRoute53DnsProviderRequest(provider string, domain string, credentials Route53DnsProviderRequestCredentials, awsRegion string, ) *Route53DnsProviderRequest`

NewRoute53DnsProviderRequest instantiates a new Route53DnsProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoute53DnsProviderRequestWithDefaults

`func NewRoute53DnsProviderRequestWithDefaults() *Route53DnsProviderRequest`

NewRoute53DnsProviderRequestWithDefaults instantiates a new Route53DnsProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *Route53DnsProviderRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Route53DnsProviderRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Route53DnsProviderRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetDomain

`func (o *Route53DnsProviderRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Route53DnsProviderRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Route53DnsProviderRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetCredentials

`func (o *Route53DnsProviderRequest) GetCredentials() Route53DnsProviderRequestCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Route53DnsProviderRequest) GetCredentialsOk() (*Route53DnsProviderRequestCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Route53DnsProviderRequest) SetCredentials(v Route53DnsProviderRequestCredentials)`

SetCredentials sets Credentials field to given value.


### GetAwsRegion

`func (o *Route53DnsProviderRequest) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *Route53DnsProviderRequest) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *Route53DnsProviderRequest) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### GetHostedZoneId

`func (o *Route53DnsProviderRequest) GetHostedZoneId() string`

GetHostedZoneId returns the HostedZoneId field if non-nil, zero value otherwise.

### GetHostedZoneIdOk

`func (o *Route53DnsProviderRequest) GetHostedZoneIdOk() (*string, bool)`

GetHostedZoneIdOk returns a tuple with the HostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneId

`func (o *Route53DnsProviderRequest) SetHostedZoneId(v string)`

SetHostedZoneId sets HostedZoneId field to given value.

### HasHostedZoneId

`func (o *Route53DnsProviderRequest) HasHostedZoneId() bool`

HasHostedZoneId returns a boolean if a field has been set.

### SetHostedZoneIdNil

`func (o *Route53DnsProviderRequest) SetHostedZoneIdNil(b bool)`

 SetHostedZoneIdNil sets the value for HostedZoneId to be an explicit nil

### UnsetHostedZoneId
`func (o *Route53DnsProviderRequest) UnsetHostedZoneId()`

UnsetHostedZoneId ensures that no value is present for HostedZoneId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


