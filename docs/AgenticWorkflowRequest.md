# AgenticWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**IpAllowlist** | Pointer to **[]string** | CIDR ranges the incoming webhook request&#39;s source IP is checked against | [optional] [default to []]
**ModelSettings** | Pointer to **string** |  | [optional] [default to ""]
**DockerFragment** | Pointer to **string** |  | [optional] [default to ""]
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**McpConnectors** | Pointer to [**[]AgenticWorkflowConnector**](AgenticWorkflowConnector.md) |  | [optional] [default to []]
**Outputs** | Pointer to [**[]AgenticWorkflowOutput**](AgenticWorkflowOutput.md) |  | [optional] [default to []]
**Model** | Pointer to [**AgenticWorkflowModel**](AgenticWorkflowModel.md) |  | [optional] [default to AGENTICWORKFLOWMODEL_CLAUDE]
**ProjectRepositories** | Pointer to [**[]AgenticWorkflowProjectRepository**](AgenticWorkflowProjectRepository.md) |  | [optional] [default to []]

## Methods

### NewAgenticWorkflowRequest

`func NewAgenticWorkflowRequest(name string, ) *AgenticWorkflowRequest`

NewAgenticWorkflowRequest instantiates a new AgenticWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowRequestWithDefaults

`func NewAgenticWorkflowRequestWithDefaults() *AgenticWorkflowRequest`

NewAgenticWorkflowRequestWithDefaults instantiates a new AgenticWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgenticWorkflowRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgenticWorkflowRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgenticWorkflowRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AgenticWorkflowRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgenticWorkflowRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgenticWorkflowRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgenticWorkflowRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAllowlist

`func (o *AgenticWorkflowRequest) GetIpAllowlist() []string`

GetIpAllowlist returns the IpAllowlist field if non-nil, zero value otherwise.

### GetIpAllowlistOk

`func (o *AgenticWorkflowRequest) GetIpAllowlistOk() (*[]string, bool)`

GetIpAllowlistOk returns a tuple with the IpAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowlist

`func (o *AgenticWorkflowRequest) SetIpAllowlist(v []string)`

SetIpAllowlist sets IpAllowlist field to given value.

### HasIpAllowlist

`func (o *AgenticWorkflowRequest) HasIpAllowlist() bool`

HasIpAllowlist returns a boolean if a field has been set.

### GetModelSettings

`func (o *AgenticWorkflowRequest) GetModelSettings() string`

GetModelSettings returns the ModelSettings field if non-nil, zero value otherwise.

### GetModelSettingsOk

`func (o *AgenticWorkflowRequest) GetModelSettingsOk() (*string, bool)`

GetModelSettingsOk returns a tuple with the ModelSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelSettings

`func (o *AgenticWorkflowRequest) SetModelSettings(v string)`

SetModelSettings sets ModelSettings field to given value.

### HasModelSettings

`func (o *AgenticWorkflowRequest) HasModelSettings() bool`

HasModelSettings returns a boolean if a field has been set.

### GetDockerFragment

`func (o *AgenticWorkflowRequest) GetDockerFragment() string`

GetDockerFragment returns the DockerFragment field if non-nil, zero value otherwise.

### GetDockerFragmentOk

`func (o *AgenticWorkflowRequest) GetDockerFragmentOk() (*string, bool)`

GetDockerFragmentOk returns a tuple with the DockerFragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerFragment

`func (o *AgenticWorkflowRequest) SetDockerFragment(v string)`

SetDockerFragment sets DockerFragment field to given value.

### HasDockerFragment

`func (o *AgenticWorkflowRequest) HasDockerFragment() bool`

HasDockerFragment returns a boolean if a field has been set.

### GetEnabled

`func (o *AgenticWorkflowRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgenticWorkflowRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgenticWorkflowRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgenticWorkflowRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMcpConnectors

`func (o *AgenticWorkflowRequest) GetMcpConnectors() []AgenticWorkflowConnector`

GetMcpConnectors returns the McpConnectors field if non-nil, zero value otherwise.

### GetMcpConnectorsOk

`func (o *AgenticWorkflowRequest) GetMcpConnectorsOk() (*[]AgenticWorkflowConnector, bool)`

GetMcpConnectorsOk returns a tuple with the McpConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpConnectors

`func (o *AgenticWorkflowRequest) SetMcpConnectors(v []AgenticWorkflowConnector)`

SetMcpConnectors sets McpConnectors field to given value.

### HasMcpConnectors

`func (o *AgenticWorkflowRequest) HasMcpConnectors() bool`

HasMcpConnectors returns a boolean if a field has been set.

### GetOutputs

`func (o *AgenticWorkflowRequest) GetOutputs() []AgenticWorkflowOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *AgenticWorkflowRequest) GetOutputsOk() (*[]AgenticWorkflowOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *AgenticWorkflowRequest) SetOutputs(v []AgenticWorkflowOutput)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *AgenticWorkflowRequest) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetModel

`func (o *AgenticWorkflowRequest) GetModel() AgenticWorkflowModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgenticWorkflowRequest) GetModelOk() (*AgenticWorkflowModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgenticWorkflowRequest) SetModel(v AgenticWorkflowModel)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgenticWorkflowRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetProjectRepositories

`func (o *AgenticWorkflowRequest) GetProjectRepositories() []AgenticWorkflowProjectRepository`

GetProjectRepositories returns the ProjectRepositories field if non-nil, zero value otherwise.

### GetProjectRepositoriesOk

`func (o *AgenticWorkflowRequest) GetProjectRepositoriesOk() (*[]AgenticWorkflowProjectRepository, bool)`

GetProjectRepositoriesOk returns a tuple with the ProjectRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRepositories

`func (o *AgenticWorkflowRequest) SetProjectRepositories(v []AgenticWorkflowProjectRepository)`

SetProjectRepositories sets ProjectRepositories field to given value.

### HasProjectRepositories

`func (o *AgenticWorkflowRequest) HasProjectRepositories() bool`

HasProjectRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


