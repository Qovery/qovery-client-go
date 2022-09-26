# OrganizationCustomRoleProjectPermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**IsAdmin** | Pointer to **bool** | If &#x60;is_admin&#x60; is &#x60;true&#x60;, the user is: - automatically &#x60;MANAGER&#x60; for each environment type - allowed to manage project deployment rules - able to delete the project Note that &#x60;permissions&#x60; can then be ignored for this project  | [optional] [default to false]
**Permissions** | Pointer to [**[]OrganizationCustomRoleRequestProjectPermissionsInnerPermissionsInner**](OrganizationCustomRoleRequestProjectPermissionsInnerPermissionsInner.md) |  | [optional] 

## Methods

### NewOrganizationCustomRoleProjectPermissionsInner

`func NewOrganizationCustomRoleProjectPermissionsInner() *OrganizationCustomRoleProjectPermissionsInner`

NewOrganizationCustomRoleProjectPermissionsInner instantiates a new OrganizationCustomRoleProjectPermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomRoleProjectPermissionsInnerWithDefaults

`func NewOrganizationCustomRoleProjectPermissionsInnerWithDefaults() *OrganizationCustomRoleProjectPermissionsInner`

NewOrganizationCustomRoleProjectPermissionsInnerWithDefaults instantiates a new OrganizationCustomRoleProjectPermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *OrganizationCustomRoleProjectPermissionsInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OrganizationCustomRoleProjectPermissionsInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OrganizationCustomRoleProjectPermissionsInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *OrganizationCustomRoleProjectPermissionsInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *OrganizationCustomRoleProjectPermissionsInner) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *OrganizationCustomRoleProjectPermissionsInner) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *OrganizationCustomRoleProjectPermissionsInner) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *OrganizationCustomRoleProjectPermissionsInner) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetIsAdmin

`func (o *OrganizationCustomRoleProjectPermissionsInner) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *OrganizationCustomRoleProjectPermissionsInner) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *OrganizationCustomRoleProjectPermissionsInner) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *OrganizationCustomRoleProjectPermissionsInner) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetPermissions

`func (o *OrganizationCustomRoleProjectPermissionsInner) GetPermissions() []OrganizationCustomRoleRequestProjectPermissionsInnerPermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OrganizationCustomRoleProjectPermissionsInner) GetPermissionsOk() (*[]OrganizationCustomRoleRequestProjectPermissionsInnerPermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OrganizationCustomRoleProjectPermissionsInner) SetPermissions(v []OrganizationCustomRoleRequestProjectPermissionsInnerPermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *OrganizationCustomRoleProjectPermissionsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


