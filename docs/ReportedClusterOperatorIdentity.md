# ReportedClusterOperatorIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseDigest** | **string** | Multi-architecture OCI index digest selected for the installed release. | 
**RuntimeImageId** | **string** | OCI image reference resolved by the Kubernetes runtime. | 
**ChartVersion** | **string** |  | 
**ObservedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewReportedClusterOperatorIdentity

`func NewReportedClusterOperatorIdentity(releaseDigest string, runtimeImageId string, chartVersion string, ) *ReportedClusterOperatorIdentity`

NewReportedClusterOperatorIdentity instantiates a new ReportedClusterOperatorIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportedClusterOperatorIdentityWithDefaults

`func NewReportedClusterOperatorIdentityWithDefaults() *ReportedClusterOperatorIdentity`

NewReportedClusterOperatorIdentityWithDefaults instantiates a new ReportedClusterOperatorIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseDigest

`func (o *ReportedClusterOperatorIdentity) GetReleaseDigest() string`

GetReleaseDigest returns the ReleaseDigest field if non-nil, zero value otherwise.

### GetReleaseDigestOk

`func (o *ReportedClusterOperatorIdentity) GetReleaseDigestOk() (*string, bool)`

GetReleaseDigestOk returns a tuple with the ReleaseDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDigest

`func (o *ReportedClusterOperatorIdentity) SetReleaseDigest(v string)`

SetReleaseDigest sets ReleaseDigest field to given value.


### GetRuntimeImageId

`func (o *ReportedClusterOperatorIdentity) GetRuntimeImageId() string`

GetRuntimeImageId returns the RuntimeImageId field if non-nil, zero value otherwise.

### GetRuntimeImageIdOk

`func (o *ReportedClusterOperatorIdentity) GetRuntimeImageIdOk() (*string, bool)`

GetRuntimeImageIdOk returns a tuple with the RuntimeImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeImageId

`func (o *ReportedClusterOperatorIdentity) SetRuntimeImageId(v string)`

SetRuntimeImageId sets RuntimeImageId field to given value.


### GetChartVersion

`func (o *ReportedClusterOperatorIdentity) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *ReportedClusterOperatorIdentity) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *ReportedClusterOperatorIdentity) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetObservedAt

`func (o *ReportedClusterOperatorIdentity) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *ReportedClusterOperatorIdentity) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *ReportedClusterOperatorIdentity) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *ReportedClusterOperatorIdentity) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.

### SetObservedAtNil

`func (o *ReportedClusterOperatorIdentity) SetObservedAtNil(b bool)`

 SetObservedAtNil sets the value for ObservedAt to be an explicit nil

### UnsetObservedAt
`func (o *ReportedClusterOperatorIdentity) UnsetObservedAt()`

UnsetObservedAt ensures that no value is present for ObservedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


