# ArgoCdUnlinkedClusterDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArgocdClusterUrl** | **string** | ArgoCD destination cluster URL for this stale mapping | 
**ArgocdClusterName** | **string** | ArgoCD cluster name (currently same as URL; will use ArgoCD alias in future) | 
**ApplicationsCount** | **int32** | Number of ArgoCD applications targeting this destination | 

## Methods

### NewArgoCdUnlinkedClusterDetails

`func NewArgoCdUnlinkedClusterDetails(argocdClusterUrl string, argocdClusterName string, applicationsCount int32, ) *ArgoCdUnlinkedClusterDetails`

NewArgoCdUnlinkedClusterDetails instantiates a new ArgoCdUnlinkedClusterDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoCdUnlinkedClusterDetailsWithDefaults

`func NewArgoCdUnlinkedClusterDetailsWithDefaults() *ArgoCdUnlinkedClusterDetails`

NewArgoCdUnlinkedClusterDetailsWithDefaults instantiates a new ArgoCdUnlinkedClusterDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgocdClusterUrl

`func (o *ArgoCdUnlinkedClusterDetails) GetArgocdClusterUrl() string`

GetArgocdClusterUrl returns the ArgocdClusterUrl field if non-nil, zero value otherwise.

### GetArgocdClusterUrlOk

`func (o *ArgoCdUnlinkedClusterDetails) GetArgocdClusterUrlOk() (*string, bool)`

GetArgocdClusterUrlOk returns a tuple with the ArgocdClusterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgocdClusterUrl

`func (o *ArgoCdUnlinkedClusterDetails) SetArgocdClusterUrl(v string)`

SetArgocdClusterUrl sets ArgocdClusterUrl field to given value.


### GetArgocdClusterName

`func (o *ArgoCdUnlinkedClusterDetails) GetArgocdClusterName() string`

GetArgocdClusterName returns the ArgocdClusterName field if non-nil, zero value otherwise.

### GetArgocdClusterNameOk

`func (o *ArgoCdUnlinkedClusterDetails) GetArgocdClusterNameOk() (*string, bool)`

GetArgocdClusterNameOk returns a tuple with the ArgocdClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgocdClusterName

`func (o *ArgoCdUnlinkedClusterDetails) SetArgocdClusterName(v string)`

SetArgocdClusterName sets ArgocdClusterName field to given value.


### GetApplicationsCount

`func (o *ArgoCdUnlinkedClusterDetails) GetApplicationsCount() int32`

GetApplicationsCount returns the ApplicationsCount field if non-nil, zero value otherwise.

### GetApplicationsCountOk

`func (o *ArgoCdUnlinkedClusterDetails) GetApplicationsCountOk() (*int32, bool)`

GetApplicationsCountOk returns a tuple with the ApplicationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsCount

`func (o *ArgoCdUnlinkedClusterDetails) SetApplicationsCount(v int32)`

SetApplicationsCount sets ApplicationsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


