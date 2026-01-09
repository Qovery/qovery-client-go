# OrganizationEnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** | name is case insensitive | 
**Organization** | [**ReferenceObject**](ReferenceObject.md) |  | 
**Project** | [**ReferenceObject**](ReferenceObject.md) |  | 
**Mode** | [**EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | 

## Methods

### NewOrganizationEnvironmentResponse

`func NewOrganizationEnvironmentResponse(id string, createdAt time.Time, name string, organization ReferenceObject, project ReferenceObject, mode EnvironmentModeEnum, ) *OrganizationEnvironmentResponse`

NewOrganizationEnvironmentResponse instantiates a new OrganizationEnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationEnvironmentResponseWithDefaults

`func NewOrganizationEnvironmentResponseWithDefaults() *OrganizationEnvironmentResponse`

NewOrganizationEnvironmentResponseWithDefaults instantiates a new OrganizationEnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationEnvironmentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationEnvironmentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationEnvironmentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrganizationEnvironmentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationEnvironmentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationEnvironmentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationEnvironmentResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationEnvironmentResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationEnvironmentResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationEnvironmentResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *OrganizationEnvironmentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationEnvironmentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationEnvironmentResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *OrganizationEnvironmentResponse) GetOrganization() ReferenceObject`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationEnvironmentResponse) GetOrganizationOk() (*ReferenceObject, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationEnvironmentResponse) SetOrganization(v ReferenceObject)`

SetOrganization sets Organization field to given value.


### GetProject

`func (o *OrganizationEnvironmentResponse) GetProject() ReferenceObject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *OrganizationEnvironmentResponse) GetProjectOk() (*ReferenceObject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *OrganizationEnvironmentResponse) SetProject(v ReferenceObject)`

SetProject sets Project field to given value.


### GetMode

`func (o *OrganizationEnvironmentResponse) GetMode() EnvironmentModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OrganizationEnvironmentResponse) GetModeOk() (*EnvironmentModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OrganizationEnvironmentResponse) SetMode(v EnvironmentModeEnum)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


