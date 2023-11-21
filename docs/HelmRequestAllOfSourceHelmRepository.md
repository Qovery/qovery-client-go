# HelmRequestAllOfSourceHelmRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **NullableString** | The id of the helm repository | [optional] 
**ChartName** | Pointer to **string** | The name of the chart in the repository | [optional] 
**ChartVersion** | Pointer to **string** | The version of the chart to use | [optional] 

## Methods

### NewHelmRequestAllOfSourceHelmRepository

`func NewHelmRequestAllOfSourceHelmRepository() *HelmRequestAllOfSourceHelmRepository`

NewHelmRequestAllOfSourceHelmRepository instantiates a new HelmRequestAllOfSourceHelmRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRequestAllOfSourceHelmRepositoryWithDefaults

`func NewHelmRequestAllOfSourceHelmRepositoryWithDefaults() *HelmRequestAllOfSourceHelmRepository`

NewHelmRequestAllOfSourceHelmRepositoryWithDefaults instantiates a new HelmRequestAllOfSourceHelmRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *HelmRequestAllOfSourceHelmRepository) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HelmRequestAllOfSourceHelmRepository) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HelmRequestAllOfSourceHelmRepository) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *HelmRequestAllOfSourceHelmRepository) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *HelmRequestAllOfSourceHelmRepository) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *HelmRequestAllOfSourceHelmRepository) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetChartName

`func (o *HelmRequestAllOfSourceHelmRepository) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *HelmRequestAllOfSourceHelmRepository) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *HelmRequestAllOfSourceHelmRepository) SetChartName(v string)`

SetChartName sets ChartName field to given value.

### HasChartName

`func (o *HelmRequestAllOfSourceHelmRepository) HasChartName() bool`

HasChartName returns a boolean if a field has been set.

### GetChartVersion

`func (o *HelmRequestAllOfSourceHelmRepository) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *HelmRequestAllOfSourceHelmRepository) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *HelmRequestAllOfSourceHelmRepository) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.

### HasChartVersion

`func (o *HelmRequestAllOfSourceHelmRepository) HasChartVersion() bool`

HasChartVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


