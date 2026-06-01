# GcpServiceAccountKeyCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CredentialType** | Pointer to **string** | Optional explicit credential type. | [optional] 
**GcpCredentials** | **string** | The json must be base64 encoded | 

## Methods

### NewGcpServiceAccountKeyCredentialsRequest

`func NewGcpServiceAccountKeyCredentialsRequest(name string, gcpCredentials string, ) *GcpServiceAccountKeyCredentialsRequest`

NewGcpServiceAccountKeyCredentialsRequest instantiates a new GcpServiceAccountKeyCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpServiceAccountKeyCredentialsRequestWithDefaults

`func NewGcpServiceAccountKeyCredentialsRequestWithDefaults() *GcpServiceAccountKeyCredentialsRequest`

NewGcpServiceAccountKeyCredentialsRequestWithDefaults instantiates a new GcpServiceAccountKeyCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GcpServiceAccountKeyCredentialsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpServiceAccountKeyCredentialsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpServiceAccountKeyCredentialsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCredentialType

`func (o *GcpServiceAccountKeyCredentialsRequest) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *GcpServiceAccountKeyCredentialsRequest) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *GcpServiceAccountKeyCredentialsRequest) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *GcpServiceAccountKeyCredentialsRequest) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetGcpCredentials

`func (o *GcpServiceAccountKeyCredentialsRequest) GetGcpCredentials() string`

GetGcpCredentials returns the GcpCredentials field if non-nil, zero value otherwise.

### GetGcpCredentialsOk

`func (o *GcpServiceAccountKeyCredentialsRequest) GetGcpCredentialsOk() (*string, bool)`

GetGcpCredentialsOk returns a tuple with the GcpCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpCredentials

`func (o *GcpServiceAccountKeyCredentialsRequest) SetGcpCredentials(v string)`

SetGcpCredentials sets GcpCredentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


