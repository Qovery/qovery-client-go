# InlineObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**UserEmail** | **string** |  | 
**TypeOfUse** | **string** |  | 
**QoveryUsage** | **string** |  | 
**CompanyName** | Pointer to **NullableString** |  | [optional] 
**CompanySize** | Pointer to **string** |  | [optional] 
**UserRole** | Pointer to **NullableString** |  | [optional] 
**QoveryUsageOther** | Pointer to **NullableString** |  | [optional] 
**UserQuestions** | Pointer to **NullableString** |  | [optional] 
**CurrentStep** | Pointer to **NullableString** |  | [optional] 
**DxAuth** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewInlineObject2

`func NewInlineObject2(firstName string, lastName string, userEmail string, typeOfUse string, qoveryUsage string, ) *InlineObject2`

NewInlineObject2 instantiates a new InlineObject2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject2WithDefaults

`func NewInlineObject2WithDefaults() *InlineObject2`

NewInlineObject2WithDefaults instantiates a new InlineObject2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *InlineObject2) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InlineObject2) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InlineObject2) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *InlineObject2) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InlineObject2) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InlineObject2) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetUserEmail

`func (o *InlineObject2) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *InlineObject2) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *InlineObject2) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetTypeOfUse

`func (o *InlineObject2) GetTypeOfUse() string`

GetTypeOfUse returns the TypeOfUse field if non-nil, zero value otherwise.

### GetTypeOfUseOk

`func (o *InlineObject2) GetTypeOfUseOk() (*string, bool)`

GetTypeOfUseOk returns a tuple with the TypeOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfUse

`func (o *InlineObject2) SetTypeOfUse(v string)`

SetTypeOfUse sets TypeOfUse field to given value.


### GetQoveryUsage

`func (o *InlineObject2) GetQoveryUsage() string`

GetQoveryUsage returns the QoveryUsage field if non-nil, zero value otherwise.

### GetQoveryUsageOk

`func (o *InlineObject2) GetQoveryUsageOk() (*string, bool)`

GetQoveryUsageOk returns a tuple with the QoveryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryUsage

`func (o *InlineObject2) SetQoveryUsage(v string)`

SetQoveryUsage sets QoveryUsage field to given value.


### GetCompanyName

`func (o *InlineObject2) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InlineObject2) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InlineObject2) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InlineObject2) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *InlineObject2) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *InlineObject2) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetCompanySize

`func (o *InlineObject2) GetCompanySize() string`

GetCompanySize returns the CompanySize field if non-nil, zero value otherwise.

### GetCompanySizeOk

`func (o *InlineObject2) GetCompanySizeOk() (*string, bool)`

GetCompanySizeOk returns a tuple with the CompanySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySize

`func (o *InlineObject2) SetCompanySize(v string)`

SetCompanySize sets CompanySize field to given value.

### HasCompanySize

`func (o *InlineObject2) HasCompanySize() bool`

HasCompanySize returns a boolean if a field has been set.

### GetUserRole

`func (o *InlineObject2) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *InlineObject2) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *InlineObject2) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *InlineObject2) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### SetUserRoleNil

`func (o *InlineObject2) SetUserRoleNil(b bool)`

 SetUserRoleNil sets the value for UserRole to be an explicit nil

### UnsetUserRole
`func (o *InlineObject2) UnsetUserRole()`

UnsetUserRole ensures that no value is present for UserRole, not even an explicit nil
### GetQoveryUsageOther

`func (o *InlineObject2) GetQoveryUsageOther() string`

GetQoveryUsageOther returns the QoveryUsageOther field if non-nil, zero value otherwise.

### GetQoveryUsageOtherOk

`func (o *InlineObject2) GetQoveryUsageOtherOk() (*string, bool)`

GetQoveryUsageOtherOk returns a tuple with the QoveryUsageOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryUsageOther

`func (o *InlineObject2) SetQoveryUsageOther(v string)`

SetQoveryUsageOther sets QoveryUsageOther field to given value.

### HasQoveryUsageOther

`func (o *InlineObject2) HasQoveryUsageOther() bool`

HasQoveryUsageOther returns a boolean if a field has been set.

### SetQoveryUsageOtherNil

`func (o *InlineObject2) SetQoveryUsageOtherNil(b bool)`

 SetQoveryUsageOtherNil sets the value for QoveryUsageOther to be an explicit nil

### UnsetQoveryUsageOther
`func (o *InlineObject2) UnsetQoveryUsageOther()`

UnsetQoveryUsageOther ensures that no value is present for QoveryUsageOther, not even an explicit nil
### GetUserQuestions

`func (o *InlineObject2) GetUserQuestions() string`

GetUserQuestions returns the UserQuestions field if non-nil, zero value otherwise.

### GetUserQuestionsOk

`func (o *InlineObject2) GetUserQuestionsOk() (*string, bool)`

GetUserQuestionsOk returns a tuple with the UserQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuestions

`func (o *InlineObject2) SetUserQuestions(v string)`

SetUserQuestions sets UserQuestions field to given value.

### HasUserQuestions

`func (o *InlineObject2) HasUserQuestions() bool`

HasUserQuestions returns a boolean if a field has been set.

### SetUserQuestionsNil

`func (o *InlineObject2) SetUserQuestionsNil(b bool)`

 SetUserQuestionsNil sets the value for UserQuestions to be an explicit nil

### UnsetUserQuestions
`func (o *InlineObject2) UnsetUserQuestions()`

UnsetUserQuestions ensures that no value is present for UserQuestions, not even an explicit nil
### GetCurrentStep

`func (o *InlineObject2) GetCurrentStep() string`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *InlineObject2) GetCurrentStepOk() (*string, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *InlineObject2) SetCurrentStep(v string)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *InlineObject2) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### SetCurrentStepNil

`func (o *InlineObject2) SetCurrentStepNil(b bool)`

 SetCurrentStepNil sets the value for CurrentStep to be an explicit nil

### UnsetCurrentStep
`func (o *InlineObject2) UnsetCurrentStep()`

UnsetCurrentStep ensures that no value is present for CurrentStep, not even an explicit nil
### GetDxAuth

`func (o *InlineObject2) GetDxAuth() bool`

GetDxAuth returns the DxAuth field if non-nil, zero value otherwise.

### GetDxAuthOk

`func (o *InlineObject2) GetDxAuthOk() (*bool, bool)`

GetDxAuthOk returns a tuple with the DxAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDxAuth

`func (o *InlineObject2) SetDxAuth(v bool)`

SetDxAuth sets DxAuth field to given value.

### HasDxAuth

`func (o *InlineObject2) HasDxAuth() bool`

HasDxAuth returns a boolean if a field has been set.

### SetDxAuthNil

`func (o *InlineObject2) SetDxAuthNil(b bool)`

 SetDxAuthNil sets the value for DxAuth to be an explicit nil

### UnsetDxAuth
`func (o *InlineObject2) UnsetDxAuth()`

UnsetDxAuth ensures that no value is present for DxAuth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


