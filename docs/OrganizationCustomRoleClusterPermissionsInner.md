# OrganizationCustomRoleClusterPermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**Permission** | Pointer to [**OrganizationCustomRoleClusterPermission**](OrganizationCustomRoleClusterPermission.md) |  | [optional] 

## Methods

### NewOrganizationCustomRoleClusterPermissionsInner

`func NewOrganizationCustomRoleClusterPermissionsInner() *OrganizationCustomRoleClusterPermissionsInner`

NewOrganizationCustomRoleClusterPermissionsInner instantiates a new OrganizationCustomRoleClusterPermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomRoleClusterPermissionsInnerWithDefaults

`func NewOrganizationCustomRoleClusterPermissionsInnerWithDefaults() *OrganizationCustomRoleClusterPermissionsInner`

NewOrganizationCustomRoleClusterPermissionsInnerWithDefaults instantiates a new OrganizationCustomRoleClusterPermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *OrganizationCustomRoleClusterPermissionsInner) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *OrganizationCustomRoleClusterPermissionsInner) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *OrganizationCustomRoleClusterPermissionsInner) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *OrganizationCustomRoleClusterPermissionsInner) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterName

`func (o *OrganizationCustomRoleClusterPermissionsInner) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *OrganizationCustomRoleClusterPermissionsInner) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *OrganizationCustomRoleClusterPermissionsInner) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *OrganizationCustomRoleClusterPermissionsInner) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetPermission

`func (o *OrganizationCustomRoleClusterPermissionsInner) GetPermission() OrganizationCustomRoleClusterPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *OrganizationCustomRoleClusterPermissionsInner) GetPermissionOk() (*OrganizationCustomRoleClusterPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *OrganizationCustomRoleClusterPermissionsInner) SetPermission(v OrganizationCustomRoleClusterPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *OrganizationCustomRoleClusterPermissionsInner) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


