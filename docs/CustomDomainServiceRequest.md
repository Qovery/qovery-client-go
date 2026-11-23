# CustomDomainServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the service. It must belong to the environment. | 
**Type** | [**ServiceTypeEnum**](ServiceTypeEnum.md) | Type of the service. Only &#x60;APPLICATION&#x60;, &#x60;CONTAINER&#x60; and &#x60;HELM&#x60; services have custom domains. | 

## Methods

### NewCustomDomainServiceRequest

`func NewCustomDomainServiceRequest(id string, type_ ServiceTypeEnum, ) *CustomDomainServiceRequest`

NewCustomDomainServiceRequest instantiates a new CustomDomainServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainServiceRequestWithDefaults

`func NewCustomDomainServiceRequestWithDefaults() *CustomDomainServiceRequest`

NewCustomDomainServiceRequestWithDefaults instantiates a new CustomDomainServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomDomainServiceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomDomainServiceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomDomainServiceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CustomDomainServiceRequest) GetType() ServiceTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomDomainServiceRequest) GetTypeOk() (*ServiceTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomDomainServiceRequest) SetType(v ServiceTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


