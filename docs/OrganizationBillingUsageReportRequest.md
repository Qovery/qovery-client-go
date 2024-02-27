# OrganizationBillingUsageReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **time.Time** | The start date of the report | [optional] 
**To** | Pointer to **time.Time** | The end date of the report | [optional] 
**ReportExpirationInSeconds** | Pointer to **int32** | The number of seconds the report will be publicly available | [optional] 

## Methods

### NewOrganizationBillingUsageReportRequest

`func NewOrganizationBillingUsageReportRequest() *OrganizationBillingUsageReportRequest`

NewOrganizationBillingUsageReportRequest instantiates a new OrganizationBillingUsageReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationBillingUsageReportRequestWithDefaults

`func NewOrganizationBillingUsageReportRequestWithDefaults() *OrganizationBillingUsageReportRequest`

NewOrganizationBillingUsageReportRequestWithDefaults instantiates a new OrganizationBillingUsageReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *OrganizationBillingUsageReportRequest) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *OrganizationBillingUsageReportRequest) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *OrganizationBillingUsageReportRequest) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *OrganizationBillingUsageReportRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *OrganizationBillingUsageReportRequest) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *OrganizationBillingUsageReportRequest) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *OrganizationBillingUsageReportRequest) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *OrganizationBillingUsageReportRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetReportExpirationInSeconds

`func (o *OrganizationBillingUsageReportRequest) GetReportExpirationInSeconds() int32`

GetReportExpirationInSeconds returns the ReportExpirationInSeconds field if non-nil, zero value otherwise.

### GetReportExpirationInSecondsOk

`func (o *OrganizationBillingUsageReportRequest) GetReportExpirationInSecondsOk() (*int32, bool)`

GetReportExpirationInSecondsOk returns a tuple with the ReportExpirationInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportExpirationInSeconds

`func (o *OrganizationBillingUsageReportRequest) SetReportExpirationInSeconds(v int32)`

SetReportExpirationInSeconds sets ReportExpirationInSeconds field to given value.

### HasReportExpirationInSeconds

`func (o *OrganizationBillingUsageReportRequest) HasReportExpirationInSeconds() bool`

HasReportExpirationInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


