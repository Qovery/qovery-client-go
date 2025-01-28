# QueuedDeploymentRequestWithStagesStagesServicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | [**QueuedDeploymentRequestWithStagesStagesServicesInnerIdentifier**](QueuedDeploymentRequestWithStagesStagesServicesInnerIdentifier.md) |  | 
**Status** | [**StageStatusEnum**](StageStatusEnum.md) |  | 
**IconUri** | Pointer to **string** |  | [optional] 

## Methods

### NewQueuedDeploymentRequestWithStagesStagesServicesInner

`func NewQueuedDeploymentRequestWithStagesStagesServicesInner(identifier QueuedDeploymentRequestWithStagesStagesServicesInnerIdentifier, status StageStatusEnum, ) *QueuedDeploymentRequestWithStagesStagesServicesInner`

NewQueuedDeploymentRequestWithStagesStagesServicesInner instantiates a new QueuedDeploymentRequestWithStagesStagesServicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuedDeploymentRequestWithStagesStagesServicesInnerWithDefaults

`func NewQueuedDeploymentRequestWithStagesStagesServicesInnerWithDefaults() *QueuedDeploymentRequestWithStagesStagesServicesInner`

NewQueuedDeploymentRequestWithStagesStagesServicesInnerWithDefaults instantiates a new QueuedDeploymentRequestWithStagesStagesServicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *QueuedDeploymentRequestWithStagesStagesServicesInner) GetIdentifier() QueuedDeploymentRequestWithStagesStagesServicesInnerIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *QueuedDeploymentRequestWithStagesStagesServicesInner) GetIdentifierOk() (*QueuedDeploymentRequestWithStagesStagesServicesInnerIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *QueuedDeploymentRequestWithStagesStagesServicesInner) SetIdentifier(v QueuedDeploymentRequestWithStagesStagesServicesInnerIdentifier)`

SetIdentifier sets Identifier field to given value.


### GetStatus

`func (o *QueuedDeploymentRequestWithStagesStagesServicesInner) GetStatus() StageStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueuedDeploymentRequestWithStagesStagesServicesInner) GetStatusOk() (*StageStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueuedDeploymentRequestWithStagesStagesServicesInner) SetStatus(v StageStatusEnum)`

SetStatus sets Status field to given value.


### GetIconUri

`func (o *QueuedDeploymentRequestWithStagesStagesServicesInner) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *QueuedDeploymentRequestWithStagesStagesServicesInner) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *QueuedDeploymentRequestWithStagesStagesServicesInner) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.

### HasIconUri

`func (o *QueuedDeploymentRequestWithStagesStagesServicesInner) HasIconUri() bool`

HasIconUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


