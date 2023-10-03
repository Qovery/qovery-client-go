# ContainerRegistryProviderDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Url** | **string** | URL of the container registry | 
**Kind** | [**ContainerRegistryKindEnum**](ContainerRegistryKindEnum.md) |  | 

## Methods

### NewContainerRegistryProviderDetailsResponse

`func NewContainerRegistryProviderDetailsResponse(id string, name string, url string, kind ContainerRegistryKindEnum, ) *ContainerRegistryProviderDetailsResponse`

NewContainerRegistryProviderDetailsResponse instantiates a new ContainerRegistryProviderDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistryProviderDetailsResponseWithDefaults

`func NewContainerRegistryProviderDetailsResponseWithDefaults() *ContainerRegistryProviderDetailsResponse`

NewContainerRegistryProviderDetailsResponseWithDefaults instantiates a new ContainerRegistryProviderDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerRegistryProviderDetailsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerRegistryProviderDetailsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerRegistryProviderDetailsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ContainerRegistryProviderDetailsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerRegistryProviderDetailsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerRegistryProviderDetailsResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *ContainerRegistryProviderDetailsResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContainerRegistryProviderDetailsResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContainerRegistryProviderDetailsResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetKind

`func (o *ContainerRegistryProviderDetailsResponse) GetKind() ContainerRegistryKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ContainerRegistryProviderDetailsResponse) GetKindOk() (*ContainerRegistryKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ContainerRegistryProviderDetailsResponse) SetKind(v ContainerRegistryKindEnum)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


