# QueuedDeploymentRequestWithStagesStagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Status** | [**StageStatusEnum**](StageStatusEnum.md) |  | 
**Services** | [**[]QueuedDeploymentRequestWithStagesStagesInnerServicesInner**](QueuedDeploymentRequestWithStagesStagesInnerServicesInner.md) |  | 

## Methods

### NewQueuedDeploymentRequestWithStagesStagesInner

`func NewQueuedDeploymentRequestWithStagesStagesInner(name string, status StageStatusEnum, services []QueuedDeploymentRequestWithStagesStagesInnerServicesInner, ) *QueuedDeploymentRequestWithStagesStagesInner`

NewQueuedDeploymentRequestWithStagesStagesInner instantiates a new QueuedDeploymentRequestWithStagesStagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuedDeploymentRequestWithStagesStagesInnerWithDefaults

`func NewQueuedDeploymentRequestWithStagesStagesInnerWithDefaults() *QueuedDeploymentRequestWithStagesStagesInner`

NewQueuedDeploymentRequestWithStagesStagesInnerWithDefaults instantiates a new QueuedDeploymentRequestWithStagesStagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QueuedDeploymentRequestWithStagesStagesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueuedDeploymentRequestWithStagesStagesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueuedDeploymentRequestWithStagesStagesInner) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *QueuedDeploymentRequestWithStagesStagesInner) GetStatus() StageStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueuedDeploymentRequestWithStagesStagesInner) GetStatusOk() (*StageStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueuedDeploymentRequestWithStagesStagesInner) SetStatus(v StageStatusEnum)`

SetStatus sets Status field to given value.


### GetServices

`func (o *QueuedDeploymentRequestWithStagesStagesInner) GetServices() []QueuedDeploymentRequestWithStagesStagesInnerServicesInner`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *QueuedDeploymentRequestWithStagesStagesInner) GetServicesOk() (*[]QueuedDeploymentRequestWithStagesStagesInnerServicesInner, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *QueuedDeploymentRequestWithStagesStagesInner) SetServices(v []QueuedDeploymentRequestWithStagesStagesInnerServicesInner)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


