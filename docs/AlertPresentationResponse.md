# AlertPresentationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to **NullableString** |  Alert summary template | [optional] 
**RunbookUrl** | Pointer to **NullableString** | URL to runbook with remediation steps | [optional] 

## Methods

### NewAlertPresentationResponse

`func NewAlertPresentationResponse() *AlertPresentationResponse`

NewAlertPresentationResponse instantiates a new AlertPresentationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertPresentationResponseWithDefaults

`func NewAlertPresentationResponseWithDefaults() *AlertPresentationResponse`

NewAlertPresentationResponseWithDefaults instantiates a new AlertPresentationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *AlertPresentationResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AlertPresentationResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AlertPresentationResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AlertPresentationResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *AlertPresentationResponse) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *AlertPresentationResponse) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetRunbookUrl

`func (o *AlertPresentationResponse) GetRunbookUrl() string`

GetRunbookUrl returns the RunbookUrl field if non-nil, zero value otherwise.

### GetRunbookUrlOk

`func (o *AlertPresentationResponse) GetRunbookUrlOk() (*string, bool)`

GetRunbookUrlOk returns a tuple with the RunbookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunbookUrl

`func (o *AlertPresentationResponse) SetRunbookUrl(v string)`

SetRunbookUrl sets RunbookUrl field to given value.

### HasRunbookUrl

`func (o *AlertPresentationResponse) HasRunbookUrl() bool`

HasRunbookUrl returns a boolean if a field has been set.

### SetRunbookUrlNil

`func (o *AlertPresentationResponse) SetRunbookUrlNil(b bool)`

 SetRunbookUrlNil sets the value for RunbookUrl to be an explicit nil

### UnsetRunbookUrl
`func (o *AlertPresentationResponse) UnsetRunbookUrl()`

UnsetRunbookUrl ensures that no value is present for RunbookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


