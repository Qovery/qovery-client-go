# DeploymentHistoryStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Status** | [**StageStatusEnum**](StageStatusEnum.md) |  | 
**Duration** | **string** |  | 
**Services** | [**[]DeploymentHistoryService**](DeploymentHistoryService.md) |  | 

## Methods

### NewDeploymentHistoryStage

`func NewDeploymentHistoryStage(name string, status StageStatusEnum, duration string, services []DeploymentHistoryService, ) *DeploymentHistoryStage`

NewDeploymentHistoryStage instantiates a new DeploymentHistoryStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryStageWithDefaults

`func NewDeploymentHistoryStageWithDefaults() *DeploymentHistoryStage`

NewDeploymentHistoryStageWithDefaults instantiates a new DeploymentHistoryStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentHistoryStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentHistoryStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentHistoryStage) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DeploymentHistoryStage) GetStatus() StageStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryStage) GetStatusOk() (*StageStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryStage) SetStatus(v StageStatusEnum)`

SetStatus sets Status field to given value.


### GetDuration

`func (o *DeploymentHistoryStage) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeploymentHistoryStage) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeploymentHistoryStage) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetServices

`func (o *DeploymentHistoryStage) GetServices() []DeploymentHistoryService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *DeploymentHistoryStage) GetServicesOk() (*[]DeploymentHistoryService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *DeploymentHistoryStage) SetServices(v []DeploymentHistoryService)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


