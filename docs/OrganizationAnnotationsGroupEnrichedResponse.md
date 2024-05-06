# OrganizationAnnotationsGroupEnrichedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Annotations** | [**[]Annotation**](Annotation.md) |  | 
**Scopes** | [**[]OrganizationAnnotationsGroupScopeEnum**](OrganizationAnnotationsGroupScopeEnum.md) |  | 
**AssociatedItemsCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewOrganizationAnnotationsGroupEnrichedResponse

`func NewOrganizationAnnotationsGroupEnrichedResponse(id string, createdAt time.Time, name string, annotations []Annotation, scopes []OrganizationAnnotationsGroupScopeEnum, ) *OrganizationAnnotationsGroupEnrichedResponse`

NewOrganizationAnnotationsGroupEnrichedResponse instantiates a new OrganizationAnnotationsGroupEnrichedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationAnnotationsGroupEnrichedResponseWithDefaults

`func NewOrganizationAnnotationsGroupEnrichedResponseWithDefaults() *OrganizationAnnotationsGroupEnrichedResponse`

NewOrganizationAnnotationsGroupEnrichedResponseWithDefaults instantiates a new OrganizationAnnotationsGroupEnrichedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationAnnotationsGroupEnrichedResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationAnnotationsGroupEnrichedResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationAnnotationsGroupEnrichedResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationAnnotationsGroupEnrichedResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationAnnotationsGroupEnrichedResponse) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetAnnotations() []Annotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetAnnotationsOk() (*[]Annotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *OrganizationAnnotationsGroupEnrichedResponse) SetAnnotations(v []Annotation)`

SetAnnotations sets Annotations field to given value.


### GetScopes

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetScopes() []OrganizationAnnotationsGroupScopeEnum`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetScopesOk() (*[]OrganizationAnnotationsGroupScopeEnum, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OrganizationAnnotationsGroupEnrichedResponse) SetScopes(v []OrganizationAnnotationsGroupScopeEnum)`

SetScopes sets Scopes field to given value.


### GetAssociatedItemsCount

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetAssociatedItemsCount() int32`

GetAssociatedItemsCount returns the AssociatedItemsCount field if non-nil, zero value otherwise.

### GetAssociatedItemsCountOk

`func (o *OrganizationAnnotationsGroupEnrichedResponse) GetAssociatedItemsCountOk() (*int32, bool)`

GetAssociatedItemsCountOk returns a tuple with the AssociatedItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedItemsCount

`func (o *OrganizationAnnotationsGroupEnrichedResponse) SetAssociatedItemsCount(v int32)`

SetAssociatedItemsCount sets AssociatedItemsCount field to given value.

### HasAssociatedItemsCount

`func (o *OrganizationAnnotationsGroupEnrichedResponse) HasAssociatedItemsCount() bool`

HasAssociatedItemsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


