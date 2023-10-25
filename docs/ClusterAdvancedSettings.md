# ClusterAdvancedSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsCloudwatchEksLogsRetentionDays** | Pointer to **int32** | Set the number of retention days for EKS Cloudwatch logs | [optional] [default to 90]
**AwsVpcEnableS3FlowLogs** | Pointer to **bool** | Enable flow logs for on the VPC and store them in an S3 bucket | [optional] [default to false]
**AwsVpcFlowLogsRetentionDays** | Pointer to **int32** | Set the number of retention days for flow logs. Disable with value \&quot;0\&quot; | [optional] [default to 365]
**LokiLogRetentionInWeek** | Pointer to **int32** | For how long in week loki is going to keep logs of your applications | [optional] [default to 12]
**RegistryImageRetentionTime** | Pointer to **int32** | Configure the number of seconds before cleaning images in the registry | [optional] [default to 31536000]
**CloudProviderContainerRegistryTags** | Pointer to **map[string]string** | Add additional tags on the cluster dedicated registry | [optional] 
**LoadBalancerSize** | Pointer to **string** | Select the size of the main load_balancer (only effective for Scaleway) | [optional] [default to "lb-s"]
**DatabasePostgresqlDenyPublicAccess** | Pointer to **bool** | Deny public access to any PostgreSQL database | [optional] [default to false]
**DatabasePostgresqlAllowedCidrs** | Pointer to **[]string** | List of CIDRs allowed to access the PostgreSQL database | [optional] [default to ["0.0.0.0/0"]]
**DatabaseMysqlDenyPublicAccess** | Pointer to **bool** | Deny public access to any MySql database | [optional] [default to false]
**DatabaseMysqlAllowedCidrs** | Pointer to **[]string** | List of CIDRs allowed to access the MySql database | [optional] [default to ["0.0.0.0/0"]]
**DatabaseMongodbDenyPublicAccess** | Pointer to **bool** | Deny public access to any MongoDB/DocumentDB database | [optional] [default to false]
**DatabaseMongodbAllowedCidrs** | Pointer to **[]string** | List of CIDRs allowed to access the MongoDB/DocumentDB database | [optional] [default to ["0.0.0.0/0"]]
**DatabaseRedisDenyPublicAccess** | Pointer to **bool** | Deny public access to any Redis database | [optional] [default to false]
**DatabaseRedisAllowedCidrs** | Pointer to **[]string** | List of CIDRs allowed to access the Redis database | [optional] [default to ["0.0.0.0/0"]]
**AwsIamAdminGroup** | Pointer to **string** | AWS IAM group name with cluster access | [optional] [default to "Admins"]
**AwsEksEc2MetadataImds** | Pointer to **string** | Specify the [IMDS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) version you want to use:   * &#x60;required&#x60;: IMDS V2 only   * &#x60;optional&#x60;: IMDS V1 + V2  | [optional] [default to "optional"]
**PlecoResourcesTtl** | Pointer to **int32** |  | [optional] [default to -1]
**RegistryMirroringMode** | Pointer to [**RegistryMirroringModeEnum**](RegistryMirroringModeEnum.md) |  | [optional] [default to REGISTRYMIRRORINGMODEENUM_SERVICE]
**NginxVcpuRequestInMilli** | Pointer to **int32** | vcpu request in millicores | [optional] [default to 100]
**NginxVcpuLimitInMilli** | Pointer to **int32** | vcpu limit in millicores | [optional] [default to 500]
**NginxMemoryRequestInMib** | Pointer to **int32** | memory request in MiB | [optional] [default to 768]
**NginxMemoryLimitInMib** | Pointer to **int32** | memory limit in MiB | [optional] [default to 768]
**NginxHpaCpuUtilizationPercentageThreshold** | Pointer to **int32** | hpa cpu threshold in percentage | [optional] [default to 50]
**NginxHpaMemoryUtilizationPercentageThreshold** | Pointer to **int32** | hpa memory threshold in percentage | [optional] [default to 50]
**NginxHpaMinNumberInstances** | Pointer to **int32** | hpa minimum number of instances | [optional] [default to 2]
**NginxHpaMaxNumberInstances** | Pointer to **int32** | hpa maximum number of instances | [optional] [default to 25]

## Methods

### NewClusterAdvancedSettings

`func NewClusterAdvancedSettings() *ClusterAdvancedSettings`

NewClusterAdvancedSettings instantiates a new ClusterAdvancedSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAdvancedSettingsWithDefaults

`func NewClusterAdvancedSettingsWithDefaults() *ClusterAdvancedSettings`

NewClusterAdvancedSettingsWithDefaults instantiates a new ClusterAdvancedSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsCloudwatchEksLogsRetentionDays

`func (o *ClusterAdvancedSettings) GetAwsCloudwatchEksLogsRetentionDays() int32`

GetAwsCloudwatchEksLogsRetentionDays returns the AwsCloudwatchEksLogsRetentionDays field if non-nil, zero value otherwise.

### GetAwsCloudwatchEksLogsRetentionDaysOk

`func (o *ClusterAdvancedSettings) GetAwsCloudwatchEksLogsRetentionDaysOk() (*int32, bool)`

GetAwsCloudwatchEksLogsRetentionDaysOk returns a tuple with the AwsCloudwatchEksLogsRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudwatchEksLogsRetentionDays

`func (o *ClusterAdvancedSettings) SetAwsCloudwatchEksLogsRetentionDays(v int32)`

SetAwsCloudwatchEksLogsRetentionDays sets AwsCloudwatchEksLogsRetentionDays field to given value.

### HasAwsCloudwatchEksLogsRetentionDays

`func (o *ClusterAdvancedSettings) HasAwsCloudwatchEksLogsRetentionDays() bool`

HasAwsCloudwatchEksLogsRetentionDays returns a boolean if a field has been set.

### GetAwsVpcEnableS3FlowLogs

`func (o *ClusterAdvancedSettings) GetAwsVpcEnableS3FlowLogs() bool`

GetAwsVpcEnableS3FlowLogs returns the AwsVpcEnableS3FlowLogs field if non-nil, zero value otherwise.

### GetAwsVpcEnableS3FlowLogsOk

`func (o *ClusterAdvancedSettings) GetAwsVpcEnableS3FlowLogsOk() (*bool, bool)`

GetAwsVpcEnableS3FlowLogsOk returns a tuple with the AwsVpcEnableS3FlowLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsVpcEnableS3FlowLogs

`func (o *ClusterAdvancedSettings) SetAwsVpcEnableS3FlowLogs(v bool)`

SetAwsVpcEnableS3FlowLogs sets AwsVpcEnableS3FlowLogs field to given value.

### HasAwsVpcEnableS3FlowLogs

`func (o *ClusterAdvancedSettings) HasAwsVpcEnableS3FlowLogs() bool`

HasAwsVpcEnableS3FlowLogs returns a boolean if a field has been set.

### GetAwsVpcFlowLogsRetentionDays

`func (o *ClusterAdvancedSettings) GetAwsVpcFlowLogsRetentionDays() int32`

GetAwsVpcFlowLogsRetentionDays returns the AwsVpcFlowLogsRetentionDays field if non-nil, zero value otherwise.

### GetAwsVpcFlowLogsRetentionDaysOk

`func (o *ClusterAdvancedSettings) GetAwsVpcFlowLogsRetentionDaysOk() (*int32, bool)`

GetAwsVpcFlowLogsRetentionDaysOk returns a tuple with the AwsVpcFlowLogsRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsVpcFlowLogsRetentionDays

`func (o *ClusterAdvancedSettings) SetAwsVpcFlowLogsRetentionDays(v int32)`

SetAwsVpcFlowLogsRetentionDays sets AwsVpcFlowLogsRetentionDays field to given value.

### HasAwsVpcFlowLogsRetentionDays

`func (o *ClusterAdvancedSettings) HasAwsVpcFlowLogsRetentionDays() bool`

HasAwsVpcFlowLogsRetentionDays returns a boolean if a field has been set.

### GetLokiLogRetentionInWeek

`func (o *ClusterAdvancedSettings) GetLokiLogRetentionInWeek() int32`

GetLokiLogRetentionInWeek returns the LokiLogRetentionInWeek field if non-nil, zero value otherwise.

### GetLokiLogRetentionInWeekOk

`func (o *ClusterAdvancedSettings) GetLokiLogRetentionInWeekOk() (*int32, bool)`

GetLokiLogRetentionInWeekOk returns a tuple with the LokiLogRetentionInWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLokiLogRetentionInWeek

`func (o *ClusterAdvancedSettings) SetLokiLogRetentionInWeek(v int32)`

SetLokiLogRetentionInWeek sets LokiLogRetentionInWeek field to given value.

### HasLokiLogRetentionInWeek

`func (o *ClusterAdvancedSettings) HasLokiLogRetentionInWeek() bool`

HasLokiLogRetentionInWeek returns a boolean if a field has been set.

### GetRegistryImageRetentionTime

`func (o *ClusterAdvancedSettings) GetRegistryImageRetentionTime() int32`

GetRegistryImageRetentionTime returns the RegistryImageRetentionTime field if non-nil, zero value otherwise.

### GetRegistryImageRetentionTimeOk

`func (o *ClusterAdvancedSettings) GetRegistryImageRetentionTimeOk() (*int32, bool)`

GetRegistryImageRetentionTimeOk returns a tuple with the RegistryImageRetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryImageRetentionTime

`func (o *ClusterAdvancedSettings) SetRegistryImageRetentionTime(v int32)`

SetRegistryImageRetentionTime sets RegistryImageRetentionTime field to given value.

### HasRegistryImageRetentionTime

`func (o *ClusterAdvancedSettings) HasRegistryImageRetentionTime() bool`

HasRegistryImageRetentionTime returns a boolean if a field has been set.

### GetCloudProviderContainerRegistryTags

`func (o *ClusterAdvancedSettings) GetCloudProviderContainerRegistryTags() map[string]string`

GetCloudProviderContainerRegistryTags returns the CloudProviderContainerRegistryTags field if non-nil, zero value otherwise.

### GetCloudProviderContainerRegistryTagsOk

`func (o *ClusterAdvancedSettings) GetCloudProviderContainerRegistryTagsOk() (*map[string]string, bool)`

GetCloudProviderContainerRegistryTagsOk returns a tuple with the CloudProviderContainerRegistryTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderContainerRegistryTags

`func (o *ClusterAdvancedSettings) SetCloudProviderContainerRegistryTags(v map[string]string)`

SetCloudProviderContainerRegistryTags sets CloudProviderContainerRegistryTags field to given value.

### HasCloudProviderContainerRegistryTags

`func (o *ClusterAdvancedSettings) HasCloudProviderContainerRegistryTags() bool`

HasCloudProviderContainerRegistryTags returns a boolean if a field has been set.

### GetLoadBalancerSize

`func (o *ClusterAdvancedSettings) GetLoadBalancerSize() string`

GetLoadBalancerSize returns the LoadBalancerSize field if non-nil, zero value otherwise.

### GetLoadBalancerSizeOk

`func (o *ClusterAdvancedSettings) GetLoadBalancerSizeOk() (*string, bool)`

GetLoadBalancerSizeOk returns a tuple with the LoadBalancerSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSize

`func (o *ClusterAdvancedSettings) SetLoadBalancerSize(v string)`

SetLoadBalancerSize sets LoadBalancerSize field to given value.

### HasLoadBalancerSize

`func (o *ClusterAdvancedSettings) HasLoadBalancerSize() bool`

HasLoadBalancerSize returns a boolean if a field has been set.

### GetDatabasePostgresqlDenyPublicAccess

`func (o *ClusterAdvancedSettings) GetDatabasePostgresqlDenyPublicAccess() bool`

GetDatabasePostgresqlDenyPublicAccess returns the DatabasePostgresqlDenyPublicAccess field if non-nil, zero value otherwise.

### GetDatabasePostgresqlDenyPublicAccessOk

`func (o *ClusterAdvancedSettings) GetDatabasePostgresqlDenyPublicAccessOk() (*bool, bool)`

GetDatabasePostgresqlDenyPublicAccessOk returns a tuple with the DatabasePostgresqlDenyPublicAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePostgresqlDenyPublicAccess

`func (o *ClusterAdvancedSettings) SetDatabasePostgresqlDenyPublicAccess(v bool)`

SetDatabasePostgresqlDenyPublicAccess sets DatabasePostgresqlDenyPublicAccess field to given value.

### HasDatabasePostgresqlDenyPublicAccess

`func (o *ClusterAdvancedSettings) HasDatabasePostgresqlDenyPublicAccess() bool`

HasDatabasePostgresqlDenyPublicAccess returns a boolean if a field has been set.

### GetDatabasePostgresqlAllowedCidrs

`func (o *ClusterAdvancedSettings) GetDatabasePostgresqlAllowedCidrs() []string`

GetDatabasePostgresqlAllowedCidrs returns the DatabasePostgresqlAllowedCidrs field if non-nil, zero value otherwise.

### GetDatabasePostgresqlAllowedCidrsOk

`func (o *ClusterAdvancedSettings) GetDatabasePostgresqlAllowedCidrsOk() (*[]string, bool)`

GetDatabasePostgresqlAllowedCidrsOk returns a tuple with the DatabasePostgresqlAllowedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePostgresqlAllowedCidrs

`func (o *ClusterAdvancedSettings) SetDatabasePostgresqlAllowedCidrs(v []string)`

SetDatabasePostgresqlAllowedCidrs sets DatabasePostgresqlAllowedCidrs field to given value.

### HasDatabasePostgresqlAllowedCidrs

`func (o *ClusterAdvancedSettings) HasDatabasePostgresqlAllowedCidrs() bool`

HasDatabasePostgresqlAllowedCidrs returns a boolean if a field has been set.

### GetDatabaseMysqlDenyPublicAccess

`func (o *ClusterAdvancedSettings) GetDatabaseMysqlDenyPublicAccess() bool`

GetDatabaseMysqlDenyPublicAccess returns the DatabaseMysqlDenyPublicAccess field if non-nil, zero value otherwise.

### GetDatabaseMysqlDenyPublicAccessOk

`func (o *ClusterAdvancedSettings) GetDatabaseMysqlDenyPublicAccessOk() (*bool, bool)`

GetDatabaseMysqlDenyPublicAccessOk returns a tuple with the DatabaseMysqlDenyPublicAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseMysqlDenyPublicAccess

`func (o *ClusterAdvancedSettings) SetDatabaseMysqlDenyPublicAccess(v bool)`

SetDatabaseMysqlDenyPublicAccess sets DatabaseMysqlDenyPublicAccess field to given value.

### HasDatabaseMysqlDenyPublicAccess

`func (o *ClusterAdvancedSettings) HasDatabaseMysqlDenyPublicAccess() bool`

HasDatabaseMysqlDenyPublicAccess returns a boolean if a field has been set.

### GetDatabaseMysqlAllowedCidrs

`func (o *ClusterAdvancedSettings) GetDatabaseMysqlAllowedCidrs() []string`

GetDatabaseMysqlAllowedCidrs returns the DatabaseMysqlAllowedCidrs field if non-nil, zero value otherwise.

### GetDatabaseMysqlAllowedCidrsOk

`func (o *ClusterAdvancedSettings) GetDatabaseMysqlAllowedCidrsOk() (*[]string, bool)`

GetDatabaseMysqlAllowedCidrsOk returns a tuple with the DatabaseMysqlAllowedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseMysqlAllowedCidrs

`func (o *ClusterAdvancedSettings) SetDatabaseMysqlAllowedCidrs(v []string)`

SetDatabaseMysqlAllowedCidrs sets DatabaseMysqlAllowedCidrs field to given value.

### HasDatabaseMysqlAllowedCidrs

`func (o *ClusterAdvancedSettings) HasDatabaseMysqlAllowedCidrs() bool`

HasDatabaseMysqlAllowedCidrs returns a boolean if a field has been set.

### GetDatabaseMongodbDenyPublicAccess

`func (o *ClusterAdvancedSettings) GetDatabaseMongodbDenyPublicAccess() bool`

GetDatabaseMongodbDenyPublicAccess returns the DatabaseMongodbDenyPublicAccess field if non-nil, zero value otherwise.

### GetDatabaseMongodbDenyPublicAccessOk

`func (o *ClusterAdvancedSettings) GetDatabaseMongodbDenyPublicAccessOk() (*bool, bool)`

GetDatabaseMongodbDenyPublicAccessOk returns a tuple with the DatabaseMongodbDenyPublicAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseMongodbDenyPublicAccess

`func (o *ClusterAdvancedSettings) SetDatabaseMongodbDenyPublicAccess(v bool)`

SetDatabaseMongodbDenyPublicAccess sets DatabaseMongodbDenyPublicAccess field to given value.

### HasDatabaseMongodbDenyPublicAccess

`func (o *ClusterAdvancedSettings) HasDatabaseMongodbDenyPublicAccess() bool`

HasDatabaseMongodbDenyPublicAccess returns a boolean if a field has been set.

### GetDatabaseMongodbAllowedCidrs

`func (o *ClusterAdvancedSettings) GetDatabaseMongodbAllowedCidrs() []string`

GetDatabaseMongodbAllowedCidrs returns the DatabaseMongodbAllowedCidrs field if non-nil, zero value otherwise.

### GetDatabaseMongodbAllowedCidrsOk

`func (o *ClusterAdvancedSettings) GetDatabaseMongodbAllowedCidrsOk() (*[]string, bool)`

GetDatabaseMongodbAllowedCidrsOk returns a tuple with the DatabaseMongodbAllowedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseMongodbAllowedCidrs

`func (o *ClusterAdvancedSettings) SetDatabaseMongodbAllowedCidrs(v []string)`

SetDatabaseMongodbAllowedCidrs sets DatabaseMongodbAllowedCidrs field to given value.

### HasDatabaseMongodbAllowedCidrs

`func (o *ClusterAdvancedSettings) HasDatabaseMongodbAllowedCidrs() bool`

HasDatabaseMongodbAllowedCidrs returns a boolean if a field has been set.

### GetDatabaseRedisDenyPublicAccess

`func (o *ClusterAdvancedSettings) GetDatabaseRedisDenyPublicAccess() bool`

GetDatabaseRedisDenyPublicAccess returns the DatabaseRedisDenyPublicAccess field if non-nil, zero value otherwise.

### GetDatabaseRedisDenyPublicAccessOk

`func (o *ClusterAdvancedSettings) GetDatabaseRedisDenyPublicAccessOk() (*bool, bool)`

GetDatabaseRedisDenyPublicAccessOk returns a tuple with the DatabaseRedisDenyPublicAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseRedisDenyPublicAccess

`func (o *ClusterAdvancedSettings) SetDatabaseRedisDenyPublicAccess(v bool)`

SetDatabaseRedisDenyPublicAccess sets DatabaseRedisDenyPublicAccess field to given value.

### HasDatabaseRedisDenyPublicAccess

`func (o *ClusterAdvancedSettings) HasDatabaseRedisDenyPublicAccess() bool`

HasDatabaseRedisDenyPublicAccess returns a boolean if a field has been set.

### GetDatabaseRedisAllowedCidrs

`func (o *ClusterAdvancedSettings) GetDatabaseRedisAllowedCidrs() []string`

GetDatabaseRedisAllowedCidrs returns the DatabaseRedisAllowedCidrs field if non-nil, zero value otherwise.

### GetDatabaseRedisAllowedCidrsOk

`func (o *ClusterAdvancedSettings) GetDatabaseRedisAllowedCidrsOk() (*[]string, bool)`

GetDatabaseRedisAllowedCidrsOk returns a tuple with the DatabaseRedisAllowedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseRedisAllowedCidrs

`func (o *ClusterAdvancedSettings) SetDatabaseRedisAllowedCidrs(v []string)`

SetDatabaseRedisAllowedCidrs sets DatabaseRedisAllowedCidrs field to given value.

### HasDatabaseRedisAllowedCidrs

`func (o *ClusterAdvancedSettings) HasDatabaseRedisAllowedCidrs() bool`

HasDatabaseRedisAllowedCidrs returns a boolean if a field has been set.

### GetAwsIamAdminGroup

`func (o *ClusterAdvancedSettings) GetAwsIamAdminGroup() string`

GetAwsIamAdminGroup returns the AwsIamAdminGroup field if non-nil, zero value otherwise.

### GetAwsIamAdminGroupOk

`func (o *ClusterAdvancedSettings) GetAwsIamAdminGroupOk() (*string, bool)`

GetAwsIamAdminGroupOk returns a tuple with the AwsIamAdminGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsIamAdminGroup

`func (o *ClusterAdvancedSettings) SetAwsIamAdminGroup(v string)`

SetAwsIamAdminGroup sets AwsIamAdminGroup field to given value.

### HasAwsIamAdminGroup

`func (o *ClusterAdvancedSettings) HasAwsIamAdminGroup() bool`

HasAwsIamAdminGroup returns a boolean if a field has been set.

### GetAwsEksEc2MetadataImds

`func (o *ClusterAdvancedSettings) GetAwsEksEc2MetadataImds() string`

GetAwsEksEc2MetadataImds returns the AwsEksEc2MetadataImds field if non-nil, zero value otherwise.

### GetAwsEksEc2MetadataImdsOk

`func (o *ClusterAdvancedSettings) GetAwsEksEc2MetadataImdsOk() (*string, bool)`

GetAwsEksEc2MetadataImdsOk returns a tuple with the AwsEksEc2MetadataImds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEksEc2MetadataImds

`func (o *ClusterAdvancedSettings) SetAwsEksEc2MetadataImds(v string)`

SetAwsEksEc2MetadataImds sets AwsEksEc2MetadataImds field to given value.

### HasAwsEksEc2MetadataImds

`func (o *ClusterAdvancedSettings) HasAwsEksEc2MetadataImds() bool`

HasAwsEksEc2MetadataImds returns a boolean if a field has been set.

### GetPlecoResourcesTtl

`func (o *ClusterAdvancedSettings) GetPlecoResourcesTtl() int32`

GetPlecoResourcesTtl returns the PlecoResourcesTtl field if non-nil, zero value otherwise.

### GetPlecoResourcesTtlOk

`func (o *ClusterAdvancedSettings) GetPlecoResourcesTtlOk() (*int32, bool)`

GetPlecoResourcesTtlOk returns a tuple with the PlecoResourcesTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlecoResourcesTtl

`func (o *ClusterAdvancedSettings) SetPlecoResourcesTtl(v int32)`

SetPlecoResourcesTtl sets PlecoResourcesTtl field to given value.

### HasPlecoResourcesTtl

`func (o *ClusterAdvancedSettings) HasPlecoResourcesTtl() bool`

HasPlecoResourcesTtl returns a boolean if a field has been set.

### GetRegistryMirroringMode

`func (o *ClusterAdvancedSettings) GetRegistryMirroringMode() RegistryMirroringModeEnum`

GetRegistryMirroringMode returns the RegistryMirroringMode field if non-nil, zero value otherwise.

### GetRegistryMirroringModeOk

`func (o *ClusterAdvancedSettings) GetRegistryMirroringModeOk() (*RegistryMirroringModeEnum, bool)`

GetRegistryMirroringModeOk returns a tuple with the RegistryMirroringMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryMirroringMode

`func (o *ClusterAdvancedSettings) SetRegistryMirroringMode(v RegistryMirroringModeEnum)`

SetRegistryMirroringMode sets RegistryMirroringMode field to given value.

### HasRegistryMirroringMode

`func (o *ClusterAdvancedSettings) HasRegistryMirroringMode() bool`

HasRegistryMirroringMode returns a boolean if a field has been set.

### GetNginxVcpuRequestInMilli

`func (o *ClusterAdvancedSettings) GetNginxVcpuRequestInMilli() int32`

GetNginxVcpuRequestInMilli returns the NginxVcpuRequestInMilli field if non-nil, zero value otherwise.

### GetNginxVcpuRequestInMilliOk

`func (o *ClusterAdvancedSettings) GetNginxVcpuRequestInMilliOk() (*int32, bool)`

GetNginxVcpuRequestInMilliOk returns a tuple with the NginxVcpuRequestInMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNginxVcpuRequestInMilli

`func (o *ClusterAdvancedSettings) SetNginxVcpuRequestInMilli(v int32)`

SetNginxVcpuRequestInMilli sets NginxVcpuRequestInMilli field to given value.

### HasNginxVcpuRequestInMilli

`func (o *ClusterAdvancedSettings) HasNginxVcpuRequestInMilli() bool`

HasNginxVcpuRequestInMilli returns a boolean if a field has been set.

### GetNginxVcpuLimitInMilli

`func (o *ClusterAdvancedSettings) GetNginxVcpuLimitInMilli() int32`

GetNginxVcpuLimitInMilli returns the NginxVcpuLimitInMilli field if non-nil, zero value otherwise.

### GetNginxVcpuLimitInMilliOk

`func (o *ClusterAdvancedSettings) GetNginxVcpuLimitInMilliOk() (*int32, bool)`

GetNginxVcpuLimitInMilliOk returns a tuple with the NginxVcpuLimitInMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNginxVcpuLimitInMilli

`func (o *ClusterAdvancedSettings) SetNginxVcpuLimitInMilli(v int32)`

SetNginxVcpuLimitInMilli sets NginxVcpuLimitInMilli field to given value.

### HasNginxVcpuLimitInMilli

`func (o *ClusterAdvancedSettings) HasNginxVcpuLimitInMilli() bool`

HasNginxVcpuLimitInMilli returns a boolean if a field has been set.

### GetNginxMemoryRequestInMib

`func (o *ClusterAdvancedSettings) GetNginxMemoryRequestInMib() int32`

GetNginxMemoryRequestInMib returns the NginxMemoryRequestInMib field if non-nil, zero value otherwise.

### GetNginxMemoryRequestInMibOk

`func (o *ClusterAdvancedSettings) GetNginxMemoryRequestInMibOk() (*int32, bool)`

GetNginxMemoryRequestInMibOk returns a tuple with the NginxMemoryRequestInMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNginxMemoryRequestInMib

`func (o *ClusterAdvancedSettings) SetNginxMemoryRequestInMib(v int32)`

SetNginxMemoryRequestInMib sets NginxMemoryRequestInMib field to given value.

### HasNginxMemoryRequestInMib

`func (o *ClusterAdvancedSettings) HasNginxMemoryRequestInMib() bool`

HasNginxMemoryRequestInMib returns a boolean if a field has been set.

### GetNginxMemoryLimitInMib

`func (o *ClusterAdvancedSettings) GetNginxMemoryLimitInMib() int32`

GetNginxMemoryLimitInMib returns the NginxMemoryLimitInMib field if non-nil, zero value otherwise.

### GetNginxMemoryLimitInMibOk

`func (o *ClusterAdvancedSettings) GetNginxMemoryLimitInMibOk() (*int32, bool)`

GetNginxMemoryLimitInMibOk returns a tuple with the NginxMemoryLimitInMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNginxMemoryLimitInMib

`func (o *ClusterAdvancedSettings) SetNginxMemoryLimitInMib(v int32)`

SetNginxMemoryLimitInMib sets NginxMemoryLimitInMib field to given value.

### HasNginxMemoryLimitInMib

`func (o *ClusterAdvancedSettings) HasNginxMemoryLimitInMib() bool`

HasNginxMemoryLimitInMib returns a boolean if a field has been set.

### GetNginxHpaCpuUtilizationPercentageThreshold

`func (o *ClusterAdvancedSettings) GetNginxHpaCpuUtilizationPercentageThreshold() int32`

GetNginxHpaCpuUtilizationPercentageThreshold returns the NginxHpaCpuUtilizationPercentageThreshold field if non-nil, zero value otherwise.

### GetNginxHpaCpuUtilizationPercentageThresholdOk

`func (o *ClusterAdvancedSettings) GetNginxHpaCpuUtilizationPercentageThresholdOk() (*int32, bool)`

GetNginxHpaCpuUtilizationPercentageThresholdOk returns a tuple with the NginxHpaCpuUtilizationPercentageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNginxHpaCpuUtilizationPercentageThreshold

`func (o *ClusterAdvancedSettings) SetNginxHpaCpuUtilizationPercentageThreshold(v int32)`

SetNginxHpaCpuUtilizationPercentageThreshold sets NginxHpaCpuUtilizationPercentageThreshold field to given value.

### HasNginxHpaCpuUtilizationPercentageThreshold

`func (o *ClusterAdvancedSettings) HasNginxHpaCpuUtilizationPercentageThreshold() bool`

HasNginxHpaCpuUtilizationPercentageThreshold returns a boolean if a field has been set.

### GetNginxHpaMemoryUtilizationPercentageThreshold

`func (o *ClusterAdvancedSettings) GetNginxHpaMemoryUtilizationPercentageThreshold() int32`

GetNginxHpaMemoryUtilizationPercentageThreshold returns the NginxHpaMemoryUtilizationPercentageThreshold field if non-nil, zero value otherwise.

### GetNginxHpaMemoryUtilizationPercentageThresholdOk

`func (o *ClusterAdvancedSettings) GetNginxHpaMemoryUtilizationPercentageThresholdOk() (*int32, bool)`

GetNginxHpaMemoryUtilizationPercentageThresholdOk returns a tuple with the NginxHpaMemoryUtilizationPercentageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNginxHpaMemoryUtilizationPercentageThreshold

`func (o *ClusterAdvancedSettings) SetNginxHpaMemoryUtilizationPercentageThreshold(v int32)`

SetNginxHpaMemoryUtilizationPercentageThreshold sets NginxHpaMemoryUtilizationPercentageThreshold field to given value.

### HasNginxHpaMemoryUtilizationPercentageThreshold

`func (o *ClusterAdvancedSettings) HasNginxHpaMemoryUtilizationPercentageThreshold() bool`

HasNginxHpaMemoryUtilizationPercentageThreshold returns a boolean if a field has been set.

### GetNginxHpaMinNumberInstances

`func (o *ClusterAdvancedSettings) GetNginxHpaMinNumberInstances() int32`

GetNginxHpaMinNumberInstances returns the NginxHpaMinNumberInstances field if non-nil, zero value otherwise.

### GetNginxHpaMinNumberInstancesOk

`func (o *ClusterAdvancedSettings) GetNginxHpaMinNumberInstancesOk() (*int32, bool)`

GetNginxHpaMinNumberInstancesOk returns a tuple with the NginxHpaMinNumberInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNginxHpaMinNumberInstances

`func (o *ClusterAdvancedSettings) SetNginxHpaMinNumberInstances(v int32)`

SetNginxHpaMinNumberInstances sets NginxHpaMinNumberInstances field to given value.

### HasNginxHpaMinNumberInstances

`func (o *ClusterAdvancedSettings) HasNginxHpaMinNumberInstances() bool`

HasNginxHpaMinNumberInstances returns a boolean if a field has been set.

### GetNginxHpaMaxNumberInstances

`func (o *ClusterAdvancedSettings) GetNginxHpaMaxNumberInstances() int32`

GetNginxHpaMaxNumberInstances returns the NginxHpaMaxNumberInstances field if non-nil, zero value otherwise.

### GetNginxHpaMaxNumberInstancesOk

`func (o *ClusterAdvancedSettings) GetNginxHpaMaxNumberInstancesOk() (*int32, bool)`

GetNginxHpaMaxNumberInstancesOk returns a tuple with the NginxHpaMaxNumberInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNginxHpaMaxNumberInstances

`func (o *ClusterAdvancedSettings) SetNginxHpaMaxNumberInstances(v int32)`

SetNginxHpaMaxNumberInstances sets NginxHpaMaxNumberInstances field to given value.

### HasNginxHpaMaxNumberInstances

`func (o *ClusterAdvancedSettings) HasNginxHpaMaxNumberInstances() bool`

HasNginxHpaMaxNumberInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


