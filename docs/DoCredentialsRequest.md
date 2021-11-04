# DoCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Token** | Pointer to **string** |  | [optional] 
**SpacesAccessId** | Pointer to **string** |  | [optional] 
**SpacesSecretKey** | Pointer to **string** |  | [optional] 

## Methods

### NewDoCredentialsRequest

`func NewDoCredentialsRequest(name string, ) *DoCredentialsRequest`

NewDoCredentialsRequest instantiates a new DoCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoCredentialsRequestWithDefaults

`func NewDoCredentialsRequestWithDefaults() *DoCredentialsRequest`

NewDoCredentialsRequestWithDefaults instantiates a new DoCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DoCredentialsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DoCredentialsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DoCredentialsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *DoCredentialsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DoCredentialsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DoCredentialsRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DoCredentialsRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetSpacesAccessId

`func (o *DoCredentialsRequest) GetSpacesAccessId() string`

GetSpacesAccessId returns the SpacesAccessId field if non-nil, zero value otherwise.

### GetSpacesAccessIdOk

`func (o *DoCredentialsRequest) GetSpacesAccessIdOk() (*string, bool)`

GetSpacesAccessIdOk returns a tuple with the SpacesAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacesAccessId

`func (o *DoCredentialsRequest) SetSpacesAccessId(v string)`

SetSpacesAccessId sets SpacesAccessId field to given value.

### HasSpacesAccessId

`func (o *DoCredentialsRequest) HasSpacesAccessId() bool`

HasSpacesAccessId returns a boolean if a field has been set.

### GetSpacesSecretKey

`func (o *DoCredentialsRequest) GetSpacesSecretKey() string`

GetSpacesSecretKey returns the SpacesSecretKey field if non-nil, zero value otherwise.

### GetSpacesSecretKeyOk

`func (o *DoCredentialsRequest) GetSpacesSecretKeyOk() (*string, bool)`

GetSpacesSecretKeyOk returns a tuple with the SpacesSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacesSecretKey

`func (o *DoCredentialsRequest) SetSpacesSecretKey(v string)`

SetSpacesSecretKey sets SpacesSecretKey field to given value.

### HasSpacesSecretKey

`func (o *DoCredentialsRequest) HasSpacesSecretKey() bool`

HasSpacesSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


