# CloneEnvironmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**ClusterId** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | [optional] 
**ApplyDeploymentRule** | Pointer to **bool** |  | [optional] [default to false]
**ProjectId** | Pointer to **string** |  | [optional] 

## Methods

### NewCloneEnvironmentRequest

`func NewCloneEnvironmentRequest(name string, ) *CloneEnvironmentRequest`

NewCloneEnvironmentRequest instantiates a new CloneEnvironmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneEnvironmentRequestWithDefaults

`func NewCloneEnvironmentRequestWithDefaults() *CloneEnvironmentRequest`

NewCloneEnvironmentRequestWithDefaults instantiates a new CloneEnvironmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloneEnvironmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloneEnvironmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloneEnvironmentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetClusterId

`func (o *CloneEnvironmentRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CloneEnvironmentRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CloneEnvironmentRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *CloneEnvironmentRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetMode

`func (o *CloneEnvironmentRequest) GetMode() EnvironmentModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CloneEnvironmentRequest) GetModeOk() (*EnvironmentModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CloneEnvironmentRequest) SetMode(v EnvironmentModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CloneEnvironmentRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetApplyDeploymentRule

`func (o *CloneEnvironmentRequest) GetApplyDeploymentRule() bool`

GetApplyDeploymentRule returns the ApplyDeploymentRule field if non-nil, zero value otherwise.

### GetApplyDeploymentRuleOk

`func (o *CloneEnvironmentRequest) GetApplyDeploymentRuleOk() (*bool, bool)`

GetApplyDeploymentRuleOk returns a tuple with the ApplyDeploymentRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyDeploymentRule

`func (o *CloneEnvironmentRequest) SetApplyDeploymentRule(v bool)`

SetApplyDeploymentRule sets ApplyDeploymentRule field to given value.

### HasApplyDeploymentRule

`func (o *CloneEnvironmentRequest) HasApplyDeploymentRule() bool`

HasApplyDeploymentRule returns a boolean if a field has been set.

### GetProjectId

`func (o *CloneEnvironmentRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CloneEnvironmentRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CloneEnvironmentRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CloneEnvironmentRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


