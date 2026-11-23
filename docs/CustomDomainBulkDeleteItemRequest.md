# CustomDomainBulkDeleteItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | Id of the environment of the domain. It must belong to the organization of the call. | 
**Domain** | **string** | The domain name. Every row of this name in the environment is deleted. | 

## Methods

### NewCustomDomainBulkDeleteItemRequest

`func NewCustomDomainBulkDeleteItemRequest(environmentId string, domain string, ) *CustomDomainBulkDeleteItemRequest`

NewCustomDomainBulkDeleteItemRequest instantiates a new CustomDomainBulkDeleteItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainBulkDeleteItemRequestWithDefaults

`func NewCustomDomainBulkDeleteItemRequestWithDefaults() *CustomDomainBulkDeleteItemRequest`

NewCustomDomainBulkDeleteItemRequestWithDefaults instantiates a new CustomDomainBulkDeleteItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *CustomDomainBulkDeleteItemRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CustomDomainBulkDeleteItemRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CustomDomainBulkDeleteItemRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetDomain

`func (o *CustomDomainBulkDeleteItemRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CustomDomainBulkDeleteItemRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CustomDomainBulkDeleteItemRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


