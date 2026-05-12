# ExternalSecretAssociatedServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**ProjectName** | **string** |  | 
**EnvironmentId** | **string** |  | 
**EnvironmentName** | **string** |  | 
**ServiceId** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to [**APIVariableScopeEnum**](APIVariableScopeEnum.md) |  | [optional] 
**VariableName** | **string** |  | 
**ExternalSecretName** | **string** |  | 

## Methods

### NewExternalSecretAssociatedServiceResponse

`func NewExternalSecretAssociatedServiceResponse(projectId string, projectName string, environmentId string, environmentName string, variableName string, externalSecretName string, ) *ExternalSecretAssociatedServiceResponse`

NewExternalSecretAssociatedServiceResponse instantiates a new ExternalSecretAssociatedServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalSecretAssociatedServiceResponseWithDefaults

`func NewExternalSecretAssociatedServiceResponseWithDefaults() *ExternalSecretAssociatedServiceResponse`

NewExternalSecretAssociatedServiceResponseWithDefaults instantiates a new ExternalSecretAssociatedServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ExternalSecretAssociatedServiceResponse) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ExternalSecretAssociatedServiceResponse) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ExternalSecretAssociatedServiceResponse) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProjectName

`func (o *ExternalSecretAssociatedServiceResponse) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ExternalSecretAssociatedServiceResponse) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ExternalSecretAssociatedServiceResponse) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetEnvironmentId

`func (o *ExternalSecretAssociatedServiceResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ExternalSecretAssociatedServiceResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ExternalSecretAssociatedServiceResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetEnvironmentName

`func (o *ExternalSecretAssociatedServiceResponse) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *ExternalSecretAssociatedServiceResponse) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *ExternalSecretAssociatedServiceResponse) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetServiceId

`func (o *ExternalSecretAssociatedServiceResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ExternalSecretAssociatedServiceResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ExternalSecretAssociatedServiceResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ExternalSecretAssociatedServiceResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *ExternalSecretAssociatedServiceResponse) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ExternalSecretAssociatedServiceResponse) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ExternalSecretAssociatedServiceResponse) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ExternalSecretAssociatedServiceResponse) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceType

`func (o *ExternalSecretAssociatedServiceResponse) GetServiceType() APIVariableScopeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ExternalSecretAssociatedServiceResponse) GetServiceTypeOk() (*APIVariableScopeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ExternalSecretAssociatedServiceResponse) SetServiceType(v APIVariableScopeEnum)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ExternalSecretAssociatedServiceResponse) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetVariableName

`func (o *ExternalSecretAssociatedServiceResponse) GetVariableName() string`

GetVariableName returns the VariableName field if non-nil, zero value otherwise.

### GetVariableNameOk

`func (o *ExternalSecretAssociatedServiceResponse) GetVariableNameOk() (*string, bool)`

GetVariableNameOk returns a tuple with the VariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableName

`func (o *ExternalSecretAssociatedServiceResponse) SetVariableName(v string)`

SetVariableName sets VariableName field to given value.


### GetExternalSecretName

`func (o *ExternalSecretAssociatedServiceResponse) GetExternalSecretName() string`

GetExternalSecretName returns the ExternalSecretName field if non-nil, zero value otherwise.

### GetExternalSecretNameOk

`func (o *ExternalSecretAssociatedServiceResponse) GetExternalSecretNameOk() (*string, bool)`

GetExternalSecretNameOk returns a tuple with the ExternalSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecretName

`func (o *ExternalSecretAssociatedServiceResponse) SetExternalSecretName(v string)`

SetExternalSecretName sets ExternalSecretName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


