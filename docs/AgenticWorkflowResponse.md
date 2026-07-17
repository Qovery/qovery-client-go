# AgenticWorkflowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** | name is case insensitive | 
**Description** | **string** |  | 
**IpAllowlist** | **[]string** | CIDR ranges the incoming webhook request&#39;s source IP is checked against | 
**ModelSettings** | **string** |  | 
**DockerFragment** | **string** |  | 
**Enabled** | **bool** |  | 
**McpConnectors** | [**[]AgenticWorkflowConnector**](AgenticWorkflowConnector.md) |  | 
**Outputs** | [**[]AgenticWorkflowOutput**](AgenticWorkflowOutput.md) |  | 
**Model** | [**AgenticWorkflowModel**](AgenticWorkflowModel.md) |  | 
**ProjectRepositories** | [**[]AgenticWorkflowProjectRepository**](AgenticWorkflowProjectRepository.md) |  | 
**Webhook** | [**AgenticWorkflowWebhook**](AgenticWorkflowWebhook.md) |  | 

## Methods

### NewAgenticWorkflowResponse

`func NewAgenticWorkflowResponse(id string, createdAt time.Time, name string, description string, ipAllowlist []string, modelSettings string, dockerFragment string, enabled bool, mcpConnectors []AgenticWorkflowConnector, outputs []AgenticWorkflowOutput, model AgenticWorkflowModel, projectRepositories []AgenticWorkflowProjectRepository, webhook AgenticWorkflowWebhook, ) *AgenticWorkflowResponse`

NewAgenticWorkflowResponse instantiates a new AgenticWorkflowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowResponseWithDefaults

`func NewAgenticWorkflowResponseWithDefaults() *AgenticWorkflowResponse`

NewAgenticWorkflowResponseWithDefaults instantiates a new AgenticWorkflowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgenticWorkflowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgenticWorkflowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgenticWorkflowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *AgenticWorkflowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgenticWorkflowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgenticWorkflowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AgenticWorkflowResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgenticWorkflowResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgenticWorkflowResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgenticWorkflowResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *AgenticWorkflowResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgenticWorkflowResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgenticWorkflowResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AgenticWorkflowResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgenticWorkflowResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgenticWorkflowResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIpAllowlist

`func (o *AgenticWorkflowResponse) GetIpAllowlist() []string`

GetIpAllowlist returns the IpAllowlist field if non-nil, zero value otherwise.

### GetIpAllowlistOk

`func (o *AgenticWorkflowResponse) GetIpAllowlistOk() (*[]string, bool)`

GetIpAllowlistOk returns a tuple with the IpAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowlist

`func (o *AgenticWorkflowResponse) SetIpAllowlist(v []string)`

SetIpAllowlist sets IpAllowlist field to given value.


### GetModelSettings

`func (o *AgenticWorkflowResponse) GetModelSettings() string`

GetModelSettings returns the ModelSettings field if non-nil, zero value otherwise.

### GetModelSettingsOk

`func (o *AgenticWorkflowResponse) GetModelSettingsOk() (*string, bool)`

GetModelSettingsOk returns a tuple with the ModelSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelSettings

`func (o *AgenticWorkflowResponse) SetModelSettings(v string)`

SetModelSettings sets ModelSettings field to given value.


### GetDockerFragment

`func (o *AgenticWorkflowResponse) GetDockerFragment() string`

GetDockerFragment returns the DockerFragment field if non-nil, zero value otherwise.

### GetDockerFragmentOk

`func (o *AgenticWorkflowResponse) GetDockerFragmentOk() (*string, bool)`

GetDockerFragmentOk returns a tuple with the DockerFragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerFragment

`func (o *AgenticWorkflowResponse) SetDockerFragment(v string)`

SetDockerFragment sets DockerFragment field to given value.


### GetEnabled

`func (o *AgenticWorkflowResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgenticWorkflowResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgenticWorkflowResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMcpConnectors

`func (o *AgenticWorkflowResponse) GetMcpConnectors() []AgenticWorkflowConnector`

GetMcpConnectors returns the McpConnectors field if non-nil, zero value otherwise.

### GetMcpConnectorsOk

`func (o *AgenticWorkflowResponse) GetMcpConnectorsOk() (*[]AgenticWorkflowConnector, bool)`

GetMcpConnectorsOk returns a tuple with the McpConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpConnectors

`func (o *AgenticWorkflowResponse) SetMcpConnectors(v []AgenticWorkflowConnector)`

SetMcpConnectors sets McpConnectors field to given value.


### GetOutputs

`func (o *AgenticWorkflowResponse) GetOutputs() []AgenticWorkflowOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *AgenticWorkflowResponse) GetOutputsOk() (*[]AgenticWorkflowOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *AgenticWorkflowResponse) SetOutputs(v []AgenticWorkflowOutput)`

SetOutputs sets Outputs field to given value.


### GetModel

`func (o *AgenticWorkflowResponse) GetModel() AgenticWorkflowModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgenticWorkflowResponse) GetModelOk() (*AgenticWorkflowModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgenticWorkflowResponse) SetModel(v AgenticWorkflowModel)`

SetModel sets Model field to given value.


### GetProjectRepositories

`func (o *AgenticWorkflowResponse) GetProjectRepositories() []AgenticWorkflowProjectRepository`

GetProjectRepositories returns the ProjectRepositories field if non-nil, zero value otherwise.

### GetProjectRepositoriesOk

`func (o *AgenticWorkflowResponse) GetProjectRepositoriesOk() (*[]AgenticWorkflowProjectRepository, bool)`

GetProjectRepositoriesOk returns a tuple with the ProjectRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRepositories

`func (o *AgenticWorkflowResponse) SetProjectRepositories(v []AgenticWorkflowProjectRepository)`

SetProjectRepositories sets ProjectRepositories field to given value.


### GetWebhook

`func (o *AgenticWorkflowResponse) GetWebhook() AgenticWorkflowWebhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *AgenticWorkflowResponse) GetWebhookOk() (*AgenticWorkflowWebhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *AgenticWorkflowResponse) SetWebhook(v AgenticWorkflowWebhook)`

SetWebhook sets Webhook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


