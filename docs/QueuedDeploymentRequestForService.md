# QueuedDeploymentRequestForService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | [**QueuedDeploymentRequestForServiceIdentifier**](QueuedDeploymentRequestForServiceIdentifier.md) |  | 
**AuditingData** | [**QueuedDeploymentRequestForServiceAuditingData**](QueuedDeploymentRequestForServiceAuditingData.md) |  | 
**StatusDetails** | [**StatusDetails**](StatusDetails.md) |  | 
**IconUri** | **string** |  | 

## Methods

### NewQueuedDeploymentRequestForService

`func NewQueuedDeploymentRequestForService(identifier QueuedDeploymentRequestForServiceIdentifier, auditingData QueuedDeploymentRequestForServiceAuditingData, statusDetails StatusDetails, iconUri string, ) *QueuedDeploymentRequestForService`

NewQueuedDeploymentRequestForService instantiates a new QueuedDeploymentRequestForService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuedDeploymentRequestForServiceWithDefaults

`func NewQueuedDeploymentRequestForServiceWithDefaults() *QueuedDeploymentRequestForService`

NewQueuedDeploymentRequestForServiceWithDefaults instantiates a new QueuedDeploymentRequestForService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *QueuedDeploymentRequestForService) GetIdentifier() QueuedDeploymentRequestForServiceIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *QueuedDeploymentRequestForService) GetIdentifierOk() (*QueuedDeploymentRequestForServiceIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *QueuedDeploymentRequestForService) SetIdentifier(v QueuedDeploymentRequestForServiceIdentifier)`

SetIdentifier sets Identifier field to given value.


### GetAuditingData

`func (o *QueuedDeploymentRequestForService) GetAuditingData() QueuedDeploymentRequestForServiceAuditingData`

GetAuditingData returns the AuditingData field if non-nil, zero value otherwise.

### GetAuditingDataOk

`func (o *QueuedDeploymentRequestForService) GetAuditingDataOk() (*QueuedDeploymentRequestForServiceAuditingData, bool)`

GetAuditingDataOk returns a tuple with the AuditingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditingData

`func (o *QueuedDeploymentRequestForService) SetAuditingData(v QueuedDeploymentRequestForServiceAuditingData)`

SetAuditingData sets AuditingData field to given value.


### GetStatusDetails

`func (o *QueuedDeploymentRequestForService) GetStatusDetails() StatusDetails`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *QueuedDeploymentRequestForService) GetStatusDetailsOk() (*StatusDetails, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *QueuedDeploymentRequestForService) SetStatusDetails(v StatusDetails)`

SetStatusDetails sets StatusDetails field to given value.


### GetIconUri

`func (o *QueuedDeploymentRequestForService) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *QueuedDeploymentRequestForService) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *QueuedDeploymentRequestForService) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


