# AccountInfoEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunicationEmail** | Pointer to **NullableString** | The email to be used for official Qovery communications | [optional] 

## Methods

### NewAccountInfoEditRequest

`func NewAccountInfoEditRequest() *AccountInfoEditRequest`

NewAccountInfoEditRequest instantiates a new AccountInfoEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInfoEditRequestWithDefaults

`func NewAccountInfoEditRequestWithDefaults() *AccountInfoEditRequest`

NewAccountInfoEditRequestWithDefaults instantiates a new AccountInfoEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunicationEmail

`func (o *AccountInfoEditRequest) GetCommunicationEmail() string`

GetCommunicationEmail returns the CommunicationEmail field if non-nil, zero value otherwise.

### GetCommunicationEmailOk

`func (o *AccountInfoEditRequest) GetCommunicationEmailOk() (*string, bool)`

GetCommunicationEmailOk returns a tuple with the CommunicationEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationEmail

`func (o *AccountInfoEditRequest) SetCommunicationEmail(v string)`

SetCommunicationEmail sets CommunicationEmail field to given value.

### HasCommunicationEmail

`func (o *AccountInfoEditRequest) HasCommunicationEmail() bool`

HasCommunicationEmail returns a boolean if a field has been set.

### SetCommunicationEmailNil

`func (o *AccountInfoEditRequest) SetCommunicationEmailNil(b bool)`

 SetCommunicationEmailNil sets the value for CommunicationEmail to be an explicit nil

### UnsetCommunicationEmail
`func (o *AccountInfoEditRequest) UnsetCommunicationEmail()`

UnsetCommunicationEmail ensures that no value is present for CommunicationEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


