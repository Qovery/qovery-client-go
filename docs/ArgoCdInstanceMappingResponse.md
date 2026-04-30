# ArgoCdInstanceMappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentClusterId** | **string** | ID of the Qovery cluster where the ArgoCD instance is running | 
**AgentClusterName** | **string** | Display name of the Qovery cluster where the ArgoCD instance is running | 
**AgentClusterCloudProvider** | [**CloudVendorEnum**](CloudVendorEnum.md) |  | 
**CredentialsId** | **string** | ID of the stored ArgoCD credentials for this instance | 
**ArgocdUrl** | **string** | URL of the ArgoCD instance | 
**Status** | [**ArgoCdConnectionStatusEnum**](ArgoCdConnectionStatusEnum.md) |  | 
**LastCheckedAt** | **time.Time** | Timestamp of the last configuration update (will use last connectivity check in future) | 
**LinkedClusters** | [**[]ArgoCdLinkedClusterDetails**](ArgoCdLinkedClusterDetails.md) | ArgoCD clusters detected and mapped to a Qovery cluster | 
**UnlinkedClusters** | [**[]ArgoCdUnlinkedClusterDetails**](ArgoCdUnlinkedClusterDetails.md) | ArgoCD clusters detected but mapping is missing | 

## Methods

### NewArgoCdInstanceMappingResponse

`func NewArgoCdInstanceMappingResponse(agentClusterId string, agentClusterName string, agentClusterCloudProvider CloudVendorEnum, credentialsId string, argocdUrl string, status ArgoCdConnectionStatusEnum, lastCheckedAt time.Time, linkedClusters []ArgoCdLinkedClusterDetails, unlinkedClusters []ArgoCdUnlinkedClusterDetails, ) *ArgoCdInstanceMappingResponse`

NewArgoCdInstanceMappingResponse instantiates a new ArgoCdInstanceMappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoCdInstanceMappingResponseWithDefaults

`func NewArgoCdInstanceMappingResponseWithDefaults() *ArgoCdInstanceMappingResponse`

NewArgoCdInstanceMappingResponseWithDefaults instantiates a new ArgoCdInstanceMappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentClusterId

`func (o *ArgoCdInstanceMappingResponse) GetAgentClusterId() string`

GetAgentClusterId returns the AgentClusterId field if non-nil, zero value otherwise.

### GetAgentClusterIdOk

`func (o *ArgoCdInstanceMappingResponse) GetAgentClusterIdOk() (*string, bool)`

GetAgentClusterIdOk returns a tuple with the AgentClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentClusterId

`func (o *ArgoCdInstanceMappingResponse) SetAgentClusterId(v string)`

SetAgentClusterId sets AgentClusterId field to given value.


### GetAgentClusterName

`func (o *ArgoCdInstanceMappingResponse) GetAgentClusterName() string`

GetAgentClusterName returns the AgentClusterName field if non-nil, zero value otherwise.

### GetAgentClusterNameOk

`func (o *ArgoCdInstanceMappingResponse) GetAgentClusterNameOk() (*string, bool)`

GetAgentClusterNameOk returns a tuple with the AgentClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentClusterName

`func (o *ArgoCdInstanceMappingResponse) SetAgentClusterName(v string)`

SetAgentClusterName sets AgentClusterName field to given value.


### GetAgentClusterCloudProvider

`func (o *ArgoCdInstanceMappingResponse) GetAgentClusterCloudProvider() CloudVendorEnum`

GetAgentClusterCloudProvider returns the AgentClusterCloudProvider field if non-nil, zero value otherwise.

### GetAgentClusterCloudProviderOk

`func (o *ArgoCdInstanceMappingResponse) GetAgentClusterCloudProviderOk() (*CloudVendorEnum, bool)`

GetAgentClusterCloudProviderOk returns a tuple with the AgentClusterCloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentClusterCloudProvider

`func (o *ArgoCdInstanceMappingResponse) SetAgentClusterCloudProvider(v CloudVendorEnum)`

SetAgentClusterCloudProvider sets AgentClusterCloudProvider field to given value.


### GetCredentialsId

`func (o *ArgoCdInstanceMappingResponse) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *ArgoCdInstanceMappingResponse) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *ArgoCdInstanceMappingResponse) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetArgocdUrl

`func (o *ArgoCdInstanceMappingResponse) GetArgocdUrl() string`

GetArgocdUrl returns the ArgocdUrl field if non-nil, zero value otherwise.

### GetArgocdUrlOk

`func (o *ArgoCdInstanceMappingResponse) GetArgocdUrlOk() (*string, bool)`

GetArgocdUrlOk returns a tuple with the ArgocdUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgocdUrl

`func (o *ArgoCdInstanceMappingResponse) SetArgocdUrl(v string)`

SetArgocdUrl sets ArgocdUrl field to given value.


### GetStatus

`func (o *ArgoCdInstanceMappingResponse) GetStatus() ArgoCdConnectionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArgoCdInstanceMappingResponse) GetStatusOk() (*ArgoCdConnectionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArgoCdInstanceMappingResponse) SetStatus(v ArgoCdConnectionStatusEnum)`

SetStatus sets Status field to given value.


### GetLastCheckedAt

`func (o *ArgoCdInstanceMappingResponse) GetLastCheckedAt() time.Time`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *ArgoCdInstanceMappingResponse) GetLastCheckedAtOk() (*time.Time, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *ArgoCdInstanceMappingResponse) SetLastCheckedAt(v time.Time)`

SetLastCheckedAt sets LastCheckedAt field to given value.


### GetLinkedClusters

`func (o *ArgoCdInstanceMappingResponse) GetLinkedClusters() []ArgoCdLinkedClusterDetails`

GetLinkedClusters returns the LinkedClusters field if non-nil, zero value otherwise.

### GetLinkedClustersOk

`func (o *ArgoCdInstanceMappingResponse) GetLinkedClustersOk() (*[]ArgoCdLinkedClusterDetails, bool)`

GetLinkedClustersOk returns a tuple with the LinkedClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedClusters

`func (o *ArgoCdInstanceMappingResponse) SetLinkedClusters(v []ArgoCdLinkedClusterDetails)`

SetLinkedClusters sets LinkedClusters field to given value.


### GetUnlinkedClusters

`func (o *ArgoCdInstanceMappingResponse) GetUnlinkedClusters() []ArgoCdUnlinkedClusterDetails`

GetUnlinkedClusters returns the UnlinkedClusters field if non-nil, zero value otherwise.

### GetUnlinkedClustersOk

`func (o *ArgoCdInstanceMappingResponse) GetUnlinkedClustersOk() (*[]ArgoCdUnlinkedClusterDetails, bool)`

GetUnlinkedClustersOk returns a tuple with the UnlinkedClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlinkedClusters

`func (o *ArgoCdInstanceMappingResponse) SetUnlinkedClusters(v []ArgoCdUnlinkedClusterDetails)`

SetUnlinkedClusters sets UnlinkedClusters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


