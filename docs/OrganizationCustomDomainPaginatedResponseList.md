# OrganizationCustomDomainPaginatedResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **float32** |  | 
**PageSize** | **float32** |  | 
**Results** | [**[]OrganizationCustomDomainResponse**](OrganizationCustomDomainResponse.md) |  | 
**TotalCount** | **int32** | The number of domains matching the query, on all pages. | 

## Methods

### NewOrganizationCustomDomainPaginatedResponseList

`func NewOrganizationCustomDomainPaginatedResponseList(page float32, pageSize float32, results []OrganizationCustomDomainResponse, totalCount int32, ) *OrganizationCustomDomainPaginatedResponseList`

NewOrganizationCustomDomainPaginatedResponseList instantiates a new OrganizationCustomDomainPaginatedResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomDomainPaginatedResponseListWithDefaults

`func NewOrganizationCustomDomainPaginatedResponseListWithDefaults() *OrganizationCustomDomainPaginatedResponseList`

NewOrganizationCustomDomainPaginatedResponseListWithDefaults instantiates a new OrganizationCustomDomainPaginatedResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *OrganizationCustomDomainPaginatedResponseList) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrganizationCustomDomainPaginatedResponseList) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrganizationCustomDomainPaginatedResponseList) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *OrganizationCustomDomainPaginatedResponseList) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OrganizationCustomDomainPaginatedResponseList) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OrganizationCustomDomainPaginatedResponseList) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *OrganizationCustomDomainPaginatedResponseList) GetResults() []OrganizationCustomDomainResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OrganizationCustomDomainPaginatedResponseList) GetResultsOk() (*[]OrganizationCustomDomainResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OrganizationCustomDomainPaginatedResponseList) SetResults(v []OrganizationCustomDomainResponse)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *OrganizationCustomDomainPaginatedResponseList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OrganizationCustomDomainPaginatedResponseList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OrganizationCustomDomainPaginatedResponseList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


