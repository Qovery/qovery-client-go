# ClusterOperatorUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartVersion** | **string** | Exact Qovery Operator Helm chart version to install. | 
**ImageVersion** | Pointer to **NullableString** | Optional Operator image tag overriding the image target selected by q-core. | [optional] 

## Methods

### NewClusterOperatorUpdateRequest

`func NewClusterOperatorUpdateRequest(chartVersion string, ) *ClusterOperatorUpdateRequest`

NewClusterOperatorUpdateRequest instantiates a new ClusterOperatorUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOperatorUpdateRequestWithDefaults

`func NewClusterOperatorUpdateRequestWithDefaults() *ClusterOperatorUpdateRequest`

NewClusterOperatorUpdateRequestWithDefaults instantiates a new ClusterOperatorUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartVersion

`func (o *ClusterOperatorUpdateRequest) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *ClusterOperatorUpdateRequest) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *ClusterOperatorUpdateRequest) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetImageVersion

`func (o *ClusterOperatorUpdateRequest) GetImageVersion() string`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *ClusterOperatorUpdateRequest) GetImageVersionOk() (*string, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *ClusterOperatorUpdateRequest) SetImageVersion(v string)`

SetImageVersion sets ImageVersion field to given value.

### HasImageVersion

`func (o *ClusterOperatorUpdateRequest) HasImageVersion() bool`

HasImageVersion returns a boolean if a field has been set.

### SetImageVersionNil

`func (o *ClusterOperatorUpdateRequest) SetImageVersionNil(b bool)`

 SetImageVersionNil sets the value for ImageVersion to be an explicit nil

### UnsetImageVersion
`func (o *ClusterOperatorUpdateRequest) UnsetImageVersion()`

UnsetImageVersion ensures that no value is present for ImageVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


