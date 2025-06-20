# OrganizationCrendentialsResponseListResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | Pointer to [**ClusterCredentials**](ClusterCredentials.md) |  | [optional] 
**Clusters** | Pointer to [**[]OrganizationCrendentialsResponseListResultsInnerClustersInner**](OrganizationCrendentialsResponseListResultsInnerClustersInner.md) |  | [optional] 

## Methods

### NewOrganizationCrendentialsResponseListResultsInner

`func NewOrganizationCrendentialsResponseListResultsInner() *OrganizationCrendentialsResponseListResultsInner`

NewOrganizationCrendentialsResponseListResultsInner instantiates a new OrganizationCrendentialsResponseListResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCrendentialsResponseListResultsInnerWithDefaults

`func NewOrganizationCrendentialsResponseListResultsInnerWithDefaults() *OrganizationCrendentialsResponseListResultsInner`

NewOrganizationCrendentialsResponseListResultsInnerWithDefaults instantiates a new OrganizationCrendentialsResponseListResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *OrganizationCrendentialsResponseListResultsInner) GetCredential() ClusterCredentials`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *OrganizationCrendentialsResponseListResultsInner) GetCredentialOk() (*ClusterCredentials, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *OrganizationCrendentialsResponseListResultsInner) SetCredential(v ClusterCredentials)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *OrganizationCrendentialsResponseListResultsInner) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetClusters

`func (o *OrganizationCrendentialsResponseListResultsInner) GetClusters() []OrganizationCrendentialsResponseListResultsInnerClustersInner`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *OrganizationCrendentialsResponseListResultsInner) GetClustersOk() (*[]OrganizationCrendentialsResponseListResultsInnerClustersInner, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *OrganizationCrendentialsResponseListResultsInner) SetClusters(v []OrganizationCrendentialsResponseListResultsInnerClustersInner)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *OrganizationCrendentialsResponseListResultsInner) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


