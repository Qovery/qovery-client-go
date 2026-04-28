# SecretManagerAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Endpoint** | [**SecretManagerEndpointConfigurationDto**](SecretManagerEndpointConfigurationDto.md) |  | 
**Authentication** | [**SecretManagerAuthenticationDto**](SecretManagerAuthenticationDto.md) |  | 

## Methods

### NewSecretManagerAccessResponse

`func NewSecretManagerAccessResponse(id string, name string, createdAt time.Time, updatedAt time.Time, endpoint SecretManagerEndpointConfigurationDto, authentication SecretManagerAuthenticationDto, ) *SecretManagerAccessResponse`

NewSecretManagerAccessResponse instantiates a new SecretManagerAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretManagerAccessResponseWithDefaults

`func NewSecretManagerAccessResponseWithDefaults() *SecretManagerAccessResponse`

NewSecretManagerAccessResponseWithDefaults instantiates a new SecretManagerAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecretManagerAccessResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretManagerAccessResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretManagerAccessResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SecretManagerAccessResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretManagerAccessResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretManagerAccessResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *SecretManagerAccessResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecretManagerAccessResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecretManagerAccessResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SecretManagerAccessResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SecretManagerAccessResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SecretManagerAccessResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEndpoint

`func (o *SecretManagerAccessResponse) GetEndpoint() SecretManagerEndpointConfigurationDto`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SecretManagerAccessResponse) GetEndpointOk() (*SecretManagerEndpointConfigurationDto, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SecretManagerAccessResponse) SetEndpoint(v SecretManagerEndpointConfigurationDto)`

SetEndpoint sets Endpoint field to given value.


### GetAuthentication

`func (o *SecretManagerAccessResponse) GetAuthentication() SecretManagerAuthenticationDto`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *SecretManagerAccessResponse) GetAuthenticationOk() (*SecretManagerAuthenticationDto, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *SecretManagerAccessResponse) SetAuthentication(v SecretManagerAuthenticationDto)`

SetAuthentication sets Authentication field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


