# ArgocdManagedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kubernetes resource kind (e.g. Deployment, Service) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**TargetState** | Pointer to **string** | JSON-encoded desired Kubernetes manifest | [optional] 
**LiveState** | Pointer to **string** | JSON-encoded live Kubernetes manifest | [optional] 

## Methods

### NewArgocdManagedResource

`func NewArgocdManagedResource() *ArgocdManagedResource`

NewArgocdManagedResource instantiates a new ArgocdManagedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgocdManagedResourceWithDefaults

`func NewArgocdManagedResourceWithDefaults() *ArgocdManagedResource`

NewArgocdManagedResourceWithDefaults instantiates a new ArgocdManagedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ArgocdManagedResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArgocdManagedResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArgocdManagedResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArgocdManagedResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ArgocdManagedResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArgocdManagedResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArgocdManagedResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArgocdManagedResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ArgocdManagedResource) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ArgocdManagedResource) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ArgocdManagedResource) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ArgocdManagedResource) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTargetState

`func (o *ArgocdManagedResource) GetTargetState() string`

GetTargetState returns the TargetState field if non-nil, zero value otherwise.

### GetTargetStateOk

`func (o *ArgocdManagedResource) GetTargetStateOk() (*string, bool)`

GetTargetStateOk returns a tuple with the TargetState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetState

`func (o *ArgocdManagedResource) SetTargetState(v string)`

SetTargetState sets TargetState field to given value.

### HasTargetState

`func (o *ArgocdManagedResource) HasTargetState() bool`

HasTargetState returns a boolean if a field has been set.

### GetLiveState

`func (o *ArgocdManagedResource) GetLiveState() string`

GetLiveState returns the LiveState field if non-nil, zero value otherwise.

### GetLiveStateOk

`func (o *ArgocdManagedResource) GetLiveStateOk() (*string, bool)`

GetLiveStateOk returns a tuple with the LiveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveState

`func (o *ArgocdManagedResource) SetLiveState(v string)`

SetLiveState sets LiveState field to given value.

### HasLiveState

`func (o *ArgocdManagedResource) HasLiveState() bool`

HasLiveState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


