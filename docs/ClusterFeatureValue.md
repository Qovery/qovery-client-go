# ClusterFeatureValue

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

### NewClusterFeatureValue

`func NewClusterFeatureValue(awsVpcEksId string, eksSubnetsZoneAIds []string, eksSubnetsZoneBIds []string, eksSubnetsZoneCIds []string, ) *ClusterFeatureValue`

NewClusterFeatureValue instantiates a new ClusterFeatureValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFeatureValueWithDefaults

`func NewClusterFeatureValueWithDefaults() *ClusterFeatureValue`

NewClusterFeatureValueWithDefaults instantiates a new ClusterFeatureValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsVpcEksId

`func (o *ClusterFeatureValue) GetAwsVpcEksId() string`

GetAwsVpcEksId returns the AwsVpcEksId field if non-nil, zero value otherwise.

### GetAwsVpcEksIdOk

`func (o *ClusterFeatureValue) GetAwsVpcEksIdOk() (*string, bool)`

GetAwsVpcEksIdOk returns a tuple with the AwsVpcEksId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsVpcEksId

`func (o *ClusterFeatureValue) SetAwsVpcEksId(v string)`

SetAwsVpcEksId sets AwsVpcEksId field to given value.


### GetEksSubnetsZoneAIds

`func (o *ClusterFeatureValue) GetEksSubnetsZoneAIds() []string`

GetEksSubnetsZoneAIds returns the EksSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetEksSubnetsZoneAIdsOk

`func (o *ClusterFeatureValue) GetEksSubnetsZoneAIdsOk() (*[]string, bool)`

GetEksSubnetsZoneAIdsOk returns a tuple with the EksSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSubnetsZoneAIds

`func (o *ClusterFeatureValue) SetEksSubnetsZoneAIds(v []string)`

SetEksSubnetsZoneAIds sets EksSubnetsZoneAIds field to given value.


### GetEksSubnetsZoneBIds

`func (o *ClusterFeatureValue) GetEksSubnetsZoneBIds() []string`

GetEksSubnetsZoneBIds returns the EksSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetEksSubnetsZoneBIdsOk

`func (o *ClusterFeatureValue) GetEksSubnetsZoneBIdsOk() (*[]string, bool)`

GetEksSubnetsZoneBIdsOk returns a tuple with the EksSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSubnetsZoneBIds

`func (o *ClusterFeatureValue) SetEksSubnetsZoneBIds(v []string)`

SetEksSubnetsZoneBIds sets EksSubnetsZoneBIds field to given value.


### GetEksSubnetsZoneCIds

`func (o *ClusterFeatureValue) GetEksSubnetsZoneCIds() []string`

GetEksSubnetsZoneCIds returns the EksSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetEksSubnetsZoneCIdsOk

`func (o *ClusterFeatureValue) GetEksSubnetsZoneCIdsOk() (*[]string, bool)`

GetEksSubnetsZoneCIdsOk returns a tuple with the EksSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksSubnetsZoneCIds

`func (o *ClusterFeatureValue) SetEksSubnetsZoneCIds(v []string)`

SetEksSubnetsZoneCIds sets EksSubnetsZoneCIds field to given value.


### GetDocumentdbSubnetsZoneAIds

`func (o *ClusterFeatureValue) GetDocumentdbSubnetsZoneAIds() []string`

GetDocumentdbSubnetsZoneAIds returns the DocumentdbSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetDocumentdbSubnetsZoneAIdsOk

`func (o *ClusterFeatureValue) GetDocumentdbSubnetsZoneAIdsOk() (*[]string, bool)`

GetDocumentdbSubnetsZoneAIdsOk returns a tuple with the DocumentdbSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentdbSubnetsZoneAIds

`func (o *ClusterFeatureValue) SetDocumentdbSubnetsZoneAIds(v []string)`

SetDocumentdbSubnetsZoneAIds sets DocumentdbSubnetsZoneAIds field to given value.

### HasDocumentdbSubnetsZoneAIds

`func (o *ClusterFeatureValue) HasDocumentdbSubnetsZoneAIds() bool`

HasDocumentdbSubnetsZoneAIds returns a boolean if a field has been set.

### SetDocumentdbSubnetsZoneAIdsNil

`func (o *ClusterFeatureValue) SetDocumentdbSubnetsZoneAIdsNil(b bool)`

 SetDocumentdbSubnetsZoneAIdsNil sets the value for DocumentdbSubnetsZoneAIds to be an explicit nil

### UnsetDocumentdbSubnetsZoneAIds
`func (o *ClusterFeatureValue) UnsetDocumentdbSubnetsZoneAIds()`

UnsetDocumentdbSubnetsZoneAIds ensures that no value is present for DocumentdbSubnetsZoneAIds, not even an explicit nil
### GetDocumentdbSubnetsZoneBIds

`func (o *ClusterFeatureValue) GetDocumentdbSubnetsZoneBIds() []string`

GetDocumentdbSubnetsZoneBIds returns the DocumentdbSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetDocumentdbSubnetsZoneBIdsOk

`func (o *ClusterFeatureValue) GetDocumentdbSubnetsZoneBIdsOk() (*[]string, bool)`

GetDocumentdbSubnetsZoneBIdsOk returns a tuple with the DocumentdbSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentdbSubnetsZoneBIds

`func (o *ClusterFeatureValue) SetDocumentdbSubnetsZoneBIds(v []string)`

SetDocumentdbSubnetsZoneBIds sets DocumentdbSubnetsZoneBIds field to given value.

### HasDocumentdbSubnetsZoneBIds

`func (o *ClusterFeatureValue) HasDocumentdbSubnetsZoneBIds() bool`

HasDocumentdbSubnetsZoneBIds returns a boolean if a field has been set.

### SetDocumentdbSubnetsZoneBIdsNil

`func (o *ClusterFeatureValue) SetDocumentdbSubnetsZoneBIdsNil(b bool)`

 SetDocumentdbSubnetsZoneBIdsNil sets the value for DocumentdbSubnetsZoneBIds to be an explicit nil

### UnsetDocumentdbSubnetsZoneBIds
`func (o *ClusterFeatureValue) UnsetDocumentdbSubnetsZoneBIds()`

UnsetDocumentdbSubnetsZoneBIds ensures that no value is present for DocumentdbSubnetsZoneBIds, not even an explicit nil
### GetDocumentdbSubnetsZoneCIds

`func (o *ClusterFeatureValue) GetDocumentdbSubnetsZoneCIds() []string`

GetDocumentdbSubnetsZoneCIds returns the DocumentdbSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetDocumentdbSubnetsZoneCIdsOk

`func (o *ClusterFeatureValue) GetDocumentdbSubnetsZoneCIdsOk() (*[]string, bool)`

GetDocumentdbSubnetsZoneCIdsOk returns a tuple with the DocumentdbSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentdbSubnetsZoneCIds

`func (o *ClusterFeatureValue) SetDocumentdbSubnetsZoneCIds(v []string)`

SetDocumentdbSubnetsZoneCIds sets DocumentdbSubnetsZoneCIds field to given value.

### HasDocumentdbSubnetsZoneCIds

`func (o *ClusterFeatureValue) HasDocumentdbSubnetsZoneCIds() bool`

HasDocumentdbSubnetsZoneCIds returns a boolean if a field has been set.

### SetDocumentdbSubnetsZoneCIdsNil

`func (o *ClusterFeatureValue) SetDocumentdbSubnetsZoneCIdsNil(b bool)`

 SetDocumentdbSubnetsZoneCIdsNil sets the value for DocumentdbSubnetsZoneCIds to be an explicit nil

### UnsetDocumentdbSubnetsZoneCIds
`func (o *ClusterFeatureValue) UnsetDocumentdbSubnetsZoneCIds()`

UnsetDocumentdbSubnetsZoneCIds ensures that no value is present for DocumentdbSubnetsZoneCIds, not even an explicit nil
### GetElasticacheSubnetsZoneAIds

`func (o *ClusterFeatureValue) GetElasticacheSubnetsZoneAIds() []string`

GetElasticacheSubnetsZoneAIds returns the ElasticacheSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetElasticacheSubnetsZoneAIdsOk

`func (o *ClusterFeatureValue) GetElasticacheSubnetsZoneAIdsOk() (*[]string, bool)`

GetElasticacheSubnetsZoneAIdsOk returns a tuple with the ElasticacheSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticacheSubnetsZoneAIds

`func (o *ClusterFeatureValue) SetElasticacheSubnetsZoneAIds(v []string)`

SetElasticacheSubnetsZoneAIds sets ElasticacheSubnetsZoneAIds field to given value.

### HasElasticacheSubnetsZoneAIds

`func (o *ClusterFeatureValue) HasElasticacheSubnetsZoneAIds() bool`

HasElasticacheSubnetsZoneAIds returns a boolean if a field has been set.

### SetElasticacheSubnetsZoneAIdsNil

`func (o *ClusterFeatureValue) SetElasticacheSubnetsZoneAIdsNil(b bool)`

 SetElasticacheSubnetsZoneAIdsNil sets the value for ElasticacheSubnetsZoneAIds to be an explicit nil

### UnsetElasticacheSubnetsZoneAIds
`func (o *ClusterFeatureValue) UnsetElasticacheSubnetsZoneAIds()`

UnsetElasticacheSubnetsZoneAIds ensures that no value is present for ElasticacheSubnetsZoneAIds, not even an explicit nil
### GetElasticacheSubnetsZoneBIds

`func (o *ClusterFeatureValue) GetElasticacheSubnetsZoneBIds() []string`

GetElasticacheSubnetsZoneBIds returns the ElasticacheSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetElasticacheSubnetsZoneBIdsOk

`func (o *ClusterFeatureValue) GetElasticacheSubnetsZoneBIdsOk() (*[]string, bool)`

GetElasticacheSubnetsZoneBIdsOk returns a tuple with the ElasticacheSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticacheSubnetsZoneBIds

`func (o *ClusterFeatureValue) SetElasticacheSubnetsZoneBIds(v []string)`

SetElasticacheSubnetsZoneBIds sets ElasticacheSubnetsZoneBIds field to given value.

### HasElasticacheSubnetsZoneBIds

`func (o *ClusterFeatureValue) HasElasticacheSubnetsZoneBIds() bool`

HasElasticacheSubnetsZoneBIds returns a boolean if a field has been set.

### SetElasticacheSubnetsZoneBIdsNil

`func (o *ClusterFeatureValue) SetElasticacheSubnetsZoneBIdsNil(b bool)`

 SetElasticacheSubnetsZoneBIdsNil sets the value for ElasticacheSubnetsZoneBIds to be an explicit nil

### UnsetElasticacheSubnetsZoneBIds
`func (o *ClusterFeatureValue) UnsetElasticacheSubnetsZoneBIds()`

UnsetElasticacheSubnetsZoneBIds ensures that no value is present for ElasticacheSubnetsZoneBIds, not even an explicit nil
### GetElasticacheSubnetsZoneCIds

`func (o *ClusterFeatureValue) GetElasticacheSubnetsZoneCIds() []string`

GetElasticacheSubnetsZoneCIds returns the ElasticacheSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetElasticacheSubnetsZoneCIdsOk

`func (o *ClusterFeatureValue) GetElasticacheSubnetsZoneCIdsOk() (*[]string, bool)`

GetElasticacheSubnetsZoneCIdsOk returns a tuple with the ElasticacheSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticacheSubnetsZoneCIds

`func (o *ClusterFeatureValue) SetElasticacheSubnetsZoneCIds(v []string)`

SetElasticacheSubnetsZoneCIds sets ElasticacheSubnetsZoneCIds field to given value.

### HasElasticacheSubnetsZoneCIds

`func (o *ClusterFeatureValue) HasElasticacheSubnetsZoneCIds() bool`

HasElasticacheSubnetsZoneCIds returns a boolean if a field has been set.

### SetElasticacheSubnetsZoneCIdsNil

`func (o *ClusterFeatureValue) SetElasticacheSubnetsZoneCIdsNil(b bool)`

 SetElasticacheSubnetsZoneCIdsNil sets the value for ElasticacheSubnetsZoneCIds to be an explicit nil

### UnsetElasticacheSubnetsZoneCIds
`func (o *ClusterFeatureValue) UnsetElasticacheSubnetsZoneCIds()`

UnsetElasticacheSubnetsZoneCIds ensures that no value is present for ElasticacheSubnetsZoneCIds, not even an explicit nil
### GetRdsSubnetsZoneAIds

`func (o *ClusterFeatureValue) GetRdsSubnetsZoneAIds() []string`

GetRdsSubnetsZoneAIds returns the RdsSubnetsZoneAIds field if non-nil, zero value otherwise.

### GetRdsSubnetsZoneAIdsOk

`func (o *ClusterFeatureValue) GetRdsSubnetsZoneAIdsOk() (*[]string, bool)`

GetRdsSubnetsZoneAIdsOk returns a tuple with the RdsSubnetsZoneAIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsSubnetsZoneAIds

`func (o *ClusterFeatureValue) SetRdsSubnetsZoneAIds(v []string)`

SetRdsSubnetsZoneAIds sets RdsSubnetsZoneAIds field to given value.

### HasRdsSubnetsZoneAIds

`func (o *ClusterFeatureValue) HasRdsSubnetsZoneAIds() bool`

HasRdsSubnetsZoneAIds returns a boolean if a field has been set.

### SetRdsSubnetsZoneAIdsNil

`func (o *ClusterFeatureValue) SetRdsSubnetsZoneAIdsNil(b bool)`

 SetRdsSubnetsZoneAIdsNil sets the value for RdsSubnetsZoneAIds to be an explicit nil

### UnsetRdsSubnetsZoneAIds
`func (o *ClusterFeatureValue) UnsetRdsSubnetsZoneAIds()`

UnsetRdsSubnetsZoneAIds ensures that no value is present for RdsSubnetsZoneAIds, not even an explicit nil
### GetRdsSubnetsZoneBIds

`func (o *ClusterFeatureValue) GetRdsSubnetsZoneBIds() []string`

GetRdsSubnetsZoneBIds returns the RdsSubnetsZoneBIds field if non-nil, zero value otherwise.

### GetRdsSubnetsZoneBIdsOk

`func (o *ClusterFeatureValue) GetRdsSubnetsZoneBIdsOk() (*[]string, bool)`

GetRdsSubnetsZoneBIdsOk returns a tuple with the RdsSubnetsZoneBIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsSubnetsZoneBIds

`func (o *ClusterFeatureValue) SetRdsSubnetsZoneBIds(v []string)`

SetRdsSubnetsZoneBIds sets RdsSubnetsZoneBIds field to given value.

### HasRdsSubnetsZoneBIds

`func (o *ClusterFeatureValue) HasRdsSubnetsZoneBIds() bool`

HasRdsSubnetsZoneBIds returns a boolean if a field has been set.

### SetRdsSubnetsZoneBIdsNil

`func (o *ClusterFeatureValue) SetRdsSubnetsZoneBIdsNil(b bool)`

 SetRdsSubnetsZoneBIdsNil sets the value for RdsSubnetsZoneBIds to be an explicit nil

### UnsetRdsSubnetsZoneBIds
`func (o *ClusterFeatureValue) UnsetRdsSubnetsZoneBIds()`

UnsetRdsSubnetsZoneBIds ensures that no value is present for RdsSubnetsZoneBIds, not even an explicit nil
### GetRdsSubnetsZoneCIds

`func (o *ClusterFeatureValue) GetRdsSubnetsZoneCIds() []string`

GetRdsSubnetsZoneCIds returns the RdsSubnetsZoneCIds field if non-nil, zero value otherwise.

### GetRdsSubnetsZoneCIdsOk

`func (o *ClusterFeatureValue) GetRdsSubnetsZoneCIdsOk() (*[]string, bool)`

GetRdsSubnetsZoneCIdsOk returns a tuple with the RdsSubnetsZoneCIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsSubnetsZoneCIds

`func (o *ClusterFeatureValue) SetRdsSubnetsZoneCIds(v []string)`

SetRdsSubnetsZoneCIds sets RdsSubnetsZoneCIds field to given value.

### HasRdsSubnetsZoneCIds

`func (o *ClusterFeatureValue) HasRdsSubnetsZoneCIds() bool`

HasRdsSubnetsZoneCIds returns a boolean if a field has been set.

### SetRdsSubnetsZoneCIdsNil

`func (o *ClusterFeatureValue) SetRdsSubnetsZoneCIdsNil(b bool)`

 SetRdsSubnetsZoneCIdsNil sets the value for RdsSubnetsZoneCIds to be an explicit nil

### UnsetRdsSubnetsZoneCIds
`func (o *ClusterFeatureValue) UnsetRdsSubnetsZoneCIds()`

UnsetRdsSubnetsZoneCIds ensures that no value is present for RdsSubnetsZoneCIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


