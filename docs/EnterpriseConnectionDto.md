# EnterpriseConnectionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionName** | **string** | The connection name | 
**DefaultRole** | **string** | The purpose of this default role is to be associated to your users if: - you choose to not expose your IDPs groups to the SAML / OIDC connection - no associated group is found in your &#x60;group_mappings&#x60; defined  You can define either a Qovery provided role (i.e &#x60;viewer&#x60;) or one of your custom role&#x60;s uuid.  | 
**EnforceGroupSync** | **bool** | * if &#x60;true&#x60;, roles will be synchronized at each user login according to your &#x60;group_mappings&#x60; configuration based on your IDP groups * if &#x60;false&#x60;, no synchronization is done for your users and &#x60;group_mappings&#x60; configuration will be ignored  | 
**GroupMappings** | **map[string][]string** | This will allow to create mapping rules based on your IDP group names.   It&#39;s a dictionnary having: - key: either a Qovery provided role (i.e &#x60;viewer&#x60;) or one of your custom role&#x60;s uuid - value: an array of your IDP group names  Example: \&quot;I want to associate the Qovery role &#x60;devops&#x60; to my IDP groups [&#39;Administrators&#39;, &#39;DevSecOps&#39;]\&quot;  | 

## Methods

### NewEnterpriseConnectionDto

`func NewEnterpriseConnectionDto(connectionName string, defaultRole string, enforceGroupSync bool, groupMappings map[string][]string, ) *EnterpriseConnectionDto`

NewEnterpriseConnectionDto instantiates a new EnterpriseConnectionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseConnectionDtoWithDefaults

`func NewEnterpriseConnectionDtoWithDefaults() *EnterpriseConnectionDto`

NewEnterpriseConnectionDtoWithDefaults instantiates a new EnterpriseConnectionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionName

`func (o *EnterpriseConnectionDto) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *EnterpriseConnectionDto) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *EnterpriseConnectionDto) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetDefaultRole

`func (o *EnterpriseConnectionDto) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *EnterpriseConnectionDto) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *EnterpriseConnectionDto) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.


### GetEnforceGroupSync

`func (o *EnterpriseConnectionDto) GetEnforceGroupSync() bool`

GetEnforceGroupSync returns the EnforceGroupSync field if non-nil, zero value otherwise.

### GetEnforceGroupSyncOk

`func (o *EnterpriseConnectionDto) GetEnforceGroupSyncOk() (*bool, bool)`

GetEnforceGroupSyncOk returns a tuple with the EnforceGroupSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceGroupSync

`func (o *EnterpriseConnectionDto) SetEnforceGroupSync(v bool)`

SetEnforceGroupSync sets EnforceGroupSync field to given value.


### GetGroupMappings

`func (o *EnterpriseConnectionDto) GetGroupMappings() map[string][]string`

GetGroupMappings returns the GroupMappings field if non-nil, zero value otherwise.

### GetGroupMappingsOk

`func (o *EnterpriseConnectionDto) GetGroupMappingsOk() (*map[string][]string, bool)`

GetGroupMappingsOk returns a tuple with the GroupMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMappings

`func (o *EnterpriseConnectionDto) SetGroupMappings(v map[string][]string)`

SetGroupMappings sets GroupMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


