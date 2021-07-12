# DeploymentHistoryPaginatedResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]DeploymentHistoryResponse**](DeploymentHistoryResponse.md) |  | [optional] 
**Page** | **float32** |  | 
**PageSize** | **float32** |  | 

## Methods

### NewDeploymentHistoryPaginatedResponseList

`func NewDeploymentHistoryPaginatedResponseList(page float32, pageSize float32, ) *DeploymentHistoryPaginatedResponseList`

NewDeploymentHistoryPaginatedResponseList instantiates a new DeploymentHistoryPaginatedResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryPaginatedResponseListWithDefaults

`func NewDeploymentHistoryPaginatedResponseListWithDefaults() *DeploymentHistoryPaginatedResponseList`

NewDeploymentHistoryPaginatedResponseListWithDefaults instantiates a new DeploymentHistoryPaginatedResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *DeploymentHistoryPaginatedResponseList) GetResults() []DeploymentHistoryResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DeploymentHistoryPaginatedResponseList) GetResultsOk() (*[]DeploymentHistoryResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DeploymentHistoryPaginatedResponseList) SetResults(v []DeploymentHistoryResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *DeploymentHistoryPaginatedResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetPage

`func (o *DeploymentHistoryPaginatedResponseList) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DeploymentHistoryPaginatedResponseList) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DeploymentHistoryPaginatedResponseList) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *DeploymentHistoryPaginatedResponseList) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DeploymentHistoryPaginatedResponseList) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DeploymentHistoryPaginatedResponseList) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


