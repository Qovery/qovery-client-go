# Route53DnsProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Domain** | **string** |  | 
**Credentials** | [**Route53DnsProviderResponseCredentials**](Route53DnsProviderResponseCredentials.md) |  | 
**AwsRegion** | **string** |  | 
**HostedZoneId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRoute53DnsProviderResponse

`func NewRoute53DnsProviderResponse(provider string, domain string, credentials Route53DnsProviderResponseCredentials, awsRegion string, ) *Route53DnsProviderResponse`

NewRoute53DnsProviderResponse instantiates a new Route53DnsProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoute53DnsProviderResponseWithDefaults

`func NewRoute53DnsProviderResponseWithDefaults() *Route53DnsProviderResponse`

NewRoute53DnsProviderResponseWithDefaults instantiates a new Route53DnsProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *Route53DnsProviderResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Route53DnsProviderResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Route53DnsProviderResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetDomain

`func (o *Route53DnsProviderResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Route53DnsProviderResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Route53DnsProviderResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetCredentials

`func (o *Route53DnsProviderResponse) GetCredentials() Route53DnsProviderResponseCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Route53DnsProviderResponse) GetCredentialsOk() (*Route53DnsProviderResponseCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Route53DnsProviderResponse) SetCredentials(v Route53DnsProviderResponseCredentials)`

SetCredentials sets Credentials field to given value.


### GetAwsRegion

`func (o *Route53DnsProviderResponse) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *Route53DnsProviderResponse) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *Route53DnsProviderResponse) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### GetHostedZoneId

`func (o *Route53DnsProviderResponse) GetHostedZoneId() string`

GetHostedZoneId returns the HostedZoneId field if non-nil, zero value otherwise.

### GetHostedZoneIdOk

`func (o *Route53DnsProviderResponse) GetHostedZoneIdOk() (*string, bool)`

GetHostedZoneIdOk returns a tuple with the HostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneId

`func (o *Route53DnsProviderResponse) SetHostedZoneId(v string)`

SetHostedZoneId sets HostedZoneId field to given value.

### HasHostedZoneId

`func (o *Route53DnsProviderResponse) HasHostedZoneId() bool`

HasHostedZoneId returns a boolean if a field has been set.

### SetHostedZoneIdNil

`func (o *Route53DnsProviderResponse) SetHostedZoneIdNil(b bool)`

 SetHostedZoneIdNil sets the value for HostedZoneId to be an explicit nil

### UnsetHostedZoneId
`func (o *Route53DnsProviderResponse) UnsetHostedZoneId()`

UnsetHostedZoneId ensures that no value is present for HostedZoneId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


