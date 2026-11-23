# ImpactedServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**ServiceTypeEnum**](ServiceTypeEnum.md) |  | 

## Methods

### NewImpactedServiceResponse

`func NewImpactedServiceResponse(id string, type_ ServiceTypeEnum, ) *ImpactedServiceResponse`

NewImpactedServiceResponse instantiates a new ImpactedServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImpactedServiceResponseWithDefaults

`func NewImpactedServiceResponseWithDefaults() *ImpactedServiceResponse`

NewImpactedServiceResponseWithDefaults instantiates a new ImpactedServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImpactedServiceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImpactedServiceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImpactedServiceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ImpactedServiceResponse) GetType() ServiceTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImpactedServiceResponse) GetTypeOk() (*ServiceTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImpactedServiceResponse) SetType(v ServiceTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


