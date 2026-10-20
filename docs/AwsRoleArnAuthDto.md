# AwsRoleArnAuthDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**RoleArn** | **string** |  | 
**AssumeRole** | Pointer to [**NullableAwsRoleArnAuthDtoAssumeRole**](AwsRoleArnAuthDtoAssumeRole.md) |  | [optional] 

## Methods

### NewAwsRoleArnAuthDto

`func NewAwsRoleArnAuthDto(mode string, roleArn string, ) *AwsRoleArnAuthDto`

NewAwsRoleArnAuthDto instantiates a new AwsRoleArnAuthDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRoleArnAuthDtoWithDefaults

`func NewAwsRoleArnAuthDtoWithDefaults() *AwsRoleArnAuthDto`

NewAwsRoleArnAuthDtoWithDefaults instantiates a new AwsRoleArnAuthDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *AwsRoleArnAuthDto) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AwsRoleArnAuthDto) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AwsRoleArnAuthDto) SetMode(v string)`

SetMode sets Mode field to given value.


### GetRoleArn

`func (o *AwsRoleArnAuthDto) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsRoleArnAuthDto) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsRoleArnAuthDto) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetAssumeRole

`func (o *AwsRoleArnAuthDto) GetAssumeRole() AwsRoleArnAuthDtoAssumeRole`

GetAssumeRole returns the AssumeRole field if non-nil, zero value otherwise.

### GetAssumeRoleOk

`func (o *AwsRoleArnAuthDto) GetAssumeRoleOk() (*AwsRoleArnAuthDtoAssumeRole, bool)`

GetAssumeRoleOk returns a tuple with the AssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumeRole

`func (o *AwsRoleArnAuthDto) SetAssumeRole(v AwsRoleArnAuthDtoAssumeRole)`

SetAssumeRole sets AssumeRole field to given value.

### HasAssumeRole

`func (o *AwsRoleArnAuthDto) HasAssumeRole() bool`

HasAssumeRole returns a boolean if a field has been set.

### SetAssumeRoleNil

`func (o *AwsRoleArnAuthDto) SetAssumeRoleNil(b bool)`

 SetAssumeRoleNil sets the value for AssumeRole to be an explicit nil

### UnsetAssumeRole
`func (o *AwsRoleArnAuthDto) UnsetAssumeRole()`

UnsetAssumeRole ensures that no value is present for AssumeRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


