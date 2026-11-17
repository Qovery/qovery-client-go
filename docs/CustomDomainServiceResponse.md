# CustomDomainServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the service. | 
**Type** | [**ServiceTypeEnum**](ServiceTypeEnum.md) | Type of the service. Only &#x60;APPLICATION&#x60;, &#x60;CONTAINER&#x60; and &#x60;HELM&#x60; services have custom domains. | 
**Name** | **string** | Name of the service. | 

## Methods

### NewCustomDomainServiceResponse

`func NewCustomDomainServiceResponse(id string, type_ ServiceTypeEnum, name string, ) *CustomDomainServiceResponse`

NewCustomDomainServiceResponse instantiates a new CustomDomainServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainServiceResponseWithDefaults

`func NewCustomDomainServiceResponseWithDefaults() *CustomDomainServiceResponse`

NewCustomDomainServiceResponseWithDefaults instantiates a new CustomDomainServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomDomainServiceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomDomainServiceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomDomainServiceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CustomDomainServiceResponse) GetType() ServiceTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomDomainServiceResponse) GetTypeOk() (*ServiceTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomDomainServiceResponse) SetType(v ServiceTypeEnum)`

SetType sets Type field to given value.


### GetName

`func (o *CustomDomainServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomDomainServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomDomainServiceResponse) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


