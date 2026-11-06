# BuildAnalysisStartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | **string** | Execution id of the build to analyse. It must be the latest build of the service, &#x60;latest_build_execution_id&#x60; in &#x60;BuildAnalysisState&#x60;. | 

## Methods

### NewBuildAnalysisStartRequest

`func NewBuildAnalysisStartRequest(executionId string, ) *BuildAnalysisStartRequest`

NewBuildAnalysisStartRequest instantiates a new BuildAnalysisStartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildAnalysisStartRequestWithDefaults

`func NewBuildAnalysisStartRequestWithDefaults() *BuildAnalysisStartRequest`

NewBuildAnalysisStartRequestWithDefaults instantiates a new BuildAnalysisStartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *BuildAnalysisStartRequest) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BuildAnalysisStartRequest) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BuildAnalysisStartRequest) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


