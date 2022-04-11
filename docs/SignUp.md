# SignUp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
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

### NewSignUp

`func NewSignUp(id string, createdAt time.Time, firstName string, lastName string, userEmail string, typeOfUse TypeOfUseEnum, qoveryUsage string, ) *SignUp`

NewSignUp instantiates a new SignUp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignUpWithDefaults

`func NewSignUpWithDefaults() *SignUp`

NewSignUpWithDefaults instantiates a new SignUp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SignUp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignUp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignUp) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *SignUp) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SignUp) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SignUp) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SignUp) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SignUp) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SignUp) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SignUp) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFirstName

`func (o *SignUp) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SignUp) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SignUp) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *SignUp) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SignUp) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SignUp) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetUserEmail

`func (o *SignUp) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *SignUp) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *SignUp) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetTypeOfUse

`func (o *SignUp) GetTypeOfUse() TypeOfUseEnum`

GetTypeOfUse returns the TypeOfUse field if non-nil, zero value otherwise.

### GetTypeOfUseOk

`func (o *SignUp) GetTypeOfUseOk() (*TypeOfUseEnum, bool)`

GetTypeOfUseOk returns a tuple with the TypeOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfUse

`func (o *SignUp) SetTypeOfUse(v TypeOfUseEnum)`

SetTypeOfUse sets TypeOfUse field to given value.


### GetQoveryUsage

`func (o *SignUp) GetQoveryUsage() string`

GetQoveryUsage returns the QoveryUsage field if non-nil, zero value otherwise.

### GetQoveryUsageOk

`func (o *SignUp) GetQoveryUsageOk() (*string, bool)`

GetQoveryUsageOk returns a tuple with the QoveryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryUsage

`func (o *SignUp) SetQoveryUsage(v string)`

SetQoveryUsage sets QoveryUsage field to given value.


### GetCompanyName

`func (o *SignUp) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *SignUp) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *SignUp) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *SignUp) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *SignUp) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *SignUp) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetCompanySize

`func (o *SignUp) GetCompanySize() CompanySizeEnum`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *SignUp) GetCompanySizeOk() (*CompanySizeEnum, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *SignUp) SetCompanySize(v CompanySizeEnum)`

SetCompanySize sets CompanySize field to given value.

### HasCompanySize

`func (o *SignUp) HasCompanySize() bool`

HasCompanySize returns a boolean if a field has been set.

### GetUserRole

`func (o *SignUp) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *SignUp) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *SignUp) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *SignUp) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### SetUserRoleNil

`func (o *SignUp) SetUserRoleNil(b bool)`

 SetUserRoleNil sets the value for UserRole to be an explicit nil

### UnsetUserRole
`func (o *SignUp) UnsetUserRole()`

UnsetUserRole ensures that no value is present for UserRole, not even an explicit nil
### GetQoveryUsageOther

`func (o *SignUp) GetQoveryUsageOther() string`

GetQoveryUsageOther returns the QoveryUsageOther field if non-nil, zero value otherwise.

### GetQoveryUsageOtherOk

`func (o *SignUp) GetQoveryUsageOtherOk() (*string, bool)`

GetQoveryUsageOtherOk returns a tuple with the QoveryUsageOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryUsageOther

`func (o *SignUp) SetQoveryUsageOther(v string)`

SetQoveryUsageOther sets QoveryUsageOther field to given value.

### HasQoveryUsageOther

`func (o *SignUp) HasQoveryUsageOther() bool`

HasQoveryUsageOther returns a boolean if a field has been set.

### SetQoveryUsageOtherNil

`func (o *SignUp) SetQoveryUsageOtherNil(b bool)`

 SetQoveryUsageOtherNil sets the value for QoveryUsageOther to be an explicit nil

### UnsetQoveryUsageOther
`func (o *SignUp) UnsetQoveryUsageOther()`

UnsetQoveryUsageOther ensures that no value is present for QoveryUsageOther, not even an explicit nil
### GetUserQuestions

`func (o *SignUp) GetUserQuestions() string`

GetUserQuestions returns the UserQuestions field if non-nil, zero value otherwise.

### GetUserQuestionsOk

`func (o *SignUp) GetUserQuestionsOk() (*string, bool)`

GetUserQuestionsOk returns a tuple with the UserQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuestions

`func (o *SignUp) SetUserQuestions(v string)`

SetUserQuestions sets UserQuestions field to given value.

### HasUserQuestions

`func (o *SignUp) HasUserQuestions() bool`

HasUserQuestions returns a boolean if a field has been set.

### SetUserQuestionsNil

`func (o *SignUp) SetUserQuestionsNil(b bool)`

 SetUserQuestionsNil sets the value for UserQuestions to be an explicit nil

### UnsetUserQuestions
`func (o *SignUp) UnsetUserQuestions()`

UnsetUserQuestions ensures that no value is present for UserQuestions, not even an explicit nil
### GetCurrentStep

`func (o *SignUp) GetCurrentStep() string`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *SignUp) GetCurrentStepOk() (*string, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *SignUp) SetCurrentStep(v string)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *SignUp) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### SetCurrentStepNil

`func (o *SignUp) SetCurrentStepNil(b bool)`

 SetCurrentStepNil sets the value for CurrentStep to be an explicit nil

### UnsetCurrentStep
`func (o *SignUp) UnsetCurrentStep()`

UnsetCurrentStep ensures that no value is present for CurrentStep, not even an explicit nil
### GetDxAuth

`func (o *SignUp) GetDxAuth() bool`

GetDxAuth returns the DxAuth field if non-nil, zero value otherwise.

### GetDxAuthOk

`func (o *SignUp) GetDxAuthOk() (*bool, bool)`

GetDxAuthOk returns a tuple with the DxAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDxAuth

`func (o *SignUp) SetDxAuth(v bool)`

SetDxAuth sets DxAuth field to given value.

### HasDxAuth

`func (o *SignUp) HasDxAuth() bool`

HasDxAuth returns a boolean if a field has been set.

### SetDxAuthNil

`func (o *SignUp) SetDxAuthNil(b bool)`

 SetDxAuthNil sets the value for DxAuth to be an explicit nil

### UnsetDxAuth
`func (o *SignUp) UnsetDxAuth()`

UnsetDxAuth ensures that no value is present for DxAuth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


