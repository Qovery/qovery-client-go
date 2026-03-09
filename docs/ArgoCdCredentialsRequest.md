# ArgoCdCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArgocdUrl** | **string** | The URL of the ArgoCD instance (e.g. https://argocd.example.com) | 
**ArgocdToken** | **string** | ArgoCD API authentication token | 

## Methods

### NewArgoCdCredentialsRequest

`func NewArgoCdCredentialsRequest(argocdUrl string, argocdToken string, ) *ArgoCdCredentialsRequest`

NewArgoCdCredentialsRequest instantiates a new ArgoCdCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoCdCredentialsRequestWithDefaults

`func NewArgoCdCredentialsRequestWithDefaults() *ArgoCdCredentialsRequest`

NewArgoCdCredentialsRequestWithDefaults instantiates a new ArgoCdCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgocdUrl

`func (o *ArgoCdCredentialsRequest) GetArgocdUrl() string`

GetArgocdUrl returns the ArgocdUrl field if non-nil, zero value otherwise.

### GetArgocdUrlOk

`func (o *ArgoCdCredentialsRequest) GetArgocdUrlOk() (*string, bool)`

GetArgocdUrlOk returns a tuple with the ArgocdUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgocdUrl

`func (o *ArgoCdCredentialsRequest) SetArgocdUrl(v string)`

SetArgocdUrl sets ArgocdUrl field to given value.


### GetArgocdToken

`func (o *ArgoCdCredentialsRequest) GetArgocdToken() string`

GetArgocdToken returns the ArgocdToken field if non-nil, zero value otherwise.

### GetArgocdTokenOk

`func (o *ArgoCdCredentialsRequest) GetArgocdTokenOk() (*string, bool)`

GetArgocdTokenOk returns a tuple with the ArgocdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgocdToken

`func (o *ArgoCdCredentialsRequest) SetArgocdToken(v string)`

SetArgocdToken sets ArgocdToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


