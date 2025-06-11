# ClusterCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**AccessKeyId** | **string** |  | 
**ObjectType** | **string** |  | 
**ScalewayAccessKey** | **string** |  | 
**ScalewayProjectId** | **string** |  | 
**ScalewayOrganizationId** | **string** |  | 
**RoleArn** | **string** |  | 
**AzureSubscriptionId** | **string** |  | 
**AzureTenantId** | **string** |  | 

## Methods

### NewClusterCredentials

`func NewClusterCredentials(id string, name string, accessKeyId string, objectType string, scalewayAccessKey string, scalewayProjectId string, scalewayOrganizationId string, roleArn string, azureSubscriptionId string, azureTenantId string, ) *ClusterCredentials`

NewClusterCredentials instantiates a new ClusterCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCredentialsWithDefaults

`func NewClusterCredentialsWithDefaults() *ClusterCredentials`

NewClusterCredentialsWithDefaults instantiates a new ClusterCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterCredentials) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterCredentials) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterCredentials) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ClusterCredentials) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterCredentials) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterCredentials) SetName(v string)`

SetName sets Name field to given value.


### GetAccessKeyId

`func (o *ClusterCredentials) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *ClusterCredentials) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *ClusterCredentials) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### GetObjectType

`func (o *ClusterCredentials) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ClusterCredentials) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ClusterCredentials) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetScalewayAccessKey

`func (o *ClusterCredentials) GetScalewayAccessKey() string`

GetScalewayAccessKey returns the ScalewayAccessKey field if non-nil, zero value otherwise.

### GetScalewayAccessKeyOk

`func (o *ClusterCredentials) GetScalewayAccessKeyOk() (*string, bool)`

GetScalewayAccessKeyOk returns a tuple with the ScalewayAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalewayAccessKey

`func (o *ClusterCredentials) SetScalewayAccessKey(v string)`

SetScalewayAccessKey sets ScalewayAccessKey field to given value.


### GetScalewayProjectId

`func (o *ClusterCredentials) GetScalewayProjectId() string`

GetScalewayProjectId returns the ScalewayProjectId field if non-nil, zero value otherwise.

### GetScalewayProjectIdOk

`func (o *ClusterCredentials) GetScalewayProjectIdOk() (*string, bool)`

GetScalewayProjectIdOk returns a tuple with the ScalewayProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalewayProjectId

`func (o *ClusterCredentials) SetScalewayProjectId(v string)`

SetScalewayProjectId sets ScalewayProjectId field to given value.


### GetScalewayOrganizationId

`func (o *ClusterCredentials) GetScalewayOrganizationId() string`

GetScalewayOrganizationId returns the ScalewayOrganizationId field if non-nil, zero value otherwise.

### GetScalewayOrganizationIdOk

`func (o *ClusterCredentials) GetScalewayOrganizationIdOk() (*string, bool)`

GetScalewayOrganizationIdOk returns a tuple with the ScalewayOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalewayOrganizationId

`func (o *ClusterCredentials) SetScalewayOrganizationId(v string)`

SetScalewayOrganizationId sets ScalewayOrganizationId field to given value.


### GetRoleArn

`func (o *ClusterCredentials) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *ClusterCredentials) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *ClusterCredentials) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetAzureSubscriptionId

`func (o *ClusterCredentials) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *ClusterCredentials) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *ClusterCredentials) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### GetAzureTenantId

`func (o *ClusterCredentials) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *ClusterCredentials) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *ClusterCredentials) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


