# ClusterOperatorUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | **string** | Identifier shared by the queued Engine execution and its infrastructure logs. | 
**ImageVersion** | **string** | Operator image tag resolved for this execution. | 
**ChartVersion** | **string** | Operator Helm chart version resolved for this execution. | 

## Methods

### NewClusterOperatorUpdateResponse

`func NewClusterOperatorUpdateResponse(executionId string, imageVersion string, chartVersion string, ) *ClusterOperatorUpdateResponse`

NewClusterOperatorUpdateResponse instantiates a new ClusterOperatorUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOperatorUpdateResponseWithDefaults

`func NewClusterOperatorUpdateResponseWithDefaults() *ClusterOperatorUpdateResponse`

NewClusterOperatorUpdateResponseWithDefaults instantiates a new ClusterOperatorUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *ClusterOperatorUpdateResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ClusterOperatorUpdateResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ClusterOperatorUpdateResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetImageVersion

`func (o *ClusterOperatorUpdateResponse) GetImageVersion() string`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *ClusterOperatorUpdateResponse) GetImageVersionOk() (*string, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *ClusterOperatorUpdateResponse) SetImageVersion(v string)`

SetImageVersion sets ImageVersion field to given value.


### GetChartVersion

`func (o *ClusterOperatorUpdateResponse) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *ClusterOperatorUpdateResponse) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *ClusterOperatorUpdateResponse) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


