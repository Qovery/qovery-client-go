# GetEnvironmentStatusesWithStages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**Status**](Status.md) |  | [optional] 
**Stages** | Pointer to [**DeploymentStageWithServiceStatusesList**](DeploymentStageWithServiceStatusesList.md) |  | [optional] 

## Methods

### NewGetEnvironmentStatusesWithStages200Response

`func NewGetEnvironmentStatusesWithStages200Response() *GetEnvironmentStatusesWithStages200Response`

NewGetEnvironmentStatusesWithStages200Response instantiates a new GetEnvironmentStatusesWithStages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEnvironmentStatusesWithStages200ResponseWithDefaults

`func NewGetEnvironmentStatusesWithStages200ResponseWithDefaults() *GetEnvironmentStatusesWithStages200Response`

NewGetEnvironmentStatusesWithStages200ResponseWithDefaults instantiates a new GetEnvironmentStatusesWithStages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *GetEnvironmentStatusesWithStages200Response) GetEnvironment() Status`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GetEnvironmentStatusesWithStages200Response) GetEnvironmentOk() (*Status, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GetEnvironmentStatusesWithStages200Response) SetEnvironment(v Status)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GetEnvironmentStatusesWithStages200Response) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetStages

`func (o *GetEnvironmentStatusesWithStages200Response) GetStages() DeploymentStageWithServiceStatusesList`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *GetEnvironmentStatusesWithStages200Response) GetStagesOk() (*DeploymentStageWithServiceStatusesList, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *GetEnvironmentStatusesWithStages200Response) SetStages(v DeploymentStageWithServiceStatusesList)`

SetStages sets Stages field to given value.

### HasStages

`func (o *GetEnvironmentStatusesWithStages200Response) HasStages() bool`

HasStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


