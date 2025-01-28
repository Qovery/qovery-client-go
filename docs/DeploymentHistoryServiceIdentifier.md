# DeploymentHistoryServiceIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ServiceId** | **string** |  | 
**ServiceType** | [**ServiceTypeEnum**](ServiceTypeEnum.md) |  | 
**ExecutionId** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentHistoryServiceIdentifier

`func NewDeploymentHistoryServiceIdentifier(name string, serviceId string, serviceType ServiceTypeEnum, ) *DeploymentHistoryServiceIdentifier`

NewDeploymentHistoryServiceIdentifier instantiates a new DeploymentHistoryServiceIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryServiceIdentifierWithDefaults

`func NewDeploymentHistoryServiceIdentifierWithDefaults() *DeploymentHistoryServiceIdentifier`

NewDeploymentHistoryServiceIdentifierWithDefaults instantiates a new DeploymentHistoryServiceIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentHistoryServiceIdentifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentHistoryServiceIdentifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentHistoryServiceIdentifier) SetName(v string)`

SetName sets Name field to given value.


### GetServiceId

`func (o *DeploymentHistoryServiceIdentifier) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeploymentHistoryServiceIdentifier) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeploymentHistoryServiceIdentifier) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceType

`func (o *DeploymentHistoryServiceIdentifier) GetServiceType() ServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *DeploymentHistoryServiceIdentifier) GetServiceTypeOk() (*ServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *DeploymentHistoryServiceIdentifier) SetServiceType(v ServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.


### GetExecutionId

`func (o *DeploymentHistoryServiceIdentifier) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *DeploymentHistoryServiceIdentifier) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *DeploymentHistoryServiceIdentifier) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *DeploymentHistoryServiceIdentifier) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


