# Link

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**InternalPort** | Pointer to **int32** | The port from which the service is reachable from within the cluster | [optional] 
**ExternalPort** | Pointer to **int32** | The port from which the service is reachable from externally (i.e: 443 for HTTPS) | [optional] 
**IsQoveryDomain** | Pointer to **bool** | True if the domain is managed by Qovery, false if it belongs to the user | [optional] 
**IsDefault** | Pointer to **bool** | Indicate if the link is using the root of the domain and not one derivated from port i.e: p8080.zxxxx.jvm.worl      &#x3D;&gt; is_default &#x3D; false, is_qovery &#x3D; true zxxxx.jvm.world           &#x3D;&gt; is_default &#x3D; true, is_qovery &#x3D; true p8080-my-super-domain.com &#x3D;&gt; is_default &#x3D; false, is_qovery &#x3D; false my-super-domain.com       &#x3D;&gt; is_default &#x3D; true, is_qovery &#x3D; false  | [optional] 

## Methods

### NewLink

`func NewLink() *Link`

NewLink instantiates a new Link object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkWithDefaults

`func NewLinkWithDefaults() *Link`

NewLinkWithDefaults instantiates a new Link object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Link) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Link) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Link) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Link) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetInternalPort

`func (o *Link) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *Link) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *Link) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.

### HasInternalPort

`func (o *Link) HasInternalPort() bool`

HasInternalPort returns a boolean if a field has been set.

### GetExternalPort

`func (o *Link) GetExternalPort() int32`

GetExternalPort returns the ExternalPort field if non-nil, zero value otherwise.

### GetExternalPortOk

`func (o *Link) GetExternalPortOk() (*int32, bool)`

GetExternalPortOk returns a tuple with the ExternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPort

`func (o *Link) SetExternalPort(v int32)`

SetExternalPort sets ExternalPort field to given value.

### HasExternalPort

`func (o *Link) HasExternalPort() bool`

HasExternalPort returns a boolean if a field has been set.

### GetIsQoveryDomain

`func (o *Link) GetIsQoveryDomain() bool`

GetIsQoveryDomain returns the IsQoveryDomain field if non-nil, zero value otherwise.

### GetIsQoveryDomainOk

`func (o *Link) GetIsQoveryDomainOk() (*bool, bool)`

GetIsQoveryDomainOk returns a tuple with the IsQoveryDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQoveryDomain

`func (o *Link) SetIsQoveryDomain(v bool)`

SetIsQoveryDomain sets IsQoveryDomain field to given value.

### HasIsQoveryDomain

`func (o *Link) HasIsQoveryDomain() bool`

HasIsQoveryDomain returns a boolean if a field has been set.

### GetIsDefault

`func (o *Link) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Link) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Link) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Link) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


