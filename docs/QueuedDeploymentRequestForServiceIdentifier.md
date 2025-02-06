# QueuedDeploymentRequestForServiceIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentRequestId** | **string** |  | 
**ServiceId** | **string** |  | 
**ServiceType** | [**ServiceTypeEnum**](ServiceTypeEnum.md) |  | 
**Name** | **string** |  | 

## Methods

### NewQueuedDeploymentRequestForServiceIdentifier

`func NewQueuedDeploymentRequestForServiceIdentifier(deploymentRequestId string, serviceId string, serviceType ServiceTypeEnum, name string, ) *QueuedDeploymentRequestForServiceIdentifier`

NewQueuedDeploymentRequestForServiceIdentifier instantiates a new QueuedDeploymentRequestForServiceIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuedDeploymentRequestForServiceIdentifierWithDefaults

`func NewQueuedDeploymentRequestForServiceIdentifierWithDefaults() *QueuedDeploymentRequestForServiceIdentifier`

NewQueuedDeploymentRequestForServiceIdentifierWithDefaults instantiates a new QueuedDeploymentRequestForServiceIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentRequestId

`func (o *QueuedDeploymentRequestForServiceIdentifier) GetDeploymentRequestId() string`

GetDeploymentRequestId returns the DeploymentRequestId field if non-nil, zero value otherwise.

### GetDeploymentRequestIdOk

`func (o *QueuedDeploymentRequestForServiceIdentifier) GetDeploymentRequestIdOk() (*string, bool)`

GetDeploymentRequestIdOk returns a tuple with the DeploymentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentRequestId

`func (o *QueuedDeploymentRequestForServiceIdentifier) SetDeploymentRequestId(v string)`

SetDeploymentRequestId sets DeploymentRequestId field to given value.


### GetServiceId

`func (o *QueuedDeploymentRequestForServiceIdentifier) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *QueuedDeploymentRequestForServiceIdentifier) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *QueuedDeploymentRequestForServiceIdentifier) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceType

`func (o *QueuedDeploymentRequestForServiceIdentifier) GetServiceType() ServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *QueuedDeploymentRequestForServiceIdentifier) GetServiceTypeOk() (*ServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *QueuedDeploymentRequestForServiceIdentifier) SetServiceType(v ServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.


### GetName

`func (o *QueuedDeploymentRequestForServiceIdentifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueuedDeploymentRequestForServiceIdentifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueuedDeploymentRequestForServiceIdentifier) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


