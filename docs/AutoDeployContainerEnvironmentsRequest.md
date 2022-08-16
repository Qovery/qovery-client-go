# AutoDeployContainerEnvironmentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageName** | Pointer to **string** | the container image name to deploy | [optional] 
**Tag** | Pointer to **string** | the new tag to deploy | [optional] 

## Methods

### NewAutoDeployContainerEnvironmentsRequest

`func NewAutoDeployContainerEnvironmentsRequest() *AutoDeployContainerEnvironmentsRequest`

NewAutoDeployContainerEnvironmentsRequest instantiates a new AutoDeployContainerEnvironmentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoDeployContainerEnvironmentsRequestWithDefaults

`func NewAutoDeployContainerEnvironmentsRequestWithDefaults() *AutoDeployContainerEnvironmentsRequest`

NewAutoDeployContainerEnvironmentsRequestWithDefaults instantiates a new AutoDeployContainerEnvironmentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageName

`func (o *AutoDeployContainerEnvironmentsRequest) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *AutoDeployContainerEnvironmentsRequest) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *AutoDeployContainerEnvironmentsRequest) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *AutoDeployContainerEnvironmentsRequest) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetTag

`func (o *AutoDeployContainerEnvironmentsRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AutoDeployContainerEnvironmentsRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AutoDeployContainerEnvironmentsRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AutoDeployContainerEnvironmentsRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


