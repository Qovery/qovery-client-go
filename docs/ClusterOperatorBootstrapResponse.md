# ClusterOperatorBootstrapResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartRepository** | **string** |  | 
**ChartName** | **string** |  | 
**ChartVersion** | **string** |  | 
**ChartReference** | **string** |  | 
**Namespace** | **string** |  | 
**ReleaseName** | **string** |  | 
**Values** | **map[string]interface{}** | Structured Helm values. This object can contain cluster credentials. | 
**ValuesYaml** | **string** | Ready-to-write Helm values file. This value can contain cluster credentials. | 
**HelmCommand** | **string** | Ready-to-run Helm upgrade/install command using values_yaml. | 

## Methods

### NewClusterOperatorBootstrapResponse

`func NewClusterOperatorBootstrapResponse(chartRepository string, chartName string, chartVersion string, chartReference string, namespace string, releaseName string, values map[string]interface{}, valuesYaml string, helmCommand string, ) *ClusterOperatorBootstrapResponse`

NewClusterOperatorBootstrapResponse instantiates a new ClusterOperatorBootstrapResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOperatorBootstrapResponseWithDefaults

`func NewClusterOperatorBootstrapResponseWithDefaults() *ClusterOperatorBootstrapResponse`

NewClusterOperatorBootstrapResponseWithDefaults instantiates a new ClusterOperatorBootstrapResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartRepository

`func (o *ClusterOperatorBootstrapResponse) GetChartRepository() string`

GetChartRepository returns the ChartRepository field if non-nil, zero value otherwise.

### GetChartRepositoryOk

`func (o *ClusterOperatorBootstrapResponse) GetChartRepositoryOk() (*string, bool)`

GetChartRepositoryOk returns a tuple with the ChartRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartRepository

`func (o *ClusterOperatorBootstrapResponse) SetChartRepository(v string)`

SetChartRepository sets ChartRepository field to given value.


### GetChartName

`func (o *ClusterOperatorBootstrapResponse) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *ClusterOperatorBootstrapResponse) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *ClusterOperatorBootstrapResponse) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartVersion

`func (o *ClusterOperatorBootstrapResponse) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *ClusterOperatorBootstrapResponse) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *ClusterOperatorBootstrapResponse) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetChartReference

`func (o *ClusterOperatorBootstrapResponse) GetChartReference() string`

GetChartReference returns the ChartReference field if non-nil, zero value otherwise.

### GetChartReferenceOk

`func (o *ClusterOperatorBootstrapResponse) GetChartReferenceOk() (*string, bool)`

GetChartReferenceOk returns a tuple with the ChartReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartReference

`func (o *ClusterOperatorBootstrapResponse) SetChartReference(v string)`

SetChartReference sets ChartReference field to given value.


### GetNamespace

`func (o *ClusterOperatorBootstrapResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ClusterOperatorBootstrapResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ClusterOperatorBootstrapResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetReleaseName

`func (o *ClusterOperatorBootstrapResponse) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *ClusterOperatorBootstrapResponse) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *ClusterOperatorBootstrapResponse) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.


### GetValues

`func (o *ClusterOperatorBootstrapResponse) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ClusterOperatorBootstrapResponse) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ClusterOperatorBootstrapResponse) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.


### GetValuesYaml

`func (o *ClusterOperatorBootstrapResponse) GetValuesYaml() string`

GetValuesYaml returns the ValuesYaml field if non-nil, zero value otherwise.

### GetValuesYamlOk

`func (o *ClusterOperatorBootstrapResponse) GetValuesYamlOk() (*string, bool)`

GetValuesYamlOk returns a tuple with the ValuesYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesYaml

`func (o *ClusterOperatorBootstrapResponse) SetValuesYaml(v string)`

SetValuesYaml sets ValuesYaml field to given value.


### GetHelmCommand

`func (o *ClusterOperatorBootstrapResponse) GetHelmCommand() string`

GetHelmCommand returns the HelmCommand field if non-nil, zero value otherwise.

### GetHelmCommandOk

`func (o *ClusterOperatorBootstrapResponse) GetHelmCommandOk() (*string, bool)`

GetHelmCommandOk returns a tuple with the HelmCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmCommand

`func (o *ClusterOperatorBootstrapResponse) SetHelmCommand(v string)`

SetHelmCommand sets HelmCommand field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


