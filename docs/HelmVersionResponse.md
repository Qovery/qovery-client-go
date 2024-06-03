# HelmVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | Pointer to **string** |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHelmVersionResponse

`func NewHelmVersionResponse() *HelmVersionResponse`

NewHelmVersionResponse instantiates a new HelmVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmVersionResponseWithDefaults

`func NewHelmVersionResponseWithDefaults() *HelmVersionResponse`

NewHelmVersionResponseWithDefaults instantiates a new HelmVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *HelmVersionResponse) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *HelmVersionResponse) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *HelmVersionResponse) SetChartName(v string)`

SetChartName sets ChartName field to given value.

### HasChartName

`func (o *HelmVersionResponse) HasChartName() bool`

HasChartName returns a boolean if a field has been set.

### GetVersions

`func (o *HelmVersionResponse) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *HelmVersionResponse) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *HelmVersionResponse) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *HelmVersionResponse) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


