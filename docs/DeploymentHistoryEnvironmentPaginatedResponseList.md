# DeploymentHistoryEnvironmentPaginatedResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **float32** |  | 
**PageSize** | **float32** |  | 
**Results** | Pointer to [**[]DeploymentHistoryEnvironmentResponse**](DeploymentHistoryEnvironmentResponse.md) |  | [optional] 

## Methods

### NewDeploymentHistoryEnvironmentPaginatedResponseList

`func NewDeploymentHistoryEnvironmentPaginatedResponseList(page float32, pageSize float32, ) *DeploymentHistoryEnvironmentPaginatedResponseList`

NewDeploymentHistoryEnvironmentPaginatedResponseList instantiates a new DeploymentHistoryEnvironmentPaginatedResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryEnvironmentPaginatedResponseListWithDefaults

`func NewDeploymentHistoryEnvironmentPaginatedResponseListWithDefaults() *DeploymentHistoryEnvironmentPaginatedResponseList`

NewDeploymentHistoryEnvironmentPaginatedResponseListWithDefaults instantiates a new DeploymentHistoryEnvironmentPaginatedResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *DeploymentHistoryEnvironmentPaginatedResponseList) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DeploymentHistoryEnvironmentPaginatedResponseList) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DeploymentHistoryEnvironmentPaginatedResponseList) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *DeploymentHistoryEnvironmentPaginatedResponseList) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DeploymentHistoryEnvironmentPaginatedResponseList) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DeploymentHistoryEnvironmentPaginatedResponseList) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *DeploymentHistoryEnvironmentPaginatedResponseList) GetResults() []DeploymentHistoryEnvironmentResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DeploymentHistoryEnvironmentPaginatedResponseList) GetResultsOk() (*[]DeploymentHistoryEnvironmentResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DeploymentHistoryEnvironmentPaginatedResponseList) SetResults(v []DeploymentHistoryEnvironmentResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *DeploymentHistoryEnvironmentPaginatedResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


