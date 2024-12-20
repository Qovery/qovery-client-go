# OrganizationEventResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**OrganizationEventResponseListLinks**](OrganizationEventResponseListLinks.md) |  | [optional] 
**OrganizationMaxLimitReached** | Pointer to **bool** | Indicates if you cannot see previous logs according to your organization max limit | [optional] 
**Events** | Pointer to [**[]OrganizationEventResponse**](OrganizationEventResponse.md) |  | [optional] 

## Methods

### NewOrganizationEventResponseList

`func NewOrganizationEventResponseList() *OrganizationEventResponseList`

NewOrganizationEventResponseList instantiates a new OrganizationEventResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationEventResponseListWithDefaults

`func NewOrganizationEventResponseListWithDefaults() *OrganizationEventResponseList`

NewOrganizationEventResponseListWithDefaults instantiates a new OrganizationEventResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *OrganizationEventResponseList) GetLinks() OrganizationEventResponseListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrganizationEventResponseList) GetLinksOk() (*OrganizationEventResponseListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrganizationEventResponseList) SetLinks(v OrganizationEventResponseListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrganizationEventResponseList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOrganizationMaxLimitReached

`func (o *OrganizationEventResponseList) GetOrganizationMaxLimitReached() bool`

GetOrganizationMaxLimitReached returns the OrganizationMaxLimitReached field if non-nil, zero value otherwise.

### GetOrganizationMaxLimitReachedOk

`func (o *OrganizationEventResponseList) GetOrganizationMaxLimitReachedOk() (*bool, bool)`

GetOrganizationMaxLimitReachedOk returns a tuple with the OrganizationMaxLimitReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationMaxLimitReached

`func (o *OrganizationEventResponseList) SetOrganizationMaxLimitReached(v bool)`

SetOrganizationMaxLimitReached sets OrganizationMaxLimitReached field to given value.

### HasOrganizationMaxLimitReached

`func (o *OrganizationEventResponseList) HasOrganizationMaxLimitReached() bool`

HasOrganizationMaxLimitReached returns a boolean if a field has been set.

### GetEvents

`func (o *OrganizationEventResponseList) GetEvents() []OrganizationEventResponse`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OrganizationEventResponseList) GetEventsOk() (*[]OrganizationEventResponse, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OrganizationEventResponseList) SetEvents(v []OrganizationEventResponse)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OrganizationEventResponseList) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


