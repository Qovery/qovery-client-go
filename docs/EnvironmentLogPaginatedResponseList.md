# EnvironmentLogPaginatedResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **float32** |  | 
**PageSize** | **float32** |  | 
**Results** | Pointer to [**[]EnvironmentLog**](EnvironmentLog.md) |  | [optional] 

## Methods

### NewEnvironmentLogPaginatedResponseList

`func NewEnvironmentLogPaginatedResponseList(page float32, pageSize float32, ) *EnvironmentLogPaginatedResponseList`

NewEnvironmentLogPaginatedResponseList instantiates a new EnvironmentLogPaginatedResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentLogPaginatedResponseListWithDefaults

`func NewEnvironmentLogPaginatedResponseListWithDefaults() *EnvironmentLogPaginatedResponseList`

NewEnvironmentLogPaginatedResponseListWithDefaults instantiates a new EnvironmentLogPaginatedResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *EnvironmentLogPaginatedResponseList) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EnvironmentLogPaginatedResponseList) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EnvironmentLogPaginatedResponseList) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *EnvironmentLogPaginatedResponseList) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EnvironmentLogPaginatedResponseList) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EnvironmentLogPaginatedResponseList) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *EnvironmentLogPaginatedResponseList) GetResults() []EnvironmentLog`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *EnvironmentLogPaginatedResponseList) GetResultsOk() (*[]EnvironmentLog, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *EnvironmentLogPaginatedResponseList) SetResults(v []EnvironmentLog)`

SetResults sets Results field to given value.

### HasResults

`func (o *EnvironmentLogPaginatedResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


