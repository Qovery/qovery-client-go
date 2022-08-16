# EnvironmentApplicationsInstanceResponseListResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **string** |  | 
**Instances** | [**[]Instance**](Instance.md) |  | 

## Methods

### NewEnvironmentApplicationsInstanceResponseListResultsInner

`func NewEnvironmentApplicationsInstanceResponseListResultsInner(application string, instances []Instance, ) *EnvironmentApplicationsInstanceResponseListResultsInner`

NewEnvironmentApplicationsInstanceResponseListResultsInner instantiates a new EnvironmentApplicationsInstanceResponseListResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentApplicationsInstanceResponseListResultsInnerWithDefaults

`func NewEnvironmentApplicationsInstanceResponseListResultsInnerWithDefaults() *EnvironmentApplicationsInstanceResponseListResultsInner`

NewEnvironmentApplicationsInstanceResponseListResultsInnerWithDefaults instantiates a new EnvironmentApplicationsInstanceResponseListResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *EnvironmentApplicationsInstanceResponseListResultsInner) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *EnvironmentApplicationsInstanceResponseListResultsInner) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *EnvironmentApplicationsInstanceResponseListResultsInner) SetApplication(v string)`

SetApplication sets Application field to given value.


### GetInstances

`func (o *EnvironmentApplicationsInstanceResponseListResultsInner) GetInstances() []Instance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *EnvironmentApplicationsInstanceResponseListResultsInner) GetInstancesOk() (*[]Instance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *EnvironmentApplicationsInstanceResponseListResultsInner) SetInstances(v []Instance)`

SetInstances sets Instances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


