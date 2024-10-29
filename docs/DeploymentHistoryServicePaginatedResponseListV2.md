# DeploymentHistoryServicePaginatedResponseListV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **float32** |  | 
**PageSize** | **float32** |  | 
**Results** | Pointer to [**[]DeploymentHistoryService**](DeploymentHistoryService.md) |  | [optional] 

## Methods

### NewDeploymentHistoryServicePaginatedResponseListV2

`func NewDeploymentHistoryServicePaginatedResponseListV2(page float32, pageSize float32, ) *DeploymentHistoryServicePaginatedResponseListV2`

NewDeploymentHistoryServicePaginatedResponseListV2 instantiates a new DeploymentHistoryServicePaginatedResponseListV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryServicePaginatedResponseListV2WithDefaults

`func NewDeploymentHistoryServicePaginatedResponseListV2WithDefaults() *DeploymentHistoryServicePaginatedResponseListV2`

NewDeploymentHistoryServicePaginatedResponseListV2WithDefaults instantiates a new DeploymentHistoryServicePaginatedResponseListV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *DeploymentHistoryServicePaginatedResponseListV2) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DeploymentHistoryServicePaginatedResponseListV2) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DeploymentHistoryServicePaginatedResponseListV2) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *DeploymentHistoryServicePaginatedResponseListV2) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DeploymentHistoryServicePaginatedResponseListV2) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DeploymentHistoryServicePaginatedResponseListV2) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *DeploymentHistoryServicePaginatedResponseListV2) GetResults() []DeploymentHistoryService`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DeploymentHistoryServicePaginatedResponseListV2) GetResultsOk() (*[]DeploymentHistoryService, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DeploymentHistoryServicePaginatedResponseListV2) SetResults(v []DeploymentHistoryService)`

SetResults sets Results field to given value.

### HasResults

`func (o *DeploymentHistoryServicePaginatedResponseListV2) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


