# ClusterFeatureAwsExistingVpc

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

## Methods

### NewClusterFeatureAwsExistingVpc

`func NewClusterFeatureAwsExistingVpc(awsVpcEksId string, eksSubnetsZoneAIds []string, eksSubnetsZoneBIds []string, eksSubnetsZoneCIds []string, ) *ClusterFeatureAwsExistingVpc`

NewClusterFeatureAwsExistingVpc instantiates a new ClusterFeatureAwsExistingVpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFeatureAwsExistingVpcWithDefaults

`func NewClusterFeatureAwsExistingVpcWithDefaults() *ClusterFeatureAwsExistingVpc`

NewClusterFeatureAwsExistingVpcWithDefaults instantiates a new ClusterFeatureAwsExistingVpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsVpcEksId

`func (o *ClusterFeatureAwsExistingVpc) GetAwsVpcEksId() string`

GetAwsVpcEksId returns the AwsVpcEksId field if non-nil, zero value otherwise.

### GetAwsVpcEksIdOk

`func (o *ClusterFeatureAwsExistingVpc) GetAwsVpcEksIdOk() (*string, bool)`

GetAwsVpcEksIdOk returns a tuple with the AwsVpcEksId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsVpcEksId

`func (o *ClusterFeatureAwsExistingVpc) SetAwsVpcEksId(v string)`

SetAwsVpcEksId sets AwsVpcEksId field to given value.


### GetEksSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) GetEksSubnetsZoneAIds() []string`

GetEksSubnetsZoneAIds returns the EksSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetEksSubnetsZoneAIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetEksSubnetsZoneAIdsOk() (*[]string, bool)`

GetEksSubnetsZoneAIdsOk returns a tuple with the EksSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) SetEksSubnetsZoneAIds(v []string)`

SetEksSubnetsZoneAIds sets EksSubnetsZoneAIds field to given value.


### GetEksSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) GetEksSubnetsZoneBIds() []string`

GetEksSubnetsZoneBIds returns the EksSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetEksSubnetsZoneBIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetEksSubnetsZoneBIdsOk() (*[]string, bool)`

GetEksSubnetsZoneBIdsOk returns a tuple with the EksSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) SetEksSubnetsZoneBIds(v []string)`

SetEksSubnetsZoneBIds sets EksSubnetsZoneBIds field to given value.


### GetEksSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) GetEksSubnetsZoneCIds() []string`

GetEksSubnetsZoneCIds returns the EksSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetEksSubnetsZoneCIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetEksSubnetsZoneCIdsOk() (*[]string, bool)`

GetEksSubnetsZoneCIdsOk returns a tuple with the EksSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) SetEksSubnetsZoneCIds(v []string)`

SetEksSubnetsZoneCIds sets EksSubnetsZoneCIds field to given value.


### GetDocumentdbSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) GetDocumentdbSubnetsZoneAIds() []string`

GetDocumentdbSubnetsZoneAIds returns the DocumentdbSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetDocumentdbSubnetsZoneAIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetDocumentdbSubnetsZoneAIdsOk() (*[]string, bool)`

GetDocumentdbSubnetsZoneAIdsOk returns a tuple with the DocumentdbSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentdbSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) SetDocumentdbSubnetsZoneAIds(v []string)`

SetDocumentdbSubnetsZoneAIds sets DocumentdbSubnetsZoneAIds field to given value.

### HasDocumentdbSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) HasDocumentdbSubnetsZoneAIds() bool`

HasDocumentdbSubnetsZoneAIds returns a boolean if a field has been set.

### GetDocumentdbSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) GetDocumentdbSubnetsZoneBIds() []string`

GetDocumentdbSubnetsZoneBIds returns the DocumentdbSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetDocumentdbSubnetsZoneBIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetDocumentdbSubnetsZoneBIdsOk() (*[]string, bool)`

GetDocumentdbSubnetsZoneBIdsOk returns a tuple with the DocumentdbSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentdbSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) SetDocumentdbSubnetsZoneBIds(v []string)`

SetDocumentdbSubnetsZoneBIds sets DocumentdbSubnetsZoneBIds field to given value.

### HasDocumentdbSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) HasDocumentdbSubnetsZoneBIds() bool`

HasDocumentdbSubnetsZoneBIds returns a boolean if a field has been set.

### GetDocumentdbSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) GetDocumentdbSubnetsZoneCIds() []string`

GetDocumentdbSubnetsZoneCIds returns the DocumentdbSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetDocumentdbSubnetsZoneCIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetDocumentdbSubnetsZoneCIdsOk() (*[]string, bool)`

GetDocumentdbSubnetsZoneCIdsOk returns a tuple with the DocumentdbSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentdbSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) SetDocumentdbSubnetsZoneCIds(v []string)`

SetDocumentdbSubnetsZoneCIds sets DocumentdbSubnetsZoneCIds field to given value.

### HasDocumentdbSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) HasDocumentdbSubnetsZoneCIds() bool`

HasDocumentdbSubnetsZoneCIds returns a boolean if a field has been set.

### GetElasticacheSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) GetElasticacheSubnetsZoneAIds() []string`

GetElasticacheSubnetsZoneAIds returns the ElasticacheSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetElasticacheSubnetsZoneAIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetElasticacheSubnetsZoneAIdsOk() (*[]string, bool)`

GetElasticacheSubnetsZoneAIdsOk returns a tuple with the ElasticacheSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticacheSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) SetElasticacheSubnetsZoneAIds(v []string)`

SetElasticacheSubnetsZoneAIds sets ElasticacheSubnetsZoneAIds field to given value.

### HasElasticacheSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) HasElasticacheSubnetsZoneAIds() bool`

HasElasticacheSubnetsZoneAIds returns a boolean if a field has been set.

### GetElasticacheSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) GetElasticacheSubnetsZoneBIds() []string`

GetElasticacheSubnetsZoneBIds returns the ElasticacheSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetElasticacheSubnetsZoneBIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetElasticacheSubnetsZoneBIdsOk() (*[]string, bool)`

GetElasticacheSubnetsZoneBIdsOk returns a tuple with the ElasticacheSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticacheSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) SetElasticacheSubnetsZoneBIds(v []string)`

SetElasticacheSubnetsZoneBIds sets ElasticacheSubnetsZoneBIds field to given value.

### HasElasticacheSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) HasElasticacheSubnetsZoneBIds() bool`

HasElasticacheSubnetsZoneBIds returns a boolean if a field has been set.

### GetElasticacheSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) GetElasticacheSubnetsZoneCIds() []string`

GetElasticacheSubnetsZoneCIds returns the ElasticacheSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetElasticacheSubnetsZoneCIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetElasticacheSubnetsZoneCIdsOk() (*[]string, bool)`

GetElasticacheSubnetsZoneCIdsOk returns a tuple with the ElasticacheSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticacheSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) SetElasticacheSubnetsZoneCIds(v []string)`

SetElasticacheSubnetsZoneCIds sets ElasticacheSubnetsZoneCIds field to given value.

### HasElasticacheSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) HasElasticacheSubnetsZoneCIds() bool`

HasElasticacheSubnetsZoneCIds returns a boolean if a field has been set.

### GetRdsSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) GetRdsSubnetsZoneAIds() []string`

GetRdsSubnetsZoneAIds returns the RdsSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetRdsSubnetsZoneAIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetRdsSubnetsZoneAIdsOk() (*[]string, bool)`

GetRdsSubnetsZoneAIdsOk returns a tuple with the RdsSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) SetRdsSubnetsZoneAIds(v []string)`

SetRdsSubnetsZoneAIds sets RdsSubnetsZoneAIds field to given value.

### HasRdsSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) HasRdsSubnetsZoneAIds() bool`

HasRdsSubnetsZoneAIds returns a boolean if a field has been set.

### GetRdsSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) GetRdsSubnetsZoneBIds() []string`

GetRdsSubnetsZoneBIds returns the RdsSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetRdsSubnetsZoneBIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetRdsSubnetsZoneBIdsOk() (*[]string, bool)`

GetRdsSubnetsZoneBIdsOk returns a tuple with the RdsSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) SetRdsSubnetsZoneBIds(v []string)`

SetRdsSubnetsZoneBIds sets RdsSubnetsZoneBIds field to given value.

### HasRdsSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) HasRdsSubnetsZoneBIds() bool`

HasRdsSubnetsZoneBIds returns a boolean if a field has been set.

### GetRdsSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) GetRdsSubnetsZoneCIds() []string`

GetRdsSubnetsZoneCIds returns the RdsSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetRdsSubnetsZoneCIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetRdsSubnetsZoneCIdsOk() (*[]string, bool)`

GetRdsSubnetsZoneCIdsOk returns a tuple with the RdsSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) SetRdsSubnetsZoneCIds(v []string)`

SetRdsSubnetsZoneCIds sets RdsSubnetsZoneCIds field to given value.

### HasRdsSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) HasRdsSubnetsZoneCIds() bool`

HasRdsSubnetsZoneCIds returns a boolean if a field has been set.

### GetEksKarpenterFargateSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) GetEksKarpenterFargateSubnetsZoneAIds() []string`

GetEksKarpenterFargateSubnetsZoneAIds returns the EksKarpenterFargateSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetEksKarpenterFargateSubnetsZoneAIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetEksKarpenterFargateSubnetsZoneAIdsOk() (*[]string, bool)`

GetEksKarpenterFargateSubnetsZoneAIdsOk returns a tuple with the EksKarpenterFargateSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksKarpenterFargateSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) SetEksKarpenterFargateSubnetsZoneAIds(v []string)`

SetEksKarpenterFargateSubnetsZoneAIds sets EksKarpenterFargateSubnetsZoneAIds field to given value.

### HasEksKarpenterFargateSubnetsZoneAIds

`func (o *ClusterFeatureAwsExistingVpc) HasEksKarpenterFargateSubnetsZoneAIds() bool`

HasEksKarpenterFargateSubnetsZoneAIds returns a boolean if a field has been set.

### GetEksKarpenterFargateSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) GetEksKarpenterFargateSubnetsZoneBIds() []string`

GetEksKarpenterFargateSubnetsZoneBIds returns the EksKarpenterFargateSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetEksKarpenterFargateSubnetsZoneBIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetEksKarpenterFargateSubnetsZoneBIdsOk() (*[]string, bool)`

GetEksKarpenterFargateSubnetsZoneBIdsOk returns a tuple with the EksKarpenterFargateSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksKarpenterFargateSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) SetEksKarpenterFargateSubnetsZoneBIds(v []string)`

SetEksKarpenterFargateSubnetsZoneBIds sets EksKarpenterFargateSubnetsZoneBIds field to given value.

### HasEksKarpenterFargateSubnetsZoneBIds

`func (o *ClusterFeatureAwsExistingVpc) HasEksKarpenterFargateSubnetsZoneBIds() bool`

HasEksKarpenterFargateSubnetsZoneBIds returns a boolean if a field has been set.

### GetEksKarpenterFargateSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) GetEksKarpenterFargateSubnetsZoneCIds() []string`

GetEksKarpenterFargateSubnetsZoneCIds returns the EksKarpenterFargateSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetEksKarpenterFargateSubnetsZoneCIdsOk

`func (o *ClusterFeatureAwsExistingVpc) GetEksKarpenterFargateSubnetsZoneCIdsOk() (*[]string, bool)`

GetEksKarpenterFargateSubnetsZoneCIdsOk returns a tuple with the EksKarpenterFargateSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksKarpenterFargateSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) SetEksKarpenterFargateSubnetsZoneCIds(v []string)`

SetEksKarpenterFargateSubnetsZoneCIds sets EksKarpenterFargateSubnetsZoneCIds field to given value.

### HasEksKarpenterFargateSubnetsZoneCIds

`func (o *ClusterFeatureAwsExistingVpc) HasEksKarpenterFargateSubnetsZoneCIds() bool`

HasEksKarpenterFargateSubnetsZoneCIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


