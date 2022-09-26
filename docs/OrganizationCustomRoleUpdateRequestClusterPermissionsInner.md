# OrganizationCustomRoleUpdateRequestClusterPermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**Permission** | Pointer to [**OrganizationCustomRoleClusterPermission**](OrganizationCustomRoleClusterPermission.md) |  | [optional] 

## Methods

### NewOrganizationCustomRoleUpdateRequestClusterPermissionsInner

`func NewOrganizationCustomRoleUpdateRequestClusterPermissionsInner() *OrganizationCustomRoleUpdateRequestClusterPermissionsInner`

NewOrganizationCustomRoleUpdateRequestClusterPermissionsInner instantiates a new OrganizationCustomRoleUpdateRequestClusterPermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomRoleUpdateRequestClusterPermissionsInnerWithDefaults

`func NewOrganizationCustomRoleUpdateRequestClusterPermissionsInnerWithDefaults() *OrganizationCustomRoleUpdateRequestClusterPermissionsInner`

NewOrganizationCustomRoleUpdateRequestClusterPermissionsInnerWithDefaults instantiates a new OrganizationCustomRoleUpdateRequestClusterPermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *OrganizationCustomRoleUpdateRequestClusterPermissionsInner) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *OrganizationCustomRoleUpdateRequestClusterPermissionsInner) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *OrganizationCustomRoleUpdateRequestClusterPermissionsInner) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *OrganizationCustomRoleUpdateRequestClusterPermissionsInner) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetPermission

`func (o *OrganizationCustomRoleUpdateRequestClusterPermissionsInner) GetPermission() OrganizationCustomRoleClusterPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *OrganizationCustomRoleUpdateRequestClusterPermissionsInner) GetPermissionOk() (*OrganizationCustomRoleClusterPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *OrganizationCustomRoleUpdateRequestClusterPermissionsInner) SetPermission(v OrganizationCustomRoleClusterPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *OrganizationCustomRoleUpdateRequestClusterPermissionsInner) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


