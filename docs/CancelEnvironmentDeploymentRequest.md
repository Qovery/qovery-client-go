# CancelEnvironmentDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceCancel** | Pointer to **bool** | Force cancel everything running in this environment if set to true (e.q lifecycle jobs triggered during the deployment). | [optional] [default to false]

## Methods

### NewCancelEnvironmentDeploymentRequest

`func NewCancelEnvironmentDeploymentRequest() *CancelEnvironmentDeploymentRequest`

NewCancelEnvironmentDeploymentRequest instantiates a new CancelEnvironmentDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelEnvironmentDeploymentRequestWithDefaults

`func NewCancelEnvironmentDeploymentRequestWithDefaults() *CancelEnvironmentDeploymentRequest`

NewCancelEnvironmentDeploymentRequestWithDefaults instantiates a new CancelEnvironmentDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceCancel

`func (o *CancelEnvironmentDeploymentRequest) GetForceCancel() bool`

GetForceCancel returns the ForceCancel field if non-nil, zero value otherwise.

### GetForceCancelOk

`func (o *CancelEnvironmentDeploymentRequest) GetForceCancelOk() (*bool, bool)`

GetForceCancelOk returns a tuple with the ForceCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceCancel

`func (o *CancelEnvironmentDeploymentRequest) SetForceCancel(v bool)`

SetForceCancel sets ForceCancel field to given value.

### HasForceCancel

`func (o *CancelEnvironmentDeploymentRequest) HasForceCancel() bool`

HasForceCancel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


