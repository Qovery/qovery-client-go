# OrganizationCustomRoleRequestProjectPermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** |  | [optional] 
**IsAdmin** | Pointer to **bool** | If &#x60;is_admin&#x60; is &#x60;true&#x60;, the user is: - automatically &#x60;MANAGER&#x60; for each environment type - allowed to manage project deployment rules - able to delete the project Note that &#x60;permissions&#x60; can then be ignored for this project  | [optional] [default to false]
**Permissions** | Pointer to [**[]OrganizationCustomRoleRequestProjectPermissionsInnerPermissionsInner**](OrganizationCustomRoleRequestProjectPermissionsInnerPermissionsInner.md) | Mandatory if &#x60;is_admin&#x60; is &#x60;false&#x60;   Should contain an entry for every environment type: - &#x60;DEVELOPMENT&#x60; - &#x60;PREVIEW&#x60; - &#x60;STAGING&#x60; - &#x60;PRODUCTION&#x60;  | [optional] 

## Methods

### NewOrganizationCustomRoleRequestProjectPermissionsInner

`func NewOrganizationCustomRoleRequestProjectPermissionsInner() *OrganizationCustomRoleRequestProjectPermissionsInner`

NewOrganizationCustomRoleRequestProjectPermissionsInner instantiates a new OrganizationCustomRoleRequestProjectPermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomRoleRequestProjectPermissionsInnerWithDefaults

`func NewOrganizationCustomRoleRequestProjectPermissionsInnerWithDefaults() *OrganizationCustomRoleRequestProjectPermissionsInner`

NewOrganizationCustomRoleRequestProjectPermissionsInnerWithDefaults instantiates a new OrganizationCustomRoleRequestProjectPermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *OrganizationCustomRoleRequestProjectPermissionsInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OrganizationCustomRoleRequestProjectPermissionsInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OrganizationCustomRoleRequestProjectPermissionsInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *OrganizationCustomRoleRequestProjectPermissionsInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetIsAdmin

`func (o *OrganizationCustomRoleRequestProjectPermissionsInner) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *OrganizationCustomRoleRequestProjectPermissionsInner) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *OrganizationCustomRoleRequestProjectPermissionsInner) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *OrganizationCustomRoleRequestProjectPermissionsInner) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetPermissions

`func (o *OrganizationCustomRoleRequestProjectPermissionsInner) GetPermissions() []OrganizationCustomRoleRequestProjectPermissionsInnerPermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OrganizationCustomRoleRequestProjectPermissionsInner) GetPermissionsOk() (*[]OrganizationCustomRoleRequestProjectPermissionsInnerPermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OrganizationCustomRoleRequestProjectPermissionsInner) SetPermissions(v []OrganizationCustomRoleRequestProjectPermissionsInnerPermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *OrganizationCustomRoleRequestProjectPermissionsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


