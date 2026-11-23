# OrganizationCustomDomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | Id of the environment of the domain. | 
**Domain** | **string** | The domain name. | 
**Project** | [**CustomDomainProjectResponse**](CustomDomainProjectResponse.md) |  | 
**Environment** | [**CustomDomainEnvironmentResponse**](CustomDomainEnvironmentResponse.md) |  | 
**Routes** | [**[]CustomDomainRouteResponse**](CustomDomainRouteResponse.md) | The routes of the domain, one per service and path. | 
**GenerateCertificate** | **bool** | Whether Qovery generates a certificate for the domain. &#x60;false&#x60; when a CDN or another party manages the certificate. | 
**UseCdn** | **bool** | Whether the domain is behind a CDN. If &#x60;true&#x60;, Qovery only checks that the domain points to an IP; otherwise it checks that the domain resolves to the service load balancer. | 
**CnameTargets** | Pointer to **[]string** | The hostnames the domain must point to with a CNAME record. One value, except for a legacy domain spread over several service routers. | [optional] 
**GeneratedHostnames** | Pointer to [**[]GeneratedHostnameResponse**](GeneratedHostnameResponse.md) | Every hostname served for the domain by its routes. | [optional] 

## Methods

### NewOrganizationCustomDomainResponse

`func NewOrganizationCustomDomainResponse(environmentId string, domain string, project CustomDomainProjectResponse, environment CustomDomainEnvironmentResponse, routes []CustomDomainRouteResponse, generateCertificate bool, useCdn bool, ) *OrganizationCustomDomainResponse`

NewOrganizationCustomDomainResponse instantiates a new OrganizationCustomDomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomDomainResponseWithDefaults

`func NewOrganizationCustomDomainResponseWithDefaults() *OrganizationCustomDomainResponse`

NewOrganizationCustomDomainResponseWithDefaults instantiates a new OrganizationCustomDomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *OrganizationCustomDomainResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *OrganizationCustomDomainResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *OrganizationCustomDomainResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetDomain

`func (o *OrganizationCustomDomainResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OrganizationCustomDomainResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OrganizationCustomDomainResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetProject

`func (o *OrganizationCustomDomainResponse) GetProject() CustomDomainProjectResponse`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *OrganizationCustomDomainResponse) GetProjectOk() (*CustomDomainProjectResponse, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *OrganizationCustomDomainResponse) SetProject(v CustomDomainProjectResponse)`

SetProject sets Project field to given value.


### GetEnvironment

`func (o *OrganizationCustomDomainResponse) GetEnvironment() CustomDomainEnvironmentResponse`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *OrganizationCustomDomainResponse) GetEnvironmentOk() (*CustomDomainEnvironmentResponse, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *OrganizationCustomDomainResponse) SetEnvironment(v CustomDomainEnvironmentResponse)`

SetEnvironment sets Environment field to given value.


### GetRoutes

`func (o *OrganizationCustomDomainResponse) GetRoutes() []CustomDomainRouteResponse`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *OrganizationCustomDomainResponse) GetRoutesOk() (*[]CustomDomainRouteResponse, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *OrganizationCustomDomainResponse) SetRoutes(v []CustomDomainRouteResponse)`

SetRoutes sets Routes field to given value.


### GetGenerateCertificate

`func (o *OrganizationCustomDomainResponse) GetGenerateCertificate() bool`

GetGenerateCertificate returns the GenerateCertificate field if non-nil, zero value otherwise.

### GetGenerateCertificateOk

`func (o *OrganizationCustomDomainResponse) GetGenerateCertificateOk() (*bool, bool)`

GetGenerateCertificateOk returns a tuple with the GenerateCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCertificate

`func (o *OrganizationCustomDomainResponse) SetGenerateCertificate(v bool)`

SetGenerateCertificate sets GenerateCertificate field to given value.


### GetUseCdn

`func (o *OrganizationCustomDomainResponse) GetUseCdn() bool`

GetUseCdn returns the UseCdn field if non-nil, zero value otherwise.

### GetUseCdnOk

`func (o *OrganizationCustomDomainResponse) GetUseCdnOk() (*bool, bool)`

GetUseCdnOk returns a tuple with the UseCdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCdn

`func (o *OrganizationCustomDomainResponse) SetUseCdn(v bool)`

SetUseCdn sets UseCdn field to given value.


### GetCnameTargets

`func (o *OrganizationCustomDomainResponse) GetCnameTargets() []string`

GetCnameTargets returns the CnameTargets field if non-nil, zero value otherwise.

### GetCnameTargetsOk

`func (o *OrganizationCustomDomainResponse) GetCnameTargetsOk() (*[]string, bool)`

GetCnameTargetsOk returns a tuple with the CnameTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameTargets

`func (o *OrganizationCustomDomainResponse) SetCnameTargets(v []string)`

SetCnameTargets sets CnameTargets field to given value.

### HasCnameTargets

`func (o *OrganizationCustomDomainResponse) HasCnameTargets() bool`

HasCnameTargets returns a boolean if a field has been set.

### GetGeneratedHostnames

`func (o *OrganizationCustomDomainResponse) GetGeneratedHostnames() []GeneratedHostnameResponse`

GetGeneratedHostnames returns the GeneratedHostnames field if non-nil, zero value otherwise.

### GetGeneratedHostnamesOk

`func (o *OrganizationCustomDomainResponse) GetGeneratedHostnamesOk() (*[]GeneratedHostnameResponse, bool)`

GetGeneratedHostnamesOk returns a tuple with the GeneratedHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedHostnames

`func (o *OrganizationCustomDomainResponse) SetGeneratedHostnames(v []GeneratedHostnameResponse)`

SetGeneratedHostnames sets GeneratedHostnames field to given value.

### HasGeneratedHostnames

`func (o *OrganizationCustomDomainResponse) HasGeneratedHostnames() bool`

HasGeneratedHostnames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


