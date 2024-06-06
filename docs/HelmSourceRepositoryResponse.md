# HelmSourceRepositoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | **string** |  | 
**ChartVersion** | **string** |  | 
**Repository** | [**HelmSourceRepositoryResponseRepository**](HelmSourceRepositoryResponseRepository.md) |  | 

## Methods

### NewHelmSourceRepositoryResponse

`func NewHelmSourceRepositoryResponse(chartName string, chartVersion string, repository HelmSourceRepositoryResponseRepository, ) *HelmSourceRepositoryResponse`

NewHelmSourceRepositoryResponse instantiates a new HelmSourceRepositoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmSourceRepositoryResponseWithDefaults

`func NewHelmSourceRepositoryResponseWithDefaults() *HelmSourceRepositoryResponse`

NewHelmSourceRepositoryResponseWithDefaults instantiates a new HelmSourceRepositoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *HelmSourceRepositoryResponse) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *HelmSourceRepositoryResponse) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *HelmSourceRepositoryResponse) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartVersion

`func (o *HelmSourceRepositoryResponse) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *HelmSourceRepositoryResponse) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *HelmSourceRepositoryResponse) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetRepository

`func (o *HelmSourceRepositoryResponse) GetRepository() HelmSourceRepositoryResponseRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HelmSourceRepositoryResponse) GetRepositoryOk() (*HelmSourceRepositoryResponseRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HelmSourceRepositoryResponse) SetRepository(v HelmSourceRepositoryResponseRepository)`

SetRepository sets Repository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


