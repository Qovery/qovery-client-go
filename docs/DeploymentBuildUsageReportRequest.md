# DeploymentBuildUsageReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | **string** | The deployment execution ID (format environment_id-version) | 
**ReportExpirationInSeconds** | **int32** | The number of seconds the report will be publicly available | 

## Methods

### NewDeploymentBuildUsageReportRequest

`func NewDeploymentBuildUsageReportRequest(executionId string, reportExpirationInSeconds int32, ) *DeploymentBuildUsageReportRequest`

NewDeploymentBuildUsageReportRequest instantiates a new DeploymentBuildUsageReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentBuildUsageReportRequestWithDefaults

`func NewDeploymentBuildUsageReportRequestWithDefaults() *DeploymentBuildUsageReportRequest`

NewDeploymentBuildUsageReportRequestWithDefaults instantiates a new DeploymentBuildUsageReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *DeploymentBuildUsageReportRequest) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *DeploymentBuildUsageReportRequest) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *DeploymentBuildUsageReportRequest) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetReportExpirationInSeconds

`func (o *DeploymentBuildUsageReportRequest) GetReportExpirationInSeconds() int32`

GetReportExpirationInSeconds returns the ReportExpirationInSeconds field if non-nil, zero value otherwise.

### GetReportExpirationInSecondsOk

`func (o *DeploymentBuildUsageReportRequest) GetReportExpirationInSecondsOk() (*int32, bool)`

GetReportExpirationInSecondsOk returns a tuple with the ReportExpirationInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportExpirationInSeconds

`func (o *DeploymentBuildUsageReportRequest) SetReportExpirationInSeconds(v int32)`

SetReportExpirationInSeconds sets ReportExpirationInSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


