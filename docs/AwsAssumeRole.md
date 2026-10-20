# AwsAssumeRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | **string** | ARN of the IAM role to assume. | 
**ExternalId** | Pointer to **NullableString** | STS external id expected by the trust policy of the assumed role, if it requires one. | [optional] 

## Methods

### NewAwsAssumeRole

`func NewAwsAssumeRole(roleArn string, ) *AwsAssumeRole`

NewAwsAssumeRole instantiates a new AwsAssumeRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAssumeRoleWithDefaults

`func NewAwsAssumeRoleWithDefaults() *AwsAssumeRole`

NewAwsAssumeRoleWithDefaults instantiates a new AwsAssumeRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleArn

`func (o *AwsAssumeRole) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsAssumeRole) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsAssumeRole) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetExternalId

`func (o *AwsAssumeRole) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AwsAssumeRole) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AwsAssumeRole) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AwsAssumeRole) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AwsAssumeRole) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AwsAssumeRole) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


