# OrganizationCustomDomainCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | Id of the environment of the domain. | 
**Domain** | **string** | The domain name as &#x60;listOrganizationCustomDomain&#x60; returns it. A wildcard domain keeps its &#x60;*.&#x60; prefix. | 
**Checks** | [**[]OrganizationCustomDomainServiceCheck**](OrganizationCustomDomainServiceCheck.md) | One check per service the domain routes to. Empty when no service of the domain can be checked. | 

## Methods

### NewOrganizationCustomDomainCheckResponse

`func NewOrganizationCustomDomainCheckResponse(environmentId string, domain string, checks []OrganizationCustomDomainServiceCheck, ) *OrganizationCustomDomainCheckResponse`

NewOrganizationCustomDomainCheckResponse instantiates a new OrganizationCustomDomainCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomDomainCheckResponseWithDefaults

`func NewOrganizationCustomDomainCheckResponseWithDefaults() *OrganizationCustomDomainCheckResponse`

NewOrganizationCustomDomainCheckResponseWithDefaults instantiates a new OrganizationCustomDomainCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *OrganizationCustomDomainCheckResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *OrganizationCustomDomainCheckResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *OrganizationCustomDomainCheckResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetDomain

`func (o *OrganizationCustomDomainCheckResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OrganizationCustomDomainCheckResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OrganizationCustomDomainCheckResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetChecks

`func (o *OrganizationCustomDomainCheckResponse) GetChecks() []OrganizationCustomDomainServiceCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *OrganizationCustomDomainCheckResponse) GetChecksOk() (*[]OrganizationCustomDomainServiceCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *OrganizationCustomDomainCheckResponse) SetChecks(v []OrganizationCustomDomainServiceCheck)`

SetChecks sets Checks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


