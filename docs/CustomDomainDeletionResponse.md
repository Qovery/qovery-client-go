# CustomDomainDeletionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImpactedServices** | [**[]ImpactedServiceResponse**](ImpactedServiceResponse.md) | The services that lost a custom domain. They need a redeploy to stop serving it. | 

## Methods

### NewCustomDomainDeletionResponse

`func NewCustomDomainDeletionResponse(impactedServices []ImpactedServiceResponse, ) *CustomDomainDeletionResponse`

NewCustomDomainDeletionResponse instantiates a new CustomDomainDeletionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainDeletionResponseWithDefaults

`func NewCustomDomainDeletionResponseWithDefaults() *CustomDomainDeletionResponse`

NewCustomDomainDeletionResponseWithDefaults instantiates a new CustomDomainDeletionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpactedServices

`func (o *CustomDomainDeletionResponse) GetImpactedServices() []ImpactedServiceResponse`

GetImpactedServices returns the ImpactedServices field if non-nil, zero value otherwise.

### GetImpactedServicesOk

`func (o *CustomDomainDeletionResponse) GetImpactedServicesOk() (*[]ImpactedServiceResponse, bool)`

GetImpactedServicesOk returns a tuple with the ImpactedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedServices

`func (o *CustomDomainDeletionResponse) SetImpactedServices(v []ImpactedServiceResponse)`

SetImpactedServices sets ImpactedServices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


