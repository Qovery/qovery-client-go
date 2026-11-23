# CustomDomainValidationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain name. | 
**GenerateCertificate** | Pointer to **bool** | Whether Qovery generates a certificate for the domain. Set it to &#x60;false&#x60; when a CDN or another party manages the certificate. | [optional] [default to true]
**UseCdn** | Pointer to **bool** | Whether the domain is behind a CDN. If &#x60;true&#x60;, Qovery only checks that the domain points to an IP; otherwise it checks that the domain resolves to the service load balancer. | [optional] [default to false]
**Services** | [**[]CustomDomainServiceRequest**](CustomDomainServiceRequest.md) | Every service the domain routes to, each with one &#x60;/&#x60; route served by its default port. Between 1 and 50 services; more than one only when the environment &#x60;routing_capabilities.can_share&#x60; is &#x60;true&#x60;. | 
**OriginalDomain** | Pointer to **NullableString** | The current name of a domain being edited, so the rows the services hold for it are ignored. Null or absent for a new domain. | [optional] 

## Methods

### NewCustomDomainValidationRequest

`func NewCustomDomainValidationRequest(domain string, services []CustomDomainServiceRequest, ) *CustomDomainValidationRequest`

NewCustomDomainValidationRequest instantiates a new CustomDomainValidationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainValidationRequestWithDefaults

`func NewCustomDomainValidationRequestWithDefaults() *CustomDomainValidationRequest`

NewCustomDomainValidationRequestWithDefaults instantiates a new CustomDomainValidationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *CustomDomainValidationRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CustomDomainValidationRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CustomDomainValidationRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetGenerateCertificate

`func (o *CustomDomainValidationRequest) GetGenerateCertificate() bool`

GetGenerateCertificate returns the GenerateCertificate field if non-nil, zero value otherwise.

### GetGenerateCertificateOk

`func (o *CustomDomainValidationRequest) GetGenerateCertificateOk() (*bool, bool)`

GetGenerateCertificateOk returns a tuple with the GenerateCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCertificate

`func (o *CustomDomainValidationRequest) SetGenerateCertificate(v bool)`

SetGenerateCertificate sets GenerateCertificate field to given value.

### HasGenerateCertificate

`func (o *CustomDomainValidationRequest) HasGenerateCertificate() bool`

HasGenerateCertificate returns a boolean if a field has been set.

### GetUseCdn

`func (o *CustomDomainValidationRequest) GetUseCdn() bool`

GetUseCdn returns the UseCdn field if non-nil, zero value otherwise.

### GetUseCdnOk

`func (o *CustomDomainValidationRequest) GetUseCdnOk() (*bool, bool)`

GetUseCdnOk returns a tuple with the UseCdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCdn

`func (o *CustomDomainValidationRequest) SetUseCdn(v bool)`

SetUseCdn sets UseCdn field to given value.

### HasUseCdn

`func (o *CustomDomainValidationRequest) HasUseCdn() bool`

HasUseCdn returns a boolean if a field has been set.

### GetServices

`func (o *CustomDomainValidationRequest) GetServices() []CustomDomainServiceRequest`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CustomDomainValidationRequest) GetServicesOk() (*[]CustomDomainServiceRequest, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CustomDomainValidationRequest) SetServices(v []CustomDomainServiceRequest)`

SetServices sets Services field to given value.


### GetOriginalDomain

`func (o *CustomDomainValidationRequest) GetOriginalDomain() string`

GetOriginalDomain returns the OriginalDomain field if non-nil, zero value otherwise.

### GetOriginalDomainOk

`func (o *CustomDomainValidationRequest) GetOriginalDomainOk() (*string, bool)`

GetOriginalDomainOk returns a tuple with the OriginalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDomain

`func (o *CustomDomainValidationRequest) SetOriginalDomain(v string)`

SetOriginalDomain sets OriginalDomain field to given value.

### HasOriginalDomain

`func (o *CustomDomainValidationRequest) HasOriginalDomain() bool`

HasOriginalDomain returns a boolean if a field has been set.

### SetOriginalDomainNil

`func (o *CustomDomainValidationRequest) SetOriginalDomainNil(b bool)`

 SetOriginalDomainNil sets the value for OriginalDomain to be an explicit nil

### UnsetOriginalDomain
`func (o *CustomDomainValidationRequest) UnsetOriginalDomain()`

UnsetOriginalDomain ensures that no value is present for OriginalDomain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


