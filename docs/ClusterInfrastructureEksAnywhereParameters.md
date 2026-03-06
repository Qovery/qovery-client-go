# ClusterInfrastructureEksAnywhereParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRepository** | [**ClusterEksAnywhereGitRepository**](ClusterEksAnywhereGitRepository.md) |  | 
**YamlFilePath** | **string** | Path to the EKS Anywhere cluster YAML file in the git repository | 

## Methods

### NewClusterInfrastructureEksAnywhereParameters

`func NewClusterInfrastructureEksAnywhereParameters(gitRepository ClusterEksAnywhereGitRepository, yamlFilePath string, ) *ClusterInfrastructureEksAnywhereParameters`

NewClusterInfrastructureEksAnywhereParameters instantiates a new ClusterInfrastructureEksAnywhereParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInfrastructureEksAnywhereParametersWithDefaults

`func NewClusterInfrastructureEksAnywhereParametersWithDefaults() *ClusterInfrastructureEksAnywhereParameters`

NewClusterInfrastructureEksAnywhereParametersWithDefaults instantiates a new ClusterInfrastructureEksAnywhereParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitRepository

`func (o *ClusterInfrastructureEksAnywhereParameters) GetGitRepository() ClusterEksAnywhereGitRepository`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *ClusterInfrastructureEksAnywhereParameters) GetGitRepositoryOk() (*ClusterEksAnywhereGitRepository, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *ClusterInfrastructureEksAnywhereParameters) SetGitRepository(v ClusterEksAnywhereGitRepository)`

SetGitRepository sets GitRepository field to given value.


### GetYamlFilePath

`func (o *ClusterInfrastructureEksAnywhereParameters) GetYamlFilePath() string`

GetYamlFilePath returns the YamlFilePath field if non-nil, zero value otherwise.

### GetYamlFilePathOk

`func (o *ClusterInfrastructureEksAnywhereParameters) GetYamlFilePathOk() (*string, bool)`

GetYamlFilePathOk returns a tuple with the YamlFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYamlFilePath

`func (o *ClusterInfrastructureEksAnywhereParameters) SetYamlFilePath(v string)`

SetYamlFilePath sets YamlFilePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


