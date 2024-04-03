# OrganizationAnnotationsGroupCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the annotations group | 
**Annotations** | [**[]Annotation**](Annotation.md) |  | 
**Scopes** | [**[]OrganizationAnnotationsGroupScopeEnum**](OrganizationAnnotationsGroupScopeEnum.md) |  | 

## Methods

### NewOrganizationAnnotationsGroupCreateRequest

`func NewOrganizationAnnotationsGroupCreateRequest(name string, annotations []Annotation, scopes []OrganizationAnnotationsGroupScopeEnum, ) *OrganizationAnnotationsGroupCreateRequest`

NewOrganizationAnnotationsGroupCreateRequest instantiates a new OrganizationAnnotationsGroupCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationAnnotationsGroupCreateRequestWithDefaults

`func NewOrganizationAnnotationsGroupCreateRequestWithDefaults() *OrganizationAnnotationsGroupCreateRequest`

NewOrganizationAnnotationsGroupCreateRequestWithDefaults instantiates a new OrganizationAnnotationsGroupCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationAnnotationsGroupCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationAnnotationsGroupCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationAnnotationsGroupCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *OrganizationAnnotationsGroupCreateRequest) GetAnnotations() []Annotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *OrganizationAnnotationsGroupCreateRequest) GetAnnotationsOk() (*[]Annotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *OrganizationAnnotationsGroupCreateRequest) SetAnnotations(v []Annotation)`

SetAnnotations sets Annotations field to given value.


### GetScopes

`func (o *OrganizationAnnotationsGroupCreateRequest) GetScopes() []OrganizationAnnotationsGroupScopeEnum`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OrganizationAnnotationsGroupCreateRequest) GetScopesOk() (*[]OrganizationAnnotationsGroupScopeEnum, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OrganizationAnnotationsGroupCreateRequest) SetScopes(v []OrganizationAnnotationsGroupScopeEnum)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


