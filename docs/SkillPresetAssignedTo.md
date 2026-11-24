# SkillPresetAssignedTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationDefault** | Pointer to **bool** | Whether the preset applies to every member with no member or role assignment of their own. | [optional] [default to false]
**RoleIds** | Pointer to **[]string** | Default or custom roles of the organization, as returned by \&quot;List organization roles\&quot;. | [optional] [default to []]
**UserIds** | Pointer to **[]string** | Members of the organization, by the id returned by \&quot;List organization members\&quot;. | [optional] [default to []]

## Methods

### NewSkillPresetAssignedTo

`func NewSkillPresetAssignedTo() *SkillPresetAssignedTo`

NewSkillPresetAssignedTo instantiates a new SkillPresetAssignedTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillPresetAssignedToWithDefaults

`func NewSkillPresetAssignedToWithDefaults() *SkillPresetAssignedTo`

NewSkillPresetAssignedToWithDefaults instantiates a new SkillPresetAssignedTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationDefault

`func (o *SkillPresetAssignedTo) GetOrganizationDefault() bool`

GetOrganizationDefault returns the OrganizationDefault field if non-nil, zero value otherwise.

### GetOrganizationDefaultOk

`func (o *SkillPresetAssignedTo) GetOrganizationDefaultOk() (*bool, bool)`

GetOrganizationDefaultOk returns a tuple with the OrganizationDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationDefault

`func (o *SkillPresetAssignedTo) SetOrganizationDefault(v bool)`

SetOrganizationDefault sets OrganizationDefault field to given value.

### HasOrganizationDefault

`func (o *SkillPresetAssignedTo) HasOrganizationDefault() bool`

HasOrganizationDefault returns a boolean if a field has been set.

### GetRoleIds

`func (o *SkillPresetAssignedTo) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *SkillPresetAssignedTo) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *SkillPresetAssignedTo) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *SkillPresetAssignedTo) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetUserIds

`func (o *SkillPresetAssignedTo) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *SkillPresetAssignedTo) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *SkillPresetAssignedTo) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *SkillPresetAssignedTo) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


