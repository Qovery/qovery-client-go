# OrganizationCustomDomainCheckListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]OrganizationCustomDomainCheckResponse**](OrganizationCustomDomainCheckResponse.md) | One result per domain name of one environment of the page, in the order of &#x60;listOrganizationCustomDomain&#x60;. | 

## Methods

### NewOrganizationCustomDomainCheckListResponse

`func NewOrganizationCustomDomainCheckListResponse(results []OrganizationCustomDomainCheckResponse, ) *OrganizationCustomDomainCheckListResponse`

NewOrganizationCustomDomainCheckListResponse instantiates a new OrganizationCustomDomainCheckListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomDomainCheckListResponseWithDefaults

`func NewOrganizationCustomDomainCheckListResponseWithDefaults() *OrganizationCustomDomainCheckListResponse`

NewOrganizationCustomDomainCheckListResponseWithDefaults instantiates a new OrganizationCustomDomainCheckListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *OrganizationCustomDomainCheckListResponse) GetResults() []OrganizationCustomDomainCheckResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OrganizationCustomDomainCheckListResponse) GetResultsOk() (*[]OrganizationCustomDomainCheckResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OrganizationCustomDomainCheckListResponse) SetResults(v []OrganizationCustomDomainCheckResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


