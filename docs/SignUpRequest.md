# SignUpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**UserEmail** | **string** |  | 
**TypeOfUse** | [**TypeOfUseEnum**](TypeOfUseEnum.md) |  | 
**QoveryUsage** | **string** |  | 
**CompanyName** | Pointer to **NullableString** |  | [optional] 
**CompanySize** | Pointer to [**CompanySizeEnum**](CompanySizeEnum.md) |  | [optional] 
**UserRole** | Pointer to **NullableString** |  | [optional] 
**QoveryUsageOther** | Pointer to **NullableString** |  | [optional] 
**UserQuestions** | Pointer to **NullableString** |  | [optional] 
**CurrentStep** | Pointer to **NullableString** |  | [optional] 
**DxAuth** | Pointer to **NullableBool** |  | [optional] 

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

### SetCompanyNameNil

`func (o *SignUpRequest) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *SignUpRequest) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
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

### SetUserRoleNil

`func (o *SignUpRequest) SetUserRoleNil(b bool)`

 SetUserRoleNil sets the value for UserRole to be an explicit nil

### UnsetUserRole
`func (o *SignUpRequest) UnsetUserRole()`

UnsetUserRole ensures that no value is present for UserRole, not even an explicit nil
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

### SetQoveryUsageOtherNil

`func (o *SignUpRequest) SetQoveryUsageOtherNil(b bool)`

 SetQoveryUsageOtherNil sets the value for QoveryUsageOther to be an explicit nil

### UnsetQoveryUsageOther
`func (o *SignUpRequest) UnsetQoveryUsageOther()`

UnsetQoveryUsageOther ensures that no value is present for QoveryUsageOther, not even an explicit nil
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

### SetUserQuestionsNil

`func (o *SignUpRequest) SetUserQuestionsNil(b bool)`

 SetUserQuestionsNil sets the value for UserQuestions to be an explicit nil

### UnsetUserQuestions
`func (o *SignUpRequest) UnsetUserQuestions()`

UnsetUserQuestions ensures that no value is present for UserQuestions, not even an explicit nil
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

### SetCurrentStepNil

`func (o *SignUpRequest) SetCurrentStepNil(b bool)`

 SetCurrentStepNil sets the value for CurrentStep to be an explicit nil

### UnsetCurrentStep
`func (o *SignUpRequest) UnsetCurrentStep()`

UnsetCurrentStep ensures that no value is present for CurrentStep, not even an explicit nil
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

### SetDxAuthNil

`func (o *SignUpRequest) SetDxAuthNil(b bool)`

 SetDxAuthNil sets the value for DxAuth to be an explicit nil

### UnsetDxAuth
`func (o *SignUpRequest) UnsetDxAuth()`

UnsetDxAuth ensures that no value is present for DxAuth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


