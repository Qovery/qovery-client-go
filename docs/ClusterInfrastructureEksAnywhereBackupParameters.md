# ClusterInfrastructureEksAnywhereBackupParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable or disable EKS Anywhere cluster backup. | [optional] [default to true]
**TimeoutSeconds** | Pointer to **int64** | Timeout in seconds for backup operations. | [optional] [default to 300]
**CertsSecretName** | Pointer to **string** | Optional Kubernetes secret name holding etcd certificates. | [optional] 
**S3** | [**ClusterInfrastructureEksAnywhereBackupS3Parameters**](ClusterInfrastructureEksAnywhereBackupS3Parameters.md) |  | 

## Methods

### NewClusterInfrastructureEksAnywhereBackupParameters

`func NewClusterInfrastructureEksAnywhereBackupParameters(s3 ClusterInfrastructureEksAnywhereBackupS3Parameters, ) *ClusterInfrastructureEksAnywhereBackupParameters`

NewClusterInfrastructureEksAnywhereBackupParameters instantiates a new ClusterInfrastructureEksAnywhereBackupParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInfrastructureEksAnywhereBackupParametersWithDefaults

`func NewClusterInfrastructureEksAnywhereBackupParametersWithDefaults() *ClusterInfrastructureEksAnywhereBackupParameters`

NewClusterInfrastructureEksAnywhereBackupParametersWithDefaults instantiates a new ClusterInfrastructureEksAnywhereBackupParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetCertsSecretName

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) GetCertsSecretName() string`

GetCertsSecretName returns the CertsSecretName field if non-nil, zero value otherwise.

### GetCertsSecretNameOk

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) GetCertsSecretNameOk() (*string, bool)`

GetCertsSecretNameOk returns a tuple with the CertsSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertsSecretName

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) SetCertsSecretName(v string)`

SetCertsSecretName sets CertsSecretName field to given value.

### HasCertsSecretName

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) HasCertsSecretName() bool`

HasCertsSecretName returns a boolean if a field has been set.

### GetS3

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) GetS3() ClusterInfrastructureEksAnywhereBackupS3Parameters`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) GetS3Ok() (*ClusterInfrastructureEksAnywhereBackupS3Parameters, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *ClusterInfrastructureEksAnywhereBackupParameters) SetS3(v ClusterInfrastructureEksAnywhereBackupS3Parameters)`

SetS3 sets S3 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


