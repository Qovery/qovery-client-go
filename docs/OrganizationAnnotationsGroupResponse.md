# OrganizationAnnotationsGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Annotations** | [**[]Annotation**](Annotation.md) |  | 
**Scopes** | [**[]OrganizationAnnotationsGroupScopeEnum**](OrganizationAnnotationsGroupScopeEnum.md) |  | 

## Methods

### NewOrganizationAnnotationsGroupResponse

`func NewOrganizationAnnotationsGroupResponse(id string, createdAt time.Time, name string, annotations []Annotation, scopes []OrganizationAnnotationsGroupScopeEnum, ) *OrganizationAnnotationsGroupResponse`

NewOrganizationAnnotationsGroupResponse instantiates a new OrganizationAnnotationsGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationAnnotationsGroupResponseWithDefaults

`func NewOrganizationAnnotationsGroupResponseWithDefaults() *OrganizationAnnotationsGroupResponse`

NewOrganizationAnnotationsGroupResponseWithDefaults instantiates a new OrganizationAnnotationsGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationAnnotationsGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationAnnotationsGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationAnnotationsGroupResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrganizationAnnotationsGroupResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationAnnotationsGroupResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationAnnotationsGroupResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationAnnotationsGroupResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationAnnotationsGroupResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationAnnotationsGroupResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationAnnotationsGroupResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *OrganizationAnnotationsGroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationAnnotationsGroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationAnnotationsGroupResponse) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *OrganizationAnnotationsGroupResponse) GetAnnotations() []Annotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *OrganizationAnnotationsGroupResponse) GetAnnotationsOk() (*[]Annotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *OrganizationAnnotationsGroupResponse) SetAnnotations(v []Annotation)`

SetAnnotations sets Annotations field to given value.


### GetScopes

`func (o *OrganizationAnnotationsGroupResponse) GetScopes() []OrganizationAnnotationsGroupScopeEnum`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OrganizationAnnotationsGroupResponse) GetScopesOk() (*[]OrganizationAnnotationsGroupScopeEnum, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OrganizationAnnotationsGroupResponse) SetScopes(v []OrganizationAnnotationsGroupScopeEnum)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


