# GcpWorkloadIdentityFederationClusterCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ObjectType** | **string** |  | 
**ProjectId** | **string** |  | 
**ServiceAccountEmail** | **string** |  | 
**WorkloadIdentityProjectNumber** | **string** |  | 
**WorkloadIdentityPoolId** | **string** |  | 
**WorkloadIdentityProviderId** | **string** |  | 
**TokenLifetimeSeconds** | **int32** |  | 

## Methods

### NewGcpWorkloadIdentityFederationClusterCredentials

`func NewGcpWorkloadIdentityFederationClusterCredentials(id string, name string, objectType string, projectId string, serviceAccountEmail string, workloadIdentityProjectNumber string, workloadIdentityPoolId string, workloadIdentityProviderId string, tokenLifetimeSeconds int32, ) *GcpWorkloadIdentityFederationClusterCredentials`

NewGcpWorkloadIdentityFederationClusterCredentials instantiates a new GcpWorkloadIdentityFederationClusterCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpWorkloadIdentityFederationClusterCredentialsWithDefaults

`func NewGcpWorkloadIdentityFederationClusterCredentialsWithDefaults() *GcpWorkloadIdentityFederationClusterCredentials`

NewGcpWorkloadIdentityFederationClusterCredentialsWithDefaults instantiates a new GcpWorkloadIdentityFederationClusterCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GcpWorkloadIdentityFederationClusterCredentials) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpWorkloadIdentityFederationClusterCredentials) SetName(v string)`

SetName sets Name field to given value.


### GetObjectType

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *GcpWorkloadIdentityFederationClusterCredentials) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProjectId

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GcpWorkloadIdentityFederationClusterCredentials) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetServiceAccountEmail

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetServiceAccountEmail() string`

GetServiceAccountEmail returns the ServiceAccountEmail field if non-nil, zero value otherwise.

### GetServiceAccountEmailOk

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetServiceAccountEmailOk() (*string, bool)`

GetServiceAccountEmailOk returns a tuple with the ServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountEmail

`func (o *GcpWorkloadIdentityFederationClusterCredentials) SetServiceAccountEmail(v string)`

SetServiceAccountEmail sets ServiceAccountEmail field to given value.


### GetWorkloadIdentityProjectNumber

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetWorkloadIdentityProjectNumber() string`

GetWorkloadIdentityProjectNumber returns the WorkloadIdentityProjectNumber field if non-nil, zero value otherwise.

### GetWorkloadIdentityProjectNumberOk

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetWorkloadIdentityProjectNumberOk() (*string, bool)`

GetWorkloadIdentityProjectNumberOk returns a tuple with the WorkloadIdentityProjectNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadIdentityProjectNumber

`func (o *GcpWorkloadIdentityFederationClusterCredentials) SetWorkloadIdentityProjectNumber(v string)`

SetWorkloadIdentityProjectNumber sets WorkloadIdentityProjectNumber field to given value.


### GetWorkloadIdentityPoolId

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetWorkloadIdentityPoolId() string`

GetWorkloadIdentityPoolId returns the WorkloadIdentityPoolId field if non-nil, zero value otherwise.

### GetWorkloadIdentityPoolIdOk

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetWorkloadIdentityPoolIdOk() (*string, bool)`

GetWorkloadIdentityPoolIdOk returns a tuple with the WorkloadIdentityPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadIdentityPoolId

`func (o *GcpWorkloadIdentityFederationClusterCredentials) SetWorkloadIdentityPoolId(v string)`

SetWorkloadIdentityPoolId sets WorkloadIdentityPoolId field to given value.


### GetWorkloadIdentityProviderId

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetWorkloadIdentityProviderId() string`

GetWorkloadIdentityProviderId returns the WorkloadIdentityProviderId field if non-nil, zero value otherwise.

### GetWorkloadIdentityProviderIdOk

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetWorkloadIdentityProviderIdOk() (*string, bool)`

GetWorkloadIdentityProviderIdOk returns a tuple with the WorkloadIdentityProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadIdentityProviderId

`func (o *GcpWorkloadIdentityFederationClusterCredentials) SetWorkloadIdentityProviderId(v string)`

SetWorkloadIdentityProviderId sets WorkloadIdentityProviderId field to given value.


### GetTokenLifetimeSeconds

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetTokenLifetimeSeconds() int32`

GetTokenLifetimeSeconds returns the TokenLifetimeSeconds field if non-nil, zero value otherwise.

### GetTokenLifetimeSecondsOk

`func (o *GcpWorkloadIdentityFederationClusterCredentials) GetTokenLifetimeSecondsOk() (*int32, bool)`

GetTokenLifetimeSecondsOk returns a tuple with the TokenLifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLifetimeSeconds

`func (o *GcpWorkloadIdentityFederationClusterCredentials) SetTokenLifetimeSeconds(v int32)`

SetTokenLifetimeSeconds sets TokenLifetimeSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


