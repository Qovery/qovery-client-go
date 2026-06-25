# ClusterAnalysisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**ClusterAnalysisKind**](ClusterAnalysisKind.md) |  | 
**OutputFormat** | [**ClusterAnalysisOutputFormat**](ClusterAnalysisOutputFormat.md) |  | 
**PrometheusUrl** | Pointer to **NullableString** | Optional Prometheus URL for COST_RECOMMENDATION analysis. When omitted, the engine resolves the default Qovery OBS endpoint. | [optional] 
**CmdArgs** | Pointer to **[]string** | Optional allowlisted KRR arguments for COST_RECOMMENDATION analysis. The engine validates and rejects unsupported or unsafe KRR flags. | [optional] 
**TargetKubernetesVersion** | Pointer to **NullableString** | Optional target Kubernetes version for DEPRECATED_API_CHECK analysis. | [optional] 

## Methods

### NewClusterAnalysisRequest

`func NewClusterAnalysisRequest(kind ClusterAnalysisKind, outputFormat ClusterAnalysisOutputFormat, ) *ClusterAnalysisRequest`

NewClusterAnalysisRequest instantiates a new ClusterAnalysisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAnalysisRequestWithDefaults

`func NewClusterAnalysisRequestWithDefaults() *ClusterAnalysisRequest`

NewClusterAnalysisRequestWithDefaults instantiates a new ClusterAnalysisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ClusterAnalysisRequest) GetKind() ClusterAnalysisKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterAnalysisRequest) GetKindOk() (*ClusterAnalysisKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterAnalysisRequest) SetKind(v ClusterAnalysisKind)`

SetKind sets Kind field to given value.


### GetOutputFormat

`func (o *ClusterAnalysisRequest) GetOutputFormat() ClusterAnalysisOutputFormat`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *ClusterAnalysisRequest) GetOutputFormatOk() (*ClusterAnalysisOutputFormat, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *ClusterAnalysisRequest) SetOutputFormat(v ClusterAnalysisOutputFormat)`

SetOutputFormat sets OutputFormat field to given value.


### GetPrometheusUrl

`func (o *ClusterAnalysisRequest) GetPrometheusUrl() string`

GetPrometheusUrl returns the PrometheusUrl field if non-nil, zero value otherwise.

### GetPrometheusUrlOk

`func (o *ClusterAnalysisRequest) GetPrometheusUrlOk() (*string, bool)`

GetPrometheusUrlOk returns a tuple with the PrometheusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusUrl

`func (o *ClusterAnalysisRequest) SetPrometheusUrl(v string)`

SetPrometheusUrl sets PrometheusUrl field to given value.

### HasPrometheusUrl

`func (o *ClusterAnalysisRequest) HasPrometheusUrl() bool`

HasPrometheusUrl returns a boolean if a field has been set.

### SetPrometheusUrlNil

`func (o *ClusterAnalysisRequest) SetPrometheusUrlNil(b bool)`

 SetPrometheusUrlNil sets the value for PrometheusUrl to be an explicit nil

### UnsetPrometheusUrl
`func (o *ClusterAnalysisRequest) UnsetPrometheusUrl()`

UnsetPrometheusUrl ensures that no value is present for PrometheusUrl, not even an explicit nil
### GetCmdArgs

`func (o *ClusterAnalysisRequest) GetCmdArgs() []string`

GetCmdArgs returns the CmdArgs field if non-nil, zero value otherwise.

### GetCmdArgsOk

`func (o *ClusterAnalysisRequest) GetCmdArgsOk() (*[]string, bool)`

GetCmdArgsOk returns a tuple with the CmdArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdArgs

`func (o *ClusterAnalysisRequest) SetCmdArgs(v []string)`

SetCmdArgs sets CmdArgs field to given value.

### HasCmdArgs

`func (o *ClusterAnalysisRequest) HasCmdArgs() bool`

HasCmdArgs returns a boolean if a field has been set.

### GetTargetKubernetesVersion

`func (o *ClusterAnalysisRequest) GetTargetKubernetesVersion() string`

GetTargetKubernetesVersion returns the TargetKubernetesVersion field if non-nil, zero value otherwise.

### GetTargetKubernetesVersionOk

`func (o *ClusterAnalysisRequest) GetTargetKubernetesVersionOk() (*string, bool)`

GetTargetKubernetesVersionOk returns a tuple with the TargetKubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKubernetesVersion

`func (o *ClusterAnalysisRequest) SetTargetKubernetesVersion(v string)`

SetTargetKubernetesVersion sets TargetKubernetesVersion field to given value.

### HasTargetKubernetesVersion

`func (o *ClusterAnalysisRequest) HasTargetKubernetesVersion() bool`

HasTargetKubernetesVersion returns a boolean if a field has been set.

### SetTargetKubernetesVersionNil

`func (o *ClusterAnalysisRequest) SetTargetKubernetesVersionNil(b bool)`

 SetTargetKubernetesVersionNil sets the value for TargetKubernetesVersion to be an explicit nil

### UnsetTargetKubernetesVersion
`func (o *ClusterAnalysisRequest) UnsetTargetKubernetesVersion()`

UnsetTargetKubernetesVersion ensures that no value is present for TargetKubernetesVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


