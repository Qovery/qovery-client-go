# ClusterInfrastructureEksAnywhereBackupS3Parameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** | S3 bucket name used to store EKS Anywhere backup artifacts. | 
**Region** | **string** | AWS region where the backup bucket is hosted. | 
**RoleArn** | **string** | IAM role ARN assumed to upload backup artifacts. | 
**KeyPrefix** | Pointer to **string** | Optional S3 key prefix used for backup object keys. | [optional] 

## Methods

### NewClusterInfrastructureEksAnywhereBackupS3Parameters

`func NewClusterInfrastructureEksAnywhereBackupS3Parameters(bucket string, region string, roleArn string, ) *ClusterInfrastructureEksAnywhereBackupS3Parameters`

NewClusterInfrastructureEksAnywhereBackupS3Parameters instantiates a new ClusterInfrastructureEksAnywhereBackupS3Parameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInfrastructureEksAnywhereBackupS3ParametersWithDefaults

`func NewClusterInfrastructureEksAnywhereBackupS3ParametersWithDefaults() *ClusterInfrastructureEksAnywhereBackupS3Parameters`

NewClusterInfrastructureEksAnywhereBackupS3ParametersWithDefaults instantiates a new ClusterInfrastructureEksAnywhereBackupS3Parameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *ClusterInfrastructureEksAnywhereBackupS3Parameters) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ClusterInfrastructureEksAnywhereBackupS3Parameters) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ClusterInfrastructureEksAnywhereBackupS3Parameters) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetRegion

`func (o *ClusterInfrastructureEksAnywhereBackupS3Parameters) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterInfrastructureEksAnywhereBackupS3Parameters) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterInfrastructureEksAnywhereBackupS3Parameters) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetRoleArn

`func (o *ClusterInfrastructureEksAnywhereBackupS3Parameters) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *ClusterInfrastructureEksAnywhereBackupS3Parameters) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *ClusterInfrastructureEksAnywhereBackupS3Parameters) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetKeyPrefix

`func (o *ClusterInfrastructureEksAnywhereBackupS3Parameters) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *ClusterInfrastructureEksAnywhereBackupS3Parameters) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *ClusterInfrastructureEksAnywhereBackupS3Parameters) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.

### HasKeyPrefix

`func (o *ClusterInfrastructureEksAnywhereBackupS3Parameters) HasKeyPrefix() bool`

HasKeyPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


