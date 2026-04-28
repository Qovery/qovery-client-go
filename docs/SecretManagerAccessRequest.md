# SecretManagerAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Endpoint** | [**SecretManagerEndpointConfigurationDto**](SecretManagerEndpointConfigurationDto.md) |  | 
**Authentication** | [**SecretManagerAuthenticationDto**](SecretManagerAuthenticationDto.md) |  | 

## Methods

### NewSecretManagerAccessRequest

`func NewSecretManagerAccessRequest(name string, endpoint SecretManagerEndpointConfigurationDto, authentication SecretManagerAuthenticationDto, ) *SecretManagerAccessRequest`

NewSecretManagerAccessRequest instantiates a new SecretManagerAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretManagerAccessRequestWithDefaults

`func NewSecretManagerAccessRequestWithDefaults() *SecretManagerAccessRequest`

NewSecretManagerAccessRequestWithDefaults instantiates a new SecretManagerAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecretManagerAccessRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretManagerAccessRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretManagerAccessRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecretManagerAccessRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SecretManagerAccessRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SecretManagerAccessRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *SecretManagerAccessRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretManagerAccessRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretManagerAccessRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEndpoint

`func (o *SecretManagerAccessRequest) GetEndpoint() SecretManagerEndpointConfigurationDto`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SecretManagerAccessRequest) GetEndpointOk() (*SecretManagerEndpointConfigurationDto, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SecretManagerAccessRequest) SetEndpoint(v SecretManagerEndpointConfigurationDto)`

SetEndpoint sets Endpoint field to given value.


### GetAuthentication

`func (o *SecretManagerAccessRequest) GetAuthentication() SecretManagerAuthenticationDto`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *SecretManagerAccessRequest) GetAuthenticationOk() (*SecretManagerAuthenticationDto, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *SecretManagerAccessRequest) SetAuthentication(v SecretManagerAuthenticationDto)`

SetAuthentication sets Authentication field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


