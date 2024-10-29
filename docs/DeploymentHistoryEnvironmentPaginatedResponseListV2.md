# DeploymentHistoryEnvironmentPaginatedResponseListV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **float32** |  | 
**PageSize** | **float32** |  | 
**Results** | Pointer to [**[]DeploymentHistoryEnvironment**](DeploymentHistoryEnvironment.md) |  | [optional] 

## Methods

### NewDeploymentHistoryEnvironmentPaginatedResponseListV2

`func NewDeploymentHistoryEnvironmentPaginatedResponseListV2(page float32, pageSize float32, ) *DeploymentHistoryEnvironmentPaginatedResponseListV2`

NewDeploymentHistoryEnvironmentPaginatedResponseListV2 instantiates a new DeploymentHistoryEnvironmentPaginatedResponseListV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryEnvironmentPaginatedResponseListV2WithDefaults

`func NewDeploymentHistoryEnvironmentPaginatedResponseListV2WithDefaults() *DeploymentHistoryEnvironmentPaginatedResponseListV2`

NewDeploymentHistoryEnvironmentPaginatedResponseListV2WithDefaults instantiates a new DeploymentHistoryEnvironmentPaginatedResponseListV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) GetResults() []DeploymentHistoryEnvironment`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) GetResultsOk() (*[]DeploymentHistoryEnvironment, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) SetResults(v []DeploymentHistoryEnvironment)`

SetResults sets Results field to given value.

### HasResults

`func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


