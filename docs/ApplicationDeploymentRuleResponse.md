# ApplicationDeploymentRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentRestrictions** | Pointer to [**[]ApplicationDeploymentRestriction**](ApplicationDeploymentRestriction.md) |  | [optional] 
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewApplicationDeploymentRuleResponse

`func NewApplicationDeploymentRuleResponse(id string, createdAt time.Time, ) *ApplicationDeploymentRuleResponse`

NewApplicationDeploymentRuleResponse instantiates a new ApplicationDeploymentRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDeploymentRuleResponseWithDefaults

`func NewApplicationDeploymentRuleResponseWithDefaults() *ApplicationDeploymentRuleResponse`

NewApplicationDeploymentRuleResponseWithDefaults instantiates a new ApplicationDeploymentRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentRestrictions

`func (o *ApplicationDeploymentRuleResponse) GetDeploymentRestrictions() []ApplicationDeploymentRestriction`

GetDeploymentRestrictions returns the DeploymentRestrictions field if non-nil, zero value otherwise.

### GetDeploymentRestrictionsOk

`func (o *ApplicationDeploymentRuleResponse) GetDeploymentRestrictionsOk() (*[]ApplicationDeploymentRestriction, bool)`

GetDeploymentRestrictionsOk returns a tuple with the DeploymentRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRestrictions

`func (o *ApplicationDeploymentRuleResponse) SetDeploymentRestrictions(v []ApplicationDeploymentRestriction)`

SetDeploymentRestrictions sets DeploymentRestrictions field to given value.

### HasDeploymentRestrictions

`func (o *ApplicationDeploymentRuleResponse) HasDeploymentRestrictions() bool`

HasDeploymentRestrictions returns a boolean if a field has been set.

### GetId

`func (o *ApplicationDeploymentRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationDeploymentRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationDeploymentRuleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ApplicationDeploymentRuleResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationDeploymentRuleResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationDeploymentRuleResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ApplicationDeploymentRuleResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApplicationDeploymentRuleResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApplicationDeploymentRuleResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApplicationDeploymentRuleResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


