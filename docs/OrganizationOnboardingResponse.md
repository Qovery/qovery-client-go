# OrganizationOnboardingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseCases** | Pointer to **NullableString** | Comma-separated list of use cases selected by the organization owner at sign-up (e.g. \&quot;ephemeral-environments,rde\&quot;) | [optional] 
**Status** | Pointer to [**OrganizationOnboardingStatusEnum**](OrganizationOnboardingStatusEnum.md) |  | [optional] 

## Methods

### NewOrganizationOnboardingResponse

`func NewOrganizationOnboardingResponse() *OrganizationOnboardingResponse`

NewOrganizationOnboardingResponse instantiates a new OrganizationOnboardingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationOnboardingResponseWithDefaults

`func NewOrganizationOnboardingResponseWithDefaults() *OrganizationOnboardingResponse`

NewOrganizationOnboardingResponseWithDefaults instantiates a new OrganizationOnboardingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseCases

`func (o *OrganizationOnboardingResponse) GetUseCases() string`

GetUseCases returns the UseCases field if non-nil, zero value otherwise.

### GetUseCasesOk

`func (o *OrganizationOnboardingResponse) GetUseCasesOk() (*string, bool)`

GetUseCasesOk returns a tuple with the UseCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCases

`func (o *OrganizationOnboardingResponse) SetUseCases(v string)`

SetUseCases sets UseCases field to given value.

### HasUseCases

`func (o *OrganizationOnboardingResponse) HasUseCases() bool`

HasUseCases returns a boolean if a field has been set.

### SetUseCasesNil

`func (o *OrganizationOnboardingResponse) SetUseCasesNil(b bool)`

 SetUseCasesNil sets the value for UseCases to be an explicit nil

### UnsetUseCases
`func (o *OrganizationOnboardingResponse) UnsetUseCases()`

UnsetUseCases ensures that no value is present for UseCases, not even an explicit nil
### GetStatus

`func (o *OrganizationOnboardingResponse) GetStatus() OrganizationOnboardingStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationOnboardingResponse) GetStatusOk() (*OrganizationOnboardingStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationOnboardingResponse) SetStatus(v OrganizationOnboardingStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrganizationOnboardingResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


