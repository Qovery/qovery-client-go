# ClusterInfrastructureChartsParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NginxParameters** | Pointer to [**ClusterInfrastructureNginxChartParameters**](ClusterInfrastructureNginxChartParameters.md) |  | [optional] 
**CertManagerParameters** | Pointer to [**ClusterInfrastructureCertManagerChartParameters**](ClusterInfrastructureCertManagerChartParameters.md) |  | [optional] 
**MetalLbParameters** | Pointer to [**ClusterInfrastructureMetalLbChartParameters**](ClusterInfrastructureMetalLbChartParameters.md) |  | [optional] 

## Methods

### NewClusterInfrastructureChartsParameters

`func NewClusterInfrastructureChartsParameters() *ClusterInfrastructureChartsParameters`

NewClusterInfrastructureChartsParameters instantiates a new ClusterInfrastructureChartsParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInfrastructureChartsParametersWithDefaults

`func NewClusterInfrastructureChartsParametersWithDefaults() *ClusterInfrastructureChartsParameters`

NewClusterInfrastructureChartsParametersWithDefaults instantiates a new ClusterInfrastructureChartsParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNginxParameters

`func (o *ClusterInfrastructureChartsParameters) GetNginxParameters() ClusterInfrastructureNginxChartParameters`

GetNginxParameters returns the NginxParameters field if non-nil, zero value otherwise.

### GetNginxParametersOk

`func (o *ClusterInfrastructureChartsParameters) GetNginxParametersOk() (*ClusterInfrastructureNginxChartParameters, bool)`

GetNginxParametersOk returns a tuple with the NginxParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNginxParameters

`func (o *ClusterInfrastructureChartsParameters) SetNginxParameters(v ClusterInfrastructureNginxChartParameters)`

SetNginxParameters sets NginxParameters field to given value.

### HasNginxParameters

`func (o *ClusterInfrastructureChartsParameters) HasNginxParameters() bool`

HasNginxParameters returns a boolean if a field has been set.

### GetCertManagerParameters

`func (o *ClusterInfrastructureChartsParameters) GetCertManagerParameters() ClusterInfrastructureCertManagerChartParameters`

GetCertManagerParameters returns the CertManagerParameters field if non-nil, zero value otherwise.

### GetCertManagerParametersOk

`func (o *ClusterInfrastructureChartsParameters) GetCertManagerParametersOk() (*ClusterInfrastructureCertManagerChartParameters, bool)`

GetCertManagerParametersOk returns a tuple with the CertManagerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertManagerParameters

`func (o *ClusterInfrastructureChartsParameters) SetCertManagerParameters(v ClusterInfrastructureCertManagerChartParameters)`

SetCertManagerParameters sets CertManagerParameters field to given value.

### HasCertManagerParameters

`func (o *ClusterInfrastructureChartsParameters) HasCertManagerParameters() bool`

HasCertManagerParameters returns a boolean if a field has been set.

### GetMetalLbParameters

`func (o *ClusterInfrastructureChartsParameters) GetMetalLbParameters() ClusterInfrastructureMetalLbChartParameters`

GetMetalLbParameters returns the MetalLbParameters field if non-nil, zero value otherwise.

### GetMetalLbParametersOk

`func (o *ClusterInfrastructureChartsParameters) GetMetalLbParametersOk() (*ClusterInfrastructureMetalLbChartParameters, bool)`

GetMetalLbParametersOk returns a tuple with the MetalLbParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalLbParameters

`func (o *ClusterInfrastructureChartsParameters) SetMetalLbParameters(v ClusterInfrastructureMetalLbChartParameters)`

SetMetalLbParameters sets MetalLbParameters field to given value.

### HasMetalLbParameters

`func (o *ClusterInfrastructureChartsParameters) HasMetalLbParameters() bool`

HasMetalLbParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


