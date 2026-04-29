# ArgoCdLinkedClusterDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArgocdClusterUrl** | **string** | ArgoCD destination cluster URL | 
**ArgocdClusterName** | **string** | ArgoCD cluster name (currently same as URL; will use ArgoCD alias in future) | 
**QoveryClusterId** | **string** | ID of the Qovery cluster this destination is mapped to | 
**QoveryClusterName** | **string** | Display name of the mapped Qovery cluster | 
**QoveryClusterCloudProvider** | [**CloudVendorEnum**](CloudVendorEnum.md) |  | 
**QoveryClusterType** | [**KubernetesEnum**](KubernetesEnum.md) |  | [default to KUBERNETESENUM_MANAGED]
**ApplicationsCount** | **int32** | Number of ArgoCD applications targeting this destination | 

## Methods

### NewArgoCdLinkedClusterDetails

`func NewArgoCdLinkedClusterDetails(argocdClusterUrl string, argocdClusterName string, qoveryClusterId string, qoveryClusterName string, qoveryClusterCloudProvider CloudVendorEnum, qoveryClusterType KubernetesEnum, applicationsCount int32, ) *ArgoCdLinkedClusterDetails`

NewArgoCdLinkedClusterDetails instantiates a new ArgoCdLinkedClusterDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoCdLinkedClusterDetailsWithDefaults

`func NewArgoCdLinkedClusterDetailsWithDefaults() *ArgoCdLinkedClusterDetails`

NewArgoCdLinkedClusterDetailsWithDefaults instantiates a new ArgoCdLinkedClusterDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgocdClusterUrl

`func (o *ArgoCdLinkedClusterDetails) GetArgocdClusterUrl() string`

GetArgocdClusterUrl returns the ArgocdClusterUrl field if non-nil, zero value otherwise.

### GetArgocdClusterUrlOk

`func (o *ArgoCdLinkedClusterDetails) GetArgocdClusterUrlOk() (*string, bool)`

GetArgocdClusterUrlOk returns a tuple with the ArgocdClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgocdClusterUrl

`func (o *ArgoCdLinkedClusterDetails) SetArgocdClusterUrl(v string)`

SetArgocdClusterUrl sets ArgocdClusterUrl field to given value.


### GetArgocdClusterName

`func (o *ArgoCdLinkedClusterDetails) GetArgocdClusterName() string`

GetArgocdClusterName returns the ArgocdClusterName field if non-nil, zero value otherwise.

### GetArgocdClusterNameOk

`func (o *ArgoCdLinkedClusterDetails) GetArgocdClusterNameOk() (*string, bool)`

GetArgocdClusterNameOk returns a tuple with the ArgocdClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgocdClusterName

`func (o *ArgoCdLinkedClusterDetails) SetArgocdClusterName(v string)`

SetArgocdClusterName sets ArgocdClusterName field to given value.


### GetQoveryClusterId

`func (o *ArgoCdLinkedClusterDetails) GetQoveryClusterId() string`

GetQoveryClusterId returns the QoveryClusterId field if non-nil, zero value otherwise.

### GetQoveryClusterIdOk

`func (o *ArgoCdLinkedClusterDetails) GetQoveryClusterIdOk() (*string, bool)`

GetQoveryClusterIdOk returns a tuple with the QoveryClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryClusterId

`func (o *ArgoCdLinkedClusterDetails) SetQoveryClusterId(v string)`

SetQoveryClusterId sets QoveryClusterId field to given value.


### GetQoveryClusterName

`func (o *ArgoCdLinkedClusterDetails) GetQoveryClusterName() string`

GetQoveryClusterName returns the QoveryClusterName field if non-nil, zero value otherwise.

### GetQoveryClusterNameOk

`func (o *ArgoCdLinkedClusterDetails) GetQoveryClusterNameOk() (*string, bool)`

GetQoveryClusterNameOk returns a tuple with the QoveryClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryClusterName

`func (o *ArgoCdLinkedClusterDetails) SetQoveryClusterName(v string)`

SetQoveryClusterName sets QoveryClusterName field to given value.


### GetQoveryClusterCloudProvider

`func (o *ArgoCdLinkedClusterDetails) GetQoveryClusterCloudProvider() CloudVendorEnum`

GetQoveryClusterCloudProvider returns the QoveryClusterCloudProvider field if non-nil, zero value otherwise.

### GetQoveryClusterCloudProviderOk

`func (o *ArgoCdLinkedClusterDetails) GetQoveryClusterCloudProviderOk() (*CloudVendorEnum, bool)`

GetQoveryClusterCloudProviderOk returns a tuple with the QoveryClusterCloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryClusterCloudProvider

`func (o *ArgoCdLinkedClusterDetails) SetQoveryClusterCloudProvider(v CloudVendorEnum)`

SetQoveryClusterCloudProvider sets QoveryClusterCloudProvider field to given value.


### GetQoveryClusterType

`func (o *ArgoCdLinkedClusterDetails) GetQoveryClusterType() KubernetesEnum`

GetQoveryClusterType returns the QoveryClusterType field if non-nil, zero value otherwise.

### GetQoveryClusterTypeOk

`func (o *ArgoCdLinkedClusterDetails) GetQoveryClusterTypeOk() (*KubernetesEnum, bool)`

GetQoveryClusterTypeOk returns a tuple with the QoveryClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryClusterType

`func (o *ArgoCdLinkedClusterDetails) SetQoveryClusterType(v KubernetesEnum)`

SetQoveryClusterType sets QoveryClusterType field to given value.


### GetApplicationsCount

`func (o *ArgoCdLinkedClusterDetails) GetApplicationsCount() int32`

GetApplicationsCount returns the ApplicationsCount field if non-nil, zero value otherwise.

### GetApplicationsCountOk

`func (o *ArgoCdLinkedClusterDetails) GetApplicationsCountOk() (*int32, bool)`

GetApplicationsCountOk returns a tuple with the ApplicationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsCount

`func (o *ArgoCdLinkedClusterDetails) SetApplicationsCount(v int32)`

SetApplicationsCount sets ApplicationsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


