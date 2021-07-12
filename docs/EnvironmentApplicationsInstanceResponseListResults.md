# EnvironmentApplicationsInstanceResponseListResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **string** |  | 
**Instances** | [**[]InstanceResponse**](InstanceResponse.md) |  | 

## Methods

### NewEnvironmentApplicationsInstanceResponseListResults

`func NewEnvironmentApplicationsInstanceResponseListResults(application string, instances []InstanceResponse, ) *EnvironmentApplicationsInstanceResponseListResults`

NewEnvironmentApplicationsInstanceResponseListResults instantiates a new EnvironmentApplicationsInstanceResponseListResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentApplicationsInstanceResponseListResultsWithDefaults

`func NewEnvironmentApplicationsInstanceResponseListResultsWithDefaults() *EnvironmentApplicationsInstanceResponseListResults`

NewEnvironmentApplicationsInstanceResponseListResultsWithDefaults instantiates a new EnvironmentApplicationsInstanceResponseListResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *EnvironmentApplicationsInstanceResponseListResults) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *EnvironmentApplicationsInstanceResponseListResults) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *EnvironmentApplicationsInstanceResponseListResults) SetApplication(v string)`

SetApplication sets Application field to given value.


### GetInstances

`func (o *EnvironmentApplicationsInstanceResponseListResults) GetInstances() []InstanceResponse`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *EnvironmentApplicationsInstanceResponseListResults) GetInstancesOk() (*[]InstanceResponse, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *EnvironmentApplicationsInstanceResponseListResults) SetInstances(v []InstanceResponse)`

SetInstances sets Instances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


