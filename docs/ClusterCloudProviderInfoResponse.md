# ClusterCloudProviderInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to [**CloudProviderEnum**](CloudProviderEnum.md) |  | [optional] 
**Credentials** | Pointer to [**ClusterCloudProviderInfoRequestCredentials**](ClusterCloudProviderInfoRequestCredentials.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterCloudProviderInfoResponse

`func NewClusterCloudProviderInfoResponse() *ClusterCloudProviderInfoResponse`

NewClusterCloudProviderInfoResponse instantiates a new ClusterCloudProviderInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCloudProviderInfoResponseWithDefaults

`func NewClusterCloudProviderInfoResponseWithDefaults() *ClusterCloudProviderInfoResponse`

NewClusterCloudProviderInfoResponseWithDefaults instantiates a new ClusterCloudProviderInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ClusterCloudProviderInfoResponse) GetCloudProvider() CloudProviderEnum`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ClusterCloudProviderInfoResponse) GetCloudProviderOk() (*CloudProviderEnum, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ClusterCloudProviderInfoResponse) SetCloudProvider(v CloudProviderEnum)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ClusterCloudProviderInfoResponse) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCredentials

`func (o *ClusterCloudProviderInfoResponse) GetCredentials() ClusterCloudProviderInfoRequestCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ClusterCloudProviderInfoResponse) GetCredentialsOk() (*ClusterCloudProviderInfoRequestCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ClusterCloudProviderInfoResponse) SetCredentials(v ClusterCloudProviderInfoRequestCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ClusterCloudProviderInfoResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetRegion

`func (o *ClusterCloudProviderInfoResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterCloudProviderInfoResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterCloudProviderInfoResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ClusterCloudProviderInfoResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


