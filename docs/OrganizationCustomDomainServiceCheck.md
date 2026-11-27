# OrganizationCustomDomainServiceCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** | Id of the service. | 
**ServiceType** | [**ServiceTypeEnum**](ServiceTypeEnum.md) | Type of the service. Only &#x60;APPLICATION&#x60;, &#x60;CONTAINER&#x60; and &#x60;HELM&#x60; services have custom domains. | 
**Status** | [**NullableCheckedCustomDomainStatus**](CheckedCustomDomainStatus.md) | The result of the check. &#x60;null&#x60; when the check could not run (cluster unreachable, timeout, service never deployed); &#x60;error_details&#x60; then gives the reason. | 
**ErrorDetails** | **NullableString** | Why the check failed or could not run. &#x60;null&#x60; when the check succeeded. | 

## Methods

### NewOrganizationCustomDomainServiceCheck

`func NewOrganizationCustomDomainServiceCheck(serviceId string, serviceType ServiceTypeEnum, status NullableCheckedCustomDomainStatus, errorDetails NullableString, ) *OrganizationCustomDomainServiceCheck`

NewOrganizationCustomDomainServiceCheck instantiates a new OrganizationCustomDomainServiceCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomDomainServiceCheckWithDefaults

`func NewOrganizationCustomDomainServiceCheckWithDefaults() *OrganizationCustomDomainServiceCheck`

NewOrganizationCustomDomainServiceCheckWithDefaults instantiates a new OrganizationCustomDomainServiceCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *OrganizationCustomDomainServiceCheck) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *OrganizationCustomDomainServiceCheck) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *OrganizationCustomDomainServiceCheck) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceType

`func (o *OrganizationCustomDomainServiceCheck) GetServiceType() ServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *OrganizationCustomDomainServiceCheck) GetServiceTypeOk() (*ServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *OrganizationCustomDomainServiceCheck) SetServiceType(v ServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.


### GetStatus

`func (o *OrganizationCustomDomainServiceCheck) GetStatus() CheckedCustomDomainStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationCustomDomainServiceCheck) GetStatusOk() (*CheckedCustomDomainStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationCustomDomainServiceCheck) SetStatus(v CheckedCustomDomainStatus)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *OrganizationCustomDomainServiceCheck) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *OrganizationCustomDomainServiceCheck) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetErrorDetails

`func (o *OrganizationCustomDomainServiceCheck) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *OrganizationCustomDomainServiceCheck) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *OrganizationCustomDomainServiceCheck) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.


### SetErrorDetailsNil

`func (o *OrganizationCustomDomainServiceCheck) SetErrorDetailsNil(b bool)`

 SetErrorDetailsNil sets the value for ErrorDetails to be an explicit nil

### UnsetErrorDetails
`func (o *OrganizationCustomDomainServiceCheck) UnsetErrorDetails()`

UnsetErrorDetails ensures that no value is present for ErrorDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


