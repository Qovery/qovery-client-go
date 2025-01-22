# EnvDeploymentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentRequestId** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StateEnum**](StateEnum.md) |  | [optional] 

## Methods

### NewEnvDeploymentStatus

`func NewEnvDeploymentStatus() *EnvDeploymentStatus`

NewEnvDeploymentStatus instantiates a new EnvDeploymentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvDeploymentStatusWithDefaults

`func NewEnvDeploymentStatusWithDefaults() *EnvDeploymentStatus`

NewEnvDeploymentStatusWithDefaults instantiates a new EnvDeploymentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentRequestId

`func (o *EnvDeploymentStatus) GetDeploymentRequestId() string`

GetDeploymentRequestId returns the DeploymentRequestId field if non-nil, zero value otherwise.

### GetDeploymentRequestIdOk

`func (o *EnvDeploymentStatus) GetDeploymentRequestIdOk() (*string, bool)`

GetDeploymentRequestIdOk returns a tuple with the DeploymentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRequestId

`func (o *EnvDeploymentStatus) SetDeploymentRequestId(v string)`

SetDeploymentRequestId sets DeploymentRequestId field to given value.

### HasDeploymentRequestId

`func (o *EnvDeploymentStatus) HasDeploymentRequestId() bool`

HasDeploymentRequestId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *EnvDeploymentStatus) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *EnvDeploymentStatus) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *EnvDeploymentStatus) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *EnvDeploymentStatus) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetStatus

`func (o *EnvDeploymentStatus) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvDeploymentStatus) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvDeploymentStatus) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnvDeploymentStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


