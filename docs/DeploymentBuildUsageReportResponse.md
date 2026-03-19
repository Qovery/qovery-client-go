# DeploymentBuildUsageReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportUrl** | Pointer to **string** | The publicly accessible URL of the Grafana snapshot report showing build pod resource usage | [optional] 
**DeleteReportUrl** | Pointer to **string** | The URL to pro-actively delete the report before it expires | [optional] 

## Methods

### NewDeploymentBuildUsageReportResponse

`func NewDeploymentBuildUsageReportResponse() *DeploymentBuildUsageReportResponse`

NewDeploymentBuildUsageReportResponse instantiates a new DeploymentBuildUsageReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentBuildUsageReportResponseWithDefaults

`func NewDeploymentBuildUsageReportResponseWithDefaults() *DeploymentBuildUsageReportResponse`

NewDeploymentBuildUsageReportResponseWithDefaults instantiates a new DeploymentBuildUsageReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportUrl

`func (o *DeploymentBuildUsageReportResponse) GetReportUrl() string`

GetReportUrl returns the ReportUrl field if non-nil, zero value otherwise.

### GetReportUrlOk

`func (o *DeploymentBuildUsageReportResponse) GetReportUrlOk() (*string, bool)`

GetReportUrlOk returns a tuple with the ReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUrl

`func (o *DeploymentBuildUsageReportResponse) SetReportUrl(v string)`

SetReportUrl sets ReportUrl field to given value.

### HasReportUrl

`func (o *DeploymentBuildUsageReportResponse) HasReportUrl() bool`

HasReportUrl returns a boolean if a field has been set.

### GetDeleteReportUrl

`func (o *DeploymentBuildUsageReportResponse) GetDeleteReportUrl() string`

GetDeleteReportUrl returns the DeleteReportUrl field if non-nil, zero value otherwise.

### GetDeleteReportUrlOk

`func (o *DeploymentBuildUsageReportResponse) GetDeleteReportUrlOk() (*string, bool)`

GetDeleteReportUrlOk returns a tuple with the DeleteReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteReportUrl

`func (o *DeploymentBuildUsageReportResponse) SetDeleteReportUrl(v string)`

SetDeleteReportUrl sets DeleteReportUrl field to given value.

### HasDeleteReportUrl

`func (o *DeploymentBuildUsageReportResponse) HasDeleteReportUrl() bool`

HasDeleteReportUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


