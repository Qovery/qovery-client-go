# GcpCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**JsonCredentials** | **string** |  | 

## Methods

### NewGcpCredentialsRequest

`func NewGcpCredentialsRequest(name string, jsonCredentials string, ) *GcpCredentialsRequest`

NewGcpCredentialsRequest instantiates a new GcpCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpCredentialsRequestWithDefaults

`func NewGcpCredentialsRequestWithDefaults() *GcpCredentialsRequest`

NewGcpCredentialsRequestWithDefaults instantiates a new GcpCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GcpCredentialsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpCredentialsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpCredentialsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetJsonCredentials

`func (o *GcpCredentialsRequest) GetJsonCredentials() string`

GetJsonCredentials returns the JsonCredentials field if non-nil, zero value otherwise.

### GetJsonCredentialsOk

`func (o *GcpCredentialsRequest) GetJsonCredentialsOk() (*string, bool)`

GetJsonCredentialsOk returns a tuple with the JsonCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonCredentials

`func (o *GcpCredentialsRequest) SetJsonCredentials(v string)`

SetJsonCredentials sets JsonCredentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


