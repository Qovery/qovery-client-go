# SignUpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**UserEmail** | **string** |  | 
**TypeOfUse** | [**TypeOfUseEnum**](TypeOfUseEnum.md) |  | 
**QoveryUsage** | **string** |  | 
**CompanyName** | Pointer to **string** |  | [optional] 
**CompanySize** | Pointer to [**CompanySizeEnum**](CompanySizeEnum.md) |  | [optional] 
**UserRole** | Pointer to **string** |  | [optional] 
**QoveryUsageOther** | Pointer to **string** |  | [optional] 
**UserQuestions** | Pointer to **string** |  | [optional] 
**CurrentStep** | Pointer to **string** |  | [optional] 
**DxAuth** | Pointer to **bool** |  | [optional] 
**InfrastructureHosting** | Pointer to **string** |  | [optional] 

## Methods

### NewSignUpRequest

`func NewSignUpRequest(firstName string, lastName string, userEmail string, typeOfUse TypeOfUseEnum, qoveryUsage string, ) *SignUpRequest`

NewSignUpRequest instantiates a new SignUpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignUpRequestWithDefaults

`func NewSignUpRequestWithDefaults() *SignUpRequest`

NewSignUpRequestWithDefaults instantiates a new SignUpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *SignUpRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SignUpRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SignUpRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *SignUpRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SignUpRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SignUpRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetUserEmail

`func (o *SignUpRequest) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *SignUpRequest) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *SignUpRequest) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetTypeOfUse

`func (o *SignUpRequest) GetTypeOfUse() TypeOfUseEnum`

GetTypeOfUse returns the TypeOfUse field if non-nil, zero value otherwise.

### GetTypeOfUseOk

`func (o *SignUpRequest) GetTypeOfUseOk() (*TypeOfUseEnum, bool)`

GetTypeOfUseOk returns a tuple with the TypeOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfUse

`func (o *SignUpRequest) SetTypeOfUse(v TypeOfUseEnum)`

SetTypeOfUse sets TypeOfUse field to given value.


### GetQoveryUsage

`func (o *SignUpRequest) GetQoveryUsage() string`

GetQoveryUsage returns the QoveryUsage field if non-nil, zero value otherwise.

### GetQoveryUsageOk

`func (o *SignUpRequest) GetQoveryUsageOk() (*string, bool)`

GetQoveryUsageOk returns a tuple with the QoveryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryUsage

`func (o *SignUpRequest) SetQoveryUsage(v string)`

SetQoveryUsage sets QoveryUsage field to given value.


### GetCompanyName

`func (o *SignUpRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *SignUpRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *SignUpRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *SignUpRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompanySize

`func (o *SignUpRequest) GetCompanySize() CompanySizeEnum`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *SignUpRequest) GetCompanySizeOk() (*CompanySizeEnum, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *SignUpRequest) SetCompanySize(v CompanySizeEnum)`

SetCompanySize sets CompanySize field to given value.

### HasCompanySize

`func (o *SignUpRequest) HasCompanySize() bool`

HasCompanySize returns a boolean if a field has been set.

### GetUserRole

`func (o *SignUpRequest) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *SignUpRequest) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *SignUpRequest) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *SignUpRequest) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### GetQoveryUsageOther

`func (o *SignUpRequest) GetQoveryUsageOther() string`

GetQoveryUsageOther returns the QoveryUsageOther field if non-nil, zero value otherwise.

### GetQoveryUsageOtherOk

`func (o *SignUpRequest) GetQoveryUsageOtherOk() (*string, bool)`

GetQoveryUsageOtherOk returns a tuple with the QoveryUsageOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryUsageOther

`func (o *SignUpRequest) SetQoveryUsageOther(v string)`

SetQoveryUsageOther sets QoveryUsageOther field to given value.

### HasQoveryUsageOther

`func (o *SignUpRequest) HasQoveryUsageOther() bool`

HasQoveryUsageOther returns a boolean if a field has been set.

### GetUserQuestions

`func (o *SignUpRequest) GetUserQuestions() string`

GetUserQuestions returns the UserQuestions field if non-nil, zero value otherwise.

### GetUserQuestionsOk

`func (o *SignUpRequest) GetUserQuestionsOk() (*string, bool)`

GetUserQuestionsOk returns a tuple with the UserQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuestions

`func (o *SignUpRequest) SetUserQuestions(v string)`

SetUserQuestions sets UserQuestions field to given value.

### HasUserQuestions

`func (o *SignUpRequest) HasUserQuestions() bool`

HasUserQuestions returns a boolean if a field has been set.

### GetCurrentStep

`func (o *SignUpRequest) GetCurrentStep() string`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *SignUpRequest) GetCurrentStepOk() (*string, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *SignUpRequest) SetCurrentStep(v string)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *SignUpRequest) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### GetDxAuth

`func (o *SignUpRequest) GetDxAuth() bool`

GetDxAuth returns the DxAuth field if non-nil, zero value otherwise.

### GetDxAuthOk

`func (o *SignUpRequest) GetDxAuthOk() (*bool, bool)`

GetDxAuthOk returns a tuple with the DxAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDxAuth

`func (o *SignUpRequest) SetDxAuth(v bool)`

SetDxAuth sets DxAuth field to given value.

### HasDxAuth

`func (o *SignUpRequest) HasDxAuth() bool`

HasDxAuth returns a boolean if a field has been set.

### GetInfrastructureHosting

`func (o *SignUpRequest) GetInfrastructureHosting() string`

GetInfrastructureHosting returns the InfrastructureHosting field if non-nil, zero value otherwise.

### GetInfrastructureHostingOk

`func (o *SignUpRequest) GetInfrastructureHostingOk() (*string, bool)`

GetInfrastructureHostingOk returns a tuple with the InfrastructureHosting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureHosting

`func (o *SignUpRequest) SetInfrastructureHosting(v string)`

SetInfrastructureHosting sets InfrastructureHosting field to given value.

### HasInfrastructureHosting

`func (o *SignUpRequest) HasInfrastructureHosting() bool`

HasInfrastructureHosting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


