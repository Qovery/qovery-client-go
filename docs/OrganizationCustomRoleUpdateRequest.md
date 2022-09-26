# OrganizationCustomRoleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ClusterPermissions** | [**[]OrganizationCustomRoleUpdateRequestClusterPermissionsInner**](OrganizationCustomRoleUpdateRequestClusterPermissionsInner.md) | Should contain an entry for every existing cluster | 
**ProjectPermissions** | [**[]OrganizationCustomRoleUpdateRequestProjectPermissionsInner**](OrganizationCustomRoleUpdateRequestProjectPermissionsInner.md) | Should contain an entry for every existing project | 

## Methods

### NewOrganizationCustomRoleUpdateRequest

`func NewOrganizationCustomRoleUpdateRequest(name string, clusterPermissions []OrganizationCustomRoleUpdateRequestClusterPermissionsInner, projectPermissions []OrganizationCustomRoleUpdateRequestProjectPermissionsInner, ) *OrganizationCustomRoleUpdateRequest`

NewOrganizationCustomRoleUpdateRequest instantiates a new OrganizationCustomRoleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomRoleUpdateRequestWithDefaults

`func NewOrganizationCustomRoleUpdateRequestWithDefaults() *OrganizationCustomRoleUpdateRequest`

NewOrganizationCustomRoleUpdateRequestWithDefaults instantiates a new OrganizationCustomRoleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationCustomRoleUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationCustomRoleUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationCustomRoleUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OrganizationCustomRoleUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationCustomRoleUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationCustomRoleUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationCustomRoleUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClusterPermissions

`func (o *OrganizationCustomRoleUpdateRequest) GetClusterPermissions() []OrganizationCustomRoleUpdateRequestClusterPermissionsInner`

GetClusterPermissions returns the ClusterPermissions field if non-nil, zero value otherwise.

### GetClusterPermissionsOk

`func (o *OrganizationCustomRoleUpdateRequest) GetClusterPermissionsOk() (*[]OrganizationCustomRoleUpdateRequestClusterPermissionsInner, bool)`

GetClusterPermissionsOk returns a tuple with the ClusterPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPermissions

`func (o *OrganizationCustomRoleUpdateRequest) SetClusterPermissions(v []OrganizationCustomRoleUpdateRequestClusterPermissionsInner)`

SetClusterPermissions sets ClusterPermissions field to given value.


### GetProjectPermissions

`func (o *OrganizationCustomRoleUpdateRequest) GetProjectPermissions() []OrganizationCustomRoleUpdateRequestProjectPermissionsInner`

GetProjectPermissions returns the ProjectPermissions field if non-nil, zero value otherwise.

### GetProjectPermissionsOk

`func (o *OrganizationCustomRoleUpdateRequest) GetProjectPermissionsOk() (*[]OrganizationCustomRoleUpdateRequestProjectPermissionsInner, bool)`

GetProjectPermissionsOk returns a tuple with the ProjectPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPermissions

`func (o *OrganizationCustomRoleUpdateRequest) SetProjectPermissions(v []OrganizationCustomRoleUpdateRequestProjectPermissionsInner)`

SetProjectPermissions sets ProjectPermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


