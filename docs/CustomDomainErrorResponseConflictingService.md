# CustomDomainErrorResponseConflictingService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the service. | 
**Type** | [**ServiceTypeEnum**](ServiceTypeEnum.md) | Type of the service. Only &#x60;APPLICATION&#x60;, &#x60;CONTAINER&#x60; and &#x60;HELM&#x60; services have custom domains. | 
**Name** | **string** | Name of the service. | 

## Methods

### NewCustomDomainErrorResponseConflictingService

`func NewCustomDomainErrorResponseConflictingService(id string, type_ ServiceTypeEnum, name string, ) *CustomDomainErrorResponseConflictingService`

NewCustomDomainErrorResponseConflictingService instantiates a new CustomDomainErrorResponseConflictingService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainErrorResponseConflictingServiceWithDefaults

`func NewCustomDomainErrorResponseConflictingServiceWithDefaults() *CustomDomainErrorResponseConflictingService`

NewCustomDomainErrorResponseConflictingServiceWithDefaults instantiates a new CustomDomainErrorResponseConflictingService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomDomainErrorResponseConflictingService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomDomainErrorResponseConflictingService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomDomainErrorResponseConflictingService) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CustomDomainErrorResponseConflictingService) GetType() ServiceTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomDomainErrorResponseConflictingService) GetTypeOk() (*ServiceTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomDomainErrorResponseConflictingService) SetType(v ServiceTypeEnum)`

SetType sets Type field to given value.


### GetName

`func (o *CustomDomainErrorResponseConflictingService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomDomainErrorResponseConflictingService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomDomainErrorResponseConflictingService) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


