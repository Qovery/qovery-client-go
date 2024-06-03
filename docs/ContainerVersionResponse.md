# ContainerVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageName** | Pointer to **string** |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewContainerVersionResponse

`func NewContainerVersionResponse() *ContainerVersionResponse`

NewContainerVersionResponse instantiates a new ContainerVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerVersionResponseWithDefaults

`func NewContainerVersionResponseWithDefaults() *ContainerVersionResponse`

NewContainerVersionResponseWithDefaults instantiates a new ContainerVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageName

`func (o *ContainerVersionResponse) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ContainerVersionResponse) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ContainerVersionResponse) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ContainerVersionResponse) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetVersions

`func (o *ContainerVersionResponse) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ContainerVersionResponse) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ContainerVersionResponse) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ContainerVersionResponse) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


