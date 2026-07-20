# AgenticWorkflowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** | name is case insensitive | 
**Description** | **string** |  | 
**WebhookIpAllowlist** | **[]string** | CIDR ranges the incoming webhook request&#39;s source IP is checked against | 
**DockerFragment** | **string** |  | 
**Enabled** | **bool** |  | 
**Mcp** | **string** | Raw JSON blob describing the MCP servers configured for this workflow | 
**Outputs** | [**[]AgenticWorkflowOutput**](AgenticWorkflowOutput.md) |  | 
**Model** | [**AgenticWorkflowModelResponse**](AgenticWorkflowModelResponse.md) |  | 
**ProjectRepositories** | [**[]AgenticWorkflowProjectRepository**](AgenticWorkflowProjectRepository.md) |  | 
**AgentPrompt** | **string** |  | 
**Governance** | [**AgenticWorkflowGovernance**](AgenticWorkflowGovernance.md) |  | 
**Webhook** | [**AgenticWorkflowWebhook**](AgenticWorkflowWebhook.md) |  | 

## Methods

### NewAgenticWorkflowResponse

`func NewAgenticWorkflowResponse(id string, createdAt time.Time, name string, description string, webhookIpAllowlist []string, dockerFragment string, enabled bool, mcp string, outputs []AgenticWorkflowOutput, model AgenticWorkflowModelResponse, projectRepositories []AgenticWorkflowProjectRepository, agentPrompt string, governance AgenticWorkflowGovernance, webhook AgenticWorkflowWebhook, ) *AgenticWorkflowResponse`

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


### GetWebhookIpAllowlist

`func (o *AgenticWorkflowResponse) GetWebhookIpAllowlist() []string`

GetWebhookIpAllowlist returns the WebhookIpAllowlist field if non-nil, zero value otherwise.

### GetWebhookIpAllowlistOk

`func (o *AgenticWorkflowResponse) GetWebhookIpAllowlistOk() (*[]string, bool)`

GetWebhookIpAllowlistOk returns a tuple with the WebhookIpAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookIpAllowlist

`func (o *AgenticWorkflowResponse) SetWebhookIpAllowlist(v []string)`

SetWebhookIpAllowlist sets WebhookIpAllowlist field to given value.


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


### GetMcp

`func (o *AgenticWorkflowResponse) GetMcp() string`

GetMcp returns the Mcp field if non-nil, zero value otherwise.

### GetMcpOk

`func (o *AgenticWorkflowResponse) GetMcpOk() (*string, bool)`

GetMcpOk returns a tuple with the Mcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcp

`func (o *AgenticWorkflowResponse) SetMcp(v string)`

SetMcp sets Mcp field to given value.


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

`func (o *AgenticWorkflowResponse) GetModel() AgenticWorkflowModelResponse`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgenticWorkflowResponse) GetModelOk() (*AgenticWorkflowModelResponse, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgenticWorkflowResponse) SetModel(v AgenticWorkflowModelResponse)`

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


### GetAgentPrompt

`func (o *AgenticWorkflowResponse) GetAgentPrompt() string`

GetAgentPrompt returns the AgentPrompt field if non-nil, zero value otherwise.

### GetAgentPromptOk

`func (o *AgenticWorkflowResponse) GetAgentPromptOk() (*string, bool)`

GetAgentPromptOk returns a tuple with the AgentPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPrompt

`func (o *AgenticWorkflowResponse) SetAgentPrompt(v string)`

SetAgentPrompt sets AgentPrompt field to given value.


### GetGovernance

`func (o *AgenticWorkflowResponse) GetGovernance() AgenticWorkflowGovernance`

GetGovernance returns the Governance field if non-nil, zero value otherwise.

### GetGovernanceOk

`func (o *AgenticWorkflowResponse) GetGovernanceOk() (*AgenticWorkflowGovernance, bool)`

GetGovernanceOk returns a tuple with the Governance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernance

`func (o *AgenticWorkflowResponse) SetGovernance(v AgenticWorkflowGovernance)`

SetGovernance sets Governance field to given value.


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


