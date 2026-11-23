# EnvironmentCustomDomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain name. | 
**GenerateCertificate** | **bool** | Whether Qovery generates a certificate for the domain. | 
**UseCdn** | **bool** | Whether the domain is behind a CDN. | 
**ImpactedServices** | [**[]ImpactedServiceResponse**](ImpactedServiceResponse.md) | The services whose custom domains changed, removed services included. They need a redeploy to serve the change. | 

## Methods

### NewEnvironmentCustomDomainResponse

`func NewEnvironmentCustomDomainResponse(domain string, generateCertificate bool, useCdn bool, impactedServices []ImpactedServiceResponse, ) *EnvironmentCustomDomainResponse`

NewEnvironmentCustomDomainResponse instantiates a new EnvironmentCustomDomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentCustomDomainResponseWithDefaults

`func NewEnvironmentCustomDomainResponseWithDefaults() *EnvironmentCustomDomainResponse`

NewEnvironmentCustomDomainResponseWithDefaults instantiates a new EnvironmentCustomDomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *EnvironmentCustomDomainResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EnvironmentCustomDomainResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EnvironmentCustomDomainResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetGenerateCertificate

`func (o *EnvironmentCustomDomainResponse) GetGenerateCertificate() bool`

GetGenerateCertificate returns the GenerateCertificate field if non-nil, zero value otherwise.

### GetGenerateCertificateOk

`func (o *EnvironmentCustomDomainResponse) GetGenerateCertificateOk() (*bool, bool)`

GetGenerateCertificateOk returns a tuple with the GenerateCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCertificate

`func (o *EnvironmentCustomDomainResponse) SetGenerateCertificate(v bool)`

SetGenerateCertificate sets GenerateCertificate field to given value.


### GetUseCdn

`func (o *EnvironmentCustomDomainResponse) GetUseCdn() bool`

GetUseCdn returns the UseCdn field if non-nil, zero value otherwise.

### GetUseCdnOk

`func (o *EnvironmentCustomDomainResponse) GetUseCdnOk() (*bool, bool)`

GetUseCdnOk returns a tuple with the UseCdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCdn

`func (o *EnvironmentCustomDomainResponse) SetUseCdn(v bool)`

SetUseCdn sets UseCdn field to given value.


### GetImpactedServices

`func (o *EnvironmentCustomDomainResponse) GetImpactedServices() []ImpactedServiceResponse`

GetImpactedServices returns the ImpactedServices field if non-nil, zero value otherwise.

### GetImpactedServicesOk

`func (o *EnvironmentCustomDomainResponse) GetImpactedServicesOk() (*[]ImpactedServiceResponse, bool)`

GetImpactedServicesOk returns a tuple with the ImpactedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedServices

`func (o *EnvironmentCustomDomainResponse) SetImpactedServices(v []ImpactedServiceResponse)`

SetImpactedServices sets ImpactedServices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


