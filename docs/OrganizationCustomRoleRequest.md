# OrganizationCustomRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ClusterPermissions** | [**[]OrganizationCustomRoleRequestClusterPermissionsInner**](OrganizationCustomRoleRequestClusterPermissionsInner.md) | Should contain an entry for every existing cluster | 
**ProjectPermissions** | [**[]OrganizationCustomRoleRequestProjectPermissionsInner**](OrganizationCustomRoleRequestProjectPermissionsInner.md) | Should contain an entry for every existing project | 

## Methods

### NewOrganizationCustomRoleRequest

`func NewOrganizationCustomRoleRequest(name string, clusterPermissions []OrganizationCustomRoleRequestClusterPermissionsInner, projectPermissions []OrganizationCustomRoleRequestProjectPermissionsInner, ) *OrganizationCustomRoleRequest`

NewOrganizationCustomRoleRequest instantiates a new OrganizationCustomRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomRoleRequestWithDefaults

`func NewOrganizationCustomRoleRequestWithDefaults() *OrganizationCustomRoleRequest`

NewOrganizationCustomRoleRequestWithDefaults instantiates a new OrganizationCustomRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationCustomRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationCustomRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationCustomRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OrganizationCustomRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationCustomRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationCustomRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationCustomRoleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClusterPermissions

`func (o *OrganizationCustomRoleRequest) GetClusterPermissions() []OrganizationCustomRoleRequestClusterPermissionsInner`

GetClusterPermissions returns the ClusterPermissions field if non-nil, zero value otherwise.

### GetClusterPermissionsOk

`func (o *OrganizationCustomRoleRequest) GetClusterPermissionsOk() (*[]OrganizationCustomRoleRequestClusterPermissionsInner, bool)`

GetClusterPermissionsOk returns a tuple with the ClusterPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPermissions

`func (o *OrganizationCustomRoleRequest) SetClusterPermissions(v []OrganizationCustomRoleRequestClusterPermissionsInner)`

SetClusterPermissions sets ClusterPermissions field to given value.


### GetProjectPermissions

`func (o *OrganizationCustomRoleRequest) GetProjectPermissions() []OrganizationCustomRoleRequestProjectPermissionsInner`

GetProjectPermissions returns the ProjectPermissions field if non-nil, zero value otherwise.

### GetProjectPermissionsOk

`func (o *OrganizationCustomRoleRequest) GetProjectPermissionsOk() (*[]OrganizationCustomRoleRequestProjectPermissionsInner, bool)`

GetProjectPermissionsOk returns a tuple with the ProjectPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPermissions

`func (o *OrganizationCustomRoleRequest) SetProjectPermissions(v []OrganizationCustomRoleRequestProjectPermissionsInner)`

SetProjectPermissions sets ProjectPermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


