# ServiceLightResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**IconUri** | **string** |  | 
**ServiceType** | [**ServiceTypeEnum**](ServiceTypeEnum.md) |  | 
**ProjectId** | **string** |  | 
**ProjectName** | **string** |  | 
**EnvironmentId** | **string** |  | 
**EnvironmentName** | **string** |  | 
**ClusterId** | **string** |  | 
**JobType** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceLightResponse

`func NewServiceLightResponse(id string, name string, description string, iconUri string, serviceType ServiceTypeEnum, projectId string, projectName string, environmentId string, environmentName string, clusterId string, ) *ServiceLightResponse`

NewServiceLightResponse instantiates a new ServiceLightResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLightResponseWithDefaults

`func NewServiceLightResponseWithDefaults() *ServiceLightResponse`

NewServiceLightResponseWithDefaults instantiates a new ServiceLightResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceLightResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceLightResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceLightResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ServiceLightResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceLightResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceLightResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServiceLightResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceLightResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceLightResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIconUri

`func (o *ServiceLightResponse) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *ServiceLightResponse) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *ServiceLightResponse) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.


### GetServiceType

`func (o *ServiceLightResponse) GetServiceType() ServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ServiceLightResponse) GetServiceTypeOk() (*ServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ServiceLightResponse) SetServiceType(v ServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.


### GetProjectId

`func (o *ServiceLightResponse) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ServiceLightResponse) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ServiceLightResponse) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProjectName

`func (o *ServiceLightResponse) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ServiceLightResponse) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ServiceLightResponse) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetEnvironmentId

`func (o *ServiceLightResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ServiceLightResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ServiceLightResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetEnvironmentName

`func (o *ServiceLightResponse) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *ServiceLightResponse) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *ServiceLightResponse) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetClusterId

`func (o *ServiceLightResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ServiceLightResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ServiceLightResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetJobType

`func (o *ServiceLightResponse) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ServiceLightResponse) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ServiceLightResponse) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *ServiceLightResponse) HasJobType() bool`

HasJobType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


