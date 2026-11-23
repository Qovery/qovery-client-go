# ServiceDomainQoveryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**ServiceDomainKindEnum**](ServiceDomainKindEnum.md) |  | 
**Domain** | **string** |  | 
**GeneratedHostnames** | [**[]GeneratedHostnameResponse**](GeneratedHostnameResponse.md) |  | 

## Methods

### NewServiceDomainQoveryResponse

`func NewServiceDomainQoveryResponse(kind ServiceDomainKindEnum, domain string, generatedHostnames []GeneratedHostnameResponse, ) *ServiceDomainQoveryResponse`

NewServiceDomainQoveryResponse instantiates a new ServiceDomainQoveryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDomainQoveryResponseWithDefaults

`func NewServiceDomainQoveryResponseWithDefaults() *ServiceDomainQoveryResponse`

NewServiceDomainQoveryResponseWithDefaults instantiates a new ServiceDomainQoveryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ServiceDomainQoveryResponse) GetKind() ServiceDomainKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServiceDomainQoveryResponse) GetKindOk() (*ServiceDomainKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServiceDomainQoveryResponse) SetKind(v ServiceDomainKindEnum)`

SetKind sets Kind field to given value.


### GetDomain

`func (o *ServiceDomainQoveryResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ServiceDomainQoveryResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ServiceDomainQoveryResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetGeneratedHostnames

`func (o *ServiceDomainQoveryResponse) GetGeneratedHostnames() []GeneratedHostnameResponse`

GetGeneratedHostnames returns the GeneratedHostnames field if non-nil, zero value otherwise.

### GetGeneratedHostnamesOk

`func (o *ServiceDomainQoveryResponse) GetGeneratedHostnamesOk() (*[]GeneratedHostnameResponse, bool)`

GetGeneratedHostnamesOk returns a tuple with the GeneratedHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedHostnames

`func (o *ServiceDomainQoveryResponse) SetGeneratedHostnames(v []GeneratedHostnameResponse)`

SetGeneratedHostnames sets GeneratedHostnames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


