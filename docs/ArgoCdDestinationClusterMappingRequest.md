# ArgoCdDestinationClusterMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentClusterId** | **string** | ID of the Qovery cluster where the ArgoCD instance is running | 
**ArgocdClusterUrl** | **string** | ArgoCD destination cluster URL as reported by ArgoCD | 
**ClusterId** | **string** | ID of the Qovery cluster to map this ArgoCD destination to | 

## Methods

### NewArgoCdDestinationClusterMappingRequest

`func NewArgoCdDestinationClusterMappingRequest(agentClusterId string, argocdClusterUrl string, clusterId string, ) *ArgoCdDestinationClusterMappingRequest`

NewArgoCdDestinationClusterMappingRequest instantiates a new ArgoCdDestinationClusterMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoCdDestinationClusterMappingRequestWithDefaults

`func NewArgoCdDestinationClusterMappingRequestWithDefaults() *ArgoCdDestinationClusterMappingRequest`

NewArgoCdDestinationClusterMappingRequestWithDefaults instantiates a new ArgoCdDestinationClusterMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentClusterId

`func (o *ArgoCdDestinationClusterMappingRequest) GetAgentClusterId() string`

GetAgentClusterId returns the AgentClusterId field if non-nil, zero value otherwise.

### GetAgentClusterIdOk

`func (o *ArgoCdDestinationClusterMappingRequest) GetAgentClusterIdOk() (*string, bool)`

GetAgentClusterIdOk returns a tuple with the AgentClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentClusterId

`func (o *ArgoCdDestinationClusterMappingRequest) SetAgentClusterId(v string)`

SetAgentClusterId sets AgentClusterId field to given value.


### GetArgocdClusterUrl

`func (o *ArgoCdDestinationClusterMappingRequest) GetArgocdClusterUrl() string`

GetArgocdClusterUrl returns the ArgocdClusterUrl field if non-nil, zero value otherwise.

### GetArgocdClusterUrlOk

`func (o *ArgoCdDestinationClusterMappingRequest) GetArgocdClusterUrlOk() (*string, bool)`

GetArgocdClusterUrlOk returns a tuple with the ArgocdClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgocdClusterUrl

`func (o *ArgoCdDestinationClusterMappingRequest) SetArgocdClusterUrl(v string)`

SetArgocdClusterUrl sets ArgocdClusterUrl field to given value.


### GetClusterId

`func (o *ArgoCdDestinationClusterMappingRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ArgoCdDestinationClusterMappingRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ArgoCdDestinationClusterMappingRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


