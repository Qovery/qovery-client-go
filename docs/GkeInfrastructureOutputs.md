# GkeInfrastructureOutputs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**ClusterName** | **string** |  | 
**ClusterSelfLink** | **string** |  | 
**ExternalSecretsAutomaticAuthentication** | Pointer to [**NullableGkeInfrastructureOutputsExternalSecretsAutomaticAuthentication**](GkeInfrastructureOutputsExternalSecretsAutomaticAuthentication.md) |  | [optional] 

## Methods

### NewGkeInfrastructureOutputs

`func NewGkeInfrastructureOutputs(kind string, clusterName string, clusterSelfLink string, ) *GkeInfrastructureOutputs`

NewGkeInfrastructureOutputs instantiates a new GkeInfrastructureOutputs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGkeInfrastructureOutputsWithDefaults

`func NewGkeInfrastructureOutputsWithDefaults() *GkeInfrastructureOutputs`

NewGkeInfrastructureOutputsWithDefaults instantiates a new GkeInfrastructureOutputs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *GkeInfrastructureOutputs) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GkeInfrastructureOutputs) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GkeInfrastructureOutputs) SetKind(v string)`

SetKind sets Kind field to given value.


### GetClusterName

`func (o *GkeInfrastructureOutputs) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *GkeInfrastructureOutputs) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *GkeInfrastructureOutputs) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetClusterSelfLink

`func (o *GkeInfrastructureOutputs) GetClusterSelfLink() string`

GetClusterSelfLink returns the ClusterSelfLink field if non-nil, zero value otherwise.

### GetClusterSelfLinkOk

`func (o *GkeInfrastructureOutputs) GetClusterSelfLinkOk() (*string, bool)`

GetClusterSelfLinkOk returns a tuple with the ClusterSelfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSelfLink

`func (o *GkeInfrastructureOutputs) SetClusterSelfLink(v string)`

SetClusterSelfLink sets ClusterSelfLink field to given value.


### GetExternalSecretsAutomaticAuthentication

`func (o *GkeInfrastructureOutputs) GetExternalSecretsAutomaticAuthentication() GkeInfrastructureOutputsExternalSecretsAutomaticAuthentication`

GetExternalSecretsAutomaticAuthentication returns the ExternalSecretsAutomaticAuthentication field if non-nil, zero value otherwise.

### GetExternalSecretsAutomaticAuthenticationOk

`func (o *GkeInfrastructureOutputs) GetExternalSecretsAutomaticAuthenticationOk() (*GkeInfrastructureOutputsExternalSecretsAutomaticAuthentication, bool)`

GetExternalSecretsAutomaticAuthenticationOk returns a tuple with the ExternalSecretsAutomaticAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecretsAutomaticAuthentication

`func (o *GkeInfrastructureOutputs) SetExternalSecretsAutomaticAuthentication(v GkeInfrastructureOutputsExternalSecretsAutomaticAuthentication)`

SetExternalSecretsAutomaticAuthentication sets ExternalSecretsAutomaticAuthentication field to given value.

### HasExternalSecretsAutomaticAuthentication

`func (o *GkeInfrastructureOutputs) HasExternalSecretsAutomaticAuthentication() bool`

HasExternalSecretsAutomaticAuthentication returns a boolean if a field has been set.

### SetExternalSecretsAutomaticAuthenticationNil

`func (o *GkeInfrastructureOutputs) SetExternalSecretsAutomaticAuthenticationNil(b bool)`

 SetExternalSecretsAutomaticAuthenticationNil sets the value for ExternalSecretsAutomaticAuthentication to be an explicit nil

### UnsetExternalSecretsAutomaticAuthentication
`func (o *GkeInfrastructureOutputs) UnsetExternalSecretsAutomaticAuthentication()`

UnsetExternalSecretsAutomaticAuthentication ensures that no value is present for ExternalSecretsAutomaticAuthentication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


