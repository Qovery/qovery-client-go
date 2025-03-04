# GetClusterKubernetesEvents200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The created date following ISO-8601 format | [optional] 
**Kind** | Pointer to **string** | The source kubernetes object related to the Event | [optional] 
**Namespace** | Pointer to **string** | The namespace of the kubernetes object related to the Event (optional) | [optional] 
**Name** | Pointer to **string** | The name of the Event | [optional] 
**Reason** | Pointer to **string** | The action that triggered the Event | [optional] 
**Message** | Pointer to **string** | A description of the Event | [optional] 
**Type** | Pointer to **string** | As of today it can be either Warning or Normal (can evolve in the next releases) | [optional] 

## Methods

### NewGetClusterKubernetesEvents200ResponseResultsInner

`func NewGetClusterKubernetesEvents200ResponseResultsInner() *GetClusterKubernetesEvents200ResponseResultsInner`

NewGetClusterKubernetesEvents200ResponseResultsInner instantiates a new GetClusterKubernetesEvents200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterKubernetesEvents200ResponseResultsInnerWithDefaults

`func NewGetClusterKubernetesEvents200ResponseResultsInnerWithDefaults() *GetClusterKubernetesEvents200ResponseResultsInner`

NewGetClusterKubernetesEvents200ResponseResultsInnerWithDefaults instantiates a new GetClusterKubernetesEvents200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetKind

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetNamespace

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetName

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReason

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetClusterKubernetesEvents200ResponseResultsInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


