# CredentialCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to [**CloudProviderEnum**](CloudProviderEnum.md) |  | [optional] 

## Methods

### NewCredentialCluster

`func NewCredentialCluster() *CredentialCluster`

NewCredentialCluster instantiates a new CredentialCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialClusterWithDefaults

`func NewCredentialClusterWithDefaults() *CredentialCluster`

NewCredentialClusterWithDefaults instantiates a new CredentialCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CredentialCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialCluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CredentialCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CredentialCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCloudProvider

`func (o *CredentialCluster) GetCloudProvider() CloudProviderEnum`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CredentialCluster) GetCloudProviderOk() (*CloudProviderEnum, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CredentialCluster) SetCloudProvider(v CloudProviderEnum)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *CredentialCluster) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


