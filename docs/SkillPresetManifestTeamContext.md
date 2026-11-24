# SkillPresetManifestTeamContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproxTokens** | **int32** | Estimated size in tokens, the number of characters divided by 4. An estimate for the member, not a bill. | 
**Sha256** | **string** | Hex-encoded SHA-256 of the UTF-8 text. | 

## Methods

### NewSkillPresetManifestTeamContext

`func NewSkillPresetManifestTeamContext(approxTokens int32, sha256 string, ) *SkillPresetManifestTeamContext`

NewSkillPresetManifestTeamContext instantiates a new SkillPresetManifestTeamContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetManifestTeamContextWithDefaults

`func NewSkillPresetManifestTeamContextWithDefaults() *SkillPresetManifestTeamContext`

NewSkillPresetManifestTeamContextWithDefaults instantiates a new SkillPresetManifestTeamContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproxTokens

`func (o *SkillPresetManifestTeamContext) GetApproxTokens() int32`

GetApproxTokens returns the ApproxTokens field if non-nil, zero value otherwise.

### GetApproxTokensOk

`func (o *SkillPresetManifestTeamContext) GetApproxTokensOk() (*int32, bool)`

GetApproxTokensOk returns a tuple with the ApproxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproxTokens

`func (o *SkillPresetManifestTeamContext) SetApproxTokens(v int32)`

SetApproxTokens sets ApproxTokens field to given value.


### GetSha256

`func (o *SkillPresetManifestTeamContext) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *SkillPresetManifestTeamContext) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *SkillPresetManifestTeamContext) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


