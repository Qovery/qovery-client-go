# EnvironmentCustomDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain name. | 
**GenerateCertificate** | Pointer to **bool** | Whether Qovery generates a certificate for the domain. Set it to &#x60;false&#x60; when a CDN or another party manages the certificate. | [optional] [default to true]
**UseCdn** | Pointer to **bool** | Whether the domain is behind a CDN. If &#x60;true&#x60;, Qovery only checks that the domain points to an IP; otherwise it checks that the domain resolves to the service load balancer. | [optional] [default to false]
**Services** | [**[]CustomDomainServiceRequest**](CustomDomainServiceRequest.md) | Every service the domain routes to, each with one &#x60;/&#x60; route served by its default port. Between 1 and 50 services; more than one only when the environment &#x60;routing_capabilities.can_share&#x60; is &#x60;true&#x60;. | 

## Methods

### NewEnvironmentCustomDomainRequest

`func NewEnvironmentCustomDomainRequest(domain string, services []CustomDomainServiceRequest, ) *EnvironmentCustomDomainRequest`

NewEnvironmentCustomDomainRequest instantiates a new EnvironmentCustomDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentCustomDomainRequestWithDefaults

`func NewEnvironmentCustomDomainRequestWithDefaults() *EnvironmentCustomDomainRequest`

NewEnvironmentCustomDomainRequestWithDefaults instantiates a new EnvironmentCustomDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *EnvironmentCustomDomainRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EnvironmentCustomDomainRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EnvironmentCustomDomainRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetGenerateCertificate

`func (o *EnvironmentCustomDomainRequest) GetGenerateCertificate() bool`

GetGenerateCertificate returns the GenerateCertificate field if non-nil, zero value otherwise.

### GetGenerateCertificateOk

`func (o *EnvironmentCustomDomainRequest) GetGenerateCertificateOk() (*bool, bool)`

GetGenerateCertificateOk returns a tuple with the GenerateCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCertificate

`func (o *EnvironmentCustomDomainRequest) SetGenerateCertificate(v bool)`

SetGenerateCertificate sets GenerateCertificate field to given value.

### HasGenerateCertificate

`func (o *EnvironmentCustomDomainRequest) HasGenerateCertificate() bool`

HasGenerateCertificate returns a boolean if a field has been set.

### GetUseCdn

`func (o *EnvironmentCustomDomainRequest) GetUseCdn() bool`

GetUseCdn returns the UseCdn field if non-nil, zero value otherwise.

### GetUseCdnOk

`func (o *EnvironmentCustomDomainRequest) GetUseCdnOk() (*bool, bool)`

GetUseCdnOk returns a tuple with the UseCdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCdn

`func (o *EnvironmentCustomDomainRequest) SetUseCdn(v bool)`

SetUseCdn sets UseCdn field to given value.

### HasUseCdn

`func (o *EnvironmentCustomDomainRequest) HasUseCdn() bool`

HasUseCdn returns a boolean if a field has been set.

### GetServices

`func (o *EnvironmentCustomDomainRequest) GetServices() []CustomDomainServiceRequest`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *EnvironmentCustomDomainRequest) GetServicesOk() (*[]CustomDomainServiceRequest, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *EnvironmentCustomDomainRequest) SetServices(v []CustomDomainServiceRequest)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


