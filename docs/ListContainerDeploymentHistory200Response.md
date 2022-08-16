# ListContainerDeploymentHistory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **float32** |  | 
**PageSize** | **float32** |  | 
**Results** | Pointer to [**[]ListContainerDeploymentHistory200ResponseAllOfResultsInner**](ListContainerDeploymentHistory200ResponseAllOfResultsInner.md) |  | [optional] 

## Methods

### NewListContainerDeploymentHistory200Response

`func NewListContainerDeploymentHistory200Response(page float32, pageSize float32, ) *ListContainerDeploymentHistory200Response`

NewListContainerDeploymentHistory200Response instantiates a new ListContainerDeploymentHistory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListContainerDeploymentHistory200ResponseWithDefaults

`func NewListContainerDeploymentHistory200ResponseWithDefaults() *ListContainerDeploymentHistory200Response`

NewListContainerDeploymentHistory200ResponseWithDefaults instantiates a new ListContainerDeploymentHistory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListContainerDeploymentHistory200Response) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListContainerDeploymentHistory200Response) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListContainerDeploymentHistory200Response) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *ListContainerDeploymentHistory200Response) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListContainerDeploymentHistory200Response) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListContainerDeploymentHistory200Response) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.


### GetResults

`func (o *ListContainerDeploymentHistory200Response) GetResults() []ListContainerDeploymentHistory200ResponseAllOfResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListContainerDeploymentHistory200Response) GetResultsOk() (*[]ListContainerDeploymentHistory200ResponseAllOfResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListContainerDeploymentHistory200Response) SetResults(v []ListContainerDeploymentHistory200ResponseAllOfResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListContainerDeploymentHistory200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


