# OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentType** | Pointer to [**EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | [optional] 
**Permission** | Pointer to [**OrganizationCustomRoleProjectPermission**](OrganizationCustomRoleProjectPermission.md) |  | [optional] 
**Shell** | **bool** | Allows opening a shell on the services of this environment type, from the CLI (&#x60;qovery shell&#x60;) or the console. Cannot be &#x60;true&#x60; with &#x60;NO_ACCESS&#x60;.  | 
**PortForward** | **bool** | Allows port-forwarding to the services of this environment type (&#x60;qovery port-forward&#x60;). Cannot be &#x60;true&#x60; with &#x60;NO_ACCESS&#x60;.  | 

## Methods

### NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner

`func NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner(shell bool, portForward bool, ) *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner`

NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner instantiates a new OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInnerWithDefaults

`func NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInnerWithDefaults() *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner`

NewOrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInnerWithDefaults instantiates a new OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentType

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) GetEnvironmentType() EnvironmentModeEnum`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) GetEnvironmentTypeOk() (*EnvironmentModeEnum, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) SetEnvironmentType(v EnvironmentModeEnum)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetPermission

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) GetPermission() OrganizationCustomRoleProjectPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) GetPermissionOk() (*OrganizationCustomRoleProjectPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) SetPermission(v OrganizationCustomRoleProjectPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetShell

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) GetShell() bool`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) GetShellOk() (*bool, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) SetShell(v bool)`

SetShell sets Shell field to given value.


### GetPortForward

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) GetPortForward() bool`

GetPortForward returns the PortForward field if non-nil, zero value otherwise.

### GetPortForwardOk

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) GetPortForwardOk() (*bool, bool)`

GetPortForwardOk returns a tuple with the PortForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortForward

`func (o *OrganizationCustomRoleUpdateRequestProjectPermissionsInnerPermissionsInner) SetPortForward(v bool)`

SetPortForward sets PortForward field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


