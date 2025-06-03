# ClusterRequestFeaturesInnerValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsVpcEksId** | **string** |  | 
**EksSubnetsZoneAIds** | **[]string** |  | 
**EksSubnetsZoneBIds** | **[]string** |  | 
**EksSubnetsZoneCIds** | **[]string** |  | 
**DocumentdbSubnetsZoneAIds** | Pointer to **[]string** |  | [optional] 
**DocumentdbSubnetsZoneBIds** | Pointer to **[]string** |  | [optional] 
**DocumentdbSubnetsZoneCIds** | Pointer to **[]string** |  | [optional] 
**ElasticacheSubnetsZoneAIds** | Pointer to **[]string** |  | [optional] 
**ElasticacheSubnetsZoneBIds** | Pointer to **[]string** |  | [optional] 
**ElasticacheSubnetsZoneCIds** | Pointer to **[]string** |  | [optional] 
**RdsSubnetsZoneAIds** | Pointer to **[]string** |  | [optional] 
**RdsSubnetsZoneBIds** | Pointer to **[]string** |  | [optional] 
**RdsSubnetsZoneCIds** | Pointer to **[]string** |  | [optional] 
**EksKarpenterFargateSubnetsZoneAIds** | Pointer to **[]string** |  | [optional] 
**EksKarpenterFargateSubnetsZoneBIds** | Pointer to **[]string** |  | [optional] 
**EksKarpenterFargateSubnetsZoneCIds** | Pointer to **[]string** |  | [optional] 
**VpcName** | **string** |  | 
**VpcProjectId** | Pointer to **string** |  | [optional] 
**SubnetworkName** | Pointer to **string** |  | [optional] 
**IpRangeServicesName** | Pointer to **string** |  | [optional] 
**IpRangePodsName** | Pointer to **string** |  | [optional] 
**AdditionalIpRangePodsNames** | Pointer to **[]string** |  | [optional] 
**SpotEnabled** | **bool** |  | 
**DiskSizeInGib** | **int32** |  | 
**DefaultServiceArchitecture** | [**CpuArchitectureEnum**](CpuArchitectureEnum.md) |  | 
**QoveryNodePools** | [**KarpenterNodePool**](KarpenterNodePool.md) |  | 

## Methods

### NewClusterRequestFeaturesInnerValue

`func NewClusterRequestFeaturesInnerValue(awsVpcEksId string, eksSubnetsZoneAIds []string, eksSubnetsZoneBIds []string, eksSubnetsZoneCIds []string, vpcName string, spotEnabled bool, diskSizeInGib int32, defaultServiceArchitecture CpuArchitectureEnum, qoveryNodePools KarpenterNodePool, ) *ClusterRequestFeaturesInnerValue`

NewClusterRequestFeaturesInnerValue instantiates a new ClusterRequestFeaturesInnerValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRequestFeaturesInnerValueWithDefaults

`func NewClusterRequestFeaturesInnerValueWithDefaults() *ClusterRequestFeaturesInnerValue`

NewClusterRequestFeaturesInnerValueWithDefaults instantiates a new ClusterRequestFeaturesInnerValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsVpcEksId

`func (o *ClusterRequestFeaturesInnerValue) GetAwsVpcEksId() string`

GetAwsVpcEksId returns the AwsVpcEksId field if non-nil, zero value otherwise.

### GetAwsVpcEksIdOk

`func (o *ClusterRequestFeaturesInnerValue) GetAwsVpcEksIdOk() (*string, bool)`

GetAwsVpcEksIdOk returns a tuple with the AwsVpcEksId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsVpcEksId

`func (o *ClusterRequestFeaturesInnerValue) SetAwsVpcEksId(v string)`

SetAwsVpcEksId sets AwsVpcEksId field to given value.


### GetEksSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) GetEksSubnetsZoneAIds() []string`

GetEksSubnetsZoneAIds returns the EksSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetEksSubnetsZoneAIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetEksSubnetsZoneAIdsOk() (*[]string, bool)`

GetEksSubnetsZoneAIdsOk returns a tuple with the EksSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) SetEksSubnetsZoneAIds(v []string)`

SetEksSubnetsZoneAIds sets EksSubnetsZoneAIds field to given value.


### GetEksSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) GetEksSubnetsZoneBIds() []string`

GetEksSubnetsZoneBIds returns the EksSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetEksSubnetsZoneBIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetEksSubnetsZoneBIdsOk() (*[]string, bool)`

GetEksSubnetsZoneBIdsOk returns a tuple with the EksSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) SetEksSubnetsZoneBIds(v []string)`

SetEksSubnetsZoneBIds sets EksSubnetsZoneBIds field to given value.


### GetEksSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) GetEksSubnetsZoneCIds() []string`

GetEksSubnetsZoneCIds returns the EksSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetEksSubnetsZoneCIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetEksSubnetsZoneCIdsOk() (*[]string, bool)`

GetEksSubnetsZoneCIdsOk returns a tuple with the EksSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) SetEksSubnetsZoneCIds(v []string)`

SetEksSubnetsZoneCIds sets EksSubnetsZoneCIds field to given value.


### GetDocumentdbSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) GetDocumentdbSubnetsZoneAIds() []string`

GetDocumentdbSubnetsZoneAIds returns the DocumentdbSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetDocumentdbSubnetsZoneAIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetDocumentdbSubnetsZoneAIdsOk() (*[]string, bool)`

GetDocumentdbSubnetsZoneAIdsOk returns a tuple with the DocumentdbSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentdbSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) SetDocumentdbSubnetsZoneAIds(v []string)`

SetDocumentdbSubnetsZoneAIds sets DocumentdbSubnetsZoneAIds field to given value.

### HasDocumentdbSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) HasDocumentdbSubnetsZoneAIds() bool`

HasDocumentdbSubnetsZoneAIds returns a boolean if a field has been set.

### GetDocumentdbSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) GetDocumentdbSubnetsZoneBIds() []string`

GetDocumentdbSubnetsZoneBIds returns the DocumentdbSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetDocumentdbSubnetsZoneBIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetDocumentdbSubnetsZoneBIdsOk() (*[]string, bool)`

GetDocumentdbSubnetsZoneBIdsOk returns a tuple with the DocumentdbSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentdbSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) SetDocumentdbSubnetsZoneBIds(v []string)`

SetDocumentdbSubnetsZoneBIds sets DocumentdbSubnetsZoneBIds field to given value.

### HasDocumentdbSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) HasDocumentdbSubnetsZoneBIds() bool`

HasDocumentdbSubnetsZoneBIds returns a boolean if a field has been set.

### GetDocumentdbSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) GetDocumentdbSubnetsZoneCIds() []string`

GetDocumentdbSubnetsZoneCIds returns the DocumentdbSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetDocumentdbSubnetsZoneCIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetDocumentdbSubnetsZoneCIdsOk() (*[]string, bool)`

GetDocumentdbSubnetsZoneCIdsOk returns a tuple with the DocumentdbSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentdbSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) SetDocumentdbSubnetsZoneCIds(v []string)`

SetDocumentdbSubnetsZoneCIds sets DocumentdbSubnetsZoneCIds field to given value.

### HasDocumentdbSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) HasDocumentdbSubnetsZoneCIds() bool`

HasDocumentdbSubnetsZoneCIds returns a boolean if a field has been set.

### GetElasticacheSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) GetElasticacheSubnetsZoneAIds() []string`

GetElasticacheSubnetsZoneAIds returns the ElasticacheSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetElasticacheSubnetsZoneAIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetElasticacheSubnetsZoneAIdsOk() (*[]string, bool)`

GetElasticacheSubnetsZoneAIdsOk returns a tuple with the ElasticacheSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticacheSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) SetElasticacheSubnetsZoneAIds(v []string)`

SetElasticacheSubnetsZoneAIds sets ElasticacheSubnetsZoneAIds field to given value.

### HasElasticacheSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) HasElasticacheSubnetsZoneAIds() bool`

HasElasticacheSubnetsZoneAIds returns a boolean if a field has been set.

### GetElasticacheSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) GetElasticacheSubnetsZoneBIds() []string`

GetElasticacheSubnetsZoneBIds returns the ElasticacheSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetElasticacheSubnetsZoneBIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetElasticacheSubnetsZoneBIdsOk() (*[]string, bool)`

GetElasticacheSubnetsZoneBIdsOk returns a tuple with the ElasticacheSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticacheSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) SetElasticacheSubnetsZoneBIds(v []string)`

SetElasticacheSubnetsZoneBIds sets ElasticacheSubnetsZoneBIds field to given value.

### HasElasticacheSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) HasElasticacheSubnetsZoneBIds() bool`

HasElasticacheSubnetsZoneBIds returns a boolean if a field has been set.

### GetElasticacheSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) GetElasticacheSubnetsZoneCIds() []string`

GetElasticacheSubnetsZoneCIds returns the ElasticacheSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetElasticacheSubnetsZoneCIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetElasticacheSubnetsZoneCIdsOk() (*[]string, bool)`

GetElasticacheSubnetsZoneCIdsOk returns a tuple with the ElasticacheSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticacheSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) SetElasticacheSubnetsZoneCIds(v []string)`

SetElasticacheSubnetsZoneCIds sets ElasticacheSubnetsZoneCIds field to given value.

### HasElasticacheSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) HasElasticacheSubnetsZoneCIds() bool`

HasElasticacheSubnetsZoneCIds returns a boolean if a field has been set.

### GetRdsSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) GetRdsSubnetsZoneAIds() []string`

GetRdsSubnetsZoneAIds returns the RdsSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetRdsSubnetsZoneAIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetRdsSubnetsZoneAIdsOk() (*[]string, bool)`

GetRdsSubnetsZoneAIdsOk returns a tuple with the RdsSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) SetRdsSubnetsZoneAIds(v []string)`

SetRdsSubnetsZoneAIds sets RdsSubnetsZoneAIds field to given value.

### HasRdsSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) HasRdsSubnetsZoneAIds() bool`

HasRdsSubnetsZoneAIds returns a boolean if a field has been set.

### GetRdsSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) GetRdsSubnetsZoneBIds() []string`

GetRdsSubnetsZoneBIds returns the RdsSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetRdsSubnetsZoneBIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetRdsSubnetsZoneBIdsOk() (*[]string, bool)`

GetRdsSubnetsZoneBIdsOk returns a tuple with the RdsSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) SetRdsSubnetsZoneBIds(v []string)`

SetRdsSubnetsZoneBIds sets RdsSubnetsZoneBIds field to given value.

### HasRdsSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) HasRdsSubnetsZoneBIds() bool`

HasRdsSubnetsZoneBIds returns a boolean if a field has been set.

### GetRdsSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) GetRdsSubnetsZoneCIds() []string`

GetRdsSubnetsZoneCIds returns the RdsSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetRdsSubnetsZoneCIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetRdsSubnetsZoneCIdsOk() (*[]string, bool)`

GetRdsSubnetsZoneCIdsOk returns a tuple with the RdsSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) SetRdsSubnetsZoneCIds(v []string)`

SetRdsSubnetsZoneCIds sets RdsSubnetsZoneCIds field to given value.

### HasRdsSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) HasRdsSubnetsZoneCIds() bool`

HasRdsSubnetsZoneCIds returns a boolean if a field has been set.

### GetEksKarpenterFargateSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) GetEksKarpenterFargateSubnetsZoneAIds() []string`

GetEksKarpenterFargateSubnetsZoneAIds returns the EksKarpenterFargateSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetEksKarpenterFargateSubnetsZoneAIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetEksKarpenterFargateSubnetsZoneAIdsOk() (*[]string, bool)`

GetEksKarpenterFargateSubnetsZoneAIdsOk returns a tuple with the EksKarpenterFargateSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksKarpenterFargateSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) SetEksKarpenterFargateSubnetsZoneAIds(v []string)`

SetEksKarpenterFargateSubnetsZoneAIds sets EksKarpenterFargateSubnetsZoneAIds field to given value.

### HasEksKarpenterFargateSubnetsZoneAIds

`func (o *ClusterRequestFeaturesInnerValue) HasEksKarpenterFargateSubnetsZoneAIds() bool`

HasEksKarpenterFargateSubnetsZoneAIds returns a boolean if a field has been set.

### GetEksKarpenterFargateSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) GetEksKarpenterFargateSubnetsZoneBIds() []string`

GetEksKarpenterFargateSubnetsZoneBIds returns the EksKarpenterFargateSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetEksKarpenterFargateSubnetsZoneBIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetEksKarpenterFargateSubnetsZoneBIdsOk() (*[]string, bool)`

GetEksKarpenterFargateSubnetsZoneBIdsOk returns a tuple with the EksKarpenterFargateSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksKarpenterFargateSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) SetEksKarpenterFargateSubnetsZoneBIds(v []string)`

SetEksKarpenterFargateSubnetsZoneBIds sets EksKarpenterFargateSubnetsZoneBIds field to given value.

### HasEksKarpenterFargateSubnetsZoneBIds

`func (o *ClusterRequestFeaturesInnerValue) HasEksKarpenterFargateSubnetsZoneBIds() bool`

HasEksKarpenterFargateSubnetsZoneBIds returns a boolean if a field has been set.

### GetEksKarpenterFargateSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) GetEksKarpenterFargateSubnetsZoneCIds() []string`

GetEksKarpenterFargateSubnetsZoneCIds returns the EksKarpenterFargateSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetEksKarpenterFargateSubnetsZoneCIdsOk

`func (o *ClusterRequestFeaturesInnerValue) GetEksKarpenterFargateSubnetsZoneCIdsOk() (*[]string, bool)`

GetEksKarpenterFargateSubnetsZoneCIdsOk returns a tuple with the EksKarpenterFargateSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksKarpenterFargateSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) SetEksKarpenterFargateSubnetsZoneCIds(v []string)`

SetEksKarpenterFargateSubnetsZoneCIds sets EksKarpenterFargateSubnetsZoneCIds field to given value.

### HasEksKarpenterFargateSubnetsZoneCIds

`func (o *ClusterRequestFeaturesInnerValue) HasEksKarpenterFargateSubnetsZoneCIds() bool`

HasEksKarpenterFargateSubnetsZoneCIds returns a boolean if a field has been set.

### GetVpcName

`func (o *ClusterRequestFeaturesInnerValue) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *ClusterRequestFeaturesInnerValue) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *ClusterRequestFeaturesInnerValue) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.


### GetVpcProjectId

`func (o *ClusterRequestFeaturesInnerValue) GetVpcProjectId() string`

GetVpcProjectId returns the VpcProjectId field if non-nil, zero value otherwise.

### GetVpcProjectIdOk

`func (o *ClusterRequestFeaturesInnerValue) GetVpcProjectIdOk() (*string, bool)`

GetVpcProjectIdOk returns a tuple with the VpcProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcProjectId

`func (o *ClusterRequestFeaturesInnerValue) SetVpcProjectId(v string)`

SetVpcProjectId sets VpcProjectId field to given value.

### HasVpcProjectId

`func (o *ClusterRequestFeaturesInnerValue) HasVpcProjectId() bool`

HasVpcProjectId returns a boolean if a field has been set.

### GetSubnetworkName

`func (o *ClusterRequestFeaturesInnerValue) GetSubnetworkName() string`

GetSubnetworkName returns the SubnetworkName field if non-nil, zero value otherwise.

### GetSubnetworkNameOk

`func (o *ClusterRequestFeaturesInnerValue) GetSubnetworkNameOk() (*string, bool)`

GetSubnetworkNameOk returns a tuple with the SubnetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetworkName

`func (o *ClusterRequestFeaturesInnerValue) SetSubnetworkName(v string)`

SetSubnetworkName sets SubnetworkName field to given value.

### HasSubnetworkName

`func (o *ClusterRequestFeaturesInnerValue) HasSubnetworkName() bool`

HasSubnetworkName returns a boolean if a field has been set.

### GetIpRangeServicesName

`func (o *ClusterRequestFeaturesInnerValue) GetIpRangeServicesName() string`

GetIpRangeServicesName returns the IpRangeServicesName field if non-nil, zero value otherwise.

### GetIpRangeServicesNameOk

`func (o *ClusterRequestFeaturesInnerValue) GetIpRangeServicesNameOk() (*string, bool)`

GetIpRangeServicesNameOk returns a tuple with the IpRangeServicesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangeServicesName

`func (o *ClusterRequestFeaturesInnerValue) SetIpRangeServicesName(v string)`

SetIpRangeServicesName sets IpRangeServicesName field to given value.

### HasIpRangeServicesName

`func (o *ClusterRequestFeaturesInnerValue) HasIpRangeServicesName() bool`

HasIpRangeServicesName returns a boolean if a field has been set.

### GetIpRangePodsName

`func (o *ClusterRequestFeaturesInnerValue) GetIpRangePodsName() string`

GetIpRangePodsName returns the IpRangePodsName field if non-nil, zero value otherwise.

### GetIpRangePodsNameOk

`func (o *ClusterRequestFeaturesInnerValue) GetIpRangePodsNameOk() (*string, bool)`

GetIpRangePodsNameOk returns a tuple with the IpRangePodsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangePodsName

`func (o *ClusterRequestFeaturesInnerValue) SetIpRangePodsName(v string)`

SetIpRangePodsName sets IpRangePodsName field to given value.

### HasIpRangePodsName

`func (o *ClusterRequestFeaturesInnerValue) HasIpRangePodsName() bool`

HasIpRangePodsName returns a boolean if a field has been set.

### GetAdditionalIpRangePodsNames

`func (o *ClusterRequestFeaturesInnerValue) GetAdditionalIpRangePodsNames() []string`

GetAdditionalIpRangePodsNames returns the AdditionalIpRangePodsNames field if non-nil, zero value otherwise.

### GetAdditionalIpRangePodsNamesOk

`func (o *ClusterRequestFeaturesInnerValue) GetAdditionalIpRangePodsNamesOk() (*[]string, bool)`

GetAdditionalIpRangePodsNamesOk returns a tuple with the AdditionalIpRangePodsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIpRangePodsNames

`func (o *ClusterRequestFeaturesInnerValue) SetAdditionalIpRangePodsNames(v []string)`

SetAdditionalIpRangePodsNames sets AdditionalIpRangePodsNames field to given value.

### HasAdditionalIpRangePodsNames

`func (o *ClusterRequestFeaturesInnerValue) HasAdditionalIpRangePodsNames() bool`

HasAdditionalIpRangePodsNames returns a boolean if a field has been set.

### GetSpotEnabled

`func (o *ClusterRequestFeaturesInnerValue) GetSpotEnabled() bool`

GetSpotEnabled returns the SpotEnabled field if non-nil, zero value otherwise.

### GetSpotEnabledOk

`func (o *ClusterRequestFeaturesInnerValue) GetSpotEnabledOk() (*bool, bool)`

GetSpotEnabledOk returns a tuple with the SpotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotEnabled

`func (o *ClusterRequestFeaturesInnerValue) SetSpotEnabled(v bool)`

SetSpotEnabled sets SpotEnabled field to given value.


### GetDiskSizeInGib

`func (o *ClusterRequestFeaturesInnerValue) GetDiskSizeInGib() int32`

GetDiskSizeInGib returns the DiskSizeInGib field if non-nil, zero value otherwise.

### GetDiskSizeInGibOk

`func (o *ClusterRequestFeaturesInnerValue) GetDiskSizeInGibOk() (*int32, bool)`

GetDiskSizeInGibOk returns a tuple with the DiskSizeInGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeInGib

`func (o *ClusterRequestFeaturesInnerValue) SetDiskSizeInGib(v int32)`

SetDiskSizeInGib sets DiskSizeInGib field to given value.


### GetDefaultServiceArchitecture

`func (o *ClusterRequestFeaturesInnerValue) GetDefaultServiceArchitecture() CpuArchitectureEnum`

GetDefaultServiceArchitecture returns the DefaultServiceArchitecture field if non-nil, zero value otherwise.

### GetDefaultServiceArchitectureOk

`func (o *ClusterRequestFeaturesInnerValue) GetDefaultServiceArchitectureOk() (*CpuArchitectureEnum, bool)`

GetDefaultServiceArchitectureOk returns a tuple with the DefaultServiceArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServiceArchitecture

`func (o *ClusterRequestFeaturesInnerValue) SetDefaultServiceArchitecture(v CpuArchitectureEnum)`

SetDefaultServiceArchitecture sets DefaultServiceArchitecture field to given value.


### GetQoveryNodePools

`func (o *ClusterRequestFeaturesInnerValue) GetQoveryNodePools() KarpenterNodePool`

GetQoveryNodePools returns the QoveryNodePools field if non-nil, zero value otherwise.

### GetQoveryNodePoolsOk

`func (o *ClusterRequestFeaturesInnerValue) GetQoveryNodePoolsOk() (*KarpenterNodePool, bool)`

GetQoveryNodePoolsOk returns a tuple with the QoveryNodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryNodePools

`func (o *ClusterRequestFeaturesInnerValue) SetQoveryNodePools(v KarpenterNodePool)`

SetQoveryNodePools sets QoveryNodePools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


