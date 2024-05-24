# OrganizationLabelsGroupCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the labels group | 
**Labels** | [**[]Label**](Label.md) |  | 

## Methods

### NewOrganizationLabelsGroupCreateRequest

`func NewOrganizationLabelsGroupCreateRequest(name string, labels []Label, ) *OrganizationLabelsGroupCreateRequest`

NewOrganizationLabelsGroupCreateRequest instantiates a new OrganizationLabelsGroupCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationLabelsGroupCreateRequestWithDefaults

`func NewOrganizationLabelsGroupCreateRequestWithDefaults() *OrganizationLabelsGroupCreateRequest`

NewOrganizationLabelsGroupCreateRequestWithDefaults instantiates a new OrganizationLabelsGroupCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationLabelsGroupCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationLabelsGroupCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationLabelsGroupCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *OrganizationLabelsGroupCreateRequest) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OrganizationLabelsGroupCreateRequest) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OrganizationLabelsGroupCreateRequest) SetLabels(v []Label)`

SetLabels sets Labels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


