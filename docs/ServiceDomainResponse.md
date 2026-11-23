# ServiceDomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**ServiceDomainKindEnum**](ServiceDomainKindEnum.md) |  | 
**Domain** | **string** |  | 
**GeneratedHostnames** | [**[]GeneratedHostnameResponse**](GeneratedHostnameResponse.md) | The hostnames served by the routes of this service only. | 
**GenerateCertificate** | **bool** | Whether Qovery generates a certificate for the domain. | 
**UseCdn** | **bool** | Whether the domain is behind a CDN. | 
**CnameTargets** | **[]string** | The hostnames the domain must point to with a CNAME record. One value, except for a legacy domain spread over several service routers. | 
**Routes** | [**[]CustomDomainRouteResponse**](CustomDomainRouteResponse.md) | Every route of the domain, the routes of the other services included. | 
**OtherServices** | [**[]CustomDomainServiceResponse**](CustomDomainServiceResponse.md) | The other services routing the domain. When it is not empty, removing this service&#39;s route keeps the domain for them. | 

## Methods

### NewServiceDomainResponse

`func NewServiceDomainResponse(kind ServiceDomainKindEnum, domain string, generatedHostnames []GeneratedHostnameResponse, generateCertificate bool, useCdn bool, cnameTargets []string, routes []CustomDomainRouteResponse, otherServices []CustomDomainServiceResponse, ) *ServiceDomainResponse`

NewServiceDomainResponse instantiates a new ServiceDomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDomainResponseWithDefaults

`func NewServiceDomainResponseWithDefaults() *ServiceDomainResponse`

NewServiceDomainResponseWithDefaults instantiates a new ServiceDomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ServiceDomainResponse) GetKind() ServiceDomainKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServiceDomainResponse) GetKindOk() (*ServiceDomainKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServiceDomainResponse) SetKind(v ServiceDomainKindEnum)`

SetKind sets Kind field to given value.


### GetDomain

`func (o *ServiceDomainResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ServiceDomainResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ServiceDomainResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetGeneratedHostnames

`func (o *ServiceDomainResponse) GetGeneratedHostnames() []GeneratedHostnameResponse`

GetGeneratedHostnames returns the GeneratedHostnames field if non-nil, zero value otherwise.

### GetGeneratedHostnamesOk

`func (o *ServiceDomainResponse) GetGeneratedHostnamesOk() (*[]GeneratedHostnameResponse, bool)`

GetGeneratedHostnamesOk returns a tuple with the GeneratedHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedHostnames

`func (o *ServiceDomainResponse) SetGeneratedHostnames(v []GeneratedHostnameResponse)`

SetGeneratedHostnames sets GeneratedHostnames field to given value.


### GetGenerateCertificate

`func (o *ServiceDomainResponse) GetGenerateCertificate() bool`

GetGenerateCertificate returns the GenerateCertificate field if non-nil, zero value otherwise.

### GetGenerateCertificateOk

`func (o *ServiceDomainResponse) GetGenerateCertificateOk() (*bool, bool)`

GetGenerateCertificateOk returns a tuple with the GenerateCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCertificate

`func (o *ServiceDomainResponse) SetGenerateCertificate(v bool)`

SetGenerateCertificate sets GenerateCertificate field to given value.


### GetUseCdn

`func (o *ServiceDomainResponse) GetUseCdn() bool`

GetUseCdn returns the UseCdn field if non-nil, zero value otherwise.

### GetUseCdnOk

`func (o *ServiceDomainResponse) GetUseCdnOk() (*bool, bool)`

GetUseCdnOk returns a tuple with the UseCdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCdn

`func (o *ServiceDomainResponse) SetUseCdn(v bool)`

SetUseCdn sets UseCdn field to given value.


### GetCnameTargets

`func (o *ServiceDomainResponse) GetCnameTargets() []string`

GetCnameTargets returns the CnameTargets field if non-nil, zero value otherwise.

### GetCnameTargetsOk

`func (o *ServiceDomainResponse) GetCnameTargetsOk() (*[]string, bool)`

GetCnameTargetsOk returns a tuple with the CnameTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameTargets

`func (o *ServiceDomainResponse) SetCnameTargets(v []string)`

SetCnameTargets sets CnameTargets field to given value.


### GetRoutes

`func (o *ServiceDomainResponse) GetRoutes() []CustomDomainRouteResponse`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *ServiceDomainResponse) GetRoutesOk() (*[]CustomDomainRouteResponse, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *ServiceDomainResponse) SetRoutes(v []CustomDomainRouteResponse)`

SetRoutes sets Routes field to given value.


### GetOtherServices

`func (o *ServiceDomainResponse) GetOtherServices() []CustomDomainServiceResponse`

GetOtherServices returns the OtherServices field if non-nil, zero value otherwise.

### GetOtherServicesOk

`func (o *ServiceDomainResponse) GetOtherServicesOk() (*[]CustomDomainServiceResponse, bool)`

GetOtherServicesOk returns a tuple with the OtherServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherServices

`func (o *ServiceDomainResponse) SetOtherServices(v []CustomDomainServiceResponse)`

SetOtherServices sets OtherServices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


