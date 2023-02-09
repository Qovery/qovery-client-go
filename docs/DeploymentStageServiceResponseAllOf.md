# DeploymentStageServiceResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** | id of the service attached to the stage | [optional] 
**ServiceType** | Pointer to **string** | type of the service (i.e APPLICATION, JOB, DATABASE, ...) | [optional] 

## Methods

### NewDeploymentStageServiceResponseAllOf

`func NewDeploymentStageServiceResponseAllOf() *DeploymentStageServiceResponseAllOf`

NewDeploymentStageServiceResponseAllOf instantiates a new DeploymentStageServiceResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStageServiceResponseAllOfWithDefaults

`func NewDeploymentStageServiceResponseAllOfWithDefaults() *DeploymentStageServiceResponseAllOf`

NewDeploymentStageServiceResponseAllOfWithDefaults instantiates a new DeploymentStageServiceResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *DeploymentStageServiceResponseAllOf) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeploymentStageServiceResponseAllOf) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeploymentStageServiceResponseAllOf) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DeploymentStageServiceResponseAllOf) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceType

`func (o *DeploymentStageServiceResponseAllOf) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *DeploymentStageServiceResponseAllOf) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *DeploymentStageServiceResponseAllOf) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *DeploymentStageServiceResponseAllOf) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


