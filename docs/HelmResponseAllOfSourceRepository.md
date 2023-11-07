# HelmResponseAllOfSourceRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | Pointer to **string** | The name of the chart in the repository | [optional] 
**ChartVersion** | Pointer to **string** | The version of the chart to use | [optional] 
**Repository** | Pointer to [**HelmResponseAllOfSourceRepositoryRepository**](HelmResponseAllOfSourceRepositoryRepository.md) |  | [optional] 

## Methods

### NewHelmResponseAllOfSourceRepository

`func NewHelmResponseAllOfSourceRepository() *HelmResponseAllOfSourceRepository`

NewHelmResponseAllOfSourceRepository instantiates a new HelmResponseAllOfSourceRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmResponseAllOfSourceRepositoryWithDefaults

`func NewHelmResponseAllOfSourceRepositoryWithDefaults() *HelmResponseAllOfSourceRepository`

NewHelmResponseAllOfSourceRepositoryWithDefaults instantiates a new HelmResponseAllOfSourceRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *HelmResponseAllOfSourceRepository) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *HelmResponseAllOfSourceRepository) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *HelmResponseAllOfSourceRepository) SetChartName(v string)`

SetChartName sets ChartName field to given value.

### HasChartName

`func (o *HelmResponseAllOfSourceRepository) HasChartName() bool`

HasChartName returns a boolean if a field has been set.

### GetChartVersion

`func (o *HelmResponseAllOfSourceRepository) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *HelmResponseAllOfSourceRepository) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *HelmResponseAllOfSourceRepository) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.

### HasChartVersion

`func (o *HelmResponseAllOfSourceRepository) HasChartVersion() bool`

HasChartVersion returns a boolean if a field has been set.

### GetRepository

`func (o *HelmResponseAllOfSourceRepository) GetRepository() HelmResponseAllOfSourceRepositoryRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HelmResponseAllOfSourceRepository) GetRepositoryOk() (*HelmResponseAllOfSourceRepositoryRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HelmResponseAllOfSourceRepository) SetRepository(v HelmResponseAllOfSourceRepositoryRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *HelmResponseAllOfSourceRepository) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


