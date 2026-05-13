# ArgocdAppManifestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManifestRevision** | Pointer to **NullableString** | Git revision the app is currently synced to | [optional] 
**ManifestMetadata** | [**ArgocdAppManifestResponseManifestMetadata**](ArgocdAppManifestResponseManifestMetadata.md) |  | 

## Methods

### NewArgocdAppManifestResponse

`func NewArgocdAppManifestResponse(manifestMetadata ArgocdAppManifestResponseManifestMetadata, ) *ArgocdAppManifestResponse`

NewArgocdAppManifestResponse instantiates a new ArgocdAppManifestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgocdAppManifestResponseWithDefaults

`func NewArgocdAppManifestResponseWithDefaults() *ArgocdAppManifestResponse`

NewArgocdAppManifestResponseWithDefaults instantiates a new ArgocdAppManifestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManifestRevision

`func (o *ArgocdAppManifestResponse) GetManifestRevision() string`

GetManifestRevision returns the ManifestRevision field if non-nil, zero value otherwise.

### GetManifestRevisionOk

`func (o *ArgocdAppManifestResponse) GetManifestRevisionOk() (*string, bool)`

GetManifestRevisionOk returns a tuple with the ManifestRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestRevision

`func (o *ArgocdAppManifestResponse) SetManifestRevision(v string)`

SetManifestRevision sets ManifestRevision field to given value.

### HasManifestRevision

`func (o *ArgocdAppManifestResponse) HasManifestRevision() bool`

HasManifestRevision returns a boolean if a field has been set.

### SetManifestRevisionNil

`func (o *ArgocdAppManifestResponse) SetManifestRevisionNil(b bool)`

 SetManifestRevisionNil sets the value for ManifestRevision to be an explicit nil

### UnsetManifestRevision
`func (o *ArgocdAppManifestResponse) UnsetManifestRevision()`

UnsetManifestRevision ensures that no value is present for ManifestRevision, not even an explicit nil
### GetManifestMetadata

`func (o *ArgocdAppManifestResponse) GetManifestMetadata() ArgocdAppManifestResponseManifestMetadata`

GetManifestMetadata returns the ManifestMetadata field if non-nil, zero value otherwise.

### GetManifestMetadataOk

`func (o *ArgocdAppManifestResponse) GetManifestMetadataOk() (*ArgocdAppManifestResponseManifestMetadata, bool)`

GetManifestMetadataOk returns a tuple with the ManifestMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestMetadata

`func (o *ArgocdAppManifestResponse) SetManifestMetadata(v ArgocdAppManifestResponseManifestMetadata)`

SetManifestMetadata sets ManifestMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


