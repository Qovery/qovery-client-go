# CustomDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | your custom domain | 
**GenerateCertificate** | **bool** | to control if a certificate has to be generated for this custom domain by Qovery. The default value is &#x60;true&#x60;. This flag should be set to &#x60;false&#x60; if a CDN or other entities are managing the certificate for the specified domain and the traffic is proxied by the CDN to Qovery. | 

## Methods

### NewCustomDomainRequest

`func NewCustomDomainRequest(domain string, generateCertificate bool, ) *CustomDomainRequest`

NewCustomDomainRequest instantiates a new CustomDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainRequestWithDefaults

`func NewCustomDomainRequestWithDefaults() *CustomDomainRequest`

NewCustomDomainRequestWithDefaults instantiates a new CustomDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *CustomDomainRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CustomDomainRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CustomDomainRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetGenerateCertificate

`func (o *CustomDomainRequest) GetGenerateCertificate() bool`

GetGenerateCertificate returns the GenerateCertificate field if non-nil, zero value otherwise.

### GetGenerateCertificateOk

`func (o *CustomDomainRequest) GetGenerateCertificateOk() (*bool, bool)`

GetGenerateCertificateOk returns a tuple with the GenerateCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCertificate

`func (o *CustomDomainRequest) SetGenerateCertificate(v bool)`

SetGenerateCertificate sets GenerateCertificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


