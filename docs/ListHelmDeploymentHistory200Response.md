# ListHelmDeploymentHistory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **float32** |  | 
**PageSize** | **float32** |  | 
**Results** | Pointer to [**[]ListHelmDeploymentHistory200ResponseAllOfResultsInner**](ListHelmDeploymentHistory200ResponseAllOfResultsInner.md) |  | [optional] 

## Methods

### NewListHelmDeploymentHistory200Response

`func NewListHelmDeploymentHistory200Response(page float32, pageSize float32, ) *ListHelmDeploymentHistory200Response`

NewListHelmDeploymentHistory200Response instantiates a new ListHelmDeploymentHistory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHelmDeploymentHistory200ResponseWithDefaults

`func NewListHelmDeploymentHistory200ResponseWithDefaults() *ListHelmDeploymentHistory200Response`

NewListHelmDeploymentHistory200ResponseWithDefaults instantiates a new ListHelmDeploymentHistory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListHelmDeploymentHistory200Response) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListHelmDeploymentHistory200Response) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListHelmDeploymentHistory200Response) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *ListHelmDeploymentHistory200Response) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListHelmDeploymentHistory200Response) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListHelmDeploymentHistory200Response) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *ListHelmDeploymentHistory200Response) GetResults() []ListHelmDeploymentHistory200ResponseAllOfResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListHelmDeploymentHistory200Response) GetResultsOk() (*[]ListHelmDeploymentHistory200ResponseAllOfResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListHelmDeploymentHistory200Response) SetResults(v []ListHelmDeploymentHistory200ResponseAllOfResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListHelmDeploymentHistory200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


