# ClusterOperatorFleetInventoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **string** |  | 
**ClusterId** | **string** |  | 
**ClusterName** | **string** |  | 
**ClusterKind** | [**SelfManagedClusterKind**](SelfManagedClusterKind.md) |  | 
**Attached** | **bool** | Whether the cluster is explicitly routed through the Operator execution path. | 
**Connected** | **bool** | Whether the last heartbeat is within the Operator presence window. | 
**LastHeartbeat** | Pointer to **NullableTime** |  | [optional] 
**DesiredImageVersion** | Pointer to **NullableString** |  | [optional] 
**ReportedImageVersion** | Pointer to **NullableString** |  | [optional] 
**DesiredChartVersion** | Pointer to **NullableString** |  | [optional] 
**ReportedChartVersion** | Pointer to **NullableString** |  | [optional] 
**Status** | [**ClusterOperatorFleetStatus**](ClusterOperatorFleetStatus.md) |  | 

## Methods

### NewClusterOperatorFleetInventoryResponse

`func NewClusterOperatorFleetInventoryResponse(organizationId string, clusterId string, clusterName string, clusterKind SelfManagedClusterKind, attached bool, connected bool, status ClusterOperatorFleetStatus, ) *ClusterOperatorFleetInventoryResponse`

NewClusterOperatorFleetInventoryResponse instantiates a new ClusterOperatorFleetInventoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOperatorFleetInventoryResponseWithDefaults

`func NewClusterOperatorFleetInventoryResponseWithDefaults() *ClusterOperatorFleetInventoryResponse`

NewClusterOperatorFleetInventoryResponseWithDefaults instantiates a new ClusterOperatorFleetInventoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *ClusterOperatorFleetInventoryResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ClusterOperatorFleetInventoryResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ClusterOperatorFleetInventoryResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetClusterId

`func (o *ClusterOperatorFleetInventoryResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterOperatorFleetInventoryResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterOperatorFleetInventoryResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *ClusterOperatorFleetInventoryResponse) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ClusterOperatorFleetInventoryResponse) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ClusterOperatorFleetInventoryResponse) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetClusterKind

`func (o *ClusterOperatorFleetInventoryResponse) GetClusterKind() SelfManagedClusterKind`

GetClusterKind returns the ClusterKind field if non-nil, zero value otherwise.

### GetClusterKindOk

`func (o *ClusterOperatorFleetInventoryResponse) GetClusterKindOk() (*SelfManagedClusterKind, bool)`

GetClusterKindOk returns a tuple with the ClusterKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterKind

`func (o *ClusterOperatorFleetInventoryResponse) SetClusterKind(v SelfManagedClusterKind)`

SetClusterKind sets ClusterKind field to given value.


### GetAttached

`func (o *ClusterOperatorFleetInventoryResponse) GetAttached() bool`

GetAttached returns the Attached field if non-nil, zero value otherwise.

### GetAttachedOk

`func (o *ClusterOperatorFleetInventoryResponse) GetAttachedOk() (*bool, bool)`

GetAttachedOk returns a tuple with the Attached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttached

`func (o *ClusterOperatorFleetInventoryResponse) SetAttached(v bool)`

SetAttached sets Attached field to given value.


### GetConnected

`func (o *ClusterOperatorFleetInventoryResponse) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ClusterOperatorFleetInventoryResponse) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ClusterOperatorFleetInventoryResponse) SetConnected(v bool)`

SetConnected sets Connected field to given value.


### GetLastHeartbeat

`func (o *ClusterOperatorFleetInventoryResponse) GetLastHeartbeat() time.Time`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *ClusterOperatorFleetInventoryResponse) GetLastHeartbeatOk() (*time.Time, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *ClusterOperatorFleetInventoryResponse) SetLastHeartbeat(v time.Time)`

SetLastHeartbeat sets LastHeartbeat field to given value.

### HasLastHeartbeat

`func (o *ClusterOperatorFleetInventoryResponse) HasLastHeartbeat() bool`

HasLastHeartbeat returns a boolean if a field has been set.

### SetLastHeartbeatNil

`func (o *ClusterOperatorFleetInventoryResponse) SetLastHeartbeatNil(b bool)`

 SetLastHeartbeatNil sets the value for LastHeartbeat to be an explicit nil

### UnsetLastHeartbeat
`func (o *ClusterOperatorFleetInventoryResponse) UnsetLastHeartbeat()`

UnsetLastHeartbeat ensures that no value is present for LastHeartbeat, not even an explicit nil
### GetDesiredImageVersion

`func (o *ClusterOperatorFleetInventoryResponse) GetDesiredImageVersion() string`

GetDesiredImageVersion returns the DesiredImageVersion field if non-nil, zero value otherwise.

### GetDesiredImageVersionOk

`func (o *ClusterOperatorFleetInventoryResponse) GetDesiredImageVersionOk() (*string, bool)`

GetDesiredImageVersionOk returns a tuple with the DesiredImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredImageVersion

`func (o *ClusterOperatorFleetInventoryResponse) SetDesiredImageVersion(v string)`

SetDesiredImageVersion sets DesiredImageVersion field to given value.

### HasDesiredImageVersion

`func (o *ClusterOperatorFleetInventoryResponse) HasDesiredImageVersion() bool`

HasDesiredImageVersion returns a boolean if a field has been set.

### SetDesiredImageVersionNil

`func (o *ClusterOperatorFleetInventoryResponse) SetDesiredImageVersionNil(b bool)`

 SetDesiredImageVersionNil sets the value for DesiredImageVersion to be an explicit nil

### UnsetDesiredImageVersion
`func (o *ClusterOperatorFleetInventoryResponse) UnsetDesiredImageVersion()`

UnsetDesiredImageVersion ensures that no value is present for DesiredImageVersion, not even an explicit nil
### GetReportedImageVersion

`func (o *ClusterOperatorFleetInventoryResponse) GetReportedImageVersion() string`

GetReportedImageVersion returns the ReportedImageVersion field if non-nil, zero value otherwise.

### GetReportedImageVersionOk

`func (o *ClusterOperatorFleetInventoryResponse) GetReportedImageVersionOk() (*string, bool)`

GetReportedImageVersionOk returns a tuple with the ReportedImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedImageVersion

`func (o *ClusterOperatorFleetInventoryResponse) SetReportedImageVersion(v string)`

SetReportedImageVersion sets ReportedImageVersion field to given value.

### HasReportedImageVersion

`func (o *ClusterOperatorFleetInventoryResponse) HasReportedImageVersion() bool`

HasReportedImageVersion returns a boolean if a field has been set.

### SetReportedImageVersionNil

`func (o *ClusterOperatorFleetInventoryResponse) SetReportedImageVersionNil(b bool)`

 SetReportedImageVersionNil sets the value for ReportedImageVersion to be an explicit nil

### UnsetReportedImageVersion
`func (o *ClusterOperatorFleetInventoryResponse) UnsetReportedImageVersion()`

UnsetReportedImageVersion ensures that no value is present for ReportedImageVersion, not even an explicit nil
### GetDesiredChartVersion

`func (o *ClusterOperatorFleetInventoryResponse) GetDesiredChartVersion() string`

GetDesiredChartVersion returns the DesiredChartVersion field if non-nil, zero value otherwise.

### GetDesiredChartVersionOk

`func (o *ClusterOperatorFleetInventoryResponse) GetDesiredChartVersionOk() (*string, bool)`

GetDesiredChartVersionOk returns a tuple with the DesiredChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredChartVersion

`func (o *ClusterOperatorFleetInventoryResponse) SetDesiredChartVersion(v string)`

SetDesiredChartVersion sets DesiredChartVersion field to given value.

### HasDesiredChartVersion

`func (o *ClusterOperatorFleetInventoryResponse) HasDesiredChartVersion() bool`

HasDesiredChartVersion returns a boolean if a field has been set.

### SetDesiredChartVersionNil

`func (o *ClusterOperatorFleetInventoryResponse) SetDesiredChartVersionNil(b bool)`

 SetDesiredChartVersionNil sets the value for DesiredChartVersion to be an explicit nil

### UnsetDesiredChartVersion
`func (o *ClusterOperatorFleetInventoryResponse) UnsetDesiredChartVersion()`

UnsetDesiredChartVersion ensures that no value is present for DesiredChartVersion, not even an explicit nil
### GetReportedChartVersion

`func (o *ClusterOperatorFleetInventoryResponse) GetReportedChartVersion() string`

GetReportedChartVersion returns the ReportedChartVersion field if non-nil, zero value otherwise.

### GetReportedChartVersionOk

`func (o *ClusterOperatorFleetInventoryResponse) GetReportedChartVersionOk() (*string, bool)`

GetReportedChartVersionOk returns a tuple with the ReportedChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedChartVersion

`func (o *ClusterOperatorFleetInventoryResponse) SetReportedChartVersion(v string)`

SetReportedChartVersion sets ReportedChartVersion field to given value.

### HasReportedChartVersion

`func (o *ClusterOperatorFleetInventoryResponse) HasReportedChartVersion() bool`

HasReportedChartVersion returns a boolean if a field has been set.

### SetReportedChartVersionNil

`func (o *ClusterOperatorFleetInventoryResponse) SetReportedChartVersionNil(b bool)`

 SetReportedChartVersionNil sets the value for ReportedChartVersion to be an explicit nil

### UnsetReportedChartVersion
`func (o *ClusterOperatorFleetInventoryResponse) UnsetReportedChartVersion()`

UnsetReportedChartVersion ensures that no value is present for ReportedChartVersion, not even an explicit nil
### GetStatus

`func (o *ClusterOperatorFleetInventoryResponse) GetStatus() ClusterOperatorFleetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterOperatorFleetInventoryResponse) GetStatusOk() (*ClusterOperatorFleetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterOperatorFleetInventoryResponse) SetStatus(v ClusterOperatorFleetStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


