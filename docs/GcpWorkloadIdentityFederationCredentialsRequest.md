# GcpWorkloadIdentityFederationCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CredentialType** | Pointer to **string** | Optional explicit credential type. | [optional] 
**ServiceAccountEmail** | **string** | GCP service account email to impersonate. | 
**WorkloadIdentityProviderResource** | **string** | Full Workload Identity Provider resource. | 

## Methods

### NewGcpWorkloadIdentityFederationCredentialsRequest

`func NewGcpWorkloadIdentityFederationCredentialsRequest(name string, serviceAccountEmail string, workloadIdentityProviderResource string, ) *GcpWorkloadIdentityFederationCredentialsRequest`

NewGcpWorkloadIdentityFederationCredentialsRequest instantiates a new GcpWorkloadIdentityFederationCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpWorkloadIdentityFederationCredentialsRequestWithDefaults

`func NewGcpWorkloadIdentityFederationCredentialsRequestWithDefaults() *GcpWorkloadIdentityFederationCredentialsRequest`

NewGcpWorkloadIdentityFederationCredentialsRequestWithDefaults instantiates a new GcpWorkloadIdentityFederationCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GcpWorkloadIdentityFederationCredentialsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpWorkloadIdentityFederationCredentialsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpWorkloadIdentityFederationCredentialsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCredentialType

`func (o *GcpWorkloadIdentityFederationCredentialsRequest) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *GcpWorkloadIdentityFederationCredentialsRequest) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *GcpWorkloadIdentityFederationCredentialsRequest) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *GcpWorkloadIdentityFederationCredentialsRequest) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetServiceAccountEmail

`func (o *GcpWorkloadIdentityFederationCredentialsRequest) GetServiceAccountEmail() string`

GetServiceAccountEmail returns the ServiceAccountEmail field if non-nil, zero value otherwise.

### GetServiceAccountEmailOk

`func (o *GcpWorkloadIdentityFederationCredentialsRequest) GetServiceAccountEmailOk() (*string, bool)`

GetServiceAccountEmailOk returns a tuple with the ServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountEmail

`func (o *GcpWorkloadIdentityFederationCredentialsRequest) SetServiceAccountEmail(v string)`

SetServiceAccountEmail sets ServiceAccountEmail field to given value.


### GetWorkloadIdentityProviderResource

`func (o *GcpWorkloadIdentityFederationCredentialsRequest) GetWorkloadIdentityProviderResource() string`

GetWorkloadIdentityProviderResource returns the WorkloadIdentityProviderResource field if non-nil, zero value otherwise.

### GetWorkloadIdentityProviderResourceOk

`func (o *GcpWorkloadIdentityFederationCredentialsRequest) GetWorkloadIdentityProviderResourceOk() (*string, bool)`

GetWorkloadIdentityProviderResourceOk returns a tuple with the WorkloadIdentityProviderResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadIdentityProviderResource

`func (o *GcpWorkloadIdentityFederationCredentialsRequest) SetWorkloadIdentityProviderResource(v string)`

SetWorkloadIdentityProviderResource sets WorkloadIdentityProviderResource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


