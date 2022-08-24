# DeployAllRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]DeployAllRequestApplicationsInner**](DeployAllRequestApplicationsInner.md) |  | [optional] 
**Containers** | Pointer to [**[]DeployAllRequestContainersInner**](DeployAllRequestContainersInner.md) |  | [optional] 

## Methods

### NewDeployAllRequest

`func NewDeployAllRequest() *DeployAllRequest`

NewDeployAllRequest instantiates a new DeployAllRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployAllRequestWithDefaults

`func NewDeployAllRequestWithDefaults() *DeployAllRequest`

NewDeployAllRequestWithDefaults instantiates a new DeployAllRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *DeployAllRequest) GetApplications() []DeployAllRequestApplicationsInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *DeployAllRequest) GetApplicationsOk() (*[]DeployAllRequestApplicationsInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *DeployAllRequest) SetApplications(v []DeployAllRequestApplicationsInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *DeployAllRequest) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetContainers

`func (o *DeployAllRequest) GetContainers() []DeployAllRequestContainersInner`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *DeployAllRequest) GetContainersOk() (*[]DeployAllRequestContainersInner, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *DeployAllRequest) SetContainers(v []DeployAllRequestContainersInner)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *DeployAllRequest) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


