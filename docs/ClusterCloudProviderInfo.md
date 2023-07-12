# ClusterCloudProviderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to [**CloudProviderEnum**](CloudProviderEnum.md) |  | [optional] 
**Credentials** | Pointer to [**GetOrganizationEventTargets200ResponseTargetsInner**](GetOrganizationEventTargets200ResponseTargetsInner.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterCloudProviderInfo

`func NewClusterCloudProviderInfo() *ClusterCloudProviderInfo`

NewClusterCloudProviderInfo instantiates a new ClusterCloudProviderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCloudProviderInfoWithDefaults

`func NewClusterCloudProviderInfoWithDefaults() *ClusterCloudProviderInfo`

NewClusterCloudProviderInfoWithDefaults instantiates a new ClusterCloudProviderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ClusterCloudProviderInfo) GetCloudProvider() CloudProviderEnum`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ClusterCloudProviderInfo) GetCloudProviderOk() (*CloudProviderEnum, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ClusterCloudProviderInfo) SetCloudProvider(v CloudProviderEnum)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ClusterCloudProviderInfo) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCredentials

`func (o *ClusterCloudProviderInfo) GetCredentials() GetOrganizationEventTargets200ResponseTargetsInner`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ClusterCloudProviderInfo) GetCredentialsOk() (*GetOrganizationEventTargets200ResponseTargetsInner, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ClusterCloudProviderInfo) SetCredentials(v GetOrganizationEventTargets200ResponseTargetsInner)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ClusterCloudProviderInfo) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetRegion

`func (o *ClusterCloudProviderInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterCloudProviderInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterCloudProviderInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ClusterCloudProviderInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


