# EventPaginatedResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **float32** |  | 
**PageSize** | **float32** |  | 
**Results** | Pointer to [**[]Event**](Event.md) |  | [optional] 

## Methods

### NewEventPaginatedResponseList

`func NewEventPaginatedResponseList(page float32, pageSize float32, ) *EventPaginatedResponseList`

NewEventPaginatedResponseList instantiates a new EventPaginatedResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPaginatedResponseListWithDefaults

`func NewEventPaginatedResponseListWithDefaults() *EventPaginatedResponseList`

NewEventPaginatedResponseListWithDefaults instantiates a new EventPaginatedResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *EventPaginatedResponseList) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EventPaginatedResponseList) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EventPaginatedResponseList) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *EventPaginatedResponseList) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EventPaginatedResponseList) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EventPaginatedResponseList) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *EventPaginatedResponseList) GetResults() []Event`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *EventPaginatedResponseList) GetResultsOk() (*[]Event, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *EventPaginatedResponseList) SetResults(v []Event)`

SetResults sets Results field to given value.

### HasResults

`func (o *EventPaginatedResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


