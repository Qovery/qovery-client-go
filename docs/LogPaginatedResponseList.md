# LogPaginatedResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]LogResponse**](LogResponse.md) |  | [optional] 
**Page** | **float32** |  | 
**PageSize** | **float32** |  | 

## Methods

### NewLogPaginatedResponseList

`func NewLogPaginatedResponseList(page float32, pageSize float32, ) *LogPaginatedResponseList`

NewLogPaginatedResponseList instantiates a new LogPaginatedResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogPaginatedResponseListWithDefaults

`func NewLogPaginatedResponseListWithDefaults() *LogPaginatedResponseList`

NewLogPaginatedResponseListWithDefaults instantiates a new LogPaginatedResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *LogPaginatedResponseList) GetResults() []LogResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *LogPaginatedResponseList) GetResultsOk() (*[]LogResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *LogPaginatedResponseList) SetResults(v []LogResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *LogPaginatedResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetPage

`func (o *LogPaginatedResponseList) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *LogPaginatedResponseList) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *LogPaginatedResponseList) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *LogPaginatedResponseList) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *LogPaginatedResponseList) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *LogPaginatedResponseList) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


