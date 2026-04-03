# ClusterEnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** |  | 
**EnvironmentName** | **string** |  | 
**ProjectId** | **string** |  | 
**ProjectName** | **string** |  | 
**Services** | [**[]ClusterEnvironmentServiceResponse**](ClusterEnvironmentServiceResponse.md) |  | 

## Methods

### NewClusterEnvironmentResponse

`func NewClusterEnvironmentResponse(environmentId string, environmentName string, projectId string, projectName string, services []ClusterEnvironmentServiceResponse, ) *ClusterEnvironmentResponse`

NewClusterEnvironmentResponse instantiates a new ClusterEnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEnvironmentResponseWithDefaults

`func NewClusterEnvironmentResponseWithDefaults() *ClusterEnvironmentResponse`

NewClusterEnvironmentResponseWithDefaults instantiates a new ClusterEnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *ClusterEnvironmentResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ClusterEnvironmentResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ClusterEnvironmentResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetEnvironmentName

`func (o *ClusterEnvironmentResponse) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *ClusterEnvironmentResponse) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *ClusterEnvironmentResponse) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetProjectId

`func (o *ClusterEnvironmentResponse) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ClusterEnvironmentResponse) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ClusterEnvironmentResponse) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProjectName

`func (o *ClusterEnvironmentResponse) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ClusterEnvironmentResponse) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ClusterEnvironmentResponse) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetServices

`func (o *ClusterEnvironmentResponse) GetServices() []ClusterEnvironmentServiceResponse`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ClusterEnvironmentResponse) GetServicesOk() (*[]ClusterEnvironmentServiceResponse, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ClusterEnvironmentResponse) SetServices(v []ClusterEnvironmentServiceResponse)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


