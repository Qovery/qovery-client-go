# SkillPresetRejectionFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**SkillPresetFindingSource**](SkillPresetFindingSource.md) |  | 
**RuleId** | **string** | For &#x60;source&#x60; RULES, the rule that matched, for example &#x60;override-instructions&#x60;, &#x60;conceal-from-user&#x60;, &#x60;skip-confirmation&#x60;, &#x60;read-credentials&#x60;, &#x60;dump-secrets&#x60;, &#x60;exfiltration-endpoint&#x60;, &#x60;upload-command&#x60;, &#x60;remote-exec&#x60;, &#x60;obfuscated-exec&#x60;, &#x60;external-url&#x60;, &#x60;invisible-characters&#x60;, &#x60;html-comment&#x60;, &#x60;encoded-blob&#x60; or &#x60;credential&#x60;. For &#x60;source&#x60; MODEL, the category: &#x60;prompt-injection&#x60;, &#x60;concealment&#x60;, &#x60;credential-access&#x60;, &#x60;data-exfiltration&#x60;, &#x60;unsafe-execution&#x60;, &#x60;approval-bypass&#x60;, &#x60;social-engineering&#x60; or &#x60;other&#x60;. New rules and categories can be added. | 
**Location** | **string** | The field that holds the finding: &#x60;description&#x60;, &#x60;team context&#x60;, &#x60;description of skill &#39;&lt;name&gt;&#39;&#x60; or &#x60;content of skill &#39;&lt;name&gt;&#39;&#x60;. | 
**Line** | **int32** | 1-based line within that field. | 
**Reason** | **string** | Why the content was refused. Never contains the offending text. | 

## Methods

### NewSkillPresetRejectionFinding

`func NewSkillPresetRejectionFinding(source SkillPresetFindingSource, ruleId string, location string, line int32, reason string, ) *SkillPresetRejectionFinding`

NewSkillPresetRejectionFinding instantiates a new SkillPresetRejectionFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetRejectionFindingWithDefaults

`func NewSkillPresetRejectionFindingWithDefaults() *SkillPresetRejectionFinding`

NewSkillPresetRejectionFindingWithDefaults instantiates a new SkillPresetRejectionFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *SkillPresetRejectionFinding) GetSource() SkillPresetFindingSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SkillPresetRejectionFinding) GetSourceOk() (*SkillPresetFindingSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SkillPresetRejectionFinding) SetSource(v SkillPresetFindingSource)`

SetSource sets Source field to given value.


### GetRuleId

`func (o *SkillPresetRejectionFinding) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *SkillPresetRejectionFinding) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *SkillPresetRejectionFinding) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetLocation

`func (o *SkillPresetRejectionFinding) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SkillPresetRejectionFinding) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SkillPresetRejectionFinding) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetLine

`func (o *SkillPresetRejectionFinding) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *SkillPresetRejectionFinding) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *SkillPresetRejectionFinding) SetLine(v int32)`

SetLine sets Line field to given value.


### GetReason

`func (o *SkillPresetRejectionFinding) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SkillPresetRejectionFinding) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SkillPresetRejectionFinding) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


