# ClusterFeatureGcpExistingVpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcName** | **string** |  | 
**VpcProjectId** | Pointer to **string** |  | [optional] 
**SubnetworkName** | Pointer to **string** |  | [optional] 
**IpRangeServicesName** | Pointer to **string** |  | [optional] 
**IpRangePodsName** | Pointer to **string** |  | [optional] 
**AdditionalIpRangePodsNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewClusterFeatureGcpExistingVpc

`func NewClusterFeatureGcpExistingVpc(vpcName string, ) *ClusterFeatureGcpExistingVpc`

NewClusterFeatureGcpExistingVpc instantiates a new ClusterFeatureGcpExistingVpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFeatureGcpExistingVpcWithDefaults

`func NewClusterFeatureGcpExistingVpcWithDefaults() *ClusterFeatureGcpExistingVpc`

NewClusterFeatureGcpExistingVpcWithDefaults instantiates a new ClusterFeatureGcpExistingVpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcName

`func (o *ClusterFeatureGcpExistingVpc) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *ClusterFeatureGcpExistingVpc) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *ClusterFeatureGcpExistingVpc) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.


### GetVpcProjectId

`func (o *ClusterFeatureGcpExistingVpc) GetVpcProjectId() string`

GetVpcProjectId returns the VpcProjectId field if non-nil, zero value otherwise.

### GetVpcProjectIdOk

`func (o *ClusterFeatureGcpExistingVpc) GetVpcProjectIdOk() (*string, bool)`

GetVpcProjectIdOk returns a tuple with the VpcProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcProjectId

`func (o *ClusterFeatureGcpExistingVpc) SetVpcProjectId(v string)`

SetVpcProjectId sets VpcProjectId field to given value.

### HasVpcProjectId

`func (o *ClusterFeatureGcpExistingVpc) HasVpcProjectId() bool`

HasVpcProjectId returns a boolean if a field has been set.

### GetSubnetworkName

`func (o *ClusterFeatureGcpExistingVpc) GetSubnetworkName() string`

GetSubnetworkName returns the SubnetworkName field if non-nil, zero value otherwise.

### GetSubnetworkNameOk

`func (o *ClusterFeatureGcpExistingVpc) GetSubnetworkNameOk() (*string, bool)`

GetSubnetworkNameOk returns a tuple with the SubnetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetworkName

`func (o *ClusterFeatureGcpExistingVpc) SetSubnetworkName(v string)`

SetSubnetworkName sets SubnetworkName field to given value.

### HasSubnetworkName

`func (o *ClusterFeatureGcpExistingVpc) HasSubnetworkName() bool`

HasSubnetworkName returns a boolean if a field has been set.

### GetIpRangeServicesName

`func (o *ClusterFeatureGcpExistingVpc) GetIpRangeServicesName() string`

GetIpRangeServicesName returns the IpRangeServicesName field if non-nil, zero value otherwise.

### GetIpRangeServicesNameOk

`func (o *ClusterFeatureGcpExistingVpc) GetIpRangeServicesNameOk() (*string, bool)`

GetIpRangeServicesNameOk returns a tuple with the IpRangeServicesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangeServicesName

`func (o *ClusterFeatureGcpExistingVpc) SetIpRangeServicesName(v string)`

SetIpRangeServicesName sets IpRangeServicesName field to given value.

### HasIpRangeServicesName

`func (o *ClusterFeatureGcpExistingVpc) HasIpRangeServicesName() bool`

HasIpRangeServicesName returns a boolean if a field has been set.

### GetIpRangePodsName

`func (o *ClusterFeatureGcpExistingVpc) GetIpRangePodsName() string`

GetIpRangePodsName returns the IpRangePodsName field if non-nil, zero value otherwise.

### GetIpRangePodsNameOk

`func (o *ClusterFeatureGcpExistingVpc) GetIpRangePodsNameOk() (*string, bool)`

GetIpRangePodsNameOk returns a tuple with the IpRangePodsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangePodsName

`func (o *ClusterFeatureGcpExistingVpc) SetIpRangePodsName(v string)`

SetIpRangePodsName sets IpRangePodsName field to given value.

### HasIpRangePodsName

`func (o *ClusterFeatureGcpExistingVpc) HasIpRangePodsName() bool`

HasIpRangePodsName returns a boolean if a field has been set.

### GetAdditionalIpRangePodsNames

`func (o *ClusterFeatureGcpExistingVpc) GetAdditionalIpRangePodsNames() []string`

GetAdditionalIpRangePodsNames returns the AdditionalIpRangePodsNames field if non-nil, zero value otherwise.

### GetAdditionalIpRangePodsNamesOk

`func (o *ClusterFeatureGcpExistingVpc) GetAdditionalIpRangePodsNamesOk() (*[]string, bool)`

GetAdditionalIpRangePodsNamesOk returns a tuple with the AdditionalIpRangePodsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIpRangePodsNames

`func (o *ClusterFeatureGcpExistingVpc) SetAdditionalIpRangePodsNames(v []string)`

SetAdditionalIpRangePodsNames sets AdditionalIpRangePodsNames field to given value.

### HasAdditionalIpRangePodsNames

`func (o *ClusterFeatureGcpExistingVpc) HasAdditionalIpRangePodsNames() bool`

HasAdditionalIpRangePodsNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


