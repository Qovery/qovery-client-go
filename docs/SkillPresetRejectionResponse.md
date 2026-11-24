# SkillPresetRejectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Status** | **int32** |  | 
**Detail** | **string** | Every finding in one sentence, safe to log. | 
**Instance** | Pointer to **string** | The request path. | [optional] 
**Findings** | [**[]SkillPresetRejectionFinding**](SkillPresetRejectionFinding.md) |  | 

## Methods

### NewSkillPresetRejectionResponse

`func NewSkillPresetRejectionResponse(title string, status int32, detail string, findings []SkillPresetRejectionFinding, ) *SkillPresetRejectionResponse`

NewSkillPresetRejectionResponse instantiates a new SkillPresetRejectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetRejectionResponseWithDefaults

`func NewSkillPresetRejectionResponseWithDefaults() *SkillPresetRejectionResponse`

NewSkillPresetRejectionResponseWithDefaults instantiates a new SkillPresetRejectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkillPresetRejectionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkillPresetRejectionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkillPresetRejectionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SkillPresetRejectionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *SkillPresetRejectionResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SkillPresetRejectionResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SkillPresetRejectionResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStatus

`func (o *SkillPresetRejectionResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SkillPresetRejectionResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SkillPresetRejectionResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetDetail

`func (o *SkillPresetRejectionResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SkillPresetRejectionResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SkillPresetRejectionResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetInstance

`func (o *SkillPresetRejectionResponse) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *SkillPresetRejectionResponse) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *SkillPresetRejectionResponse) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *SkillPresetRejectionResponse) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetFindings

`func (o *SkillPresetRejectionResponse) GetFindings() []SkillPresetRejectionFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *SkillPresetRejectionResponse) GetFindingsOk() (*[]SkillPresetRejectionFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *SkillPresetRejectionResponse) SetFindings(v []SkillPresetRejectionFinding)`

SetFindings sets Findings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


