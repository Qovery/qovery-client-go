# AwsRoleArnAuthDtoAssumeRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | **string** | ARN of the IAM role to assume. | 
**ExternalId** | Pointer to **NullableString** | STS external id expected by the trust policy of the assumed role, if it requires one. | [optional] 

## Methods

### NewAwsRoleArnAuthDtoAssumeRole

`func NewAwsRoleArnAuthDtoAssumeRole(roleArn string, ) *AwsRoleArnAuthDtoAssumeRole`

NewAwsRoleArnAuthDtoAssumeRole instantiates a new AwsRoleArnAuthDtoAssumeRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRoleArnAuthDtoAssumeRoleWithDefaults

`func NewAwsRoleArnAuthDtoAssumeRoleWithDefaults() *AwsRoleArnAuthDtoAssumeRole`

NewAwsRoleArnAuthDtoAssumeRoleWithDefaults instantiates a new AwsRoleArnAuthDtoAssumeRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleArn

`func (o *AwsRoleArnAuthDtoAssumeRole) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsRoleArnAuthDtoAssumeRole) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsRoleArnAuthDtoAssumeRole) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetExternalId

`func (o *AwsRoleArnAuthDtoAssumeRole) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AwsRoleArnAuthDtoAssumeRole) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AwsRoleArnAuthDtoAssumeRole) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AwsRoleArnAuthDtoAssumeRole) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AwsRoleArnAuthDtoAssumeRole) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AwsRoleArnAuthDtoAssumeRole) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


