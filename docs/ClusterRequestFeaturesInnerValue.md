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

## Methods

### NewClusterRequestFeaturesInnerValue

`func NewClusterRequestFeaturesInnerValue(awsVpcEksId string, eksSubnetsZoneAIds []string, eksSubnetsZoneBIds []string, eksSubnetsZoneCIds []string, ) *ClusterRequestFeaturesInnerValue`

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

### SetDocumentdbSubnetsZoneAIdsNil

`func (o *ClusterRequestFeaturesInnerValue) SetDocumentdbSubnetsZoneAIdsNil(b bool)`

 SetDocumentdbSubnetsZoneAIdsNil sets the value for DocumentdbSubnetsZoneAIds to be an explicit nil

### UnsetDocumentdbSubnetsZoneAIds
`func (o *ClusterRequestFeaturesInnerValue) UnsetDocumentdbSubnetsZoneAIds()`

UnsetDocumentdbSubnetsZoneAIds ensures that no value is present for DocumentdbSubnetsZoneAIds, not even an explicit nil
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

### SetDocumentdbSubnetsZoneBIdsNil

`func (o *ClusterRequestFeaturesInnerValue) SetDocumentdbSubnetsZoneBIdsNil(b bool)`

 SetDocumentdbSubnetsZoneBIdsNil sets the value for DocumentdbSubnetsZoneBIds to be an explicit nil

### UnsetDocumentdbSubnetsZoneBIds
`func (o *ClusterRequestFeaturesInnerValue) UnsetDocumentdbSubnetsZoneBIds()`

UnsetDocumentdbSubnetsZoneBIds ensures that no value is present for DocumentdbSubnetsZoneBIds, not even an explicit nil
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

### SetDocumentdbSubnetsZoneCIdsNil

`func (o *ClusterRequestFeaturesInnerValue) SetDocumentdbSubnetsZoneCIdsNil(b bool)`

 SetDocumentdbSubnetsZoneCIdsNil sets the value for DocumentdbSubnetsZoneCIds to be an explicit nil

### UnsetDocumentdbSubnetsZoneCIds
`func (o *ClusterRequestFeaturesInnerValue) UnsetDocumentdbSubnetsZoneCIds()`

UnsetDocumentdbSubnetsZoneCIds ensures that no value is present for DocumentdbSubnetsZoneCIds, not even an explicit nil
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

### SetElasticacheSubnetsZoneAIdsNil

`func (o *ClusterRequestFeaturesInnerValue) SetElasticacheSubnetsZoneAIdsNil(b bool)`

 SetElasticacheSubnetsZoneAIdsNil sets the value for ElasticacheSubnetsZoneAIds to be an explicit nil

### UnsetElasticacheSubnetsZoneAIds
`func (o *ClusterRequestFeaturesInnerValue) UnsetElasticacheSubnetsZoneAIds()`

UnsetElasticacheSubnetsZoneAIds ensures that no value is present for ElasticacheSubnetsZoneAIds, not even an explicit nil
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

### SetElasticacheSubnetsZoneBIdsNil

`func (o *ClusterRequestFeaturesInnerValue) SetElasticacheSubnetsZoneBIdsNil(b bool)`

 SetElasticacheSubnetsZoneBIdsNil sets the value for ElasticacheSubnetsZoneBIds to be an explicit nil

### UnsetElasticacheSubnetsZoneBIds
`func (o *ClusterRequestFeaturesInnerValue) UnsetElasticacheSubnetsZoneBIds()`

UnsetElasticacheSubnetsZoneBIds ensures that no value is present for ElasticacheSubnetsZoneBIds, not even an explicit nil
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

### SetElasticacheSubnetsZoneCIdsNil

`func (o *ClusterRequestFeaturesInnerValue) SetElasticacheSubnetsZoneCIdsNil(b bool)`

 SetElasticacheSubnetsZoneCIdsNil sets the value for ElasticacheSubnetsZoneCIds to be an explicit nil

### UnsetElasticacheSubnetsZoneCIds
`func (o *ClusterRequestFeaturesInnerValue) UnsetElasticacheSubnetsZoneCIds()`

UnsetElasticacheSubnetsZoneCIds ensures that no value is present for ElasticacheSubnetsZoneCIds, not even an explicit nil
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

### SetRdsSubnetsZoneAIdsNil

`func (o *ClusterRequestFeaturesInnerValue) SetRdsSubnetsZoneAIdsNil(b bool)`

 SetRdsSubnetsZoneAIdsNil sets the value for RdsSubnetsZoneAIds to be an explicit nil

### UnsetRdsSubnetsZoneAIds
`func (o *ClusterRequestFeaturesInnerValue) UnsetRdsSubnetsZoneAIds()`

UnsetRdsSubnetsZoneAIds ensures that no value is present for RdsSubnetsZoneAIds, not even an explicit nil
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

### SetRdsSubnetsZoneBIdsNil

`func (o *ClusterRequestFeaturesInnerValue) SetRdsSubnetsZoneBIdsNil(b bool)`

 SetRdsSubnetsZoneBIdsNil sets the value for RdsSubnetsZoneBIds to be an explicit nil

### UnsetRdsSubnetsZoneBIds
`func (o *ClusterRequestFeaturesInnerValue) UnsetRdsSubnetsZoneBIds()`

UnsetRdsSubnetsZoneBIds ensures that no value is present for RdsSubnetsZoneBIds, not even an explicit nil
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

### SetRdsSubnetsZoneCIdsNil

`func (o *ClusterRequestFeaturesInnerValue) SetRdsSubnetsZoneCIdsNil(b bool)`

 SetRdsSubnetsZoneCIdsNil sets the value for RdsSubnetsZoneCIds to be an explicit nil

### UnsetRdsSubnetsZoneCIds
`func (o *ClusterRequestFeaturesInnerValue) UnsetRdsSubnetsZoneCIds()`

UnsetRdsSubnetsZoneCIds ensures that no value is present for RdsSubnetsZoneCIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


