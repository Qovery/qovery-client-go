# BlueprintManifestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]BlueprintManifestResponseResultsInner**](BlueprintManifestResponseResultsInner.md) | Form fields the console renders for the user (variables + auto-sourced context variables). | 
**Engine** | Pointer to [**BlueprintManifestEngineSpec**](BlueprintManifestEngineSpec.md) |  | [optional] 

## Methods

### NewBlueprintManifestResponse

`func NewBlueprintManifestResponse(results []BlueprintManifestResponseResultsInner, ) *BlueprintManifestResponse`

NewBlueprintManifestResponse instantiates a new BlueprintManifestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintManifestResponseWithDefaults

`func NewBlueprintManifestResponseWithDefaults() *BlueprintManifestResponse`

NewBlueprintManifestResponseWithDefaults instantiates a new BlueprintManifestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *BlueprintManifestResponse) GetResults() []BlueprintManifestResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BlueprintManifestResponse) GetResultsOk() (*[]BlueprintManifestResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BlueprintManifestResponse) SetResults(v []BlueprintManifestResponseResultsInner)`

SetResults sets Results field to given value.


### GetEngine

`func (o *BlueprintManifestResponse) GetEngine() BlueprintManifestEngineSpec`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *BlueprintManifestResponse) GetEngineOk() (*BlueprintManifestEngineSpec, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *BlueprintManifestResponse) SetEngine(v BlueprintManifestEngineSpec)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *BlueprintManifestResponse) HasEngine() bool`

HasEngine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


