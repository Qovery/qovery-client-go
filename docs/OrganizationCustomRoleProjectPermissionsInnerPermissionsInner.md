# OrganizationCustomRoleProjectPermissionsInnerPermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentType** | Pointer to [**EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | [optional] 
**Permission** | Pointer to [**OrganizationCustomRoleProjectPermission**](OrganizationCustomRoleProjectPermission.md) |  | [optional] 
**Shell** | **bool** | Allows opening a shell on the services of this environment type, from the CLI (&#x60;qovery shell&#x60;) or the console | 
**PortForward** | **bool** | Allows port-forwarding to the services of this environment type (&#x60;qovery port-forward&#x60;) | 

## Methods

### NewOrganizationCustomRoleProjectPermissionsInnerPermissionsInner

`func NewOrganizationCustomRoleProjectPermissionsInnerPermissionsInner(shell bool, portForward bool, ) *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner`

NewOrganizationCustomRoleProjectPermissionsInnerPermissionsInner instantiates a new OrganizationCustomRoleProjectPermissionsInnerPermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomRoleProjectPermissionsInnerPermissionsInnerWithDefaults

`func NewOrganizationCustomRoleProjectPermissionsInnerPermissionsInnerWithDefaults() *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner`

NewOrganizationCustomRoleProjectPermissionsInnerPermissionsInnerWithDefaults instantiates a new OrganizationCustomRoleProjectPermissionsInnerPermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentType

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) GetEnvironmentType() EnvironmentModeEnum`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) GetEnvironmentTypeOk() (*EnvironmentModeEnum, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) SetEnvironmentType(v EnvironmentModeEnum)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetPermission

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) GetPermission() OrganizationCustomRoleProjectPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) GetPermissionOk() (*OrganizationCustomRoleProjectPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) SetPermission(v OrganizationCustomRoleProjectPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetShell

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) GetShell() bool`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) GetShellOk() (*bool, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) SetShell(v bool)`

SetShell sets Shell field to given value.


### GetPortForward

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) GetPortForward() bool`

GetPortForward returns the PortForward field if non-nil, zero value otherwise.

### GetPortForwardOk

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) GetPortForwardOk() (*bool, bool)`

GetPortForwardOk returns a tuple with the PortForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortForward

`func (o *OrganizationCustomRoleProjectPermissionsInnerPermissionsInner) SetPortForward(v bool)`

SetPortForward sets PortForward field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


