# OrganizationLabelsGroupEnrichedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Labels** | [**[]Label**](Label.md) |  | 
**AssociatedItemsCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewOrganizationLabelsGroupEnrichedResponse

`func NewOrganizationLabelsGroupEnrichedResponse(id string, createdAt time.Time, name string, labels []Label, ) *OrganizationLabelsGroupEnrichedResponse`

NewOrganizationLabelsGroupEnrichedResponse instantiates a new OrganizationLabelsGroupEnrichedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationLabelsGroupEnrichedResponseWithDefaults

`func NewOrganizationLabelsGroupEnrichedResponseWithDefaults() *OrganizationLabelsGroupEnrichedResponse`

NewOrganizationLabelsGroupEnrichedResponseWithDefaults instantiates a new OrganizationLabelsGroupEnrichedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationLabelsGroupEnrichedResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationLabelsGroupEnrichedResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationLabelsGroupEnrichedResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrganizationLabelsGroupEnrichedResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationLabelsGroupEnrichedResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationLabelsGroupEnrichedResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationLabelsGroupEnrichedResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationLabelsGroupEnrichedResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationLabelsGroupEnrichedResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationLabelsGroupEnrichedResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *OrganizationLabelsGroupEnrichedResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationLabelsGroupEnrichedResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationLabelsGroupEnrichedResponse) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *OrganizationLabelsGroupEnrichedResponse) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OrganizationLabelsGroupEnrichedResponse) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OrganizationLabelsGroupEnrichedResponse) SetLabels(v []Label)`

SetLabels sets Labels field to given value.


### GetAssociatedItemsCount

`func (o *OrganizationLabelsGroupEnrichedResponse) GetAssociatedItemsCount() int32`

GetAssociatedItemsCount returns the AssociatedItemsCount field if non-nil, zero value otherwise.

### GetAssociatedItemsCountOk

`func (o *OrganizationLabelsGroupEnrichedResponse) GetAssociatedItemsCountOk() (*int32, bool)`

GetAssociatedItemsCountOk returns a tuple with the AssociatedItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedItemsCount

`func (o *OrganizationLabelsGroupEnrichedResponse) SetAssociatedItemsCount(v int32)`

SetAssociatedItemsCount sets AssociatedItemsCount field to given value.

### HasAssociatedItemsCount

`func (o *OrganizationLabelsGroupEnrichedResponse) HasAssociatedItemsCount() bool`

HasAssociatedItemsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


