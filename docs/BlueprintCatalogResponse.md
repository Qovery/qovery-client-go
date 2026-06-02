# BlueprintCatalogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**GeneratedAt** | **time.Time** |  | 
**Blueprints** | [**[]BlueprintItem**](BlueprintItem.md) |  | 

## Methods

### NewBlueprintCatalogResponse

`func NewBlueprintCatalogResponse(version string, generatedAt time.Time, blueprints []BlueprintItem, ) *BlueprintCatalogResponse`

NewBlueprintCatalogResponse instantiates a new BlueprintCatalogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintCatalogResponseWithDefaults

`func NewBlueprintCatalogResponseWithDefaults() *BlueprintCatalogResponse`

NewBlueprintCatalogResponseWithDefaults instantiates a new BlueprintCatalogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *BlueprintCatalogResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlueprintCatalogResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlueprintCatalogResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetGeneratedAt

`func (o *BlueprintCatalogResponse) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *BlueprintCatalogResponse) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *BlueprintCatalogResponse) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.


### GetBlueprints

`func (o *BlueprintCatalogResponse) GetBlueprints() []BlueprintItem`

GetBlueprints returns the Blueprints field if non-nil, zero value otherwise.

### GetBlueprintsOk

`func (o *BlueprintCatalogResponse) GetBlueprintsOk() (*[]BlueprintItem, bool)`

GetBlueprintsOk returns a tuple with the Blueprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprints

`func (o *BlueprintCatalogResponse) SetBlueprints(v []BlueprintItem)`

SetBlueprints sets Blueprints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


