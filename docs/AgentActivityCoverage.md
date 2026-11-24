# AgentActivityCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | [**[]AgentActivitySource**](AgentActivitySource.md) | The sources the view is built from. | 
**Note** | **string** | The coverage limit, in plain words, to show next to the data. | 

## Methods

### NewAgentActivityCoverage

`func NewAgentActivityCoverage(sources []AgentActivitySource, note string, ) *AgentActivityCoverage`

NewAgentActivityCoverage instantiates a new AgentActivityCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActivityCoverageWithDefaults

`func NewAgentActivityCoverageWithDefaults() *AgentActivityCoverage`

NewAgentActivityCoverageWithDefaults instantiates a new AgentActivityCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *AgentActivityCoverage) GetSources() []AgentActivitySource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AgentActivityCoverage) GetSourcesOk() (*[]AgentActivitySource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AgentActivityCoverage) SetSources(v []AgentActivitySource)`

SetSources sets Sources field to given value.


### GetNote

`func (o *AgentActivityCoverage) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AgentActivityCoverage) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AgentActivityCoverage) SetNote(v string)`

SetNote sets Note field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


