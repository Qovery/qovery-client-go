# ArgocdAppResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** | name is case insensitive | 
**ServiceType** | [**ServiceTypeEnum**](ServiceTypeEnum.md) |  | 
**Namespace** | **string** |  | 
**EnvironmentId** | **string** |  | 
**ClusterId** | **string** |  | 
**LastSyncedAt** | Pointer to **NullableTime** |  | [optional] 
**ManifestRevision** | Pointer to **NullableString** |  | [optional] 
**SourceRepoUrl** | Pointer to **NullableString** |  | [optional] 
**SourceTargetRevision** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewArgocdAppResponse

`func NewArgocdAppResponse(id string, createdAt time.Time, name string, serviceType ServiceTypeEnum, namespace string, environmentId string, clusterId string, ) *ArgocdAppResponse`

NewArgocdAppResponse instantiates a new ArgocdAppResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgocdAppResponseWithDefaults

`func NewArgocdAppResponseWithDefaults() *ArgocdAppResponse`

NewArgocdAppResponseWithDefaults instantiates a new ArgocdAppResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArgocdAppResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArgocdAppResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArgocdAppResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ArgocdAppResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArgocdAppResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArgocdAppResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ArgocdAppResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArgocdAppResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArgocdAppResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ArgocdAppResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *ArgocdAppResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArgocdAppResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArgocdAppResponse) SetName(v string)`

SetName sets Name field to given value.


### GetServiceType

`func (o *ArgocdAppResponse) GetServiceType() ServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ArgocdAppResponse) GetServiceTypeOk() (*ServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ArgocdAppResponse) SetServiceType(v ServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.


### GetNamespace

`func (o *ArgocdAppResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ArgocdAppResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ArgocdAppResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetEnvironmentId

`func (o *ArgocdAppResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ArgocdAppResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ArgocdAppResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetClusterId

`func (o *ArgocdAppResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ArgocdAppResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ArgocdAppResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetLastSyncedAt

`func (o *ArgocdAppResponse) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *ArgocdAppResponse) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *ArgocdAppResponse) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *ArgocdAppResponse) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### SetLastSyncedAtNil

`func (o *ArgocdAppResponse) SetLastSyncedAtNil(b bool)`

 SetLastSyncedAtNil sets the value for LastSyncedAt to be an explicit nil

### UnsetLastSyncedAt
`func (o *ArgocdAppResponse) UnsetLastSyncedAt()`

UnsetLastSyncedAt ensures that no value is present for LastSyncedAt, not even an explicit nil
### GetManifestRevision

`func (o *ArgocdAppResponse) GetManifestRevision() string`

GetManifestRevision returns the ManifestRevision field if non-nil, zero value otherwise.

### GetManifestRevisionOk

`func (o *ArgocdAppResponse) GetManifestRevisionOk() (*string, bool)`

GetManifestRevisionOk returns a tuple with the ManifestRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestRevision

`func (o *ArgocdAppResponse) SetManifestRevision(v string)`

SetManifestRevision sets ManifestRevision field to given value.

### HasManifestRevision

`func (o *ArgocdAppResponse) HasManifestRevision() bool`

HasManifestRevision returns a boolean if a field has been set.

### SetManifestRevisionNil

`func (o *ArgocdAppResponse) SetManifestRevisionNil(b bool)`

 SetManifestRevisionNil sets the value for ManifestRevision to be an explicit nil

### UnsetManifestRevision
`func (o *ArgocdAppResponse) UnsetManifestRevision()`

UnsetManifestRevision ensures that no value is present for ManifestRevision, not even an explicit nil
### GetSourceRepoUrl

`func (o *ArgocdAppResponse) GetSourceRepoUrl() string`

GetSourceRepoUrl returns the SourceRepoUrl field if non-nil, zero value otherwise.

### GetSourceRepoUrlOk

`func (o *ArgocdAppResponse) GetSourceRepoUrlOk() (*string, bool)`

GetSourceRepoUrlOk returns a tuple with the SourceRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepoUrl

`func (o *ArgocdAppResponse) SetSourceRepoUrl(v string)`

SetSourceRepoUrl sets SourceRepoUrl field to given value.

### HasSourceRepoUrl

`func (o *ArgocdAppResponse) HasSourceRepoUrl() bool`

HasSourceRepoUrl returns a boolean if a field has been set.

### SetSourceRepoUrlNil

`func (o *ArgocdAppResponse) SetSourceRepoUrlNil(b bool)`

 SetSourceRepoUrlNil sets the value for SourceRepoUrl to be an explicit nil

### UnsetSourceRepoUrl
`func (o *ArgocdAppResponse) UnsetSourceRepoUrl()`

UnsetSourceRepoUrl ensures that no value is present for SourceRepoUrl, not even an explicit nil
### GetSourceTargetRevision

`func (o *ArgocdAppResponse) GetSourceTargetRevision() string`

GetSourceTargetRevision returns the SourceTargetRevision field if non-nil, zero value otherwise.

### GetSourceTargetRevisionOk

`func (o *ArgocdAppResponse) GetSourceTargetRevisionOk() (*string, bool)`

GetSourceTargetRevisionOk returns a tuple with the SourceTargetRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTargetRevision

`func (o *ArgocdAppResponse) SetSourceTargetRevision(v string)`

SetSourceTargetRevision sets SourceTargetRevision field to given value.

### HasSourceTargetRevision

`func (o *ArgocdAppResponse) HasSourceTargetRevision() bool`

HasSourceTargetRevision returns a boolean if a field has been set.

### SetSourceTargetRevisionNil

`func (o *ArgocdAppResponse) SetSourceTargetRevisionNil(b bool)`

 SetSourceTargetRevisionNil sets the value for SourceTargetRevision to be an explicit nil

### UnsetSourceTargetRevision
`func (o *ArgocdAppResponse) UnsetSourceTargetRevision()`

UnsetSourceTargetRevision ensures that no value is present for SourceTargetRevision, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


