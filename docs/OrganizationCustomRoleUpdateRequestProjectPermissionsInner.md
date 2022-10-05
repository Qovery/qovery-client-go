# OrganizationCustomRoleUpdateRequestProjectPermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** |  | [optional] 
**IsAdmin** | Pointer to **bool** | If &#x60;is_admin&#x60; is &#x60;true&#x60;, the user is: - automatically &#x60;MANAGER&#x60; for each environment type - allowed to manage project deployment rules - able to delete the project    Note that &#x60;permissions&#x60; can then be ignored for this project  | [optional] [default to false]
**Permissions** | Pointer to [**[]OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner**](OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner.md) | Mandatory if &#x60;is_admin&#x60; is &#x60;false&#x60;   Should contain an entry for every environment type: - &#x60;DEVELOPMENT&#x60; - &#x60;PREVIEW&#x60; - &#x60;STAGING&#x60; - &#x60;PRODUCTION&#x60;  | [optional] 

## Methods

### NewOrganizationCustomRoleUpdateRequestProjectPermissionsInner

`func NewOrganizationCustomRoleUpdateRequestProjectPermissionsInner() *OrganizationCustomRoleUpdateRequestProjectPermissionsInner`

NewOrganizationCustomRoleUpdateRequestProjectPermissionsInner instantiates a new OrganizationCustomRoleUpdateRequestProjectPermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerWithDefaults

`func NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerWithDefaults() *OrganizationCustomRoleUpdateRequestProjectPermissionsInner`

NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerWithDefaults instantiates a new OrganizationCustomRoleUpdateRequestProjectPermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetIsAdmin

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetPermissions

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) GetPermissions() []OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) GetPermissionsOk() (*[]OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) SetPermissions(v []OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


