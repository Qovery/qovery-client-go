# OrganizationCustomRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ClusterPermissions** | Pointer to [**[]OrganizationCustomRoleClusterPermissionsInner**](OrganizationCustomRoleClusterPermissionsInner.md) |  | [optional] 
**ProjectPermissions** | Pointer to [**[]OrganizationCustomRoleProjectPermissionsInner**](OrganizationCustomRoleProjectPermissionsInner.md) |  | [optional] 

## Methods

### NewOrganizationCustomRole

`func NewOrganizationCustomRole() *OrganizationCustomRole`

NewOrganizationCustomRole instantiates a new OrganizationCustomRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomRoleWithDefaults

`func NewOrganizationCustomRoleWithDefaults() *OrganizationCustomRole`

NewOrganizationCustomRoleWithDefaults instantiates a new OrganizationCustomRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationCustomRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationCustomRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationCustomRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationCustomRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationCustomRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationCustomRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationCustomRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationCustomRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationCustomRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationCustomRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationCustomRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationCustomRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClusterPermissions

`func (o *OrganizationCustomRole) GetClusterPermissions() []OrganizationCustomRoleClusterPermissionsInner`

GetClusterPermissions returns the ClusterPermissions field if non-nil, zero value otherwise.

### GetClusterPermissionsOk

`func (o *OrganizationCustomRole) GetClusterPermissionsOk() (*[]OrganizationCustomRoleClusterPermissionsInner, bool)`

GetClusterPermissionsOk returns a tuple with the ClusterPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPermissions

`func (o *OrganizationCustomRole) SetClusterPermissions(v []OrganizationCustomRoleClusterPermissionsInner)`

SetClusterPermissions sets ClusterPermissions field to given value.

### HasClusterPermissions

`func (o *OrganizationCustomRole) HasClusterPermissions() bool`

HasClusterPermissions returns a boolean if a field has been set.

### GetProjectPermissions

`func (o *OrganizationCustomRole) GetProjectPermissions() []OrganizationCustomRoleProjectPermissionsInner`

GetProjectPermissions returns the ProjectPermissions field if non-nil, zero value otherwise.

### GetProjectPermissionsOk

`func (o *OrganizationCustomRole) GetProjectPermissionsOk() (*[]OrganizationCustomRoleProjectPermissionsInner, bool)`

GetProjectPermissionsOk returns a tuple with the ProjectPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPermissions

`func (o *OrganizationCustomRole) SetProjectPermissions(v []OrganizationCustomRoleProjectPermissionsInner)`

SetProjectPermissions sets ProjectPermissions field to given value.

### HasProjectPermissions

`func (o *OrganizationCustomRole) HasProjectPermissions() bool`

HasProjectPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


