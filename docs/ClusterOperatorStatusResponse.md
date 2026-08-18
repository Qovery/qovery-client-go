# ClusterOperatorStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **string** |  | 
**ClusterId** | **string** |  | 
**OperatorConnected** | **bool** | Whether the last heartbeat is within the Operator presence window. | 
**LastHeartbeat** | Pointer to **NullableTime** |  | [optional] 
**OperatorVersion** | Pointer to **NullableString** | Display version reported by the Operator. For the POC version-reporting heartbeat, the official chart sets this to the exact installed image tag. Legacy Operators can report opaque build metadata instead. | [optional] 
**ControllerVersion** | Pointer to **NullableString** |  | [optional] 
**RequestSchemaVersion** | Pointer to **NullableString** |  | [optional] 
**DesiredImageVersion** | Pointer to **NullableString** | Image tag currently selected for a newly compiled Operator bootstrap. | [optional] 
**DesiredChartVersion** | Pointer to **NullableString** | Helm chart version currently selected for a newly compiled Operator bootstrap. | [optional] 
**Status** | [**ClusterOperatorFleetStatus**](ClusterOperatorFleetStatus.md) |  | 
**ReportedChartVersion** | Pointer to **NullableString** | Helm chart version reported by the Operator, even without immutable identity. | [optional] 
**ReportedIdentity** | Pointer to [**NullableReportedClusterOperatorIdentity**](ReportedClusterOperatorIdentity.md) |  | [optional] 

## Methods

### NewClusterOperatorStatusResponse

`func NewClusterOperatorStatusResponse(organizationId string, clusterId string, operatorConnected bool, status ClusterOperatorFleetStatus, ) *ClusterOperatorStatusResponse`

NewClusterOperatorStatusResponse instantiates a new ClusterOperatorStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOperatorStatusResponseWithDefaults

`func NewClusterOperatorStatusResponseWithDefaults() *ClusterOperatorStatusResponse`

NewClusterOperatorStatusResponseWithDefaults instantiates a new ClusterOperatorStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *ClusterOperatorStatusResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ClusterOperatorStatusResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ClusterOperatorStatusResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetClusterId

`func (o *ClusterOperatorStatusResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterOperatorStatusResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterOperatorStatusResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetOperatorConnected

`func (o *ClusterOperatorStatusResponse) GetOperatorConnected() bool`

GetOperatorConnected returns the OperatorConnected field if non-nil, zero value otherwise.

### GetOperatorConnectedOk

`func (o *ClusterOperatorStatusResponse) GetOperatorConnectedOk() (*bool, bool)`

GetOperatorConnectedOk returns a tuple with the OperatorConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorConnected

`func (o *ClusterOperatorStatusResponse) SetOperatorConnected(v bool)`

SetOperatorConnected sets OperatorConnected field to given value.


### GetLastHeartbeat

`func (o *ClusterOperatorStatusResponse) GetLastHeartbeat() time.Time`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *ClusterOperatorStatusResponse) GetLastHeartbeatOk() (*time.Time, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *ClusterOperatorStatusResponse) SetLastHeartbeat(v time.Time)`

SetLastHeartbeat sets LastHeartbeat field to given value.

### HasLastHeartbeat

`func (o *ClusterOperatorStatusResponse) HasLastHeartbeat() bool`

HasLastHeartbeat returns a boolean if a field has been set.

### SetLastHeartbeatNil

`func (o *ClusterOperatorStatusResponse) SetLastHeartbeatNil(b bool)`

 SetLastHeartbeatNil sets the value for LastHeartbeat to be an explicit nil

### UnsetLastHeartbeat
`func (o *ClusterOperatorStatusResponse) UnsetLastHeartbeat()`

UnsetLastHeartbeat ensures that no value is present for LastHeartbeat, not even an explicit nil
### GetOperatorVersion

`func (o *ClusterOperatorStatusResponse) GetOperatorVersion() string`

GetOperatorVersion returns the OperatorVersion field if non-nil, zero value otherwise.

### GetOperatorVersionOk

`func (o *ClusterOperatorStatusResponse) GetOperatorVersionOk() (*string, bool)`

GetOperatorVersionOk returns a tuple with the OperatorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorVersion

`func (o *ClusterOperatorStatusResponse) SetOperatorVersion(v string)`

SetOperatorVersion sets OperatorVersion field to given value.

### HasOperatorVersion

`func (o *ClusterOperatorStatusResponse) HasOperatorVersion() bool`

HasOperatorVersion returns a boolean if a field has been set.

### SetOperatorVersionNil

`func (o *ClusterOperatorStatusResponse) SetOperatorVersionNil(b bool)`

 SetOperatorVersionNil sets the value for OperatorVersion to be an explicit nil

### UnsetOperatorVersion
`func (o *ClusterOperatorStatusResponse) UnsetOperatorVersion()`

UnsetOperatorVersion ensures that no value is present for OperatorVersion, not even an explicit nil
### GetControllerVersion

`func (o *ClusterOperatorStatusResponse) GetControllerVersion() string`

GetControllerVersion returns the ControllerVersion field if non-nil, zero value otherwise.

### GetControllerVersionOk

`func (o *ClusterOperatorStatusResponse) GetControllerVersionOk() (*string, bool)`

GetControllerVersionOk returns a tuple with the ControllerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVersion

`func (o *ClusterOperatorStatusResponse) SetControllerVersion(v string)`

SetControllerVersion sets ControllerVersion field to given value.

### HasControllerVersion

`func (o *ClusterOperatorStatusResponse) HasControllerVersion() bool`

HasControllerVersion returns a boolean if a field has been set.

### SetControllerVersionNil

`func (o *ClusterOperatorStatusResponse) SetControllerVersionNil(b bool)`

 SetControllerVersionNil sets the value for ControllerVersion to be an explicit nil

### UnsetControllerVersion
`func (o *ClusterOperatorStatusResponse) UnsetControllerVersion()`

UnsetControllerVersion ensures that no value is present for ControllerVersion, not even an explicit nil
### GetRequestSchemaVersion

`func (o *ClusterOperatorStatusResponse) GetRequestSchemaVersion() string`

GetRequestSchemaVersion returns the RequestSchemaVersion field if non-nil, zero value otherwise.

### GetRequestSchemaVersionOk

`func (o *ClusterOperatorStatusResponse) GetRequestSchemaVersionOk() (*string, bool)`

GetRequestSchemaVersionOk returns a tuple with the RequestSchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSchemaVersion

`func (o *ClusterOperatorStatusResponse) SetRequestSchemaVersion(v string)`

SetRequestSchemaVersion sets RequestSchemaVersion field to given value.

### HasRequestSchemaVersion

`func (o *ClusterOperatorStatusResponse) HasRequestSchemaVersion() bool`

HasRequestSchemaVersion returns a boolean if a field has been set.

### SetRequestSchemaVersionNil

`func (o *ClusterOperatorStatusResponse) SetRequestSchemaVersionNil(b bool)`

 SetRequestSchemaVersionNil sets the value for RequestSchemaVersion to be an explicit nil

### UnsetRequestSchemaVersion
`func (o *ClusterOperatorStatusResponse) UnsetRequestSchemaVersion()`

UnsetRequestSchemaVersion ensures that no value is present for RequestSchemaVersion, not even an explicit nil
### GetDesiredImageVersion

`func (o *ClusterOperatorStatusResponse) GetDesiredImageVersion() string`

GetDesiredImageVersion returns the DesiredImageVersion field if non-nil, zero value otherwise.

### GetDesiredImageVersionOk

`func (o *ClusterOperatorStatusResponse) GetDesiredImageVersionOk() (*string, bool)`

GetDesiredImageVersionOk returns a tuple with the DesiredImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredImageVersion

`func (o *ClusterOperatorStatusResponse) SetDesiredImageVersion(v string)`

SetDesiredImageVersion sets DesiredImageVersion field to given value.

### HasDesiredImageVersion

`func (o *ClusterOperatorStatusResponse) HasDesiredImageVersion() bool`

HasDesiredImageVersion returns a boolean if a field has been set.

### SetDesiredImageVersionNil

`func (o *ClusterOperatorStatusResponse) SetDesiredImageVersionNil(b bool)`

 SetDesiredImageVersionNil sets the value for DesiredImageVersion to be an explicit nil

### UnsetDesiredImageVersion
`func (o *ClusterOperatorStatusResponse) UnsetDesiredImageVersion()`

UnsetDesiredImageVersion ensures that no value is present for DesiredImageVersion, not even an explicit nil
### GetDesiredChartVersion

`func (o *ClusterOperatorStatusResponse) GetDesiredChartVersion() string`

GetDesiredChartVersion returns the DesiredChartVersion field if non-nil, zero value otherwise.

### GetDesiredChartVersionOk

`func (o *ClusterOperatorStatusResponse) GetDesiredChartVersionOk() (*string, bool)`

GetDesiredChartVersionOk returns a tuple with the DesiredChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredChartVersion

`func (o *ClusterOperatorStatusResponse) SetDesiredChartVersion(v string)`

SetDesiredChartVersion sets DesiredChartVersion field to given value.

### HasDesiredChartVersion

`func (o *ClusterOperatorStatusResponse) HasDesiredChartVersion() bool`

HasDesiredChartVersion returns a boolean if a field has been set.

### SetDesiredChartVersionNil

`func (o *ClusterOperatorStatusResponse) SetDesiredChartVersionNil(b bool)`

 SetDesiredChartVersionNil sets the value for DesiredChartVersion to be an explicit nil

### UnsetDesiredChartVersion
`func (o *ClusterOperatorStatusResponse) UnsetDesiredChartVersion()`

UnsetDesiredChartVersion ensures that no value is present for DesiredChartVersion, not even an explicit nil
### GetStatus

`func (o *ClusterOperatorStatusResponse) GetStatus() ClusterOperatorFleetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterOperatorStatusResponse) GetStatusOk() (*ClusterOperatorFleetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterOperatorStatusResponse) SetStatus(v ClusterOperatorFleetStatus)`

SetStatus sets Status field to given value.


### GetReportedChartVersion

`func (o *ClusterOperatorStatusResponse) GetReportedChartVersion() string`

GetReportedChartVersion returns the ReportedChartVersion field if non-nil, zero value otherwise.

### GetReportedChartVersionOk

`func (o *ClusterOperatorStatusResponse) GetReportedChartVersionOk() (*string, bool)`

GetReportedChartVersionOk returns a tuple with the ReportedChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedChartVersion

`func (o *ClusterOperatorStatusResponse) SetReportedChartVersion(v string)`

SetReportedChartVersion sets ReportedChartVersion field to given value.

### HasReportedChartVersion

`func (o *ClusterOperatorStatusResponse) HasReportedChartVersion() bool`

HasReportedChartVersion returns a boolean if a field has been set.

### SetReportedChartVersionNil

`func (o *ClusterOperatorStatusResponse) SetReportedChartVersionNil(b bool)`

 SetReportedChartVersionNil sets the value for ReportedChartVersion to be an explicit nil

### UnsetReportedChartVersion
`func (o *ClusterOperatorStatusResponse) UnsetReportedChartVersion()`

UnsetReportedChartVersion ensures that no value is present for ReportedChartVersion, not even an explicit nil
### GetReportedIdentity

`func (o *ClusterOperatorStatusResponse) GetReportedIdentity() ReportedClusterOperatorIdentity`

GetReportedIdentity returns the ReportedIdentity field if non-nil, zero value otherwise.

### GetReportedIdentityOk

`func (o *ClusterOperatorStatusResponse) GetReportedIdentityOk() (*ReportedClusterOperatorIdentity, bool)`

GetReportedIdentityOk returns a tuple with the ReportedIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedIdentity

`func (o *ClusterOperatorStatusResponse) SetReportedIdentity(v ReportedClusterOperatorIdentity)`

SetReportedIdentity sets ReportedIdentity field to given value.

### HasReportedIdentity

`func (o *ClusterOperatorStatusResponse) HasReportedIdentity() bool`

HasReportedIdentity returns a boolean if a field has been set.

### SetReportedIdentityNil

`func (o *ClusterOperatorStatusResponse) SetReportedIdentityNil(b bool)`

 SetReportedIdentityNil sets the value for ReportedIdentity to be an explicit nil

### UnsetReportedIdentity
`func (o *ClusterOperatorStatusResponse) UnsetReportedIdentity()`

UnsetReportedIdentity ensures that no value is present for ReportedIdentity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


