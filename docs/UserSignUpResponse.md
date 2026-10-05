# UserSignUpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**UserEmail** | Pointer to **NullableString** |  | [optional] 
**TypeOfUse** | Pointer to **NullableString** |  | [optional] 
**CompanyName** | Pointer to **NullableString** |  | [optional] 
**CompanySize** | Pointer to **NullableString** |  | [optional] 
**UserRole** | Pointer to **NullableString** |  | [optional] 
**QoveryUsage** | Pointer to **NullableString** |  | [optional] 
**QoveryUsageOther** | Pointer to **NullableString** |  | [optional] 
**UserQuestions** | Pointer to **NullableString** |  | [optional] 
**DxAuth** | Pointer to **NullableBool** |  | [optional] 
**CurrentStep** | Pointer to **NullableString** |  | [optional] 
**InfrastructureHosting** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserSignUpResponse

`func NewUserSignUpResponse(id string, ) *UserSignUpResponse`

NewUserSignUpResponse instantiates a new UserSignUpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSignUpResponseWithDefaults

`func NewUserSignUpResponseWithDefaults() *UserSignUpResponse`

NewUserSignUpResponseWithDefaults instantiates a new UserSignUpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSignUpResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSignUpResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSignUpResponse) SetId(v string)`

SetId sets Id field to given value.


### GetFirstName

`func (o *UserSignUpResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserSignUpResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserSignUpResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserSignUpResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UserSignUpResponse) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UserSignUpResponse) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UserSignUpResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserSignUpResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserSignUpResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserSignUpResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UserSignUpResponse) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UserSignUpResponse) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetUserEmail

`func (o *UserSignUpResponse) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *UserSignUpResponse) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *UserSignUpResponse) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *UserSignUpResponse) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *UserSignUpResponse) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *UserSignUpResponse) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetTypeOfUse

`func (o *UserSignUpResponse) GetTypeOfUse() string`

GetTypeOfUse returns the TypeOfUse field if non-nil, zero value otherwise.

### GetTypeOfUseOk

`func (o *UserSignUpResponse) GetTypeOfUseOk() (*string, bool)`

GetTypeOfUseOk returns a tuple with the TypeOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfUse

`func (o *UserSignUpResponse) SetTypeOfUse(v string)`

SetTypeOfUse sets TypeOfUse field to given value.

### HasTypeOfUse

`func (o *UserSignUpResponse) HasTypeOfUse() bool`

HasTypeOfUse returns a boolean if a field has been set.

### SetTypeOfUseNil

`func (o *UserSignUpResponse) SetTypeOfUseNil(b bool)`

 SetTypeOfUseNil sets the value for TypeOfUse to be an explicit nil

### UnsetTypeOfUse
`func (o *UserSignUpResponse) UnsetTypeOfUse()`

UnsetTypeOfUse ensures that no value is present for TypeOfUse, not even an explicit nil
### GetCompanyName

`func (o *UserSignUpResponse) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UserSignUpResponse) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UserSignUpResponse) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UserSignUpResponse) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *UserSignUpResponse) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *UserSignUpResponse) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetCompanySize

`func (o *UserSignUpResponse) GetCompanySize() string`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *UserSignUpResponse) GetCompanySizeOk() (*string, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *UserSignUpResponse) SetCompanySize(v string)`

SetCompanySize sets CompanySize field to given value.

### HasCompanySize

`func (o *UserSignUpResponse) HasCompanySize() bool`

HasCompanySize returns a boolean if a field has been set.

### SetCompanySizeNil

`func (o *UserSignUpResponse) SetCompanySizeNil(b bool)`

 SetCompanySizeNil sets the value for CompanySize to be an explicit nil

### UnsetCompanySize
`func (o *UserSignUpResponse) UnsetCompanySize()`

UnsetCompanySize ensures that no value is present for CompanySize, not even an explicit nil
### GetUserRole

`func (o *UserSignUpResponse) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *UserSignUpResponse) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *UserSignUpResponse) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *UserSignUpResponse) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### SetUserRoleNil

`func (o *UserSignUpResponse) SetUserRoleNil(b bool)`

 SetUserRoleNil sets the value for UserRole to be an explicit nil

### UnsetUserRole
`func (o *UserSignUpResponse) UnsetUserRole()`

UnsetUserRole ensures that no value is present for UserRole, not even an explicit nil
### GetQoveryUsage

`func (o *UserSignUpResponse) GetQoveryUsage() string`

GetQoveryUsage returns the QoveryUsage field if non-nil, zero value otherwise.

### GetQoveryUsageOk

`func (o *UserSignUpResponse) GetQoveryUsageOk() (*string, bool)`

GetQoveryUsageOk returns a tuple with the QoveryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryUsage

`func (o *UserSignUpResponse) SetQoveryUsage(v string)`

SetQoveryUsage sets QoveryUsage field to given value.

### HasQoveryUsage

`func (o *UserSignUpResponse) HasQoveryUsage() bool`

HasQoveryUsage returns a boolean if a field has been set.

### SetQoveryUsageNil

`func (o *UserSignUpResponse) SetQoveryUsageNil(b bool)`

 SetQoveryUsageNil sets the value for QoveryUsage to be an explicit nil

### UnsetQoveryUsage
`func (o *UserSignUpResponse) UnsetQoveryUsage()`

UnsetQoveryUsage ensures that no value is present for QoveryUsage, not even an explicit nil
### GetQoveryUsageOther

`func (o *UserSignUpResponse) GetQoveryUsageOther() string`

GetQoveryUsageOther returns the QoveryUsageOther field if non-nil, zero value otherwise.

### GetQoveryUsageOtherOk

`func (o *UserSignUpResponse) GetQoveryUsageOtherOk() (*string, bool)`

GetQoveryUsageOtherOk returns a tuple with the QoveryUsageOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryUsageOther

`func (o *UserSignUpResponse) SetQoveryUsageOther(v string)`

SetQoveryUsageOther sets QoveryUsageOther field to given value.

### HasQoveryUsageOther

`func (o *UserSignUpResponse) HasQoveryUsageOther() bool`

HasQoveryUsageOther returns a boolean if a field has been set.

### SetQoveryUsageOtherNil

`func (o *UserSignUpResponse) SetQoveryUsageOtherNil(b bool)`

 SetQoveryUsageOtherNil sets the value for QoveryUsageOther to be an explicit nil

### UnsetQoveryUsageOther
`func (o *UserSignUpResponse) UnsetQoveryUsageOther()`

UnsetQoveryUsageOther ensures that no value is present for QoveryUsageOther, not even an explicit nil
### GetUserQuestions

`func (o *UserSignUpResponse) GetUserQuestions() string`

GetUserQuestions returns the UserQuestions field if non-nil, zero value otherwise.

### GetUserQuestionsOk

`func (o *UserSignUpResponse) GetUserQuestionsOk() (*string, bool)`

GetUserQuestionsOk returns a tuple with the UserQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuestions

`func (o *UserSignUpResponse) SetUserQuestions(v string)`

SetUserQuestions sets UserQuestions field to given value.

### HasUserQuestions

`func (o *UserSignUpResponse) HasUserQuestions() bool`

HasUserQuestions returns a boolean if a field has been set.

### SetUserQuestionsNil

`func (o *UserSignUpResponse) SetUserQuestionsNil(b bool)`

 SetUserQuestionsNil sets the value for UserQuestions to be an explicit nil

### UnsetUserQuestions
`func (o *UserSignUpResponse) UnsetUserQuestions()`

UnsetUserQuestions ensures that no value is present for UserQuestions, not even an explicit nil
### GetDxAuth

`func (o *UserSignUpResponse) GetDxAuth() bool`

GetDxAuth returns the DxAuth field if non-nil, zero value otherwise.

### GetDxAuthOk

`func (o *UserSignUpResponse) GetDxAuthOk() (*bool, bool)`

GetDxAuthOk returns a tuple with the DxAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDxAuth

`func (o *UserSignUpResponse) SetDxAuth(v bool)`

SetDxAuth sets DxAuth field to given value.

### HasDxAuth

`func (o *UserSignUpResponse) HasDxAuth() bool`

HasDxAuth returns a boolean if a field has been set.

### SetDxAuthNil

`func (o *UserSignUpResponse) SetDxAuthNil(b bool)`

 SetDxAuthNil sets the value for DxAuth to be an explicit nil

### UnsetDxAuth
`func (o *UserSignUpResponse) UnsetDxAuth()`

UnsetDxAuth ensures that no value is present for DxAuth, not even an explicit nil
### GetCurrentStep

`func (o *UserSignUpResponse) GetCurrentStep() string`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *UserSignUpResponse) GetCurrentStepOk() (*string, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *UserSignUpResponse) SetCurrentStep(v string)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *UserSignUpResponse) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### SetCurrentStepNil

`func (o *UserSignUpResponse) SetCurrentStepNil(b bool)`

 SetCurrentStepNil sets the value for CurrentStep to be an explicit nil

### UnsetCurrentStep
`func (o *UserSignUpResponse) UnsetCurrentStep()`

UnsetCurrentStep ensures that no value is present for CurrentStep, not even an explicit nil
### GetInfrastructureHosting

`func (o *UserSignUpResponse) GetInfrastructureHosting() string`

GetInfrastructureHosting returns the InfrastructureHosting field if non-nil, zero value otherwise.

### GetInfrastructureHostingOk

`func (o *UserSignUpResponse) GetInfrastructureHostingOk() (*string, bool)`

GetInfrastructureHostingOk returns a tuple with the InfrastructureHosting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureHosting

`func (o *UserSignUpResponse) SetInfrastructureHosting(v string)`

SetInfrastructureHosting sets InfrastructureHosting field to given value.

### HasInfrastructureHosting

`func (o *UserSignUpResponse) HasInfrastructureHosting() bool`

HasInfrastructureHosting returns a boolean if a field has been set.

### SetInfrastructureHostingNil

`func (o *UserSignUpResponse) SetInfrastructureHostingNil(b bool)`

 SetInfrastructureHostingNil sets the value for InfrastructureHosting to be an explicit nil

### UnsetInfrastructureHosting
`func (o *UserSignUpResponse) UnsetInfrastructureHosting()`

UnsetInfrastructureHosting ensures that no value is present for InfrastructureHosting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


